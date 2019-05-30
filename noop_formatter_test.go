package hook_test

import (
	"testing"

	hook "github.com/innopals/sls-logrus-hook"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNoopFormatter(t *testing.T) {
	formatter := &hook.NoopFormatter{}
	bytes, err := formatter.Format(&logrus.Entry{})
	assert.Nil(t, bytes)
	assert.Nil(t, err)
}
