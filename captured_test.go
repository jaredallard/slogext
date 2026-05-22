package slogext_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"testing/synctest"
	"time"

	"github.com/charmbracelet/x/ansi"
	"github.com/creack/pty"
	"go.rgst.io/jaredallard/slogext/v2"
	"gotest.tools/v3/assert"
)

func TestCanCaptureWithCapturedLogger(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		log, buf := slogext.NewCapturedLogger()
		log.Info("hello world")

		assert.Equal(t, strings.TrimSpace(buf.String()),
			fmt.Sprintf(
				`{"time":%q,"level":"info","msg":"hello world"}`,
				time.Now().Format(time.RFC3339),
			),
		)
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

	time.Sleep(200 * time.Millisecond)
	tty.Close()

	assert.Equal(t, strings.TrimSpace(ansi.Strip(buf.String())), "INFO hello world")
}

func TestCanForceColoredLogs(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		t.Setenv("CLICOLOR_FORCE", "1")

		log, buf := slogext.NewCapturedLogger()
		log.Info("hello world")

		assert.Equal(t, buf.String(), "\x1b[1;38;5;86mINFO\x1b[m hello world\n")
	})
}
