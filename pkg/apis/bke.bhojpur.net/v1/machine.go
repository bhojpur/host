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
	"github.com/bhojpur/host/pkg/common/genericcondition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	capi "sigs.k8s.io/cluster-api/api/v1beta1"
)

type BKECommonNodeConfig struct {
	Labels                    map[string]string `json:"labels,omitempty"`
	Taints                    []corev1.Taint    `json:"taints,omitempty"`
	CloudCredentialSecretName string            `json:"cloudCredentialSecretName,omitempty"`
}

type BKEMachineStatus struct {
	Conditions                []genericcondition.GenericCondition `json:"conditions,omitempty"`
	JobComplete               bool                                `json:"jobComplete,omitempty"`
	JobName                   string                              `json:"jobName,omitempty"`
	Ready                     bool                                `json:"ready,omitempty"`
	DriverHash                string                              `json:"driverHash,omitempty"`
	DriverURL                 string                              `json:"driverUrl,omitempty"`
	CloudCredentialSecretName string                              `json:"cloudCredentialSecretName,omitempty"`
	FailureReason             string                              `json:"failureReason,omitempty"`
	FailureMessage            string                              `json:"failureMessage,omitempty"`
	Addresses                 []capi.MachineAddress               `json:"addresses,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CustomMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomMachineSpec   `json:"spec,omitempty"`
	Status CustomMachineStatus `json:"status,omitempty"`
}

type CustomMachineSpec struct {
	ProviderID string `json:"providerID,omitempty"`
}

type CustomMachineStatus struct {
	Conditions []genericcondition.GenericCondition `json:"conditions,omitempty"`
	Ready      bool                                `json:"ready,omitempty"`
	Addresses  []capi.MachineAddress               `json:"addresses,omitempty"`
}
