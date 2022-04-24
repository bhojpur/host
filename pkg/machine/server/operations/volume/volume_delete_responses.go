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

// VolumeDeleteNoContentCode is the HTTP code returned for type VolumeDeleteNoContent
const VolumeDeleteNoContentCode int = 204

/*VolumeDeleteNoContent The volume was removed

swagger:response volumeDeleteNoContent
*/
type VolumeDeleteNoContent struct {
}

// NewVolumeDeleteNoContent creates VolumeDeleteNoContent with default headers values
func NewVolumeDeleteNoContent() *VolumeDeleteNoContent {

	return &VolumeDeleteNoContent{}
}

// WriteResponse to the client
func (o *VolumeDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// VolumeDeleteNotFoundCode is the HTTP code returned for type VolumeDeleteNotFound
const VolumeDeleteNotFoundCode int = 404

/*VolumeDeleteNotFound No such volume or volume driver

swagger:response volumeDeleteNotFound
*/
type VolumeDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewVolumeDeleteNotFound creates VolumeDeleteNotFound with default headers values
func NewVolumeDeleteNotFound() *VolumeDeleteNotFound {

	return &VolumeDeleteNotFound{}
}

// WithPayload adds the payload to the volume delete not found response
func (o *VolumeDeleteNotFound) WithPayload(payload *models.ErrorResponse) *VolumeDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the volume delete not found response
func (o *VolumeDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VolumeDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VolumeDeleteConflictCode is the HTTP code returned for type VolumeDeleteConflict
const VolumeDeleteConflictCode int = 409

/*VolumeDeleteConflict Volume is in use and cannot be removed

swagger:response volumeDeleteConflict
*/
type VolumeDeleteConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewVolumeDeleteConflict creates VolumeDeleteConflict with default headers values
func NewVolumeDeleteConflict() *VolumeDeleteConflict {

	return &VolumeDeleteConflict{}
}

// WithPayload adds the payload to the volume delete conflict response
func (o *VolumeDeleteConflict) WithPayload(payload *models.ErrorResponse) *VolumeDeleteConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the volume delete conflict response
func (o *VolumeDeleteConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VolumeDeleteConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// VolumeDeleteInternalServerErrorCode is the HTTP code returned for type VolumeDeleteInternalServerError
const VolumeDeleteInternalServerErrorCode int = 500

/*VolumeDeleteInternalServerError Server error

swagger:response volumeDeleteInternalServerError
*/
type VolumeDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewVolumeDeleteInternalServerError creates VolumeDeleteInternalServerError with default headers values
func NewVolumeDeleteInternalServerError() *VolumeDeleteInternalServerError {

	return &VolumeDeleteInternalServerError{}
}

// WithPayload adds the payload to the volume delete internal server error response
func (o *VolumeDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *VolumeDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the volume delete internal server error response
func (o *VolumeDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *VolumeDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}