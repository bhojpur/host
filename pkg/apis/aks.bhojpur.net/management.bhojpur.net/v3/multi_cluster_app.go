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

	v3 "github.com/bhojpur/host/pkg/apis/project.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/condition"
	"github.com/bhojpur/host/pkg/core/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	MultiClusterAppConditionInstalled condition.Cond = "Installed"
	MultiClusterAppConditionDeployed  condition.Cond = "Deployed"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MultiClusterApp struct {
	types.Namespaced
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status

	Spec   MultiClusterAppSpec   `json:"spec"`
	Status MultiClusterAppStatus `json:"status"`
}

type MultiClusterAppSpec struct {
	TemplateVersionName  string          `json:"templateVersionName,omitempty" bhojpur:"type=reference[templateVersion],required"`
	Answers              []Answer        `json:"answers,omitempty"`
	Wait                 bool            `json:"wait,omitempty"`
	Timeout              int             `json:"timeout,omitempty" bhojpur:"min=1,default=300"`
	Targets              []Target        `json:"targets,omitempty" bhojpur:"required,noupdate"`
	Members              []Member        `json:"members,omitempty"`
	Roles                []string        `json:"roles,omitempty" bhojpur:"type=array[reference[roleTemplate]],required"`
	RevisionHistoryLimit int             `json:"revisionHistoryLimit,omitempty" bhojpur:"default=10"`
	UpgradeStrategy      UpgradeStrategy `json:"upgradeStrategy,omitempty"`
}

type MultiClusterAppStatus struct {
	Conditions   []v3.AppCondition `json:"conditions,omitempty"`
	RevisionName string            `json:"revisionName,omitempty" bhojpur:"type=reference[multiClusterAppRevision],required"`
	HelmVersion  string            `json:"helmVersion,omitempty" bhojpur:"nocreate,noupdate"`
}

type Target struct {
	ProjectName string `json:"projectName,omitempty" bhojpur:"type=reference[project],required"`
	AppName     string `json:"appName,omitempty" bhojpur:"type=reference[v3/projects/schemas/app]"`
	State       string `json:"state,omitempty"`
	Healthstate string `json:"healthState,omitempty"`
}

func (t *Target) ObjClusterName() string {
	if parts := strings.SplitN(t.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type Answer struct {
	ProjectName     string            `json:"projectName,omitempty" bhojpur:"type=reference[project]"`
	ClusterName     string            `json:"clusterName,omitempty" bhojpur:"type=reference[cluster]"`
	Values          map[string]string `json:"values,omitempty"`
	ValuesSetString map[string]string `json:"valuesSetString,omitempty"`
}

func (a *Answer) ObjClusterName() string {
	return a.ClusterName
}

type Member struct {
	UserName           string `json:"userName,omitempty" bhojpur:"type=reference[user]"`
	UserPrincipalName  string `json:"userPrincipalName,omitempty" bhojpur:"type=reference[principal]"`
	DisplayName        string `json:"displayName,omitempty"`
	GroupPrincipalName string `json:"groupPrincipalName,omitempty" bhojpur:"type=reference[principal]"`
	AccessType         string `json:"accessType,omitempty" bhojpur:"type=enum,options=owner|member|read-only"`
}

type UpgradeStrategy struct {
	RollingUpdate *RollingUpdate `json:"rollingUpdate,omitempty"`
}

type RollingUpdate struct {
	BatchSize int `json:"batchSize,omitempty"`
	Interval  int `json:"interval,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MultiClusterAppRevision struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	TemplateVersionName string   `json:"templateVersionName,omitempty" bhojpur:"type=reference[templateVersion]"`
	Answers             []Answer `json:"answers,omitempty"`
}

type MultiClusterAppRollbackInput struct {
	RevisionName string `json:"revisionName,omitempty" bhojpur:"type=reference[multiClusterAppRevision]"`
}

type UpdateMultiClusterAppTargetsInput struct {
	Projects []string `json:"projects" bhojpur:"type=array[reference[project]],required"`
	Answers  []Answer `json:"answers"`
}
