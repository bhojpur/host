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

package system

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

// NewSystemInfoParams creates a new SystemInfoParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSystemInfoParams() *SystemInfoParams {
	return &SystemInfoParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSystemInfoParamsWithTimeout creates a new SystemInfoParams object
// with the ability to set a timeout on a request.
func NewSystemInfoParamsWithTimeout(timeout time.Duration) *SystemInfoParams {
	return &SystemInfoParams{
		timeout: timeout,
	}
}

// NewSystemInfoParamsWithContext creates a new SystemInfoParams object
// with the ability to set a context for a request.
func NewSystemInfoParamsWithContext(ctx context.Context) *SystemInfoParams {
	return &SystemInfoParams{
		Context: ctx,
	}
}

// NewSystemInfoParamsWithHTTPClient creates a new SystemInfoParams object
// with the ability to set a custom HTTPClient for a request.
func NewSystemInfoParamsWithHTTPClient(client *http.Client) *SystemInfoParams {
	return &SystemInfoParams{
		HTTPClient: client,
	}
}

/* SystemInfoParams contains all the parameters to send to the API endpoint
   for the system info operation.

   Typically these are written to a http.Request.
*/
type SystemInfoParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the system info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemInfoParams) WithDefaults() *SystemInfoParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the system info params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SystemInfoParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the system info params
func (o *SystemInfoParams) WithTimeout(timeout time.Duration) *SystemInfoParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the system info params
func (o *SystemInfoParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the system info params
func (o *SystemInfoParams) WithContext(ctx context.Context) *SystemInfoParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the system info params
func (o *SystemInfoParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the system info params
func (o *SystemInfoParams) WithHTTPClient(client *http.Client) *SystemInfoParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the system info params
func (o *SystemInfoParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *SystemInfoParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
