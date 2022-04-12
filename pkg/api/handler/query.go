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
	"sort"

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

func QueryFilter(opts *types.QueryOptions, schema *types.Schema, data []map[string]interface{}) []map[string]interface{} {
	return ApplyQueryOptions(opts, schema, data)
}

func ApplyQueryOptions(options *types.QueryOptions, schema *types.Schema, data []map[string]interface{}) []map[string]interface{} {
	data = ApplyQueryConditions(options.Conditions, schema, data)
	data = ApplySort(options.Sort, data)
	return ApplyPagination(options.Pagination, data)
}

func ApplySort(sortOpts types.Sort, data []map[string]interface{}) []map[string]interface{} {
	name := sortOpts.Name
	if name == "" {
		name = "id"
	}

	sort.Slice(data, func(i, j int) bool {
		left, right := i, j
		if sortOpts.Order == types.DESC {
			left, right = j, i
		}

		return convert.ToString(data[left][name]) < convert.ToString(data[right][name])
	})

	return data
}

func ApplyQueryConditions(conditions []*types.QueryCondition, schema *types.Schema, data []map[string]interface{}) []map[string]interface{} {
	var result []map[string]interface{}

outer:
	for _, item := range data {
		for _, condition := range conditions {
			if !condition.Valid(schema, item) {
				continue outer
			}
		}

		result = append(result, item)
	}

	return result
}

func ApplyPagination(pagination *types.Pagination, data []map[string]interface{}) []map[string]interface{} {
	if pagination == nil || pagination.Limit == nil {
		return data
	}

	limit := *pagination.Limit
	if limit < 0 {
		limit = 0
	}

	total := int64(len(data))

	// Reset fields
	pagination.Next = ""
	pagination.Previous = ""
	pagination.Partial = false
	pagination.Total = &total
	pagination.First = ""

	if len(data) == 0 {
		return data
	}

	// startIndex is guaranteed to be a valid index
	startIndex := int64(0)
	if pagination.Marker != "" {
		for i, item := range data {
			id, _ := item["id"].(string)
			if id == pagination.Marker {
				startIndex = int64(i)
				break
			}
		}
	}

	previousIndex := startIndex - limit
	if previousIndex <= 0 {
		previousIndex = 0
	}
	nextIndex := startIndex + limit
	if nextIndex > int64(len(data)) {
		nextIndex = int64(len(data))
	}

	if previousIndex < startIndex {
		pagination.Previous, _ = data[previousIndex]["id"].(string)
	}

	if nextIndex > startIndex && nextIndex < int64(len(data)) {
		pagination.Next, _ = data[nextIndex]["id"].(string)
	}

	if startIndex > 0 || nextIndex < int64(len(data)) {
		pagination.Partial = true
	}

	if pagination.Partial {
		pagination.First, _ = data[0]["id"].(string)

		lastIndex := int64(len(data)) - limit
		if lastIndex > 0 && lastIndex < int64(len(data)) {
			pagination.Last, _ = data[lastIndex]["id"].(string)
		}
	}

	return data[startIndex:nextIndex]
}
