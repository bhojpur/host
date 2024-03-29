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

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// TaskInspectReader is a Reader for the TaskInspect structure.
type TaskInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TaskInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTaskInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewTaskInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTaskInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewTaskInspectServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTaskInspectOK creates a TaskInspectOK with default headers values
func NewTaskInspectOK() *TaskInspectOK {
	return &TaskInspectOK{}
}

/* TaskInspectOK describes a response with status code 200, with default header values.

no error
*/
type TaskInspectOK struct {
	Payload *models.Task
}

func (o *TaskInspectOK) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] taskInspectOK  %+v", 200, o.Payload)
}
func (o *TaskInspectOK) GetPayload() *models.Task {
	return o.Payload
}

func (o *TaskInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Task)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskInspectNotFound creates a TaskInspectNotFound with default headers values
func NewTaskInspectNotFound() *TaskInspectNotFound {
	return &TaskInspectNotFound{}
}

/* TaskInspectNotFound describes a response with status code 404, with default header values.

no such task
*/
type TaskInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *TaskInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] taskInspectNotFound  %+v", 404, o.Payload)
}
func (o *TaskInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskInspectInternalServerError creates a TaskInspectInternalServerError with default headers values
func NewTaskInspectInternalServerError() *TaskInspectInternalServerError {
	return &TaskInspectInternalServerError{}
}

/* TaskInspectInternalServerError describes a response with status code 500, with default header values.

server error
*/
type TaskInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *TaskInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] taskInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *TaskInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTaskInspectServiceUnavailable creates a TaskInspectServiceUnavailable with default headers values
func NewTaskInspectServiceUnavailable() *TaskInspectServiceUnavailable {
	return &TaskInspectServiceUnavailable{}
}

/* TaskInspectServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type TaskInspectServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *TaskInspectServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /tasks/{id}][%d] taskInspectServiceUnavailable  %+v", 503, o.Payload)
}
func (o *TaskInspectServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *TaskInspectServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
