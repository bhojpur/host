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

package task

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"

	"github.com/go-openapi/swag"
)

// TaskLogsURL generates an URL for the task logs operation
type TaskLogsURL struct {
	ID string

	Details    *bool
	Follow     *bool
	Since      *int64
	Stderr     *bool
	Stdout     *bool
	Tail       *string
	Timestamps *bool

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TaskLogsURL) WithBasePath(bp string) *TaskLogsURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TaskLogsURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *TaskLogsURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/tasks/{id}/logs"

	id := o.ID
	if id != "" {
		_path = strings.Replace(_path, "{id}", id, -1)
	} else {
		return nil, errors.New("id is required on TaskLogsURL")
	}

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/v1.42"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var detailsQ string
	if o.Details != nil {
		detailsQ = swag.FormatBool(*o.Details)
	}
	if detailsQ != "" {
		qs.Set("details", detailsQ)
	}

	var followQ string
	if o.Follow != nil {
		followQ = swag.FormatBool(*o.Follow)
	}
	if followQ != "" {
		qs.Set("follow", followQ)
	}

	var sinceQ string
	if o.Since != nil {
		sinceQ = swag.FormatInt64(*o.Since)
	}
	if sinceQ != "" {
		qs.Set("since", sinceQ)
	}

	var stderrQ string
	if o.Stderr != nil {
		stderrQ = swag.FormatBool(*o.Stderr)
	}
	if stderrQ != "" {
		qs.Set("stderr", stderrQ)
	}

	var stdoutQ string
	if o.Stdout != nil {
		stdoutQ = swag.FormatBool(*o.Stdout)
	}
	if stdoutQ != "" {
		qs.Set("stdout", stdoutQ)
	}

	var tailQ string
	if o.Tail != nil {
		tailQ = *o.Tail
	}
	if tailQ != "" {
		qs.Set("tail", tailQ)
	}

	var timestampsQ string
	if o.Timestamps != nil {
		timestampsQ = swag.FormatBool(*o.Timestamps)
	}
	if timestampsQ != "" {
		qs.Set("timestamps", timestampsQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *TaskLogsURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *TaskLogsURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *TaskLogsURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on TaskLogsURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on TaskLogsURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *TaskLogsURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
