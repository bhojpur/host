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

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// PluginDisableReader is a Reader for the PluginDisable structure.
type PluginDisableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PluginDisableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPluginDisableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewPluginDisableNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPluginDisableInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPluginDisableOK creates a PluginDisableOK with default headers values
func NewPluginDisableOK() *PluginDisableOK {
	return &PluginDisableOK{}
}

/* PluginDisableOK describes a response with status code 200, with default header values.

no error
*/
type PluginDisableOK struct {
}

func (o *PluginDisableOK) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/disable][%d] pluginDisableOK ", 200)
}

func (o *PluginDisableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPluginDisableNotFound creates a PluginDisableNotFound with default headers values
func NewPluginDisableNotFound() *PluginDisableNotFound {
	return &PluginDisableNotFound{}
}

/* PluginDisableNotFound describes a response with status code 404, with default header values.

plugin is not installed
*/
type PluginDisableNotFound struct {
	Payload *models.ErrorResponse
}

func (o *PluginDisableNotFound) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/disable][%d] pluginDisableNotFound  %+v", 404, o.Payload)
}
func (o *PluginDisableNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginDisableNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPluginDisableInternalServerError creates a PluginDisableInternalServerError with default headers values
func NewPluginDisableInternalServerError() *PluginDisableInternalServerError {
	return &PluginDisableInternalServerError{}
}

/* PluginDisableInternalServerError describes a response with status code 500, with default header values.

server error
*/
type PluginDisableInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *PluginDisableInternalServerError) Error() string {
	return fmt.Sprintf("[POST /plugins/{name}/disable][%d] pluginDisableInternalServerError  %+v", 500, o.Payload)
}
func (o *PluginDisableInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PluginDisableInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
