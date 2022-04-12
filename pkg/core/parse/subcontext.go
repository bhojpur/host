package parse

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
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

type DefaultSubContextAttributeProvider struct {
}

func (d *DefaultSubContextAttributeProvider) Query(apiContext *types.APIContext, schema *types.Schema) []*types.QueryCondition {
	var result []*types.QueryCondition

	for name, value := range d.create(apiContext, schema) {
		result = append(result, types.NewConditionFromString(name, types.ModifierEQ, value))
	}

	return result
}

func (d *DefaultSubContextAttributeProvider) Create(apiContext *types.APIContext, schema *types.Schema) map[string]interface{} {
	result := map[string]interface{}{}
	for key, value := range d.create(apiContext, schema) {
		result[key] = value
	}
	return result
}

func (d *DefaultSubContextAttributeProvider) create(apiContext *types.APIContext, schema *types.Schema) map[string]string {
	result := map[string]string{}

	for subContextSchemaID, value := range apiContext.SubContext {
		subContextSchema := apiContext.Schemas.Schema(nil, subContextSchemaID)
		if subContextSchema == nil {
			continue
		}

		ref := convert.ToReference(subContextSchema.ID)
		fullRef := convert.ToFullReference(subContextSchema.Version.Path, subContextSchema.ID)

		for name, field := range schema.ResourceFields {
			if field.Type == ref || field.Type == fullRef {
				result[name] = value
				break
			}
		}
	}

	return result
}
