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

package plugin

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

	"github.com/bhojpur/host/pkg/machine/models"
)

// NewPluginPullParams creates a new PluginPullParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPluginPullParams() *PluginPullParams {
	return &PluginPullParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPluginPullParamsWithTimeout creates a new PluginPullParams object
// with the ability to set a timeout on a request.
func NewPluginPullParamsWithTimeout(timeout time.Duration) *PluginPullParams {
	return &PluginPullParams{
		timeout: timeout,
	}
}

// NewPluginPullParamsWithContext creates a new PluginPullParams object
// with the ability to set a context for a request.
func NewPluginPullParamsWithContext(ctx context.Context) *PluginPullParams {
	return &PluginPullParams{
		Context: ctx,
	}
}

// NewPluginPullParamsWithHTTPClient creates a new PluginPullParams object
// with the ability to set a custom HTTPClient for a request.
func NewPluginPullParamsWithHTTPClient(client *http.Client) *PluginPullParams {
	return &PluginPullParams{
		HTTPClient: client,
	}
}

/* PluginPullParams contains all the parameters to send to the API endpoint
   for the plugin pull operation.

   Typically these are written to a http.Request.
*/
type PluginPullParams struct {

	/* XRegistryAuth.

	     A base64url-encoded auth configuration to use when pulling a plugin
	from a registry.

	Refer to the [authentication section](#section/Authentication) for
	details.

	*/
	XRegistryAuth *string

	// Body.
	Body []*models.PluginPrivilege

	/* Name.

	     Local name for the pulled plugin.

	The `:latest` tag is optional, and is used as the default if omitted.

	*/
	Name *string

	/* Remote.

	     Remote reference for plugin to install.

	The `:latest` tag is optional, and is used as the default if omitted.

	*/
	Remote string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the plugin pull params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginPullParams) WithDefaults() *PluginPullParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the plugin pull params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PluginPullParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the plugin pull params
func (o *PluginPullParams) WithTimeout(timeout time.Duration) *PluginPullParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the plugin pull params
func (o *PluginPullParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the plugin pull params
func (o *PluginPullParams) WithContext(ctx context.Context) *PluginPullParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the plugin pull params
func (o *PluginPullParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the plugin pull params
func (o *PluginPullParams) WithHTTPClient(client *http.Client) *PluginPullParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the plugin pull params
func (o *PluginPullParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRegistryAuth adds the xRegistryAuth to the plugin pull params
func (o *PluginPullParams) WithXRegistryAuth(xRegistryAuth *string) *PluginPullParams {
	o.SetXRegistryAuth(xRegistryAuth)
	return o
}

// SetXRegistryAuth adds the xRegistryAuth to the plugin pull params
func (o *PluginPullParams) SetXRegistryAuth(xRegistryAuth *string) {
	o.XRegistryAuth = xRegistryAuth
}

// WithBody adds the body to the plugin pull params
func (o *PluginPullParams) WithBody(body []*models.PluginPrivilege) *PluginPullParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the plugin pull params
func (o *PluginPullParams) SetBody(body []*models.PluginPrivilege) {
	o.Body = body
}

// WithName adds the name to the plugin pull params
func (o *PluginPullParams) WithName(name *string) *PluginPullParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the plugin pull params
func (o *PluginPullParams) SetName(name *string) {
	o.Name = name
}

// WithRemote adds the remote to the plugin pull params
func (o *PluginPullParams) WithRemote(remote string) *PluginPullParams {
	o.SetRemote(remote)
	return o
}

// SetRemote adds the remote to the plugin pull params
func (o *PluginPullParams) SetRemote(remote string) {
	o.Remote = remote
}

// WriteToRequest writes these params to a swagger request
func (o *PluginPullParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRegistryAuth != nil {

		// header param X-Registry-Auth
		if err := r.SetHeaderParam("X-Registry-Auth", *o.XRegistryAuth); err != nil {
			return err
		}
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	// query param remote
	qrRemote := o.Remote
	qRemote := qrRemote
	if qRemote != "" {

		if err := r.SetQueryParam("remote", qRemote); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}