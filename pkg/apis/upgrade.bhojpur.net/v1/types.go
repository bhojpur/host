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
	"time"

	"github.com/bhojpur/host/pkg/common/condition"
	"github.com/bhojpur/host/pkg/common/genericcondition"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	// PlanLatestResolved indicates that the latest version as per the spec has been determined.
	PlanLatestResolved = condition.Cond("LatestResolved")
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Plan represents a "JobSet" of ApplyingNodes
type Plan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PlanSpec   `json:"spec,omitempty"`
	Status PlanStatus `json:"status,omitempty"`
}

// PlanSpec represents the user-configurable details of a Plan.
type PlanSpec struct {
	Concurrency        int64                 `json:"concurrency,omitempty"`
	NodeSelector       *metav1.LabelSelector `json:"nodeSelector,omitempty"`
	ServiceAccountName string                `json:"serviceAccountName,omitempty"`

	Channel string       `json:"channel,omitempty"`
	Version string       `json:"version,omitempty"`
	Secrets []SecretSpec `json:"secrets,omitempty"`

	Tolerations []corev1.Toleration `json:"tolerations,omitempty"`

	Prepare *ContainerSpec `json:"prepare,omitempty"`
	Cordon  bool           `json:"cordon,omitempty"`
	Drain   *DrainSpec     `json:"drain,omitempty"`
	Upgrade *ContainerSpec `json:"upgrade,omitempty" bhojpur:"required"`
}

// PlanStatus represents the resulting state from processing Plan events.
type PlanStatus struct {
	Conditions    []genericcondition.GenericCondition `json:"conditions,omitempty"`
	LatestVersion string                              `json:"latestVersion,omitempty"`
	LatestHash    string                              `json:"latestHash,omitempty"`
	Applying      []string                            `json:"applying,omitempty"`
}

// ContainerSpec is a simplified container template.
type ContainerSpec struct {
	Image   string                 `json:"image,omitempty"`
	Command []string               `json:"command,omitempty"`
	Args    []string               `json:"args,omitempty"`
	Env     []corev1.EnvVar        `json:"envs,omitempty"`
	EnvFrom []corev1.EnvFromSource `json:"envFrom,omitempty"`
}

// DrainSpec encapsulates `kubectl drain` parameters minus node/pod selectors.
type DrainSpec struct {
	Timeout                  *time.Duration `json:"timeout,omitempty"`
	GracePeriod              *int32         `json:"gracePeriod,omitempty"`
	DeleteLocalData          *bool          `json:"deleteLocalData,omitempty"`
	IgnoreDaemonSets         *bool          `json:"ignoreDaemonSets,omitempty"`
	Force                    bool           `json:"force,omitempty"`
	DisableEviction          bool           `json:"disableEviction,omitempty"`
	SkipWaitForDeleteTimeout int            `json:"skipWaitForDeleteTimeout,omitempty"`
}

// SecretSpec describes a secret to be mounted for prepare/upgrade containers.
type SecretSpec struct {
	Name string `json:"name,omitempty"`
	Path string `json:"path,omitempty"`
}
