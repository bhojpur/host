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
	"github.com/bhojpur/host/pkg/core/store/empty"
	"github.com/bhojpur/host/pkg/core/types"
)

func APIRootFormatter(apiContext *types.APIContext, resource *types.RawResource) {
	path, _ := resource.Values["path"].(string)
	if path == "" {
		return
	}

	delete(resource.Values, "path")

	resource.Links["root"] = apiContext.URLBuilder.RelativeToRoot(path)

	data, _ := resource.Values["apiVersion"].(map[string]interface{})
	apiVersion := apiVersionFromMap(apiContext.Schemas, data)

	resource.Links["self"] = apiContext.URLBuilder.Version(apiVersion)

	for _, schema := range apiContext.Schemas.SchemasForVersion(apiVersion) {
		addCollectionLink(apiContext, schema, resource.Links)
	}

	return
}

func addCollectionLink(apiContext *types.APIContext, schema *types.Schema, links map[string]string) {
	collectionLink := getSchemaCollectionLink(apiContext, schema, nil)
	if collectionLink != "" {
		links[schema.PluralName] = collectionLink
	}
}

type APIRootStore struct {
	empty.Store
	roots []string
}

func NewAPIRootStore(roots []string) types.Store {
	return &APIRootStore{roots: roots}
}

func (a *APIRootStore) ByID(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	for _, version := range apiContext.Schemas.Versions() {
		if version.Path == id {
			return apiVersionToAPIRootMap(version), nil
		}
	}
	return nil, nil
}

func (a *APIRootStore) List(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) ([]map[string]interface{}, error) {
	var roots []map[string]interface{}

	for _, version := range apiContext.Schemas.Versions() {
		roots = append(roots, apiVersionToAPIRootMap(version))
	}

	for _, root := range a.roots {
		roots = append(roots, map[string]interface{}{
			"path": root,
		})
	}

	return roots, nil
}

func apiVersionToAPIRootMap(version types.APIVersion) map[string]interface{} {
	return map[string]interface{}{
		"type": "/meta/schemas/apiRoot",
		"apiVersion": map[string]interface{}{
			"version": version.Version,
			"group":   version.Group,
			"path":    version.Path,
		},
		"path": version.Path,
	}
}
