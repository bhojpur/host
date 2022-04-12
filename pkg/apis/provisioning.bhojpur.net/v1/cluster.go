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
	bkev1 "github.com/bhojpur/host/pkg/apis/bke.bhojpur.net/v1"
	"github.com/bhojpur/host/pkg/common/genericcondition"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ClusterSpec   `json:"spec"`
	Status            ClusterStatus `json:"status,omitempty"`
}

type ClusterSpec struct {
	CloudCredentialSecretName string `json:"cloudCredentialSecretName,omitempty"`
	KubernetesVersion         string `json:"kubernetesVersion,omitempty"`

	ClusterAPIConfig         *ClusterAPIConfig              `json:"clusterAPIConfig,omitempty"`
	BKEConfig                *BKEConfig                     `json:"bkeConfig,omitempty"`
	LocalClusterAuthEndpoint bkev1.LocalClusterAuthEndpoint `json:"localClusterAuthEndpoint,omitempty"`

	AgentEnvVars                         []bkev1.EnvVar `json:"agentEnvVars,omitempty"`
	DefaultPodSecurityPolicyTemplateName string         `json:"defaultPodSecurityPolicyTemplateName,omitempty" bhojpur:"type=reference[podSecurityPolicyTemplate]"`
	DefaultClusterRoleForProjectMembers  string         `json:"defaultClusterRoleForProjectMembers,omitempty" bhojpur:"type=reference[roleTemplate]"`
	EnableNetworkPolicy                  *bool          `json:"enableNetworkPolicy,omitempty" bhojpur:"default=false"`

	RedeploySystemAgentGeneration int64 `json:"redeploySystemAgentGeneration,omitempty"`
}

type ClusterStatus struct {
	Ready              bool                                `json:"ready,omitempty"`
	ClusterName        string                              `json:"clusterName,omitempty"`
	ClientSecretName   string                              `json:"clientSecretName,omitempty"`
	AgentDeployed      bool                                `json:"agentDeployed,omitempty"`
	ObservedGeneration int64                               `json:"observedGeneration"`
	Conditions         []genericcondition.GenericCondition `json:"conditions,omitempty"`
}

type ImportedConfig struct {
	KubeConfigSecretName string `json:"kubeConfigSecretName,omitempty"`
}

type ClusterAPIConfig struct {
	ClusterName string `json:"clusterName,omitempty"`
}
