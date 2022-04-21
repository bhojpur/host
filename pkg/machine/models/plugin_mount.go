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

// PluginMount plugin mount
//
// swagger:model PluginMount
type PluginMount struct {

	// description
	// Example: This is a mount that's used by the plugin.
	// Required: true
	Description string `json:"Description"`

	// destination
	// Example: /mnt/state
	// Required: true
	Destination string `json:"Destination"`

	// name
	// Example: some-mount
	// Required: true
	Name string `json:"Name"`

	// options
	// Example: ["rbind","rw"]
	// Required: true
	Options []string `json:"Options"`

	// settable
	// Required: true
	Settable []string `json:"Settable"`

	// source
	// Example: /var/lib/bhojpur/plugins/
	// Required: true
	Source *string `json:"Source"`

	// type
	// Example: bind
	// Required: true
	Type string `json:"Type"`
}

// Validate validates this plugin mount
func (m *PluginMount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDestination(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSettable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PluginMount) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("Description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateDestination(formats strfmt.Registry) error {

	if err := validate.RequiredString("Destination", "body", m.Destination); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateOptions(formats strfmt.Registry) error {

	if err := validate.Required("Options", "body", m.Options); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateSettable(formats strfmt.Registry) error {

	if err := validate.Required("Settable", "body", m.Settable); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("Source", "body", m.Source); err != nil {
		return err
	}

	return nil
}

func (m *PluginMount) validateType(formats strfmt.Registry) error {

	if err := validate.RequiredString("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this plugin mount based on context it is used
func (m *PluginMount) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PluginMount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PluginMount) UnmarshalBinary(b []byte) error {
	var res PluginMount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
