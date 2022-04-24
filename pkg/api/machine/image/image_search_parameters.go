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

// NewImageSearchParams creates a new ImageSearchParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewImageSearchParams() *ImageSearchParams {
	return &ImageSearchParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewImageSearchParamsWithTimeout creates a new ImageSearchParams object
// with the ability to set a timeout on a request.
func NewImageSearchParamsWithTimeout(timeout time.Duration) *ImageSearchParams {
	return &ImageSearchParams{
		timeout: timeout,
	}
}

// NewImageSearchParamsWithContext creates a new ImageSearchParams object
// with the ability to set a context for a request.
func NewImageSearchParamsWithContext(ctx context.Context) *ImageSearchParams {
	return &ImageSearchParams{
		Context: ctx,
	}
}

// NewImageSearchParamsWithHTTPClient creates a new ImageSearchParams object
// with the ability to set a custom HTTPClient for a request.
func NewImageSearchParamsWithHTTPClient(client *http.Client) *ImageSearchParams {
	return &ImageSearchParams{
		HTTPClient: client,
	}
}

/* ImageSearchParams contains all the parameters to send to the API endpoint
   for the image search operation.

   Typically these are written to a http.Request.
*/
type ImageSearchParams struct {

	/* Filters.

	     A JSON encoded value of the filters (a `map[string][]string`) to process on the images list. Available filters:

	- `is-automated=(true|false)`
	- `is-official=(true|false)`
	- `stars=<number>` Matches images that has at least 'number' stars.

	*/
	Filters *string

	/* Limit.

	   Maximum number of results to return
	*/
	Limit *int64

	/* Term.

	   Term to search
	*/
	Term string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the image search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageSearchParams) WithDefaults() *ImageSearchParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the image search params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ImageSearchParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the image search params
func (o *ImageSearchParams) WithTimeout(timeout time.Duration) *ImageSearchParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the image search params
func (o *ImageSearchParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the image search params
func (o *ImageSearchParams) WithContext(ctx context.Context) *ImageSearchParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the image search params
func (o *ImageSearchParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the image search params
func (o *ImageSearchParams) WithHTTPClient(client *http.Client) *ImageSearchParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the image search params
func (o *ImageSearchParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the image search params
func (o *ImageSearchParams) WithFilters(filters *string) *ImageSearchParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the image search params
func (o *ImageSearchParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WithLimit adds the limit to the image search params
func (o *ImageSearchParams) WithLimit(limit *int64) *ImageSearchParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the image search params
func (o *ImageSearchParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithTerm adds the term to the image search params
func (o *ImageSearchParams) WithTerm(term string) *ImageSearchParams {
	o.SetTerm(term)
	return o
}

// SetTerm adds the term to the image search params
func (o *ImageSearchParams) SetTerm(term string) {
	o.Term = term
}

// WriteToRequest writes these params to a swagger request
func (o *ImageSearchParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Filters != nil {

		// query param filters
		var qrFilters string

		if o.Filters != nil {
			qrFilters = *o.Filters
		}
		qFilters := qrFilters
		if qFilters != "" {

			if err := r.SetQueryParam("filters", qFilters); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	// query param term
	qrTerm := o.Term
	qTerm := qrTerm
	if qTerm != "" {

		if err := r.SetQueryParam("term", qTerm); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}