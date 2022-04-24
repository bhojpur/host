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
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// ImagePruneReader is a Reader for the ImagePrune structure.
type ImagePruneReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ImagePruneReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewImagePruneOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewImagePruneInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewImagePruneOK creates a ImagePruneOK with default headers values
func NewImagePruneOK() *ImagePruneOK {
	return &ImagePruneOK{}
}

/* ImagePruneOK describes a response with status code 200, with default header values.

No error
*/
type ImagePruneOK struct {
	Payload *ImagePruneOKBody
}

func (o *ImagePruneOK) Error() string {
	return fmt.Sprintf("[POST /images/prune][%d] imagePruneOK  %+v", 200, o.Payload)
}
func (o *ImagePruneOK) GetPayload() *ImagePruneOKBody {
	return o.Payload
}

func (o *ImagePruneOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ImagePruneOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewImagePruneInternalServerError creates a ImagePruneInternalServerError with default headers values
func NewImagePruneInternalServerError() *ImagePruneInternalServerError {
	return &ImagePruneInternalServerError{}
}

/* ImagePruneInternalServerError describes a response with status code 500, with default header values.

Server error
*/
type ImagePruneInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ImagePruneInternalServerError) Error() string {
	return fmt.Sprintf("[POST /images/prune][%d] imagePruneInternalServerError  %+v", 500, o.Payload)
}
func (o *ImagePruneInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ImagePruneInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ImagePruneOKBody ImagePruneResponse
swagger:model ImagePruneOKBody
*/
type ImagePruneOKBody struct {

	// Images that were deleted
	ImagesDeleted []*models.ImageDeleteResponseItem `json:"ImagesDeleted"`

	// Disk space reclaimed in bytes
	SpaceReclaimed int64 `json:"SpaceReclaimed,omitempty"`
}

// Validate validates this image prune o k body
func (o *ImagePruneOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateImagesDeleted(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImagePruneOKBody) validateImagesDeleted(formats strfmt.Registry) error {
	if swag.IsZero(o.ImagesDeleted) { // not required
		return nil
	}

	for i := 0; i < len(o.ImagesDeleted); i++ {
		if swag.IsZero(o.ImagesDeleted[i]) { // not required
			continue
		}

		if o.ImagesDeleted[i] != nil {
			if err := o.ImagesDeleted[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePruneOK" + "." + "ImagesDeleted" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("imagePruneOK" + "." + "ImagesDeleted" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this image prune o k body based on the context it is used
func (o *ImagePruneOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateImagesDeleted(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ImagePruneOKBody) contextValidateImagesDeleted(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ImagesDeleted); i++ {

		if o.ImagesDeleted[i] != nil {
			if err := o.ImagesDeleted[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("imagePruneOK" + "." + "ImagesDeleted" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("imagePruneOK" + "." + "ImagesDeleted" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *ImagePruneOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ImagePruneOKBody) UnmarshalBinary(b []byte) error {
	var res ImagePruneOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}