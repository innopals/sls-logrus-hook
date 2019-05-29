package hook

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
)

const (
	DefaultTimeout           = 5 * time.Second
	SlsVersion               = "0.6.0"
	SlsSignatureMethod       = "hmac-sha1"
	HeaderAuthorization      = "Authorization"
	HeaderContentType        = "Content-Type"
	HeaderContentLength      = "Content-Length"
	HeaderContentMd5         = "Content-MD5"
	HeaderDate               = "Date"
	HeaderHost               = "Host"
	HeaderLogVersion         = "x-log-apiversion"
	HeaderLogSignatureMethod = "x-log-signaturemethod"
	HeaderLogBodyRawSize     = "x-log-bodyrawsize"
)

type SlsClient struct {
	endpoint     string
	accessKey    string
	accessSecret string
	logStore     string
	lock         *sync.Mutex
	client       *http.Client
}

func NewSlsClient(endpoint string, accessKey string, accessSecret string, logStore string) (*SlsClient, error) {
	if len(endpoint) == 0 {
		return nil, errors.New("Sls endpoint should not be empty")
	}
	if len(accessKey) == 0 {
		return nil, errors.New("Sls access key should not be empty")
	}
	if len(accessSecret) == 0 {
		return nil, errors.New("Sls access secret should not be empty")
	}
	if len(logStore) == 0 {
		return nil, errors.New("Sls log store should not be empty")
	}
	return &SlsClient{
		endpoint:     endpoint,
		accessKey:    accessKey,
		accessSecret: accessSecret,
		logStore:     logStore,
		lock:         &sync.Mutex{},
		client: &http.Client{
			Timeout: DefaultTimeout,
		},
	}, nil
}

func (client *SlsClient) Ping() error {
	// TODO use get log store to ping connection
	group := LogGroup{}
	group.Logs = []*Log{{
		Time: proto.Uint32(uint32(time.Now().Unix())),
		Contents: []*LogContent{{
			Key:   proto.String("__topic__"),
			Value: proto.String("status"),
		}, {
			Key:   proto.String("message"),
			Value: proto.String("Status check by sls-logrus-hook."),
		}},
	}}
	body, err := proto.Marshal(&group)
	if err != nil {
		return err
	}
	err = client.sendPb(body)
	if err != nil {
		return err
	}
	return nil
}

func (client *SlsClient) SendLogs(logs []*Log) error {
	// TODO Split logs when log size exceeds 1M or log group size exceed 5M.
	for _, log := range logs {
		fmt.Println(log)
	}
	group := LogGroup{}
	group.Logs = logs
	body, err := proto.Marshal(&group)
	if err != nil {
		return err
	}
	err = client.sendPb(body)
	if err != nil {
		return err
	}
	return nil
}

func (client *SlsClient) sendPb(logContent []byte) error {
	method := "POST"
	resource := "logstores/" + client.logStore + "/shards/lb"
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

	if sign, e := client.Sign(method, headers, fmt.Sprintf("/%s", resource)); e != nil {
		return errors.WithMessage(e, "Fail to create sign for sls")
	} else {
		headers[HeaderAuthorization] = fmt.Sprintf("LOG %s:%s", client.accessKey, sign)
	}

	url := client.endpoint + "/" + resource
	if !strings.HasPrefix(client.endpoint, "http://") || strings.HasPrefix(client.endpoint, "https://") {
		url = "http://" + client.endpoint + "/" + resource
	}
	postBodyReader := bytes.NewBuffer(logContent)

	req, err := http.NewRequest(method, url, postBodyReader)
	if err != nil {
		return errors.WithMessage(err, "Error creating http request for sls")
	}
	for header, value := range headers {
		req.Header.Add(header, value)
	}

	resp, err := client.client.Do(req)
	if err != nil {
		return errors.WithMessage(err, "Error sending log with http client")
	}
	if resp.StatusCode != 200 {
		defer func() { _ = resp.Body.Close() }()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}
	return nil
}

func (client *SlsClient) Sign(method string, headers map[string]string, resource string) (string, error) {

	signItems := make([]string, 0)
	signItems = append(signItems, method)

	var contentMD5, contentType string
	date := time.Now().UTC().Format(http.TimeFormat)

	if v, exist := headers[HeaderContentMd5]; exist {
		contentMD5 = v
	}
	if v, exist := headers[HeaderContentType]; exist {
		contentType = v
	}
	if v, exist := headers[HeaderDate]; exist {
		date = v
	}

	logHeaders := make([]string, 0)
	for k, v := range headers {
		if strings.HasPrefix(k, "x-log") || strings.HasPrefix(k, "x-acs") {
			logHeaders = append(logHeaders, k+":"+strings.TrimSpace(v))
		}
	}

	sort.Sort(sort.StringSlice(logHeaders))

	stringToSign := method + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		date + "\n" +
		strings.Join(logHeaders, "\n") + "\n" +
		resource

	sha1Hash := hmac.New(sha1.New, []byte(client.accessSecret))
	if _, e := sha1Hash.Write([]byte(stringToSign)); e != nil {
		return "", e
	}
	return base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil)), nil
}
