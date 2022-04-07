package digitalocean

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

func TestSetConfigFromFlags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"digitalocean-access-token": "TOKEN",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)

	assert.NoError(t, err)
	assert.Empty(t, checkFlags.InvalidFlags)

	assert.Equal(t, driver.ResolveStorePath("id_rsa"), driver.GetSSHKeyPath())
}

func TestDefaultSSHUserAndPort(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"digitalocean-access-token": "TOKEN",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)

	sshPort, err := driver.GetSSHPort()
	assert.Equal(t, "root", driver.GetSSHUsername())
	assert.Equal(t, 22, sshPort)
	assert.NoError(t, err)
}

func TestCustomSSHUserAndPort(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"digitalocean-access-token": "TOKEN",
			"digitalocean-ssh-user":     "user",
			"digitalocean-ssh-port":     2222,
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)

	sshPort, err := driver.GetSSHPort()
	assert.Equal(t, "user", driver.GetSSHUsername())
	assert.Equal(t, 2222, sshPort)
	assert.NoError(t, err)
}

func TestSSHKeyFingerprint(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"digitalocean-access-token":        "TOKEN",
			"digitalocean-ssh-key-fingerprint": "64:51:2b:9b:8b:f0:95:3c:f9:36:4d:8b:80:a8:8f:1e",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)

	assert.Equal(t, "64:51:2b:9b:8b:f0:95:3c:f9:36:4d:8b:80:a8:8f:1e", driver.SSHKeyFingerprint)
	assert.Equal(t, "", driver.GetSSHKeyPath())
}

func TestTags(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"digitalocean-access-token": "TOKEN",
			"digitalocean-tags":         "bhojpur,swarm, no-leading-space,,",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)
	assert.Equal(t, []string{"bhojpur", "swarm", "no-leading-space"}, driver.getTags())
}

func TestTagsEmpty(t *testing.T) {
	driver := NewDriver("default", "path")

	checkFlags := &drivers.CheckDriverOptions{
		FlagsValues: map[string]interface{}{
			"digitalocean-access-token": "TOKEN",
		},
		CreateFlags: driver.GetCreateFlags(),
	}

	err := driver.SetConfigFromFlags(checkFlags)
	assert.NoError(t, err)
	assert.Nil(t, driver.getTags())
}
