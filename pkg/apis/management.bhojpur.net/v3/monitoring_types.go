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

type MonitoringStatus struct {
	GrafanaEndpoint string                `json:"grafanaEndpoint,omitempty"`
	Conditions      []MonitoringCondition `json:"conditions,omitempty"`
}

type MonitoringCondition struct {
	// Type of cluster condition.
	Type ClusterConditionType `json:"type"`
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

const (
	MonitoringConditionGrafanaDeployed           condition.Cond = "GrafanaDeployed"
	MonitoringConditionPrometheusDeployed        condition.Cond = "PrometheusDeployed"
	MonitoringConditionAlertmaanagerDeployed     condition.Cond = "AlertmanagerDeployed"
	MonitoringConditionNodeExporterDeployed      condition.Cond = "NodeExporterDeployed"
	MonitoringConditionKubeStateExporterDeployed condition.Cond = "KubeStateExporterDeployed"
	MonitoringConditionMetricExpressionDeployed  condition.Cond = "MetricExpressionDeployed"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterMonitorGraph struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterMonitorGraphSpec `json:"spec"`
}

func (c *ClusterMonitorGraph) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectMonitorGraph struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectMonitorGraphSpec `json:"spec"`
}

func (p *ProjectMonitorGraph) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type ClusterMonitorGraphSpec struct {
	ClusterName         string `json:"clusterName" bhojpur:"type=reference[cluster]"`
	ResourceType        string `json:"resourceType,omitempty"  bhojpur:"type=enum,options=node|cluster|etcd|apiserver|scheduler|controllermanager|fluentd|istiocluster|istioproject"`
	DisplayResourceType string `json:"displayResourceType,omitempty" bhojpur:"type=enum,options=node|cluster|etcd|kube-component|bhojpur-component"`
	CommonMonitorGraphSpec
}

func (c *ClusterMonitorGraphSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectMonitorGraphSpec struct {
	ProjectName         string `json:"projectName" bhojpur:"type=reference[project]"`
	ResourceType        string `json:"resourceType,omitempty" bhojpur:"type=enum,options=workload|pod|container"`
	DisplayResourceType string `json:"displayResourceType,omitempty" bhojpur:"type=enum,options=workload|pod|container"`
	CommonMonitorGraphSpec
}

func (p *ProjectMonitorGraphSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type CommonMonitorGraphSpec struct {
	Description            string            `json:"description,omitempty"`
	MetricsSelector        map[string]string `json:"metricsSelector,omitempty"`
	DetailsMetricsSelector map[string]string `json:"detailsMetricsSelector,omitempty"`
	YAxis                  YAxis             `json:"yAxis,omitempty"`
	Priority               int               `json:"priority,omitempty"`
	GraphType              string            `json:"graphType,omitempty" bhojpur:"type=enum,options=graph|singlestat"`
}

type YAxis struct {
	Unit string `json:"unit,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MonitorMetric struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MonitorMetricSpec `json:"spec"`
}

type MonitorMetricSpec struct {
	Expression   string `json:"expression,omitempty" bhojpur:"required"`
	LegendFormat string `json:"legendFormat,omitempty"`
	Description  string `json:"description,omitempty"`
}

type QueryGraphInput struct {
	From         string            `json:"from,omitempty"`
	To           string            `json:"to,omitempty"`
	Interval     string            `json:"interval,omitempty"`
	MetricParams map[string]string `json:"metricParams,omitempty"`
	Filters      map[string]string `json:"filters,omitempty"`
	IsDetails    bool              `json:"isDetails,omitempty"`
}

type QueryClusterGraphOutput struct {
	Type string              `json:"type,omitempty"`
	Data []QueryClusterGraph `json:"data,omitempty"`
}

type QueryClusterGraph struct {
	GraphName string        `json:"graphID" bhojpur:"type=reference[clusterMonitorGraph]"`
	Series    []*TimeSeries `json:"series" bhojpur:"type=array[reference[timeSeries]]"`
}

type QueryProjectGraphOutput struct {
	Type string              `json:"type,omitempty"`
	Data []QueryProjectGraph `json:"data,omitempty"`
}

type QueryProjectGraph struct {
	GraphName string        `json:"graphID" bhojpur:"type=reference[projectMonitorGraph]"`
	Series    []*TimeSeries `json:"series" bhojpur:"type=array[reference[timeSeries]]"`
}

type QueryClusterMetricInput struct {
	ClusterName string `json:"clusterId" bhojpur:"type=reference[cluster]"`
	CommonQueryMetricInput
}

func (q *QueryClusterMetricInput) ObjClusterName() string {
	return q.ClusterName
}

type QueryProjectMetricInput struct {
	ProjectName string `json:"projectId" bhojpur:"type=reference[project]"`
	CommonQueryMetricInput
}

func (q *QueryProjectMetricInput) ObjClusterName() string {
	if parts := strings.SplitN(q.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type CommonQueryMetricInput struct {
	From     string `json:"from,omitempty"`
	To       string `json:"to,omitempty"`
	Interval string `json:"interval,omitempty"`
	Expr     string `json:"expr,omitempty" bhojpur:"required"`
}

type QueryMetricOutput struct {
	Type   string        `json:"type,omitempty"`
	Series []*TimeSeries `json:"series" bhojpur:"type=array[reference[timeSeries]]"`
}

type TimeSeries struct {
	Name   string      `json:"name"`
	Points [][]float64 `json:"points" bhojpur:"type=array[array[float]]"`
}

type MetricNamesOutput struct {
	Type  string   `json:"type,omitempty"`
	Names []string `json:"names" bhojpur:"type=array[string]"`
}

type ClusterMetricNamesInput struct {
	ClusterName string `json:"clusterId" bhojpur:"type=reference[cluster]"`
}

func (c *ClusterMetricNamesInput) ObjClusterName() string {
	return c.ClusterName
}

type ProjectMetricNamesInput struct {
	ProjectName string `json:"projectId" bhojpur:"type=reference[project]"`
}

func (p *ProjectMetricNamesInput) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}
