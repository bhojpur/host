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
	"strings"
	"syscall"

	"github.com/bhojpur/host/pkg/core/log"
)

// IsVTXDisabled checks if VT-X is disabled in the BIOS. If it is, the vm will fail to start.
// If we can't be sure it is disabled, we carry on and will check the vm logs after it's started.
func (d *Driver) IsVTXDisabled() bool {
	features, err := syscall.Sysctl("machdep.cpu.features")
	if err != nil {
		log.Debugf("Couldn't check that VT-X/AMD-v is enabled. Will check that the vm is properly created: %v", err)
		return false
	}
	return isVTXDisabled(features)
}

func isVTXDisabled(features string) bool {
	return !strings.Contains(features, "VMX")
}

func detectVBoxManageCmd() string {
	return detectVBoxManageCmdInPath()
}

func getShareDriveAndName() (string, string) {
	return "Users", "/Users"
}

func isHyperVInstalled() bool {
	return false
}
