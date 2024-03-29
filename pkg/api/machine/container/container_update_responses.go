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
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ContainerUpdateReader is a Reader for the ContainerUpdate structure.
type ContainerUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerUpdateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerUpdateOK creates a ContainerUpdateOK with default headers values
func NewContainerUpdateOK() *ContainerUpdateOK {
	return &ContainerUpdateOK{}
}

/* ContainerUpdateOK describes a response with status code 200, with default header values.

The container has been updated.
*/
type ContainerUpdateOK struct {
	Payload *ContainerUpdateOKBody
}

func (o *ContainerUpdateOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/update][%d] containerUpdateOK  %+v", 200, o.Payload)
}
func (o *ContainerUpdateOK) GetPayload() *ContainerUpdateOKBody {
	return o.Payload
}

func (o *ContainerUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerUpdateOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerUpdateNotFound creates a ContainerUpdateNotFound with default headers values
func NewContainerUpdateNotFound() *ContainerUpdateNotFound {
	return &ContainerUpdateNotFound{}
}

/* ContainerUpdateNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerUpdateNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerUpdateNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/update][%d] containerUpdateNotFound  %+v", 404, o.Payload)
}
func (o *ContainerUpdateNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerUpdateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerUpdateInternalServerError creates a ContainerUpdateInternalServerError with default headers values
func NewContainerUpdateInternalServerError() *ContainerUpdateInternalServerError {
	return &ContainerUpdateInternalServerError{}
}

/* ContainerUpdateInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/update][%d] containerUpdateInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerUpdateBody container update body
// Example: {"BlkioWeight":300,"CpuPeriod":100000,"CpuQuota":50000,"CpuRealtimePeriod":1000000,"CpuRealtimeRuntime":10000,"CpuShares":512,"CpusetCpus":"0,1","CpusetMems":"0","Memory":314572800,"MemoryReservation":209715200,"MemorySwap":514288000,"RestartPolicy":{"MaximumRetryCount":4,"Name":"on-failure"}}
swagger:model ContainerUpdateBody
*/
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

/*ContainerUpdateOKBody ContainerUpdateResponse
//
// OK response to ContainerUpdate operation
swagger:model ContainerUpdateOKBody
*/
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
