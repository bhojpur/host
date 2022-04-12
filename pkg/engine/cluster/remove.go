package cluster

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

	"github.com/bhojpur/host/pkg/container/log"
	"github.com/bhojpur/host/pkg/engine/hosts"
	"github.com/bhojpur/host/pkg/engine/k8s"
	"github.com/bhojpur/host/pkg/engine/pki"
	"github.com/bhojpur/host/pkg/engine/services"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/bhojpur/host/pkg/engine/util"
	"golang.org/x/sync/errgroup"
)

func (c *Cluster) ClusterRemove(ctx context.Context) error {
	if err := c.CleanupNodes(ctx); err != nil {
		return err
	}
	c.CleanupFiles(ctx)
	return nil
}

func cleanUpHosts(ctx context.Context, cpHosts, workerHosts, etcdHosts []*hosts.Host, cleanerImage string, prsMap map[string]v3.PrivateRegistry, externalEtcd bool, k8sVersion string) error {

	uniqueHosts := hosts.GetUniqueHostList(cpHosts, workerHosts, etcdHosts)

	var errgrp errgroup.Group
	hostsQueue := util.GetObjectQueue(uniqueHosts)
	for w := 0; w < WorkerThreads; w++ {
		errgrp.Go(func() error {
			var errList []error
			for host := range hostsQueue {
				runHost := host.(*hosts.Host)
				if err := runHost.CleanUpAll(ctx, cleanerImage, prsMap, externalEtcd, k8sVersion); err != nil {
					errList = append(errList, err)
				}
			}
			return util.ErrList(errList)
		})
	}

	return errgrp.Wait()
}

func (c *Cluster) CleanupNodes(ctx context.Context) error {
	externalEtcd := false
	if len(c.Services.Etcd.ExternalURLs) > 0 {
		externalEtcd = true
	}
	// Remove Worker Plane
	if err := services.RemoveWorkerPlane(ctx, c.WorkerHosts, true); err != nil {
		return err
	}
	// Remove Contol Plane
	if err := services.RemoveControlPlane(ctx, c.ControlPlaneHosts, true); err != nil {
		return err
	}

	// Remove Etcd Plane
	if !externalEtcd {
		if err := services.RemoveEtcdPlane(ctx, c.EtcdHosts, true); err != nil {
			return err
		}
	}

	// Clean up all hosts
	return cleanUpHosts(ctx, c.ControlPlaneHosts, c.WorkerHosts, c.EtcdHosts, c.SystemImages.Alpine, c.PrivateRegistriesMap, externalEtcd, c.Version)
}

func (c *Cluster) CleanupFiles(ctx context.Context) error {
	pki.RemoveAdminConfig(ctx, c.LocalKubeConfigPath)
	RemoveStateFile(ctx, c.StateFilePath)
	return nil
}

func (c *Cluster) RemoveOldNodes(ctx context.Context) error {
	kubeClient, err := k8s.NewClient(c.LocalKubeConfigPath, c.K8sWrapTransport)
	if err != nil {
		return err
	}
	nodeList, err := k8s.GetNodeList(kubeClient)
	if err != nil {
		return err
	}
	uniqueHosts := hosts.GetUniqueHostList(c.EtcdHosts, c.ControlPlaneHosts, c.WorkerHosts)
	for _, node := range nodeList.Items {
		_, isEtcd := node.Labels[etcdRoleLabel]
		if k8s.IsNodeReady(node) && !isEtcd {
			continue
		}
		host := &hosts.Host{}
		host.HostnameOverride = node.Name
		if !hosts.IsNodeInList(host, uniqueHosts) {
			if err := k8s.DeleteNode(kubeClient, node.Name, c.CloudProvider.Name); err != nil {
				log.Warnf(ctx, "Failed to delete old node [%s] from kubernetes")
			}
		}
	}
	return nil
}
