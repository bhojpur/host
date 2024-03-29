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
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// VolumeInspectOKCode is the HTTP code returned for type VolumeInspectOK
const VolumeInspectOKCode int = 200

/*VolumeInspectOK No error

swagger:response volumeInspectOK
*/
type VolumeInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Volume `json:"body,omitempty"`
}

// NewVolumeInspectOK creates VolumeInspectOK with default headers values
func NewVolumeInspectOK() *VolumeInspectOK {

	return &VolumeInspectOK{}
}

// WithPayload adds the payload to the volume inspect o k response
func (o *VolumeInspectOK) WithPayload(payload *models.Volume) *VolumeInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the volume inspect o k response
func (o *VolumeInspectOK) SetPayload(payload *models.Volume) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VolumeInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VolumeInspectNotFoundCode is the HTTP code returned for type VolumeInspectNotFound
const VolumeInspectNotFoundCode int = 404

/*VolumeInspectNotFound No such volume

swagger:response volumeInspectNotFound
*/
type VolumeInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewVolumeInspectNotFound creates VolumeInspectNotFound with default headers values
func NewVolumeInspectNotFound() *VolumeInspectNotFound {

	return &VolumeInspectNotFound{}
}

// WithPayload adds the payload to the volume inspect not found response
func (o *VolumeInspectNotFound) WithPayload(payload *models.ErrorResponse) *VolumeInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the volume inspect not found response
func (o *VolumeInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VolumeInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VolumeInspectInternalServerErrorCode is the HTTP code returned for type VolumeInspectInternalServerError
const VolumeInspectInternalServerErrorCode int = 500

/*VolumeInspectInternalServerError Server error

swagger:response volumeInspectInternalServerError
*/
type VolumeInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewVolumeInspectInternalServerError creates VolumeInspectInternalServerError with default headers values
func NewVolumeInspectInternalServerError() *VolumeInspectInternalServerError {

	return &VolumeInspectInternalServerError{}
}

// WithPayload adds the payload to the volume inspect internal server error response
func (o *VolumeInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *VolumeInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the volume inspect internal server error response
func (o *VolumeInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VolumeInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
