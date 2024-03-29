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

package distribution

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// DistributionInspectOKCode is the HTTP code returned for type DistributionInspectOK
const DistributionInspectOKCode int = 200

/*DistributionInspectOK descriptor and platform information

swagger:response distributionInspectOK
*/
type DistributionInspectOK struct {

	/*
	  In: Body
	*/
	Payload *models.DistributionInspect `json:"body,omitempty"`
}

// NewDistributionInspectOK creates DistributionInspectOK with default headers values
func NewDistributionInspectOK() *DistributionInspectOK {

	return &DistributionInspectOK{}
}

// WithPayload adds the payload to the distribution inspect o k response
func (o *DistributionInspectOK) WithPayload(payload *models.DistributionInspect) *DistributionInspectOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the distribution inspect o k response
func (o *DistributionInspectOK) SetPayload(payload *models.DistributionInspect) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DistributionInspectOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DistributionInspectUnauthorizedCode is the HTTP code returned for type DistributionInspectUnauthorized
const DistributionInspectUnauthorizedCode int = 401

/*DistributionInspectUnauthorized Failed authentication or no image found

swagger:response distributionInspectUnauthorized
*/
type DistributionInspectUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDistributionInspectUnauthorized creates DistributionInspectUnauthorized with default headers values
func NewDistributionInspectUnauthorized() *DistributionInspectUnauthorized {

	return &DistributionInspectUnauthorized{}
}

// WithPayload adds the payload to the distribution inspect unauthorized response
func (o *DistributionInspectUnauthorized) WithPayload(payload *models.ErrorResponse) *DistributionInspectUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the distribution inspect unauthorized response
func (o *DistributionInspectUnauthorized) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DistributionInspectUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DistributionInspectInternalServerErrorCode is the HTTP code returned for type DistributionInspectInternalServerError
const DistributionInspectInternalServerErrorCode int = 500

/*DistributionInspectInternalServerError Server error

swagger:response distributionInspectInternalServerError
*/
type DistributionInspectInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewDistributionInspectInternalServerError creates DistributionInspectInternalServerError with default headers values
func NewDistributionInspectInternalServerError() *DistributionInspectInternalServerError {

	return &DistributionInspectInternalServerError{}
}

// WithPayload adds the payload to the distribution inspect internal server error response
func (o *DistributionInspectInternalServerError) WithPayload(payload *models.ErrorResponse) *DistributionInspectInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the distribution inspect internal server error response
func (o *DistributionInspectInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DistributionInspectInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
