package softlayer

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
	"io/ioutil"
	"os"
	"testing"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	ctest "github.com/bhojpur/host/cmd/machine/commands/test"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/stretchr/testify/assert"
)

const (
	testStoreDir          = ".store-test"
	machineTestName       = "test-host"
	machineTestCaCert     = "test-cert"
	machineTestPrivateKey = "test-key"
)

func cleanup() error {
	return os.RemoveAll(testStoreDir)
}

func getTestStorePath() (string, error) {
	tmpDir, err := ioutil.TempDir("", "machine-test-")
	if err != nil {
		return "", err
	}
	mdirs.BaseDir = tmpDir
	return tmpDir, nil
}

func getDefaultTestDriverFlags() *ctest.FakeFlagger {
	return &ctest.FakeFlagger{
		Data: map[string]interface{}{
			"name":                   "test",
			"url":                    "unix:///var/run/bhojpur.sock",
			"softlayer-api-key":      "12345",
			"softlayer-user":         "abcdefg",
			"softlayer-api-endpoint": "https://api.softlayer.com/rest/v3",
			"softlayer-image":        "MY_TEST_IMAGE",
		},
	}
}

func getTestDriver() (*Driver, error) {
	storePath, err := getTestStorePath()
	if err != nil {
		return nil, err
	}
	defer cleanup()

	d := NewDriver(machineTestName, storePath)
	d.SetConfigFromFlags(getDefaultTestDriverFlags())
	drv := d.(*Driver)
	return drv, nil
}

func TestSetConfigFromFlagsSetsImage(t *testing.T) {
	d, err := getTestDriver()

	if assert.NoError(t, err) {
		assert.Equal(t, "MY_TEST_IMAGE", d.deviceConfig.Image)
	}
}

func TestHostnameDefaultsToMachineName(t *testing.T) {
	d, err := getTestDriver()
	if assert.NoError(t, err) {
		assert.Equal(t, machineTestName, d.deviceConfig.Hostname)
	}
}

func TestSetConfigFromFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"softlayer-api-key":      "KEY",
			"softlayer-user":         "user",
			"softlayer-api-endpoint": "ENDPOINT",
			"softlayer-domain":       "DOMAIN",
			"softlayer-region":       "REGION",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)

	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)
}
