package generator

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

var typeTemplate = `package client

{{$intstr := false}}
{{- range $key, $value := .structFields}}
	{{if eq $value.Type "intstr.IntOrString" }}
		{{$intstr = true}}
	{{end}}
{{end}}

import (
{{- if .schema | hasGet }}
	"github.com/bhojpur/host/pkg/core/types"
{{- end}}
	{{if $intstr  }}
		"k8s.io/apimachinery/pkg/util/intstr"
	{{end}}
)

const (
    {{.schema.CodeName}}Type = "{{.schema.ID}}"
{{- range $key, $value := .structFields}}
	{{$.schema.CodeName}}Field{{$key}} = "{{$value.Name}}"
{{- end}}
)

type {{.schema.CodeName}} struct {
{{- if .schema | hasGet }}
    types.Resource
{{- end}}
    {{- range $key, $value := .structFields}}
        {{$key}} {{$value.Type}} %BACK%json:"{{$value.Name}},omitempty" yaml:"{{$value.Name}},omitempty"%BACK%
    {{- end}}
}

{{ if .schema | hasGet }}
type {{.schema.CodeName}}Collection struct {
    types.Collection
    Data []{{.schema.CodeName}} %BACK%json:"data,omitempty"%BACK%
    client *{{.schema.CodeName}}Client
}

type {{.schema.CodeName}}Client struct {
    apiClient *Client
}

type {{.schema.CodeName}}Operations interface {
    List(opts *types.ListOpts) (*{{.schema.CodeName}}Collection, error)
    ListAll(opts *types.ListOpts) (*{{.schema.CodeName}}Collection, error)
    Create(opts *{{.schema.CodeName}}) (*{{.schema.CodeName}}, error)
    Update(existing *{{.schema.CodeName}}, updates interface{}) (*{{.schema.CodeName}}, error)
    Replace(existing *{{.schema.CodeName}}) (*{{.schema.CodeName}}, error)
    ByID(id string) (*{{.schema.CodeName}}, error)
    Delete(container *{{.schema.CodeName}}) error
    {{range $key, $value := .resourceActions}}
        {{if (and (eq $value.Input "") (eq $value.Output ""))}}
            Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}) (error)
        {{else if (and (eq $value.Input "") (ne $value.Output ""))}}
            Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}) (*{{.Output | capitalize}}, error)
        {{else if (and (ne $value.Input "") (eq $value.Output ""))}}
            Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}, input *{{$value.Input | capitalize}}) (error)
        {{else}}
            Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}, input *{{$value.Input | capitalize}}) (*{{.Output | capitalize}}, error)
        {{end}}
	{{end}}
    {{range $key, $value := .collectionActions}}
        {{if (and (eq $value.Input "") (eq $value.Output ""))}}
            CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection) (error)
        {{else if (and (eq $value.Input "") (ne $value.Output ""))}}
            CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection) (*{{getCollectionOutput $value.Output $.schema.CodeName}}, error)
        {{else if (and (ne $value.Input "") (eq $value.Output ""))}}
            CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection, input *{{$value.Input | capitalize}}) (error)
        {{else}}
            CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection, input *{{$value.Input | capitalize}}) (*{{getCollectionOutput $value.Output $.schema.CodeName}}, error)
        {{end}}
	{{end}}
}

func new{{.schema.CodeName}}Client(apiClient *Client) *{{.schema.CodeName}}Client {
    return &{{.schema.CodeName}}Client{
        apiClient: apiClient,
    }
}

func (c *{{.schema.CodeName}}Client) Create(container *{{.schema.CodeName}}) (*{{.schema.CodeName}}, error) {
    resp := &{{.schema.CodeName}}{}
    err := c.apiClient.Ops.DoCreate({{.schema.CodeName}}Type, container, resp)
    return resp, err
}

func (c *{{.schema.CodeName}}Client) Update(existing *{{.schema.CodeName}}, updates interface{}) (*{{.schema.CodeName}}, error) {
    resp := &{{.schema.CodeName}}{}
    err := c.apiClient.Ops.DoUpdate({{.schema.CodeName}}Type, &existing.Resource, updates, resp)
    return resp, err
}

func (c *{{.schema.CodeName}}Client) Replace(obj *{{.schema.CodeName}}) (*{{.schema.CodeName}}, error) {
	resp := &{{.schema.CodeName}}{}
	err := c.apiClient.Ops.DoReplace({{.schema.CodeName}}Type, &obj.Resource, obj, resp)
	return resp, err
}

func (c *{{.schema.CodeName}}Client) List(opts *types.ListOpts) (*{{.schema.CodeName}}Collection, error) {
    resp := &{{.schema.CodeName}}Collection{}
    err := c.apiClient.Ops.DoList({{.schema.CodeName}}Type, opts, resp)
    resp.client = c
    return resp, err
}

func (c *{{.schema.CodeName}}Client) ListAll(opts *types.ListOpts) (*{{.schema.CodeName}}Collection, error) {
    resp := &{{.schema.CodeName}}Collection{}
    resp, err := c.List(opts)
    if err != nil {
        return resp, err
    }
    data := resp.Data
    for next, err := resp.Next(); next != nil && err == nil; next, err = next.Next() {
        data = append(data, next.Data...)
        resp = next
        resp.Data = data
    }
    if err != nil {
        return resp, err
    }
    return resp, err
}

func (cc *{{.schema.CodeName}}Collection) Next() (*{{.schema.CodeName}}Collection, error) {
    if cc != nil && cc.Pagination != nil && cc.Pagination.Next != "" {
        resp := &{{.schema.CodeName}}Collection{}
        err := cc.client.apiClient.Ops.DoNext(cc.Pagination.Next, resp)
        resp.client = cc.client
        return resp, err
    }
    return nil, nil
}

func (c *{{.schema.CodeName}}Client) ByID(id string) (*{{.schema.CodeName}}, error) {
    resp := &{{.schema.CodeName}}{}
    err := c.apiClient.Ops.DoByID({{.schema.CodeName}}Type, id, resp)
    return resp, err
}

func (c *{{.schema.CodeName}}Client) Delete(container *{{.schema.CodeName}}) error {
    return c.apiClient.Ops.DoResourceDelete({{.schema.CodeName}}Type, &container.Resource)
}

{{range $key, $value := .resourceActions}}
    {{if (and (eq $value.Input "") (eq $value.Output ""))}}
        func (c *{{$.schema.CodeName}}Client) Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}) (error) {
            err := c.apiClient.Ops.DoAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Resource, nil, nil)
            return err
    {{else if (and (eq $value.Input "") (ne $value.Output ""))}}
        func (c *{{$.schema.CodeName}}Client) Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}) (*{{.Output | capitalize}}, error) {
            resp := &{{.Output | capitalize}}{}
            err := c.apiClient.Ops.DoAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Resource, nil, resp)
            return resp, err
    {{else if (and (ne $value.Input "") (eq $value.Output ""))}}
        func (c *{{$.schema.CodeName}}Client) Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}, input *{{$value.Input | capitalize}}) (error) {
            err := c.apiClient.Ops.DoAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Resource, input, nil)
            return err
    {{else}}
        func (c *{{$.schema.CodeName}}Client) Action{{$key | capitalize}} (resource *{{$.schema.CodeName}}, input *{{$value.Input | capitalize}}) (*{{.Output | capitalize}}, error) {
            resp := &{{.Output | capitalize}}{}
            err := c.apiClient.Ops.DoAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Resource, input, resp)
            return resp, err
    {{- end -}}
    }
{{end}}

{{range $key, $value := .collectionActions}}
    {{if (and (eq $value.Input "") (eq $value.Output ""))}}
        func (c *{{$.schema.CodeName}}Client) CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection) (error) {
			err := c.apiClient.Ops.DoCollectionAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Collection, nil, nil)
			return err
    {{else if (and (eq $value.Input "") (ne $value.Output ""))}}
        func (c *{{$.schema.CodeName}}Client) CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection) (*{{getCollectionOutput $value.Output $.schema.CodeName}}, error) {
			resp := &{{getCollectionOutput $value.Output $.schema.CodeName}}{}
			err := c.apiClient.Ops.DoCollectionAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Collection, nil, resp)
			return resp, err
	{{else if (and (ne $value.Input "") (eq $value.Output ""))}}
		func (c *{{$.schema.CodeName}}Client) CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection, input *{{$value.Input | capitalize}}) (error) {
			err := c.apiClient.Ops.DoCollectionAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Collection, input, nil)
    		return err
	{{else}}
        func (c *{{$.schema.CodeName}}Client) CollectionAction{{$key | capitalize}} (resource *{{$.schema.CodeName}}Collection, input *{{$value.Input | capitalize}}) (*{{getCollectionOutput $value.Output $.schema.CodeName}}, error) {
			resp := &{{getCollectionOutput $value.Output $.schema.CodeName}}{}
			err := c.apiClient.Ops.DoCollectionAction({{$.schema.CodeName}}Type, "{{$key}}", &resource.Collection, input, resp)
    		return resp, err
    {{- end -}}
    }
{{end}}
{{end}}`
