package hosttest

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
	"github.com/bhojpur/host/pkg/core/host"
	"github.com/bhojpur/host/pkg/core/swarm"
	"github.com/bhojpur/host/pkg/core/version"
	"github.com/bhojpur/host/pkg/drivers/none"
)

const (
	DefaultHostName    = "test-host"
	HostTestCaCert     = "test-cert"
	HostTestPrivateKey = "test-key"
)

type DriverOptionsMock struct {
	Data map[string]interface{}
}

func (d DriverOptionsMock) String(key string) string {
	return d.Data[key].(string)
}

func (d DriverOptionsMock) StringSlice(key string) []string {
	return d.Data[key].([]string)
}

func (d DriverOptionsMock) Int(key string) int {
	return d.Data[key].(int)
}

func (d DriverOptionsMock) Bool(key string) bool {
	return d.Data[key].(bool)
}

func GetTestDriverFlags() *DriverOptionsMock {
	flags := &DriverOptionsMock{
		Data: map[string]interface{}{
			"name":            DefaultHostName,
			"url":             "unix:///var/run/bhojpur.sock",
			"swarm":           false,
			"swarm-host":      "",
			"swarm-master":    false,
			"swarm-discovery": "",
		},
	}
	return flags
}

func GetDefaultTestHost() (*host.Host, error) {
	hostOptions := &host.Options{
		EngineOptions: &engine.Options{},
		SwarmOptions:  &swarm.Options{},
		AuthOptions: &auth.Options{
			CaCertPath:       HostTestCaCert,
			CaPrivateKeyPath: HostTestPrivateKey,
		},
	}

	driver := none.NewDriver(DefaultHostName, "/tmp/artifacts")

	host := &host.Host{
		ConfigVersion: version.ConfigVersion,
		Name:          DefaultHostName,
		Driver:        driver,
		DriverName:    "none",
		HostOptions:   hostOptions,
	}

	flags := GetTestDriverFlags()
	if err := host.Driver.SetConfigFromFlags(flags); err != nil {
		return nil, err
	}

	return host, nil
}
