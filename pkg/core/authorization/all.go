package authorization

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

	"github.com/bhojpur/host/pkg/core/httperror"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/slice"
)

type AllAccess struct {
}

func (*AllAccess) CanCreate(apiContext *types.APIContext, schema *types.Schema) error {
	if slice.ContainsString(schema.CollectionMethods, http.MethodPost) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not create "+schema.ID)
}

func (*AllAccess) CanGet(apiContext *types.APIContext, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, http.MethodGet) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not get "+schema.ID)
}

func (*AllAccess) CanList(apiContext *types.APIContext, schema *types.Schema) error {
	if slice.ContainsString(schema.CollectionMethods, http.MethodGet) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not list "+schema.ID)
}

func (*AllAccess) CanUpdate(apiContext *types.APIContext, obj map[string]interface{}, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, http.MethodPut) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not update "+schema.ID)
}

func (*AllAccess) CanDelete(apiContext *types.APIContext, obj map[string]interface{}, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, http.MethodDelete) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not delete "+schema.ID)
}

func (*AllAccess) CanDo(apiGroup, resource, verb string, apiContext *types.APIContext, obj map[string]interface{}, schema *types.Schema) error {
	if slice.ContainsString(schema.ResourceMethods, verb) {
		return nil
	}
	return httperror.NewAPIError(httperror.PermissionDenied, "can not perform "+verb+" "+schema.ID)
}

func (*AllAccess) Filter(apiContext *types.APIContext, schema *types.Schema, obj map[string]interface{}, context map[string]string) map[string]interface{} {
	return obj
}

func (*AllAccess) FilterList(apiContext *types.APIContext, schema *types.Schema, obj []map[string]interface{}, context map[string]string) []map[string]interface{} {
	return obj
}
