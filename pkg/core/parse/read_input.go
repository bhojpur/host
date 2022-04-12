package parse

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
	"fmt"
	"io"
	"net/http"

	"github.com/bhojpur/host/pkg/core/httperror"
	"k8s.io/apimachinery/pkg/util/yaml"
)

const reqMaxSize = (2 * 1 << 20) + 1

var bodyMethods = map[string]bool{
	http.MethodPut:  true,
	http.MethodPost: true,
}

type Decode func(interface{}) error

func ReadBody(req *http.Request) (map[string]interface{}, error) {
	if !bodyMethods[req.Method] {
		return nil, nil
	}

	decode := getDecoder(req, io.LimitReader(req.Body, maxFormSize))

	data := map[string]interface{}{}
	if err := decode(&data); err != nil {
		return nil, httperror.NewAPIError(httperror.InvalidBodyContent,
			fmt.Sprintf("Failed to parse body: %v", err))
	}

	return data, nil
}

func getDecoder(req *http.Request, reader io.Reader) Decode {
	if req.Header.Get("Content-type") == "application/yaml" {
		return yaml.NewYAMLToJSONDecoder(reader).Decode
	}
	decoder := json.NewDecoder(reader)
	decoder.UseNumber()
	return decoder.Decode
}
