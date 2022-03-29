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
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
)

// re-implementation of private function in https://github.com/golang/go/blob/master/src/syscall/syscall_windows.go#L945
func getProcessEntry(pid int) (pe *syscall.ProcessEntry32, err error) {
	snapshot, err := syscall.CreateToolhelp32Snapshot(syscall.TH32CS_SNAPPROCESS, 0)
	if err != nil {
		return nil, err
	}
	defer syscall.CloseHandle(syscall.Handle(snapshot))

	var processEntry syscall.ProcessEntry32
	processEntry.Size = uint32(unsafe.Sizeof(processEntry))
	err = syscall.Process32First(snapshot, &processEntry)
	if err != nil {
		return nil, err
	}

	for {
		if processEntry.ProcessID == uint32(pid) {
			pe = &processEntry
			return
		}

		err = syscall.Process32Next(snapshot, &processEntry)
		if err != nil {
			return nil, err
		}
	}
}

// getNameAndItsPpid returns the exe file name its parent process id.
func getNameAndItsPpid(pid int) (exefile string, parentid int, err error) {
	pe, err := getProcessEntry(pid)
	if err != nil {
		return "", 0, err
	}

	name := syscall.UTF16ToString(pe.ExeFile[:])
	return name, int(pe.ParentProcessID), nil
}

func Detect() (string, error) {
	shell := os.Getenv("SHELL")

	if shell == "" {
		shell, shellppid, err := getNameAndItsPpid(os.Getppid())
		if err != nil {
			return "cmd", err // defaulting to cmd
		}
		if strings.Contains(strings.ToLower(shell), "powershell") {
			return "powershell", nil
		} else if strings.Contains(strings.ToLower(shell), "cmd") {
			return "cmd", nil
		} else {
			shell, _, err := getNameAndItsPpid(shellppid)
			if err != nil {
				return "cmd", err // defaulting to cmd
			}
			if strings.Contains(strings.ToLower(shell), "powershell") {
				return "powershell", nil
			} else if strings.Contains(strings.ToLower(shell), "cmd") {
				return "cmd", nil
			} else {
				fmt.Printf("You can further specify your shell with either 'cmd' or 'powershell' with the --shell flag.\n\n")
				return "cmd", nil // this could be either powershell or cmd, defaulting to cmd
			}
		}
	}

	if os.Getenv("__fish_bin_dir") != "" {
		return "fish", nil
	}

	return filepath.Base(shell), nil
}
