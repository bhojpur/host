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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new config API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for config API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ConfigCreate(params *ConfigCreateParams, opts ...ClientOption) (*ConfigCreateCreated, error)

	ConfigDelete(params *ConfigDeleteParams, opts ...ClientOption) (*ConfigDeleteNoContent, error)

	ConfigInspect(params *ConfigInspectParams, opts ...ClientOption) (*ConfigInspectOK, error)

	ConfigList(params *ConfigListParams, opts ...ClientOption) (*ConfigListOK, error)

	ConfigUpdate(params *ConfigUpdateParams, opts ...ClientOption) (*ConfigUpdateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ConfigCreate creates a config
*/
func (a *Client) ConfigCreate(params *ConfigCreateParams, opts ...ClientOption) (*ConfigCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigCreateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConfigCreate",
		Method:             "POST",
		PathPattern:        "/configs/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfigCreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConfigCreate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConfigDelete deletes a config
*/
func (a *Client) ConfigDelete(params *ConfigDeleteParams, opts ...ClientOption) (*ConfigDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigDeleteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConfigDelete",
		Method:             "DELETE",
		PathPattern:        "/configs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfigDeleteNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConfigDelete: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConfigInspect inspects a config
*/
func (a *Client) ConfigInspect(params *ConfigInspectParams, opts ...ClientOption) (*ConfigInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigInspectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConfigInspect",
		Method:             "GET",
		PathPattern:        "/configs/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfigInspectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConfigInspect: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConfigList lists configs
*/
func (a *Client) ConfigList(params *ConfigListParams, opts ...ClientOption) (*ConfigListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConfigList",
		Method:             "GET",
		PathPattern:        "/configs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfigListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConfigList: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ConfigUpdate updates a config
*/
func (a *Client) ConfigUpdate(params *ConfigUpdateParams, opts ...ClientOption) (*ConfigUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConfigUpdateParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ConfigUpdate",
		Method:             "POST",
		PathPattern:        "/configs/{id}/update",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConfigUpdateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ConfigUpdateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for ConfigUpdate: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
