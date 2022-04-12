package summary

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

	"github.com/bhojpur/host/pkg/common/data"
	unstructured2 "github.com/bhojpur/host/pkg/common/unstructured"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type Summary struct {
	State         string                 `json:"state,omitempty"`
	Error         bool                   `json:"error,omitempty"`
	Transitioning bool                   `json:"transitioning,omitempty"`
	Message       []string               `json:"message,omitempty"`
	Attributes    map[string]interface{} `json:"-"`
	Relationships []Relationship         `json:"-"`
}

type Relationship struct {
	Name         string
	Namespace    string
	ControlledBy bool
	Kind         string
	APIVersion   string
	Inbound      bool
	Type         string
	Selector     *metav1.LabelSelector
}

func (s Summary) String() string {
	if !s.Transitioning && !s.Error {
		return s.State
	}
	var msg string
	if s.Transitioning {
		msg = "[progressing"
	}
	if s.Error {
		if len(msg) > 0 {
			msg += ",error]"
		} else {
			msg = "error]"
		}
	} else {
		msg += "]"
	}
	if len(s.Message) > 0 {
		msg = msg + " " + strings.Join(s.Message, ", ")
	}
	return msg
}

func (s Summary) IsReady() bool {
	return !s.Error && !s.Transitioning
}

func (s *Summary) DeepCopy() *Summary {
	v := *s
	return &v
}

func (s *Summary) DeepCopyInto(v *Summary) {
	*v = *s
}

func dedupMessage(messages []string) []string {
	seen := map[string]bool{}
	var result []string

	for _, message := range messages {
		message = strings.TrimSpace(message)
		if message == "" {
			continue
		}
		if seen[message] {
			continue
		}
		seen[message] = true
		result = append(result, message)
	}

	return result
}

func Summarize(runtimeObj runtime.Object) Summary {
	var (
		obj     data.Object
		err     error
		summary Summary
	)

	if s, ok := runtimeObj.(*SummarizedObject); ok {
		return s.Summary
	}

	unstr, ok := runtimeObj.(*unstructured.Unstructured)
	if !ok {
		unstr, err = unstructured2.ToUnstructured(runtimeObj)
		if err != nil {
			return summary
		}
	}

	if unstr != nil {
		obj = unstr.Object
	}

	conditions := getConditions(obj)

	for _, summarizer := range Summarizers {
		summary = summarizer(obj, conditions, summary)
	}

	if summary.State == "" {
		summary.State = "active"
	}

	summary.State = strings.ToLower(summary.State)
	summary.Message = dedupMessage(summary.Message)
	return summary
}
