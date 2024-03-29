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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewNodeInspectParams creates a new NodeInspectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeInspectParams() *NodeInspectParams {
	return &NodeInspectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeInspectParamsWithTimeout creates a new NodeInspectParams object
// with the ability to set a timeout on a request.
func NewNodeInspectParamsWithTimeout(timeout time.Duration) *NodeInspectParams {
	return &NodeInspectParams{
		timeout: timeout,
	}
}

// NewNodeInspectParamsWithContext creates a new NodeInspectParams object
// with the ability to set a context for a request.
func NewNodeInspectParamsWithContext(ctx context.Context) *NodeInspectParams {
	return &NodeInspectParams{
		Context: ctx,
	}
}

// NewNodeInspectParamsWithHTTPClient creates a new NodeInspectParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeInspectParamsWithHTTPClient(client *http.Client) *NodeInspectParams {
	return &NodeInspectParams{
		HTTPClient: client,
	}
}

/* NodeInspectParams contains all the parameters to send to the API endpoint
   for the node inspect operation.

   Typically these are written to a http.Request.
*/
type NodeInspectParams struct {

	/* ID.

	   The ID or name of the node
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeInspectParams) WithDefaults() *NodeInspectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node inspect params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeInspectParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the node inspect params
func (o *NodeInspectParams) WithTimeout(timeout time.Duration) *NodeInspectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node inspect params
func (o *NodeInspectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node inspect params
func (o *NodeInspectParams) WithContext(ctx context.Context) *NodeInspectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node inspect params
func (o *NodeInspectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node inspect params
func (o *NodeInspectParams) WithHTTPClient(client *http.Client) *NodeInspectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node inspect params
func (o *NodeInspectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the node inspect params
func (o *NodeInspectParams) WithID(id string) *NodeInspectParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the node inspect params
func (o *NodeInspectParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NodeInspectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
