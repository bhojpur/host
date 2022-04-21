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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// DistributionInspect DistributionInspectResponse
//
// Describes the result obtained from contacting the registry to retrieve
// image metadata.
//
//
// swagger:model DistributionInspect
type DistributionInspect struct {

	// descriptor
	// Required: true
	Descriptor *Descriptor `json:"Descriptor"`

	// An array containing all platforms supported by the image.
	//
	// Required: true
	Platforms []*Platform `json:"Platforms"`
}

// Validate validates this distribution inspect
func (m *DistributionInspect) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlatforms(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistributionInspect) validateDescriptor(formats strfmt.Registry) error {

	if err := validate.Required("Descriptor", "body", m.Descriptor); err != nil {
		return err
	}

	if m.Descriptor != nil {
		if err := m.Descriptor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Descriptor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Descriptor")
			}
			return err
		}
	}

	return nil
}

func (m *DistributionInspect) validatePlatforms(formats strfmt.Registry) error {

	if err := validate.Required("Platforms", "body", m.Platforms); err != nil {
		return err
	}

	for i := 0; i < len(m.Platforms); i++ {
		if swag.IsZero(m.Platforms[i]) { // not required
			continue
		}

		if m.Platforms[i] != nil {
			if err := m.Platforms[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Platforms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Platforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this distribution inspect based on the context it is used
func (m *DistributionInspect) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDescriptor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlatforms(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DistributionInspect) contextValidateDescriptor(ctx context.Context, formats strfmt.Registry) error {

	if m.Descriptor != nil {
		if err := m.Descriptor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Descriptor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("Descriptor")
			}
			return err
		}
	}

	return nil
}

func (m *DistributionInspect) contextValidatePlatforms(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Platforms); i++ {

		if m.Platforms[i] != nil {
			if err := m.Platforms[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Platforms" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("Platforms" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *DistributionInspect) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DistributionInspect) UnmarshalBinary(b []byte) error {
	var res DistributionInspect
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
