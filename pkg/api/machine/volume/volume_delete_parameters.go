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

package volume

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

// NewVolumeDeleteParams creates a new VolumeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeDeleteParams() *VolumeDeleteParams {
	return &VolumeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeDeleteParamsWithTimeout creates a new VolumeDeleteParams object
// with the ability to set a timeout on a request.
func NewVolumeDeleteParamsWithTimeout(timeout time.Duration) *VolumeDeleteParams {
	return &VolumeDeleteParams{
		timeout: timeout,
	}
}

// NewVolumeDeleteParamsWithContext creates a new VolumeDeleteParams object
// with the ability to set a context for a request.
func NewVolumeDeleteParamsWithContext(ctx context.Context) *VolumeDeleteParams {
	return &VolumeDeleteParams{
		Context: ctx,
	}
}

// NewVolumeDeleteParamsWithHTTPClient creates a new VolumeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeDeleteParamsWithHTTPClient(client *http.Client) *VolumeDeleteParams {
	return &VolumeDeleteParams{
		HTTPClient: client,
	}
}

/* VolumeDeleteParams contains all the parameters to send to the API endpoint
   for the volume delete operation.

   Typically these are written to a http.Request.
*/
type VolumeDeleteParams struct {

	/* Force.

	   Force the removal of the volume
	*/
	Force *bool

	/* Name.

	   Volume name or ID
	*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeDeleteParams) WithDefaults() *VolumeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeDeleteParams) SetDefaults() {
	var (
		forceDefault = bool(false)
	)

	val := VolumeDeleteParams{
		Force: &forceDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the volume delete params
func (o *VolumeDeleteParams) WithTimeout(timeout time.Duration) *VolumeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume delete params
func (o *VolumeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume delete params
func (o *VolumeDeleteParams) WithContext(ctx context.Context) *VolumeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume delete params
func (o *VolumeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume delete params
func (o *VolumeDeleteParams) WithHTTPClient(client *http.Client) *VolumeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume delete params
func (o *VolumeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForce adds the force to the volume delete params
func (o *VolumeDeleteParams) WithForce(force *bool) *VolumeDeleteParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the volume delete params
func (o *VolumeDeleteParams) SetForce(force *bool) {
	o.Force = force
}

// WithName adds the name to the volume delete params
func (o *VolumeDeleteParams) WithName(name string) *VolumeDeleteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the volume delete params
func (o *VolumeDeleteParams) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Force != nil {

		// query param force
		var qrForce bool

		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}