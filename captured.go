// Copyright (C) 2025 slogext contributors
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public
// License along with this program. If not, see
// <https://www.gnu.org/licenses/>.
//
// SPDX-License-Identifier: LGPL-3.0

package slogext

import "bytes"

// NewCapturedLogger returns a [Logger] who's output is captured for
// future usage.
func NewCapturedLogger() (Logger, *bytes.Buffer) {
	b := new(bytes.Buffer)
	return NewWithWriter(b), b
}
