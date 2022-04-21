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

package volume

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/bhojpur/host/pkg/machine/models"
)

// VolumeInspectReader is a Reader for the VolumeInspect structure.
type VolumeInspectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeInspectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewVolumeInspectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewVolumeInspectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewVolumeInspectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeInspectOK creates a VolumeInspectOK with default headers values
func NewVolumeInspectOK() *VolumeInspectOK {
	return &VolumeInspectOK{}
}

/* VolumeInspectOK describes a response with status code 200, with default header values.

No error
*/
type VolumeInspectOK struct {
	Payload *models.Volume
}

func (o *VolumeInspectOK) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectOK  %+v", 200, o.Payload)
}
func (o *VolumeInspectOK) GetPayload() *models.Volume {
	return o.Payload
}

func (o *VolumeInspectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectNotFound creates a VolumeInspectNotFound with default headers values
func NewVolumeInspectNotFound() *VolumeInspectNotFound {
	return &VolumeInspectNotFound{}
}

/* VolumeInspectNotFound describes a response with status code 404, with default header values.

No such volume
*/
type VolumeInspectNotFound struct {
	Payload *models.ErrorResponse
}

func (o *VolumeInspectNotFound) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectNotFound  %+v", 404, o.Payload)
}
func (o *VolumeInspectNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeInspectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeInspectInternalServerError creates a VolumeInspectInternalServerError with default headers values
func NewVolumeInspectInternalServerError() *VolumeInspectInternalServerError {
	return &VolumeInspectInternalServerError{}
}

/* VolumeInspectInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type VolumeInspectInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *VolumeInspectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /volumes/{name}][%d] volumeInspectInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeInspectInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeInspectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
