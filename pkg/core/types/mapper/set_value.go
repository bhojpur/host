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
	"github.com/bhojpur/host/pkg/core/types/values"
)

type SetValue struct {
	Field, To        string
	Value            interface{}
	IfEq             interface{}
	IgnoreDefinition bool
}

func (s SetValue) FromInternal(data map[string]interface{}) {
	if s.IfEq == nil {
		values.PutValue(data, s.Value, strings.Split(s.getTo(), "/")...)
		return
	}

	v, ok := values.GetValue(data, strings.Split(s.Field, "/")...)
	if !ok {
		return
	}

	if v == s.IfEq {
		values.PutValue(data, s.Value, strings.Split(s.getTo(), "/")...)
	}
}

func (s SetValue) getTo() string {
	if s.To == "" {
		return s.Field
	}
	return s.To
}

func (s SetValue) ToInternal(data map[string]interface{}) error {
	v, ok := values.GetValue(data, strings.Split(s.getTo(), "/")...)
	if !ok {
		return nil
	}

	if s.IfEq == nil {
		values.RemoveValue(data, strings.Split(s.Field, "/")...)
	} else if v == s.Value {
		values.PutValue(data, s.IfEq, strings.Split(s.Field, "/")...)
	}

	return nil
}

func (s SetValue) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	if s.IgnoreDefinition {
		return nil
	}

	_, _, _, ok, err := getField(schema, schemas, s.getTo())
	if err != nil {
		return err
	}
	if !ok {
		return fmt.Errorf("failed to find defined field for %s on schemas %s", s.getTo(), schema.ID)
	}

	return nil
}
