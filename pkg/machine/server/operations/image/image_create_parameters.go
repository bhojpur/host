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
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewImageCreateParams creates a new ImageCreateParams object
// with the default values initialized.
func NewImageCreateParams() ImageCreateParams {

	var (
		// initialize parameters with default values

		platformDefault = string("")
	)

	return ImageCreateParams{
		Platform: &platformDefault,
	}
}

// ImageCreateParams contains all the bound params for the image create operation
// typically these are obtained from a http.Request
//
// swagger:parameters ImageCreate
type ImageCreateParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*A base64url-encoded auth configuration.

	Refer to the [authentication section](#section/Authentication) for
	details.

	  In: header
	*/
	XRegistryAuth *string
	/*Apply `Bhojpurfile` instructions to the image that is created,
	for example: `changes=ENV DEBUG=true`.
	Note that `ENV DEBUG=true` should be URI component encoded.

	Supported `Bhojpurfile` instructions:
	`CMD`|`ENTRYPOINT`|`ENV`|`EXPOSE`|`ONBUILD`|`USER`|`VOLUME`|`WORKDIR`

	  In: query
	*/
	Changes []string
	/*Name of the image to pull. The name may include a tag or digest. This parameter may only be used when pulling an image. The pull is cancelled if the HTTP connection is closed.
	  In: query
	*/
	FromImage *string
	/*Source to import. The value may be a URL from which the image can be retrieved or `-` to read the image from the request body. This parameter may only be used when importing an image.
	  In: query
	*/
	FromSrc *string
	/*Image content if the value `-` has been specified in fromSrc query parameter
	  In: body
	*/
	InputImage string
	/*Set commit message for imported image.
	  In: query
	*/
	Message *string
	/*Platform in the format os[/arch[/variant]].

	When used in combination with the `fromImage` option, the daemon checks
	if the given image is present in the local image cache with the given
	OS and Architecture, and otherwise attempts to pull the image. If the
	option is not set, the host's native OS and Architecture are used.
	If the given image does not exist in the local image cache, the daemon
	attempts to pull the image with the host's native OS and Architecture.
	If the given image does exists in the local image cache, but its OS or
	architecture does not match, a warning is produced.

	When used with the `fromSrc` option to import an image from an archive,
	this option sets the platform information for the imported image. If
	the option is not set, the host's native OS and Architecture are used
	for the imported image.

	  In: query
	  Default: ""
	*/
	Platform *string
	/*Repository name given to an image when it is imported. The repo may include a tag. This parameter may only be used when importing an image.
	  In: query
	*/
	Repo *string
	/*Tag or digest. If empty when pulling an image, this causes all tags for the given image to be pulled.
	  In: query
	*/
	Tag *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewImageCreateParams() beforehand.
func (o *ImageCreateParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	if err := o.bindXRegistryAuth(r.Header[http.CanonicalHeaderKey("X-Registry-Auth")], true, route.Formats); err != nil {
		res = append(res, err)
	}

	qChanges, qhkChanges, _ := qs.GetOK("changes")
	if err := o.bindChanges(qChanges, qhkChanges, route.Formats); err != nil {
		res = append(res, err)
	}

	qFromImage, qhkFromImage, _ := qs.GetOK("fromImage")
	if err := o.bindFromImage(qFromImage, qhkFromImage, route.Formats); err != nil {
		res = append(res, err)
	}

	qFromSrc, qhkFromSrc, _ := qs.GetOK("fromSrc")
	if err := o.bindFromSrc(qFromSrc, qhkFromSrc, route.Formats); err != nil {
		res = append(res, err)
	}

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body string
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("inputImage", "body", "", err))
		} else {
			// no validation required on inline body
			o.InputImage = body
		}
	}

	qMessage, qhkMessage, _ := qs.GetOK("message")
	if err := o.bindMessage(qMessage, qhkMessage, route.Formats); err != nil {
		res = append(res, err)
	}

	qPlatform, qhkPlatform, _ := qs.GetOK("platform")
	if err := o.bindPlatform(qPlatform, qhkPlatform, route.Formats); err != nil {
		res = append(res, err)
	}

	qRepo, qhkRepo, _ := qs.GetOK("repo")
	if err := o.bindRepo(qRepo, qhkRepo, route.Formats); err != nil {
		res = append(res, err)
	}

	qTag, qhkTag, _ := qs.GetOK("tag")
	if err := o.bindTag(qTag, qhkTag, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindXRegistryAuth binds and validates parameter XRegistryAuth from header.
func (o *ImageCreateParams) bindXRegistryAuth(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.XRegistryAuth = &raw

	return nil
}

// bindChanges binds and validates array parameter Changes from query.
//
// Arrays are parsed according to CollectionFormat: "" (defaults to "csv" when empty).
func (o *ImageCreateParams) bindChanges(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var qvChanges string
	if len(rawData) > 0 {
		qvChanges = rawData[len(rawData)-1]
	}

	// CollectionFormat:
	changesIC := swag.SplitByFormat(qvChanges, "")
	if len(changesIC) == 0 {
		return nil
	}

	var changesIR []string
	for _, changesIV := range changesIC {
		changesI := changesIV

		changesIR = append(changesIR, changesI)
	}

	o.Changes = changesIR

	return nil
}

// bindFromImage binds and validates parameter FromImage from query.
func (o *ImageCreateParams) bindFromImage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FromImage = &raw

	return nil
}

// bindFromSrc binds and validates parameter FromSrc from query.
func (o *ImageCreateParams) bindFromSrc(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.FromSrc = &raw

	return nil
}

// bindMessage binds and validates parameter Message from query.
func (o *ImageCreateParams) bindMessage(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Message = &raw

	return nil
}

// bindPlatform binds and validates parameter Platform from query.
func (o *ImageCreateParams) bindPlatform(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		// Default values have been previously initialized by NewImageCreateParams()
		return nil
	}
	o.Platform = &raw

	return nil
}

// bindRepo binds and validates parameter Repo from query.
func (o *ImageCreateParams) bindRepo(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Repo = &raw

	return nil
}

// bindTag binds and validates parameter Tag from query.
func (o *ImageCreateParams) bindTag(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false

	if raw == "" { // empty values pass all other validations
		return nil
	}
	o.Tag = &raw

	return nil
}