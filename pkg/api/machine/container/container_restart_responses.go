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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ContainerRestartReader is a Reader for the ContainerRestart structure.
type ContainerRestartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRestartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContainerRestartNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewContainerRestartNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerRestartInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerRestartNoContent creates a ContainerRestartNoContent with default headers values
func NewContainerRestartNoContent() *ContainerRestartNoContent {
	return &ContainerRestartNoContent{}
}

/* ContainerRestartNoContent describes a response with status code 204, with default header values.

no error
*/
type ContainerRestartNoContent struct {
}

func (o *ContainerRestartNoContent) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/restart][%d] containerRestartNoContent ", 204)
}

func (o *ContainerRestartNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerRestartNotFound creates a ContainerRestartNotFound with default headers values
func NewContainerRestartNotFound() *ContainerRestartNotFound {
	return &ContainerRestartNotFound{}
}

/* ContainerRestartNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerRestartNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerRestartNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/restart][%d] containerRestartNotFound  %+v", 404, o.Payload)
}
func (o *ContainerRestartNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerRestartNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRestartInternalServerError creates a ContainerRestartInternalServerError with default headers values
func NewContainerRestartInternalServerError() *ContainerRestartInternalServerError {
	return &ContainerRestartInternalServerError{}
}

/* ContainerRestartInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerRestartInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerRestartInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/restart][%d] containerRestartInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerRestartInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerRestartInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}