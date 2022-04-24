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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ConfigInspectReader is a Reader for the ConfigInspect structure.
type ConfigInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConfigInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConfigInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewConfigInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewConfigInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewConfigInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewConfigInspectOK creates a ConfigInspectOK with default headers values
func NewConfigInspectOK() *ConfigInspectOK {
	return &ConfigInspectOK{}
}

/* ConfigInspectOK describes a response with status code 200, with default header values.

no error
*/
type ConfigInspectOK struct {
	Payload *models.Config
}

func (o *ConfigInspectOK) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] configInspectOK  %+v", 200, o.Payload)
}
func (o *ConfigInspectOK) GetPayload() *models.Config {
	return o.Payload
}

func (o *ConfigInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Config)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigInspectNotFound creates a ConfigInspectNotFound with default headers values
func NewConfigInspectNotFound() *ConfigInspectNotFound {
	return &ConfigInspectNotFound{}
}

/* ConfigInspectNotFound describes a response with status code 404, with default header values.

config not found
*/
type ConfigInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ConfigInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] configInspectNotFound  %+v", 404, o.Payload)
}
func (o *ConfigInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigInspectInternalServerError creates a ConfigInspectInternalServerError with default headers values
func NewConfigInspectInternalServerError() *ConfigInspectInternalServerError {
	return &ConfigInspectInternalServerError{}
}

/* ConfigInspectInternalServerError describes a response with status code 500, with default header values.

server error
*/
type ConfigInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ConfigInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] configInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *ConfigInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConfigInspectServiceUnavailable creates a ConfigInspectServiceUnavailable with default headers values
func NewConfigInspectServiceUnavailable() *ConfigInspectServiceUnavailable {
	return &ConfigInspectServiceUnavailable{}
}

/* ConfigInspectServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type ConfigInspectServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *ConfigInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] configInspectServiceUnavailable  %+v", 503, o.Payload)
}
func (o *ConfigInspectServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ConfigInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}