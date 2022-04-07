package drivers

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

	mflag "github.com/bhojpur/host/pkg/machine/flag"
	"github.com/stretchr/testify/assert"
)

func TestIP(t *testing.T) {
	cases := []struct {
		baseDriver  *BaseDriver
		expectedIP  string
		expectedErr error
	}{
		{&BaseDriver{}, "", errors.New("IP address is not set")},
		{&BaseDriver{IPAddress: "2001:4860:0:2001::68"}, "2001:4860:0:2001::68", nil},
		{&BaseDriver{IPAddress: "192.168.0.1"}, "192.168.0.1", nil},
		{&BaseDriver{IPAddress: "::1"}, "::1", nil},
		{&BaseDriver{IPAddress: "hostname"}, "hostname", nil},
	}

	for _, c := range cases {
		ip, err := c.baseDriver.GetIP()
		assert.Equal(t, c.expectedIP, ip)
		assert.Equal(t, c.expectedErr, err)
	}
}

func TestEngineInstallUrlFlagEmpty(t *testing.T) {
	assert.False(t, EngineInstallURLFlagSet(&CheckDriverOptions{}))
}

func createDriverOptionWithEngineInstall(url string) *CheckDriverOptions {
	return &CheckDriverOptions{
		FlagsValues: map[string]interface{}{"engine-install-url": url},
		CreateFlags: []mflag.Flag{mflag.StringFlag{Name: "engine-install-url", Value: ""}},
	}
}

func TestEngineInstallUrlFlagDefault(t *testing.T) {
	options := createDriverOptionWithEngineInstall(DefaultEngineInstallURL)
	assert.False(t, EngineInstallURLFlagSet(options))
}

func TestEngineInstallUrlFlagSet(t *testing.T) {
	options := createDriverOptionWithEngineInstall("https://test.bhojpur.net")
	assert.True(t, EngineInstallURLFlagSet(options))
}
