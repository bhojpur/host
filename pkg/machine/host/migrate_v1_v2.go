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

	"github.com/bhojpur/host/pkg/machine/auth"
	"github.com/bhojpur/host/pkg/machine/drivers"
	"github.com/bhojpur/host/pkg/machine/engine"
	"github.com/bhojpur/host/pkg/machine/swarm"
)

type AuthOptionsV1 struct {
	StorePath            string
	CaCertPath           string
	CaCertRemotePath     string
	ServerCertPath       string
	ServerKeyPath        string
	ClientKeyPath        string
	ServerCertRemotePath string
	ServerKeyRemotePath  string
	PrivateKeyPath       string
	ClientCertPath       string
}

type OptionsV1 struct {
	Driver        string
	Memory        int
	Disk          int
	EngineOptions *engine.Options
	SwarmOptions  *swarm.Options
	AuthOptions   *AuthOptionsV1
}

type V1 struct {
	ConfigVersion int
	Driver        drivers.Driver
	DriverName    string
	HostOptions   *OptionsV1
	Name          string `json:"-"`
	StorePath     string
}

func MigrateHostV1ToHostV2(hostV1 *V1) *V2 {
	// Changed:  Put StorePath directly in AuthOptions (for provisioning),
	// and AuthOptions.PrivateKeyPath => AuthOptions.CaPrivateKeyPath
	// Also, CertDir has been added.

	globalStorePath := filepath.Dir(filepath.Dir(hostV1.StorePath))

	h := &V2{
		ConfigVersion: hostV1.ConfigVersion,
		Driver:        hostV1.Driver,
		Name:          hostV1.Driver.GetMachineName(),
		DriverName:    hostV1.DriverName,
		HostOptions: &Options{
			Driver:        hostV1.HostOptions.Driver,
			Memory:        hostV1.HostOptions.Memory,
			Disk:          hostV1.HostOptions.Disk,
			EngineOptions: hostV1.HostOptions.EngineOptions,
			SwarmOptions:  hostV1.HostOptions.SwarmOptions,
			AuthOptions: &auth.Options{
				CertDir:              filepath.Join(globalStorePath, "certs"),
				CaCertPath:           hostV1.HostOptions.AuthOptions.CaCertPath,
				CaPrivateKeyPath:     hostV1.HostOptions.AuthOptions.PrivateKeyPath,
				CaCertRemotePath:     hostV1.HostOptions.AuthOptions.CaCertRemotePath,
				ServerCertPath:       hostV1.HostOptions.AuthOptions.ServerCertPath,
				ServerKeyPath:        hostV1.HostOptions.AuthOptions.ServerKeyPath,
				ClientKeyPath:        hostV1.HostOptions.AuthOptions.ClientKeyPath,
				ServerCertRemotePath: hostV1.HostOptions.AuthOptions.ServerCertRemotePath,
				ServerKeyRemotePath:  hostV1.HostOptions.AuthOptions.ServerKeyRemotePath,
				ClientCertPath:       hostV1.HostOptions.AuthOptions.ClientCertPath,
				StorePath:            globalStorePath,
			},
		},
	}

	return h
}
