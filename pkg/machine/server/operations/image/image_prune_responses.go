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

// ImagePruneOKCode is the HTTP code returned for type ImagePruneOK
const ImagePruneOKCode int = 200

/*ImagePruneOK No error

swagger:response imagePruneOK
*/
type ImagePruneOK struct {

	/*
	  In: Body
	*/
	Payload *ImagePruneOKBody `json:"body,omitempty"`
}

// NewImagePruneOK creates ImagePruneOK with default headers values
func NewImagePruneOK() *ImagePruneOK {

	return &ImagePruneOK{}
}

// WithPayload adds the payload to the image prune o k response
func (o *ImagePruneOK) WithPayload(payload *ImagePruneOKBody) *ImagePruneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image prune o k response
func (o *ImagePruneOK) SetPayload(payload *ImagePruneOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImagePruneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ImagePruneInternalServerErrorCode is the HTTP code returned for type ImagePruneInternalServerError
const ImagePruneInternalServerErrorCode int = 500

/*ImagePruneInternalServerError Server error

swagger:response imagePruneInternalServerError
*/
type ImagePruneInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewImagePruneInternalServerError creates ImagePruneInternalServerError with default headers values
func NewImagePruneInternalServerError() *ImagePruneInternalServerError {

	return &ImagePruneInternalServerError{}
}

// WithPayload adds the payload to the image prune internal server error response
func (o *ImagePruneInternalServerError) WithPayload(payload *models.ErrorResponse) *ImagePruneInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the image prune internal server error response
func (o *ImagePruneInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ImagePruneInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
