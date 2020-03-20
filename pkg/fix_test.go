/*
Copyright 2020 The Kubernetes Authors.

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

package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestFix(t *testing.T) {
	for _, module := range []bool{true, false} {
		for _, vendor := range []bool{true, false} {
			for _, version := range []string{"v0.17.4", "v0.18.0-beta.2"} {
				name := ifelse(module, "mod", "gopath") + "_" + ifelse(vendor, "vendor", "novendor") + "_" + version
				t.Run(name, func(t *testing.T) {
					defer restoreEnv()()
					tempDir, cleanupTempDir := tempDir(t, name)
					defer cleanupTempDir()
					copyDir(t, "testdata/input/", tempDir)
					prepDeps(t, tempDir, module, vendor, version)

					subdirs, err := ioutil.ReadDir(filepath.Join(tempDir, "src", "example.com"))
					if err != nil {
						t.Fatal(err)
					}
					for _, subdir := range subdirs {
						if !subdir.IsDir() || subdir.Name() == "vendor" {
							continue
						}
						t.Run(subdir.Name(), func(t *testing.T) {
							runFix(t, filepath.Join(tempDir, "src", "example.com"), fmt.Sprintf("./%s/...", subdir.Name()))
							diffOutput(t, filepath.Join(tempDir, "src", "example.com"), subdir.Name())
						})
					}
				})
			}
		}
	}
}

func ifelse(b bool, t, f string) string {
	if b {
		return t
	}
	return f
}

func restoreEnv() func() {
	gopath := os.Getenv("GOPATH")
	goflags := os.Getenv("GOFLAGS")
	goproxy := os.Getenv("GOPROXY")
	gomodule := os.Getenv("GO111MODULE")
	return func() {
		os.Setenv("GOPATH", gopath)
		os.Setenv("GOFLAGS", goflags)
		os.Setenv("GOPROXY", goproxy)
		os.Setenv("GO111MODULE", gomodule)
	}
}

func tempDir(t *testing.T, name string) (string, func()) {
	tempDir, err := ioutil.TempDir("", name)
	if err != nil {
		t.Fatal(err)
	}
	return tempDir, func() { os.RemoveAll(tempDir) }
}

func copyDir(t *testing.T, in, out string) {
	err := filepath.Walk(in, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if in == path || !strings.HasPrefix(path, in) {
			return nil
		}
		rel := strings.TrimPrefix(path, in)
		if info.IsDir() {
			return os.MkdirAll(filepath.Join(out, rel), os.FileMode(0755))
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filepath.Join(out, rel), data, os.FileMode(0644))
	})
	if err != nil {
		t.Error(err)
	}
}

func prepDeps(t *testing.T, tempDir string, module, vendor bool, version string) {
	srcDir := filepath.Join(tempDir, "src")
	exampleDir := filepath.Join(srcDir, "example.com")
	vendorDir := filepath.Join(exampleDir, "vendor")
	goModFile := filepath.Join(exampleDir, "go.mod")
	writeGoMod(t, goModFile, version)
	switch {
	case module && vendor:
		// enable modules and vendor and build vendor dir
		os.Setenv("GO111MODULE", "on")
		os.Setenv("GOFLAGS", "-mod=vendor")
		runGoModVendor(t, exampleDir)

	case module && !vendor:
		// enable modules and disable vendor
		os.Setenv("GO111MODULE", "on")
		os.Setenv("GOFLAGS", "")

	case !module && vendor:
		// enable modules and build vendor dir to assemble required dependencies
		os.Setenv("GO111MODULE", "on")
		runGoModVendor(t, exampleDir)

		// set $GOPATH to tempDir
		os.Setenv("GOPATH", tempDir)

		// disable modules and remove go.mod file
		os.Setenv("GO111MODULE", "off")
		if err := os.Remove(goModFile); err != nil {
			t.Fatal(err)
		}

	case !module && !vendor:
		// enable modules and build vendor dir to assemble required dependencies
		os.Setenv("GO111MODULE", "on")
		runGoModVendor(t, exampleDir)

		// set $GOPATH to tempDir and move dependencies from vendor there
		os.Setenv("GOPATH", tempDir)
		dependencies, err := ioutil.ReadDir(vendorDir)
		if err != nil {
			t.Fatal(err)
		}
		for _, dependency := range dependencies {
			err := os.Rename(
				filepath.Join(vendorDir, dependency.Name()),
				filepath.Join(srcDir, dependency.Name()),
			)
			if err != nil {
				t.Fatal(err)
			}
		}

		// disable modules and remove go.mod file and vendor dir
		os.Setenv("GO111MODULE", "off")
		if err := os.Remove(goModFile); err != nil {
			t.Fatal(err)
		}
		if err := os.Remove(filepath.Join(exampleDir, "vendor")); err != nil {
			t.Fatal(err)
		}
	}
}

func writeGoMod(t *testing.T, goMod string, version string) {
	err := ioutil.WriteFile(goMod, []byte(fmt.Sprintf(`module example.com

go 1.13
	
require (
	k8s.io/api %[1]s
	k8s.io/apimachinery %[1]s
	k8s.io/client-go %[1]s
)`, version)), os.FileMode(0644))
	if err != nil {
		t.Fatal(err)
	}
}

func runGoModVendor(t *testing.T, dir string) {
	cmd := exec.Command("go", "mod", "vendor")
	cmd.Dir = dir
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Log(string(output))
		t.Fatal(err)
	}
}

func runFix(t *testing.T, dir, pkg string) {
	b := bytes.NewBuffer([]byte{})
	o := DefaultFixOptions()
	o.Dir = dir
	o.Out = b
	o.Packages = []string{pkg}
	err := o.Run()
	if err != nil {
		t.Log(b.String())
		t.Fatal(err)
	}
}

func diffOutput(t *testing.T, resultDir, subDir string) {
	err := filepath.Walk(filepath.Join(resultDir, subDir), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name() == "vendor" {
				return filepath.SkipDir
			}
			return nil
		}
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		rel := strings.TrimPrefix(path, resultDir)
		name := filepath.Base(path)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			t.Errorf("%s: %v", name, err)
			return nil
		}

		expectFile := filepath.Join("testdata", "expect", "src", "example.com", rel)
		expectData, err := ioutil.ReadFile(expectFile)
		ok := err == nil && bytes.Equal(data, expectData)
		if !ok {
			if os.Getenv("UPDATE_FIXTURE_DATA") == "true" {
				os.MkdirAll(filepath.Dir(expectFile), os.FileMode(0755))
				if err := ioutil.WriteFile(expectFile, data, os.FileMode(0644)); err != nil {
					t.Errorf("%s: %v", name, err)
					return nil
				}
				t.Errorf("%s: wrote testdata, rerun test", name)
				return nil
			} else if err != nil {
				t.Log("set UPDATE_FIXTURE_DATA=true to write expected testdata")
				t.Errorf("%s: %v", name, err)
				return nil
			} else {
				t.Log("set UPDATE_FIXTURE_DATA=true to write expected testdata")
				t.Log(cmp.Diff(string(data), string(expectData)))
				t.Errorf("%s: diff", name)
				return nil
			}
		}
		return nil
	})
	if err != nil {
		t.Error(err)
	}
}
