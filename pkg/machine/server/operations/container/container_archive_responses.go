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

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ContainerArchiveOKCode is the HTTP code returned for type ContainerArchiveOK
const ContainerArchiveOKCode int = 200

/*ContainerArchiveOK no error

swagger:response containerArchiveOK
*/
type ContainerArchiveOK struct {
}

// NewContainerArchiveOK creates ContainerArchiveOK with default headers values
func NewContainerArchiveOK() *ContainerArchiveOK {

	return &ContainerArchiveOK{}
}

// WriteResponse to the client
func (o *ContainerArchiveOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ContainerArchiveBadRequestCode is the HTTP code returned for type ContainerArchiveBadRequest
const ContainerArchiveBadRequestCode int = 400

/*ContainerArchiveBadRequest Bad parameter

swagger:response containerArchiveBadRequest
*/
type ContainerArchiveBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerArchiveBadRequest creates ContainerArchiveBadRequest with default headers values
func NewContainerArchiveBadRequest() *ContainerArchiveBadRequest {

	return &ContainerArchiveBadRequest{}
}

// WithPayload adds the payload to the container archive bad request response
func (o *ContainerArchiveBadRequest) WithPayload(payload *models.ErrorResponse) *ContainerArchiveBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container archive bad request response
func (o *ContainerArchiveBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerArchiveBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerArchiveNotFoundCode is the HTTP code returned for type ContainerArchiveNotFound
const ContainerArchiveNotFoundCode int = 404

/*ContainerArchiveNotFound Container or path does not exist

swagger:response containerArchiveNotFound
*/
type ContainerArchiveNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerArchiveNotFound creates ContainerArchiveNotFound with default headers values
func NewContainerArchiveNotFound() *ContainerArchiveNotFound {

	return &ContainerArchiveNotFound{}
}

// WithPayload adds the payload to the container archive not found response
func (o *ContainerArchiveNotFound) WithPayload(payload *models.ErrorResponse) *ContainerArchiveNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container archive not found response
func (o *ContainerArchiveNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerArchiveNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerArchiveInternalServerErrorCode is the HTTP code returned for type ContainerArchiveInternalServerError
const ContainerArchiveInternalServerErrorCode int = 500

/*ContainerArchiveInternalServerError server error

swagger:response containerArchiveInternalServerError
*/
type ContainerArchiveInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerArchiveInternalServerError creates ContainerArchiveInternalServerError with default headers values
func NewContainerArchiveInternalServerError() *ContainerArchiveInternalServerError {

	return &ContainerArchiveInternalServerError{}
}

// WithPayload adds the payload to the container archive internal server error response
func (o *ContainerArchiveInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerArchiveInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container archive internal server error response
func (o *ContainerArchiveInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerArchiveInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
