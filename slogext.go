// Copyright (C) 2026 slogext contributors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

// Package slogext is a small wrapper around the [log/slog] package
// focused on providing consistency in logging across the stencil
// codebase.
package slogext

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"time"

	charmlog "charm.land/log/v2"
	"github.com/charmbracelet/x/term"
	"github.com/muesli/termenv"
)

// _ ensures that the logger struct satisfies the Logger interface.
var _ Logger = &logger{}

// Logger is a [slog.Logger] interface with extra functionality.
type Logger interface {
	Info(string, ...any)
	Infof(string, ...any)
	Debug(string, ...any)
	Debugf(string, ...any)
	Error(string, ...any)
	Errorf(string, ...any)
	Warn(string, ...any)
	Warnf(string, ...any)
	With(...any) Logger
	WithError(error) Logger
	SetLevel(charmlog.Level)
	GetHandler() slog.Handler
}

// Level is a logging level.
type Level = charmlog.Level

// Contains valid [Level] values.
const (
	// DebugLevel is for debug level logs.
	DebugLevel Level = charmlog.DebugLevel

	// InfoLevel is for info level logs.
	InfoLevel Level = charmlog.InfoLevel

	// WarnLevel is for warn level logs.
	WarnLevel Level = charmlog.WarnLevel

	// ErrorLevel is for error level logs
	ErrorLevel Level = charmlog.ErrorLevel

	// FatalLevel is for fatal level logs.
	FatalLevel Level = charmlog.FatalLevel
)

// New creates a new [Logger] using the slog package.
func New() Logger {
	return NewWithWriter(os.Stdout)
}

// NewWithWriter creates a new [Logger] using the slog package with a
// custom writer target. If the writer does not appear to be a terminal,
// then JSON is used.
func NewWithWriter(w io.Writer) Logger {
	var prettyPrint bool
	if f, ok := w.(*os.File); ok && term.IsTerminal(f.Fd()) {
		prettyPrint = true
	}

	// Essentially; if CLICOLOR_FORCE is set this'll returning something
	// other than ASCII, but also allows us to forcibly disable. This
	// matches how charm does the logic on their side as well.
	if profile := termenv.EnvColorProfile(); profile != termenv.Ascii {
		prettyPrint = true
	}

	opts := charmlog.Options{Formatter: charmlog.JSONFormatter}
	if prettyPrint {
		opts.Formatter = charmlog.TextFormatter
	} else {
		// JSON options
		opts.ReportTimestamp = true
		opts.TimeFormat = time.RFC3339
	}

	handler := charmlog.NewWithOptions(w, opts)
	return &logger{slog.New(handler), handler}
}

// NewWithHandler creates a new [Logger] using the slog package. Levels
// do not work since it is not a slog native concept.
func NewWithHandler(h slog.Handler) Logger {
	// TODO(jaredallard): Consider exposing a leveler function.
	return &logger{slog.New(h), charmlog.New(io.Discard)}
}

// logger is a simple wrapper around the slog.Logger interface. Use
// [Logger] when passing around loggers in the stencil codebase.
type logger struct {
	*slog.Logger
	handler *charmlog.Logger
}

// With wraps the slog.With method to return a new logger with the
// provided arguments while satisfying the Logger interface.
func (l *logger) With(args ...any) Logger {
	return &logger{l.Logger.With(args...), l.handler}
}

// WithError wraps the slog.With method using a consistent key for
// errors, "error".
func (l *logger) WithError(err error) Logger {
	return &logger{l.Logger.With("error", err), l.handler}
}

// SetLevel updates the level of the current logger to the provided
// level.
func (l *logger) SetLevel(level Level) {
	l.handler.SetLevel(level)
}

// Infof wraps Info with a formatted message.
func (l *logger) Infof(format string, args ...any) {
	l.Info(fmt.Sprintf(format, args...))
}

// Debugf wraps Debug with a formatted message.
func (l *logger) Debugf(format string, args ...any) {
	l.Debug(fmt.Sprintf(format, args...))
}

// Errorf wraps Error with a formatted message.
func (l *logger) Errorf(format string, args ...any) {
	l.Error(fmt.Sprintf(format, args...))
}

// Warnf wraps Warn with a formatted message.
func (l *logger) Warnf(format string, args ...any) {
	l.Warn(fmt.Sprintf(format, args...))
}

// GetHandler returns the [slog.Handler] used by this logger, for
// interoperability purposes.
func (l *logger) GetHandler() slog.Handler {
	return l.handler
}
