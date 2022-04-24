// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Commit Commit holds the Git-commit (SHA1) that a binary was built from, as
// reported in the version-string of external tools, such as `containerd`,
// or `runC`.
//
//
// swagger:model Commit
type Commit struct {

	// Commit ID of external tool expected by hostutld as set at build time.
	//
	// Example: 2d41c047c83e09a6d61d464906feb2a2f3c52aa4
	Expected string `json:"Expected,omitempty"`

	// Actual commit ID of external tool.
	// Example: cfb82a876ecc11b5ca0977d1733adbe58599088a
	ID string `json:"ID,omitempty"`
}

// Validate validates this commit
func (m *Commit) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this commit based on context it is used
func (m *Commit) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Commit) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Commit) UnmarshalBinary(b []byte) error {
	var res Commit
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}