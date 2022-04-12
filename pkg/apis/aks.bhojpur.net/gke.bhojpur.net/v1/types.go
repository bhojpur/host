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

type GKEClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GKEClusterConfigSpec   `json:"spec"`
	Status GKEClusterConfigStatus `json:"status"`
}

// GKEClusterConfigSpec is the spec for a GKEClusterConfig resource
type GKEClusterConfigSpec struct {
	Region                         string                             `json:"region" bhojpur:"noupdate"`
	Zone                           string                             `json:"zone" bhojpur:"noupdate"`
	Imported                       bool                               `json:"imported" bhojpur:"noupdate"`
	Description                    string                             `json:"description"`
	Labels                         map[string]string                  `json:"labels"`
	EnableKubernetesAlpha          *bool                              `json:"enableKubernetesAlpha"`
	ClusterAddons                  *GKEClusterAddons                  `json:"clusterAddons"`
	ClusterIpv4CidrBlock           *string                            `json:"clusterIpv4Cidr" bhojpur:"pointer"`
	ProjectID                      string                             `json:"projectID"`
	GoogleCredentialSecret         string                             `json:"googleCredentialSecret"`
	ClusterName                    string                             `json:"clusterName"`
	KubernetesVersion              *string                            `json:"kubernetesVersion" bhojpur:"pointer"`
	LoggingService                 *string                            `json:"loggingService" bhojpur:"pointer"`
	MonitoringService              *string                            `json:"monitoringService" bhojpur:"pointer"`
	NodePools                      []GKENodePoolConfig                `json:"nodePools"`
	Network                        *string                            `json:"network,omitempty" bhojpur:"pointer"`
	Subnetwork                     *string                            `json:"subnetwork,omitempty" bhojpur:"pointer"`
	NetworkPolicyEnabled           *bool                              `json:"networkPolicyEnabled,omitempty"`
	PrivateClusterConfig           *GKEPrivateClusterConfig           `json:"privateClusterConfig,omitempty"`
	IPAllocationPolicy             *GKEIPAllocationPolicy             `json:"ipAllocationPolicy,omitempty"`
	MasterAuthorizedNetworksConfig *GKEMasterAuthorizedNetworksConfig `json:"masterAuthorizedNetworks,omitempty"`
	Locations                      []string                           `json:"locations"`
	MaintenanceWindow              *string                            `json:"maintenanceWindow,omitempty" bhojpur:"pointer"`
}

type GKEIPAllocationPolicy struct {
	ClusterIpv4CidrBlock       string `json:"clusterIpv4CidrBlock,omitempty"`
	ClusterSecondaryRangeName  string `json:"clusterSecondaryRangeName,omitempty"`
	CreateSubnetwork           bool   `json:"createSubnetwork,omitempty"`
	NodeIpv4CidrBlock          string `json:"nodeIpv4CidrBlock,omitempty"`
	ServicesIpv4CidrBlock      string `json:"servicesIpv4CidrBlock,omitempty"`
	ServicesSecondaryRangeName string `json:"servicesSecondaryRangeName,omitempty"`
	SubnetworkName             string `json:"subnetworkName,omitempty"`
	UseIPAliases               bool   `json:"useIpAliases,omitempty"`
}

type GKEPrivateClusterConfig struct {
	EnablePrivateEndpoint bool   `json:"enablePrivateEndpoint,omitempty"`
	EnablePrivateNodes    bool   `json:"enablePrivateNodes,omitempty"`
	MasterIpv4CidrBlock   string `json:"masterIpv4CidrBlock,omitempty"`
}

type GKEClusterConfigStatus struct {
	Phase          string `json:"phase"`
	FailureMessage string `json:"failureMessage"`
}

type GKEClusterAddons struct {
	HTTPLoadBalancing        bool `json:"httpLoadBalancing,omitempty"`
	HorizontalPodAutoscaling bool `json:"horizontalPodAutoscaling,omitempty"`
	NetworkPolicyConfig      bool `json:"networkPolicyConfig,omitempty"`
}

type GKENodePoolConfig struct {
	Autoscaling       *GKENodePoolAutoscaling `json:"autoscaling,omitempty"`
	Config            *GKENodeConfig          `json:"config,omitempty"`
	InitialNodeCount  *int64                  `json:"initialNodeCount,omitempty"`
	MaxPodsConstraint *int64                  `json:"maxPodsConstraint,omitempty"`
	Name              *string                 `json:"name,omitempty" bhojpur:"pointer"`
	Version           *string                 `json:"version,omitempty" bhojpur:"pointer"`
	Management        *GKENodePoolManagement  `json:"management,omitempty"`
}

type GKENodePoolAutoscaling struct {
	Enabled      bool  `json:"enabled,omitempty"`
	MaxNodeCount int64 `json:"maxNodeCount,omitempty"`
	MinNodeCount int64 `json:"minNodeCount,omitempty"`
}

type GKENodeConfig struct {
	DiskSizeGb    int64                `json:"diskSizeGb,omitempty"`
	DiskType      string               `json:"diskType,omitempty"`
	ImageType     string               `json:"imageType,omitempty"`
	Labels        map[string]string    `json:"labels,omitempty"`
	LocalSsdCount int64                `json:"localSsdCount,omitempty"`
	MachineType   string               `json:"machineType,omitempty"`
	OauthScopes   []string             `json:"oauthScopes,omitempty"`
	Preemptible   bool                 `json:"preemptible,omitempty"`
	Tags          []string             `json:"tags,omitempty"`
	Taints        []GKENodeTaintConfig `json:"taints,omitempty"`
}

type GKENodeTaintConfig struct {
	Effect string `json:"effect,omitempty"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
}

type GKEMasterAuthorizedNetworksConfig struct {
	CidrBlocks []*GKECidrBlock `json:"cidrBlocks,omitempty"`
	Enabled    bool            `json:"enabled,omitempty"`
}

type GKECidrBlock struct {
	CidrBlock   string `json:"cidrBlock,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type GKENodePoolManagement struct {
	AutoRepair  bool `json:"autoRepair,omitempty"`
	AutoUpgrade bool `json:"autoUpgrade,omitempty"`
}
