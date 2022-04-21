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

// NetworkingConfig NetworkingConfig represents the container's networking configuration for
// each of its interfaces.
// It is used for the networking configs specified in the `hostutl create`
// and `hostutl network connect` commands.
//
// Example: {"EndpointsConfig":{"isolated_nw":{"Aliases":["server_x","server_y"],"IPAMConfig":{"IPv4Address":"172.20.30.33","IPv6Address":"2001:db8:abcd::3033","LinkLocalIPs":["169.254.34.68","fe80::3468"]},"Links":["container_1","container_2"]}}}
//
// swagger:model NetworkingConfig
type NetworkingConfig struct {

	// A mapping of network name to endpoint configuration for that network.
	//
	EndpointsConfig map[string]EndpointSettings `json:"EndpointsConfig,omitempty"`
}

// Validate validates this networking config
func (m *NetworkingConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndpointsConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkingConfig) validateEndpointsConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.EndpointsConfig) { // not required
		return nil
	}

	for k := range m.EndpointsConfig {

		if err := validate.Required("EndpointsConfig"+"."+k, "body", m.EndpointsConfig[k]); err != nil {
			return err
		}
		if val, ok := m.EndpointsConfig[k]; ok {
			if err := val.Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("EndpointsConfig" + "." + k)
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("EndpointsConfig" + "." + k)
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this networking config based on the context it is used
func (m *NetworkingConfig) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateEndpointsConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NetworkingConfig) contextValidateEndpointsConfig(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.EndpointsConfig {

		if val, ok := m.EndpointsConfig[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *NetworkingConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NetworkingConfig) UnmarshalBinary(b []byte) error {
	var res NetworkingConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
