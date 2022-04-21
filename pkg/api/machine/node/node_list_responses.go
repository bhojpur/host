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

// NodeListReader is a Reader for the NodeList structure.
type NodeListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NodeListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNodeListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewNodeListInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewNodeListServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewNodeListOK creates a NodeListOK with default headers values
func NewNodeListOK() *NodeListOK {
	return &NodeListOK{}
}

/* NodeListOK describes a response with status code 200, with default header values.

no error
*/
type NodeListOK struct {
	Payload []*models.Node
}

func (o *NodeListOK) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListOK  %+v", 200, o.Payload)
}
func (o *NodeListOK) GetPayload() []*models.Node {
	return o.Payload
}

func (o *NodeListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeListInternalServerError creates a NodeListInternalServerError with default headers values
func NewNodeListInternalServerError() *NodeListInternalServerError {
	return &NodeListInternalServerError{}
}

/* NodeListInternalServerError describes a response with status code 500, with default header values.

server error
*/
type NodeListInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *NodeListInternalServerError) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListInternalServerError  %+v", 500, o.Payload)
}
func (o *NodeListInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeListInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNodeListServiceUnavailable creates a NodeListServiceUnavailable with default headers values
func NewNodeListServiceUnavailable() *NodeListServiceUnavailable {
	return &NodeListServiceUnavailable{}
}

/* NodeListServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type NodeListServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *NodeListServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /nodes][%d] nodeListServiceUnavailable  %+v", 503, o.Payload)
}
func (o *NodeListServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NodeListServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
