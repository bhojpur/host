package services

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

	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/hosts"
	v3 "github.com/bhojpur/host/pkg/engine/types"
)

const (
	NginxProxyImage   = "bhojpur/bke-nginx-proxy:0.1.0"
	NginxProxyEnvName = "CP_HOSTS"
)

func runNginxProxy(ctx context.Context, host *hosts.Host, prsMap map[string]v3.PrivateRegistry, proxyProcess v3.Process, alpineImage, k8sVersion string) error {
	imageCfg, hostCfg, _ := GetProcessConfig(proxyProcess, host, k8sVersion)
	if err := docker.DoRunContainer(ctx, host.DClient, imageCfg, hostCfg, NginxProxyContainerName, host.Address, WorkerRole, prsMap); err != nil {
		return err
	}
	return createLogLink(ctx, host, NginxProxyContainerName, WorkerRole, alpineImage, prsMap)
}

func removeNginxProxy(ctx context.Context, host *hosts.Host) error {
	return docker.DoRemoveContainer(ctx, host.DClient, NginxProxyContainerName, host.Address)
}

func RestartNginxProxy(ctx context.Context, host *hosts.Host) error {
	return docker.DoRestartContainer(ctx, host.DClient, NginxProxyContainerName, host.Address)
}
