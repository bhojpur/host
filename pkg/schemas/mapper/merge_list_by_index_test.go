package mapper

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
	"reflect"
	"testing"
)

var (
	metrics = []map[string]interface{}{
		{"type": "Resource", "resource": "abc"},
		{"type": "Object", "object": "def"},
	}
	currentMetrics = []map[string]interface{}{
		{"type": "Resource", "currentResource": "tuvw"},
		{"type": "Object", "currentObject": "xyz"},
	}
	origin = map[string]interface{}{
		"metrics":        metrics,
		"currentMetrics": currentMetrics,
	}
)

func Test_MergeList(t *testing.T) {
	mapper := NewMergeListByIndexMapper("currentMetrics", "metrics", "type")
	mapper.fromFields = []string{"type", "currentResource", "currentObject"}
	internal := map[string]interface{}{
		"metrics":        metrics,
		"currentMetrics": currentMetrics,
	}
	mapper.FromInternal(internal)

	if err := mapper.ToInternal(internal); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(internal, origin) {
		t.Fatal("merge list not match after parse")
	}
}
