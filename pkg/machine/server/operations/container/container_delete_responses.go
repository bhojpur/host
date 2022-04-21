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

// ContainerDeleteNoContentCode is the HTTP code returned for type ContainerDeleteNoContent
const ContainerDeleteNoContentCode int = 204

/*ContainerDeleteNoContent no error

swagger:response containerDeleteNoContent
*/
type ContainerDeleteNoContent struct {
}

// NewContainerDeleteNoContent creates ContainerDeleteNoContent with default headers values
func NewContainerDeleteNoContent() *ContainerDeleteNoContent {

	return &ContainerDeleteNoContent{}
}

// WriteResponse to the client
func (o *ContainerDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ContainerDeleteBadRequestCode is the HTTP code returned for type ContainerDeleteBadRequest
const ContainerDeleteBadRequestCode int = 400

/*ContainerDeleteBadRequest bad parameter

swagger:response containerDeleteBadRequest
*/
type ContainerDeleteBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerDeleteBadRequest creates ContainerDeleteBadRequest with default headers values
func NewContainerDeleteBadRequest() *ContainerDeleteBadRequest {

	return &ContainerDeleteBadRequest{}
}

// WithPayload adds the payload to the container delete bad request response
func (o *ContainerDeleteBadRequest) WithPayload(payload *models.ErrorResponse) *ContainerDeleteBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container delete bad request response
func (o *ContainerDeleteBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerDeleteBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerDeleteNotFoundCode is the HTTP code returned for type ContainerDeleteNotFound
const ContainerDeleteNotFoundCode int = 404

/*ContainerDeleteNotFound no such container

swagger:response containerDeleteNotFound
*/
type ContainerDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerDeleteNotFound creates ContainerDeleteNotFound with default headers values
func NewContainerDeleteNotFound() *ContainerDeleteNotFound {

	return &ContainerDeleteNotFound{}
}

// WithPayload adds the payload to the container delete not found response
func (o *ContainerDeleteNotFound) WithPayload(payload *models.ErrorResponse) *ContainerDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container delete not found response
func (o *ContainerDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerDeleteConflictCode is the HTTP code returned for type ContainerDeleteConflict
const ContainerDeleteConflictCode int = 409

/*ContainerDeleteConflict conflict

swagger:response containerDeleteConflict
*/
type ContainerDeleteConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerDeleteConflict creates ContainerDeleteConflict with default headers values
func NewContainerDeleteConflict() *ContainerDeleteConflict {

	return &ContainerDeleteConflict{}
}

// WithPayload adds the payload to the container delete conflict response
func (o *ContainerDeleteConflict) WithPayload(payload *models.ErrorResponse) *ContainerDeleteConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container delete conflict response
func (o *ContainerDeleteConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerDeleteConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerDeleteInternalServerErrorCode is the HTTP code returned for type ContainerDeleteInternalServerError
const ContainerDeleteInternalServerErrorCode int = 500

/*ContainerDeleteInternalServerError server error

swagger:response containerDeleteInternalServerError
*/
type ContainerDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerDeleteInternalServerError creates ContainerDeleteInternalServerError with default headers values
func NewContainerDeleteInternalServerError() *ContainerDeleteInternalServerError {

	return &ContainerDeleteInternalServerError{}
}

// WithPayload adds the payload to the container delete internal server error response
func (o *ContainerDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container delete internal server error response
func (o *ContainerDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
