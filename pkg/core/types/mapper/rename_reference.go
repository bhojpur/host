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

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"github.com/bhojpur/host/pkg/core/types/definition"
)

type RenameReference struct {
	mapper types.Mapper
}

func (r *RenameReference) FromInternal(data map[string]interface{}) {
	if r.mapper != nil {
		r.mapper.FromInternal(data)
	}
}

func (r *RenameReference) ToInternal(data map[string]interface{}) error {
	if r.mapper != nil {
		return r.mapper.ToInternal(data)
	}
	return nil
}

func (r *RenameReference) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	var mappers []types.Mapper
	for name, field := range schema.ResourceFields {
		if definition.IsReferenceType(field.Type) && strings.HasSuffix(name, "Name") {
			newName := strings.TrimSuffix(name, "Name") + "Id"
			newCodeName := convert.Capitalize(strings.TrimSuffix(name, "Name") + "ID")
			move := Move{From: name, To: newName, CodeName: newCodeName}
			if err := move.ModifySchema(schema, schemas); err != nil {
				return err
			}

			mappers = append(mappers, move)
		} else if definition.IsArrayType(field.Type) && definition.IsReferenceType(definition.SubType(field.Type)) && strings.HasSuffix(name, "Names") {
			newName := strings.TrimSuffix(name, "Names") + "Ids"
			newCodeName := convert.Capitalize(strings.TrimSuffix(name, "Names") + "IDs")
			move := Move{From: name, To: newName, CodeName: newCodeName}
			if err := move.ModifySchema(schema, schemas); err != nil {
				return err
			}

			mappers = append(mappers, move)
		}
	}

	if len(mappers) > 0 {
		r.mapper = types.Mappers(mappers)
	}

	return nil
}
