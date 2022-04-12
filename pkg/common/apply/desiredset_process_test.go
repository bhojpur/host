package apply

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
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic/fake"
	k8stesting "k8s.io/client-go/testing"
)

func Test_multiNamespaceList(t *testing.T) {
	results := map[string]*unstructured.UnstructuredList{
		"ns1": {Items: []unstructured.Unstructured{
			{Object: map[string]interface{}{"name": "o1", "namespace": "ns1"}},
			{Object: map[string]interface{}{"name": "o2", "namespace": "ns1"}},
			{Object: map[string]interface{}{"name": "o3", "namespace": "ns1"}},
		}},
		"ns2": {Items: []unstructured.Unstructured{
			{Object: map[string]interface{}{"name": "o4", "namespace": "ns2"}},
			{Object: map[string]interface{}{"name": "o5", "namespace": "ns2"}},
		}},
		"ns3": {Items: []unstructured.Unstructured{}},
	}

	baseClient := fake.NewSimpleDynamicClient(runtime.NewScheme())
	baseClient.PrependReactor("list", "*", func(action k8stesting.Action) (handled bool, ret runtime.Object, err error) {
		if strings.Contains(action.GetNamespace(), "error") {
			return true, nil, errors.New("simulated failure")
		}

		return true, results[action.GetNamespace()], nil
	})

	type args struct {
		namespaces []string
	}
	tests := []struct {
		name          string
		args          args
		expectedCalls int
		expectError   bool
	}{
		{
			name: "no namespaces",
			args: args{
				namespaces: []string{},
			},
			expectError:   false,
			expectedCalls: 0,
		},
		{
			name: "1 namespace",
			args: args{
				namespaces: []string{"ns1"},
			},
			expectError:   false,
			expectedCalls: 3,
		},
		{
			name: "many namespaces",
			args: args{
				namespaces: []string{"ns1", "ns2", "ns3"},
			},
			expectError:   false,
			expectedCalls: 5,
		},
		{
			name: "1 namespace error",
			args: args{
				namespaces: []string{"error", "ns2", "ns3"},
			},
			expectError:   true,
			expectedCalls: -1,
		},
		{
			name: "many namespace errors",
			args: args{
				namespaces: []string{"error", "error1", "error2"},
			},
			expectError:   true,
			expectedCalls: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var calls int
			err := multiNamespaceList(context.TODO(), tt.args.namespaces, baseClient.Resource(schema.GroupVersionResource{}), labels.NewSelector(), func(obj unstructured.Unstructured) {
				calls += 1
			})

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tt.expectedCalls >= 0 {
				assert.Equal(t, tt.expectedCalls, calls)
			}
		})
	}
}
