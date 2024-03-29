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

// ContainerPauseNoContentCode is the HTTP code returned for type ContainerPauseNoContent
const ContainerPauseNoContentCode int = 204

/*ContainerPauseNoContent no error

swagger:response containerPauseNoContent
*/
type ContainerPauseNoContent struct {
}

// NewContainerPauseNoContent creates ContainerPauseNoContent with default headers values
func NewContainerPauseNoContent() *ContainerPauseNoContent {

	return &ContainerPauseNoContent{}
}

// WriteResponse to the client
func (o *ContainerPauseNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// ContainerPauseNotFoundCode is the HTTP code returned for type ContainerPauseNotFound
const ContainerPauseNotFoundCode int = 404

/*ContainerPauseNotFound no such container

swagger:response containerPauseNotFound
*/
type ContainerPauseNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerPauseNotFound creates ContainerPauseNotFound with default headers values
func NewContainerPauseNotFound() *ContainerPauseNotFound {

	return &ContainerPauseNotFound{}
}

// WithPayload adds the payload to the container pause not found response
func (o *ContainerPauseNotFound) WithPayload(payload *models.ErrorResponse) *ContainerPauseNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container pause not found response
func (o *ContainerPauseNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerPauseNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerPauseInternalServerErrorCode is the HTTP code returned for type ContainerPauseInternalServerError
const ContainerPauseInternalServerErrorCode int = 500

/*ContainerPauseInternalServerError server error

swagger:response containerPauseInternalServerError
*/
type ContainerPauseInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerPauseInternalServerError creates ContainerPauseInternalServerError with default headers values
func NewContainerPauseInternalServerError() *ContainerPauseInternalServerError {

	return &ContainerPauseInternalServerError{}
}

// WithPayload adds the payload to the container pause internal server error response
func (o *ContainerPauseInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerPauseInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container pause internal server error response
func (o *ContainerPauseInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerPauseInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
