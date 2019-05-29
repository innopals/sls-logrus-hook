package hook_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"github.com/gogo/protobuf/proto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"sls-logrus-hook"
	"testing"
	"time"
)

func TestLogs(t *testing.T) {
	requests := make(chan *http.Request, 3)
	contents := make(chan []byte, 3)
	mockServer := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			requests <- req
			if req.Method == "POST" {
				bytes, err := ioutil.ReadAll(req.Body)
				assert.Nil(t, err)
				contents <- bytes
			}
			writer.WriteHeader(200)
		}),
		ReadTimeout:    3 * time.Second,
		WriteTimeout:   3 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal("Fail to find a free port", err)
	}
	go func() {
		_ = mockServer.Serve(listener)
	}()
	port := listener.Addr().(*net.TCPAddr).Port
	endpoint := fmt.Sprintf("127.0.0.1:%d", port)
	slsLogrusHook, err := hook.NewSlsLogrusHook(endpoint, "test", "test", "test", "test")
	assert.Nil(t, err)
	slsLogrusHook.SetSendInterval(100 * time.Millisecond)
	defer func() {
		slsLogrusHook.Flush(time.Second)
		_ = mockServer.Close()
		_ = listener.Close()
	}()

	select {
	case req := <-requests:
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, "/logstores/test", req.RequestURI)
		date := req.Header.Get("Date")
		stringToSign := "GET\n\n\n" + date + "\nx-log-apiversion:0.6.0\nx-log-signaturemethod:hmac-sha1\n/logstores/test"
		sha1Hash := hmac.New(sha1.New, []byte("test"))
		_, e := sha1Hash.Write([]byte(stringToSign))
		assert.Nil(t, e)
		assert.Equal(t, "LOG test:"+base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil)), req.Header.Get("Authorization"))
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Mock server should have received a request.")
	}

	logrus.AddHook(slsLogrusHook)
	logrus.SetFormatter(&hook.NoopFormatter{})
	logrus.SetOutput(ioutil.Discard)

	logrus.Info("Hello world!")
	var date string
	select {
	case req := <-requests:
		assert.Equal(t, "POST", req.Method)
		assert.Equal(t, "/logstores/test/shards/lb", req.RequestURI)
		date = req.Header.Get("Date")
		md5 := req.Header.Get("Content-Md5")
		stringToSign := "POST\n" + md5 + "\napplication/x-protobuf\n" + date + "\nx-log-apiversion:0.6.0\nx-log-bodyrawsize:0\nx-log-signaturemethod:hmac-sha1\n/logstores/test/shards/lb"
		sha1Hash := hmac.New(sha1.New, []byte("test"))
		_, e := sha1Hash.Write([]byte(stringToSign))
		assert.Nil(t, e)
		assert.Equal(t, "LOG test:"+base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil)), req.Header.Get("Authorization"))
	case <-time.After(300 * time.Millisecond):
		t.Errorf("Mock server should have received a request.")
	}

	select {
	case content := <-contents:
		group := new(hook.LogGroup)
		assert.Nil(t, proto.Unmarshal(content, group))
		assert.Equal(t, 1, len(group.Logs))
		assert.Equal(t, date, time.Unix(int64(*group.Logs[0].Time), 0).UTC().Format(http.TimeFormat))
		assert.Equal(t, "__topic__", *group.Logs[0].Contents[0].Key)
		assert.Equal(t, "test", *group.Logs[0].Contents[0].Value)
		assert.Equal(t, "level", *group.Logs[0].Contents[1].Key)
		assert.Equal(t, "info", *group.Logs[0].Contents[1].Value)
		assert.Equal(t, "message", *group.Logs[0].Contents[2].Key)
		assert.Equal(t, "Hello world!", *group.Logs[0].Contents[2].Value)
	case <-time.After(300 * time.Millisecond):
		t.Errorf("Mock server should have received a request content.")
	}
}
