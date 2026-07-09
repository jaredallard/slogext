// Copyright (C) 2026 slogext contributors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package slogext

import (
	"io"
	"log/slog"
)

// NewNullLogger returns a [Logger] that does not output anywhere.
func NewNullLogger() Logger {
	return NewWithHandler(slog.NewTextHandler(io.Discard, nil))
}
