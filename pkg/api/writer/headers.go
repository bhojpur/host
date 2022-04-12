package writer

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
	"github.com/bhojpur/host/pkg/api/builtin"
	"github.com/bhojpur/host/pkg/core/types"
)

func AddCommonResponseHeader(apiContext *types.APIContext) error {
	addExpires(apiContext)
	return addSchemasHeader(apiContext)
}

func addSchemasHeader(apiContext *types.APIContext) error {
	schema := apiContext.Schemas.Schema(&builtin.Version, "schema")
	if schema == nil {
		return nil
	}

	version := apiContext.SchemasVersion
	if version == nil {
		version = apiContext.Version
	}

	apiContext.Response.Header().Set("X-Api-Schemas", apiContext.URLBuilder.Collection(schema, version))
	return nil
}

func addExpires(apiContext *types.APIContext) {
	apiContext.Response.Header().Set("Expires", "Wed 24 Feb 1982 18:42:00 GMT")
}
