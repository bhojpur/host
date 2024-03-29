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

package service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// ServiceLogsHandlerFunc turns a function with the right signature into a service logs handler
type ServiceLogsHandlerFunc func(ServiceLogsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceLogsHandlerFunc) Handle(params ServiceLogsParams) middleware.Responder {
	return fn(params)
}

// ServiceLogsHandler interface for that can handle valid service logs params
type ServiceLogsHandler interface {
	Handle(ServiceLogsParams) middleware.Responder
}

// NewServiceLogs creates a new http.Handler for the service logs operation
func NewServiceLogs(ctx *middleware.Context, handler ServiceLogsHandler) *ServiceLogs {
	return &ServiceLogs{Context: ctx, Handler: handler}
}

/* ServiceLogs swagger:route GET /services/{id}/logs Service serviceLogs

Get service logs

Get `stdout` and `stderr` logs from a service. See also
[`/containers/{id}/logs`](#operation/ContainerLogs).

**Note**: This endpoint works only for services with the `local`,
`json-file` or `journald` logging drivers.


*/
type ServiceLogs struct {
	Context *middleware.Context
	Handler ServiceLogsHandler
}

func (o *ServiceLogs) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceLogsParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
