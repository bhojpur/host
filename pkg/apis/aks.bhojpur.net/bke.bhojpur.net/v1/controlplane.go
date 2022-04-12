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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type BKEControlPlane struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BKEControlPlaneSpec   `json:"spec"`
	Status            BKEControlPlaneStatus `json:"status,omitempty"`
}

type EnvVar struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type BKEControlPlaneSpec struct {
	BKEClusterSpecCommon

	AgentEnvVars             []EnvVar                 `json:"agentEnvVars,omitempty"`
	LocalClusterAuthEndpoint LocalClusterAuthEndpoint `json:"localClusterAuthEndpoint"`
	ETCDSnapshotCreate       *ETCDSnapshotCreate      `json:"etcdSnapshotCreate,omitempty"`
	ETCDSnapshotRestore      *ETCDSnapshotRestore     `json:"etcdSnapshotRestore,omitempty"`
	RotateCertificates       *RotateCertificates      `json:"rotateCertificates,omitempty"`
	RotateEncryptionKeys     *RotateEncryptionKeys    `json:"rotateEncryptionKeys,omitempty"`
	KubernetesVersion        string                   `json:"kubernetesVersion,omitempty"`
	ClusterName              string                   `json:"clusterName,omitempty" bhojpur:"required"`
	ManagementClusterName    string                   `json:"managementClusterName,omitempty" bhojpur:"required"`
	UnmanagedConfig          bool                     `json:"unmanagedConfig,omitempty"`
}

type ETCDSnapshotPhase string

var (
	ETCDSnapshotPhaseStarted  ETCDSnapshotPhase = "Started"
	ETCDSnapshotPhaseShutdown ETCDSnapshotPhase = "Shutdown"
	ETCDSnapshotPhaseRestore  ETCDSnapshotPhase = "Restore"
	ETCDSnapshotPhaseFinished ETCDSnapshotPhase = "Finished"
	ETCDSnapshotPhaseFailed   ETCDSnapshotPhase = "Failed"
)

type RotateEncryptionKeysPhase string

const (
	RotateEncryptionKeysPhaseStart              RotateEncryptionKeysPhase = "Start"
	RotateEncryptionKeysPhaseRestartLeader      RotateEncryptionKeysPhase = "RestartLeader"
	RotateEncryptionKeysPhaseVerifyLeaderStatus RotateEncryptionKeysPhase = "VerifyLeaderStatus"
	RotateEncryptionKeysPhaseRestartFollowers   RotateEncryptionKeysPhase = "RestartFollowers"
	RotateEncryptionKeysPhaseApplyLeader        RotateEncryptionKeysPhase = "ApplyLeader"
	RotateEncryptionKeysPhaseDone               RotateEncryptionKeysPhase = "Done"
)

type BKEControlPlaneStatus struct {
	Conditions                     []genericcondition.GenericCondition `json:"conditions,omitempty"`
	Ready                          bool                                `json:"ready,omitempty"`
	ObservedGeneration             int64                               `json:"observedGeneration"`
	CertificateRotationGeneration  int64                               `json:"certificateRotationGeneration"`
	RotateEncryptionKeysGeneration int64                               `json:"rotateEncryptionKeysGeneration"`
	RotateEncryptionKeysPhase      RotateEncryptionKeysPhase           `json:"rotateEncryptionKeysPhase"`
	ETCDSnapshotRestore            *ETCDSnapshotRestore                `json:"etcdSnapshotRestore,omitempty"`
	ETCDSnapshotRestorePhase       ETCDSnapshotPhase                   `json:"etcdSnapshotRestorePhase,omitempty"`
	ETCDSnapshotCreate             *ETCDSnapshotCreate                 `json:"etcdSnapshotCreate,omitempty"`
	ETCDSnapshotCreatePhase        ETCDSnapshotPhase                   `json:"etcdSnapshotCreatePhase,omitempty"`
	ConfigGeneration               int64                               `json:"configGeneration,omitempty"`
	Initialized                    bool                                `json:"initialized,omitempty"`
	AgentConnected                 bool                                `json:"agentConnected,omitempty"`
}
