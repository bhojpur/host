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
	"github.com/bhojpur/host/pkg/core/types/definition"
	"github.com/bhojpur/host/pkg/core/types/values"
)

type Move struct {
	From, To, CodeName string
	DestDefined        bool
	NoDeleteFromField  bool
}

func (m Move) FromInternal(data map[string]interface{}) {
	if v, ok := values.RemoveValue(data, strings.Split(m.From, "/")...); ok {
		values.PutValue(data, v, strings.Split(m.To, "/")...)
	}
}

func (m Move) ToInternal(data map[string]interface{}) error {
	if v, ok := values.RemoveValue(data, strings.Split(m.To, "/")...); ok {
		values.PutValue(data, v, strings.Split(m.From, "/")...)
	}
	return nil
}

func (m Move) ModifySchema(s *types.Schema, schemas *types.Schemas) error {
	fromSchema, _, fromField, ok, err := getField(s, schemas, m.From)
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("failed to find field %s on schema %s", m.From, s.ID)
	}

	toSchema, toFieldName, _, ok, err := getField(s, schemas, m.To)
	if err != nil {
		return err
	}
	_, ok = toSchema.ResourceFields[toFieldName]
	if ok && !strings.Contains(m.To, "/") && !m.DestDefined {
		return fmt.Errorf("field %s already exists on schema %s", m.To, s.ID)
	}

	if !m.NoDeleteFromField {
		delete(fromSchema.ResourceFields, m.From)
	}

	if !m.DestDefined {
		if m.CodeName == "" {
			fromField.CodeName = convert.Capitalize(toFieldName)
		} else {
			fromField.CodeName = m.CodeName
		}
		toSchema.ResourceFields[toFieldName] = fromField
	}

	return nil
}

func getField(schema *types.Schema, schemas *types.Schemas, target string) (*types.Schema, string, types.Field, bool, error) {
	parts := strings.Split(target, "/")
	for i, part := range parts {
		if i == len(parts)-1 {
			continue
		}

		fieldType := schema.ResourceFields[part].Type
		if definition.IsArrayType(fieldType) {
			fieldType = definition.SubType(fieldType)
		}
		subSchema := schemas.Schema(&schema.Version, fieldType)
		if subSchema == nil {
			return nil, "", types.Field{}, false, fmt.Errorf("failed to find field or schema for %s on %s", part, schema.ID)
		}

		schema = subSchema
	}

	name := parts[len(parts)-1]
	f, ok := schema.ResourceFields[name]
	return schema, name, f, ok, nil
}
