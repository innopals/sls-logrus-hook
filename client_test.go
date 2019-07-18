package hook_test

import (
	"testing"

	hook "github.com/innopals/sls-logrus-hook"
	"github.com/stretchr/testify/assert"
)

func TestCreateClient(t *testing.T) {
	var err error
	_, err = hook.NewSlsClient("", "", "", "", "", hook.DefaultTimeout)
	assert.NotNil(t, err)
	assert.Equal(t, "Sls endpoint should not be empty", err.Error())
	_, err = hook.NewSlsClient("http://localhost", "", "", "", "", hook.DefaultTimeout)
	assert.NotNil(t, err)
	assert.Equal(t, "Sls access key should not be empty", err.Error())
	_, err = hook.NewSlsClient("http://localhost", "access_key", "", "", "", hook.DefaultTimeout)
	assert.NotNil(t, err)
	assert.Equal(t, "Sls access secret should not be empty", err.Error())
	_, err = hook.NewSlsClient("http://localhost", "access_key", "access_secret", "", "", hook.DefaultTimeout)
	assert.NotNil(t, err)
	assert.Equal(t, "Sls log store should not be empty", err.Error())
}
