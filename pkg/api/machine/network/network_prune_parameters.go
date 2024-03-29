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

// NewNetworkPruneParams creates a new NetworkPruneParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetworkPruneParams() *NetworkPruneParams {
	return &NetworkPruneParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetworkPruneParamsWithTimeout creates a new NetworkPruneParams object
// with the ability to set a timeout on a request.
func NewNetworkPruneParamsWithTimeout(timeout time.Duration) *NetworkPruneParams {
	return &NetworkPruneParams{
		timeout: timeout,
	}
}

// NewNetworkPruneParamsWithContext creates a new NetworkPruneParams object
// with the ability to set a context for a request.
func NewNetworkPruneParamsWithContext(ctx context.Context) *NetworkPruneParams {
	return &NetworkPruneParams{
		Context: ctx,
	}
}

// NewNetworkPruneParamsWithHTTPClient creates a new NetworkPruneParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetworkPruneParamsWithHTTPClient(client *http.Client) *NetworkPruneParams {
	return &NetworkPruneParams{
		HTTPClient: client,
	}
}

/* NetworkPruneParams contains all the parameters to send to the API endpoint
   for the network prune operation.

   Typically these are written to a http.Request.
*/
type NetworkPruneParams struct {

	/* Filters.

	     Filters to process on the prune list, encoded as JSON (a `map[string][]string`).

	Available filters:
	- `until=<timestamp>` Prune networks created before this timestamp. The `<timestamp>` can be Unix timestamps, date formatted timestamps, or Go duration strings (e.g. `10m`, `1h30m`) computed relative to the daemon machine’s time.
	- `label` (`label=<key>`, `label=<key>=<value>`, `label!=<key>`, or `label!=<key>=<value>`) Prune networks with (or without, in case `label!=...` is used) the specified labels.

	*/
	Filters *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the network prune params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkPruneParams) WithDefaults() *NetworkPruneParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the network prune params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetworkPruneParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the network prune params
func (o *NetworkPruneParams) WithTimeout(timeout time.Duration) *NetworkPruneParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the network prune params
func (o *NetworkPruneParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the network prune params
func (o *NetworkPruneParams) WithContext(ctx context.Context) *NetworkPruneParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the network prune params
func (o *NetworkPruneParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the network prune params
func (o *NetworkPruneParams) WithHTTPClient(client *http.Client) *NetworkPruneParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the network prune params
func (o *NetworkPruneParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFilters adds the filters to the network prune params
func (o *NetworkPruneParams) WithFilters(filters *string) *NetworkPruneParams {
	o.SetFilters(filters)
	return o
}

// SetFilters adds the filters to the network prune params
func (o *NetworkPruneParams) SetFilters(filters *string) {
	o.Filters = filters
}

// WriteToRequest writes these params to a swagger request
func (o *NetworkPruneParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
