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

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

type UnionMapping struct {
	FieldName   string
	CheckFields []string
}

type UnionEmbed struct {
	Fields []UnionMapping
	embeds map[string]Embed
}

func (u *UnionEmbed) FromInternal(data map[string]interface{}) {
	for _, embed := range u.embeds {
		embed.FromInternal(data)
	}
}

func (u *UnionEmbed) ToInternal(data map[string]interface{}) error {
outer:
	for _, mapper := range u.Fields {
		if len(mapper.CheckFields) == 0 {
			continue
		}

		for _, check := range mapper.CheckFields {
			v, ok := data[check]
			if !ok || convert.IsAPIObjectEmpty(v) {
				continue outer
			}
		}

		embed := u.embeds[mapper.FieldName]
		return embed.ToInternal(data)
	}

	return nil
}

func (u *UnionEmbed) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	u.embeds = map[string]Embed{}

	for _, mapping := range u.Fields {
		embed := Embed{
			Field:          mapping.FieldName,
			ignoreOverride: true,
		}
		if err := embed.ModifySchema(schema, schemas); err != nil {
			return err
		}

		for _, checkField := range mapping.CheckFields {
			if _, ok := schema.ResourceFields[checkField]; !ok {
				return fmt.Errorf("missing check field %s on schema %s", checkField, schema.ID)
			}
		}

		u.embeds[mapping.FieldName] = embed
	}

	return nil
}
