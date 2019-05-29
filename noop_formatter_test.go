package hook_test

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"sls-logrus-hook"
	"testing"
)

func TestNoopFormatter(t *testing.T) {
	formatter := &hook.NoopFormatter{}
	bytes, err := formatter.Format(&logrus.Entry{})
	assert.Nil(t, bytes)
	assert.Nil(t, err)
}
