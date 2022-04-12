package data

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
	"github.com/bhojpur/host/pkg/common/data/convert"
)

type List []map[string]interface{}

type Object map[string]interface{}

func New() Object {
	return map[string]interface{}{}
}

func Convert(obj interface{}) (Object, error) {
	data, err := convert.EncodeToMap(obj)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (o Object) Map(names ...string) Object {
	v := GetValueN(o, names...)
	m := convert.ToMapInterface(v)
	return m
}

func (o Object) Slice(names ...string) (result []Object) {
	v := GetValueN(o, names...)
	for _, item := range convert.ToInterfaceSlice(v) {
		result = append(result, convert.ToMapInterface(item))
	}
	return
}

func (o Object) Values() (result []Object) {
	for k := range o {
		result = append(result, o.Map(k))
	}
	return
}

func (o Object) String(names ...string) string {
	v := GetValueN(o, names...)
	return convert.ToString(v)
}

func (o Object) StringSlice(names ...string) []string {
	v := GetValueN(o, names...)
	return convert.ToStringSlice(v)
}

func (o Object) Set(key string, obj interface{}) {
	if o == nil {
		return
	}
	o[key] = obj
}

func (o Object) SetNested(obj interface{}, key ...string) {
	PutValue(o, obj, key...)
}

func (o Object) Bool(key ...string) bool {
	return convert.ToBool(GetValueN(o, key...))
}
