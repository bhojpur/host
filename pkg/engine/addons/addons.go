package addons

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
	"fmt"
	"strconv"

	"k8s.io/client-go/transport"

	"github.com/bhojpur/host/pkg/engine/k8s"
	"github.com/bhojpur/host/pkg/engine/templates"
	"github.com/blang/semver"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func GetAddonsExecuteJob(addonName, nodeName, image, k8sVersion string) (string, error) {
	return getAddonJob(addonName, nodeName, image, k8sVersion, false)
}

func GetAddonsDeleteJob(addonName, nodeName, image, k8sVersion string) (string, error) {
	return getAddonJob(addonName, nodeName, image, k8sVersion, true)
}

func getAddonJob(addonName, nodeName, image, k8sVersion string, isDelete bool) (string, error) {
	OSLabel := "beta.kubernetes.io/os"
	toMatch, err := semver.Make(k8sVersion[1:])
	if err != nil {
		return "", fmt.Errorf("Cluster version [%s] can not be parsed as semver: %v", k8sVersion, err)
	}

	logrus.Debugf("Checking addon job OS label for cluster version [%s]", k8sVersion)
	// kubernetes.io/os should be used 1.22.0 and up
	OSLabelRange, err := semver.ParseRange(">=1.22.0-bhojpur0")
	if err != nil {
		return "", fmt.Errorf("Failed to parse semver range for checking OS label for addon job: %v", err)
	}
	if OSLabelRange(toMatch) {
		logrus.Debugf("Cluster version [%s] needs to use new OS label", k8sVersion)
		OSLabel = "kubernetes.io/os"
	}

	jobConfig := map[string]string{
		"AddonName": addonName,
		"NodeName":  nodeName,
		"Image":     image,
		"DeleteJob": strconv.FormatBool(isDelete),
		"OSLabel":   OSLabel,
	}
	template, err := templates.CompileTemplateFromMap(templates.AddonJobTemplate, jobConfig)
	logrus.Tracef("template for [%s] is: [%s]", addonName, template)
	return template, err
}

func AddonJobExists(addonJobName, kubeConfigPath string, k8sWrapTransport transport.WrapperFunc) (bool, error) {
	k8sClient, err := k8s.NewClient(kubeConfigPath, k8sWrapTransport)
	if err != nil {
		return false, err
	}
	addonJobStatus, err := k8s.GetK8sJobStatus(k8sClient, addonJobName, metav1.NamespaceSystem)
	if err != nil {
		return false, fmt.Errorf("Failed to get job [%s] status: %v", addonJobName, err)
	}
	return addonJobStatus.Created, nil
}
