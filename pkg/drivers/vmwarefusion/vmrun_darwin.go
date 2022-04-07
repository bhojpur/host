package vmwarefusion

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
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/bhojpur/host/pkg/machine/log"
)

var (
	vmrunbin    = setVmwareCmd("vmrun")
	vdiskmanbin = setVmwareCmd("vmware-vdiskmanager")
)

var (
	ErrMachineExist    = errors.New("machine already exists")
	ErrMachineNotExist = errors.New("machine does not exist")
	ErrVMRUNNotFound   = errors.New("VMRUN not found")
)

// detect the vmrun and vmware-vdiskmanager cmds' path if needed
func setVmwareCmd(cmd string) string {
	if path, err := exec.LookPath(cmd); err == nil {
		return path
	}
	return filepath.Join("/Applications/VMware Fusion.app/Contents/Library/", cmd)
}

func vmrun(args ...string) (string, string, error) {
	// vmrun with nogui on VMware Fusion through at least 8.0.1 doesn't work right
	// if the umask is set to not allow world-readable permissions
	_ = syscall.Umask(022)
	cmd := exec.Command(vmrunbin, args...)
	if os.Getenv("MACHINE_DEBUG") != "" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout, cmd.Stderr = &stdout, &stderr
	log.Debugf("executing: %v %v", vmrunbin, strings.Join(args, " "))

	err := cmd.Run()
	if err != nil {
		if ee, ok := err.(*exec.Error); ok && ee == exec.ErrNotFound {
			err = ErrVMRUNNotFound
		}
	}

	return stdout.String(), stderr.String(), err
}

// Make a vmdk disk image with the given size (in MB).
func vdiskmanager(dest string, size int) error {
	cmd := exec.Command(vdiskmanbin, "-c", "-t", "0", "-s", fmt.Sprintf("%dMB", size), "-a", "lsilogic", dest)
	if os.Getenv("MACHINE_DEBUG") != "" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	if stdout := cmd.Run(); stdout != nil {
		if ee, ok := stdout.(*exec.Error); ok && ee == exec.ErrNotFound {
			return ErrVMRUNNotFound
		}
	}
	return nil
}
