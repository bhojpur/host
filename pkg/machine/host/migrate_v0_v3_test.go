package host

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

import "testing"

var (
	v0conf = []byte(`{"DriverName":"virtualbox","Driver":{"IPAddress":"192.168.99.100","SSHUser":"bhojpur","SSHPort":53507,"MachineName":"dev","CaCertPath":"/Users/shashi.rai/.bhojpur/machine/certs/ca.pem","PrivateKeyPath":"/Users/shashi.rai/.bhojpur/machine/certs/ca-key.pem","SwarmMaster":false,"SwarmHost":"tcp://0.0.0.0:3376","SwarmDiscovery":"","CPU":-1,"Memory":1024,"DiskSize":20000,"Boot2DockerURL":"","Boot2DockerImportVM":"","HostOnlyCIDR":""},"StorePath":"/Users/shashi.rai/.bhojpur/machine/machines/dev","HostOptions":{"Driver":"","Memory":0,"Disk":0,"EngineOptions":{"ArbitraryFlags":null,"Dns":null,"GraphDir":"","Ipv6":false,"InsecureRegistry":null,"Labels":null,"LogLevel":"","StorageDriver":"","SelinuxEnabled":false,"TlsCaCert":"","TlsCert":"","TlsKey":"","TlsVerify":false,"RegistryMirror":null,"InstallURL":""},"SwarmOptions":{"IsSwarm":false,"Address":"","Discovery":"","Master":false,"Host":"tcp://0.0.0.0:3376","Image":"","Strategy":"","Heartbeat":0,"Overcommit":0,"TlsCaCert":"","TlsCert":"","TlsKey":"","TlsVerify":false,"ArbitraryFlags":null},"AuthOptions":{"StorePath":"/Users/shashi.rai/.bhojpur/machine/machines/dev","CaCertPath":"/Users/shashi.rai/.bhojpur/machine/certs/ca.pem","CaCertRemotePath":"","ServerCertPath":"/Users/shashi.rai/.bhojpur/machine/certs/server.pem","ServerKeyPath":"/Users/shashi.rai/.bhojpur/machine/certs/server-key.pem","ClientKeyPath":"/Users/shashi.rai/.bhojpur/machine/certs/key.pem","ServerCertRemotePath":"","ServerKeyRemotePath":"","PrivateKeyPath":"/Users/shashi.rai/.bhojpur/machine/certs/ca-key.pem","ClientCertPath":"/Users/shashi.rai/.bhojpur/machine/certs/cert.pem"}}}`)
)

func TestMigrateHostV0ToHostV3(t *testing.T) {
	h := &Host{}
	migratedHost, migrationPerformed, err := MigrateHost(h, v0conf)
	if err != nil {
		t.Fatalf("Error attempting to migrate host: %s", err)
	}

	if !migrationPerformed {
		t.Fatal("Expected a migration to be reported as performed but it was not")
	}

	if migratedHost.DriverName != "virtualbox" {
		t.Fatalf("Expected %q, got %q for the driver name", "virtualbox", migratedHost.DriverName)
	}
}
