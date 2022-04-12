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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type DynamicSchema struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object’s metadata. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#metadata
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// Specification of the desired behavior of the the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Spec DynamicSchemaSpec `json:"spec"`
	// Most recent observed status of the cluster. More info:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/api-conventions.md#spec-and-status
	Status DynamicSchemaStatus `json:"status"`
}

type DynamicSchemaSpec struct {
	SchemaName           string            `json:"schemaName,omitempty"`
	Embed                bool              `json:"embed,omitempty"`
	EmbedType            string            `json:"embedType,omitempty"`
	PluralName           string            `json:"pluralName,omitempty"`
	ResourceMethods      []string          `json:"resourceMethods,omitempty"`
	ResourceFields       map[string]Field  `json:"resourceFields,omitempty"`
	ResourceActions      map[string]Action `json:"resourceActions,omitempty"`
	CollectionMethods    []string          `json:"collectionMethods,omitempty"`
	CollectionFields     map[string]Field  `json:"collectionFields,omitempty"`
	CollectionActions    map[string]Action `json:"collectionActions,omitempty"`
	CollectionFilters    map[string]Filter `json:"collectionFilters,omitempty"`
	IncludeableLinks     []string          `json:"includeableLinks,omitempty"`
	DynamicSchemaVersion string            `json:"dynamicSchemaVersion,omitempty"`
}

type DynamicSchemaStatus struct {
	Fake string `json:"fake,omitempty"`
}

type Field struct {
	Type         string   `json:"type,omitempty"`
	Default      Values   `json:"default,omitempty"`
	Unique       bool     `json:"unique,omitempty"`
	Nullable     bool     `json:"nullable,omitempty"`
	Create       bool     `json:"create,omitempty"`
	Required     bool     `json:"required,omitempty"`
	Update       bool     `json:"update,omitempty"`
	MinLength    int64    `json:"minLength,omitempty"`
	MaxLength    int64    `json:"maxLength,omitempty"`
	Min          int64    `json:"min,omitempty"`
	Max          int64    `json:"max,omitempty"`
	Options      []string `json:"options,omitempty"`
	ValidChars   string   `json:"validChars,omitempty"`
	InvalidChars string   `json:"invalidChars,omitempty"`
	Description  string   `json:"description,omitempty"`
	DynamicField bool     `json:"dynamicField,omitempty"`
}

type Values struct {
	StringValue      string   `json:"stringValue"`
	IntValue         int      `json:"intValue"`
	BoolValue        bool     `json:"boolValue"`
	StringSliceValue []string `json:"stringSliceValue"`
}

type Action struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}

type Filter struct {
	Modifiers []string `json:"modifiers,omitempty"`
}

type ListOpts struct {
	Filters map[string]string `json:"filters,omitempty"`
}
