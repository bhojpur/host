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
	"net/http"
	"strconv"
	"strings"

	"github.com/bhojpur/host/pkg/core/types"
)

var (
	defaultLimit = int64(1000)
	maxLimit     = int64(10000)
)

func QueryOptions(apiContext *types.APIContext, schema *types.Schema) types.QueryOptions {
	req := apiContext.Request
	if req.Method != http.MethodGet {
		return types.QueryOptions{}
	}

	result := &types.QueryOptions{}

	result.Sort = parseSort(schema, apiContext)
	result.Pagination = parsePagination(apiContext)
	result.Conditions = parseFilters(schema, apiContext)

	return *result
}

func parseOrder(apiContext *types.APIContext) types.SortOrder {
	order := apiContext.Query.Get("order")
	if types.SortOrder(order) == types.DESC {
		return types.DESC
	}
	return types.ASC
}

func parseSort(schema *types.Schema, apiContext *types.APIContext) types.Sort {
	sortField := apiContext.Query.Get("sort")
	if _, ok := schema.CollectionFilters[sortField]; !ok {
		sortField = ""
	}
	return types.Sort{
		Order: parseOrder(apiContext),
		Name:  sortField,
	}
}

func parsePagination(apiContext *types.APIContext) *types.Pagination {
	if apiContext.Pagination != nil {
		return apiContext.Pagination
	}

	q := apiContext.Query
	limit := q.Get("limit")
	marker := q.Get("marker")

	result := &types.Pagination{
		Limit:  &defaultLimit,
		Marker: marker,
	}

	if limit != "" {
		limitInt, err := strconv.ParseInt(limit, 10, 64)
		if err != nil {
			return result
		}

		if limitInt > maxLimit || limitInt == -1 {
			result.Limit = &maxLimit
		} else if limitInt >= 0 {
			result.Limit = &limitInt
		}
	}

	return result
}

func parseNameAndOp(value string) (string, types.ModifierType) {
	name := value
	op := "eq"

	idx := strings.LastIndex(value, "_")
	if idx > 0 {
		op = value[idx+1:]
		name = value[0:idx]
	}

	return name, types.ModifierType(op)
}

func parseFilters(schema *types.Schema, apiContext *types.APIContext) []*types.QueryCondition {
	var conditions []*types.QueryCondition
	for key, values := range apiContext.Query {
		name, op := parseNameAndOp(key)
		filter, ok := schema.CollectionFilters[name]
		if !ok {
			continue
		}

		for _, mod := range filter.Modifiers {
			if op != mod || !types.ValidMod(op) {
				continue
			}

			conditions = append(conditions, types.NewConditionFromString(name, mod, values...))
		}
	}

	return conditions
}
