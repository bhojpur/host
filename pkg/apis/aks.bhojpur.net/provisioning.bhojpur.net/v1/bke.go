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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

type BKEMachinePool struct {
	bkev1.BKECommonNodeConfig

	Paused                       bool                         `json:"paused,omitempty"`
	EtcdRole                     bool                         `json:"etcdRole,omitempty"`
	ControlPlaneRole             bool                         `json:"controlPlaneRole,omitempty"`
	WorkerRole                   bool                         `json:"workerRole,omitempty"`
	DrainBeforeDelete            bool                         `json:"drainBeforeDelete,omitempty"`
	DrainBeforeDeleteTimeout     *metav1.Duration             `json:"drainBeforeDeleteTimeout,omitempty"`
	NodeConfig                   *corev1.ObjectReference      `json:"machineConfigRef,omitempty" bhojpur:"required"`
	Name                         string                       `json:"name,omitempty" bhojpur:"required"`
	DisplayName                  string                       `json:"displayName,omitempty"`
	Quantity                     *int32                       `json:"quantity,omitempty"`
	RollingUpdate                *BKEMachinePoolRollingUpdate `json:"rollingUpdate,omitempty"`
	MachineDeploymentLabels      map[string]string            `json:"machineDeploymentLabels,omitempty"`
	MachineDeploymentAnnotations map[string]string            `json:"machineDeploymentAnnotations,omitempty"`
	NodeStartupTimeout           *metav1.Duration             `json:"nodeStartupTimeout,omitempty"`
	UnhealthyNodeTimeout         *metav1.Duration             `json:"unhealthyNodeTimeout,omitempty"`
	MaxUnhealthy                 *string                      `json:"maxUnhealthy,omitempty"`
	UnhealthyRange               *string                      `json:"unhealthyRange,omitempty"`
	MachineOS                    string                       `json:"machineOS,omitempty"`
}

type BKEMachinePoolRollingUpdate struct {
	// The maximum number of machines that can be unavailable during the update.
	// Value can be an absolute number (ex: 5) or a percentage of desired
	// machines (ex: 10%).
	// Absolute number is calculated from percentage by rounding down.
	// This can not be 0 if MaxSurge is 0.
	// Defaults to 0.
	// Example: when this is set to 30%, the old MachineSet can be scaled
	// down to 70% of desired machines immediately when the rolling update
	// starts. Once new machines are ready, old MachineSet can be scaled
	// down further, followed by scaling up the new MachineSet, ensuring
	// that the total number of machines available at all times
	// during the update is at least 70% of desired machines.
	// +optional
	MaxUnavailable *intstr.IntOrString `json:"maxUnavailable,omitempty"`

	// The maximum number of machines that can be scheduled above the
	// desired number of machines.
	// Value can be an absolute number (ex: 5) or a percentage of
	// desired machines (ex: 10%).
	// This can not be 0 if MaxUnavailable is 0.
	// Absolute number is calculated from percentage by rounding up.
	// Defaults to 1.
	// Example: when this is set to 30%, the new MachineSet can be scaled
	// up immediately when the rolling update starts, such that the total
	// number of old and new machines do not exceed 130% of desired
	// machines. Once old machines have been killed, new MachineSet can
	// be scaled up further, ensuring that total number of machines running
	// at any time during the update is at most 130% of desired machines.
	// +optional
	MaxSurge *intstr.IntOrString `json:"maxSurge,omitempty"`
}

type BKEConfig struct {
	bkev1.BKEClusterSpecCommon

	ETCDSnapshotCreate   *bkev1.ETCDSnapshotCreate   `json:"etcdSnapshotCreate,omitempty"`
	ETCDSnapshotRestore  *bkev1.ETCDSnapshotRestore  `json:"etcdSnapshotRestore,omitempty"`
	RotateCertificates   *bkev1.RotateCertificates   `json:"rotateCertificates,omitempty"`
	RotateEncryptionKeys *bkev1.RotateEncryptionKeys `json:"rotateEncryptionKeys,omitempty"`

	MachinePools      []BKEMachinePool        `json:"machinePools,omitempty"`
	InfrastructureRef *corev1.ObjectReference `json:"infrastructureRef,omitempty"`
}
