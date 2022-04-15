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

	"github.com/bhojpur/host/pkg/cluster/log"
	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/pki"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/sirupsen/logrus"
)

func runKubeAPI(ctx context.Context, host *hosts.Host, df hosts.DialerFactory, prsMap map[string]v3.PrivateRegistry, kubeAPIProcess v3.Process, alpineImage string, certMap map[string]pki.CertificatePKI, k8sVersion string) error {
	imageCfg, hostCfg, healthCheckURL := GetProcessConfig(kubeAPIProcess, host, k8sVersion)
	if err := docker.DoRunContainer(ctx, host.DClient, imageCfg, hostCfg, KubeAPIContainerName, host.Address, ControlRole, prsMap); err != nil {
		return err
	}
	if err := runHealthcheck(ctx, host, KubeAPIContainerName, df, healthCheckURL, certMap); err != nil {
		return err
	}
	return createLogLink(ctx, host, KubeAPIContainerName, ControlRole, alpineImage, prsMap)
}

func removeKubeAPI(ctx context.Context, host *hosts.Host) error {
	return docker.DoRemoveContainer(ctx, host.DClient, KubeAPIContainerName, host.Address)
}

func RestartKubeAPI(ctx context.Context, host *hosts.Host) error {
	return docker.DoRestartContainer(ctx, host.DClient, KubeAPIContainerName, host.Address)
}

func RestartKubeAPIWithHealthcheck(ctx context.Context, hostList []*hosts.Host, df hosts.DialerFactory, certMap map[string]pki.CertificatePKI) error {
	log.Infof(ctx, "[%s] Restarting %s on %s nodes..", ControlRole, KubeAPIContainerName, ControlRole)
	for _, runHost := range hostList {
		logrus.Debugf("[%s] Restarting %s on node [%s]", ControlRole, KubeAPIContainerName, runHost.Address)
		if err := RestartKubeAPI(ctx, runHost); err != nil {
			return err
		}
		logrus.Debugf("[%s] Running healthcheck for %s on node [%s]", ControlRole, KubeAPIContainerName, runHost.Address)
		if err := runHealthcheck(ctx, runHost, KubeAPIContainerName,
			df, GetHealthCheckURL(true, KubeAPIPort),
			certMap); err != nil {
			return err
		}
	}
	log.Infof(ctx, "[%s] Restarted %s on %s nodes successfully", ControlRole, KubeAPIContainerName, ControlRole)
	return nil
}
