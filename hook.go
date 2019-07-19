package hook

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Default config for sls logrus hooks
const (
	BufferSize          = 4096
	DefaultSendInterval = 300 * time.Millisecond
	MaxBatchSize        = 300
)

type Config struct {
	Endpoint     string
	AccessKey    string
	AccessSecret string
	LogStore     string
	Topic        string
	Timeout      time.Duration
}

// SlsLogrusHook logrus hook for sls
type SlsLogrusHook struct {
	client       *SlsClient
	sendInterval time.Duration
	c            chan *Log
	lock         *sync.Mutex
	sending      bool
	realSendLogs func(logs []*Log) error
}

func New(config *Config) (*SlsLogrusHook, error) {
	client, err := NewSlsClient(config)
	if err != nil {
		return nil, errors.WithMessage(err, "Unable to create sls logrus hook")
	}
	if len(config.Topic) == 0 {
		return nil, errors.New("Sls topic should not be empty")
	}
	hook := &SlsLogrusHook{
		client:       client,
		c:            make(chan *Log, BufferSize),
		lock:         &sync.Mutex{},
		sending:      false,
		sendInterval: DefaultSendInterval,
	}
	err = client.Ping()
	if err != nil {
		hook.realSendLogs = fallbackSendLogs
		_, _ = fmt.Fprintf(os.Stderr, "Fail to send logs to sls, fallback to stdout. error: %v", err.Error())
	} else {
		hook.realSendLogs = client.SendLogs
	}
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		<-gracefulStop
		fmt.Println("Flushing logs")
		time.Sleep(time.Second)
	}()
	return hook, err
}

// NewSlsLogrusHook create logrus hook
func NewSlsLogrusHook(endpoint string, accessKey string, accessSecret string, logStore string, topic string) (*SlsLogrusHook, error) {
	return New(&Config{
		endpoint,
		accessKey,
		accessSecret,
		logStore,
		topic,
		DefaultTimeout,
	})
}

// SetSendInterval change batch send interval
func (hook *SlsLogrusHook) SetSendInterval(interval time.Duration) {
	hook.sendInterval = interval
}

// Fire implement logrus Hook interface
func (hook *SlsLogrusHook) Fire(entry *logrus.Entry) error {
	const depth = 16
	var pcs [depth]uintptr
	n := runtime.Callers(5, pcs[:])
	var location string
	for _, pc := range pcs[0:n] {
		if !strings.HasPrefix(getFunctionName(pc), "github.com/sirupsen/logrus") {
			file, line := getFileLocation(pc)
			location = fmt.Sprintf("%s#%d", file, line)
			break
		}
	}

	log := &Log{
		Time: proto.Uint32(uint32(time.Now().Unix())),
		Contents: []*LogContent{
			{
				Key:   proto.String("level"),
				Value: proto.String(strings.ToUpper(entry.Level.String())),
			},
			{
				Key:   proto.String("location"),
				Value: proto.String(location),
			},
			{
				Key:   proto.String("message"),
				Value: proto.String(entry.Message),
			},
		},
	}
	for k, v := range entry.Data {
		if k == "__topic__" || k == "__source__" || k == "level" || k == "message" {
			k = "field_" + k
		}
		var value string
		switch v := v.(type) {
		case string:
			value = v
		case error:
			value = fmt.Sprintf("%+v", v)
		default:
			bytes, err := json.Marshal(v)
			if err != nil {
				value = fmt.Sprintf("%+v", v)
			} else {
				value = string(bytes)
			}
		}
		if len(value) == 0 {
			continue
		}
		log.Contents = append(log.Contents, &LogContent{
			Key:   proto.String(k),
			Value: proto.String(value),
		})
	}
	hook.c <- log
	if !hook.sending {
		hook.startWork()
	}
	return nil
}

// Levels implement logrus Hook interface
func (hook *SlsLogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// Flush ensure logs are flush through sls api
func (hook *SlsLogrusHook) Flush(timeout time.Duration) {
	until := time.Now().UnixNano() + int64(timeout)
	for (hook.sending || len(hook.c) > 0) && time.Now().UnixNano() < until {
		time.Sleep(10 * time.Millisecond)
	}
}

func (hook *SlsLogrusHook) startWork() {
	hook.lock.Lock()
	defer hook.lock.Unlock()
	if hook.sending {
		return
	}
	hook.sending = true
	go hook.work()
}

func (hook *SlsLogrusHook) work() {
	for {
		if !hook.sending {
			return
		}
		deadline := time.After(hook.sendInterval)
		logs := make([]*Log, MaxBatchSize)
		count := 0
	waitLoop:
		for count < MaxBatchSize {
			select {
			case log := <-hook.c:
				logs[count] = log
				count++
			case <-deadline:
				break waitLoop
			}
		}
		if count == 0 {
			time.Sleep(hook.sendInterval)
			if len(hook.c) == 0 {
				break
			}
		}
		if err := hook.realSendLogs(logs[0:count]); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error sending logs, error: %+v\n", err)
			_ = fallbackSendLogs(logs)
		}
	}
	hook.sending = false
	// if new logs pushed to channel before setting sending to false.
	if len(hook.c) > 0 {
		hook.startWork()
	}
}

func fallbackSendLogs(logs []*Log) error {
	for _, log := range logs {
		_, _ = fmt.Fprint(os.Stdout, log, "\n")
	}
	return nil
}

func getFileLocation(f uintptr) (string, int) {
	fn := runtime.FuncForPC(f - 1)
	if fn == nil {
		return "unknown", 0
	}
	return fn.FileLine(f - 1)
}

func getFunctionName(f uintptr) string {
	fn := runtime.FuncForPC(f - 1)
	if fn == nil {
		return "unknown"
	}
	return fn.Name()
}
