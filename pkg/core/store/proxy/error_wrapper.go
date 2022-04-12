package proxy

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
	"k8s.io/apimachinery/pkg/api/errors"
)

type errorStore struct {
	types.Store
}

func (e *errorStore) ByID(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	data, err := e.Store.ByID(apiContext, schema, id)
	return data, translateError(err)
}

func (e *errorStore) List(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) ([]map[string]interface{}, error) {
	data, err := e.Store.List(apiContext, schema, opt)
	return data, translateError(err)
}

func (e *errorStore) Create(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}) (map[string]interface{}, error) {
	data, err := e.Store.Create(apiContext, schema, data)
	return data, translateError(err)

}

func (e *errorStore) Update(apiContext *types.APIContext, schema *types.Schema, data map[string]interface{}, id string) (map[string]interface{}, error) {
	data, err := e.Store.Update(apiContext, schema, data, id)
	return data, translateError(err)

}

func (e *errorStore) Delete(apiContext *types.APIContext, schema *types.Schema, id string) (map[string]interface{}, error) {
	data, err := e.Store.Delete(apiContext, schema, id)
	return data, translateError(err)

}

func (e *errorStore) Watch(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) (chan map[string]interface{}, error) {
	data, err := e.Store.Watch(apiContext, schema, opt)
	return data, translateError(err)
}

func translateError(err error) error {
	if apiError, ok := err.(errors.APIStatus); ok {
		status := apiError.Status()
		return httperror.NewAPIErrorLong(int(status.Code), string(status.Reason), status.Message)
	}
	return err
}
