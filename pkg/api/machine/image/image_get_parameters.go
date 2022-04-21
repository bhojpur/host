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
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewImageGetParams creates a new ImageGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageGetParams() *ImageGetParams {
	return &ImageGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageGetParamsWithTimeout creates a new ImageGetParams object
// with the ability to set a timeout on a request.
func NewImageGetParamsWithTimeout(timeout time.Duration) *ImageGetParams {
	return &ImageGetParams{
		timeout: timeout,
	}
}

// NewImageGetParamsWithContext creates a new ImageGetParams object
// with the ability to set a context for a request.
func NewImageGetParamsWithContext(ctx context.Context) *ImageGetParams {
	return &ImageGetParams{
		Context: ctx,
	}
}

// NewImageGetParamsWithHTTPClient creates a new ImageGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageGetParamsWithHTTPClient(client *http.Client) *ImageGetParams {
	return &ImageGetParams{
		HTTPClient: client,
	}
}

/* ImageGetParams contains all the parameters to send to the API endpoint
   for the image get operation.

   Typically these are written to a http.Request.
*/
type ImageGetParams struct {

	/* Name.

	   Image name or ID
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageGetParams) WithDefaults() *ImageGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image get params
func (o *ImageGetParams) WithTimeout(timeout time.Duration) *ImageGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image get params
func (o *ImageGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image get params
func (o *ImageGetParams) WithContext(ctx context.Context) *ImageGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image get params
func (o *ImageGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image get params
func (o *ImageGetParams) WithHTTPClient(client *http.Client) *ImageGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image get params
func (o *ImageGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the image get params
func (o *ImageGetParams) WithName(name string) *ImageGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the image get params
func (o *ImageGetParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *ImageGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
