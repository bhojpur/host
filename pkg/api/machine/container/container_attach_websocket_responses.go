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

// ContainerAttachWebsocketReader is a Reader for the ContainerAttachWebsocket structure.
type ContainerAttachWebsocketReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerAttachWebsocketReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 101:
		result := NewContainerAttachWebsocketSwitchingProtocols()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 200:
		result := NewContainerAttachWebsocketOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewContainerAttachWebsocketBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewContainerAttachWebsocketNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewContainerAttachWebsocketInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewContainerAttachWebsocketSwitchingProtocols creates a ContainerAttachWebsocketSwitchingProtocols with default headers values
func NewContainerAttachWebsocketSwitchingProtocols() *ContainerAttachWebsocketSwitchingProtocols {
	return &ContainerAttachWebsocketSwitchingProtocols{}
}

/* ContainerAttachWebsocketSwitchingProtocols describes a response with status code 101, with default header values.

no error, hints proxy about hijacking
*/
type ContainerAttachWebsocketSwitchingProtocols struct {
}

func (o *ContainerAttachWebsocketSwitchingProtocols) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/attach/ws][%d] containerAttachWebsocketSwitchingProtocols ", 101)
}

func (o *ContainerAttachWebsocketSwitchingProtocols) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachWebsocketOK creates a ContainerAttachWebsocketOK with default headers values
func NewContainerAttachWebsocketOK() *ContainerAttachWebsocketOK {
	return &ContainerAttachWebsocketOK{}
}

/* ContainerAttachWebsocketOK describes a response with status code 200, with default header values.

no error, no upgrade header found
*/
type ContainerAttachWebsocketOK struct {
}

func (o *ContainerAttachWebsocketOK) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/attach/ws][%d] containerAttachWebsocketOK ", 200)
}

func (o *ContainerAttachWebsocketOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewContainerAttachWebsocketBadRequest creates a ContainerAttachWebsocketBadRequest with default headers values
func NewContainerAttachWebsocketBadRequest() *ContainerAttachWebsocketBadRequest {
	return &ContainerAttachWebsocketBadRequest{}
}

/* ContainerAttachWebsocketBadRequest describes a response with status code 400, with default header values.

bad parameter
*/
type ContainerAttachWebsocketBadRequest struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachWebsocketBadRequest) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/attach/ws][%d] containerAttachWebsocketBadRequest  %+v", 400, o.Payload)
}
func (o *ContainerAttachWebsocketBadRequest) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachWebsocketBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachWebsocketNotFound creates a ContainerAttachWebsocketNotFound with default headers values
func NewContainerAttachWebsocketNotFound() *ContainerAttachWebsocketNotFound {
	return &ContainerAttachWebsocketNotFound{}
}

/* ContainerAttachWebsocketNotFound describes a response with status code 404, with default header values.

no such container
*/
type ContainerAttachWebsocketNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachWebsocketNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/attach/ws][%d] containerAttachWebsocketNotFound  %+v", 404, o.Payload)
}
func (o *ContainerAttachWebsocketNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachWebsocketNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerAttachWebsocketInternalServerError creates a ContainerAttachWebsocketInternalServerError with default headers values
func NewContainerAttachWebsocketInternalServerError() *ContainerAttachWebsocketInternalServerError {
	return &ContainerAttachWebsocketInternalServerError{}
}

/* ContainerAttachWebsocketInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ContainerAttachWebsocketInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ContainerAttachWebsocketInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/{id}/attach/ws][%d] containerAttachWebsocketInternalServerError  %+v", 500, o.Payload)
}
func (o *ContainerAttachWebsocketInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ContainerAttachWebsocketInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
