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

// SystemAuthOKCode is the HTTP code returned for type SystemAuthOK
const SystemAuthOKCode int = 200

/*SystemAuthOK An identity token was generated successfully.

swagger:response systemAuthOK
*/
type SystemAuthOK struct {

	/*
	  In: Body
	*/
	Payload *SystemAuthOKBody `json:"body,omitempty"`
}

// NewSystemAuthOK creates SystemAuthOK with default headers values
func NewSystemAuthOK() *SystemAuthOK {

	return &SystemAuthOK{}
}

// WithPayload adds the payload to the system auth o k response
func (o *SystemAuthOK) WithPayload(payload *SystemAuthOKBody) *SystemAuthOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the system auth o k response
func (o *SystemAuthOK) SetPayload(payload *SystemAuthOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SystemAuthOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// SystemAuthNoContentCode is the HTTP code returned for type SystemAuthNoContent
const SystemAuthNoContentCode int = 204

/*SystemAuthNoContent No error

swagger:response systemAuthNoContent
*/
type SystemAuthNoContent struct {
}

// NewSystemAuthNoContent creates SystemAuthNoContent with default headers values
func NewSystemAuthNoContent() *SystemAuthNoContent {

	return &SystemAuthNoContent{}
}

// WriteResponse to the client
func (o *SystemAuthNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

// SystemAuthInternalServerErrorCode is the HTTP code returned for type SystemAuthInternalServerError
const SystemAuthInternalServerErrorCode int = 500

/*SystemAuthInternalServerError Server error

swagger:response systemAuthInternalServerError
*/
type SystemAuthInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewSystemAuthInternalServerError creates SystemAuthInternalServerError with default headers values
func NewSystemAuthInternalServerError() *SystemAuthInternalServerError {

	return &SystemAuthInternalServerError{}
}

// WithPayload adds the payload to the system auth internal server error response
func (o *SystemAuthInternalServerError) WithPayload(payload *models.ErrorResponse) *SystemAuthInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the system auth internal server error response
func (o *SystemAuthInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *SystemAuthInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
