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

package container

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ContainerArchiveInfoOKCode is the HTTP code returned for type ContainerArchiveInfoOK
const ContainerArchiveInfoOKCode int = 200

/*ContainerArchiveInfoOK no error

swagger:response containerArchiveInfoOK
*/
type ContainerArchiveInfoOK struct {
	/*A base64 - encoded JSON object with some filesystem header
	information about the path


	*/
	XBhojpurContainerPathStat string `json:"X-Bhojpur-Container-Path-Stat"`
}

// NewContainerArchiveInfoOK creates ContainerArchiveInfoOK with default headers values
func NewContainerArchiveInfoOK() *ContainerArchiveInfoOK {

	return &ContainerArchiveInfoOK{}
}

// WithXBhojpurContainerPathStat adds the xBhojpurContainerPathStat to the container archive info o k response
func (o *ContainerArchiveInfoOK) WithXBhojpurContainerPathStat(xBhojpurContainerPathStat string) *ContainerArchiveInfoOK {
	o.XBhojpurContainerPathStat = xBhojpurContainerPathStat
	return o
}

// SetXBhojpurContainerPathStat sets the xBhojpurContainerPathStat to the container archive info o k response
func (o *ContainerArchiveInfoOK) SetXBhojpurContainerPathStat(xBhojpurContainerPathStat string) {
	o.XBhojpurContainerPathStat = xBhojpurContainerPathStat
}

// WriteResponse to the client
func (o *ContainerArchiveInfoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	// response header X-Bhojpur-Container-Path-Stat

	xBhojpurContainerPathStat := o.XBhojpurContainerPathStat
	if xBhojpurContainerPathStat != "" {
		rw.Header().Set("X-Bhojpur-Container-Path-Stat", xBhojpurContainerPathStat)
	}

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// ContainerArchiveInfoBadRequestCode is the HTTP code returned for type ContainerArchiveInfoBadRequest
const ContainerArchiveInfoBadRequestCode int = 400

/*ContainerArchiveInfoBadRequest Bad parameter

swagger:response containerArchiveInfoBadRequest
*/
type ContainerArchiveInfoBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerArchiveInfoBadRequest creates ContainerArchiveInfoBadRequest with default headers values
func NewContainerArchiveInfoBadRequest() *ContainerArchiveInfoBadRequest {

	return &ContainerArchiveInfoBadRequest{}
}

// WithPayload adds the payload to the container archive info bad request response
func (o *ContainerArchiveInfoBadRequest) WithPayload(payload *models.ErrorResponse) *ContainerArchiveInfoBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container archive info bad request response
func (o *ContainerArchiveInfoBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerArchiveInfoBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerArchiveInfoNotFoundCode is the HTTP code returned for type ContainerArchiveInfoNotFound
const ContainerArchiveInfoNotFoundCode int = 404

/*ContainerArchiveInfoNotFound Container or path does not exist

swagger:response containerArchiveInfoNotFound
*/
type ContainerArchiveInfoNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerArchiveInfoNotFound creates ContainerArchiveInfoNotFound with default headers values
func NewContainerArchiveInfoNotFound() *ContainerArchiveInfoNotFound {

	return &ContainerArchiveInfoNotFound{}
}

// WithPayload adds the payload to the container archive info not found response
func (o *ContainerArchiveInfoNotFound) WithPayload(payload *models.ErrorResponse) *ContainerArchiveInfoNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container archive info not found response
func (o *ContainerArchiveInfoNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerArchiveInfoNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ContainerArchiveInfoInternalServerErrorCode is the HTTP code returned for type ContainerArchiveInfoInternalServerError
const ContainerArchiveInfoInternalServerErrorCode int = 500

/*ContainerArchiveInfoInternalServerError Server error

swagger:response containerArchiveInfoInternalServerError
*/
type ContainerArchiveInfoInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewContainerArchiveInfoInternalServerError creates ContainerArchiveInfoInternalServerError with default headers values
func NewContainerArchiveInfoInternalServerError() *ContainerArchiveInfoInternalServerError {

	return &ContainerArchiveInfoInternalServerError{}
}

// WithPayload adds the payload to the container archive info internal server error response
func (o *ContainerArchiveInfoInternalServerError) WithPayload(payload *models.ErrorResponse) *ContainerArchiveInfoInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container archive info internal server error response
func (o *ContainerArchiveInfoInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerArchiveInfoInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}