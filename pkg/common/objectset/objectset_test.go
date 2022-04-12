package objectset

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
	"testing"

	"github.com/stretchr/testify/assert"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func TestObjectSet_Namespaces(t *testing.T) {
	type fields struct {
		errs        []error
		objects     ObjectByGVK
		objectsByGK ObjectByGK
		order       []runtime.Object
		gvkOrder    []schema.GroupVersionKind
		gvkSeen     map[schema.GroupVersionKind]bool
	}
	tests := []struct {
		name           string
		fields         fields
		wantNamespaces []string
	}{
		{
			name: "empty",
			fields: fields{
				objects: map[schema.GroupVersionKind]map[ObjectKey]runtime.Object{},
			},
			wantNamespaces: nil,
		},
		{
			name: "1 namespace",
			fields: fields{
				objects: map[schema.GroupVersionKind]map[ObjectKey]runtime.Object{
					schema.GroupVersionKind{}: {
						ObjectKey{Namespace: "ns1", Name: "a"}: nil,
						ObjectKey{Namespace: "ns1", Name: "b"}: nil,
					},
				},
			},
			wantNamespaces: []string{"ns1"},
		},
		{
			name: "many namespace",
			fields: fields{
				objects: map[schema.GroupVersionKind]map[ObjectKey]runtime.Object{
					schema.GroupVersionKind{}: {
						ObjectKey{Namespace: "ns1", Name: "a"}: nil,
						ObjectKey{Namespace: "ns2", Name: "b"}: nil,
					},
				},
			},
			wantNamespaces: []string{"ns1", "ns2"},
		},
		{
			name: "missing namespace",
			fields: fields{
				objects: map[schema.GroupVersionKind]map[ObjectKey]runtime.Object{
					schema.GroupVersionKind{}: {
						ObjectKey{Namespace: "ns1", Name: "a"}: nil,
						ObjectKey{Name: "b"}:                   nil,
					},
				},
			},
			wantNamespaces: []string{"", "ns1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &ObjectSet{
				errs:        tt.fields.errs,
				objects:     tt.fields.objects,
				objectsByGK: tt.fields.objectsByGK,
				order:       tt.fields.order,
				gvkOrder:    tt.fields.gvkOrder,
				gvkSeen:     tt.fields.gvkSeen,
			}

			gotNamespaces := o.Namespaces()
			assert.ElementsMatchf(t, tt.wantNamespaces, gotNamespaces, "Namespaces() = %v, want %v", gotNamespaces, tt.wantNamespaces)
		})
	}
}

func TestObjectByKey_Namespaces(t *testing.T) {
	tests := []struct {
		name           string
		objects        ObjectByKey
		wantNamespaces []string
	}{
		{
			name:           "empty",
			objects:        ObjectByKey{},
			wantNamespaces: nil,
		},
		{
			name: "1 namespace",
			objects: ObjectByKey{
				ObjectKey{Namespace: "ns1", Name: "a"}: nil,
				ObjectKey{Namespace: "ns1", Name: "b"}: nil,
			},
			wantNamespaces: []string{"ns1"},
		},
		{
			name: "many namespaces",
			objects: ObjectByKey{
				ObjectKey{Namespace: "ns1", Name: "a"}: nil,
				ObjectKey{Namespace: "ns2", Name: "b"}: nil,
			},
			wantNamespaces: []string{"ns1", "ns2"},
		},
		{
			name: "many namespaces with duplicates",
			objects: ObjectByKey{
				ObjectKey{Namespace: "ns1", Name: "a"}: nil,
				ObjectKey{Namespace: "ns2", Name: "b"}: nil,
				ObjectKey{Namespace: "ns1", Name: "c"}: nil,
			},
			wantNamespaces: []string{"ns1", "ns2"},
		},
		{
			name: "missing namespace",
			objects: ObjectByKey{
				ObjectKey{Namespace: "ns1", Name: "a"}: nil,
				ObjectKey{Name: "b"}:                   nil,
			},
			wantNamespaces: []string{"", "ns1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotNamespaces := tt.objects.Namespaces()
			assert.ElementsMatchf(t, tt.wantNamespaces, gotNamespaces, "Namespaces() = %v, want %v", gotNamespaces, tt.wantNamespaces)
		})
	}
}
