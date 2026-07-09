// Copyright (C) 2026 slogext contributors
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.
//
// SPDX-License-Identifier: MPL-2.0

package slogext

import "bytes"

// NewCapturedLogger returns a [Logger] who's output is captured for
// future usage.
func NewCapturedLogger() (Logger, *bytes.Buffer) {
	b := new(bytes.Buffer)
	return NewWithWriter(b), b
}
