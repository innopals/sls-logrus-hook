package hook_test

import (
	hook "github.com/innopals/sls-logrus-hook"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateClient(t *testing.T) {
	var err error
	_, err = hook.NewSlsClient("", "", "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, "Sls endpoint should not be empty", err.Error())
	_, err = hook.NewSlsClient("http://localhost", "", "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, "Sls access key should not be empty", err.Error())
	_, err = hook.NewSlsClient("http://localhost", "access_key", "", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, "Sls access secret should not be empty", err.Error())
	_, err = hook.NewSlsClient("http://localhost", "access_key", "access_secret", "", "")
	assert.NotNil(t, err)
	assert.Equal(t, "Sls log store should not be empty", err.Error())
}
