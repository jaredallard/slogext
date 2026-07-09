// Copyright (C) 2026 slogext contributors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package slogext

import "testing"

// NewTestLogger creates a new logger for testing purposes. The logging
// level is set to DebugLevel to ensure all logs are captured.
func NewTestLogger(t *testing.T) Logger {
	logger := New()
	logger.SetLevel(DebugLevel)
	return logger.With("test.name", t.Name())
}
