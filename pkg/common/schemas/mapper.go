package schemas

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
	"github.com/bhojpur/host/pkg/common/data"
	"github.com/bhojpur/host/pkg/common/data/convert"
	"github.com/bhojpur/host/pkg/common/merr"
	"github.com/bhojpur/host/pkg/common/schemas/definition"
)

type Mapper interface {
	FromInternal(data data.Object)
	ToInternal(data data.Object) error
	ModifySchema(schema *Schema, schemas *Schemas) error
}

type Mappers []Mapper

func (m Mappers) FromInternal(data data.Object) {
	for _, mapper := range m {
		mapper.FromInternal(data)
	}
}

func (m Mappers) ToInternal(data data.Object) error {
	var errors []error
	for i := len(m) - 1; i >= 0; i-- {
		errors = append(errors, m[i].ToInternal(data))
	}
	return merr.NewErrors(errors...)
}

func (m Mappers) ModifySchema(schema *Schema, schemas *Schemas) error {
	for _, mapper := range m {
		if err := mapper.ModifySchema(schema, schemas); err != nil {
			return err
		}
	}
	return nil
}

type typeMapper struct {
	Mappers         []Mapper
	root            bool
	typeName        string
	subSchemas      map[string]*Schema
	subArraySchemas map[string]*Schema
	subMapSchemas   map[string]*Schema
}

func (t *typeMapper) FromInternal(data data.Object) {
	for fieldName, schema := range t.subSchemas {
		if schema.Mapper == nil {
			continue
		}
		schema.Mapper.FromInternal(data.Map(fieldName))
	}

	for fieldName, schema := range t.subMapSchemas {
		if schema.Mapper == nil {
			continue
		}
		for _, fieldData := range data.Map(fieldName).Values() {
			schema.Mapper.FromInternal(fieldData)
		}
	}

	for fieldName, schema := range t.subArraySchemas {
		if schema.Mapper == nil {
			continue
		}
		for _, fieldData := range data.Slice(fieldName) {
			schema.Mapper.FromInternal(fieldData)
		}
	}

	Mappers(t.Mappers).FromInternal(data)
}

func addError(errors []error, err error) []error {
	if err == nil {
		return errors
	}
	return append(errors, err)
}

func (t *typeMapper) ToInternal(data data.Object) error {
	var errors []error
	errors = addError(errors, Mappers(t.Mappers).ToInternal(data))

	for fieldName, schema := range t.subArraySchemas {
		if schema.Mapper == nil {
			continue
		}
		for _, fieldData := range data.Slice(fieldName) {
			errors = addError(errors, schema.Mapper.ToInternal(fieldData))
		}
	}

	for fieldName, schema := range t.subMapSchemas {
		if schema.Mapper == nil {
			continue
		}
		for _, fieldData := range data.Map(fieldName) {
			errors = addError(errors, schema.Mapper.ToInternal(convert.ToMapInterface(fieldData)))
		}
	}

	for fieldName, schema := range t.subSchemas {
		if schema.Mapper == nil {
			continue
		}
		errors = addError(errors, schema.Mapper.ToInternal(data.Map(fieldName)))
	}

	return merr.NewErrors(errors...)
}

func (t *typeMapper) ModifySchema(schema *Schema, schemas *Schemas) error {
	t.subSchemas = map[string]*Schema{}
	t.subArraySchemas = map[string]*Schema{}
	t.subMapSchemas = map[string]*Schema{}
	t.typeName = schema.ID

	mapperSchema := schema
	if schema.InternalSchema != nil {
		mapperSchema = schema.InternalSchema
	}
	for name, field := range mapperSchema.ResourceFields {
		fieldType := field.Type
		targetMap := t.subSchemas
		if definition.IsArrayType(fieldType) {
			fieldType = definition.SubType(fieldType)
			targetMap = t.subArraySchemas
		} else if definition.IsMapType(fieldType) {
			fieldType = definition.SubType(fieldType)
			targetMap = t.subMapSchemas
		}

		schema := schemas.doSchema(fieldType, false)
		if schema != nil {
			targetMap[name] = schema
		}
	}

	return Mappers(t.Mappers).ModifySchema(schema, schemas)
}
