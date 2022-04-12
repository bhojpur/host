package api

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
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/bhojpur/host/pkg/core/httperror"
	"github.com/bhojpur/host/pkg/core/parse"
	"github.com/bhojpur/host/pkg/core/types"
)

const (
	csrfCookie = "CSRF"
	csrfHeader = "X-API-CSRF"
)

func ValidateAction(request *types.APIContext) (*types.Action, error) {
	if request.Action == "" || request.Link != "" || request.Method != http.MethodPost {
		return nil, nil
	}

	actions := request.Schema.CollectionActions
	if request.ID != "" {
		actions = request.Schema.ResourceActions
	}

	action, ok := actions[request.Action]
	if !ok {
		return nil, httperror.NewAPIError(httperror.InvalidAction, fmt.Sprintf("Invalid action: %s", request.Action))
	}

	if request.ID != "" && request.ReferenceValidator != nil {
		resource := request.ReferenceValidator.Lookup(request.Type, request.ID)
		if resource == nil {
			return nil, httperror.NewAPIError(httperror.NotFound, fmt.Sprintf("Failed to find type: %s id: %s", request.Type, request.ID))
		}

		if _, ok := resource.Actions[request.Action]; !ok {
			return nil, httperror.NewAPIError(httperror.InvalidAction, fmt.Sprintf("Invalid action: %s", request.Action))
		}
	}

	return &action, nil
}

func CheckCSRF(apiContext *types.APIContext) error {
	if !parse.IsBrowser(apiContext.Request, false) {
		return nil
	}

	cookie, err := apiContext.Request.Cookie(csrfCookie)
	if err == http.ErrNoCookie {
		bytes := make([]byte, 5)
		_, err := rand.Read(bytes)
		if err != nil {
			return httperror.WrapAPIError(err, httperror.ServerError, "Failed in CSRF processing")
		}

		cookie = &http.Cookie{
			Name:   csrfCookie,
			Value:  hex.EncodeToString(bytes),
			Path:   "/",
			Secure: true,
		}

		http.SetCookie(apiContext.Response, cookie)
	} else if err != nil {
		return httperror.NewAPIError(httperror.InvalidCSRFToken, "Failed to parse cookies")
	}

	// Not an else-if, because this should happen even if there was no cookie to begin with.
	if apiContext.Method != http.MethodGet {
		/*
		 * Very important to use apiContext.Method and not apiContext.Request.Method. The client can override the HTTP method with _method
		 */
		if cookie.Value == apiContext.Request.Header.Get(csrfHeader) {
			// Good
		} else if cookie.Value == apiContext.Request.URL.Query().Get(csrfCookie) {
			// Good
		} else {
			return httperror.NewAPIError(httperror.InvalidCSRFToken, "Invalid CSRF token")
		}
	}

	return nil
}
