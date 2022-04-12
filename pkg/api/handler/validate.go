package handler

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
	"github.com/bhojpur/host/pkg/core/parse"
	"github.com/bhojpur/host/pkg/core/parse/builder"
	"github.com/bhojpur/host/pkg/core/types"
)

func ParseAndValidateBody(apiContext *types.APIContext, create bool) (map[string]interface{}, error) {
	data, err := parse.Body(apiContext.Request)
	if err != nil {
		return nil, err
	}

	if create {
		for key, value := range apiContext.SubContextAttributeProvider.Create(apiContext, apiContext.Schema) {
			if data == nil {
				data = map[string]interface{}{}
			}
			data[key] = value
		}
	}

	b := builder.NewBuilder(apiContext)

	op := builder.Create
	if !create {
		op = builder.Update
	}
	if apiContext.Schema.InputFormatter != nil {
		err = apiContext.Schema.InputFormatter(apiContext, apiContext.Schema, data, create)
		if err != nil {
			return nil, err
		}
	}
	data, err = b.Construct(apiContext.Schema, data, op)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ParseAndValidateActionBody(apiContext *types.APIContext, actionInputSchema *types.Schema) (map[string]interface{}, error) {
	data, err := parse.Body(apiContext.Request)
	if err != nil {
		return nil, err
	}

	b := builder.NewBuilder(apiContext)

	op := builder.Create
	data, err = b.Construct(actionInputSchema, data, op)
	if err != nil {
		return nil, err
	}

	return data, nil
}
