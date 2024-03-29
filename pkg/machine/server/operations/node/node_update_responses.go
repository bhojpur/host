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

package node

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/bhojpur/host/pkg/machine/models"
)

// NodeUpdateOKCode is the HTTP code returned for type NodeUpdateOK
const NodeUpdateOKCode int = 200

/*NodeUpdateOK no error

swagger:response nodeUpdateOK
*/
type NodeUpdateOK struct {
}

// NewNodeUpdateOK creates NodeUpdateOK with default headers values
func NewNodeUpdateOK() *NodeUpdateOK {

	return &NodeUpdateOK{}
}

// WriteResponse to the client
func (o *NodeUpdateOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// NodeUpdateBadRequestCode is the HTTP code returned for type NodeUpdateBadRequest
const NodeUpdateBadRequestCode int = 400

/*NodeUpdateBadRequest bad parameter

swagger:response nodeUpdateBadRequest
*/
type NodeUpdateBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodeUpdateBadRequest creates NodeUpdateBadRequest with default headers values
func NewNodeUpdateBadRequest() *NodeUpdateBadRequest {

	return &NodeUpdateBadRequest{}
}

// WithPayload adds the payload to the node update bad request response
func (o *NodeUpdateBadRequest) WithPayload(payload *models.ErrorResponse) *NodeUpdateBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node update bad request response
func (o *NodeUpdateBadRequest) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeUpdateBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodeUpdateNotFoundCode is the HTTP code returned for type NodeUpdateNotFound
const NodeUpdateNotFoundCode int = 404

/*NodeUpdateNotFound no such node

swagger:response nodeUpdateNotFound
*/
type NodeUpdateNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodeUpdateNotFound creates NodeUpdateNotFound with default headers values
func NewNodeUpdateNotFound() *NodeUpdateNotFound {

	return &NodeUpdateNotFound{}
}

// WithPayload adds the payload to the node update not found response
func (o *NodeUpdateNotFound) WithPayload(payload *models.ErrorResponse) *NodeUpdateNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node update not found response
func (o *NodeUpdateNotFound) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeUpdateNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodeUpdateInternalServerErrorCode is the HTTP code returned for type NodeUpdateInternalServerError
const NodeUpdateInternalServerErrorCode int = 500

/*NodeUpdateInternalServerError server error

swagger:response nodeUpdateInternalServerError
*/
type NodeUpdateInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodeUpdateInternalServerError creates NodeUpdateInternalServerError with default headers values
func NewNodeUpdateInternalServerError() *NodeUpdateInternalServerError {

	return &NodeUpdateInternalServerError{}
}

// WithPayload adds the payload to the node update internal server error response
func (o *NodeUpdateInternalServerError) WithPayload(payload *models.ErrorResponse) *NodeUpdateInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node update internal server error response
func (o *NodeUpdateInternalServerError) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeUpdateInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// NodeUpdateServiceUnavailableCode is the HTTP code returned for type NodeUpdateServiceUnavailable
const NodeUpdateServiceUnavailableCode int = 503

/*NodeUpdateServiceUnavailable node is not part of a swarm

swagger:response nodeUpdateServiceUnavailable
*/
type NodeUpdateServiceUnavailable struct {

	/*
	  In: Body
	*/
	Payload *models.ErrorResponse `json:"body,omitempty"`
}

// NewNodeUpdateServiceUnavailable creates NodeUpdateServiceUnavailable with default headers values
func NewNodeUpdateServiceUnavailable() *NodeUpdateServiceUnavailable {

	return &NodeUpdateServiceUnavailable{}
}

// WithPayload adds the payload to the node update service unavailable response
func (o *NodeUpdateServiceUnavailable) WithPayload(payload *models.ErrorResponse) *NodeUpdateServiceUnavailable {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the node update service unavailable response
func (o *NodeUpdateServiceUnavailable) SetPayload(payload *models.ErrorResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *NodeUpdateServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
