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
	"testing"

	"github.com/stretchr/testify/assert"
)

const stdOutDiskInfo = `
storagecontrollerbootable0="on"
"SATA-0-0"="/home/ehazlett/.boot2docker/boot2docker.iso"
"SATA-IsEjected"="off"
"SATA-1-0"="/home/ehazlett/vm/test/disk.vmdk"
"SATA-ImageUUID-1-0"="12345-abcdefg"
"SATA-2-0"="none"
nic1="nat"`

func TestVMDiskInfo(t *testing.T) {
	vbox := &VBoxManagerMock{
		args:   "showvminfo default --machinereadable",
		stdOut: stdOutDiskInfo,
	}

	disk, err := getVMDiskInfo("default", vbox)

	assert.Equal(t, "/home/ehazlett/vm/test/disk.vmdk", disk.Path)
	assert.Equal(t, "12345-abcdefg", disk.UUID)
	assert.NoError(t, err)
}

func TestVMDiskInfoError(t *testing.T) {
	vbox := &VBoxManagerMock{
		args: "showvminfo default --machinereadable",
		err:  errors.New("BUG"),
	}

	disk, err := getVMDiskInfo("default", vbox)

	assert.Nil(t, disk)
	assert.EqualError(t, err, "BUG")
}

func TestVMDiskInfoInvalidOutput(t *testing.T) {
	vbox := &VBoxManagerMock{
		args:   "showvminfo default --machinereadable",
		stdOut: "INVALID",
	}

	disk, err := getVMDiskInfo("default", vbox)

	assert.Empty(t, disk.Path)
	assert.Empty(t, disk.UUID)
	assert.NoError(t, err)
}
