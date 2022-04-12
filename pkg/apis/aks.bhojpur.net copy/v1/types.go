package v1

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AKSClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AKSClusterConfigSpec   `json:"spec"`
	Status AKSClusterConfigStatus `json:"status"`
}

// AKSClusterConfigSpec is the spec for a AKSClusterConfig resource
type AKSClusterConfigSpec struct {
	Imported                    bool              `json:"imported" bhojpur:"noupdate"`
	ResourceLocation            string            `json:"resourceLocation" bhojpur:"noupdate"`
	ResourceGroup               string            `json:"resourceGroup" bhojpur:"noupdate"`
	ClusterName                 string            `json:"clusterName" bhojpur:"noupdate"`
	AzureCredentialSecret       string            `json:"azureCredentialSecret"`
	BaseURL                     *string           `json:"baseUrl" bhojpur:"pointer"`
	AuthBaseURL                 *string           `json:"authBaseUrl" bhojpur:"pointer"`
	NetworkPlugin               *string           `json:"networkPlugin" bhojpur:"pointer"`
	VirtualNetworkResourceGroup *string           `json:"virtualNetworkResourceGroup" bhojpur:"pointer"`
	VirtualNetwork              *string           `json:"virtualNetwork" bhojpur:"pointer"`
	Subnet                      *string           `json:"subnet" bhojpur:"pointer"`
	NetworkDNSServiceIP         *string           `json:"dnsServiceIp" bhojpur:"pointer"`
	NetworkServiceCIDR          *string           `json:"serviceCidr" bhojpur:"pointer"`
	NetworkDockerBridgeCIDR     *string           `json:"dockerBridgeCidr" bhojpur:"pointer"`
	NetworkPodCIDR              *string           `json:"podCidr" bhojpur:"pointer"`
	LoadBalancerSKU             *string           `json:"loadBalancerSku" bhojpur:"pointer"`
	NetworkPolicy               *string           `json:"networkPolicy" bhojpur:"pointer"`
	LinuxAdminUsername          *string           `json:"linuxAdminUsername,omitempty" bhojpur:"pointer"`
	LinuxSSHPublicKey           *string           `json:"sshPublicKey,omitempty" bhojpur:"pointer"`
	DNSPrefix                   *string           `json:"dnsPrefix,omitempty" bhojpur:"pointer"`
	KubernetesVersion           *string           `json:"kubernetesVersion" bhojpur:"pointer"`
	Tags                        map[string]string `json:"tags"`
	NodePools                   []AKSNodePool     `json:"nodePools"`
	PrivateCluster              *bool             `json:"privateCluster"`
	AuthorizedIPRanges          *[]string         `json:"authorizedIpRanges" bhojpur:"pointer"`
	HTTPApplicationRouting      *bool             `json:"httpApplicationRouting"`
	Monitoring                  *bool             `json:"monitoring"`
	LogAnalyticsWorkspaceGroup  *string           `json:"logAnalyticsWorkspaceGroup" bhojpur:"pointer"`
	LogAnalyticsWorkspaceName   *string           `json:"logAnalyticsWorkspaceName" bhojpur:"pointer"`
}

type AKSClusterConfigStatus struct {
	Phase          string `json:"phase"`
	FailureMessage string `json:"failureMessage"`
	RBACEnabled    *bool  `json:"rbacEnabled"`
}

type AKSNodePool struct {
	Name                *string   `json:"name,omitempty" bhojpur:"pointer"`
	Count               *int32    `json:"count,omitempty"`
	MaxPods             *int32    `json:"maxPods,omitempty"`
	VMSize              string    `json:"vmSize,omitempty"`
	OsDiskSizeGB        *int32    `json:"osDiskSizeGB,omitempty"`
	OsDiskType          string    `json:"osDiskType,omitempty"`
	Mode                string    `json:"mode,omitempty"`
	OsType              string    `json:"osType,omitempty"`
	OrchestratorVersion *string   `json:"orchestratorVersion,omitempty" bhojpur:"pointer"`
	AvailabilityZones   *[]string `json:"availabilityZones,omitempty" bhojpur:"pointer"`
	MaxCount            *int32    `json:"maxCount,omitempty"`
	MinCount            *int32    `json:"minCount,omitempty"`
	EnableAutoScaling   *bool     `json:"enableAutoScaling,omitempty"`
}
