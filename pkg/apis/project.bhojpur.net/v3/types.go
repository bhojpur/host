package v3

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
	"github.com/bhojpur/host/pkg/core/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ServiceAccountToken struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	AccountName string `json:"accountName"`
	AccountUID  string `json:"accountUid"`
	Description string `json:"description"`
	Token       string `json:"token" bhojpur:"writeOnly"`
	CACRT       string `json:"caCrt"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespacedServiceAccountToken ServiceAccountToken

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DockerCredential struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string                        `json:"description"`
	Registries  map[string]RegistryCredential `json:"registries"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespacedDockerCredential DockerCredential

type RegistryCredential struct {
	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password" bhojpur:"writeOnly"`
	Auth        string `json:"auth" bhojpur:"writeOnly"`
	Email       string `json:"email"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Certificate struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string `json:"description"`
	Certs       string `json:"certs"`
	Key         string `json:"key" bhojpur:"writeOnly"`

	CertFingerprint         string   `json:"certFingerprint" bhojpur:"nocreate,noupdate"`
	CN                      string   `json:"cn" bhojpur:"nocreate,noupdate"`
	Version                 string   `json:"version" bhojpur:"nocreate,noupdate"`
	ExpiresAt               string   `json:"expiresAt" bhojpur:"nocreate,noupdate"`
	Issuer                  string   `json:"issuer" bhojpur:"nocreate,noupdate"`
	IssuedAt                string   `json:"issuedAt" bhojpur:"nocreate,noupdate"`
	Algorithm               string   `json:"algorithm" bhojpur:"nocreate,noupdate"`
	SerialNumber            string   `json:"serialNumber" bhojpur:"nocreate,noupdate"`
	KeySize                 string   `json:"keySize" bhojpur:"nocreate,noupdate"`
	SubjectAlternativeNames []string `json:"subjectAlternativeNames" bhojpur:"nocreate,noupdate"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespacedCertificate Certificate

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BasicAuth struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string `json:"description"`
	Username    string `json:"username"`
	Password    string `json:"password" bhojpur:"writeOnly"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespacedBasicAuth BasicAuth

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SSHAuth struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Description string `json:"description"`
	PrivateKey  string `json:"privateKey" bhojpur:"writeOnly"`
	Fingerprint string `json:"certFingerprint" bhojpur:"nocreate,noupdate"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NamespacedSSHAuth SSHAuth

type PublicEndpoint struct {
	NodeName  string   `json:"nodeName,omitempty" bhojpur:"type=reference[/v3/schemas/node],nocreate,noupdate"`
	Addresses []string `json:"addresses,omitempty" bhojpur:"nocreate,noupdate"`
	Port      int32    `json:"port,omitempty" bhojpur:"nocreate,noupdate"`
	Protocol  string   `json:"protocol,omitempty" bhojpur:"nocreate,noupdate"`
	// for node port service endpoint
	ServiceName string `json:"serviceName,omitempty" bhojpur:"type=reference[service],nocreate,noupdate"`
	// for host port endpoint
	PodName string `json:"podName,omitempty" bhojpur:"type=reference[pod],nocreate,noupdate"`
	// for ingress endpoint. ServiceName, podName, ingressName are mutually exclusive
	IngressName string `json:"ingressName,omitempty" bhojpur:"type=reference[ingress],nocreate,noupdate"`
	// Hostname/path are set for Ingress endpoints
	Hostname string `json:"hostname,omitempty" bhojpur:"nocreate,noupdate"`
	Path     string `json:"path,omitempty" bhojpur:"nocreate,noupdate"`
	// True when endpoint is exposed on every node
	AllNodes bool `json:"allNodes" bhojpur:"nocreate,noupdate"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Workload struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
}

type DeploymentRollbackInput struct {
	ReplicaSetID string `json:"replicaSetId" bhojpur:"type=reference[replicaSet]"`
}

type WorkloadMetric struct {
	Port   int32  `json:"port,omitempty"`
	Path   string `json:"path,omitempty"`
	Schema string `json:"schema,omitempty" bhojpur:"type=enum,options=HTTP|HTTPS"`
}
