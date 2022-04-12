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
	fleet "github.com/bhojpur/host/pkg/apis/fleet.bhojpur.net/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ManagedChart struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ManagedChartSpec   `json:"spec"`
	Status ManagedChartStatus `json:"status"`
}

type ManagedChartSpec struct {
	Paused           bool               `json:"paused,omitempty"`
	Chart            string             `json:"chart,omitempty"`
	RepoName         string             `json:"repoName,omitempty"`
	ReleaseName      string             `json:"releaseName,omitempty"`
	Version          string             `json:"version,omitempty"`
	TimeoutSeconds   int                `json:"timeoutSeconds,omitempty"`
	Values           *fleet.GenericMap  `json:"values,omitempty"`
	Force            bool               `json:"force,omitempty"`
	TakeOwnership    bool               `json:"takeOwnership,omitempty"`
	MaxHistory       int                `json:"maxHistory,omitempty"`
	DefaultNamespace string             `json:"defaultNamespace,omitempty"`
	TargetNamespace  string             `json:"namespace,omitempty"`
	ServiceAccount   string             `json:"serviceAccount,omitempty"`
	Diff             *fleet.DiffOptions `json:"diff,omitempty"`

	RolloutStrategy *fleet.RolloutStrategy `json:"rolloutStrategy,omitempty"`
	Targets         []fleet.BundleTarget   `json:"targets,omitempty"`
}

type ManagedChartStatus struct {
	fleet.BundleStatus
}
