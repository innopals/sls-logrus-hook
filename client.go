package hook

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

// Default config for sls client
const (
	DefaultTimeout  = 5 * time.Second
	MaxLogItemSize  = 512 * 1024      // Safe value for maximum 1M log item.
	MaxLogGroupSize = 4 * 1024 * 1024 // Safe value for maximum 5M log group
	MaxLogBatchSize = 1024            // Safe value for batch send size
)

var logSource string

func init() {
	var err error
	logSource, err = os.Hostname()
	if err != nil {
		logSource = "unknown_source"
	}
}

// SlsClient the client struct for sls connection
type SlsClient struct {
	endpoint     string
	accessKey    string
	accessSecret string
	logStore     string
	topic        string
	lock         *sync.Mutex
	client       *http.Client
}

// NewSlsClient create a new sls client
func NewSlsClient(config *Config) (*SlsClient, error) {
	if len(config.Endpoint) == 0 {
		return nil, errors.New("Sls endpoint should not be empty")
	}
	if len(config.AccessKey) == 0 {
		return nil, errors.New("Sls access key should not be empty")
	}
	if len(config.AccessSecret) == 0 {
		return nil, errors.New("Sls access secret should not be empty")
	}
	if len(config.LogStore) == 0 {
		return nil, errors.New("Sls log store should not be empty")
	}
	endpoint := config.Endpoint
	if !strings.HasPrefix(endpoint, "http://") || !strings.HasPrefix(endpoint, "https://") {
		endpoint = "http://" + endpoint
	}
	if strings.HasSuffix(endpoint, "/") {
		endpoint = endpoint[:len(endpoint)-1]
	}
	return &SlsClient{
		endpoint:     endpoint,
		accessKey:    config.AccessKey,
		accessSecret: config.AccessSecret,
		logStore:     config.LogStore,
		topic:        config.Topic,
		lock:         &sync.Mutex{},
		client: &http.Client{
			Timeout: config.Timeout,
		},
	}, nil
}

// Ping sls api auth & connection
func (client *SlsClient) Ping() error {
	method := "GET"
	resource := "/logstores/" + client.logStore
	headers := make(map[string]string)

	headers[HeaderLogVersion] = SlsVersion
	headers[HeaderLogSignatureMethod] = SlsSignatureMethod
	headers[HeaderHost] = client.endpoint
	headers[HeaderDate] = time.Now().UTC().Format(http.TimeFormat)

	sign := APISign(client.accessSecret, method, headers, resource)
	headers[HeaderAuthorization] = fmt.Sprintf("LOG %s:%s", client.accessKey, sign)

	url := client.endpoint + resource

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return errors.WithMessage(err, "Error creating http request for sls")
	}
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	resp, err := client.client.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return errors.WithMessage(err, "Error sending log with http client")
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}

// SendLogs using sls api & handle extreme cases
func (client *SlsClient) SendLogs(logs []*Log) error {
	if len(logs) == 0 {
		return nil
	}
	if len(logs) > MaxLogBatchSize {
		return errors.Errorf("Log batch size should not exceed %d.", MaxLogBatchSize)
	}
	group := LogGroup{}
	group.Logs = logs
	group.Topic = proto.String(client.topic)
	group.Source = proto.String(logSource)
	body, err := proto.Marshal(&group)
	if len(body) > MaxLogGroupSize {
		// Extreme cases when log group size exceed the maximum
		return client.splitSendLogs(logs)
	}
	if err != nil {
		return err
	}
	err = client.sendPb(body)
	if err != nil {
		return err
	}
	return nil
}

func logSize(log *Log) int {
	// Estimate log size
	size := 4
	for _, content := range log.Contents {
		size += len(*content.Key) + len(*content.Value) + 8
	}
	return size
}

func (client *SlsClient) splitSendLogs(logs []*Log) error {
	var errorList []error
	cursor := 0
	for cursor < len(logs) {
		groupSize := 0
		group := LogGroup{
			Logs: make([]*Log, 0),
		}
		group.Topic = proto.String(client.topic)
		group.Source = proto.String(logSource)
		for cursor < len(logs) {
			log := logs[cursor]
			size := logSize(log)
			if groupSize+size > MaxLogGroupSize {
				break
			}
			cursor++
			if size > MaxLogItemSize {
				// Print huge single log to stdout
				_, _ = fmt.Fprintf(os.Stdout, "[HUGE SLS LOG] %+v", log)
				continue
			}
			groupSize += size
			group.Logs = append(group.Logs, log)
		}

		body, err := proto.Marshal(&group)
		if len(body) > MaxLogGroupSize {
			// Extreme cases when log group size exceed the maximum
			_, _ = fmt.Fprintf(os.Stdout, "[HUGE SLS LOG GROUP] %+v", group)
			continue
		}
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
		err = client.sendPb(body)
		if err != nil {
			errorList = append(errorList, err)
			continue
		}
	}
	if len(errorList) == 0 {
		return nil
	}
	return errors.Errorf("Fail to send logs due to the following errors: %+v", errorList)
}

func (client *SlsClient) sendPb(logContent []byte) error {
	method := "POST"
	resource := "/logstores/" + client.logStore + "/shards/lb"
	headers := make(map[string]string)
	logMD5 := md5.Sum(logContent)
	strMd5 := strings.ToUpper(fmt.Sprintf("%x", logMD5))

	headers[HeaderLogVersion] = SlsVersion
	headers[HeaderContentType] = "application/x-protobuf"
	headers[HeaderContentMd5] = strMd5
	headers[HeaderLogSignatureMethod] = SlsSignatureMethod
	headers[HeaderContentLength] = fmt.Sprintf("%v", len(logContent))
	headers[HeaderLogBodyRawSize] = "0"
	headers[HeaderHost] = client.endpoint
	headers[HeaderDate] = time.Now().UTC().Format(http.TimeFormat)
	sign := APISign(client.accessSecret, method, headers, resource)
	headers[HeaderAuthorization] = fmt.Sprintf("LOG %s:%s", client.accessKey, sign)

	url := client.endpoint + resource
	postBodyReader := bytes.NewBuffer(logContent)

	req, err := http.NewRequest(method, url, postBodyReader)
	if err != nil {
		return errors.WithMessage(err, "Error creating http request for sls")
	}
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	resp, err := client.client.Do(req)
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()
	if err != nil {
		return errors.WithMessage(err, "Error sending log with http client")
	}
	if resp.StatusCode != 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}
