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

type ClusterLogging struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec ClusterLoggingSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status ClusterLoggingStatus `json:"status"`
}

func (c *ClusterLogging) ObjClusterName() string {
	return c.Spec.ObjClusterName()
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProjectLogging struct {
	types.Namespaced

	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec ProjectLoggingSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status ProjectLoggingStatus `json:"status"`
}

func (p *ProjectLogging) ObjClusterName() string {
	return p.Spec.ObjClusterName()
}

type LoggingCommonField struct {
	DisplayName         string            `json:"displayName,omitempty"`
	OutputFlushInterval int               `json:"outputFlushInterval,omitempty" bhojpur:"default=60"`
	OutputTags          map[string]string `json:"outputTags,omitempty"`
	EnableJSONParsing   bool              `json:"enableJSONParsing,omitempty"`
}

type LoggingTargets struct {
	ElasticsearchConfig   *ElasticsearchConfig   `json:"elasticsearchConfig,omitempty"`
	SplunkConfig          *SplunkConfig          `json:"splunkConfig,omitempty"`
	KafkaConfig           *KafkaConfig           `json:"kafkaConfig,omitempty"`
	SyslogConfig          *SyslogConfig          `json:"syslogConfig,omitempty"`
	FluentForwarderConfig *FluentForwarderConfig `json:"fluentForwarderConfig,omitempty"`
	CustomTargetConfig    *CustomTargetConfig    `json:"customTargetConfig,omitempty"`
}

type ClusterLoggingSpec struct {
	LoggingTargets
	LoggingCommonField
	ClusterName            string `json:"clusterName" bhojpur:"type=reference[cluster]"`
	IncludeSystemComponent *bool  `json:"includeSystemComponent,omitempty" bhojpur:"default=true"`
}

func (c *ClusterLoggingSpec) ObjClusterName() string {
	return c.ClusterName
}

type ProjectLoggingSpec struct {
	LoggingTargets
	LoggingCommonField
	ProjectName string `json:"projectName" bhojpur:"type=reference[project]"`
}

func (p *ProjectLoggingSpec) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}

type ClusterLoggingStatus struct {
	Conditions  []LoggingCondition  `json:"conditions,omitempty"`
	AppliedSpec ClusterLoggingSpec  `json:"appliedSpec,omitempty"`
	FailedSpec  *ClusterLoggingSpec `json:"failedSpec,omitempty"`
}

type ProjectLoggingStatus struct {
	Conditions  []LoggingCondition `json:"conditions,omitempty"`
	AppliedSpec ProjectLoggingSpec `json:"appliedSpec,omitempty"`
}

var (
	LoggingConditionProvisioned condition.Cond = "Provisioned"
	LoggingConditionUpdated     condition.Cond = "Updated"
)

type LoggingCondition struct {
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

type ElasticsearchConfig struct {
	Endpoint      string `json:"endpoint,omitempty" bhojpur:"required"`
	IndexPrefix   string `json:"indexPrefix,omitempty" bhojpur:"required"`
	DateFormat    string `json:"dateFormat,omitempty" bhojpur:"required,type=enum,options=YYYY-MM-DD|YYYY-MM|YYYY,default=YYYY-MM-DD"`
	AuthUserName  string `json:"authUsername,omitempty"`
	AuthPassword  string `json:"authPassword,omitempty" bhojpur:"type=password"`
	Certificate   string `json:"certificate,omitempty"`
	ClientCert    string `json:"clientCert,omitempty"`
	ClientKey     string `json:"clientKey,omitempty"`
	ClientKeyPass string `json:"clientKeyPass,omitempty"`
	SSLVerify     bool   `json:"sslVerify,omitempty"`
	SSLVersion    string `json:"sslVersion,omitempty" bhojpur:"type=enum,options=SSLv23|TLSv1|TLSv1_1|TLSv1_2,default=TLSv1_2"`
}

type SplunkConfig struct {
	Endpoint      string `json:"endpoint,omitempty" bhojpur:"required"`
	Source        string `json:"source,omitempty"`
	Token         string `json:"token,omitempty" bhojpur:"required,type=password"`
	Certificate   string `json:"certificate,omitempty"`
	ClientCert    string `json:"clientCert,omitempty"`
	ClientKey     string `json:"clientKey,omitempty"`
	ClientKeyPass string `json:"clientKeyPass,omitempty"`
	SSLVerify     bool   `json:"sslVerify,omitempty"`
	Index         string `json:"index,omitempty"`
}

type KafkaConfig struct {
	ZookeeperEndpoint  string   `json:"zookeeperEndpoint,omitempty"`
	BrokerEndpoints    []string `json:"brokerEndpoints,omitempty"`
	Topic              string   `json:"topic,omitempty" bhojpur:"required"`
	Certificate        string   `json:"certificate,omitempty"`
	ClientCert         string   `json:"clientCert,omitempty"`
	ClientKey          string   `json:"clientKey,omitempty"`
	SaslUsername       string   `json:"saslUsername,omitempty"`
	SaslPassword       string   `json:"saslPassword,omitempty" bhojpur:"type=password"`
	SaslScramMechanism string   `json:"saslScramMechanism,omitempty" bhojpur:"type=enum,options=sha256|sha512"`
	SaslType           string   `json:"saslType,omitempty" bhojpur:"type=enum,options=plain|scram"`
}

type SyslogConfig struct {
	Endpoint    string `json:"endpoint,omitempty" bhojpur:"required"`
	Severity    string `json:"severity,omitempty" bhojpur:"default=notice,type=enum,options=emerg|alert|crit|err|warning|notice|info|debug"`
	Program     string `json:"program,omitempty"`
	Protocol    string `json:"protocol,omitempty" bhojpur:"default=udp,type=enum,options=udp|tcp"`
	Token       string `json:"token,omitempty" bhojpur:"type=password"`
	EnableTLS   bool   `json:"enableTls,omitempty" bhojpur:"default=false"`
	Certificate string `json:"certificate,omitempty"`
	ClientCert  string `json:"clientCert,omitempty"`
	ClientKey   string `json:"clientKey,omitempty"`
	SSLVerify   bool   `json:"sslVerify,omitempty"`
}

type FluentForwarderConfig struct {
	EnableTLS     bool           `json:"enableTls,omitempty" bhojpur:"default=false"`
	Certificate   string         `json:"certificate,omitempty"`
	ClientCert    string         `json:"clientCert,omitempty"`
	ClientKey     string         `json:"clientKey,omitempty"`
	ClientKeyPass string         `json:"clientKeyPass,omitempty"`
	SSLVerify     bool           `json:"sslVerify,omitempty"`
	Compress      *bool          `json:"compress,omitempty" bhojpur:"default=true"`
	FluentServers []FluentServer `json:"fluentServers,omitempty" bhojpur:"required"`
}

type FluentServer struct {
	Endpoint  string `json:"endpoint,omitempty" bhojpur:"required"`
	Hostname  string `json:"hostname,omitempty"`
	Weight    int    `json:"weight,omitempty" bhojpur:"default=100"`
	Standby   bool   `json:"standby,omitempty" bhojpur:"default=false"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty" bhojpur:"type=password"`
	SharedKey string `json:"sharedKey,omitempty" bhojpur:"type=password"`
}

type CustomTargetConfig struct {
	Content     string `json:"content,omitempty"`
	Certificate string `json:"certificate,omitempty"`
	ClientCert  string `json:"clientCert,omitempty"`
	ClientKey   string `json:"clientKey,omitempty"`
}

type ClusterTestInput struct {
	ClusterName string `json:"clusterId" bhojpur:"required,type=reference[cluster]"`
	LoggingTargets
	OutputTags map[string]string `json:"outputTags,omitempty"`
}

func (c *ClusterTestInput) ObjClusterName() string {
	return c.ClusterName
}

type ProjectTestInput struct {
	ProjectName string `json:"projectId" bhojpur:"required,type=reference[project]"`
	LoggingTargets
	OutputTags map[string]string `json:"outputTags,omitempty"`
}

func (p *ProjectTestInput) ObjClusterName() string {
	if parts := strings.SplitN(p.ProjectName, ":", 2); len(parts) == 2 {
		return parts[0]
	}
	return ""
}
