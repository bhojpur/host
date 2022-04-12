package templates

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
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"github.com/bhojpur/host/pkg/engine/metadata"
	"github.com/bhojpur/host/pkg/engine/types/kdm"
	"github.com/blang/semver"
	"github.com/ghodss/yaml"
	"github.com/sirupsen/logrus"
)

func CompileTemplateFromMap(tmplt string, configMap interface{}) (string, error) {
	out := new(bytes.Buffer)
	templateFuncMap := sprig.TxtFuncMap()
	templateFuncMap["GetKubednsStubDomains"] = GetKubednsStubDomains
	templateFuncMap["toYaml"] = ToYAML
	t := template.Must(template.New("compiled_template").Funcs(templateFuncMap).Parse(tmplt))
	if err := t.Execute(out, configMap); err != nil {
		return "", err
	}
	return out.String(), nil
}

func GetVersionedTemplates(templateName string, data map[string]interface{}, k8sVersion string) (string, error) {
	if template, ok := data[templateName]; ok {
		return convert.ToString(template), nil
	}
	return getTemplate(templateName, k8sVersion)
}

func GetKubednsStubDomains(stubDomains map[string][]string) string {
	json, _ := json.Marshal(stubDomains)
	return string(json)
}

func ToYAML(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		// Swallow errors inside of a template so it doesn't affect remaining template lines
		logrus.Errorf("[ToYAML] Error marshaling %v: %v", v, err)
		return ""
	}
	yamlData, err := yaml.JSONToYAML(data)
	if err != nil {
		// Swallow errors inside of a template so it doesn't affect remaining template lines
		logrus.Errorf("[ToYAML] Error converting json to yaml for %v: %v ", string(data), err)
		return ""
	}
	return strings.TrimSuffix(string(yamlData), "\n")
}

func getTemplate(templateName, k8sVersion string) (string, error) {
	versionData := metadata.K8sVersionToTemplates[templateName]
	toMatch, err := semver.Make(k8sVersion[1:])
	if err != nil {
		return "", fmt.Errorf("k8sVersion not sem-ver %s %v", k8sVersion, err)
	}
	for k := range versionData {
		testRange, err := semver.ParseRange(k)
		if err != nil {
			logrus.Errorf("range for %s not sem-ver %v %v", templateName, testRange, err)
			continue
		}
		if testRange(toMatch) {
			return metadata.K8sVersionToTemplates[kdm.TemplateKeys][versionData[k]], nil
		}
	}
	return "", fmt.Errorf("no %s template found for k8sVersion %s", templateName, k8sVersion)
}
