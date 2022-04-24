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

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ImageInspectReader is a Reader for the ImageInspect structure.
type ImageInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageInspectOK creates a ImageInspectOK with default headers values
func NewImageInspectOK() *ImageInspectOK {
	return &ImageInspectOK{}
}

/* ImageInspectOK describes a response with status code 200, with default header values.

No error
*/
type ImageInspectOK struct {
	Payload *models.ImageInspect
}

func (o *ImageInspectOK) Error() string {
	return fmt.Sprintf("[GET /images/{name}/json][%d] imageInspectOK  %+v", 200, o.Payload)
}
func (o *ImageInspectOK) GetPayload() *models.ImageInspect {
	return o.Payload
}

func (o *ImageInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ImageInspect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectNotFound creates a ImageInspectNotFound with default headers values
func NewImageInspectNotFound() *ImageInspectNotFound {
	return &ImageInspectNotFound{}
}

/* ImageInspectNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ImageInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /images/{name}/json][%d] imageInspectNotFound  %+v", 404, o.Payload)
}
func (o *ImageInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageInspectInternalServerError creates a ImageInspectInternalServerError with default headers values
func NewImageInspectInternalServerError() *ImageInspectInternalServerError {
	return &ImageInspectInternalServerError{}
}

/* ImageInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImageInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{name}/json][%d] imageInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}