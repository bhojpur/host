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

package secret

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

	"github.com/bhojpur/host/pkg/machine/models"
)

// SecretCreateReader is a Reader for the SecretCreate structure.
type SecretCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecretCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSecretCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 409:
		result := NewSecretCreateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSecretCreateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSecretCreateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSecretCreateCreated creates a SecretCreateCreated with default headers values
func NewSecretCreateCreated() *SecretCreateCreated {
	return &SecretCreateCreated{}
}

/* SecretCreateCreated describes a response with status code 201, with default header values.

no error
*/
type SecretCreateCreated struct {
	Payload *models.IDResponse
}

func (o *SecretCreateCreated) Error() string {
	return fmt.Sprintf("[POST /secrets/create][%d] secretCreateCreated  %+v", 201, o.Payload)
}
func (o *SecretCreateCreated) GetPayload() *models.IDResponse {
	return o.Payload
}

func (o *SecretCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretCreateConflict creates a SecretCreateConflict with default headers values
func NewSecretCreateConflict() *SecretCreateConflict {
	return &SecretCreateConflict{}
}

/* SecretCreateConflict describes a response with status code 409, with default header values.

name conflicts with an existing object
*/
type SecretCreateConflict struct {
	Payload *models.ErrorResponse
}

func (o *SecretCreateConflict) Error() string {
	return fmt.Sprintf("[POST /secrets/create][%d] secretCreateConflict  %+v", 409, o.Payload)
}
func (o *SecretCreateConflict) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecretCreateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretCreateInternalServerError creates a SecretCreateInternalServerError with default headers values
func NewSecretCreateInternalServerError() *SecretCreateInternalServerError {
	return &SecretCreateInternalServerError{}
}

/* SecretCreateInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SecretCreateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SecretCreateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /secrets/create][%d] secretCreateInternalServerError  %+v", 500, o.Payload)
}
func (o *SecretCreateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecretCreateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecretCreateServiceUnavailable creates a SecretCreateServiceUnavailable with default headers values
func NewSecretCreateServiceUnavailable() *SecretCreateServiceUnavailable {
	return &SecretCreateServiceUnavailable{}
}

/* SecretCreateServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type SecretCreateServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *SecretCreateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /secrets/create][%d] secretCreateServiceUnavailable  %+v", 503, o.Payload)
}
func (o *SecretCreateServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecretCreateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SecretCreateBody secret create body
swagger:model SecretCreateBody
*/
type SecretCreateBody struct {
	models.SecretSpec

	SecretCreateParamsBodyAllOf1
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *SecretCreateBody) UnmarshalJSON(raw []byte) error {
	// SecretCreateParamsBodyAO0
	var secretCreateParamsBodyAO0 models.SecretSpec
	if err := swag.ReadJSON(raw, &secretCreateParamsBodyAO0); err != nil {
		return err
	}
	o.SecretSpec = secretCreateParamsBodyAO0

	// SecretCreateParamsBodyAO1
	var secretCreateParamsBodyAO1 SecretCreateParamsBodyAllOf1
	if err := swag.ReadJSON(raw, &secretCreateParamsBodyAO1); err != nil {
		return err
	}
	o.SecretCreateParamsBodyAllOf1 = secretCreateParamsBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o SecretCreateBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	secretCreateParamsBodyAO0, err := swag.WriteJSON(o.SecretSpec)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, secretCreateParamsBodyAO0)

	secretCreateParamsBodyAO1, err := swag.WriteJSON(o.SecretCreateParamsBodyAllOf1)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, secretCreateParamsBodyAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this secret create body
func (o *SecretCreateBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SecretSpec
	if err := o.SecretSpec.Validate(formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SecretCreateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this secret create body based on the context it is used
func (o *SecretCreateBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.SecretSpec
	if err := o.SecretSpec.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}
	// validation for a type composition with SecretCreateParamsBodyAllOf1

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *SecretCreateBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SecretCreateBody) UnmarshalBinary(b []byte) error {
	var res SecretCreateBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*SecretCreateParamsBodyAllOf1 secret create params body all of1
// Example: {"Data":"VEhJUyBJUyBOT1QgQSBSRUFMIENFUlRJRklDQVRFCg==","Driver":{"Name":"secret-bucket","Options":{"OptionA":"value for driver option A","OptionB":"value for driver option B"}},"Labels":{"foo":"bar"},"Name":"app-key.crt"}
swagger:model SecretCreateParamsBodyAllOf1
*/
type SecretCreateParamsBodyAllOf1 interface{}
