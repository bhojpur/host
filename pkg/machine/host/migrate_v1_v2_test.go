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

import (
	"path/filepath"
	"testing"
)

var (
	v1conf = []byte(`{
    "ConfigVersion": 1,
    "Driver": {
        "IPAddress": "192.168.99.100",
        "SSHUser": "bhojpur",
        "SSHPort": 64477,
        "MachineName": "foobar",
        "CaCertPath": "/Users/shashi.rai/.bhojpur/machine/certs/ca.pem",
        "PrivateKeyPath": "/Users/shashi.rai/.bhojpur/machine/certs/ca-key.pem",
        "SwarmMaster": false,
        "SwarmHost": "tcp://0.0.0.0:3376",
        "SwarmDiscovery": "",
        "CPU": 1,
        "Memory": 1024,
        "DiskSize": 20000,
        "Boot2DockerURL": "",
        "Boot2DockerImportVM": "",
        "HostOnlyCIDR": "192.168.99.1/24"
    },
    "DriverName": "virtualbox",
    "HostOptions": {
        "Driver": "",
        "Memory": 0,
        "Disk": 0,
        "EngineOptions": {
            "ArbitraryFlags": [],
            "Dns": null,
            "GraphDir": "",
            "Env": [],
            "Ipv6": false,
            "InsecureRegistry": [],
            "Labels": [],
            "LogLevel": "",
            "StorageDriver": "",
            "SelinuxEnabled": false,
            "TlsCaCert": "",
            "TlsCert": "",
            "TlsKey": "",
            "TlsVerify": true,
            "RegistryMirror": [],
            "InstallURL": "https://get.docker.com"
        },
        "SwarmOptions": {
            "IsSwarm": false,
            "Address": "",
            "Discovery": "",
            "Master": false,
            "Host": "tcp://0.0.0.0:3376",
            "Image": "swarm:latest",
            "Strategy": "spread",
            "Heartbeat": 0,
            "Overcommit": 0,
            "TlsCaCert": "",
            "TlsCert": "",
            "TlsKey": "",
            "TlsVerify": false,
            "ArbitraryFlags": []
        },
        "AuthOptions": {
            "StorePath": "",
            "CaCertPath": "/Users/shashi.rai/.bhojpur/machine/certs/ca.pem",
            "CaCertRemotePath": "",
            "ServerCertPath": "/Users/shashi.rai/.bhojpur/machine/machines/foobar/server.pem",
            "ServerKeyPath": "/Users/shashi.rai/.bhojpur/machine/machines/foobar/server-key.pem",
            "ClientKeyPath": "/Users/shashi.rai/.bhojpur/machine/certs/key.pem",
            "ServerCertRemotePath": "",
            "ServerKeyRemotePath": "",
            "PrivateKeyPath": "/Users/shashi.rai/.bhojpur/machine/certs/ca-key.pem",
            "ClientCertPath": "/Users/shashi.rai/.bhojpur/machine/certs/cert.pem"
        }
    },
    "StorePath": "/Users/shashi.rai/.bhojpur/machine/machines/foobar"
}`)
)

func TestMigrateHostV1ToHostV2(t *testing.T) {
	h := &Host{}
	expectedGlobalStorePath := "/Users/shashi.rai/.bhojpur/machine"
	expectedCaPrivateKeyPath := "/Users/shashi.rai/.bhojpur/machine/certs/ca-key.pem"
	migratedHost, migrationPerformed, err := MigrateHost(h, v1conf)
	if err != nil {
		t.Fatalf("Error attempting to migrate host: %s", err)
	}

	if !migrationPerformed {
		t.Fatal("Expected a migration to be reported as performed but it was not")
	}

	if migratedHost.HostOptions.AuthOptions.StorePath != expectedGlobalStorePath {
		t.Fatalf("Expected %q, got %q for the store path in AuthOptions", migratedHost.HostOptions.AuthOptions.StorePath, expectedGlobalStorePath)
	}

	if migratedHost.HostOptions.AuthOptions.CaPrivateKeyPath != expectedCaPrivateKeyPath {
		t.Fatalf("Expected %q, got %q for the private key path in AuthOptions", migratedHost.HostOptions.AuthOptions.CaPrivateKeyPath, expectedCaPrivateKeyPath)
	}

	if migratedHost.HostOptions.AuthOptions.CertDir != filepath.Join(expectedGlobalStorePath, "certs") {
		t.Fatalf("Expected %q, got %q for the cert dir in AuthOptions", migratedHost.HostOptions.AuthOptions.CaPrivateKeyPath, expectedGlobalStorePath)
	}
}
