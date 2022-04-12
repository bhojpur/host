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

import (
	"fmt"
	"path"
	"strings"

	"net/http"
	"os"
	"path/filepath"
	"text/template"

	"github.com/bhojpur/host/pkg/core/generator"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/gengo/args"
)

var (
	outputDir   = "./pkg/generated"
	basePackage = "github.com/bhojpur/host/pkg/apis"
	baseBhojpur = "../client/generated"
	baseK8s     = "bhojpur"
	baseCompose = "compose"
)

func funcs() template.FuncMap {
	return template.FuncMap{
		"capitalize":   convert.Capitalize,
		"unCapitalize": convert.Uncapitalize,
		"upper":        strings.ToUpper,
		"toLower":      strings.ToLower,
		"hasGet":       hasGet,
		"hasPost":      hasPost,
	}
}

func hasGet(schema *types.Schema) bool {
	return contains(schema.CollectionMethods, http.MethodGet)
}

func hasPost(schema *types.Schema) bool {
	return contains(schema.CollectionMethods, http.MethodPost)
}

func contains(list []string, needle string) bool {
	for _, i := range list {
		if i == needle {
			return true
		}
	}
	return false
}

func Generate(schemas *types.Schemas, backendTypes map[string]bool) {
	version := getVersion(schemas)
	group := strings.Split(version.Group, ".")[0]

	bhojpurOutputPackage := path.Join(baseBhojpur, group, version.Version)
	k8sOutputPackage := path.Join(baseK8s, version.Group, version.Version)

	if err := generator.Generate(schemas, backendTypes, basePackage, outputDir, bhojpurOutputPackage, k8sOutputPackage); err != nil {
		panic(err)
	}
}

func GenerateClient(schemas *types.Schemas, backendTypes map[string]bool) {
	version := getVersion(schemas)
	group := strings.Split(version.Group, ".")[0]

	bhojpurOutputPackage := path.Join(baseBhojpur, group, version.Version)

	if err := generator.GenerateClient(schemas, backendTypes, outputDir, bhojpurOutputPackage); err != nil {
		panic(err)
	}
}

func GenerateComposeType(projectSchemas *types.Schemas, managementSchemas *types.Schemas, clusterSchemas *types.Schemas) {
	if err := generateComposeType(filepath.Join(outputDir, baseCompose), projectSchemas, managementSchemas, clusterSchemas); err != nil {
		panic(err)
	}
}

func generateComposeType(baseCompose string, projectSchemas *types.Schemas, managementSchemas *types.Schemas, clusterSchemas *types.Schemas) error {
	outputDir := filepath.Join(args.DefaultSourceTree(), baseCompose)
	if err := os.MkdirAll(outputDir, 0755); err != nil {
		return err
	}
	filePath := "zz_generated_compose.go"
	output, err := os.Create(path.Join(outputDir, filePath))
	if err != nil {
		return err
	}
	defer output.Close()

	typeTemplate, err := template.New("compose.template").
		Funcs(funcs()).
		Parse(strings.Replace(composeTemplate, "%BACK%", "`", -1))
	if err != nil {
		return err
	}

	if err := typeTemplate.Execute(output, map[string]interface{}{
		"managementSchemas": managementSchemas.Schemas(),
		"projectSchemas":    projectSchemas.Schemas(),
		"clusterSchemas":    clusterSchemas.Schemas(),
	}); err != nil {
		return err
	}
	if err := output.Close(); err != nil {
		return err
	}

	return generator.Gofmt(args.DefaultSourceTree(), baseCompose)
}

func GenerateNativeTypes(gv schema.GroupVersion, nsObjs []interface{}, objs []interface{}) {
	version := gv.Version
	group := gv.Group
	groupPath := group

	if groupPath == "" {
		groupPath = "core"
	}

	k8sOutputPackage := path.Join(outputDir, baseK8s, groupPath, version)

	if err := generator.GenerateControllerForTypes(&types.APIVersion{
		Version: version,
		Group:   group,
		Path:    fmt.Sprintf("/k8s/%s-%s", groupPath, version),
	}, k8sOutputPackage, nsObjs, objs); err != nil {
		panic(err)
	}
}

func getVersion(schemas *types.Schemas) *types.APIVersion {
	var version types.APIVersion
	for _, schema := range schemas.Schemas() {
		if version.Group == "" {
			version = schema.Version
			continue
		}
		if version.Group != schema.Version.Group ||
			version.Version != schema.Version.Version {
			panic("schema set contains two APIVersions")
		}
	}

	return &version
}
