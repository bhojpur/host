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

var composeTemplate = `package compose

import (
	clusterClient "github.com/bhojpur/host/pkg/client/generated/cluster/v3"
	managementClient "github.com/bhojpur/host/pkg/client/generated/management/v3"
	projectClient "github.com/bhojpur/host/pkg/client/generated/project/v3"
)

type Config struct {
	Version string %BACK%yaml:"version,omitempty"%BACK%

	// Management Client
	{{range .managementSchemas}}
    {{- if . | hasPost }}{{.CodeName}}s map[string]managementClient.{{.CodeName}} %BACK%json:"{{.PluralName}},omitempty" yaml:"{{.PluralName}},omitempty"%BACK%
{{end}}{{end}}

	// Cluster Client
	{{range .clusterSchemas}}
	{{- if . | hasGet }}{{.CodeName}}s map[string]clusterClient.{{.CodeName}} %BACK%json:"{{.PluralName}},omitempty" yaml:"{{.PluralName}},omitempty"%BACK%
{{end}}{{end}}

	// Project Client
	{{range .projectSchemas}}
	{{- if . | hasGet }}{{.CodeName}}s map[string]projectClient.{{.CodeName}} %BACK%json:"{{.PluralName}},omitempty" yaml:"{{.PluralName}},omitempty"%BACK%
{{end}}{{end}}}`
