package wrapper

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
	"github.com/bhojpur/host/pkg/core/httperror"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

func Wrap(store types.Store) types.Store {
	if _, ok := store.(*StoreWrapper); ok {
		return store
	}

	return &StoreWrapper{
		store: store,
	}
}

type StoreWrapper struct {
	store types.Store
}

func (s *StoreWrapper) Context() types.StorageContext {
	return s.store.Context()
}

func (s *StoreWrapper) ByID(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	data, err := s.store.ByID(apiContext, schema, id)
	if err != nil {
		return nil, err
	}

	return apiContext.FilterObject(&types.QueryOptions{
		Conditions: apiContext.SubContextAttributeProvider.Query(apiContext, schema),
	}, schema, data), nil
}

func (s *StoreWrapper) List(apiContext *types.APIContext, schema *types.Schema, opts *types.QueryOptions) ([]map[string]interface{}, error) {
	opts.Conditions = append(opts.Conditions, apiContext.SubContextAttributeProvider.Query(apiContext, schema)...)
	data, err := s.store.List(apiContext, schema, opts)
	if err != nil {
		return nil, err
	}

	return apiContext.FilterList(opts, schema, data), nil
}

func (s *StoreWrapper) Watch(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) (chan map[string]interface{}, error) {
	c, err := s.store.Watch(apiContext, schema, opt)
	if err != nil || c == nil {
		return nil, err
	}

	return convert.Chan(c, func(data map[string]interface{}) map[string]interface{} {
		return apiContext.FilterObject(&types.QueryOptions{
			Conditions: apiContext.SubContextAttributeProvider.Query(apiContext, schema),
		}, schema, data)
	}), nil
}

func (s *StoreWrapper) Create(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}) (map[string]interface{}, error) {
	for key, value := range apiContext.SubContextAttributeProvider.Create(apiContext, schema) {
		if data == nil {
			data = map[string]interface{}{}
		}
		data[key] = value
	}

	data, err := s.store.Create(apiContext, schema, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *StoreWrapper) Update(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}, id string) (map[string]interface{}, error) {
	err := validateGet(apiContext, schema, id)
	if err != nil {
		return nil, err
	}

	data, err = s.store.Update(apiContext, schema, data, id)
	if err != nil {
		return nil, err
	}

	return apiContext.FilterObject(&types.QueryOptions{
		Conditions: apiContext.SubContextAttributeProvider.Query(apiContext, schema),
	}, schema, data), nil
}

func (s *StoreWrapper) Delete(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	if err := validateGet(apiContext, schema, id); err != nil {
		return nil, err
	}

	return s.store.Delete(apiContext, schema, id)
}

func validateGet(apiContext *types.APIContext, schema *types.Schema, id string) error {
	store := schema.Store
	if store == nil {
		return nil
	}

	existing, err := store.ByID(apiContext, schema, id)
	if err != nil {
		return err
	}

	if apiContext.Filter(&types.QueryOptions{
		Conditions: apiContext.SubContextAttributeProvider.Query(apiContext, schema),
	}, schema, existing) == nil {
		return httperror.NewAPIError(httperror.NotFound, "failed to find "+id)
	}

	return nil
}
