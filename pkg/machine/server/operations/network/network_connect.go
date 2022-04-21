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

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// NetworkConnectHandlerFunc turns a function with the right signature into a network connect handler
type NetworkConnectHandlerFunc func(NetworkConnectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NetworkConnectHandlerFunc) Handle(params NetworkConnectParams) middleware.Responder {
	return fn(params)
}

// NetworkConnectHandler interface for that can handle valid network connect params
type NetworkConnectHandler interface {
	Handle(NetworkConnectParams) middleware.Responder
}

// NewNetworkConnect creates a new http.Handler for the network connect operation
func NewNetworkConnect(ctx *middleware.Context, handler NetworkConnectHandler) *NetworkConnect {
	return &NetworkConnect{Context: ctx, Handler: handler}
}

/* NetworkConnect swagger:route POST /networks/{id}/connect Network networkConnect

Connect a container to a network

*/
type NetworkConnect struct {
	Context *middleware.Context
	Handler NetworkConnectHandler
}

func (o *NetworkConnect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewNetworkConnectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NetworkConnectBody NetworkConnectRequest
// Example: {"Container":"3613f73ba0e4","EndpointConfig":{"IPAMConfig":{"IPv4Address":"172.24.56.89","IPv6Address":"2001:db8::5689"}}}
//
// swagger:model NetworkConnectBody
type NetworkConnectBody struct {

	// The ID or name of the container to connect to the network.
	Container string `json:"Container,omitempty"`

	// endpoint config
	EndpointConfig *models.EndpointSettings `json:"EndpointConfig,omitempty"`
}

// Validate validates this network connect body
func (o *NetworkConnectBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEndpointConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkConnectBody) validateEndpointConfig(formats strfmt.Registry) error {
	if swag.IsZero(o.EndpointConfig) { // not required
		return nil
	}

	if o.EndpointConfig != nil {
		if err := o.EndpointConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container" + "." + "EndpointConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container" + "." + "EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this network connect body based on the context it is used
func (o *NetworkConnectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEndpointConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *NetworkConnectBody) contextValidateEndpointConfig(ctx context.Context, formats strfmt.Registry) error {

	if o.EndpointConfig != nil {
		if err := o.EndpointConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("container" + "." + "EndpointConfig")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("container" + "." + "EndpointConfig")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *NetworkConnectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkConnectBody) UnmarshalBinary(b []byte) error {
	var res NetworkConnectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
