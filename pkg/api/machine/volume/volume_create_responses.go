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

// VolumeCreateReader is a Reader for the VolumeCreate structure.
type VolumeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewVolumeCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewVolumeCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewVolumeCreateCreated creates a VolumeCreateCreated with default headers values
func NewVolumeCreateCreated() *VolumeCreateCreated {
	return &VolumeCreateCreated{}
}

/* VolumeCreateCreated describes a response with status code 201, with default header values.

The volume was created successfully
*/
type VolumeCreateCreated struct {
	Payload *models.Volume
}

func (o *VolumeCreateCreated) Error() string {
	return fmt.Sprintf("[POST /volumes/create][%d] volumeCreateCreated  %+v", 201, o.Payload)
}
func (o *VolumeCreateCreated) GetPayload() *models.Volume {
	return o.Payload
}

func (o *VolumeCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Volume)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateInternalServerError creates a VolumeCreateInternalServerError with default headers values
func NewVolumeCreateInternalServerError() *VolumeCreateInternalServerError {
	return &VolumeCreateInternalServerError{}
}

/* VolumeCreateInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type VolumeCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *VolumeCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /volumes/create][%d] volumeCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *VolumeCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
