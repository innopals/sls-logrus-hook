package hook

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"sort"
	"strings"
)

// APISign Create signature for sls api
func APISign(secret string, method string, headers map[string]string, resource string) (string, error) {
	var contentMD5, contentType, date string
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

	logHeaders = sort.StringSlice(logHeaders)

	stringToSign := method + "\n" +
		contentMD5 + "\n" +
		contentType + "\n" +
		date + "\n" +
		strings.Join(logHeaders, "\n") + "\n" +
		resource

	sha1Hash := hmac.New(sha1.New, []byte(secret))
	if _, e := sha1Hash.Write([]byte(stringToSign)); e != nil {
		return "", e
	}
	return base64.StdEncoding.EncodeToString(sha1Hash.Sum(nil)), nil
}
