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

// NewContainerStatsParams creates a new ContainerStatsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContainerStatsParams() *ContainerStatsParams {
	return &ContainerStatsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContainerStatsParamsWithTimeout creates a new ContainerStatsParams object
// with the ability to set a timeout on a request.
func NewContainerStatsParamsWithTimeout(timeout time.Duration) *ContainerStatsParams {
	return &ContainerStatsParams{
		timeout: timeout,
	}
}

// NewContainerStatsParamsWithContext creates a new ContainerStatsParams object
// with the ability to set a context for a request.
func NewContainerStatsParamsWithContext(ctx context.Context) *ContainerStatsParams {
	return &ContainerStatsParams{
		Context: ctx,
	}
}

// NewContainerStatsParamsWithHTTPClient creates a new ContainerStatsParams object
// with the ability to set a custom HTTPClient for a request.
func NewContainerStatsParamsWithHTTPClient(client *http.Client) *ContainerStatsParams {
	return &ContainerStatsParams{
		HTTPClient: client,
	}
}

/* ContainerStatsParams contains all the parameters to send to the API endpoint
   for the container stats operation.

   Typically these are written to a http.Request.
*/
type ContainerStatsParams struct {

	/* ID.

	   ID or name of the container
	*/
	ID string

	/* OneShot.

	     Only get a single stat instead of waiting for 2 cycles. Must be used
	with `stream=false`.

	*/
	OneShot *bool

	/* Stream.

	     Stream the output. If false, the stats will be output once and then
	it will disconnect.


	     Default: true
	*/
	Stream *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the container stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerStatsParams) WithDefaults() *ContainerStatsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the container stats params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContainerStatsParams) SetDefaults() {
	var (
		oneShotDefault = bool(false)

		streamDefault = bool(true)
	)

	val := ContainerStatsParams{
		OneShot: &oneShotDefault,
		Stream:  &streamDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the container stats params
func (o *ContainerStatsParams) WithTimeout(timeout time.Duration) *ContainerStatsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container stats params
func (o *ContainerStatsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container stats params
func (o *ContainerStatsParams) WithContext(ctx context.Context) *ContainerStatsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container stats params
func (o *ContainerStatsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container stats params
func (o *ContainerStatsParams) WithHTTPClient(client *http.Client) *ContainerStatsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container stats params
func (o *ContainerStatsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the container stats params
func (o *ContainerStatsParams) WithID(id string) *ContainerStatsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container stats params
func (o *ContainerStatsParams) SetID(id string) {
	o.ID = id
}

// WithOneShot adds the oneShot to the container stats params
func (o *ContainerStatsParams) WithOneShot(oneShot *bool) *ContainerStatsParams {
	o.SetOneShot(oneShot)
	return o
}

// SetOneShot adds the oneShot to the container stats params
func (o *ContainerStatsParams) SetOneShot(oneShot *bool) {
	o.OneShot = oneShot
}

// WithStream adds the stream to the container stats params
func (o *ContainerStatsParams) WithStream(stream *bool) *ContainerStatsParams {
	o.SetStream(stream)
	return o
}

// SetStream adds the stream to the container stats params
func (o *ContainerStatsParams) SetStream(stream *bool) {
	o.Stream = stream
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerStatsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if o.OneShot != nil {

		// query param one-shot
		var qrOneShot bool

		if o.OneShot != nil {
			qrOneShot = *o.OneShot
		}
		qOneShot := swag.FormatBool(qrOneShot)
		if qOneShot != "" {

			if err := r.SetQueryParam("one-shot", qOneShot); err != nil {
				return err
			}
		}
	}

	if o.Stream != nil {

		// query param stream
		var qrStream bool

		if o.Stream != nil {
			qrStream = *o.Stream
		}
		qStream := swag.FormatBool(qrStream)
		if qStream != "" {

			if err := r.SetQueryParam("stream", qStream); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
