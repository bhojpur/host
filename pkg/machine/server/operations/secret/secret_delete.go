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

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// SecretDeleteHandlerFunc turns a function with the right signature into a secret delete handler
type SecretDeleteHandlerFunc func(SecretDeleteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn SecretDeleteHandlerFunc) Handle(params SecretDeleteParams) middleware.Responder {
	return fn(params)
}

// SecretDeleteHandler interface for that can handle valid secret delete params
type SecretDeleteHandler interface {
	Handle(SecretDeleteParams) middleware.Responder
}

// NewSecretDelete creates a new http.Handler for the secret delete operation
func NewSecretDelete(ctx *middleware.Context, handler SecretDeleteHandler) *SecretDelete {
	return &SecretDelete{Context: ctx, Handler: handler}
}

/* SecretDelete swagger:route DELETE /secrets/{id} Secret secretDelete

Delete a secret

*/
type SecretDelete struct {
	Context *middleware.Context
	Handler SecretDeleteHandler
}

func (o *SecretDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewSecretDeleteParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}