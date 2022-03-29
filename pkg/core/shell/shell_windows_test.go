package shell

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
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDetect(t *testing.T) {
	defer func(shell string) { os.Setenv("SHELL", shell) }(os.Getenv("SHELL"))
	os.Setenv("SHELL", "")

	shell, err := Detect()

	assert.Equal(t, "powershell", shell)
	assert.NoError(t, err)
}

func TestGetNameAndItsPpidOfCurrent(t *testing.T) {
	shell, shellppid, err := getNameAndItsPpid(os.Getpid())

	assert.Equal(t, "shell.test.exe", shell)
	assert.Equal(t, os.Getppid(), shellppid)
	assert.NoError(t, err)
}

func TestGetNameAndItsPpidOfParent(t *testing.T) {
	shell, _, err := getNameAndItsPpid(os.Getppid())

	assert.Equal(t, "go.exe", shell)
	assert.NoError(t, err)
}

func TestGetNameAndItsPpidOfGrandParent(t *testing.T) {
	shell, shellppid, err := getNameAndItsPpid(os.Getppid())
	shell, shellppid, err = getNameAndItsPpid(shellppid)

	assert.Equal(t, "powershell.exe", shell)
	assert.NoError(t, err)
}
