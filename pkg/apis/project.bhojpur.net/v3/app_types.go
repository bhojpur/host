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
	"strings"

	"github.com/bhojpur/host/pkg/core/condition"
	"github.com/bhojpur/host/pkg/core/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type App struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppSpec   `json:"spec,omitempty"`
	Status AppStatus `json:"status,omitempty"`
}

func (a *App) ObjClusterName() string {
	return a.Spec.ObjClusterName()
}

type AppSpec struct {
	ProjectName         string            `json:"projectName,omitempty" bhojpur:"type=reference[/v3/schemas/project]"`
	Description         string            `json:"description,omitempty"`
	TargetNamespace     string            `json:"targetNamespace,omitempty"`
	ExternalID          string            `json:"externalId,omitempty"`
	Files               map[string]string `json:"files,omitempty"`
	Answers             map[string]string `json:"answers,omitempty"`
	AnswersSetString    map[string]string `json:"answersSetString,omitempty"`
	Wait                bool              `json:"wait,omitempty"`
	Timeout             int               `json:"timeout,omitempty" bhojpur:"min=1,default=300"`
	AppRevisionName     string            `json:"appRevisionName,omitempty" bhojpur:"type=reference[/v3/project/schemas/apprevision]"`
	Prune               bool              `json:"prune,omitempty"`
	MultiClusterAppName string            `json:"multiClusterAppName,omitempty" bhojpur:"type=reference[/v3/schemas/multiclusterapp]"`
	ValuesYaml          string            `json:"valuesYaml,omitempty"`
	MaxRevisionCount    int               `json:"maxRevisionCount,omitempty"`
}

func (a *AppSpec) ObjClusterName() string {
	if parts := strings.SplitN(a.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

var (
	AppConditionInstalled                  condition.Cond = "Installed"
	AppConditionMigrated                   condition.Cond = "Migrated"
	AppConditionDeployed                   condition.Cond = "Deployed"
	AppConditionForceUpgrade               condition.Cond = "ForceUpgrade"
	AppConditionUserTriggeredAction        condition.Cond = "UserTriggeredAction"
	IstioConditionMetricExpressionDeployed condition.Cond = "MetricExpressionDeployed"
)

type AppStatus struct {
	AppliedFiles         map[string]string `json:"appliedFiles,omitempty"`
	Notes                string            `json:"notes,omitempty"`
	Conditions           []AppCondition    `json:"conditions,omitempty"`
	LastAppliedTemplates string            `json:"lastAppliedTemplate,omitempty"`
	HelmVersion          string            `json:"helmVersion,omitempty" bhojpur:"noupdate,nocreate"`
}

type AppCondition struct {
	// Type of cluster condition.
	Type condition.Cond `json:"type"`
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

type AppRevision struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AppRevisionSpec   `json:"spec,omitempty"`
	Status AppRevisionStatus `json:"status,omitempty"`
}

type AppRevisionSpec struct {
	ProjectName string `json:"projectName,omitempty" bhojpur:"type=reference[/v3/schemas/project]"`
}

func (a *AppRevisionSpec) ObjClusterName() string {
	if parts := strings.SplitN(a.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type AppRevisionStatus struct {
	ProjectName      string            `json:"projectName,omitempty" bhojpur:"type=reference[/v3/schemas/project]"`
	ExternalID       string            `json:"externalId"`
	Answers          map[string]string `json:"answers"`
	AnswersSetString map[string]string `json:"answersSetString"`
	Digest           string            `json:"digest"`
	ValuesYaml       string            `json:"valuesYaml,omitempty"`
	Files            map[string]string `json:"files,omitempty"`
}

func (a *AppRevisionStatus) ObjClusterName() string {
	if parts := strings.SplitN(a.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type AppUpgradeConfig struct {
	ExternalID       string            `json:"externalId,omitempty"`
	Answers          map[string]string `json:"answers,omitempty"`
	AnswersSetString map[string]string `json:"answersSetString,omitempty"`
	ForceUpgrade     bool              `json:"forceUpgrade,omitempty"`
	Files            map[string]string `json:"files,omitempty"`
	ValuesYaml       string            `json:"valuesYaml,omitempty"`
}

type RollbackRevision struct {
	RevisionName string `json:"revisionName,omitempty" bhojpur:"type=reference[/v3/project/schemas/apprevision]"`
	ForceUpgrade bool   `json:"forceUpgrade,omitempty"`
}
