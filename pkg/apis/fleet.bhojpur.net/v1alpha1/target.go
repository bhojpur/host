package v1alpha1

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
	"github.com/bhojpur/host/pkg/common/genericcondition"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	ClusterConditionReady                  = "Ready"
	ClusterGroupAnnotation                 = "fleet.bhojpur.net/cluster-group"
	ClusterNamespaceAnnotation             = "fleet.bhojpur.net/cluster-namespace"
	ClusterAnnotation                      = "fleet.bhojpur.net/cluster"
	ClusterRegistrationAnnotation          = "fleet.bhojpur.net/cluster-registration"
	ClusterRegistrationNamespaceAnnotation = "fleet.bhojpur.net/cluster-registration-namespace"
	ManagedLabel                           = "fleet.bhojpur.net/managed"

	BootstrapToken = "fleet.bhojpur.net/bootstrap-token"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterGroupSpec   `json:"spec"`
	Status ClusterGroupStatus `json:"status"`
}

type ClusterGroupSpec struct {
	Selector *metav1.LabelSelector `json:"selector,omitempty"`
}

type ClusterGroupStatus struct {
	ClusterCount         int                                 `json:"clusterCount"`
	NonReadyClusterCount int                                 `json:"nonReadyClusterCount"`
	NonReadyClusters     []string                            `json:"nonReadyClusters,omitempty"`
	Conditions           []genericcondition.GenericCondition `json:"conditions,omitempty"`
	Summary              BundleSummary                       `json:"summary,omitempty"`
	Display              ClusterGroupDisplay                 `json:"display,omitempty"`
	ResourceCounts       GitRepoResourceCounts               `json:"resourceCounts,omitempty"`
}

type ClusterGroupDisplay struct {
	ReadyClusters string `json:"readyClusters,omitempty"`
	ReadyBundles  string `json:"readyBundles,omitempty"`
	State         string `json:"state,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

type ClusterSpec struct {
	Paused                  bool        `json:"paused,omitempty"`
	ClientID                string      `json:"clientID,omitempty"`
	KubeConfigSecret        string      `json:"kubeConfigSecret,omitempty"`
	RedeployAgentGeneration int64       `json:"redeployAgentGeneration,omitempty"`
	AgentEnvVars            []v1.EnvVar `json:"agentEnvVars,omitempty"`
	AgentNamespace          string      `json:"agentNamespace,omitempty"`
}

type ClusterStatus struct {
	Conditions           []genericcondition.GenericCondition `json:"conditions,omitempty"`
	Namespace            string                              `json:"namespace,omitempty"`
	Summary              BundleSummary                       `json:"summary,omitempty"`
	ResourceCounts       GitRepoResourceCounts               `json:"resourceCounts,omitempty"`
	ReadyGitRepos        int                                 `json:"readyGitRepos"`
	DesiredReadyGitRepos int                                 `json:"desiredReadyGitRepos"`

	AgentEnvVarsHash         string `json:"agentEnvVarsHash,omitempty"`
	AgentDeployedGeneration  *int64 `json:"agentDeployedGeneration,omitempty"`
	AgentMigrated            bool   `json:"agentMigrated,omitempty"`
	AgentNamespaceMigrated   bool   `json:"agentNamespaceMigrated,omitempty"`
	BhojpurNamespaceMigrated bool   `json:"bhojpurNamespaceMigrated,omitempty"`

	Display ClusterDisplay `json:"display,omitempty"`
	Agent   AgentStatus    `json:"agent,omitempty"`
}

type ClusterDisplay struct {
	ReadyBundles string `json:"readyBundles,omitempty"`
	ReadyNodes   string `json:"readyNodes,omitempty"`
	SampleNode   string `json:"sampleNode,omitempty"`
	State        string `json:"state,omitempty"`
}

type AgentStatus struct {
	LastSeen      metav1.Time `json:"lastSeen"`
	Namespace     string      `json:"namespace"`
	NonReadyNodes int         `json:"nonReadyNodes"`
	ReadyNodes    int         `json:"readyNodes"`
	// At most 3 nodes
	NonReadyNodeNames []string `json:"nonReadyNodeNames"`
	// At most 3 nodes
	ReadyNodeNames []string `json:"readyNodeNames"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterRegistration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRegistrationSpec   `json:"spec,omitempty"`
	Status ClusterRegistrationStatus `json:"status,omitempty"`
}

type ClusterRegistrationSpec struct {
	ClientID      string            `json:"clientID,omitempty"`
	ClientRandom  string            `json:"clientRandom,omitempty"`
	ClusterLabels map[string]string `json:"clusterLabels,omitempty"`
}

type ClusterRegistrationStatus struct {
	ClusterName string `json:"clusterName,omitempty"`
	Granted     bool   `json:"granted,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterRegistrationToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterRegistrationTokenSpec   `json:"spec,omitempty"`
	Status ClusterRegistrationTokenStatus `json:"status,omitempty"`
}

type ClusterRegistrationTokenSpec struct {
	TTL *metav1.Duration `json:"ttl,omitempty"`
}

type ClusterRegistrationTokenStatus struct {
	Expires    *metav1.Time `json:"expires,omitempty"`
	SecretName string       `json:"secretName,omitempty"`
}
