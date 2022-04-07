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
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/bhojpur/host/pkg/machine/log"
	mutils "github.com/bhojpur/host/pkg/machine/utils"
)

type VirtualDisk struct {
	UUID string
	Path string
}

type DiskCreator interface {
	Create(size int, publicSSHKeyPath, diskPath string) error
}

func NewDiskCreator() DiskCreator {
	return &defaultDiskCreator{}
}

type defaultDiskCreator struct{}

// Make a boot2docker VM disk image.
func (c *defaultDiskCreator) Create(size int, publicSSHKeyPath, diskPath string) error {
	log.Debugf("Creating %d MB hard disk image...", size)

	tarBuf, err := mutils.MakeDiskImage(publicSSHKeyPath)
	if err != nil {
		return err
	}

	log.Debug("Calling inner createDiskImage")

	return createDiskImage(diskPath, size, tarBuf)
}

// createDiskImage makes a disk image at dest with the given size in MB. If r is
// not nil, it will be read as a raw disk image to convert from.
func createDiskImage(dest string, size int, r io.Reader) error {
	// Convert a raw image from stdin to the dest VMDK image.
	sizeBytes := int64(size) << 20 // usually won't fit in 32-bit int (max 2GB)
	// FIXME: why isn't this just using the vbm*() functions?
	cmd := exec.Command(vboxManageCmd, "convertfromraw", "stdin", dest,
		fmt.Sprintf("%d", sizeBytes), "--format", "VMDK")

	log.Debug(cmd)

	if os.Getenv("MACHINE_DEBUG") != "" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}

	log.Debug("Starting command")

	if err := cmd.Start(); err != nil {
		return err
	}

	log.Debug("Copying to stdin")

	n, err := io.Copy(stdin, r)
	if err != nil {
		return err
	}

	log.Debug("Filling zeroes")

	// The total number of bytes written to stdin must match sizeBytes, or
	// VBoxManage.exe on Windows will fail. Fill remaining with zeros.
	if left := sizeBytes - n; left > 0 {
		if err := zeroFill(stdin, left); err != nil {
			return err
		}
	}

	log.Debug("Closing STDIN")

	// cmd won't exit until the stdin is closed.
	if err := stdin.Close(); err != nil {
		return err
	}

	log.Debug("Waiting on cmd")

	return cmd.Wait()
}

// zeroFill writes n zero bytes into w.
func zeroFill(w io.Writer, n int64) error {
	const blocksize = 32 << 10
	zeros := make([]byte, blocksize)
	var k int
	var err error
	for n > 0 {
		if n > blocksize {
			k, err = w.Write(zeros)
		} else {
			k, err = w.Write(zeros[:n])
		}
		if err != nil {
			return err
		}
		n -= int64(k)
	}
	return nil
}

func getVMDiskInfo(name string, vbox VBoxManager) (*VirtualDisk, error) {
	out, err := vbox.vbmOut("showvminfo", name, "--machinereadable")
	if err != nil {
		return nil, err
	}

	disk := &VirtualDisk{}

	err = parseKeyValues(out, reEqualQuoteLine, func(key, val string) error {
		switch key {
		case "SATA-1-0":
			disk.Path = val
		case "SATA-ImageUUID-1-0":
			disk.UUID = val
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return disk, nil
}
