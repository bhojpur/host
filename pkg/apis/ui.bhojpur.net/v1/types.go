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
	"k8s.io/apimachinery/pkg/util/intstr"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NavLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NavLinkSpec `json:"spec"`
}

type NavLinkSpec struct {
	Label       string                `json:"label,omitempty"`
	Description string                `json:"description,omitempty"`
	SideLabel   string                `json:"sideLabel,omitempty"`
	IconSrc     string                `json:"iconSrc,omitempty"`
	Group       string                `json:"group,omitempty"`
	Target      string                `json:"target,omitempty"`
	ToURL       string                `json:"toURL,omitempty"`
	ToService   *NavLinkTargetService `json:"toService,omitempty"`
}

type NavLinkTargetService struct {
	Namespace string              `json:"namespace,omitempty" bhojpur:"required"`
	Name      string              `json:"name,omitempty" bhojpur:"required"`
	Scheme    string              `json:"scheme,omitempty" bhojpur:"default=http,options=http|https,type=enum"`
	Port      *intstr.IntOrString `json:"port,omitempty"`
	Path      string              `json:"path,omitempty"`
}
