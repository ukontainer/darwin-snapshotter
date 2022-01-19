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

package hdiu

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/containerd/containerd/log"
)

func CreateSparseBundle(image string, fileSizeMB int) error {
	// create a disk image file from a snapshot
	args := []string{
		"create",
		"-size", strconv.Itoa(fileSizeMB) + "m",
		image,
		"-ov",
		"-type", "SPARSEBUNDLE",
		"-fs", "HFS+",
	}

	return hdiutilCommand(args...)
}

func AttachSparseBundle(target, source string, options []string) error {
	args := []string{
		"attach",
		source,
		"-mountpoint", target,
		"-nobrowse",
		"-owners", "on",
		"-noverify",
		"-noautofsck",
	}
	args = append(args, options...)

	return hdiutilCommand(args...)
}

func DetachSparseBundle(target string, flags int) error {
	if _, err := os.Stat(target); err != nil {
		return nil
	}

	return hdiutilCommand("detach", target)
}

func hdiutilCommand(args ...string) error {
	log.L.Debugf("hdiutil %s", args)
	cmd := exec.Command("hdiutil", args...)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, out)
	}

	return nil
}
