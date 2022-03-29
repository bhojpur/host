package commands

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
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMountCmd(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		ip:          "12.34.56.78",
		sshPort:     234,
		sshUsername: "root",
		sshKeyPath:  "/fake/keypath/id_rsa",
	}}

	path, err := exec.LookPath("sshfs")
	if err != nil {
		t.Skip("sshfs not found (install sshfs ?)")
	}
	cmd, err := getMountCmd("myfunhost:/home/bhojpur/foo", "/tmp/foo", false, &hostInfoLoader)

	expectedArgs := append(
		baseSSHFSArgs,
		"-o",
		"IdentitiesOnly=yes",
		"-o",
		"Port=234",
		"-o",
		"IdentityFile=/fake/keypath/id_rsa",
		"root@12.34.56.78:/home/bhojpur/foo",
		"/tmp/foo",
	)
	expectedCmd := exec.Command(path, expectedArgs...)

	assert.Equal(t, expectedCmd, cmd)
	assert.NoError(t, err)
}

func TestGetMountCmdWithoutSshKey(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		ip:          "1.2.3.4",
		sshUsername: "user",
	}}

	path, err := exec.LookPath("sshfs")
	if err != nil {
		t.Skip("sshfs not found (install sshfs ?)")
	}
	cmd, err := getMountCmd("myfunhost:/home/bhojpur/foo", "", false, &hostInfoLoader)

	expectedArgs := append(
		baseSSHFSArgs,
		"user@1.2.3.4:/home/bhojpur/foo",
		"/home/bhojpur/foo",
	)
	expectedCmd := exec.Command(path, expectedArgs...)

	assert.Equal(t, expectedCmd, cmd)
	assert.NoError(t, err)
}

func TestGetMountCmdUnmount(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		ip:          "1.2.3.4",
		sshUsername: "user",
	}}

	path, err := exec.LookPath("fusermount")
	if err != nil {
		t.Skip("fusermount not found (install fuse ?)")
	}
	cmd, err := getMountCmd("myfunhost:/home/bhojpur/foo", "/tmp/foo", true, &hostInfoLoader)

	expectedArgs := []string{
		"-u",
		"/tmp/foo",
	}
	expectedCmd := exec.Command(path, expectedArgs...)

	assert.Equal(t, expectedCmd, cmd)
	assert.NoError(t, err)
}
