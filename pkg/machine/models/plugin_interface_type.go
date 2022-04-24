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

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PluginInterfaceType plugin interface type
//
// swagger:model PluginInterfaceType
type PluginInterfaceType struct {

	// capability
	// Required: true
	Capability string `json:"Capability"`

	// prefix
	// Required: true
	Prefix string `json:"Prefix"`

	// version
	// Required: true
	Version string `json:"Version"`
}

// Validate validates this plugin interface type
func (m *PluginInterfaceType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCapability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrefix(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginInterfaceType) validateCapability(formats strfmt.Registry) error {

	if err := validate.RequiredString("Capability", "body", m.Capability); err != nil {
		return err
	}

	return nil
}

func (m *PluginInterfaceType) validatePrefix(formats strfmt.Registry) error {

	if err := validate.RequiredString("Prefix", "body", m.Prefix); err != nil {
		return err
	}

	return nil
}

func (m *PluginInterfaceType) validateVersion(formats strfmt.Registry) error {

	if err := validate.RequiredString("Version", "body", m.Version); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin interface type based on context it is used
func (m *PluginInterfaceType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginInterfaceType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginInterfaceType) UnmarshalBinary(b []byte) error {
	var res PluginInterfaceType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}