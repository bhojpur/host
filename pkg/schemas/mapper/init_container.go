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
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

type InitContainerMapper struct {
}

func (e InitContainerMapper) FromInternal(data map[string]interface{}) {
	containers, _ := data["containers"].([]interface{})

	for _, initContainer := range convert.ToMapSlice(data["initContainers"]) {
		if initContainer == nil {
			continue
		}
		initContainer["initContainer"] = true
		containers = append(containers, initContainer)
	}

	if data != nil {
		data["containers"] = containers
	}
}

func (e InitContainerMapper) ToInternal(data map[string]interface{}) error {
	var newContainers []interface{}
	var newInitContainers []interface{}

	for _, container := range convert.ToMapSlice(data["containers"]) {
		if convert.ToBool(container["initContainer"]) {
			newInitContainers = append(newInitContainers, container)
		} else {
			newContainers = append(newContainers, container)
		}
		delete(container, "initContainer")
	}

	if data != nil {
		data["containers"] = newContainers
		data["initContainers"] = newInitContainers
	}

	return nil
}

func (e InitContainerMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	delete(schema.ResourceFields, "initContainers")
	return nil
}
