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

package image

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ImageTagCreatedCode is the HTTP code returned for type ImageTagCreated
const ImageTagCreatedCode int = 201

/*ImageTagCreated No error

swagger:response imageTagCreated
*/
type ImageTagCreated struct {
}

// NewImageTagCreated creates ImageTagCreated with default headers values
func NewImageTagCreated() *ImageTagCreated {

	return &ImageTagCreated{}
}

// WriteResponse to the client
func (o *ImageTagCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

// ImageTagBadRequestCode is the HTTP code returned for type ImageTagBadRequest
const ImageTagBadRequestCode int = 400

/*ImageTagBadRequest Bad parameter

swagger:response imageTagBadRequest
*/
type ImageTagBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageTagBadRequest creates ImageTagBadRequest with default headers values
func NewImageTagBadRequest() *ImageTagBadRequest {

	return &ImageTagBadRequest{}
}

// WithPayload adds the payload to the image tag bad request response
func (o *ImageTagBadRequest) WithPayload(payload *models.ErrorResponse) *ImageTagBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image tag bad request response
func (o *ImageTagBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageTagBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageTagNotFoundCode is the HTTP code returned for type ImageTagNotFound
const ImageTagNotFoundCode int = 404

/*ImageTagNotFound No such image

swagger:response imageTagNotFound
*/
type ImageTagNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageTagNotFound creates ImageTagNotFound with default headers values
func NewImageTagNotFound() *ImageTagNotFound {

	return &ImageTagNotFound{}
}

// WithPayload adds the payload to the image tag not found response
func (o *ImageTagNotFound) WithPayload(payload *models.ErrorResponse) *ImageTagNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image tag not found response
func (o *ImageTagNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageTagNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageTagConflictCode is the HTTP code returned for type ImageTagConflict
const ImageTagConflictCode int = 409

/*ImageTagConflict Conflict

swagger:response imageTagConflict
*/
type ImageTagConflict struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageTagConflict creates ImageTagConflict with default headers values
func NewImageTagConflict() *ImageTagConflict {

	return &ImageTagConflict{}
}

// WithPayload adds the payload to the image tag conflict response
func (o *ImageTagConflict) WithPayload(payload *models.ErrorResponse) *ImageTagConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image tag conflict response
func (o *ImageTagConflict) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageTagConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageTagInternalServerErrorCode is the HTTP code returned for type ImageTagInternalServerError
const ImageTagInternalServerErrorCode int = 500

/*ImageTagInternalServerError Server error

swagger:response imageTagInternalServerError
*/
type ImageTagInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageTagInternalServerError creates ImageTagInternalServerError with default headers values
func NewImageTagInternalServerError() *ImageTagInternalServerError {

	return &ImageTagInternalServerError{}
}

// WithPayload adds the payload to the image tag internal server error response
func (o *ImageTagInternalServerError) WithPayload(payload *models.ErrorResponse) *ImageTagInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image tag internal server error response
func (o *ImageTagInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageTagInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
