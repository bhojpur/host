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

package exec

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ExecInspectOKCode is the HTTP code returned for type ExecInspectOK
const ExecInspectOKCode int = 200

/*ExecInspectOK No error

swagger:response execInspectOK
*/
type ExecInspectOK struct {

	/*
	  In: Body
	*/
	Payload *ExecInspectOKBody `json:"body,omitempty"`
}

// NewExecInspectOK creates ExecInspectOK with default headers values
func NewExecInspectOK() *ExecInspectOK {

	return &ExecInspectOK{}
}

// WithPayload adds the payload to the exec inspect o k response
func (o *ExecInspectOK) WithPayload(payload *ExecInspectOKBody) *ExecInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the exec inspect o k response
func (o *ExecInspectOK) SetPayload(payload *ExecInspectOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExecInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ExecInspectNotFoundCode is the HTTP code returned for type ExecInspectNotFound
const ExecInspectNotFoundCode int = 404

/*ExecInspectNotFound No such exec instance

swagger:response execInspectNotFound
*/
type ExecInspectNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewExecInspectNotFound creates ExecInspectNotFound with default headers values
func NewExecInspectNotFound() *ExecInspectNotFound {

	return &ExecInspectNotFound{}
}

// WithPayload adds the payload to the exec inspect not found response
func (o *ExecInspectNotFound) WithPayload(payload *models.ErrorResponse) *ExecInspectNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the exec inspect not found response
func (o *ExecInspectNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExecInspectNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ExecInspectInternalServerErrorCode is the HTTP code returned for type ExecInspectInternalServerError
const ExecInspectInternalServerErrorCode int = 500

/*ExecInspectInternalServerError Server error

swagger:response execInspectInternalServerError
*/
type ExecInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewExecInspectInternalServerError creates ExecInspectInternalServerError with default headers values
func NewExecInspectInternalServerError() *ExecInspectInternalServerError {

	return &ExecInspectInternalServerError{}
}

// WithPayload adds the payload to the exec inspect internal server error response
func (o *ExecInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *ExecInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the exec inspect internal server error response
func (o *ExecInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ExecInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}