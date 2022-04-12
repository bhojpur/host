package mapper

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
	"strings"

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"github.com/bhojpur/host/pkg/core/types/values"
)

type CredentialMapper struct {
}

func (s CredentialMapper) FromInternal(data map[string]interface{}) {
	formatData(data)
	name := convert.ToString(values.GetValueN(data, "annotations", "field.bhojpur.net/name"))
	if name == "" {
		id := convert.ToString(values.GetValueN(data, "id"))
		if id != "" {
			values.PutValue(data, id, "annotations", "field.bhojpur.net/name")
		}
	}
	delete(data, "data")
}

func (s CredentialMapper) ToInternal(data map[string]interface{}) error {
	updateData(data)
	return nil
}

func (s CredentialMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

func updateData(data map[string]interface{}) {
	stringData := map[string]string{}
	for key, val := range data {
		if val == nil {
			continue
		}
		if strings.HasSuffix(key, "Config") {
			for key2, val2 := range convert.ToMapInterface(val) {
				stringData[fmt.Sprintf("%s-%s", key, key2)] = convert.ToString(val2)
			}
			values.PutValue(data, stringData, "stringData")
			delete(data, key)
			return
		}
	}
}

func formatData(data map[string]interface{}) {
	secretData := convert.ToMapInterface(data["data"])
	getKey := func(data map[string]interface{}) string {
		for key := range data {
			splitKeys := strings.Split(key, "-")
			if len(splitKeys) != 2 {
				continue
			}
			if strings.HasSuffix(splitKeys[0], "Config") {
				return splitKeys[0]
			}
		}
		return ""
	}
	config := getKey(secretData)
	if config == "" {
		return
	}
	for key, val := range secretData {
		splitKeys := strings.Split(key, "-")
		if len(splitKeys) != 2 {
			continue
		}
		values.PutValue(data, convert.ToString(val), config, splitKeys[1])
	}
}
