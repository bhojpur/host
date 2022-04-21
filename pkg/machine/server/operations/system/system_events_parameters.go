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

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// NewSystemEventsParams creates a new SystemEventsParams object
//
// There are no default values defined in the spec.
func NewSystemEventsParams() SystemEventsParams {

	return SystemEventsParams{}
}

// SystemEventsParams contains all the bound params for the system events operation
// typically these are obtained from a http.Request
//
// swagger:parameters SystemEvents
type SystemEventsParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A JSON encoded value of filters (a `map[string][]string`) to process on the event list. Available filters:

	- `config=<string>` config name or ID
	- `container=<string>` container name or ID
	- `daemon=<string>` daemon name or ID
	- `event=<string>` event type
	- `image=<string>` image name or ID
	- `label=<string>` image or container label
	- `network=<string>` network name or ID
	- `node=<string>` node ID
	- `plugin`=<string> plugin name or ID
	- `scope`=<string> local or swarm
	- `secret=<string>` secret name or ID
	- `service=<string>` service name or ID
	- `type=<string>` object to filter by, one of `container`, `image`, `volume`, `network`, `daemon`, `plugin`, `node`, `service`, `secret` or `config`
	- `volume=<string>` volume name

	  In: query
	*/
	Filters *string
	/*Show events created since this timestamp then stream new events.
	  In: query
	*/
	Since *string
	/*Show events created until this timestamp then stop streaming.
	  In: query
	*/
	Until *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewSystemEventsParams() beforehand.
func (o *SystemEventsParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qFilters, qhkFilters, _ := qs.GetOK("filters")
	if err := o.bindFilters(qFilters, qhkFilters, route.Formats); err != nil {
		res = append(res, err)
	}

	qSince, qhkSince, _ := qs.GetOK("since")
	if err := o.bindSince(qSince, qhkSince, route.Formats); err != nil {
		res = append(res, err)
	}

	qUntil, qhkUntil, _ := qs.GetOK("until")
	if err := o.bindUntil(qUntil, qhkUntil, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindFilters binds and validates parameter Filters from query.
func (o *SystemEventsParams) bindFilters(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Filters = &raw

	return nil
}

// bindSince binds and validates parameter Since from query.
func (o *SystemEventsParams) bindSince(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Since = &raw

	return nil
}

// bindUntil binds and validates parameter Until from query.
func (o *SystemEventsParams) bindUntil(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Until = &raw

	return nil
}
