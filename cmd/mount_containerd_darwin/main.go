//go:build darwin
// +build darwin

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
	"strconv"

	"github.com/ukontainer/darwin-snapshotter/hdiu"
)

func usage() {
	fmt.Printf("invalid args: usage:\n \t %s mount target source [option]\n \t %s unmount target [flags]\n", os.Args[0], os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage();
	}

	var target = os.Args[2]

	switch os.Args[1] {
	case "mount":
		if (len(os.Args) < 4) {
			usage()
		}
		var (
			source = os.Args[3]
			options []string
		)
		if (len(os.Args) == 5) {
			options = append(options, os.Args[4])
		}

		if err := hdiu.AttachSparseBundle(target, source, options); err != nil {
			fmt.Printf("failed to mount\n")
			os.Exit(1)
		}
	case "unmount":
		var flags int

		if (len(os.Args) == 4) {
			flags, _ = strconv.Atoi(os.Args[3])
		}

		if err := hdiu.DetachSparseBundle(target, flags); err != nil {
			fmt.Printf("failed to unmount\n")
			os.Exit(1)
		}
	default:
		usage();
	}


}
