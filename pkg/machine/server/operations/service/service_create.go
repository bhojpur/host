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
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ServiceCreateHandlerFunc turns a function with the right signature into a service create handler
type ServiceCreateHandlerFunc func(ServiceCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ServiceCreateHandlerFunc) Handle(params ServiceCreateParams) middleware.Responder {
	return fn(params)
}

// ServiceCreateHandler interface for that can handle valid service create params
type ServiceCreateHandler interface {
	Handle(ServiceCreateParams) middleware.Responder
}

// NewServiceCreate creates a new http.Handler for the service create operation
func NewServiceCreate(ctx *middleware.Context, handler ServiceCreateHandler) *ServiceCreate {
	return &ServiceCreate{Context: ctx, Handler: handler}
}

/* ServiceCreate swagger:route POST /services/create Service serviceCreate

Create a service

*/
type ServiceCreate struct {
	Context *middleware.Context
	Handler ServiceCreateHandler
}

func (o *ServiceCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewServiceCreateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ServiceCreateBody service create body
//
// swagger:model ServiceCreateBody
type ServiceCreateBody struct {
	models.ServiceSpec

	ServiceCreateParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ServiceCreateBody) UnmarshalJSON(raw []byte) error {
	// ServiceCreateParamsBodyAO0
	var serviceCreateParamsBodyAO0 models.ServiceSpec
	if err := swag.ReadJSON(raw, &serviceCreateParamsBodyAO0); err != nil {
		return err
	}
	o.ServiceSpec = serviceCreateParamsBodyAO0

	// ServiceCreateParamsBodyAO1
	var serviceCreateParamsBodyAO1 ServiceCreateParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &serviceCreateParamsBodyAO1); err != nil {
		return err
	}
	o.ServiceCreateParamsBodyAllOf1 = serviceCreateParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ServiceCreateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	serviceCreateParamsBodyAO0, err := swag.WriteJSON(o.ServiceSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, serviceCreateParamsBodyAO0)

	serviceCreateParamsBodyAO1, err := swag.WriteJSON(o.ServiceCreateParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, serviceCreateParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this service create body
func (o *ServiceCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ServiceSpec
	if err := o.ServiceSpec.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ServiceCreateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this service create body based on the context it is used
func (o *ServiceCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.ServiceSpec
	if err := o.ServiceSpec.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with ServiceCreateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceCreateBody) UnmarshalBinary(b []byte) error {
	var res ServiceCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ServiceCreateCreatedBody ServiceCreateResponse
// Example: {"ID":"ak7w3gjqoa3kuz8xcpnyy0pvl","Warning":"unable to pin image doesnotexist:latest to digest: image library/doesnotexist:latest not found"}
//
// swagger:model ServiceCreateCreatedBody
type ServiceCreateCreatedBody struct {

	// The ID of the created service.
	ID string `json:"ID,omitempty"`

	// Optional warning message
	Warning string `json:"Warning,omitempty"`
}

// Validate validates this service create created body
func (o *ServiceCreateCreatedBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this service create created body based on context it is used
func (o *ServiceCreateCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ServiceCreateCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ServiceCreateCreatedBody) UnmarshalBinary(b []byte) error {
	var res ServiceCreateCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ServiceCreateParamsBodyAllOf1 service create params body all of1
// Example: {"EndpointSpec":{"Ports":[{"Protocol":"tcp","PublishedPort":8080,"TargetPort":80}]},"Labels":{"foo":"bar"},"Mode":{"Replicated":{"Replicas":4}},"Name":"web","RollbackConfig":{"Delay":1000000000,"FailureAction":"pause","MaxFailureRatio":0.15,"Monitor":15000000000,"Parallelism":1},"TaskTemplate":{"ContainerSpec":{"DNSConfig":{"Nameservers":["8.8.8.8"],"Options":["timeout:3"],"Search":["example.org"]},"Hosts":["10.10.10.10 host1","ABCD:EF01:2345:6789:ABCD:EF01:2345:6789 host2"],"Image":"nginx:alpine","Mounts":[{"ReadOnly":true,"Source":"web-data","Target":"/usr/share/nginx/html","Type":"volume","VolumeOptions":{"DriverConfig":{},"Labels":{"com.example.something":"something-value"}}}],"Secrets":[{"File":{"GID":"33","Mode":384,"Name":"www.example.org.key","UID":"33"},"SecretID":"fpjqlhnwb19zds35k8wn80lq9","SecretName":"example_org_domain_key"}],"User":"33"},"LogDriver":{"Name":"json-file","Options":{"max-file":"3","max-size":"10M"}},"Placement":{},"Resources":{"Limits":{"MemoryBytes":104857600},"Reservations":{}},"RestartPolicy":{"Condition":"on-failure","Delay":10000000000,"MaxAttempts":10}},"UpdateConfig":{"Delay":1000000000,"FailureAction":"pause","MaxFailureRatio":0.15,"Monitor":15000000000,"Parallelism":2}}
//
// swagger:model ServiceCreateParamsBodyAllOf1
type ServiceCreateParamsBodyAllOf1 interface{}