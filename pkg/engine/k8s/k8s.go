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
	"bytes"
	"time"

	"k8s.io/client-go/transport"

	yamlutil "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

const (
	DefaultRetries          = 5
	DefaultSleepSeconds     = 5
	DefaultTimeout          = 45
	K8sWrapTransportTimeout = 30
)

type k8sCall func(*kubernetes.Clientset, interface{}) error

func NewClient(kubeConfigPath string, k8sWrapTransport transport.WrapperFunc) (*kubernetes.Clientset, error) {
	// use the current admin kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
	if err != nil {
		return nil, err
	}
	if k8sWrapTransport != nil {
		config.WrapTransport = k8sWrapTransport
	}
	config.Timeout = time.Second * time.Duration(K8sWrapTransportTimeout)
	K8sClientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return K8sClientSet, nil
}

func DecodeYamlResource(resource interface{}, yamlManifest string) error {
	decoder := yamlutil.NewYAMLToJSONDecoder(bytes.NewReader([]byte(yamlManifest)))
	return decoder.Decode(&resource)
}

func retryToWithTimeout(runFunc k8sCall, k8sClient *kubernetes.Clientset, resource interface{}, timeout int) error {
	var err error
	timePassed := 0
	for timePassed < timeout {
		if err = runFunc(k8sClient, resource); err != nil {
			time.Sleep(time.Second * time.Duration(DefaultSleepSeconds))
			timePassed += DefaultSleepSeconds
			continue
		}
		return nil
	}
	return err
}

func retryTo(runFunc k8sCall, k8sClient *kubernetes.Clientset, resource interface{}, retries, sleepSeconds int) error {
	var err error
	if retries == 0 {
		retries = DefaultRetries
	}
	if sleepSeconds == 0 {
		sleepSeconds = DefaultSleepSeconds
	}
	for i := 0; i < retries; i++ {
		if err = runFunc(k8sClient, resource); err != nil {
			time.Sleep(time.Second * time.Duration(sleepSeconds))
			continue
		}
		return nil
	}
	return err
}
