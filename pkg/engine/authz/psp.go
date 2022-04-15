package authz

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

	"k8s.io/client-go/transport"

	"github.com/bhojpur/host/pkg/cluster/log"
	"github.com/bhojpur/host/pkg/engine/k8s"
	"github.com/bhojpur/host/pkg/engine/templates"
)

func ApplyDefaultPodSecurityPolicy(ctx context.Context, kubeConfigPath string, k8sWrapTransport transport.WrapperFunc) error {
	log.Infof(ctx, "[authz] Applying default PodSecurityPolicy")
	k8sClient, err := k8s.NewClient(kubeConfigPath, k8sWrapTransport)
	if err != nil {
		return err
	}
	if err := k8s.UpdatePodSecurityPolicyFromYaml(k8sClient, templates.DefaultPodSecurityPolicy); err != nil {
		return err
	}
	log.Infof(ctx, "[authz] Default PodSecurityPolicy applied successfully")
	return nil
}

func ApplyDefaultPodSecurityPolicyRole(ctx context.Context, kubeConfigPath, namespace string, k8sWrapTransport transport.WrapperFunc) error {
	log.Infof(ctx, "[authz] Applying default PodSecurityPolicy Role and RoleBinding in %s", namespace)
	k8sClient, err := k8s.NewClient(kubeConfigPath, k8sWrapTransport)
	if err != nil {
		return err
	}
	if err := k8s.UpdateRoleFromYaml(k8sClient, templates.DefaultPodSecurityRole, namespace); err != nil {
		return err
	}
	if err := k8s.UpdateRoleBindingFromYaml(k8sClient, templates.DefaultPodSecurityRoleBinding, namespace); err != nil {
		return err
	}
	log.Infof(ctx, "[authz] Default PodSecurityPolicy Role and RoleBinding applied successfully")
	return nil
}
