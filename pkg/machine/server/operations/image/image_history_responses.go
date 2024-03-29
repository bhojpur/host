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

// ImageHistoryOKCode is the HTTP code returned for type ImageHistoryOK
const ImageHistoryOKCode int = 200

/*ImageHistoryOK List of image layers

swagger:response imageHistoryOK
*/
type ImageHistoryOK struct {

	/*
	  In: Body
	*/
	Payload []*HistoryResponseItem `json:"body,omitempty"`
}

// NewImageHistoryOK creates ImageHistoryOK with default headers values
func NewImageHistoryOK() *ImageHistoryOK {

	return &ImageHistoryOK{}
}

// WithPayload adds the payload to the image history o k response
func (o *ImageHistoryOK) WithPayload(payload []*HistoryResponseItem) *ImageHistoryOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image history o k response
func (o *ImageHistoryOK) SetPayload(payload []*HistoryResponseItem) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageHistoryOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*HistoryResponseItem, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// ImageHistoryNotFoundCode is the HTTP code returned for type ImageHistoryNotFound
const ImageHistoryNotFoundCode int = 404

/*ImageHistoryNotFound No such image

swagger:response imageHistoryNotFound
*/
type ImageHistoryNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageHistoryNotFound creates ImageHistoryNotFound with default headers values
func NewImageHistoryNotFound() *ImageHistoryNotFound {

	return &ImageHistoryNotFound{}
}

// WithPayload adds the payload to the image history not found response
func (o *ImageHistoryNotFound) WithPayload(payload *models.ErrorResponse) *ImageHistoryNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image history not found response
func (o *ImageHistoryNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageHistoryNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImageHistoryInternalServerErrorCode is the HTTP code returned for type ImageHistoryInternalServerError
const ImageHistoryInternalServerErrorCode int = 500

/*ImageHistoryInternalServerError Server error

swagger:response imageHistoryInternalServerError
*/
type ImageHistoryInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImageHistoryInternalServerError creates ImageHistoryInternalServerError with default headers values
func NewImageHistoryInternalServerError() *ImageHistoryInternalServerError {

	return &ImageHistoryInternalServerError{}
}

// WithPayload adds the payload to the image history internal server error response
func (o *ImageHistoryInternalServerError) WithPayload(payload *models.ErrorResponse) *ImageHistoryInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image history internal server error response
func (o *ImageHistoryInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImageHistoryInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
