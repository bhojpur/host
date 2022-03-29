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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockHostInfo struct {
	name        string
	ip          string
	sshPort     int
	sshUsername string
	sshKeyPath  string
}

func (h *MockHostInfo) GetMachineName() string {
	return h.name
}

func (h *MockHostInfo) GetSSHHostname() (string, error) {
	return h.ip, nil
}

func (h *MockHostInfo) GetSSHPort() (int, error) {
	return h.sshPort, nil
}

func (h *MockHostInfo) GetSSHUsername() string {
	return h.sshUsername
}

func (h *MockHostInfo) GetSSHKeyPath() string {
	return h.sshKeyPath
}

type MockHostInfoLoader struct {
	hostInfo MockHostInfo
}

func (l *MockHostInfoLoader) load(name string) (HostInfo, error) {
	info := l.hostInfo
	info.name = name
	return &info, nil
}

func TestGetInfoForLocalScpArg(t *testing.T) {
	host, user, path, opts, err := getInfoForScpArg("/tmp/foo", nil)
	assert.Nil(t, host)
	assert.Empty(t, user)
	assert.Equal(t, "/tmp/foo", path)
	assert.Nil(t, opts)
	assert.NoError(t, err)

	host, user, path, opts, err = getInfoForScpArg("localhost:C:\\path", nil)
	assert.Nil(t, host)
	assert.Empty(t, user)
	assert.Equal(t, "C:\\path", path)
	assert.Nil(t, opts)
	assert.NoError(t, err)
}

func TestGetInfoForRemoteScpArg(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		sshKeyPath: "/fake/keypath/id_rsa",
	}}

	host, user, path, opts, err := getInfoForScpArg("myuser@myfunhost:/home/bhojpur/foo", &hostInfoLoader)
	assert.Equal(t, "myfunhost", host.GetMachineName())
	assert.Equal(t, "myuser", user)
	assert.Equal(t, "/home/bhojpur/foo", path)
	assert.Equal(t, []string{"-o", `IdentityFile="/fake/keypath/id_rsa"`}, opts)
	assert.NoError(t, err)

	host, user, path, opts, err = getInfoForScpArg("myfunhost:C:\\path", &hostInfoLoader)
	assert.Equal(t, "myfunhost", host.GetMachineName())
	assert.Empty(t, user)
	assert.Equal(t, "C:\\path", path)
	assert.Equal(t, []string{"-o", `IdentityFile="/fake/keypath/id_rsa"`}, opts)
	assert.NoError(t, err)
}

func TestHostLocation(t *testing.T) {
	arg, err := generateLocationArg(nil, "user1", "/home/bhojpur/foo")

	assert.Equal(t, "/home/bhojpur/foo", arg)
	assert.NoError(t, err)
}

func TestRemoteLocation(t *testing.T) {
	hostInfo := MockHostInfo{
		ip:          "12.34.56.78",
		sshUsername: "root",
	}

	arg, err := generateLocationArg(&hostInfo, "", "/home/bhojpur/foo")

	assert.Equal(t, "root@12.34.56.78:/home/bhojpur/foo", arg)
	assert.NoError(t, err)

	argWithUser, err := generateLocationArg(&hostInfo, "user1", "/home/bhojpur/foo")

	assert.Equal(t, "user1@12.34.56.78:/home/bhojpur/foo", argWithUser)
	assert.NoError(t, err)
}

func TestGetScpCmd(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		ip:          "12.34.56.78",
		sshPort:     234,
		sshUsername: "root",
		sshKeyPath:  "/fake/keypath/id_rsa",
	}}

	cmd, err := getScpCmd("/tmp/foo", "myfunhost:/home/bhojpur/foo", true, false, false, &hostInfoLoader)

	expectedArgs := append(
		baseSSHArgs,
		"-3",
		"-r",
		"-o",
		"IdentitiesOnly=yes",
		"-o",
		"Port=234",
		"-o",
		`IdentityFile="/fake/keypath/id_rsa"`,
		"/tmp/foo",
		"root@12.34.56.78:/home/bhojpur/foo",
	)
	expectedCmd := exec.Command("/usr/bin/scp", expectedArgs...)

	assert.Equal(t, expectedCmd, cmd)
	assert.NoError(t, err)
}

func TestGetScpCmdWithoutSshKey(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		ip:          "1.2.3.4",
		sshUsername: "user",
	}}

	cmd, err := getScpCmd("/tmp/foo", "myfunhost:/home/bhojpur/foo", true, false, false, &hostInfoLoader)

	expectedArgs := append(
		baseSSHArgs,
		"-3",
		"-r",
		"/tmp/foo",
		"user@1.2.3.4:/home/bhojpur/foo",
	)
	expectedCmd := exec.Command("/usr/bin/scp", expectedArgs...)

	assert.Equal(t, expectedCmd, cmd)
	assert.NoError(t, err)
}

func TestGetScpCmdWithDelta(t *testing.T) {
	hostInfoLoader := MockHostInfoLoader{MockHostInfo{
		ip:          "1.2.3.4",
		sshUsername: "user",
	}}

	cmd, err := getScpCmd("/tmp/foo", "myfunhost:/home/bhojpur/foo", true, true, false, &hostInfoLoader)

	expectedArgs := append(
		[]string{"--progress"},
		"-e",
		"ssh "+strings.Join(baseSSHArgs, " "),
		"-r",
		"/tmp/foo",
		"user@1.2.3.4:/home/bhojpur/foo",
	)
	expectedCmd := exec.Command("/usr/bin/rsync", expectedArgs...)

	assert.Equal(t, expectedCmd, cmd)
	assert.NoError(t, err)
}
