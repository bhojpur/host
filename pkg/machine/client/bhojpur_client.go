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
	"context"
	"fmt"

	"github.com/bhojpur/host/pkg/machine/cert"
	docker "github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

// BhojpurClient creates a Bhojpur Host client for a given host.
func BhojpurClient(bhojpurHost BhojpurHost) (*docker.Client, error) {
	url, err := bhojpurHost.URL()
	if err != nil {
		return nil, err
	}

	tlsConfig, err := cert.ReadTLSConfig(url, bhojpurHost.AuthOptions())
	if err != nil {
		return nil, fmt.Errorf("Unable to read TLS config: %s", err)
	}
	fmt.Println("read TLS config: %s", tlsConfig)

	// set DOCKER_HOST, DOCKER_API_VERSION, DOCKER_CERT_PATH, DOCKER_TLS_VERIFY
	return docker.NewClientWithOpts(docker.FromEnv)
}

// CreateContainer creates a Bhojpur Host container.
func CreateContainer(bhojpurHost BhojpurHost, config *container.Config, name string) error {
	engine, err := BhojpurClient(bhojpurHost)
	if err != nil {
		return err
	}

	imageName := config.Image
	if _, err = engine.ImagePull(context.Background(), imageName, types.ImagePullOptions{}); err != nil {
		return fmt.Errorf("Unable to pull Bhojpur Host image: %s", err)
	}

	//var authConfig *types.AuthConfig
	//containerID, err := engine.ContainerCreate(config, name, authConfig)
	cntr, err := engine.ContainerCreate(context.Background(), config, nil, nil, nil, name)
	if err != nil {
		return fmt.Errorf("Error while creating Bhojpur Host container: %s", err)
	}

	if err = engine.ContainerStart(context.Background(), cntr.ID, types.ContainerStartOptions{} /*&config.HostConfig*/); err != nil {
		return fmt.Errorf("Error while starting Bhojpur Host container: %s", err)
	}

	return nil
}
