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

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// SwarmInitOKCode is the HTTP code returned for type SwarmInitOK
const SwarmInitOKCode int = 200

/*SwarmInitOK no error

swagger:response swarmInitOK
*/
type SwarmInitOK struct {

	/*The node ID
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewSwarmInitOK creates SwarmInitOK with default headers values
func NewSwarmInitOK() *SwarmInitOK {

	return &SwarmInitOK{}
}

// WithPayload adds the payload to the swarm init o k response
func (o *SwarmInitOK) WithPayload(payload string) *SwarmInitOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm init o k response
func (o *SwarmInitOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmInitOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// SwarmInitBadRequestCode is the HTTP code returned for type SwarmInitBadRequest
const SwarmInitBadRequestCode int = 400

/*SwarmInitBadRequest bad parameter

swagger:response swarmInitBadRequest
*/
type SwarmInitBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSwarmInitBadRequest creates SwarmInitBadRequest with default headers values
func NewSwarmInitBadRequest() *SwarmInitBadRequest {

	return &SwarmInitBadRequest{}
}

// WithPayload adds the payload to the swarm init bad request response
func (o *SwarmInitBadRequest) WithPayload(payload *models.ErrorResponse) *SwarmInitBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm init bad request response
func (o *SwarmInitBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmInitBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SwarmInitInternalServerErrorCode is the HTTP code returned for type SwarmInitInternalServerError
const SwarmInitInternalServerErrorCode int = 500

/*SwarmInitInternalServerError server error

swagger:response swarmInitInternalServerError
*/
type SwarmInitInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSwarmInitInternalServerError creates SwarmInitInternalServerError with default headers values
func NewSwarmInitInternalServerError() *SwarmInitInternalServerError {

	return &SwarmInitInternalServerError{}
}

// WithPayload adds the payload to the swarm init internal server error response
func (o *SwarmInitInternalServerError) WithPayload(payload *models.ErrorResponse) *SwarmInitInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm init internal server error response
func (o *SwarmInitInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmInitInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SwarmInitServiceUnavailableCode is the HTTP code returned for type SwarmInitServiceUnavailable
const SwarmInitServiceUnavailableCode int = 503

/*SwarmInitServiceUnavailable node is already part of a swarm

swagger:response swarmInitServiceUnavailable
*/
type SwarmInitServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSwarmInitServiceUnavailable creates SwarmInitServiceUnavailable with default headers values
func NewSwarmInitServiceUnavailable() *SwarmInitServiceUnavailable {

	return &SwarmInitServiceUnavailable{}
}

// WithPayload adds the payload to the swarm init service unavailable response
func (o *SwarmInitServiceUnavailable) WithPayload(payload *models.ErrorResponse) *SwarmInitServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the swarm init service unavailable response
func (o *SwarmInitServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SwarmInitServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
