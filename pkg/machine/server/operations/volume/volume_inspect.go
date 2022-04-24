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

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// VolumeInspectHandlerFunc turns a function with the right signature into a volume inspect handler
type VolumeInspectHandlerFunc func(VolumeInspectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VolumeInspectHandlerFunc) Handle(params VolumeInspectParams) middleware.Responder {
	return fn(params)
}

// VolumeInspectHandler interface for that can handle valid volume inspect params
type VolumeInspectHandler interface {
	Handle(VolumeInspectParams) middleware.Responder
}

// NewVolumeInspect creates a new http.Handler for the volume inspect operation
func NewVolumeInspect(ctx *middleware.Context, handler VolumeInspectHandler) *VolumeInspect {
	return &VolumeInspect{Context: ctx, Handler: handler}
}

/* VolumeInspect swagger:route GET /volumes/{name} Volume volumeInspect

Inspect a volume

*/
type VolumeInspect struct {
	Context *middleware.Context
	Handler VolumeInspectHandler
}

func (o *VolumeInspect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewVolumeInspectParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}