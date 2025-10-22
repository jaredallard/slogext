package slogext_test

import (
	"testing"

	"github.com/jaredallard/slogext"
	"gotest.tools/v3/assert"
)

func TestCanCaptureWithCapturedLogger(t *testing.T) {
	log, buf := slogext.NewCapturedLogger()
	log.Info("hello world")

	assert.Equal(t, buf.String(), "INFO hello world\n")
}
