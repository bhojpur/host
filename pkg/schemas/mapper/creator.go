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
	"github.com/bhojpur/host/pkg/core/types/mapper"
)

type Creator struct {
	m types.Mapper
}

func (c *Creator) FromInternal(data map[string]interface{}) {
	if c.m != nil {
		c.m.FromInternal(data)
	}
}

func (c *Creator) ToInternal(data map[string]interface{}) error {
	if c.m != nil {
		return c.m.ToInternal(data)
	}
	return nil
}

func (c *Creator) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	if schema.CanList(nil) == nil && schema.CanCreate(nil) == nil {
		schema.ResourceFields["creatorId"] = types.Field{
			Type:     "reference[/v3/schemas/user]",
			CodeName: "CreatorID",
		}
		c.m = &mapper.AnnotationField{Field: "creatorId"}
		return c.m.ModifySchema(schema, schemas)
	}
	return nil
}
