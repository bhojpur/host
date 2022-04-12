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
)

type Scope struct {
	If      types.TypeScope
	IfNot   types.TypeScope
	Mappers []types.Mapper
	run     bool
}

func (s *Scope) FromInternal(data map[string]interface{}) {
	if s.run {
		types.Mappers(s.Mappers).FromInternal(data)
	}
}

func (s *Scope) ToInternal(data map[string]interface{}) error {
	if s.run {
		return types.Mappers(s.Mappers).ToInternal(data)
	}
	return nil
}

func (s *Scope) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	if s.If != "" {
		s.run = schema.Scope == s.If
	}
	if s.IfNot != "" {
		s.run = schema.Scope != s.IfNot
	}
	if s.run {
		return types.Mappers(s.Mappers).ModifySchema(schema, schemas)
	}
	return nil
}
