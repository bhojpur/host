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

package secret

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

// NewSecretListParams creates a new SecretListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecretListParams() *SecretListParams {
	return &SecretListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecretListParamsWithTimeout creates a new SecretListParams object
// with the ability to set a timeout on a request.
func NewSecretListParamsWithTimeout(timeout time.Duration) *SecretListParams {
	return &SecretListParams{
		timeout: timeout,
	}
}

// NewSecretListParamsWithContext creates a new SecretListParams object
// with the ability to set a context for a request.
func NewSecretListParamsWithContext(ctx context.Context) *SecretListParams {
	return &SecretListParams{
		Context: ctx,
	}
}

// NewSecretListParamsWithHTTPClient creates a new SecretListParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecretListParamsWithHTTPClient(client *http.Client) *SecretListParams {
	return &SecretListParams{
		HTTPClient: client,
	}
}

/* SecretListParams contains all the parameters to send to the API endpoint
   for the secret list operation.

   Typically these are written to a http.Request.
*/
type SecretListParams struct {

	/* Filters.

	     A JSON encoded value of the filters (a `map[string][]string`) to
	process on the secrets list.

	Available filters:

	- `id=<secret id>`
	- `label=<key> or label=<key>=value`
	- `name=<secret name>`
	- `names=<secret name>`

	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the secret list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretListParams) WithDefaults() *SecretListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the secret list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecretListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the secret list params
func (o *SecretListParams) WithTimeout(timeout time.Duration) *SecretListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the secret list params
func (o *SecretListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the secret list params
func (o *SecretListParams) WithContext(ctx context.Context) *SecretListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the secret list params
func (o *SecretListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the secret list params
func (o *SecretListParams) WithHTTPClient(client *http.Client) *SecretListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the secret list params
func (o *SecretListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the secret list params
func (o *SecretListParams) WithFilters(filters *string) *SecretListParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the secret list params
func (o *SecretListParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *SecretListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
