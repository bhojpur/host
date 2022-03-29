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
	"github.com/bhojpur/host/pkg/core/auth"
	"github.com/bhojpur/host/pkg/core/engine"
	"github.com/bhojpur/host/pkg/core/swarm"
)

// In the 0.1.0 => 0.2.0 transition, the JSON representation of
// machines changed from a "flat" to a more "nested" structure
// for various options and configuration settings.  To preserve
// compatibility with existing machines, these migration functions
// have been introduced.  They preserve backwards compat at the expense
// of some duplicated information.

// MigrateHostV0ToHostV1 validates host config and modifies if needed
// this is used for configuration updates
func MigrateHostV0ToHostV1(hostV0 *V0) *V1 {
	hostV1 := &V1{
		Driver:     hostV0.Driver,
		DriverName: hostV0.DriverName,
	}

	hostV1.HostOptions = &OptionsV1{}
	hostV1.HostOptions.EngineOptions = &engine.Options{
		TLSVerify:  true,
		InstallURL: "https://get.docker.com",
	}
	hostV1.HostOptions.SwarmOptions = &swarm.Options{
		Address:   "",
		Discovery: hostV0.SwarmDiscovery,
		Host:      hostV0.SwarmHost,
		Master:    hostV0.SwarmMaster,
	}
	hostV1.HostOptions.AuthOptions = &AuthOptionsV1{
		StorePath:            hostV0.StorePath,
		CaCertPath:           hostV0.CaCertPath,
		CaCertRemotePath:     "",
		ServerCertPath:       hostV0.ServerCertPath,
		ServerKeyPath:        hostV0.ServerKeyPath,
		ClientKeyPath:        hostV0.ClientKeyPath,
		ServerCertRemotePath: "",
		ServerKeyRemotePath:  "",
		PrivateKeyPath:       hostV0.PrivateKeyPath,
		ClientCertPath:       hostV0.ClientCertPath,
	}

	return hostV1
}

// MigrateHostMetadataV0ToHostMetadataV1 fills nested host metadata and modifies if needed
// this is used for configuration updates
func MigrateHostMetadataV0ToHostMetadataV1(m *MetadataV0) *Metadata {
	hostMetadata := &Metadata{}
	hostMetadata.DriverName = m.DriverName
	hostMetadata.HostOptions.EngineOptions = &engine.Options{}
	hostMetadata.HostOptions.AuthOptions = &auth.Options{
		StorePath:            m.StorePath,
		CaCertPath:           m.CaCertPath,
		CaCertRemotePath:     "",
		ServerCertPath:       m.ServerCertPath,
		ServerKeyPath:        m.ServerKeyPath,
		ClientKeyPath:        "",
		ServerCertRemotePath: "",
		ServerKeyRemotePath:  "",
		CaPrivateKeyPath:     m.PrivateKeyPath,
		ClientCertPath:       m.ClientCertPath,
	}

	hostMetadata.ConfigVersion = m.ConfigVersion

	return hostMetadata
}
