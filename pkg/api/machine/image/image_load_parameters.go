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
	"io"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewImageLoadParams creates a new ImageLoadParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageLoadParams() *ImageLoadParams {
	return &ImageLoadParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageLoadParamsWithTimeout creates a new ImageLoadParams object
// with the ability to set a timeout on a request.
func NewImageLoadParamsWithTimeout(timeout time.Duration) *ImageLoadParams {
	return &ImageLoadParams{
		timeout: timeout,
	}
}

// NewImageLoadParamsWithContext creates a new ImageLoadParams object
// with the ability to set a context for a request.
func NewImageLoadParamsWithContext(ctx context.Context) *ImageLoadParams {
	return &ImageLoadParams{
		Context: ctx,
	}
}

// NewImageLoadParamsWithHTTPClient creates a new ImageLoadParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageLoadParamsWithHTTPClient(client *http.Client) *ImageLoadParams {
	return &ImageLoadParams{
		HTTPClient: client,
	}
}

/* ImageLoadParams contains all the parameters to send to the API endpoint
   for the image load operation.

   Typically these are written to a http.Request.
*/
type ImageLoadParams struct {

	/* ImagesTarball.

	   Tar archive containing images

	   Format: binary
	*/
	ImagesTarball io.ReadCloser

	/* Quiet.

	   Suppress progress details during load.
	*/
	Quiet *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image load params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageLoadParams) WithDefaults() *ImageLoadParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image load params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageLoadParams) SetDefaults() {
	var (
		quietDefault = bool(false)
	)

	val := ImageLoadParams{
		Quiet: &quietDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the image load params
func (o *ImageLoadParams) WithTimeout(timeout time.Duration) *ImageLoadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image load params
func (o *ImageLoadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image load params
func (o *ImageLoadParams) WithContext(ctx context.Context) *ImageLoadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image load params
func (o *ImageLoadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image load params
func (o *ImageLoadParams) WithHTTPClient(client *http.Client) *ImageLoadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image load params
func (o *ImageLoadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithImagesTarball adds the imagesTarball to the image load params
func (o *ImageLoadParams) WithImagesTarball(imagesTarball io.ReadCloser) *ImageLoadParams {
	o.SetImagesTarball(imagesTarball)
	return o
}

// SetImagesTarball adds the imagesTarball to the image load params
func (o *ImageLoadParams) SetImagesTarball(imagesTarball io.ReadCloser) {
	o.ImagesTarball = imagesTarball
}

// WithQuiet adds the quiet to the image load params
func (o *ImageLoadParams) WithQuiet(quiet *bool) *ImageLoadParams {
	o.SetQuiet(quiet)
	return o
}

// SetQuiet adds the quiet to the image load params
func (o *ImageLoadParams) SetQuiet(quiet *bool) {
	o.Quiet = quiet
}

// WriteToRequest writes these params to a swagger request
func (o *ImageLoadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.ImagesTarball != nil {
		if err := r.SetBodyParam(o.ImagesTarball); err != nil {
			return err
		}
	}

	if o.Quiet != nil {

		// query param quiet
		var qrQuiet bool

		if o.Quiet != nil {
			qrQuiet = *o.Quiet
		}
		qQuiet := swag.FormatBool(qrQuiet)
		if qQuiet != "" {

			if err := r.SetQueryParam("quiet", qQuiet); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
