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
	"github.com/bhojpur/host/pkg/core/types/values"
)

const (
	extIPField = "externalIpAddress"
)

type NodeAddressMapper struct {
}

func (n NodeAddressMapper) FromInternal(data map[string]interface{}) {
	addresses, _ := values.GetSlice(data, "addresses")
	for _, address := range addresses {
		t := address["type"]
		a := address["address"]
		if t == "InternalIP" {
			data["ipAddress"] = a
		} else if t == "ExternalIP" {
			data[extIPField] = a
		} else if t == "Hostname" {
			data["hostname"] = a
		}
	}
}

func (n NodeAddressMapper) ToInternal(data map[string]interface{}) error {
	return nil
}

func (n NodeAddressMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}

type NodeAddressAnnotationMapper struct {
}

func (n NodeAddressAnnotationMapper) FromInternal(data map[string]interface{}) {
	externalIP, ok := values.GetValue(data, "status", "nodeAnnotations", "bke.bhojpur.net/external-ip")
	if ok {
		data[extIPField] = externalIP
	}
}

func (n NodeAddressAnnotationMapper) ToInternal(data map[string]interface{}) error {
	return nil
}

func (n NodeAddressAnnotationMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
