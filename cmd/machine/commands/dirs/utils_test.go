package dirs

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"path"
	"strings"
	"testing"

	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

func TestGetBaseDir(t *testing.T) {
	// reset any override env var
	BaseDir = ""

	homeDir := mutils.GetHomeDir()
	baseDir := GetBaseDir()

	if strings.Index(baseDir, homeDir) != 0 {
		t.Fatalf("expected base dir with prefix %s; received %s", homeDir, baseDir)
	}
}

func TestGetCustomBaseDir(t *testing.T) {
	root := "/tmp"
	BaseDir = root
	baseDir := GetBaseDir()

	if strings.Index(baseDir, root) != 0 {
		t.Fatalf("expected base dir with prefix %s; received %s", root, baseDir)
	}
	BaseDir = ""
}

func TestGetMachineDir(t *testing.T) {
	root := "/tmp"
	BaseDir = root
	machineDir := GetMachineDir()

	if strings.Index(machineDir, root) != 0 {
		t.Fatalf("expected machine dir with prefix %s; received %s", root, machineDir)
	}

	path, filename := path.Split(machineDir)
	if strings.Index(path, root) != 0 {
		t.Fatalf("expected base path of %s; received %s", root, path)
	}
	if filename != "machines" {
		t.Fatalf("expected machine dir \"machines\"; received %s", filename)
	}
	BaseDir = ""
}

func TestGetMachineCertDir(t *testing.T) {
	root := "/tmp"
	BaseDir = root
	clientDir := GetMachineCertDir()

	if strings.Index(clientDir, root) != 0 {
		t.Fatalf("expected machine client cert dir with prefix %s; received %s", root, clientDir)
	}

	path, filename := path.Split(clientDir)
	if strings.Index(path, root) != 0 {
		t.Fatalf("expected base path of %s; received %s", root, path)
	}
	if filename != "certs" {
		t.Fatalf("expected machine client dir \"certs\"; received %s", filename)
	}
	BaseDir = ""
}
