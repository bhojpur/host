package schemas

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

type Schema struct {
	ID                string                 `json:"-"`
	Description       string                 `json:"description,omitempty"`
	CodeName          string                 `json:"-"`
	CodeNamePlural    string                 `json:"-"`
	PkgName           string                 `json:"-"`
	PluralName        string                 `json:"pluralName,omitempty"`
	ResourceMethods   []string               `json:"resourceMethods,omitempty"`
	ResourceFields    map[string]Field       `json:"resourceFields"`
	ResourceActions   map[string]Action      `json:"resourceActions,omitempty"`
	CollectionMethods []string               `json:"collectionMethods,omitempty"`
	CollectionFields  map[string]Field       `json:"collectionFields,omitempty"`
	CollectionActions map[string]Action      `json:"collectionActions,omitempty"`
	Attributes        map[string]interface{} `json:"attributes,omitempty"`

	InternalSchema *Schema `json:"-"`
	Mapper         Mapper  `json:"-"`
}

func (s *Schema) DeepCopy() *Schema {
	r := *s

	if s.ResourceFields != nil {
		r.ResourceFields = map[string]Field{}
		for k, v := range s.ResourceFields {
			r.ResourceFields[k] = v
		}
	}

	if s.ResourceActions != nil {
		r.ResourceActions = map[string]Action{}
		for k, v := range s.ResourceActions {
			r.ResourceActions[k] = v
		}
	}

	if s.CollectionFields != nil {
		r.CollectionFields = map[string]Field{}
		for k, v := range s.CollectionFields {
			r.CollectionFields[k] = v
		}
	}

	if s.CollectionActions != nil {
		r.CollectionActions = map[string]Action{}
		for k, v := range s.CollectionActions {
			r.CollectionActions[k] = v
		}
	}

	if s.Attributes != nil {
		r.Attributes = map[string]interface{}{}
		for k, v := range s.Attributes {
			r.Attributes[k] = v
		}
	}

	if s.InternalSchema != nil {
		r.InternalSchema = r.InternalSchema.DeepCopy()
	}

	return &r
}

type Field struct {
	Type         string      `json:"type,omitempty"`
	Default      interface{} `json:"default,omitempty"`
	Nullable     bool        `json:"nullable,omitempty"`
	Create       bool        `json:"create"`
	WriteOnly    bool        `json:"writeOnly,omitempty"`
	Required     bool        `json:"required,omitempty"`
	Update       bool        `json:"update"`
	MinLength    *int64      `json:"minLength,omitempty"`
	MaxLength    *int64      `json:"maxLength,omitempty"`
	Min          *int64      `json:"min,omitempty"`
	Max          *int64      `json:"max,omitempty"`
	Options      []string    `json:"options,omitempty"`
	ValidChars   string      `json:"validChars,omitempty"`
	InvalidChars string      `json:"invalidChars,omitempty"`
	Description  string      `json:"description,omitempty"`
	CodeName     string      `json:"-"`
}

type Action struct {
	Input  string `json:"input,omitempty"`
	Output string `json:"output,omitempty"`
}
