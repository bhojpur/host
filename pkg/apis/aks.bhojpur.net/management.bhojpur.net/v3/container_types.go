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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KontainerDriver struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec KontainerDriverSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status KontainerDriverStatus `json:"status"`
}

type KontainerDriverStatus struct {
	ActualURL      string      `json:"actualUrl"`
	ExecutablePath string      `json:"executablePath"`
	Conditions     []Condition `json:"conditions"`
	DisplayName    string      `json:"displayName"`
}

type KontainerDriverSpec struct {
	URL              string   `json:"url" bhojpur:"required"`
	Checksum         string   `json:"checksum"`
	BuiltIn          bool     `json:"builtIn" bhojpur:"noupdate"`
	Active           bool     `json:"active"`
	UIURL            string   `json:"uiUrl"`
	WhitelistDomains []string `json:"whitelistDomains,omitempty"`
}

var (
	KontainerDriverConditionDownloaded condition.Cond = "Downloaded"
	KontainerDriverConditionInstalled  condition.Cond = "Installed"
	KontainerDriverConditionActive     condition.Cond = "Active"
	KontainerDriverConditionInactive   condition.Cond = "Inactive"
)
