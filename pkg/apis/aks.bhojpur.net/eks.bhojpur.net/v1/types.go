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

type EKSClusterConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EKSClusterConfigSpec   `json:"spec"`
	Status EKSClusterConfigStatus `json:"status"`
}

// EKSClusterConfigSpec is the spec for a EKSClusterConfig resource
type EKSClusterConfigSpec struct {
	AmazonCredentialSecret string            `json:"amazonCredentialSecret"`
	DisplayName            string            `json:"displayName" bhojpur:"noupdate"`
	Region                 string            `json:"region" bhojpur:"noupdate"`
	Imported               bool              `json:"imported" bhojpur:"noupdate"`
	KubernetesVersion      *string           `json:"kubernetesVersion" bhojpur:"pointer"`
	Tags                   map[string]string `json:"tags"`
	SecretsEncryption      *bool             `json:"secretsEncryption" bhojpur:"noupdate"`
	KmsKey                 *string           `json:"kmsKey" bhojpur:"noupdate,pointer"`
	PublicAccess           *bool             `json:"publicAccess"`
	PrivateAccess          *bool             `json:"privateAccess"`
	PublicAccessSources    []string          `json:"publicAccessSources"`
	LoggingTypes           []string          `json:"loggingTypes"`
	Subnets                []string          `json:"subnets" bhojpur:"noupdate"`
	SecurityGroups         []string          `json:"securityGroups" bhojpur:"noupdate"`
	ServiceRole            *string           `json:"serviceRole" bhojpur:"noupdate,pointer"`
	NodeGroups             []NodeGroup       `json:"nodeGroups"`
}

type EKSClusterConfigStatus struct {
	Phase                         string            `json:"phase"`
	VirtualNetwork                string            `json:"virtualNetwork"`
	Subnets                       []string          `json:"subnets"`
	SecurityGroups                []string          `json:"securityGroups"`
	ManagedLaunchTemplateID       string            `json:"managedLaunchTemplateID"`
	ManagedLaunchTemplateVersions map[string]string `json:"managedLaunchTemplateVersions"`
	TemplateVersionsToDelete      []string          `json:"templateVersionsToDelete"`
	// describes how the above network fields were provided. Valid values are provided and generated
	NetworkFieldsSource string `json:"networkFieldsSource"`
	FailureMessage      string `json:"failureMessage"`
}

type NodeGroup struct {
	Gpu                  *bool              `json:"gpu"`
	ImageID              *string            `json:"imageId" bhojpur:"pointer"`
	NodegroupName        *string            `json:"nodegroupName" bhojpur:"required,pointer" bhojpur:"required"`
	DiskSize             *int64             `json:"diskSize"`
	InstanceType         *string            `json:"instanceType" bhojpur:"pointer"`
	Labels               map[string]*string `json:"labels"`
	Ec2SshKey            *string            `json:"ec2SshKey" bhojpur:"pointer"`
	DesiredSize          *int64             `json:"desiredSize"`
	MaxSize              *int64             `json:"maxSize"`
	MinSize              *int64             `json:"minSize"`
	Subnets              []string           `json:"subnets"`
	Tags                 map[string]*string `json:"tags"`
	ResourceTags         map[string]*string `json:"resourceTags"`
	UserData             *string            `json:"userData" bhojpur:"pointer"`
	Version              *string            `json:"version" bhojpur:"pointer"`
	LaunchTemplate       *LaunchTemplate    `json:"launchTemplate"`
	RequestSpotInstances *bool              `json:"requestSpotInstances"`
	SpotInstanceTypes    []*string          `json:"spotInstanceTypes"`
}

type LaunchTemplate struct {
	ID      *string `json:"id" bhojpur:"pointer"`
	Name    *string `json:"name" bhojpur:"pointer"`
	Version *int64  `json:"version"`
}
