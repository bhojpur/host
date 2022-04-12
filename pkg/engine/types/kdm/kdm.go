package kdm

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
	"encoding/json"

	v3 "github.com/bhojpur/host/pkg/engine/types"
)

const (
	Calico        = "calico"
	Canal         = "canal"
	Flannel       = "flannel"
	Weave         = "weave"
	Aci           = "aci"
	CoreDNS       = "coreDNS"
	KubeDNS       = "kubeDNS"
	MetricsServer = "metricsServer"
	NginxIngress  = "nginxIngress"
	Nodelocal     = "nodelocal"
	TemplateKeys  = "templateKeys"
)

// +k8s:deepcopy-gen=false

type Data struct {
	// K8sVersionServiceOptions - service options per k8s version
	K8sVersionServiceOptions  map[string]v3.KubernetesServicesOptions
	K8sVersionBKESystemImages map[string]v3.BKESystemImages

	// Addon Templates per K8s version ("default" where nothing changes for k8s version)
	K8sVersionedTemplates map[string]map[string]string

	// K8sVersionInfo - min/max BKE+Bhojur versions per k8s version
	K8sVersionInfo map[string]v3.K8sVersionInfo

	//Default K8s version for every Bhojpur version
	BhojpurDefaultK8sVersions map[string]string

	//Default K8s version for every bke version
	BKEDefaultK8sVersions map[string]string

	K8sVersionDockerInfo map[string][]string

	// K8sVersionWindowsServiceOptions - service options per windows k8s version
	K8sVersionWindowsServiceOptions map[string]v3.KubernetesServicesOptions

	CisConfigParams         map[string]CisConfigParams
	CisBenchmarkVersionInfo map[string]CisBenchmarkVersionInfo

	// DCP specific data, opaque and defined by the config file in kdm
	DCP map[string]interface{} `json:"dcp,omitempty"`
	// UKE specific data, defined by the config file in kdm
	UKE map[string]interface{} `json:"uke,omitempty"`
}

func FromData(b []byte) (Data, error) {
	d := &Data{}

	if err := json.Unmarshal(b, d); err != nil {
		return Data{}, err
	}
	return *d, nil
}

type CisBenchmarkVersionInfo struct {
	Managed              bool              `yaml:"managed" json:"managed"`
	MinKubernetesVersion string            `yaml:"min_kubernetes_version" json:"minKubernetesVersion"`
	SkippedChecks        map[string]string `yaml:"skipped_checks" json:"skippedChecks"`
	NotApplicableChecks  map[string]string `yaml:"not_applicable_checks" json:"notApplicableChecks"`
}

type CisConfigParams struct {
	BenchmarkVersion string `yaml:"benchmark_version" json:"benchmarkVersion"`
}
