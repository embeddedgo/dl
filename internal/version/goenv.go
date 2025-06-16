// Copyright 2025 The Embedded Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version

import (
	"bytes"
	"errors"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func goGetenv(name string) string {
	out, err := exec.Command("go", "env", name).Output()
	if err != nil {
		log.Fatalf("go env: %v", err)
	}
	return string(bytes.TrimSpace(out))
}

const goenvName = "go.env"

func findGoenv() (goenvPath string) {
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "build", "install":
			// ok
		default:
			return ""
		}
	}
	//if goGetenv("GOOS") == "noos" && os.Getenv("GOENV") != "" {
	if os.Getenv("GOOS") == "noos" && os.Getenv("GOENV") != "" {
		// GOENV is probably set by user manualy to be used together with GOOS=noos
		// TODO: what about go env -w GOENV
		return ""
	}
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	for {
		goenvPath = filepath.Join(wd, goenvName)
		fi, err := os.Stat(goenvPath)
		if err == nil {
			if !fi.Mode().IsRegular() {
				log.Fatalf("%s: not a regular file", goenvPath)
			}
			return goenvPath
		}
		if !errors.Is(err, fs.ErrNotExist) {
			log.Fatal(err)
		}
		_, err = os.Stat(filepath.Join(wd, "go.mod"))
		if err == nil {
			return "" // found go.mod but no goenvName, stop here
		}
		if !errors.Is(err, fs.ErrNotExist) {
			log.Fatal(err)
		}
		wd = filepath.Dir(wd)
		if wd == "/" || wd == "." { // FIXME: windows?
			return ""
		}
	}
}
