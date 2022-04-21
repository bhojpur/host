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

// SecretUpdateOKCode is the HTTP code returned for type SecretUpdateOK
const SecretUpdateOKCode int = 200

/*SecretUpdateOK no error

swagger:response secretUpdateOK
*/
type SecretUpdateOK struct {
}

// NewSecretUpdateOK creates SecretUpdateOK with default headers values
func NewSecretUpdateOK() *SecretUpdateOK {

	return &SecretUpdateOK{}
}

// WriteResponse to the client
func (o *SecretUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// SecretUpdateBadRequestCode is the HTTP code returned for type SecretUpdateBadRequest
const SecretUpdateBadRequestCode int = 400

/*SecretUpdateBadRequest bad parameter

swagger:response secretUpdateBadRequest
*/
type SecretUpdateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretUpdateBadRequest creates SecretUpdateBadRequest with default headers values
func NewSecretUpdateBadRequest() *SecretUpdateBadRequest {

	return &SecretUpdateBadRequest{}
}

// WithPayload adds the payload to the secret update bad request response
func (o *SecretUpdateBadRequest) WithPayload(payload *models.ErrorResponse) *SecretUpdateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret update bad request response
func (o *SecretUpdateBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretUpdateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretUpdateNotFoundCode is the HTTP code returned for type SecretUpdateNotFound
const SecretUpdateNotFoundCode int = 404

/*SecretUpdateNotFound no such secret

swagger:response secretUpdateNotFound
*/
type SecretUpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretUpdateNotFound creates SecretUpdateNotFound with default headers values
func NewSecretUpdateNotFound() *SecretUpdateNotFound {

	return &SecretUpdateNotFound{}
}

// WithPayload adds the payload to the secret update not found response
func (o *SecretUpdateNotFound) WithPayload(payload *models.ErrorResponse) *SecretUpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret update not found response
func (o *SecretUpdateNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretUpdateInternalServerErrorCode is the HTTP code returned for type SecretUpdateInternalServerError
const SecretUpdateInternalServerErrorCode int = 500

/*SecretUpdateInternalServerError server error

swagger:response secretUpdateInternalServerError
*/
type SecretUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretUpdateInternalServerError creates SecretUpdateInternalServerError with default headers values
func NewSecretUpdateInternalServerError() *SecretUpdateInternalServerError {

	return &SecretUpdateInternalServerError{}
}

// WithPayload adds the payload to the secret update internal server error response
func (o *SecretUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *SecretUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret update internal server error response
func (o *SecretUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SecretUpdateServiceUnavailableCode is the HTTP code returned for type SecretUpdateServiceUnavailable
const SecretUpdateServiceUnavailableCode int = 503

/*SecretUpdateServiceUnavailable node is not part of a swarm

swagger:response secretUpdateServiceUnavailable
*/
type SecretUpdateServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSecretUpdateServiceUnavailable creates SecretUpdateServiceUnavailable with default headers values
func NewSecretUpdateServiceUnavailable() *SecretUpdateServiceUnavailable {

	return &SecretUpdateServiceUnavailable{}
}

// WithPayload adds the payload to the secret update service unavailable response
func (o *SecretUpdateServiceUnavailable) WithPayload(payload *models.ErrorResponse) *SecretUpdateServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the secret update service unavailable response
func (o *SecretUpdateServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SecretUpdateServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
