package schema

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
	"github.com/bhojpur/host/pkg/core/types"
	m "github.com/bhojpur/host/pkg/core/types/mapper"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var (
	handlerMapper = &m.UnionEmbed{
		Fields: []m.UnionMapping{
			{
				FieldName:   "exec",
				CheckFields: []string{"command"},
			},
			{
				FieldName:   "tcpSocket",
				CheckFields: []string{"tcp", "port"},
			},
			{
				FieldName:   "httpGet",
				CheckFields: []string{"port"},
			},
		},
	}
)

type ScalingGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              interface{} `json:"spec"`
	Status            interface{} `json:"status"`
}

type handlerOverride struct {
	TCP bool
}

type EnvironmentFrom struct {
	Source     string `bhojpur:"type=enum,options=field|resource|configMap|secret"`
	SourceName string
	SourceKey  string
	Prefix     string
	Optional   bool
	TargetKey  string
}

type Scheduling struct {
	Node              *NodeScheduling
	Tolerate          []v1.Toleration
	Scheduler         string
	Priority          *int64
	PriorityClassName string
}

type NodeScheduling struct {
	NodeName   string `json:"nodeName" bhojpur:"type=reference[/v3/schemas/node]"`
	RequireAll []string
	RequireAny []string
	Preferred  []string
}

type projectOverride struct {
	types.Namespaced
	ProjectID string `bhojpur:"type=reference[/v3/schemas/project],noupdate"`
}

type Target struct {
	Addresses         []string `json:"addresses"`
	NotReadyAddresses []string `json:"notReadyAddresses"`
	Port              *int32   `json:"port"`
	Protocol          string   `json:"protocol" bhojpur:"type=enum,options=TCP|UDP"`
}
