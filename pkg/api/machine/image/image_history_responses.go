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
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ImageHistoryReader is a Reader for the ImageHistory structure.
type ImageHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImageHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImageHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewImageHistoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewImageHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImageHistoryOK creates a ImageHistoryOK with default headers values
func NewImageHistoryOK() *ImageHistoryOK {
	return &ImageHistoryOK{}
}

/* ImageHistoryOK describes a response with status code 200, with default header values.

List of image layers
*/
type ImageHistoryOK struct {
	Payload []*HistoryResponseItem
}

func (o *ImageHistoryOK) Error() string {
	return fmt.Sprintf("[GET /images/{name}/history][%d] imageHistoryOK  %+v", 200, o.Payload)
}
func (o *ImageHistoryOK) GetPayload() []*HistoryResponseItem {
	return o.Payload
}

func (o *ImageHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageHistoryNotFound creates a ImageHistoryNotFound with default headers values
func NewImageHistoryNotFound() *ImageHistoryNotFound {
	return &ImageHistoryNotFound{}
}

/* ImageHistoryNotFound describes a response with status code 404, with default header values.

No such image
*/
type ImageHistoryNotFound struct {
	Payload *models.ErrorResponse
}

func (o *ImageHistoryNotFound) Error() string {
	return fmt.Sprintf("[GET /images/{name}/history][%d] imageHistoryNotFound  %+v", 404, o.Payload)
}
func (o *ImageHistoryNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageHistoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImageHistoryInternalServerError creates a ImageHistoryInternalServerError with default headers values
func NewImageHistoryInternalServerError() *ImageHistoryInternalServerError {
	return &ImageHistoryInternalServerError{}
}

/* ImageHistoryInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImageHistoryInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImageHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{name}/history][%d] imageHistoryInternalServerError  %+v", 500, o.Payload)
}
func (o *ImageHistoryInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImageHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*HistoryResponseItem HistoryResponseItem
//
// individual image layer information in response to ImageHistory operation
swagger:model HistoryResponseItem
*/
type HistoryResponseItem struct {

	// comment
	// Required: true
	Comment string `json:"Comment"`

	// created
	// Required: true
	Created int64 `json:"Created"`

	// created by
	// Required: true
	CreatedBy string `json:"CreatedBy"`

	// Id
	// Required: true
	ID string `json:"Id"`

	// size
	// Required: true
	Size int64 `json:"Size"`

	// tags
	// Required: true
	Tags []string `json:"Tags"`
}

// Validate validates this history response item
func (o *HistoryResponseItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSize(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HistoryResponseItem) validateComment(formats strfmt.Registry) error {

	if err := validate.RequiredString("Comment", "body", o.Comment); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateCreated(formats strfmt.Registry) error {

	if err := validate.Required("Created", "body", int64(o.Created)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.RequiredString("CreatedBy", "body", o.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("Id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateSize(formats strfmt.Registry) error {

	if err := validate.Required("Size", "body", int64(o.Size)); err != nil {
		return err
	}

	return nil
}

func (o *HistoryResponseItem) validateTags(formats strfmt.Registry) error {

	if err := validate.Required("Tags", "body", o.Tags); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this history response item based on context it is used
func (o *HistoryResponseItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *HistoryResponseItem) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *HistoryResponseItem) UnmarshalBinary(b []byte) error {
	var res HistoryResponseItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}