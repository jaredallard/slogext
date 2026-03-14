package slogext_test

import (
	"testing"

	"go.rgst.io/jaredallard/slogext/v2"
	"gotest.tools/v3/assert"
)

func TestCanCaptureWithCapturedLogger(t *testing.T) {
	log, buf := slogext.NewCapturedLogger()
	log.Info("hello world")

	assert.Equal(t, buf.String(), "INFO hello world\n")
}
