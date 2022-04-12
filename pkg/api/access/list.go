package access

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

	"github.com/bhojpur/host/pkg/core/parse/builder"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

func Create(context *types.APIContext, version *types.APIVersion, typeName string, data map[string]interface{}, into interface{}) error {
	schema := context.Schemas.Schema(version, typeName)
	if schema == nil {
		return fmt.Errorf("failed to find schema " + typeName)
	}

	item, err := schema.Store.Create(context, schema, data)
	if err != nil {
		return err
	}

	b := builder.NewBuilder(context)
	b.Version = version

	item, err = b.Construct(schema, item, builder.List)
	if err != nil {
		return err
	}

	if into == nil {
		return nil
	}

	return convert.ToObj(item, into)
}

func ByID(context *types.APIContext, version *types.APIVersion, typeName string, id string, into interface{}) error {
	schema := context.Schemas.Schema(version, typeName)
	if schema == nil {
		return fmt.Errorf("failed to find schema " + typeName)
	}

	item, err := schema.Store.ByID(context, schema, id)
	if err != nil {
		return err
	}

	b := builder.NewBuilder(context)
	b.Version = version

	item, err = b.Construct(schema, item, builder.List)
	if err != nil {
		return err
	}

	if into == nil {
		return nil
	}

	return convert.ToObj(item, into)
}

func List(context *types.APIContext, version *types.APIVersion, typeName string, opts *types.QueryOptions, into interface{}) error {
	schema := context.Schemas.Schema(version, typeName)
	if schema == nil {
		return fmt.Errorf("failed to find schema " + typeName)
	}

	data, err := schema.Store.List(context, schema, opts)
	if err != nil {
		return err
	}

	b := builder.NewBuilder(context)
	b.Version = version

	var newData []map[string]interface{}
	for _, item := range data {
		item, err = b.Construct(schema, item, builder.List)
		if err != nil {
			return err
		}
		newData = append(newData, item)
	}

	return convert.ToObj(newData, into)
}
