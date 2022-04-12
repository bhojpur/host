package handler

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

import (
	"errors"
	"net/url"

	"github.com/bhojpur/host/pkg/core/httperror"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/sirupsen/logrus"
)

func ErrorHandler(request *types.APIContext, err error) {
	error := &httperror.APIError{}
	if errors.As(err, &error) {
		if error.Cause != nil {
			url, _ := url.PathUnescape(request.Request.URL.String())
			if url == "" {
				url = request.Request.URL.String()
			}
			logrus.Errorf("API error response %v for %v %v. Cause: %v", error.Code.Status, request.Request.Method,
				url, error.Cause)
		}
	} else {
		logrus.Errorf("Unknown error: %v", err)
		error = &httperror.APIError{
			Code:    httperror.ServerError,
			Message: err.Error(),
		}
	}

	data := toError(error)
	request.WriteResponse(error.Code.Status, data)
}

func toError(apiError *httperror.APIError) map[string]interface{} {
	e := map[string]interface{}{
		"type":    "/meta/schemas/error",
		"status":  apiError.Code.Status,
		"code":    apiError.Code.Code,
		"message": apiError.Message,
	}
	if apiError.FieldName != "" {
		e["fieldName"] = apiError.FieldName
	}

	return e
}
