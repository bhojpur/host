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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ContainerTopReader is a Reader for the ContainerTop structure.
type ContainerTopReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerTopReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContainerTopOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerTopNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerTopInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerTopOK creates a ContainerTopOK with default headers values
func NewContainerTopOK() *ContainerTopOK {
	return &ContainerTopOK{}
}

/* ContainerTopOK describes a response with status code 200, with default header values.

no error
*/
type ContainerTopOK struct {
	Payload *ContainerTopOKBody
}

func (o *ContainerTopOK) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/top][%d] containerTopOK  %+v", 200, o.Payload)
}
func (o *ContainerTopOK) GetPayload() *ContainerTopOKBody {
	return o.Payload
}

func (o *ContainerTopOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ContainerTopOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerTopNotFound creates a ContainerTopNotFound with default headers values
func NewContainerTopNotFound() *ContainerTopNotFound {
	return &ContainerTopNotFound{}
}

/* ContainerTopNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerTopNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerTopNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/top][%d] containerTopNotFound  %+v", 404, o.Payload)
}
func (o *ContainerTopNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerTopNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerTopInternalServerError creates a ContainerTopInternalServerError with default headers values
func NewContainerTopInternalServerError() *ContainerTopInternalServerError {
	return &ContainerTopInternalServerError{}
}

/* ContainerTopInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerTopInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerTopInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/top][%d] containerTopInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerTopInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerTopInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ContainerTopOKBody ContainerTopResponse
//
// OK response to ContainerTop operation
swagger:model ContainerTopOKBody
*/
type ContainerTopOKBody struct {

	// Each process running in the container, where each is process
	// is an array of values corresponding to the titles.
	//
	Processes [][]string `json:"Processes"`

	// The ps column titles
	Titles []string `json:"Titles"`
}

// Validate validates this container top o k body
func (o *ContainerTopOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this container top o k body based on context it is used
func (o *ContainerTopOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ContainerTopOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ContainerTopOKBody) UnmarshalBinary(b []byte) error {
	var res ContainerTopOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
