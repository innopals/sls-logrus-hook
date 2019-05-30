package hook_test

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
	"time"

	"github.com/gogo/protobuf/proto"
	hook "github.com/innopals/sls-logrus-hook"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type RequestWrapper struct {
	Request *http.Request
	Body    []byte
}

func TestLogs(t *testing.T) {
	requests := make(chan *RequestWrapper, 3)
	mockServer := &http.Server{
		Addr: ":8080",
		Handler: http.HandlerFunc(func(writer http.ResponseWriter, req *http.Request) {
			wrapper := &RequestWrapper{}
			wrapper.Request = req
			if req.Method == "POST" {
				bytes, err := ioutil.ReadAll(req.Body)
				assert.Nil(t, err)
				wrapper.Body = bytes
			}
			requests <- wrapper
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
	case wrapper := <-requests:
		assert.Equal(t, "GET", wrapper.Request.Method)
		assert.Equal(t, "/logstores/test", wrapper.Request.RequestURI)
		date := wrapper.Request.Header.Get("Date")
		stringToSign := "GET\n\n\n" + date + "\nx-log-apiversion:0.6.0\nx-log-signaturemethod:hmac-sha1\n/logstores/test"
		sha1Hash := hmac.New(sha1.New, []byte("test"))
		_, e := sha1Hash.Write([]byte(stringToSign))
		assert.Nil(t, e)
		assert.Equal(t, "LOG test:"+base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil)), wrapper.Request.Header.Get("Authorization"))
	case <-time.After(100 * time.Millisecond):
		t.Errorf("Mock server should have received a request.")
	}

	logrus.AddHook(slsLogrusHook)
	logrus.SetFormatter(&hook.NoopFormatter{})
	logrus.SetOutput(ioutil.Discard)

	logrus.Info("Hello world!")
	var date string
	select {
	case wrapper := <-requests:
		assert.Equal(t, "POST", wrapper.Request.Method)
		assert.Equal(t, "/logstores/test/shards/lb", wrapper.Request.RequestURI)
		date = wrapper.Request.Header.Get("Date")
		md5 := wrapper.Request.Header.Get("Content-Md5")
		stringToSign := "POST\n" + md5 + "\napplication/x-protobuf\n" + date + "\nx-log-apiversion:0.6.0\nx-log-bodyrawsize:0\nx-log-signaturemethod:hmac-sha1\n/logstores/test/shards/lb"
		sha1Hash := hmac.New(sha1.New, []byte("test"))
		_, e := sha1Hash.Write([]byte(stringToSign))
		assert.Nil(t, e)
		assert.Equal(t, "LOG test:"+base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil)), wrapper.Request.Header.Get("Authorization"))

		group := new(hook.LogGroup)
		assert.Nil(t, proto.Unmarshal(wrapper.Body, group))
		assert.Equal(t, 1, len(group.Logs))
		apiTime, err := time.Parse(http.TimeFormat, date)
		assert.Nil(t, err)
		assert.InDelta(t, apiTime.Unix(), int32(*group.Logs[0].Time), 1)
		assert.Equal(t, "__topic__", *group.Logs[0].Contents[0].Key)
		assert.Equal(t, "test", *group.Logs[0].Contents[0].Value)
		assert.Equal(t, "level", *group.Logs[0].Contents[1].Key)
		assert.Equal(t, "info", *group.Logs[0].Contents[1].Value)
		assert.Equal(t, "message", *group.Logs[0].Contents[2].Key)
		assert.Equal(t, "Hello world!", *group.Logs[0].Contents[2].Value)
	case <-time.After(300 * time.Millisecond):
		t.Errorf("Mock server should have received a request.")
	}
}
