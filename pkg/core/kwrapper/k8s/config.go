package k8s

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

	"github.com/bhojpur/host/pkg/common/kubeconfig"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func GetConfig(ctx context.Context, k8sMode string, kubeConfig string) (bool, clientcmd.ClientConfig, error) {
	var (
		cfg clientcmd.ClientConfig
		err error
	)

	switch k8sMode {
	case "auto":
		return getAuto(ctx, kubeConfig)
	case "embedded":
		return getEmbedded(ctx)
	case "external":
		cfg = getExternal(kubeConfig)
	default:
		return false, nil, fmt.Errorf("invalid Kubernetes mode %s", k8sMode)
	}

	return false, cfg, err
}

func getAuto(ctx context.Context, kubeConfig string) (bool, clientcmd.ClientConfig, error) {
	if isManual(kubeConfig) {
		return false, kubeconfig.GetNonInteractiveClientConfig(kubeConfig), nil
	}

	return getEmbedded(ctx)
}

func isManual(kubeConfig string) bool {
	if kubeConfig != "" {
		return true
	}
	_, inClusterErr := rest.InClusterConfig()
	return inClusterErr == nil
}

func getExternal(kubeConfig string) clientcmd.ClientConfig {
	return kubeconfig.GetNonInteractiveClientConfig(kubeConfig)
}
