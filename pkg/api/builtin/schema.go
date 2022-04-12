package builtin

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
	"net/http"

	"github.com/bhojpur/host/pkg/core/store/schema"
	"github.com/bhojpur/host/pkg/core/types"
)

var (
	Version = types.APIVersion{
		Group:   "meta.bhojpur.net",
		Version: "v1",
		Path:    "/meta",
	}

	Schema = types.Schema{
		ID:                "schema",
		PluralName:        "schemas",
		Version:           Version,
		CollectionMethods: []string{"GET"},
		ResourceMethods:   []string{"GET"},
		ResourceFields: map[string]types.Field{
			"collectionActions": {Type: "map[json]"},
			"collectionFields":  {Type: "map[json]"},
			"collectionFilters": {Type: "map[json]"},
			"collectionMethods": {Type: "array[string]"},
			"pluralName":        {Type: "string"},
			"resourceActions":   {Type: "map[json]"},
			"resourceFields":    {Type: "map[json]"},
			"resourceMethods":   {Type: "array[string]"},
			"version":           {Type: "map[json]"},
		},
		Formatter: SchemaFormatter,
		Store:     schema.NewSchemaStore(),
	}

	Error = types.Schema{
		ID:                "error",
		Version:           Version,
		ResourceMethods:   []string{},
		CollectionMethods: []string{},
		ResourceFields: map[string]types.Field{
			"code":      {Type: "string"},
			"detail":    {Type: "string", Nullable: true},
			"message":   {Type: "string", Nullable: true},
			"fieldName": {Type: "string", Nullable: true},
			"status":    {Type: "int"},
		},
	}

	Collection = types.Schema{
		ID:                "collection",
		Version:           Version,
		ResourceMethods:   []string{},
		CollectionMethods: []string{},
		ResourceFields: map[string]types.Field{
			"data":       {Type: "array[json]"},
			"pagination": {Type: "map[json]"},
			"sort":       {Type: "map[json]"},
			"filters":    {Type: "map[json]"},
		},
	}

	APIRoot = types.Schema{
		ID:                "apiRoot",
		Version:           Version,
		CollectionMethods: []string{"GET"},
		ResourceMethods:   []string{"GET"},
		ResourceFields: map[string]types.Field{
			"apiVersion": {Type: "map[json]"},
			"path":       {Type: "string"},
		},
		Formatter: APIRootFormatter,
		Store:     NewAPIRootStore(nil),
	}

	Schemas = types.NewSchemas().
		AddSchema(Schema).
		AddSchema(Error).
		AddSchema(Collection).
		AddSchema(APIRoot)
)

func apiVersionFromMap(schemas *types.Schemas, apiVersion map[string]interface{}) types.APIVersion {
	path, _ := apiVersion["path"].(string)
	version, _ := apiVersion["version"].(string)
	group, _ := apiVersion["group"].(string)

	apiVersionObj := types.APIVersion{
		Path:    path,
		Version: version,
		Group:   group,
	}

	for _, testVersion := range schemas.Versions() {
		if testVersion.Equals(&apiVersionObj) {
			return testVersion
		}
	}

	return apiVersionObj
}

func SchemaFormatter(apiContext *types.APIContext, resource *types.RawResource) {
	data, _ := resource.Values["version"].(map[string]interface{})
	apiVersion := apiVersionFromMap(apiContext.Schemas, data)

	schema := apiContext.Schemas.Schema(&apiVersion, resource.ID)
	if schema == nil {
		return
	}

	collectionLink := getSchemaCollectionLink(apiContext, schema, &apiVersion)
	if collectionLink != "" {
		resource.Links["collection"] = collectionLink
	}

	resource.Links["self"] = apiContext.URLBuilder.SchemaLink(schema)
}

func getSchemaCollectionLink(apiContext *types.APIContext, schema *types.Schema, apiVersion *types.APIVersion) string {
	if schema != nil && contains(schema.CollectionMethods, http.MethodGet) {
		return apiContext.URLBuilder.Collection(schema, apiVersion)
	}
	return ""
}

func contains(list []string, needle string) bool {
	for _, v := range list {
		if v == needle {
			return true
		}
	}
	return false
}
