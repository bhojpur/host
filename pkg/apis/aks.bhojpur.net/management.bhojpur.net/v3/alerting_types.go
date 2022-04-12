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

	"github.com/bhojpur/host/pkg/core/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAlert struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterAlertSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (c *ClusterAlert) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectAlert struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectAlertSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (p *ProjectAlert) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type AlertCommonSpec struct {
	DisplayName           string      `json:"displayName,omitempty" bhojpur:"required"`
	Description           string      `json:"description,omitempty"`
	Severity              string      `json:"severity,omitempty" bhojpur:"required,options=info|critical|warning,default=critical"`
	Recipients            []Recipient `json:"recipients,omitempty" bhojpur:"required"`
	InitialWaitSeconds    int         `json:"initialWaitSeconds,omitempty" bhojpur:"required,default=180,min=0"`
	RepeatIntervalSeconds int         `json:"repeatIntervalSeconds,omitempty"  bhojpur:"required,default=3600,min=0"`
}

type ClusterAlertSpec struct {
	AlertCommonSpec

	ClusterName         string               `json:"clusterName" bhojpur:"type=reference[cluster]"`
	TargetNode          *TargetNode          `json:"targetNode,omitempty"`
	TargetSystemService *TargetSystemService `json:"targetSystemService,omitempty"`
	TargetEvent         *TargetEvent         `json:"targetEvent,omitempty"`
}

func (c *ClusterAlertSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectAlertSpec struct {
	AlertCommonSpec

	ProjectName    string          `json:"projectName" bhojpur:"type=reference[project]"`
	TargetWorkload *TargetWorkload `json:"targetWorkload,omitempty"`
	TargetPod      *TargetPod      `json:"targetPod,omitempty"`
}

func (p *ProjectAlertSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type Recipient struct {
	Recipient    string `json:"recipient,omitempty"`
	NotifierName string `json:"notifierName,omitempty" bhojpur:"required,type=reference[notifier]"`
	NotifierType string `json:"notifierType,omitempty" bhojpur:"required,options=slack|email|pagerduty|webhook|wechat|dingtalk|msteams"`
}

type TargetNode struct {
	NodeName     string            `json:"nodeName,omitempty" bhojpur:"type=reference[node]"`
	Selector     map[string]string `json:"selector,omitempty"`
	Condition    string            `json:"condition,omitempty" bhojpur:"required,options=notready|mem|cpu,default=notready"`
	MemThreshold int               `json:"memThreshold,omitempty" bhojpur:"min=1,max=100,default=70"`
	CPUThreshold int               `json:"cpuThreshold,omitempty" bhojpur:"min=1,default=70"`
}

type TargetPod struct {
	PodName                string `json:"podName,omitempty" bhojpur:"required,type=reference[/v3/projects/schemas/pod]"`
	Condition              string `json:"condition,omitempty" bhojpur:"required,options=notrunning|notscheduled|restarts,default=notrunning"`
	RestartTimes           int    `json:"restartTimes,omitempty" bhojpur:"min=1,default=3"`
	RestartIntervalSeconds int    `json:"restartIntervalSeconds,omitempty"  bhojpur:"min=1,default=300"`
}

type TargetEvent struct {
	EventType    string `json:"eventType,omitempty" bhojpur:"required,options=Normal|Warning,default=Warning"`
	ResourceKind string `json:"resourceKind,omitempty" bhojpur:"required,options=Pod|Node|Deployment|StatefulSet|DaemonSet"`
}

type TargetWorkload struct {
	WorkloadID          string            `json:"workloadId,omitempty"`
	Selector            map[string]string `json:"selector,omitempty"`
	AvailablePercentage int               `json:"availablePercentage,omitempty" bhojpur:"required,min=1,max=100,default=70"`
}

type TargetSystemService struct {
	Condition string `json:"condition,omitempty" bhojpur:"required,options=etcd|controller-manager|scheduler,default=scheduler"`
}

type AlertStatus struct {
	AlertState string `json:"alertState,omitempty" bhojpur:"options=active|inactive|alerting|muted,default=active"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAlertGroup struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterGroupSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (c *ClusterAlertGroup) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectAlertGroup struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectGroupSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (p *ProjectAlertGroup) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type ClusterGroupSpec struct {
	ClusterName string      `json:"clusterName" bhojpur:"type=reference[cluster]"`
	Recipients  []Recipient `json:"recipients,omitempty"`
	CommonGroupField
}

func (c *ClusterGroupSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectGroupSpec struct {
	ProjectName string      `json:"projectName" bhojpur:"type=reference[project]"`
	Recipients  []Recipient `json:"recipients,omitempty"`
	CommonGroupField
}

func (p *ProjectGroupSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ClusterAlertRule struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterAlertRuleSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (c *ClusterAlertRule) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

type ClusterAlertRuleSpec struct {
	CommonRuleField
	ClusterName       string             `json:"clusterName" bhojpur:"type=reference[cluster]"`
	GroupName         string             `json:"groupName" bhojpur:"type=reference[clusterAlertGroup]"`
	NodeRule          *NodeRule          `json:"nodeRule,omitempty"`
	EventRule         *EventRule         `json:"eventRule,omitempty"`
	SystemServiceRule *SystemServiceRule `json:"systemServiceRule,omitempty"`
	MetricRule        *MetricRule        `json:"metricRule,omitempty"`
	ClusterScanRule   *ClusterScanRule   `json:"clusterScanRule,omitempty"`
}

func (c *ClusterAlertRuleSpec) ObjClusterName() string {
	return c.ClusterName
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectAlertRule struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ProjectAlertRuleSpec `json:"spec"`
	// Most recent observed status of the alert. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status AlertStatus `json:"status"`
}

func (p *ProjectAlertRule) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type ProjectAlertRuleSpec struct {
	CommonRuleField
	ProjectName  string        `json:"projectName" bhojpur:"type=reference[project]"`
	GroupName    string        `json:"groupName" bhojpur:"type=reference[projectAlertGroup]"`
	PodRule      *PodRule      `json:"podRule,omitempty"`
	WorkloadRule *WorkloadRule `json:"workloadRule,omitempty"`
	MetricRule   *MetricRule   `json:"metricRule,omitempty"`
}

func (p *ProjectAlertRuleSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type CommonGroupField struct {
	DisplayName string `json:"displayName,omitempty" bhojpur:"required"`
	Description string `json:"description,omitempty"`
	TimingField
}

type CommonRuleField struct {
	DisplayName string `json:"displayName,omitempty"`
	Severity    string `json:"severity,omitempty" bhojpur:"required,options=info|critical|warning,default=critical"`
	Inherited   *bool  `json:"inherited,omitempty" bhojpur:"default=true"`
	TimingField
}

type ClusterScanRule struct {
	ScanRunType  ClusterScanRunType `json:"scanRunType,omitempty" bhojpur:"required,options=manual|scheduled,default=scheduled"`
	FailuresOnly bool               `json:"failuresOnly,omitempty"`
}

type MetricRule struct {
	Expression     string  `json:"expression,omitempty" bhojpur:"required"`
	Description    string  `json:"description,omitempty"`
	Duration       string  `json:"duration,omitempty" bhojpur:"required"`
	Comparison     string  `json:"comparison,omitempty" bhojpur:"type=enum,options=equal|not-equal|greater-than|less-than|greater-or-equal|less-or-equal|has-value,default=equal"`
	ThresholdValue float64 `json:"thresholdValue,omitempty" bhojpur:"type=float"`
}

type TimingField struct {
	GroupWaitSeconds      int `json:"groupWaitSeconds,omitempty" bhojpur:"required,default=30,min=1"`
	GroupIntervalSeconds  int `json:"groupIntervalSeconds,omitempty" bhojpur:"required,default=180,min=1"`
	RepeatIntervalSeconds int `json:"repeatIntervalSeconds,omitempty"  bhojpur:"required,default=3600,min=1"`
}

type NodeRule struct {
	NodeName     string            `json:"nodeName,omitempty" bhojpur:"type=reference[node]"`
	Selector     map[string]string `json:"selector,omitempty"`
	Condition    string            `json:"condition,omitempty" bhojpur:"required,options=notready|mem|cpu,default=notready"`
	MemThreshold int               `json:"memThreshold,omitempty" bhojpur:"min=1,max=100,default=70"`
	CPUThreshold int               `json:"cpuThreshold,omitempty" bhojpur:"min=1,default=70"`
}

type PodRule struct {
	PodName                string `json:"podName,omitempty" bhojpur:"required,type=reference[/v3/projects/schemas/pod]"`
	Condition              string `json:"condition,omitempty" bhojpur:"required,options=notrunning|notscheduled|restarts,default=notrunning"`
	RestartTimes           int    `json:"restartTimes,omitempty" bhojpur:"min=1,default=3"`
	RestartIntervalSeconds int    `json:"restartIntervalSeconds,omitempty"  bhojpur:"min=1,default=300"`
}

type EventRule struct {
	EventType    string `json:"eventType,omitempty" bhojpur:"required,options=Normal|Warning,default=Warning"`
	ResourceKind string `json:"resourceKind,omitempty" bhojpur:"required,options=Pod|Node|Deployment|StatefulSet|DaemonSet"`
}

type WorkloadRule struct {
	WorkloadID          string            `json:"workloadId,omitempty"`
	Selector            map[string]string `json:"selector,omitempty"`
	AvailablePercentage int               `json:"availablePercentage,omitempty" bhojpur:"required,min=1,max=100,default=70"`
}

type SystemServiceRule struct {
	Condition string `json:"condition,omitempty" bhojpur:"required,options=etcd|controller-manager|scheduler,default=scheduler"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Notifier struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec NotifierSpec `json:"spec"`
	// Most recent observed status of the notifier. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status NotifierStatus `json:"status"`
}

func (n *Notifier) ObjClusterName() string {
	return n.Spec.ObjClusterName()
}

type NotifierSpec struct {
	ClusterName string `json:"clusterName" bhojpur:"type=reference[cluster]"`

	DisplayName     string           `json:"displayName,omitempty" bhojpur:"required"`
	Description     string           `json:"description,omitempty"`
	SendResolved    bool             `json:"sendResolved,omitempty"`
	SMTPConfig      *SMTPConfig      `json:"smtpConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty"`
	WechatConfig    *WechatConfig    `json:"wechatConfig,omitempty"`
	DingtalkConfig  *DingtalkConfig  `json:"dingtalkConfig,omitempty"`
	MSTeamsConfig   *MSTeamsConfig   `json:"msteamsConfig,omitempty"`
}

func (n *NotifierSpec) ObjClusterName() string {
	return n.ClusterName
}

type Notification struct {
	Message         string           `json:"message,omitempty"`
	SMTPConfig      *SMTPConfig      `json:"smtpConfig,omitempty"`
	SlackConfig     *SlackConfig     `json:"slackConfig,omitempty"`
	PagerdutyConfig *PagerdutyConfig `json:"pagerdutyConfig,omitempty"`
	WebhookConfig   *WebhookConfig   `json:"webhookConfig,omitempty"`
	WechatConfig    *WechatConfig    `json:"wechatConfig,omitempty"`
	DingtalkConfig  *DingtalkConfig  `json:"dingtalkConfig,omitempty"`
	MSTeamsConfig   *MSTeamsConfig   `json:"msteamsConfig,omitempty"`
}

type SMTPConfig struct {
	Host             string `json:"host,omitempty" bhojpur:"required,type=hostname"`
	Port             int    `json:"port,omitempty" bhojpur:"required,min=1,max=65535,default=587"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty" bhojpur:"type=password"`
	Sender           string `json:"sender,omitempty" bhojpur:"required"`
	DefaultRecipient string `json:"defaultRecipient,omitempty" bhojpur:"required"`
	TLS              *bool  `json:"tls,omitempty" bhojpur:"required,default=true"`
}

type SlackConfig struct {
	DefaultRecipient string `json:"defaultRecipient,omitempty"`
	URL              string `json:"url,omitempty" bhojpur:"required"`
	*HTTPClientConfig
}

type PagerdutyConfig struct {
	ServiceKey string `json:"serviceKey,omitempty" bhojpur:"required"`
	*HTTPClientConfig
}

type WebhookConfig struct {
	URL string `json:"url,omitempty" bhojpur:"required"`
	*HTTPClientConfig
}

type WechatConfig struct {
	DefaultRecipient string `json:"defaultRecipient,omitempty" bhojpur:"required"`
	Secret           string `json:"secret,omitempty" bhojpur:"type=password,required"`
	Agent            string `json:"agent,omitempty" bhojpur:"required"`
	Corp             string `json:"corp,omitempty" bhojpur:"required"`
	RecipientType    string `json:"recipientType,omitempty" bhojpur:"required,options=tag|party|user,default=party"`
	APIURL           string `json:"apiUrl,omitempty"`
	*HTTPClientConfig
}

type DingtalkConfig struct {
	URL    string `json:"url,omitempty" bhojpur:"required"`
	Secret string `json:"secret,omitempty" bhojpur:"type=password"`
	*HTTPClientConfig
}

type MSTeamsConfig struct {
	URL string `json:"url,omitempty" bhojpur:"required"`
	*HTTPClientConfig
}

type NotifierStatus struct {
	SMTPCredentialSecret     string `json:"smtpCredentialSecret,omitempty" bhojpur:"nocreate,noupdate"`
	WechatCredentialSecret   string `json:"wechatCredentialSecret,omitempty" bhojpur:"nocreate,noupdate"`
	DingtalkCredentialSecret string `json:"dingtalkCredentialSecret,omitempty" bhojpur:"nocreate,noupdate"`
}

// HTTPClientConfig configures an HTTP client.
type HTTPClientConfig struct {
	// HTTP proxy server to use to connect to the targets.
	ProxyURL string `json:"proxyUrl,omitempty"`
}
