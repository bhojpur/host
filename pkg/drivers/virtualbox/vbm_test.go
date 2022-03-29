package virtualbox

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
	"errors"
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidCheckVBoxManageVersion(t *testing.T) {
	var tests = []struct {
		version string
	}{
		{"5.1"},
		{"5.0.8r103449"},
		{"5.0"},
		{"4.10"},
		{"4.3.1"},
	}

	for _, test := range tests {
		err := checkVBoxManageVersion(test.version)

		assert.NoError(t, err)
	}
}

func TestInvalidCheckVBoxManageVersion(t *testing.T) {
	var tests = []struct {
		version       string
		expectedError string
	}{
		{"3.9", `We support Virtualbox starting with version 5. Your VirtualBox install is "3.9". Please upgrade at https://www.virtualbox.org`},
		{"4.0", `We support Virtualbox starting with version 5. Your VirtualBox install is "4.0". Please upgrade at https://www.virtualbox.org`},
		{"4.1.1", `We support Virtualbox starting with version 5. Your VirtualBox install is "4.1.1". Please upgrade at https://www.virtualbox.org`},
		{"4.2.36-104064", `We support Virtualbox starting with version 5. Your VirtualBox install is "4.2.36-104064". Please upgrade at https://www.virtualbox.org`},
		{"X.Y", `We support Virtualbox starting with version 5. Your VirtualBox install is "X.Y". Please upgrade at https://www.virtualbox.org`},
		{"", `We support Virtualbox starting with version 5. Your VirtualBox install is "". Please upgrade at https://www.virtualbox.org`},
	}

	for _, test := range tests {
		err := checkVBoxManageVersion(test.version)

		assert.EqualError(t, err, test.expectedError)
	}
}

func TestVbmOutErr(t *testing.T) {
	var cmdRun *exec.Cmd
	vBoxManager := NewVBoxManager()
	vBoxManager.runCmd = func(cmd *exec.Cmd) error {
		cmdRun = cmd
		fmt.Fprint(cmd.Stdout, "Printed to StdOut")
		fmt.Fprint(cmd.Stderr, "Printed to StdErr")
		return nil
	}

	stdOut, stdErr, err := vBoxManager.vbmOutErr("arg1", "arg2")

	assert.Equal(t, []string{vboxManageCmd, "arg1", "arg2"}, cmdRun.Args)
	assert.Equal(t, "Printed to StdOut", stdOut)
	assert.Equal(t, "Printed to StdErr", stdErr)
	assert.NoError(t, err)
}

func TestVbmOutErrError(t *testing.T) {
	vBoxManager := NewVBoxManager()
	vBoxManager.runCmd = func(cmd *exec.Cmd) error { return errors.New("BUG") }

	_, _, err := vBoxManager.vbmOutErr("arg1", "arg2")

	assert.EqualError(t, err, "BUG")
}

func TestVbmOutErrNotFound(t *testing.T) {
	vBoxManager := NewVBoxManager()
	vBoxManager.runCmd = func(cmd *exec.Cmd) error { return &exec.Error{Err: exec.ErrNotFound} }

	_, _, err := vBoxManager.vbmOutErr("arg1", "arg2")

	assert.Equal(t, ErrVBMNotFound, err)
}

func TestVbmOutErrFailingWithExitStatus(t *testing.T) {
	vBoxManager := NewVBoxManager()
	vBoxManager.runCmd = func(cmd *exec.Cmd) error {
		fmt.Fprint(cmd.Stderr, "error: Unable to run vbox")
		return errors.New("exit status BUG")
	}

	_, _, err := vBoxManager.vbmOutErr("arg1", "arg2", "arg3")

	assert.EqualError(t, err, vboxManageCmd+" arg1 arg2 arg3 failed:\nerror: Unable to run vbox")
}

func TestVbmOutErrRetryOnce(t *testing.T) {
	var cmdRun *exec.Cmd
	var runCount int
	vBoxManager := NewVBoxManager()
	vBoxManager.runCmd = func(cmd *exec.Cmd) error {
		cmdRun = cmd

		runCount++
		if runCount == 1 {
			fmt.Fprint(cmd.Stderr, "error: The object is not ready")
			return errors.New("Fail the first time it's called")
		}

		fmt.Fprint(cmd.Stdout, "Printed to StdOut")
		return nil
	}

	stdOut, stdErr, err := vBoxManager.vbmOutErr("command", "arg")

	assert.Equal(t, 2, runCount)
	assert.Equal(t, []string{vboxManageCmd, "command", "arg"}, cmdRun.Args)
	assert.Equal(t, "Printed to StdOut", stdOut)
	assert.Empty(t, stdErr)
	assert.NoError(t, err)
}

func TestVbmOutErrRetryMax(t *testing.T) {
	var runCount int
	vBoxManager := NewVBoxManager()
	vBoxManager.runCmd = func(cmd *exec.Cmd) error {
		runCount++
		fmt.Fprint(cmd.Stderr, "error: The object is not ready")
		return errors.New("Always fail")
	}

	stdOut, stdErr, err := vBoxManager.vbmOutErr("command", "arg")

	assert.Equal(t, 5, runCount)
	assert.Empty(t, stdOut)
	assert.Equal(t, "error: The object is not ready", stdErr)
	assert.Error(t, err)
}
