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

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// SystemInfoOKCode is the HTTP code returned for type SystemInfoOK
const SystemInfoOKCode int = 200

/*SystemInfoOK No error

swagger:response systemInfoOK
*/
type SystemInfoOK struct {

	/*
	  In: Body
	*/
	Payload *models.SystemInfo `json:"body,omitempty"`
}

// NewSystemInfoOK creates SystemInfoOK with default headers values
func NewSystemInfoOK() *SystemInfoOK {

	return &SystemInfoOK{}
}

// WithPayload adds the payload to the system info o k response
func (o *SystemInfoOK) WithPayload(payload *models.SystemInfo) *SystemInfoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the system info o k response
func (o *SystemInfoOK) SetPayload(payload *models.SystemInfo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SystemInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SystemInfoInternalServerErrorCode is the HTTP code returned for type SystemInfoInternalServerError
const SystemInfoInternalServerErrorCode int = 500

/*SystemInfoInternalServerError Server error

swagger:response systemInfoInternalServerError
*/
type SystemInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSystemInfoInternalServerError creates SystemInfoInternalServerError with default headers values
func NewSystemInfoInternalServerError() *SystemInfoInternalServerError {

	return &SystemInfoInternalServerError{}
}

// WithPayload adds the payload to the system info internal server error response
func (o *SystemInfoInternalServerError) WithPayload(payload *models.ErrorResponse) *SystemInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the system info internal server error response
func (o *SystemInfoInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SystemInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}