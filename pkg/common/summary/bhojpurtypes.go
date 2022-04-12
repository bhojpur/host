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
)

func checkBhojpurReady(obj data.Object, condition []Condition, summary Summary) Summary {
	if strings.Contains(obj.String("apiVersion"), "bhojpur.net/") {
		for _, condition := range condition {
			if condition.Type() == "Ready" && condition.Status() == "False" && condition.Message() != "" {
				summary.Message = append(summary.Message, condition.Message())
				summary.Error = true
				return summary
			}
		}
	}

	return summary
}

func checkBhojpurTypes(obj data.Object, condition []Condition, summary Summary) Summary {
	return checkRelease(obj, condition, summary)
}

func checkRelease(obj data.Object, _ []Condition, summary Summary) Summary {
	if !isKind(obj, "App", "catalog.bhojpur.net") {
		return summary
	}
	if obj.String("status", "summary", "state") != "deployed" {
		return summary
	}
	for _, resources := range obj.Slice("spec", "resources") {
		summary.Relationships = append(summary.Relationships, Relationship{
			Name:       resources.String("name"),
			Kind:       resources.String("kind"),
			APIVersion: resources.String("apiVersion"),
			Type:       "helmresource",
		})
	}
	return summary
}
