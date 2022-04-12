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
	"strings"

	"github.com/bhojpur/host/pkg/common/name"
	"github.com/bhojpur/host/pkg/core/types"
)

type SelfLink struct {
	resource string
}

func (s *SelfLink) FromInternal(data map[string]interface{}) {
	if data != nil {
		sl, ok := data["selfLink"].(string)
		if !ok || sl == "" {
			data["selfLink"] = s.selflink(data)
		}
	}
}

func (s *SelfLink) ToInternal(data map[string]interface{}) error {
	return nil
}

func (s *SelfLink) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	s.resource = name.GuessPluralName(strings.ToLower(schema.ID))
	return nil
}

func (s *SelfLink) selflink(data map[string]interface{}) string {
	buf := &strings.Builder{}
	name, ok := data["name"].(string)
	if !ok || name == "" {
		return ""
	}
	apiVersion, ok := data["apiVersion"].(string)
	if !ok || apiVersion == "v1" {
		buf.WriteString("/api/v1/")
	} else {
		buf.WriteString("/apis/")
		buf.WriteString(apiVersion)
		buf.WriteString("/")
	}
	namespace, ok := data["namespace"].(string)
	if ok && namespace != "" {
		buf.WriteString("namespaces/")
		buf.WriteString(namespace)
		buf.WriteString("/")
	}
	buf.WriteString(s.resource)
	buf.WriteString("/")
	buf.WriteString(name)
	return buf.String()
}
