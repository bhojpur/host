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
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"

	"github.com/bhojpur/host/pkg/engine/docker"
	"github.com/bhojpur/host/pkg/engine/hosts"
	v3 "github.com/bhojpur/host/pkg/engine/types"
	"github.com/bhojpur/host/pkg/engine/util"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

const (
	SELinuxCheckContainer = "bke-selinux-checker"
)

func (c *Cluster) RunSELinuxCheck(ctx context.Context) error {
	// We only need to check this on k8s 1.22 and higher
	matchedRange, err := util.SemVerMatchRange(c.Version, util.SemVerK8sVersion122OrHigher)
	if err != nil {
		return err
	}

	if matchedRange {
		var errgrp errgroup.Group
		allHosts := hosts.GetUniqueHostList(c.EtcdHosts, c.ControlPlaneHosts, c.WorkerHosts)
		hostsQueue := util.GetObjectQueue(allHosts)
		for w := 0; w < WorkerThreads; w++ {
			errgrp.Go(func() error {
				var errList []error
				for host := range hostsQueue {
					if hosts.IsDockerSELinuxEnabled(host.(*hosts.Host)) {
						err := checkSELinuxLabelOnHost(ctx, host.(*hosts.Host), c.SystemImages.Alpine, c.PrivateRegistriesMap)
						if err != nil {
							errList = append(errList, err)
						}
					}
				}
				return util.ErrList(errList)
			})
		}
		if err := errgrp.Wait(); err != nil {
			return err
		}
	}
	return nil
}

func checkSELinuxLabelOnHost(ctx context.Context, host *hosts.Host, image string, prsMap map[string]v3.PrivateRegistry) error {
	var err error
	imageCfg := &container.Config{
		Image: image,
	}
	hostCfg := &container.HostConfig{
		SecurityOpt: []string{SELinuxLabel},
	}
	for retries := 0; retries < 3; retries++ {
		logrus.Infof("[selinux] Checking if host [%s] recognizes SELinux label [%s], try #%d", host.Address, SELinuxLabel, retries+1)
		if err = docker.DoRemoveContainer(ctx, host.DClient, SELinuxCheckContainer, host.Address); err != nil {
			return err
		}
		if err = docker.DoRunOnetimeContainer(ctx, host.DClient, imageCfg, hostCfg, SELinuxCheckContainer, host.Address, "selinux", prsMap); err != nil {
			// If we hit the error that indicates that the bhojpur-selinux RPM package is not installed (SELinux label is not recognized), we immediately return
			// Else we keep trying as there might be an error with Docker (slow system for example)
			if strings.Contains(err.Error(), "invalid argument") {
				return fmt.Errorf("[selinux] Host [%s] does not recognize SELinux label [%s]. This is required for Kubernetes version [%s]. Please install bhojpur-selinux RPM package and try again", host.Address, SELinuxLabel, util.SemVerK8sVersion122OrHigher)
			}
			continue
		}
		return nil
	}
	if err != nil {
		return fmt.Errorf("[selinux] Host [%s] was not able to correctly perform SELinux label check: [%v]", host.Address, err)
	}
	return nil
}
