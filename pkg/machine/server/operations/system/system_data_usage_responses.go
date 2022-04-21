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

// SystemDataUsageOKCode is the HTTP code returned for type SystemDataUsageOK
const SystemDataUsageOKCode int = 200

/*SystemDataUsageOK no error

swagger:response systemDataUsageOK
*/
type SystemDataUsageOK struct {

	/*
	  In: Body
	*/
	Payload *SystemDataUsageOKBody `json:"body,omitempty"`
}

// NewSystemDataUsageOK creates SystemDataUsageOK with default headers values
func NewSystemDataUsageOK() *SystemDataUsageOK {

	return &SystemDataUsageOK{}
}

// WithPayload adds the payload to the system data usage o k response
func (o *SystemDataUsageOK) WithPayload(payload *SystemDataUsageOKBody) *SystemDataUsageOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the system data usage o k response
func (o *SystemDataUsageOK) SetPayload(payload *SystemDataUsageOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SystemDataUsageOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SystemDataUsageInternalServerErrorCode is the HTTP code returned for type SystemDataUsageInternalServerError
const SystemDataUsageInternalServerErrorCode int = 500

/*SystemDataUsageInternalServerError server error

swagger:response systemDataUsageInternalServerError
*/
type SystemDataUsageInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSystemDataUsageInternalServerError creates SystemDataUsageInternalServerError with default headers values
func NewSystemDataUsageInternalServerError() *SystemDataUsageInternalServerError {

	return &SystemDataUsageInternalServerError{}
}

// WithPayload adds the payload to the system data usage internal server error response
func (o *SystemDataUsageInternalServerError) WithPayload(payload *models.ErrorResponse) *SystemDataUsageInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the system data usage internal server error response
func (o *SystemDataUsageInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SystemDataUsageInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
