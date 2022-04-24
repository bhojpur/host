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

package plugin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// PluginEnableOKCode is the HTTP code returned for type PluginEnableOK
const PluginEnableOKCode int = 200

/*PluginEnableOK no error

swagger:response pluginEnableOK
*/
type PluginEnableOK struct {
}

// NewPluginEnableOK creates PluginEnableOK with default headers values
func NewPluginEnableOK() *PluginEnableOK {

	return &PluginEnableOK{}
}

// WriteResponse to the client
func (o *PluginEnableOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// PluginEnableNotFoundCode is the HTTP code returned for type PluginEnableNotFound
const PluginEnableNotFoundCode int = 404

/*PluginEnableNotFound plugin is not installed

swagger:response pluginEnableNotFound
*/
type PluginEnableNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPluginEnableNotFound creates PluginEnableNotFound with default headers values
func NewPluginEnableNotFound() *PluginEnableNotFound {

	return &PluginEnableNotFound{}
}

// WithPayload adds the payload to the plugin enable not found response
func (o *PluginEnableNotFound) WithPayload(payload *models.ErrorResponse) *PluginEnableNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin enable not found response
func (o *PluginEnableNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginEnableNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginEnableInternalServerErrorCode is the HTTP code returned for type PluginEnableInternalServerError
const PluginEnableInternalServerErrorCode int = 500

/*PluginEnableInternalServerError server error

swagger:response pluginEnableInternalServerError
*/
type PluginEnableInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPluginEnableInternalServerError creates PluginEnableInternalServerError with default headers values
func NewPluginEnableInternalServerError() *PluginEnableInternalServerError {

	return &PluginEnableInternalServerError{}
}

// WithPayload adds the payload to the plugin enable internal server error response
func (o *PluginEnableInternalServerError) WithPayload(payload *models.ErrorResponse) *PluginEnableInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin enable internal server error response
func (o *PluginEnableInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginEnableInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}