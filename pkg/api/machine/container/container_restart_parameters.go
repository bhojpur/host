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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewContainerRestartParams creates a new ContainerRestartParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerRestartParams() *ContainerRestartParams {
	return &ContainerRestartParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerRestartParamsWithTimeout creates a new ContainerRestartParams object
// with the ability to set a timeout on a request.
func NewContainerRestartParamsWithTimeout(timeout time.Duration) *ContainerRestartParams {
	return &ContainerRestartParams{
		timeout: timeout,
	}
}

// NewContainerRestartParamsWithContext creates a new ContainerRestartParams object
// with the ability to set a context for a request.
func NewContainerRestartParamsWithContext(ctx context.Context) *ContainerRestartParams {
	return &ContainerRestartParams{
		Context: ctx,
	}
}

// NewContainerRestartParamsWithHTTPClient creates a new ContainerRestartParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerRestartParamsWithHTTPClient(client *http.Client) *ContainerRestartParams {
	return &ContainerRestartParams{
		HTTPClient: client,
	}
}

/* ContainerRestartParams contains all the parameters to send to the API endpoint
   for the container restart operation.

   Typically these are written to a http.Request.
*/
type ContainerRestartParams struct {

	/* ID.

	   ID or name of the container
	*/
	ID string

	/* T.

	   Number of seconds to wait before killing the container
	*/
	T *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container restart params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerRestartParams) WithDefaults() *ContainerRestartParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container restart params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerRestartParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the container restart params
func (o *ContainerRestartParams) WithTimeout(timeout time.Duration) *ContainerRestartParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container restart params
func (o *ContainerRestartParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container restart params
func (o *ContainerRestartParams) WithContext(ctx context.Context) *ContainerRestartParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container restart params
func (o *ContainerRestartParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container restart params
func (o *ContainerRestartParams) WithHTTPClient(client *http.Client) *ContainerRestartParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container restart params
func (o *ContainerRestartParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the container restart params
func (o *ContainerRestartParams) WithID(id string) *ContainerRestartParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container restart params
func (o *ContainerRestartParams) SetID(id string) {
	o.ID = id
}

// WithT adds the t to the container restart params
func (o *ContainerRestartParams) WithT(t *int64) *ContainerRestartParams {
	o.SetT(t)
	return o
}

// SetT adds the t to the container restart params
func (o *ContainerRestartParams) SetT(t *int64) {
	o.T = t
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerRestartParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.T != nil {

		// query param t
		var qrT int64

		if o.T != nil {
			qrT = *o.T
		}
		qT := swag.FormatInt64(qrT)
		if qT != "" {

			if err := r.SetQueryParam("t", qT); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
