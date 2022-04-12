package mapper

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
	"strings"

	"fmt"

	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	"github.com/bhojpur/host/pkg/core/types/definition"
)

type NamespaceReference struct {
	fields      [][]string
	VersionPath string
}

func (n *NamespaceReference) FromInternal(data map[string]interface{}) {
	namespaceID, ok := data["namespaceId"]
	if ok {
		for _, path := range n.fields {
			convert.Transform(data, path, func(input interface{}) interface{} {
				parts := strings.SplitN(convert.ToString(input), ":", 2)
				if len(parts) == 2 {
					return input
				}
				return fmt.Sprintf("%s:%v", namespaceID, input)
			})
		}
	}
}

func (n *NamespaceReference) ToInternal(data map[string]interface{}) error {
	namespaceID, ok := data["namespaceId"]
	for _, path := range n.fields {
		convert.Transform(data, path, func(input interface{}) interface{} {
			parts := strings.SplitN(convert.ToString(input), ":", 2)
			if len(parts) == 2 && (!ok || parts[0] == namespaceID) {
				return parts[1]
			}
			return input
		})
	}

	return nil
}

func (n *NamespaceReference) ModifySchema(schema *types.Schema, schemas *types.Schemas) error {
	_, hasNamespace := schema.ResourceFields["namespaceId"]
	if schema.Version.Path != n.VersionPath || !hasNamespace {
		return nil
	}
	n.fields = traverse(nil, schema, schemas)
	return nil
}

func traverse(prefix []string, schema *types.Schema, schemas *types.Schemas) [][]string {
	var result [][]string

	for name, field := range schema.ResourceFields {
		localPrefix := []string{name}
		subType := field.Type
		if definition.IsArrayType(field.Type) {
			localPrefix = append(localPrefix, "{ARRAY}")
			subType = definition.SubType(field.Type)
		} else if definition.IsMapType(field.Type) {
			localPrefix = append(localPrefix, "{MAP}")
			subType = definition.SubType(field.Type)
		}
		if definition.IsReferenceType(subType) {
			result = appendReference(result, prefix, localPrefix, field, schema, schemas)
			continue
		}

		subSchema := schemas.Schema(&schema.Version, subType)
		if subSchema != nil {
			result = append(result, traverse(append(prefix, localPrefix...), subSchema, schemas)...)
		}
	}

	return result
}

func appendReference(result [][]string, prefix []string, name []string, field types.Field, schema *types.Schema, schemas *types.Schemas) [][]string {
	targetSchema := schemas.Schema(&schema.Version, definition.SubType(field.Type))
	if targetSchema != nil && targetSchema.Scope == types.NamespaceScope {
		result = append(result, append(prefix, name...))
	}
	return result
}
