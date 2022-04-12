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
	"github.com/bhojpur/host/pkg/engine/types/kdm"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ClusterScanRunType string
type CisScanProfileType string

const (
	ClusterScanConditionCreated      condition.Cond = Created
	ClusterScanConditionRunCompleted condition.Cond = RunCompleted
	ClusterScanConditionCompleted    condition.Cond = Completed
	ClusterScanConditionFailed       condition.Cond = Failed
	ClusterScanConditionAlerted      condition.Cond = Alerted

	ClusterScanTypeCis         = "cis"
	DefaultNamespaceForCis     = "security-scan"
	DefaultSonobuoyPodName     = "security-scan-runner"
	ConfigMapNameForUserConfig = "security-scan-cfg"

	SonobuoyCompletionAnnotation = "field.bhojpur.net/sonobuoyDone"
	CisHelmChartOwner            = "field.bhojpur.net/clusterScanOwner"

	ClusterScanRunTypeManual    ClusterScanRunType = "manual"
	ClusterScanRunTypeScheduled ClusterScanRunType = "scheduled"

	CisScanProfileTypePermissive CisScanProfileType = "permissive"
	CisScanProfileTypeHardened   CisScanProfileType = "hardened"

	DefaultScanOutputFileName string = "output.json"
)

type CisScanConfig struct {
	// IDs of the checks that need to be skipped in the final report
	OverrideSkip []string `json:"overrideSkip"`
	// Override the CIS benchmark version to use for the scan (instead of latest)
	OverrideBenchmarkVersion string `json:"overrideBenchmarkVersion,omitempty"`
	// scan profile to use
	Profile CisScanProfileType `json:"profile,omitempty" bhojpur:"required,options=permissive|hardened,default=permissive"`
	// Internal flag for debugging master component of the scan
	DebugMaster bool `json:"debugMaster"`
	// Internal flag for debugging worker component of the scan
	DebugWorker bool `json:"debugWorker"`
}

type CisScanStatus struct {
	Total         int `json:"total"`
	Pass          int `json:"pass"`
	Fail          int `json:"fail"`
	Skip          int `json:"skip"`
	NotApplicable int `json:"notApplicable"`
}

type ClusterScanConfig struct {
	CisScanConfig *CisScanConfig `json:"cisScanConfig"`
}

type ClusterScanCondition struct {
	// Type of condition.
	Type string `json:"type"`
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

type ClusterScanSpec struct {
	ScanType string `json:"scanType"`
	// cluster ID
	ClusterID string `json:"clusterId,omitempty" bhojpur:"required,type=reference[cluster]"`
	// Run type
	RunType ClusterScanRunType `json:"runType,omitempty"`
	// scanConfig
	ScanConfig ClusterScanConfig `yaml:",omitempty" json:"scanConfig,omitempty"`
}

type ClusterScanStatus struct {
	Conditions    []ClusterScanCondition `json:"conditions"`
	CisScanStatus *CisScanStatus         `json:"cisScanStatus"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterScan struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterScanSpec   `json:"spec"`
	Status ClusterScanStatus `yaml:"status" json:"status,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CisConfig struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Params kdm.CisConfigParams `yaml:"params" json:"params,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type CisBenchmarkVersion struct {
	types.Namespaced

	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Info kdm.CisBenchmarkVersionInfo `json:"info" yaml:"info"`
}

type ScheduledClusterScanConfig struct {
	// Cron Expression for Schedule
	CronSchedule string `yaml:"cron_schedule" json:"cronSchedule,omitempty"`
	// Number of past scans to keep
	Retention int `yaml:"retention" json:"retention,omitempty"`
}

type ScheduledClusterScan struct {
	// Enable or disable scheduled scans
	Enabled        bool                        `yaml:"enabled" json:"enabled,omitempty" bhojpur:"default=false"`
	ScheduleConfig *ScheduledClusterScanConfig `yaml:"schedule_config" json:"scheduleConfig,omitempty"`
	ScanConfig     *ClusterScanConfig          `yaml:"scan_config,omitempty" json:"scanConfig,omitempty"`
}

type ScheduledClusterScanStatus struct {
	Enabled          bool   `yaml:"enabled" json:"enabled,omitempty"`
	LastRunTimestamp string `yaml:"last_run_timestamp" json:"lastRunTimestamp"`
}
