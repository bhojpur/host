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

// ContainerAttachReader is a Reader for the ContainerAttach structure.
type ContainerAttachReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerAttachReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 101:
		result := NewContainerAttachSwitchingProtocols()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 200:
		result := NewContainerAttachOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerAttachBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerAttachNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerAttachInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerAttachSwitchingProtocols creates a ContainerAttachSwitchingProtocols with default headers values
func NewContainerAttachSwitchingProtocols() *ContainerAttachSwitchingProtocols {
	return &ContainerAttachSwitchingProtocols{}
}

/* ContainerAttachSwitchingProtocols describes a response with status code 101, with default header values.

no error, hints proxy about hijacking
*/
type ContainerAttachSwitchingProtocols struct {
}

func (o *ContainerAttachSwitchingProtocols) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachSwitchingProtocols ", 101)
}

func (o *ContainerAttachSwitchingProtocols) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachOK creates a ContainerAttachOK with default headers values
func NewContainerAttachOK() *ContainerAttachOK {
	return &ContainerAttachOK{}
}

/* ContainerAttachOK describes a response with status code 200, with default header values.

no error, no upgrade header found
*/
type ContainerAttachOK struct {
}

func (o *ContainerAttachOK) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachOK ", 200)
}

func (o *ContainerAttachOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachBadRequest creates a ContainerAttachBadRequest with default headers values
func NewContainerAttachBadRequest() *ContainerAttachBadRequest {
	return &ContainerAttachBadRequest{}
}

/* ContainerAttachBadRequest describes a response with status code 400, with default header values.

bad parameter
*/
type ContainerAttachBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachBadRequest) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachBadRequest  %+v", 400, o.Payload)
}
func (o *ContainerAttachBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachNotFound creates a ContainerAttachNotFound with default headers values
func NewContainerAttachNotFound() *ContainerAttachNotFound {
	return &ContainerAttachNotFound{}
}

/* ContainerAttachNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerAttachNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachNotFound) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachNotFound  %+v", 404, o.Payload)
}
func (o *ContainerAttachNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachInternalServerError creates a ContainerAttachInternalServerError with default headers values
func NewContainerAttachInternalServerError() *ContainerAttachInternalServerError {
	return &ContainerAttachInternalServerError{}
}

/* ContainerAttachInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerAttachInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachInternalServerError) Error() string {
	return fmt.Sprintf("[POST /containers/{id}/attach][%d] containerAttachInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerAttachInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}