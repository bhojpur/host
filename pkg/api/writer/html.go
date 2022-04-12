package writer

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
	"encoding/json"
	"strings"

	"github.com/bhojpur/host/pkg/api/builtin"
	"github.com/bhojpur/host/pkg/core/types"
)

const (
	JSURL          = "https://releases.bhojpur.net/api-ui/%API_UI_VERSION%/ui.min.js"
	CSSURL         = "https://releases.bhojpur.net/api-ui/%API_UI_VERSION%/ui.min.css"
	DefaultVersion = "1.1.6"
)

var (
	start = `
<!DOCTYPE html>
<!-- If you are reading this, there is a good chance you would prefer sending an
"Accept: application/json" header and receiving actual JSON responses. -->
<link rel="stylesheet" type="text/css" href="%CSSURL%" />
<script src="%JSURL%"></script>
<script>
var user = "admin";
var curlUser='${BHOJPUR_ACCESS_KEY}:${BHOJPUR_SECRET_KEY}';
var schemas=%SCHEMAS%;
var data =
`
	end = []byte(`</script>
`)
)

type StringGetter func() string

type HTMLResponseWriter struct {
	EncodingResponseWriter
	CSSURL       StringGetter
	JSURL        StringGetter
	APIUIVersion StringGetter
}

func (h *HTMLResponseWriter) start(apiContext *types.APIContext, code int, obj interface{}) {
	AddCommonResponseHeader(apiContext)
	apiContext.Response.Header().Set("content-type", "text/html")
	apiContext.Response.Header().Set("X-Frame-Options", "SAMEORIGIN")
	apiContext.Response.WriteHeader(code)
}

func (h *HTMLResponseWriter) Write(apiContext *types.APIContext, code int, obj interface{}) {
	h.start(apiContext, code, obj)
	schemaSchema := apiContext.Schemas.Schema(&builtin.Version, "schema")
	headerString := start
	if schemaSchema != nil {
		headerString = strings.Replace(headerString, "%SCHEMAS%", jsonEncodeURL(apiContext.URLBuilder.Collection(schemaSchema, apiContext.Version)), 1)
	}
	var jsurl, cssurl string
	if h.CSSURL != nil && h.JSURL != nil && h.CSSURL() != "" && h.JSURL() != "" {
		jsurl = h.JSURL()
		cssurl = h.CSSURL()
	} else if h.APIUIVersion != nil && h.APIUIVersion() != "" {
		jsurl = strings.Replace(JSURL, "%API_UI_VERSION%", h.APIUIVersion(), 1)
		cssurl = strings.Replace(CSSURL, "%API_UI_VERSION%", h.APIUIVersion(), 1)
	} else {
		jsurl = strings.Replace(JSURL, "%API_UI_VERSION%", DefaultVersion, 1)
		cssurl = strings.Replace(CSSURL, "%API_UI_VERSION%", DefaultVersion, 1)
	}
	headerString = strings.Replace(headerString, "%JSURL%", jsurl, 1)
	headerString = strings.Replace(headerString, "%CSSURL%", cssurl, 1)

	apiContext.Response.Write([]byte(headerString))
	h.Body(apiContext, apiContext.Response, obj)
	if schemaSchema != nil {
		apiContext.Response.Write(end)
	}
}

func jsonEncodeURL(str string) string {
	data, _ := json.Marshal(str)
	return string(data)
}
