package slogext_test

import (
	"bytes"
	"io"
	"strings"
	"testing"
	"testing/synctest"

	"github.com/charmbracelet/x/ansi"
	"github.com/creack/pty"
	"go.rgst.io/jaredallard/slogext/v2"
	"gotest.tools/v3/assert"
)

func TestCanCaptureWithCapturedLogger(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		log, buf := slogext.NewCapturedLogger()
		log.Info("hello world")

		assert.Equal(t, strings.TrimSpace(buf.String()), `{"time":"1999-12-31T16:00:00-08:00","level":"info","msg":"hello world"}`)
	})
}

func TestDisplaysFormattedLogsTTY(t *testing.T) {
	p, tty, err := pty.Open()
	assert.NilError(t, err, "expected pty.Open to not fail")
	defer p.Close()

	buf := bytes.Buffer{}
	go io.Copy(&buf, p)

	log := slogext.NewWithWriter(tty)
	log.Info("hello world")
	tty.Close()

	assert.Equal(t, strings.TrimSpace(ansi.Strip(buf.String())), "INFO hello world")
}
