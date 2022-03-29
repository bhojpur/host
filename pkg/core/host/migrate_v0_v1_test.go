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
	"reflect"
	"testing"

	mdirs "github.com/bhojpur/host/cmd/machine/commands/dirs"
	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/engine"
	"github.com/bhojpur/host/pkg/core/swarm"
)

func TestMigrateHostV0ToV1(t *testing.T) {
	mdirs.BaseDir = "/tmp/migration"
	originalHost := &V0{
		HostOptions:    nil,
		SwarmDiscovery: "token://foobar",
		SwarmHost:      "1.2.3.4:2376",
		SwarmMaster:    true,
		CaCertPath:     "/tmp/migration/certs/ca.pem",
		PrivateKeyPath: "/tmp/migration/certs/ca-key.pem",
		ClientCertPath: "/tmp/migration/certs/cert.pem",
		ClientKeyPath:  "/tmp/migration/certs/key.pem",
		ServerCertPath: "/tmp/migration/certs/server.pem",
		ServerKeyPath:  "/tmp/migration/certs/server-key.pem",
	}
	hostOptions := &OptionsV1{
		SwarmOptions: &swarm.Options{
			Master:    true,
			Discovery: "token://foobar",
			Host:      "1.2.3.4:2376",
		},
		AuthOptions: &AuthOptionsV1{
			CaCertPath:     "/tmp/migration/certs/ca.pem",
			PrivateKeyPath: "/tmp/migration/certs/ca-key.pem",
			ClientCertPath: "/tmp/migration/certs/cert.pem",
			ClientKeyPath:  "/tmp/migration/certs/key.pem",
			ServerCertPath: "/tmp/migration/certs/server.pem",
			ServerKeyPath:  "/tmp/migration/certs/server-key.pem",
		},
		EngineOptions: &engine.Options{
			InstallURL: "https://get.docker.com",
			TLSVerify:  true,
		},
	}

	expectedHost := &V1{
		HostOptions: hostOptions,
	}

	host := MigrateHostV0ToHostV1(originalHost)

	if !reflect.DeepEqual(host, expectedHost) {
		t.Logf("\n%+v\n%+v", host, expectedHost)
		t.Logf("\n%+v\n%+v", host.HostOptions, expectedHost.HostOptions)
		t.Fatal("Expected these structs to be equal, they were different")
	}
}

func TestMigrateHostMetadataV0ToV1(t *testing.T) {
	metadata := &MetadataV0{
		HostOptions: Options{
			EngineOptions: nil,
			AuthOptions:   nil,
		},
		StorePath:      "/tmp/store",
		CaCertPath:     "/tmp/store/certs/ca.pem",
		ServerCertPath: "/tmp/store/certs/server.pem",
	}
	expectedAuthOptions := &auth.Options{
		StorePath:      "/tmp/store",
		CaCertPath:     "/tmp/store/certs/ca.pem",
		ServerCertPath: "/tmp/store/certs/server.pem",
	}

	expectedMetadata := &Metadata{
		HostOptions: Options{
			EngineOptions: &engine.Options{},
			AuthOptions:   expectedAuthOptions,
		},
	}

	m := MigrateHostMetadataV0ToHostMetadataV1(metadata)

	if !reflect.DeepEqual(m, expectedMetadata) {
		t.Logf("\n%+v\n%+v", m, expectedMetadata)
		t.Fatal("Expected these structs to be equal, they were different")
	}
}
