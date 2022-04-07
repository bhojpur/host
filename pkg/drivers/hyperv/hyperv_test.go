package hyperv

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
	"testing"

	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/stretchr/testify/assert"
)

func TestSetConfigFromDefaultFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)

	sshPort, err := driver.GetSSHPort()
	assert.Equal(t, 22, sshPort)
	assert.NoError(t, err)

	assert.Equal(t, "", driver.Boot2DockerURL)
	assert.Equal(t, "", driver.VSwitch)
	assert.Equal(t, defaultDiskSize, driver.DiskSize)
	assert.Equal(t, defaultMemory, driver.MemSize)
	assert.Equal(t, defaultCPU, driver.CPU)
	assert.Equal(t, "", driver.MacAddr)
	assert.Equal(t, defaultVLanID, driver.VLanID)
	assert.Equal(t, "bhojpur", driver.GetSSHUsername())
	assert.Equal(t, defaultDisableDynamicMemory, driver.DisableDynamicMemory)
}

func TestSetConfigFromCustomFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"hyperv-boot2docker-url":        "B2D_URL",
			"hyperv-virtual-switch":         "TheSwitch",
			"hyperv-disk-size":              100000,
			"hyperv-memory":                 4096,
			"hyperv-cpu-count":              4,
			"hyperv-static-macaddress":      "00:0a:95:9d:68:16",
			"hyperv-vlan-id":                2,
			"hyperv-disable-dynamic-memory": true,
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)

	sshPort, err := driver.GetSSHPort()
	assert.Equal(t, 22, sshPort)
	assert.NoError(t, err)

	assert.Equal(t, "B2D_URL", driver.Boot2DockerURL)
	assert.Equal(t, "TheSwitch", driver.VSwitch)
	assert.Equal(t, 100000, driver.DiskSize)
	assert.Equal(t, 4096, driver.MemSize)
	assert.Equal(t, 4, driver.CPU)
	assert.Equal(t, "00:0a:95:9d:68:16", driver.MacAddr)
	assert.Equal(t, 2, driver.VLanID)
	assert.Equal(t, "bhojpur", driver.GetSSHUsername())
	assert.Equal(t, true, driver.DisableDynamicMemory)
}
