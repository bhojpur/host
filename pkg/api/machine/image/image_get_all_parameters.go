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
	"github.com/go-openapi/swag"
)

// NewImageGetAllParams creates a new ImageGetAllParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageGetAllParams() *ImageGetAllParams {
	return &ImageGetAllParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageGetAllParamsWithTimeout creates a new ImageGetAllParams object
// with the ability to set a timeout on a request.
func NewImageGetAllParamsWithTimeout(timeout time.Duration) *ImageGetAllParams {
	return &ImageGetAllParams{
		timeout: timeout,
	}
}

// NewImageGetAllParamsWithContext creates a new ImageGetAllParams object
// with the ability to set a context for a request.
func NewImageGetAllParamsWithContext(ctx context.Context) *ImageGetAllParams {
	return &ImageGetAllParams{
		Context: ctx,
	}
}

// NewImageGetAllParamsWithHTTPClient creates a new ImageGetAllParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageGetAllParamsWithHTTPClient(client *http.Client) *ImageGetAllParams {
	return &ImageGetAllParams{
		HTTPClient: client,
	}
}

/* ImageGetAllParams contains all the parameters to send to the API endpoint
   for the image get all operation.

   Typically these are written to a http.Request.
*/
type ImageGetAllParams struct {

	/* Names.

	   Image names to filter by
	*/
	Names []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageGetAllParams) WithDefaults() *ImageGetAllParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image get all params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageGetAllParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image get all params
func (o *ImageGetAllParams) WithTimeout(timeout time.Duration) *ImageGetAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image get all params
func (o *ImageGetAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image get all params
func (o *ImageGetAllParams) WithContext(ctx context.Context) *ImageGetAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image get all params
func (o *ImageGetAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image get all params
func (o *ImageGetAllParams) WithHTTPClient(client *http.Client) *ImageGetAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image get all params
func (o *ImageGetAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNames adds the names to the image get all params
func (o *ImageGetAllParams) WithNames(names []string) *ImageGetAllParams {
	o.SetNames(names)
	return o
}

// SetNames adds the names to the image get all params
func (o *ImageGetAllParams) SetNames(names []string) {
	o.Names = names
}

// WriteToRequest writes these params to a swagger request
func (o *ImageGetAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Names != nil {

		// binding items for names
		joinedNames := o.bindParamNames(reg)

		// query array param names
		if err := r.SetQueryParam("names", joinedNames...); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamImageGetAll binds the parameter names
func (o *ImageGetAllParams) bindParamNames(formats strfmt.Registry) []string {
	namesIR := o.Names

	var namesIC []string
	for _, namesIIR := range namesIR { // explode []string

		namesIIV := namesIIR // string as string
		namesIC = append(namesIC, namesIIV)
	}

	// items.CollectionFormat: ""
	namesIS := swag.JoinByFormat(namesIC, "")

	return namesIS
}
