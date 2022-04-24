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

package secret

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// SecretDeleteNoContentCode is the HTTP code returned for type SecretDeleteNoContent
const SecretDeleteNoContentCode int = 204

/*SecretDeleteNoContent no error

swagger:response secretDeleteNoContent
*/
type SecretDeleteNoContent struct {
}

// NewSecretDeleteNoContent creates SecretDeleteNoContent with default headers values
func NewSecretDeleteNoContent() *SecretDeleteNoContent {

	return &SecretDeleteNoContent{}
}

// WriteResponse to the client
func (o *SecretDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// SecretDeleteNotFoundCode is the HTTP code returned for type SecretDeleteNotFound
const SecretDeleteNotFoundCode int = 404

/*SecretDeleteNotFound secret not found

swagger:response secretDeleteNotFound
*/
type SecretDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretDeleteNotFound creates SecretDeleteNotFound with default headers values
func NewSecretDeleteNotFound() *SecretDeleteNotFound {

	return &SecretDeleteNotFound{}
}

// WithPayload adds the payload to the secret delete not found response
func (o *SecretDeleteNotFound) WithPayload(payload *models.ErrorResponse) *SecretDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret delete not found response
func (o *SecretDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretDeleteInternalServerErrorCode is the HTTP code returned for type SecretDeleteInternalServerError
const SecretDeleteInternalServerErrorCode int = 500

/*SecretDeleteInternalServerError server error

swagger:response secretDeleteInternalServerError
*/
type SecretDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretDeleteInternalServerError creates SecretDeleteInternalServerError with default headers values
func NewSecretDeleteInternalServerError() *SecretDeleteInternalServerError {

	return &SecretDeleteInternalServerError{}
}

// WithPayload adds the payload to the secret delete internal server error response
func (o *SecretDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *SecretDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret delete internal server error response
func (o *SecretDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretDeleteServiceUnavailableCode is the HTTP code returned for type SecretDeleteServiceUnavailable
const SecretDeleteServiceUnavailableCode int = 503

/*SecretDeleteServiceUnavailable node is not part of a swarm

swagger:response secretDeleteServiceUnavailable
*/
type SecretDeleteServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretDeleteServiceUnavailable creates SecretDeleteServiceUnavailable with default headers values
func NewSecretDeleteServiceUnavailable() *SecretDeleteServiceUnavailable {

	return &SecretDeleteServiceUnavailable{}
}

// WithPayload adds the payload to the secret delete service unavailable response
func (o *SecretDeleteServiceUnavailable) WithPayload(payload *models.ErrorResponse) *SecretDeleteServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret delete service unavailable response
func (o *SecretDeleteServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretDeleteServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}