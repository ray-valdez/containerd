//go:build linux
// +build linux

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"fmt"
	"os"

	"github.com/containerd/containerd/runtime/v2/shim"
	split "github.com/containerd/containerd/runtime/v2/split/pkg"
)

// name of the project tags
const (
	PROJECT                = "Split API"
	VERSION                = "0.1"
	COMMIT                 = "TBD"
	DefaultKataRuntimeName = "io.containerd.split.v2"
)

func main() {

	if len(os.Args) == 2 && os.Args[1] == "--version" {
		fmt.Printf("%s containerd shim: id: %q, version: %s, commit: %s\n", PROJECT, DefaultKataRuntimeName,
			VERSION, COMMIT)
		os.Exit(0)
	}

	// init and execute the shim
	shim.Run(DefaultKataRuntimeName, split.New)
}
