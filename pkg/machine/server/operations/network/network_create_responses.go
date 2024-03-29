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

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// NetworkCreateCreatedCode is the HTTP code returned for type NetworkCreateCreated
const NetworkCreateCreatedCode int = 201

/*NetworkCreateCreated No error

swagger:response networkCreateCreated
*/
type NetworkCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *NetworkCreateCreatedBody `json:"body,omitempty"`
}

// NewNetworkCreateCreated creates NetworkCreateCreated with default headers values
func NewNetworkCreateCreated() *NetworkCreateCreated {

	return &NetworkCreateCreated{}
}

// WithPayload adds the payload to the network create created response
func (o *NetworkCreateCreated) WithPayload(payload *NetworkCreateCreatedBody) *NetworkCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network create created response
func (o *NetworkCreateCreated) SetPayload(payload *NetworkCreateCreatedBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkCreateForbiddenCode is the HTTP code returned for type NetworkCreateForbidden
const NetworkCreateForbiddenCode int = 403

/*NetworkCreateForbidden operation not supported for pre-defined networks

swagger:response networkCreateForbidden
*/
type NetworkCreateForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkCreateForbidden creates NetworkCreateForbidden with default headers values
func NewNetworkCreateForbidden() *NetworkCreateForbidden {

	return &NetworkCreateForbidden{}
}

// WithPayload adds the payload to the network create forbidden response
func (o *NetworkCreateForbidden) WithPayload(payload *models.ErrorResponse) *NetworkCreateForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network create forbidden response
func (o *NetworkCreateForbidden) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkCreateForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkCreateNotFoundCode is the HTTP code returned for type NetworkCreateNotFound
const NetworkCreateNotFoundCode int = 404

/*NetworkCreateNotFound plugin not found

swagger:response networkCreateNotFound
*/
type NetworkCreateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkCreateNotFound creates NetworkCreateNotFound with default headers values
func NewNetworkCreateNotFound() *NetworkCreateNotFound {

	return &NetworkCreateNotFound{}
}

// WithPayload adds the payload to the network create not found response
func (o *NetworkCreateNotFound) WithPayload(payload *models.ErrorResponse) *NetworkCreateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network create not found response
func (o *NetworkCreateNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkCreateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NetworkCreateInternalServerErrorCode is the HTTP code returned for type NetworkCreateInternalServerError
const NetworkCreateInternalServerErrorCode int = 500

/*NetworkCreateInternalServerError Server error

swagger:response networkCreateInternalServerError
*/
type NetworkCreateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNetworkCreateInternalServerError creates NetworkCreateInternalServerError with default headers values
func NewNetworkCreateInternalServerError() *NetworkCreateInternalServerError {

	return &NetworkCreateInternalServerError{}
}

// WithPayload adds the payload to the network create internal server error response
func (o *NetworkCreateInternalServerError) WithPayload(payload *models.ErrorResponse) *NetworkCreateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the network create internal server error response
func (o *NetworkCreateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NetworkCreateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
