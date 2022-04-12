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
	"testing"

	"github.com/bhojpur/host/pkg/api/handler"
	"github.com/bhojpur/host/pkg/core/parse"
	"github.com/bhojpur/host/pkg/core/store/empty"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/stretchr/testify/assert"
)

type testStore struct {
	empty.Store
}

func (t *testStore) List(apiContext *types.APIContext, schema *types.Schema, opt *types.QueryOptions) ([]map[string]interface{}, error) {
	return []map[string]interface{}{{"1": "1"}, {"2": "2"}, {"3": "3"}}, nil
}

func TestWrap(t *testing.T) {
	store := &testStore{}
	limit := int64(1)
	opt := &types.QueryOptions{
		Pagination: &types.Pagination{
			Limit: &limit,
		},
	}
	apiContext := &types.APIContext{
		SubContextAttributeProvider: &parse.DefaultSubContextAttributeProvider{},
		QueryFilter:                 handler.QueryFilter,
		Pagination:                  opt.Pagination,
	}

	wrapped := Wrap(store)
	if _, err := wrapped.List(apiContext, &types.Schema{}, opt); err != nil {
		t.Fatal(err)
	}
	assert.True(t, apiContext.Pagination.Partial)
	assert.Equal(t, int64(3), *apiContext.Pagination.Total)
	assert.Equal(t, int64(1), *apiContext.Pagination.Limit)

	wrappedTwice := Wrap(wrapped)
	apiContext.Pagination = opt.Pagination
	if _, err := wrappedTwice.List(apiContext, &types.Schema{}, opt); err != nil {
		t.Fatal(err)
	}
	assert.True(t, apiContext.Pagination.Partial)
	assert.Equal(t, int64(3), *apiContext.Pagination.Total)
	assert.Equal(t, int64(1), *apiContext.Pagination.Limit)

}
