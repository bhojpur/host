package client

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

	"github.com/bhojpur/host/pkg/machine/cert"
	"github.com/samalba/dockerclient"
)

// BhojpurClient creates a Bhojpur Host client for a given host.
func BhojpurClient(bhojpurHost BhojpurHost) (*dockerclient.DockerClient, error) {
	url, err := bhojpurHost.URL()
	if err != nil {
		return nil, err
	}

	tlsConfig, err := cert.ReadTLSConfig(url, bhojpurHost.AuthOptions())
	if err != nil {
		return nil, fmt.Errorf("Unable to read TLS config: %s", err)
	}

	return dockerclient.NewDockerClient(url, tlsConfig)
}

// CreateContainer creates a Bhojpur Host container.
func CreateContainer(bhojpurHost BhojpurHost, config *dockerclient.ContainerConfig, name string) error {
	engine, err := BhojpurClient(bhojpurHost)
	if err != nil {
		return err
	}

	if err = engine.PullImage(config.Image, nil); err != nil {
		return fmt.Errorf("Unable to pull Bhojpur Host image: %s", err)
	}

	var authConfig *dockerclient.AuthConfig
	containerID, err := engine.CreateContainer(config, name, authConfig)
	if err != nil {
		return fmt.Errorf("Error while creating Bhojpur Host container: %s", err)
	}

	if err = engine.StartContainer(containerID, &config.HostConfig); err != nil {
		return fmt.Errorf("Error while starting Bhojpur Host container: %s", err)
	}

	return nil
}
