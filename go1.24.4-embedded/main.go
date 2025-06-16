// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The go1.24.4-embedded command compiles and runs the go command from Embedded Go
// 1.24.4
//
// To install, run:
//
//	$ go install github.com/embeddedgo/dl/go1.24.4-embedded@latest
//	$ go1.24.4-embedded download
//
// Next you can select the newly installed Embedded Go toolchain by setting
//
//  GOTOOLCHAIN=go1.24.4-embedded
//
// in your environment or use the go1.24.4-embedded command directly.
package main

import "github.com/embeddedgo/dl/internal/version"

func main() {
	version.RunTag("go1.24.4-embedded")
}