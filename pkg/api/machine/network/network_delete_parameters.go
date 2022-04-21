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

package network

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

// NewNetworkDeleteParams creates a new NetworkDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkDeleteParams() *NetworkDeleteParams {
	return &NetworkDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkDeleteParamsWithTimeout creates a new NetworkDeleteParams object
// with the ability to set a timeout on a request.
func NewNetworkDeleteParamsWithTimeout(timeout time.Duration) *NetworkDeleteParams {
	return &NetworkDeleteParams{
		timeout: timeout,
	}
}

// NewNetworkDeleteParamsWithContext creates a new NetworkDeleteParams object
// with the ability to set a context for a request.
func NewNetworkDeleteParamsWithContext(ctx context.Context) *NetworkDeleteParams {
	return &NetworkDeleteParams{
		Context: ctx,
	}
}

// NewNetworkDeleteParamsWithHTTPClient creates a new NetworkDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkDeleteParamsWithHTTPClient(client *http.Client) *NetworkDeleteParams {
	return &NetworkDeleteParams{
		HTTPClient: client,
	}
}

/* NetworkDeleteParams contains all the parameters to send to the API endpoint
   for the network delete operation.

   Typically these are written to a http.Request.
*/
type NetworkDeleteParams struct {

	/* ID.

	   Network ID or name
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkDeleteParams) WithDefaults() *NetworkDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network delete params
func (o *NetworkDeleteParams) WithTimeout(timeout time.Duration) *NetworkDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network delete params
func (o *NetworkDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network delete params
func (o *NetworkDeleteParams) WithContext(ctx context.Context) *NetworkDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network delete params
func (o *NetworkDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network delete params
func (o *NetworkDeleteParams) WithHTTPClient(client *http.Client) *NetworkDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network delete params
func (o *NetworkDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the network delete params
func (o *NetworkDeleteParams) WithID(id string) *NetworkDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the network delete params
func (o *NetworkDeleteParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
