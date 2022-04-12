package schema

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
	"encoding/base64"
	"fmt"

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
)

type RegistryCredentialMapper struct {
}

func (e RegistryCredentialMapper) FromInternal(data map[string]interface{}) {
}

func (e RegistryCredentialMapper) ToInternal(data map[string]interface{}) error {
	if data == nil {
		return nil
	}

	if data["kind"] != "dockerCredential" {
		return nil
	}

	addAuthInfo(data)

	return nil
}

func addAuthInfo(data map[string]interface{}) error {

	registryMap := convert.ToMapInterface(data["registries"])
	for _, regCred := range registryMap {
		regCredMap := convert.ToMapInterface(regCred)

		username := convert.ToString(regCredMap["username"])
		if username == "" {
			continue
		}
		password := convert.ToString(regCredMap["password"])
		if password == "" {
			continue
		}
		auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))
		regCredMap["auth"] = auth
	}

	return nil
}
func (e RegistryCredentialMapper) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	return nil
}
