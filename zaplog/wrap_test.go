package zaplog_test

import (
	"testing"

	"cloud.google.com/go/logging"
	"github.com/savaki/stackdriver/zaplog"
	"github.com/tj/assert"
	"go.uber.org/zap"
)

func TestWrap(t *testing.T) {
	target, err := zap.NewDevelopmentConfig().Build()
	assert.Nil(t, err)

	logger := zaplog.Wrap(target)
	logger.Log(logging.Entry{
		Payload: map[string]interface{}{
			"message": "hello",
			"string":  "string",
			"int":     1,
			"bool":    true,
		},
		Labels: map[string]string{
			"foo": "bar",
		},
	})
}
