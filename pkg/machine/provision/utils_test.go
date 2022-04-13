package provision

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
	"regexp"
	"strings"
	"testing"

	"github.com/bhojpur/host/pkg/drivers/fakedriver"
	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/engine"
	"github.com/bhojpur/host/pkg/machine/provision/pkgaction"
	"github.com/bhojpur/host/pkg/machine/provision/provisiontest"
	"github.com/bhojpur/host/pkg/machine/provision/serviceaction"
	"github.com/bhojpur/host/pkg/machine/swarm"
	"github.com/stretchr/testify/assert"
)

var (
	reDaemonListening = ":2376\\s+.*:.*"
)

func TestMatchNetstatOutMissing(t *testing.T) {
	nsOut := `Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN
tcp        0      0 0.0.0.0:237             0.0.0.0:*               LISTEN
tcp6       0      0 :::22                   :::*                    LISTEN
tcp6       0      0 :::23760                :::*                    LISTEN`
	if matchNetstatOut(reDaemonListening, nsOut) {
		t.Fatal("Expected not to match the netstat output as showing the daemon listening but got a match")
	}
}

func TestMatchNetstatOutPresent(t *testing.T) {
	nsOut := `Active Internet connections (servers and established)
Proto Recv-Q Send-Q Local Address           Foreign Address         State
tcp        0      0 0.0.0.0:22              0.0.0.0:*               LISTEN
tcp6       0      0 :::2376                 :::*                    LISTEN
tcp6       0      0 :::22                   :::*                    LISTEN`
	if !matchNetstatOut(reDaemonListening, nsOut) {
		t.Fatal("Expected to match the netstat output as showing the daemon listening but didn't")
	}
}

func TestMatchSsOutMissing(t *testing.T) {
	ssOut := `State      Recv-Q Send-Q Local Address:Port               Peer Address:Port              
LISTEN     0      128          *:22                       *:*                  
LISTEN     0      128         :::22                      :::*                  
LISTEN     0      128         :::23760                   :::*                  `
	if matchNetstatOut(reDaemonListening, ssOut) {
		t.Fatal("Expected not to match the ss output as showing the daemon listening but got a match")
	}
}

func TestMatchSsOutPresent(t *testing.T) {
	ssOut := `State      Recv-Q Send-Q Local Address:Port               Peer Address:Port              
LISTEN     0      128          *:22                       *:*                  
LISTEN     0      128         :::22                      :::*                  
LISTEN     0      128         :::2376                    :::*                  `
	if !matchNetstatOut(reDaemonListening, ssOut) {
		t.Fatal("Expected to match the ss output as showing the daemon listening but didn't")
	}
}

func TestGenerateBhojpurOptionsBoot2Docker(t *testing.T) {
	p := &Boot2DockerProvisioner{
		Driver: &fakedriver.Driver{},
	}
	bhojpurPort := 1234
	p.AuthOptions = auth.Options{
		CaCertRemotePath:     "/test/ca-cert",
		ServerKeyRemotePath:  "/test/server-key",
		ServerCertRemotePath: "/test/server-cert",
	}
	engineConfigPath := "/var/lib/boot2docker/profile"

	bhojpurCfg, err := p.GenerateBhojpurOptions(bhojpurPort)
	if err != nil {
		t.Fatal(err)
	}

	if bhojpurCfg.EngineOptionsPath != engineConfigPath {
		t.Fatalf("expected engine path %s; received %s", engineConfigPath, bhojpurCfg.EngineOptionsPath)
	}

	if strings.Index(bhojpurCfg.EngineOptions, fmt.Sprintf("-H tcp://0.0.0.0:%d", bhojpurPort)) == -1 {
		t.Fatalf("-H bhojpur port invalid; expected %d", bhojpurPort)
	}

	if strings.Index(bhojpurCfg.EngineOptions, fmt.Sprintf("CACERT=%s", p.AuthOptions.CaCertRemotePath)) == -1 {
		t.Fatalf("CACERT option invalid; expected %s", p.AuthOptions.CaCertRemotePath)
	}

	if strings.Index(bhojpurCfg.EngineOptions, fmt.Sprintf("SERVERKEY=%s", p.AuthOptions.ServerKeyRemotePath)) == -1 {
		t.Fatalf("SERVERKEY option invalid; expected %s", p.AuthOptions.ServerKeyRemotePath)
	}

	if strings.Index(bhojpurCfg.EngineOptions, fmt.Sprintf("SERVERCERT=%s", p.AuthOptions.ServerCertRemotePath)) == -1 {
		t.Fatalf("SERVERCERT option invalid; expected %s", p.AuthOptions.ServerCertRemotePath)
	}
}

func TestMachinePortBoot2Docker(t *testing.T) {
	p := &Boot2DockerProvisioner{
		Driver: &fakedriver.Driver{},
	}
	bhojpurPort := engine.DefaultPort
	bindURL := fmt.Sprintf("tcp://0.0.0.0:%d", bhojpurPort)
	p.AuthOptions = auth.Options{
		CaCertRemotePath:     "/test/ca-cert",
		ServerKeyRemotePath:  "/test/server-key",
		ServerCertRemotePath: "/test/server-cert",
	}

	cfg, err := p.GenerateBhojpurOptions(bhojpurPort)
	if err != nil {
		t.Fatal(err)
	}
	re := regexp.MustCompile("-H tcp://.*:(.+)")
	m := re.FindStringSubmatch(cfg.EngineOptions)
	if len(m) == 0 {
		t.Errorf("could not find port %d in engine config", bhojpurPort)
	}

	b := m[0]
	u := strings.Split(b, " ")
	url := u[1]
	url = strings.Replace(url, "'", "", -1)
	url = strings.Replace(url, "\\\"", "", -1)
	if url != bindURL {
		t.Errorf("expected url %s; received %s", bindURL, url)
	}
}

func TestMachineCustomPortBoot2Docker(t *testing.T) {
	p := &Boot2DockerProvisioner{
		Driver: &fakedriver.Driver{},
	}
	bhojpurPort := 3376
	bindURL := fmt.Sprintf("tcp://0.0.0.0:%d", bhojpurPort)
	p.AuthOptions = auth.Options{
		CaCertRemotePath:     "/test/ca-cert",
		ServerKeyRemotePath:  "/test/server-key",
		ServerCertRemotePath: "/test/server-cert",
	}

	cfg, err := p.GenerateBhojpurOptions(bhojpurPort)
	if err != nil {
		t.Fatal(err)
	}

	re := regexp.MustCompile("-H tcp://.*:(.+)")
	m := re.FindStringSubmatch(cfg.EngineOptions)
	if len(m) == 0 {
		t.Errorf("could not find port %d in engine config", bhojpurPort)
	}

	b := m[0]
	u := strings.Split(b, " ")
	url := u[1]
	url = strings.Replace(url, "'", "", -1)
	url = strings.Replace(url, "\\\"", "", -1)
	if url != bindURL {
		t.Errorf("expected url %s; received %s", bindURL, url)
	}
}

func TestUbuntuSystemdDaemonBinary(t *testing.T) {
	p := NewUbuntuSystemdProvisioner(&fakedriver.Driver{}).(*UbuntuSystemdProvisioner)
	cases := []struct {
		output, want string
	}{
		{"hostutl version 1.9.1\n", "hostutl daemon"},
		{"hostutl version 1.11.2\n", "hostutl daemon"},
		{"hostutl version 1.12.0\n", "hostd"},
		{"hostutl version 1.13.0\n", "hostd"},
	}

	sshCmder := &provisiontest.FakeSSHCommander{
		Responses: make(map[string]string),
	}
	p.SSHCommander = sshCmder

	for _, tc := range cases {
		sshCmder.Responses["hostutl --version"] = tc.output
		opts, err := p.GenerateBhojpurOptions(1234)
		if err != nil {
			t.Fatal(err)
		}
		if !strings.Contains(opts.EngineOptions, tc.want) {
			t.Fatal("incorrect Bhojpur Host daemon binary in engine options")
		}
	}
}

type fakeProvisioner struct {
	GenericProvisioner
}

func (provisioner *fakeProvisioner) Package(name string, action pkgaction.PackageAction) error {
	return nil
}

func (provisioner *fakeProvisioner) Provision(swarmOptions swarm.Options, authOptions auth.Options, engineOptions engine.Options) error {
	return nil
}

func (provisioner *fakeProvisioner) Service(name string, action serviceaction.ServiceAction) error {
	return nil
}

func (provisioner *fakeProvisioner) String() string {
	return "fake"
}

func TestDecideStorageDriver(t *testing.T) {
	var tests = []struct {
		suppliedDriver       string
		defaultDriver        string
		remoteFilesystemType string
		expectedDriver       string
	}{
		{"", "aufs", "ext4", "aufs"},
		{"", "aufs", "btrfs", "btrfs"},
		{"", "overlay", "btrfs", "overlay"},
		{"devicemapper", "aufs", "ext4", "devicemapper"},
		{"devicemapper", "aufs", "btrfs", "devicemapper"},
	}

	p := &fakeProvisioner{GenericProvisioner{
		Driver: &fakedriver.Driver{},
	}}
	for _, test := range tests {
		p.SSHCommander = provisiontest.NewFakeSSHCommander(
			provisiontest.FakeSSHCommanderOptions{
				FilesystemType: test.remoteFilesystemType,
			},
		)
		storageDriver, err := decideStorageDriver(p, test.defaultDriver, test.suppliedDriver)
		assert.NoError(t, err)
		assert.Equal(t, test.expectedDriver, storageDriver)
	}
}

func TestGetFilesystemType(t *testing.T) {
	p := &fakeProvisioner{GenericProvisioner{
		Driver: &fakedriver.Driver{},
	}}
	p.SSHCommander = &provisiontest.FakeSSHCommander{
		Responses: map[string]string{
			"stat -f -c %T /var/lib": "btrfs\n",
		},
	}
	fsType, err := getFilesystemType(p, "/var/lib")
	assert.NoError(t, err)
	assert.Equal(t, "btrfs", fsType)
}

func TestBhojpurClientVersion(t *testing.T) {
	cases := []struct {
		output, want string
	}{
		{"Bhojpur Host version 1.9.1, build a34a1d5\n", "1.9.1"},
		{"Bhojpur Host version 1.9.1\n", "1.9.1"},
		{"Bhojpur Host version 1.13.0-rc1, build deadbeef\n", "1.13.0-rc1"},
		{"Bhojpur Host version 1.13.0-dev, build deadbeef\n", "1.13.0-dev"},
	}

	sshCmder := &provisiontest.FakeSSHCommander{
		Responses: make(map[string]string),
	}

	for _, tc := range cases {
		sshCmder.Responses["hostutl --version"] = tc.output
		got, err := BhojpurClientVersion(sshCmder)
		if err != nil {
			t.Fatal(err)
		}
		if got != tc.want {
			t.Errorf("Unexpected version string from %q; got %q, want %q", tc.output, tc.want, got)
		}
	}
}