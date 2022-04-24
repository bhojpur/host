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

package container

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

// ContainerUpdateHandlerFunc turns a function with the right signature into a container update handler
type ContainerUpdateHandlerFunc func(ContainerUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn ContainerUpdateHandlerFunc) Handle(params ContainerUpdateParams) middleware.Responder {
	return fn(params)
}

// ContainerUpdateHandler interface for that can handle valid container update params
type ContainerUpdateHandler interface {
	Handle(ContainerUpdateParams) middleware.Responder
}

// NewContainerUpdate creates a new http.Handler for the container update operation
func NewContainerUpdate(ctx *middleware.Context, handler ContainerUpdateHandler) *ContainerUpdate {
	return &ContainerUpdate{Context: ctx, Handler: handler}
}

/* ContainerUpdate swagger:route POST /containers/{id}/update Container containerUpdate

Update a container

Change various configuration options of a container without having to
recreate it.


*/
type ContainerUpdate struct {
	Context *middleware.Context
	Handler ContainerUpdateHandler
}

func (o *ContainerUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewContainerUpdateParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// ContainerUpdateBody container update body
// Example: {"BlkioWeight":300,"CpuPeriod":100000,"CpuQuota":50000,"CpuRealtimePeriod":1000000,"CpuRealtimeRuntime":10000,"CpuShares":512,"CpusetCpus":"0,1","CpusetMems":"0","Memory":314572800,"MemoryReservation":209715200,"MemorySwap":514288000,"RestartPolicy":{"MaximumRetryCount":4,"Name":"on-failure"}}
//
// swagger:model ContainerUpdateBody
type ContainerUpdateBody struct {
	models.Resources

	// restart policy
	RestartPolicy *models.RestartPolicy `json:"RestartPolicy,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *ContainerUpdateBody) UnmarshalJSON(raw []byte) error {
	// ContainerUpdateParamsBodyAO0
	var containerUpdateParamsBodyAO0 models.Resources
	if err := swag.ReadJSON(raw, &containerUpdateParamsBodyAO0); err != nil {
		return err
	}
	o.Resources = containerUpdateParamsBodyAO0

	// ContainerUpdateParamsBodyAO1
	var dataContainerUpdateParamsBodyAO1 struct {
		RestartPolicy *models.RestartPolicy `json:"RestartPolicy,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataContainerUpdateParamsBodyAO1); err != nil {
		return err
	}

	o.RestartPolicy = dataContainerUpdateParamsBodyAO1.RestartPolicy

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o ContainerUpdateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	containerUpdateParamsBodyAO0, err := swag.WriteJSON(o.Resources)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, containerUpdateParamsBodyAO0)
	var dataContainerUpdateParamsBodyAO1 struct {
		RestartPolicy *models.RestartPolicy `json:"RestartPolicy,omitempty"`
	}

	dataContainerUpdateParamsBodyAO1.RestartPolicy = o.RestartPolicy

	jsonDataContainerUpdateParamsBodyAO1, errContainerUpdateParamsBodyAO1 := swag.WriteJSON(dataContainerUpdateParamsBodyAO1)
	if errContainerUpdateParamsBodyAO1 != nil {
		return nil, errContainerUpdateParamsBodyAO1
	}
	_parts = append(_parts, jsonDataContainerUpdateParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this container update body
func (o *ContainerUpdateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Resources
	if err := o.Resources.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRestartPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerUpdateBody) validateRestartPolicy(formats strfmt.Registry) error {

	if swag.IsZero(o.RestartPolicy) { // not required
		return nil
	}

	if o.RestartPolicy != nil {
		if err := o.RestartPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update" + "." + "RestartPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update" + "." + "RestartPolicy")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this container update body based on the context it is used
func (o *ContainerUpdateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Resources
	if err := o.Resources.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRestartPolicy(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ContainerUpdateBody) contextValidateRestartPolicy(ctx context.Context, formats strfmt.Registry) error {

	if o.RestartPolicy != nil {
		if err := o.RestartPolicy.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("update" + "." + "RestartPolicy")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("update" + "." + "RestartPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ContainerUpdateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerUpdateBody) UnmarshalBinary(b []byte) error {
	var res ContainerUpdateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

// ContainerUpdateOKBody ContainerUpdateResponse
//
// OK response to ContainerUpdate operation
//
// swagger:model ContainerUpdateOKBody
type ContainerUpdateOKBody struct {

	// warnings
	Warnings []string `json:"Warnings"`
}

// Validate validates this container update o k body
func (o *ContainerUpdateOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container update o k body based on context it is used
func (o *ContainerUpdateOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerUpdateOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerUpdateOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerUpdateOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}