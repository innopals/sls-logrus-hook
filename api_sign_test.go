package hook_test

import (
	"testing"

	hook "github.com/innopals/sls-logrus-hook"
	"github.com/stretchr/testify/assert"
)

func TestApiSign(t *testing.T) {
	// GET
	headers := make(map[string]string)
	headers[hook.HeaderLogVersion] = "0.6.0"
	headers[hook.HeaderLogSignatureMethod] = "hmac-sha1"
	headers[hook.HeaderHost] = "test.cn-hangzhou.log.aliyuncs.com"
	headers[hook.HeaderDate] = "Wed, 29 May 2019 16:00:00 GMT"

	sign, err := hook.APISign("63D0CEA9D550E495FDE1B81310951BD7", "POST", headers, "logstores/test")
	assert.Nil(t, err)
	assert.Equal(t, "KYke+ObziWJpOk305xnagI/omps=", sign)

	// POST
	headers = make(map[string]string)
	headers[hook.HeaderLogVersion] = "0.6.0"
	headers[hook.HeaderLogSignatureMethod] = "hmac-sha1"
	headers[hook.HeaderHost] = "test.cn-hangzhou.log.aliyuncs.com"
	headers[hook.HeaderDate] = "Wed, 29 May 2019 16:00:00 GMT"
	// Content: "Hello world!"
	headers[hook.HeaderContentMd5] = "86FB269D190D2C85F6E0468CECA42A20"
	headers[hook.HeaderContentLength] = "12"
	headers[hook.HeaderLogBodyRawSize] = "0"

	sign, err = hook.APISign("2974A71FE7FCCEF63C436826DD53BA6D", "POST", headers, "logstores/test")
	assert.Nil(t, err)
	assert.Equal(t, "rAU8ZoY8y3GM9cA8XDkf6GiJVi8=", sign)
}
