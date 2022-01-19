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
	"os"
	"path/filepath"
	"testing"

	"github.com/containerd/continuity/testutil"
)

var randomData = []byte("randomdata")

func createTempDir(t *testing.T) string {
	t.Helper()

	dir, err := os.MkdirTemp("", "test-image")
	if err != nil {
		t.Fatal(err)
	}

	return dir
}

func testSparseBundle(t *testing.T, option string, testerFn func(file string)) {
	tmpDir := createTempDir(t)
	defer os.RemoveAll(tmpDir)

	mountPoint := filepath.Join(tmpDir, "mount"+option)
	sparseBundle := filepath.Join(tmpDir, "test"+option+".sparsebundle")
	err := CreateSparseBundle(sparseBundle, 64)
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if err := os.RemoveAll(sparseBundle); err != nil {
			t.Fatal(err)
		}
	}()

	err = AttachSparseBundle(mountPoint, sparseBundle, []string{option})
	if err != nil {
		t.Fatal(err)
	}

	testfile := filepath.Join(mountPoint, "testfile")
	testerFn(testfile)

	err = DetachSparseBundle(mountPoint, 0)
	if err != nil {
		t.Fatal(err)
	}

}

func TestSparseRO(t *testing.T) {
	testerFn := func(testfile string) {
		if err := os.WriteFile(testfile, randomData, 0777); err != nil {
			t.Logf("write %q failed with %v (EROFS is expected)", testfile, err)
		} else {
			t.Fatalf("write %q should fail (EROFS) but did not fail", testfile)
		}
	}
	testSparseBundle(t, "-readonly", testerFn)
}

func TestSparseRW(t *testing.T) {
	testutil.RequiresRoot(t)
	testerFn := func(testfile string) {
		if err := os.WriteFile(testfile, randomData, 0777); err != nil {
			t.Fatalf("write to %q failed", testfile)
		}
	}
	testSparseBundle(t, "-readwrite", testerFn)
}
