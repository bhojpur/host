package schema

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

type NodeInfo struct {
	CPU        CPUInfo
	Memory     MemoryInfo
	OS         OSInfo
	Kubernetes KubernetesInfo
}

type CPUInfo struct {
	Count int64
}

type MemoryInfo struct {
	MemTotalKiB int64
}

type OSInfo struct {
	DockerVersion   string
	KernelVersion   string
	OperatingSystem string
}

type KubernetesInfo struct {
	KubeletVersion   string
	KubeProxyVersion string
}

type NamespaceResourceQuota struct {
	Limit ResourceQuotaLimit `json:"limit,omitempty"`
}

type ResourceQuotaLimit struct {
	Pods                   string `json:"pods,omitempty"`
	Services               string `json:"services,omitempty"`
	ReplicationControllers string `json:"replicationControllers,omitempty"`
	Secrets                string `json:"secrets,omitempty"`
	ConfigMaps             string `json:"configMaps,omitempty"`
	PersistentVolumeClaims string `json:"persistentVolumeClaims,omitempty"`
	ServicesNodePorts      string `json:"servicesNodePorts,omitempty"`
	ServicesLoadBalancers  string `json:"servicesLoadBalancers,omitempty"`
	RequestsCPU            string `json:"requestsCpu,omitempty"`
	RequestsMemory         string `json:"requestsMemory,omitempty"`
	RequestsStorage        string `json:"requestsStorage,omitempty"`
	LimitsCPU              string `json:"limitsCpu,omitempty"`
	LimitsMemory           string `json:"limitsMemory,omitempty"`
}

type NamespaceMove struct {
	ProjectID string `json:"projectId,omitempty"`
}

type ContainerResourceLimit struct {
	RequestsCPU    string `json:"requestsCpu,omitempty"`
	RequestsMemory string `json:"requestsMemory,omitempty"`
	LimitsCPU      string `json:"limitsCpu,omitempty"`
	LimitsMemory   string `json:"limitsMemory,omitempty"`
}
