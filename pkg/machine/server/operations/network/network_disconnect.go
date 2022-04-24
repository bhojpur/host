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

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NetworkDisconnectHandlerFunc turns a function with the right signature into a network disconnect handler
type NetworkDisconnectHandlerFunc func(NetworkDisconnectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn NetworkDisconnectHandlerFunc) Handle(params NetworkDisconnectParams) middleware.Responder {
	return fn(params)
}

// NetworkDisconnectHandler interface for that can handle valid network disconnect params
type NetworkDisconnectHandler interface {
	Handle(NetworkDisconnectParams) middleware.Responder
}

// NewNetworkDisconnect creates a new http.Handler for the network disconnect operation
func NewNetworkDisconnect(ctx *middleware.Context, handler NetworkDisconnectHandler) *NetworkDisconnect {
	return &NetworkDisconnect{Context: ctx, Handler: handler}
}

/* NetworkDisconnect swagger:route POST /networks/{id}/disconnect Network networkDisconnect

Disconnect a container from a network

*/
type NetworkDisconnect struct {
	Context *middleware.Context
	Handler NetworkDisconnectHandler
}

func (o *NetworkDisconnect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewNetworkDisconnectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// NetworkDisconnectBody NetworkDisconnectRequest
//
// swagger:model NetworkDisconnectBody
type NetworkDisconnectBody struct {

	// The ID or name of the container to disconnect from the network.
	//
	Container string `json:"Container,omitempty"`

	// Force the container to disconnect from the network.
	//
	Force bool `json:"Force,omitempty"`
}

// Validate validates this network disconnect body
func (o *NetworkDisconnectBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this network disconnect body based on context it is used
func (o *NetworkDisconnectBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *NetworkDisconnectBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *NetworkDisconnectBody) UnmarshalBinary(b []byte) error {
	var res NetworkDisconnectBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}