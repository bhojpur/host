package transform

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

	"github.com/bhojpur/host/pkg/core/httperror"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

type TransformerFunc func(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}, opt *types.QueryOptions) (map[string]interface{}, error)

type ListTransformerFunc func(apiContext *types.APIContext, schema *types.Schema, data []map[string]interface{}, opt *types.QueryOptions) ([]map[string]interface{}, error)

type StreamTransformerFunc func(apiContext *types.APIContext, schema *types.Schema, data chan map[string]interface{}, opt *types.QueryOptions) (chan map[string]interface{}, error)

type Store struct {
	Store             types.Store
	Transformer       TransformerFunc
	ListTransformer   ListTransformerFunc
	StreamTransformer StreamTransformerFunc
}

func (s *Store) Context() types.StorageContext {
	return s.Store.Context()
}

func (s *Store) ByID(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	data, err := s.Store.ByID(apiContext, schema, id)
	if err != nil {
		return nil, err
	}
	if s.Transformer == nil {
		return data, nil
	}
	obj, err := s.Transformer(apiContext, schema, data, &types.QueryOptions{
		Options: map[string]string{
			"ByID": "true",
		},
	})
	if obj == nil && err == nil {
		return obj, httperror.NewAPIError(httperror.NotFound, fmt.Sprintf("%s not found", id))
	}
	return obj, err
}

func (s *Store) Watch(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) (chan map[string]interface{}, error) {
	c, err := s.Store.Watch(apiContext, schema, opt)
	if err != nil {
		return nil, err
	}

	if s.StreamTransformer != nil {
		return s.StreamTransformer(apiContext, schema, c, opt)
	}

	return convert.Chan(c, func(data map[string]interface{}) map[string]interface{} {
		item, err := s.Transformer(apiContext, schema, data, opt)
		if err != nil {
			return nil
		}
		return item
	}), nil
}

func (s *Store) List(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) ([]map[string]interface{}, error) {
	data, err := s.Store.List(apiContext, schema, opt)
	if err != nil {
		return nil, err
	}

	if s.ListTransformer != nil {
		return s.ListTransformer(apiContext, schema, data, opt)
	}

	if s.Transformer == nil {
		return data, nil
	}

	var result []map[string]interface{}
	for _, item := range data {
		item, err := s.Transformer(apiContext, schema, item, opt)
		if err != nil {
			return nil, err
		}
		if item != nil {
			result = append(result, item)
		}
	}

	return result, nil
}

func (s *Store) Create(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}) (map[string]interface{}, error) {
	data, err := s.Store.Create(apiContext, schema, data)
	if err != nil {
		return nil, err
	}
	if s.Transformer == nil {
		return data, nil
	}
	return s.Transformer(apiContext, schema, data, nil)
}

func (s *Store) Update(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}, id string) (map[string]interface{}, error) {
	data, err := s.Store.Update(apiContext, schema, data, id)
	if err != nil {
		return nil, err
	}
	if s.Transformer == nil {
		return data, nil
	}
	return s.Transformer(apiContext, schema, data, nil)
}

func (s *Store) Delete(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	obj, err := s.Store.Delete(apiContext, schema, id)
	if err != nil || obj == nil {
		return obj, err
	}
	return s.Transformer(apiContext, schema, obj, nil)
}
