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

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// NodeDeleteReader is a Reader for the NodeDelete structure.
type NodeDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewNodeDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewNodeDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNodeDeleteServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodeDeleteOK creates a NodeDeleteOK with default headers values
func NewNodeDeleteOK() *NodeDeleteOK {
	return &NodeDeleteOK{}
}

/* NodeDeleteOK describes a response with status code 200, with default header values.

no error
*/
type NodeDeleteOK struct {
}

func (o *NodeDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{id}][%d] nodeDeleteOK ", 200)
}

func (o *NodeDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNodeDeleteNotFound creates a NodeDeleteNotFound with default headers values
func NewNodeDeleteNotFound() *NodeDeleteNotFound {
	return &NodeDeleteNotFound{}
}

/* NodeDeleteNotFound describes a response with status code 404, with default header values.

no such node
*/
type NodeDeleteNotFound struct {
	Payload *models.ErrorResponse
}

func (o *NodeDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{id}][%d] nodeDeleteNotFound  %+v", 404, o.Payload)
}
func (o *NodeDeleteNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeDeleteInternalServerError creates a NodeDeleteInternalServerError with default headers values
func NewNodeDeleteInternalServerError() *NodeDeleteInternalServerError {
	return &NodeDeleteInternalServerError{}
}

/* NodeDeleteInternalServerError describes a response with status code 500, with default header values.

server error
*/
type NodeDeleteInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *NodeDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{id}][%d] nodeDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *NodeDeleteInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeDeleteServiceUnavailable creates a NodeDeleteServiceUnavailable with default headers values
func NewNodeDeleteServiceUnavailable() *NodeDeleteServiceUnavailable {
	return &NodeDeleteServiceUnavailable{}
}

/* NodeDeleteServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type NodeDeleteServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *NodeDeleteServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /nodes/{id}][%d] nodeDeleteServiceUnavailable  %+v", 503, o.Payload)
}
func (o *NodeDeleteServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeDeleteServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}