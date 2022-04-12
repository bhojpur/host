package kubeconfig

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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
)

func GetNonInteractiveClientConfig(kubeConfig string) clientcmd.ClientConfig {
	return GetClientConfig(kubeConfig, nil)
}

func GetNonInteractiveClientConfigWithContext(kubeConfig, currentContext string) clientcmd.ClientConfig {
	return GetClientConfigWithContext(kubeConfig, currentContext, nil)
}

func GetInteractiveClientConfig(kubeConfig string) clientcmd.ClientConfig {
	return GetClientConfig(kubeConfig, os.Stdin)
}

func GetClientConfigWithContext(kubeConfig, currentContext string, reader io.Reader) clientcmd.ClientConfig {
	loadingRules := GetLoadingRules(kubeConfig)
	overrides := &clientcmd.ConfigOverrides{ClusterDefaults: clientcmd.ClusterDefaults, CurrentContext: currentContext}
	return clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, overrides, reader)
}

func GetClientConfig(kubeConfig string, reader io.Reader) clientcmd.ClientConfig {
	return GetClientConfigWithContext(kubeConfig, "", reader)
}

func GetLoadingRules(kubeConfig string) *clientcmd.ClientConfigLoadingRules {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	if kubeConfig != "" {
		loadingRules.ExplicitPath = kubeConfig
	}

	var otherFiles []string
	homeDir, err := os.UserHomeDir()
	if err == nil {
		otherFiles = append(otherFiles, filepath.Join(homeDir, ".kube", "dcp.yaml"))
	}
	otherFiles = append(otherFiles, "/etc/bhojpur/dcp/dcp.yaml")
	loadingRules.Precedence = append(loadingRules.Precedence, canRead(otherFiles)...)

	return loadingRules
}

func canRead(files []string) (result []string) {
	for _, f := range files {
		_, err := ioutil.ReadFile(f)
		if err == nil {
			result = append(result, f)
		}
	}
	return
}
