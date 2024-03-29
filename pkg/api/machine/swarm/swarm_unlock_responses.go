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

package swarm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/bhojpur/host/pkg/machine/models"
)

// SwarmUnlockReader is a Reader for the SwarmUnlock structure.
type SwarmUnlockReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SwarmUnlockReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSwarmUnlockOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewSwarmUnlockInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewSwarmUnlockServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewSwarmUnlockOK creates a SwarmUnlockOK with default headers values
func NewSwarmUnlockOK() *SwarmUnlockOK {
	return &SwarmUnlockOK{}
}

/* SwarmUnlockOK describes a response with status code 200, with default header values.

no error
*/
type SwarmUnlockOK struct {
}

func (o *SwarmUnlockOK) Error() string {
	return fmt.Sprintf("[POST /swarm/unlock][%d] swarmUnlockOK ", 200)
}

func (o *SwarmUnlockOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSwarmUnlockInternalServerError creates a SwarmUnlockInternalServerError with default headers values
func NewSwarmUnlockInternalServerError() *SwarmUnlockInternalServerError {
	return &SwarmUnlockInternalServerError{}
}

/* SwarmUnlockInternalServerError describes a response with status code 500, with default header values.

server error
*/
type SwarmUnlockInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SwarmUnlockInternalServerError) Error() string {
	return fmt.Sprintf("[POST /swarm/unlock][%d] swarmUnlockInternalServerError  %+v", 500, o.Payload)
}
func (o *SwarmUnlockInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwarmUnlockInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSwarmUnlockServiceUnavailable creates a SwarmUnlockServiceUnavailable with default headers values
func NewSwarmUnlockServiceUnavailable() *SwarmUnlockServiceUnavailable {
	return &SwarmUnlockServiceUnavailable{}
}

/* SwarmUnlockServiceUnavailable describes a response with status code 503, with default header values.

node is not part of a swarm
*/
type SwarmUnlockServiceUnavailable struct {
	Payload *models.ErrorResponse
}

func (o *SwarmUnlockServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /swarm/unlock][%d] swarmUnlockServiceUnavailable  %+v", 503, o.Payload)
}
func (o *SwarmUnlockServiceUnavailable) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SwarmUnlockServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*SwarmUnlockBody SwarmUnlockRequest
// Example: {"UnlockKey":"SWMKEY-1-7c37Cc8654o6p38HnroywCi19pllOnGtbdZEgtKxZu8"}
swagger:model SwarmUnlockBody
*/
type SwarmUnlockBody struct {

	// The swarm's unlock key.
	UnlockKey string `json:"UnlockKey,omitempty"`
}

// Validate validates this swarm unlock body
func (o *SwarmUnlockBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this swarm unlock body based on context it is used
func (o *SwarmUnlockBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SwarmUnlockBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SwarmUnlockBody) UnmarshalBinary(b []byte) error {
	var res SwarmUnlockBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
