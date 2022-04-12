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
	"github.com/bhojpur/host/pkg/core/condition"
	"github.com/bhojpur/host/pkg/core/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const ClusterTemplateRevisionConditionSecretsMigrated condition.Cond = "SecretsMigrated"

type ClusterTemplateRevisionConditionType string

type ClusterTemplateRevisionCondition struct {
	// Type of cluster template revision condition.
	Type ClusterTemplateRevisionConditionType `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterTemplate struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterTemplateSpec `json:"spec"`
}

type ClusterTemplateSpec struct {
	DisplayName         string `json:"displayName" bhojpur:"required"`
	Description         string `json:"description"`
	DefaultRevisionName string `json:"defaultRevisionName,omitempty" bhojpur:"type=reference[clusterTemplateRevision]"`

	Members []Member `json:"members,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterTemplateRevision struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec   ClusterTemplateRevisionSpec   `json:"spec"`
	Status ClusterTemplateRevisionStatus `json:"status"`
}

type ClusterTemplateRevisionSpec struct {
	DisplayName         string `json:"displayName" bhojpur:"required"`
	Enabled             *bool  `json:"enabled,omitempty" bhojpur:"default=true"`
	ClusterTemplateName string `json:"clusterTemplateName,omitempty" bhojpur:"type=reference[clusterTemplate],required,noupdate"`

	Questions     []Question       `json:"questions,omitempty"`
	ClusterConfig *ClusterSpecBase `json:"clusterConfig" bhojpur:"required"`
}

type ClusterTemplateRevisionStatus struct {
	PrivateRegistrySecret string                             `json:"privateRegistrySecret,omitempty" bhojpur:"nocreate,noupdate"`
	S3CredentialSecret    string                             `json:"s3CredentialSecret,omitempty" bhojpur:"nocreate,noupdate"`
	WeavePasswordSecret   string                             `json:"weavePasswordSecret,omitempty" bhojpur:"nocreate,noupdate"`
	VsphereSecret         string                             `json:"vsphereSecret,omitempty" bhojpur:"nocreate,noupdate"`
	VirtualCenterSecret   string                             `json:"virtualCenterSecret,omitempty" bhojpur:"nocreate,noupdate"`
	OpenStackSecret       string                             `json:"openStackSecret,omitempty" bhojpur:"nocreate,noupdate"`
	AADClientSecret       string                             `json:"aadClientSecret,omitempty" bhojpur:"nocreate,noupdate"`
	AADClientCertSecret   string                             `json:"aadClientCertSecret,omitempty" bhojpur:"nocreate,noupdate"`
	Conditions            []ClusterTemplateRevisionCondition `json:"conditions,omitempty"`
}

type ClusterTemplateQuestionsOutput struct {
	Questions []Question `json:"questions,omitempty"`
}
