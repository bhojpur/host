package types

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
	"net/http"

	"github.com/bhojpur/host/pkg/core/httperror"
	"github.com/bhojpur/host/pkg/core/types/slice"
)

func (s *Schema) MustCustomizeField(name string, f func(f Field) Field) *Schema {
	field, ok := s.ResourceFields[name]
	if !ok {
		panic("Failed to find field " + name + " on schema " + s.ID)
	}
	s.ResourceFields[name] = f(field)
	return s
}

func (v *APIVersion) Equals(other *APIVersion) bool {
	return v.Version == other.Version &&
		v.Group == other.Group &&
		v.Path == other.Path
}

func (s *Schema) CanList(context *APIContext) error {
	if context == nil {
		if slice.ContainsString(s.CollectionMethods, http.MethodGet) {
			return nil
		}
		return httperror.NewAPIError(httperror.PermissionDenied, "can not list "+s.ID)
	}
	return context.AccessControl.CanList(context, s)
}

func (s *Schema) CanGet(context *APIContext) error {
	if context == nil {
		if slice.ContainsString(s.ResourceMethods, http.MethodGet) {
			return nil
		}
		return httperror.NewAPIError(httperror.PermissionDenied, "can not get "+s.ID)
	}
	return context.AccessControl.CanGet(context, s)
}

func (s *Schema) CanCreate(context *APIContext) error {
	if context == nil {
		if slice.ContainsString(s.CollectionMethods, http.MethodPost) {
			return nil
		}
		return httperror.NewAPIError(httperror.PermissionDenied, "can not create "+s.ID)
	}
	return context.AccessControl.CanCreate(context, s)
}

func (s *Schema) CanUpdate(context *APIContext) error {
	if context == nil {
		if slice.ContainsString(s.ResourceMethods, http.MethodPut) {
			return nil
		}
		return httperror.NewAPIError(httperror.PermissionDenied, "can not update "+s.ID)
	}
	return context.AccessControl.CanUpdate(context, nil, s)
}

func (s *Schema) CanDelete(context *APIContext) error {
	if context == nil {
		if slice.ContainsString(s.ResourceMethods, http.MethodDelete) {
			return nil
		}
		return httperror.NewAPIError(httperror.PermissionDenied, "can not delete "+s.ID)
	}
	return context.AccessControl.CanDelete(context, nil, s)
}
