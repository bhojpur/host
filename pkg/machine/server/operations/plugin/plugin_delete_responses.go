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

// PluginDeleteOKCode is the HTTP code returned for type PluginDeleteOK
const PluginDeleteOKCode int = 200

/*PluginDeleteOK no error

swagger:response pluginDeleteOK
*/
type PluginDeleteOK struct {

	/*
	  In: Body
	*/
	Payload *models.Plugin `json:"body,omitempty"`
}

// NewPluginDeleteOK creates PluginDeleteOK with default headers values
func NewPluginDeleteOK() *PluginDeleteOK {

	return &PluginDeleteOK{}
}

// WithPayload adds the payload to the plugin delete o k response
func (o *PluginDeleteOK) WithPayload(payload *models.Plugin) *PluginDeleteOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin delete o k response
func (o *PluginDeleteOK) SetPayload(payload *models.Plugin) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginDeleteOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginDeleteNotFoundCode is the HTTP code returned for type PluginDeleteNotFound
const PluginDeleteNotFoundCode int = 404

/*PluginDeleteNotFound plugin is not installed

swagger:response pluginDeleteNotFound
*/
type PluginDeleteNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPluginDeleteNotFound creates PluginDeleteNotFound with default headers values
func NewPluginDeleteNotFound() *PluginDeleteNotFound {

	return &PluginDeleteNotFound{}
}

// WithPayload adds the payload to the plugin delete not found response
func (o *PluginDeleteNotFound) WithPayload(payload *models.ErrorResponse) *PluginDeleteNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin delete not found response
func (o *PluginDeleteNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginDeleteNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PluginDeleteInternalServerErrorCode is the HTTP code returned for type PluginDeleteInternalServerError
const PluginDeleteInternalServerErrorCode int = 500

/*PluginDeleteInternalServerError server error

swagger:response pluginDeleteInternalServerError
*/
type PluginDeleteInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewPluginDeleteInternalServerError creates PluginDeleteInternalServerError with default headers values
func NewPluginDeleteInternalServerError() *PluginDeleteInternalServerError {

	return &PluginDeleteInternalServerError{}
}

// WithPayload adds the payload to the plugin delete internal server error response
func (o *PluginDeleteInternalServerError) WithPayload(payload *models.ErrorResponse) *PluginDeleteInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the plugin delete internal server error response
func (o *PluginDeleteInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PluginDeleteInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
