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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ConfigInspectOKCode is the HTTP code returned for type ConfigInspectOK
const ConfigInspectOKCode int = 200

/*ConfigInspectOK no error

swagger:response configInspectOK
*/
type ConfigInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.Config `json:"body,omitempty"`
}

// NewConfigInspectOK creates ConfigInspectOK with default headers values
func NewConfigInspectOK() *ConfigInspectOK {

	return &ConfigInspectOK{}
}

// WithPayload adds the payload to the config inspect o k response
func (o *ConfigInspectOK) WithPayload(payload *models.Config) *ConfigInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config inspect o k response
func (o *ConfigInspectOK) SetPayload(payload *models.Config) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConfigInspectNotFoundCode is the HTTP code returned for type ConfigInspectNotFound
const ConfigInspectNotFoundCode int = 404

/*ConfigInspectNotFound config not found

swagger:response configInspectNotFound
*/
type ConfigInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewConfigInspectNotFound creates ConfigInspectNotFound with default headers values
func NewConfigInspectNotFound() *ConfigInspectNotFound {

	return &ConfigInspectNotFound{}
}

// WithPayload adds the payload to the config inspect not found response
func (o *ConfigInspectNotFound) WithPayload(payload *models.ErrorResponse) *ConfigInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config inspect not found response
func (o *ConfigInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConfigInspectInternalServerErrorCode is the HTTP code returned for type ConfigInspectInternalServerError
const ConfigInspectInternalServerErrorCode int = 500

/*ConfigInspectInternalServerError server error

swagger:response configInspectInternalServerError
*/
type ConfigInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewConfigInspectInternalServerError creates ConfigInspectInternalServerError with default headers values
func NewConfigInspectInternalServerError() *ConfigInspectInternalServerError {

	return &ConfigInspectInternalServerError{}
}

// WithPayload adds the payload to the config inspect internal server error response
func (o *ConfigInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *ConfigInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config inspect internal server error response
func (o *ConfigInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ConfigInspectServiceUnavailableCode is the HTTP code returned for type ConfigInspectServiceUnavailable
const ConfigInspectServiceUnavailableCode int = 503

/*ConfigInspectServiceUnavailable node is not part of a swarm

swagger:response configInspectServiceUnavailable
*/
type ConfigInspectServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewConfigInspectServiceUnavailable creates ConfigInspectServiceUnavailable with default headers values
func NewConfigInspectServiceUnavailable() *ConfigInspectServiceUnavailable {

	return &ConfigInspectServiceUnavailable{}
}

// WithPayload adds the payload to the config inspect service unavailable response
func (o *ConfigInspectServiceUnavailable) WithPayload(payload *models.ErrorResponse) *ConfigInspectServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the config inspect service unavailable response
func (o *ConfigInspectServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ConfigInspectServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}