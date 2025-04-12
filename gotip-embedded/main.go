// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The gotip-embedded command compiles and runs the go command from the development tree.
//
// To install, run:
//
//	$ go install github.com/embeddedgo/dl/gotip-embedded@latest
//	$ gotip-embedded download
//
// And then use the gotip-embedded command as if it were your normal go command.
//
// To update, run "gotip-embedded download" again. This will always download the main branch.
// To download an alternative branch, run "gotip-embedded download BRANCH".
package main

import "github.com/embeddedgo/dl/internal/version"

func main() {
	version.RunTip()
}
