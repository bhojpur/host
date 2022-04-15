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
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeConfig = ".kube/dcp.yaml"
)

func getEmbedded(ctx context.Context) (bool, clientcmd.ClientConfig, error) {
	var (
		err error
	)

	kubeConfig, err := dcpServer(ctx)
	if err != nil {
		return false, nil, err
	}

	os.Setenv("KUBECONFIG", kubeConfig)
	clientConfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
		&clientcmd.ClientConfigLoadingRules{ExplicitPath: kubeConfig}, &clientcmd.ConfigOverrides{})

	return true, clientConfig, nil
}

func dcpServer(ctx context.Context) (string, error) {
	cmd := exec.Command("dcp", "server",
		"--cluster-init",
		"--disable=traefik,servicelb,metrics-server,local-storage",
		"--node-name=local-node",
		"--log=./dcp.log")

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Pdeathsig: syscall.SIGKILL,
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	go func() {
		err := cmd.Run()
		logrus.Fatalf("Bhojpur DCP instance exited with: %v", err)
	}()

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	kubeConfig := filepath.Join(home, kubeConfig)

	for {
		if _, err := os.Stat(kubeConfig); err == nil {
			return kubeConfig, nil
		}
		logrus.Infof("Waiting for Bhojpur DCP instance to start")
		select {
		case <-ctx.Done():
			return "", fmt.Errorf("startup interrupted")
		case <-time.After(time.Second):
		}
	}
}
