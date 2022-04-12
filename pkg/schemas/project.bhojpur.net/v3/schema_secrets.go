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
	v3 "github.com/bhojpur/host/pkg/apis/project.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/core/types/convert"
	m "github.com/bhojpur/host/pkg/core/types/mapper"
	v1 "k8s.io/api/core/v1"
)

func secretTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		AddMapperForType(&Version, v1.Secret{},
			&m.AnnotationField{Field: "description"},
			m.AnnotationField{Field: "projectId", IgnoreDefinition: true},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/service-account-token",
				Value: "serviceAccountToken",
			},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/dockercfg",
				Value: "dockerCredential",
			},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/dockerconfigjson",
				Value: "dockerCredential",
			},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/basic-auth",
				Value: "basicAuth",
			},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/ssh-auth",
				Value: "sshAuth",
			},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/ssh-auth",
				Value: "sshAuth",
			},
			m.SetValue{
				Field: "type",
				IfEq:  "kubernetes.io/tls",
				Value: "certificate",
			},
			&m.Move{From: "type", To: "kind"},
			m.Condition{
				Field: "kind",
				Value: "sshAuth",
				Mapper: types.Mappers{
					m.UntypedMove{
						From: "data/ssh-privatekey",
						To:   "privateKey",
					},
					m.Base64{
						Field:            "privateKey",
						IgnoreDefinition: true,
					},
					m.SetValue{
						Field:            "type",
						Value:            "sshAuth",
						IgnoreDefinition: true,
					},
					m.AnnotationField{Field: "fingerprint", IgnoreDefinition: true},
				},
			},
			m.Condition{
				Field: "kind",
				Value: "basicAuth",
				Mapper: types.Mappers{
					m.UntypedMove{
						From: "data/username",
						To:   "username",
					},
					m.UntypedMove{
						From: "data/password",
						To:   "password",
					},
					m.Base64{
						Field:            "username",
						IgnoreDefinition: true,
					},
					m.Base64{
						Field:            "password",
						IgnoreDefinition: true,
					},
					m.SetValue{
						Field:            "type",
						Value:            "basicAuth",
						IgnoreDefinition: true,
					},
				},
			},
			m.Condition{
				Field: "kind",
				Value: "certificate",
				Mapper: types.Mappers{
					m.UntypedMove{
						From: "data/tls.crt",
						To:   "certs",
					},
					m.UntypedMove{
						From: "data/tls.key",
						To:   "key",
					},
					m.Base64{
						Field:            "certs",
						IgnoreDefinition: true,
					},
					m.Base64{
						Field:            "key",
						IgnoreDefinition: true,
					},
					m.AnnotationField{Field: "certFingerprint", IgnoreDefinition: true},
					m.AnnotationField{Field: "cn", IgnoreDefinition: true},
					m.AnnotationField{Field: "version", IgnoreDefinition: true},
					m.AnnotationField{Field: "issuer", IgnoreDefinition: true},
					m.AnnotationField{Field: "issuedAt", IgnoreDefinition: true},
					m.AnnotationField{Field: "expiresAt", IgnoreDefinition: true},
					m.AnnotationField{Field: "algorithm", IgnoreDefinition: true},
					m.AnnotationField{Field: "serialNumber", IgnoreDefinition: true},
					m.AnnotationField{Field: "keySize", IgnoreDefinition: true},
					m.AnnotationField{Field: "subjectAlternativeNames", IgnoreDefinition: true, List: true},
					m.SetValue{
						Field:            "type",
						Value:            "certificate",
						IgnoreDefinition: true,
					},
				},
			},
			m.Condition{
				Field: "kind",
				Value: "dockerCredential",
				Mapper: types.Mappers{
					m.Base64{
						Field:            "data/.dockercfg",
						IgnoreDefinition: true,
					},
					m.JSONEncode{
						Field:            "data/.dockercfg",
						IgnoreDefinition: true,
					},
					m.UntypedMove{
						From: "data/.dockercfg",
						To:   "registries",
					},
					m.Base64{
						Field:            "data/.dockerconfigjson",
						IgnoreDefinition: true,
					},
					m.JSONEncode{
						Field:            "data/.dockerconfigjson",
						IgnoreDefinition: true,
					},
					m.UntypedMove{
						From: "data/.dockerconfigjson/auths",
						To:   "registries",
					},
					m.SetValue{
						Field:            "type",
						Value:            "dockerCredential",
						IgnoreDefinition: true,
					},
				},
			},
			m.Condition{
				Field: "kind",
				Value: "serviceAccountToken",
				Mapper: types.Mappers{
					m.UntypedMove{
						From:      "annotations!kubernetes.io/service-account.name",
						To:        "accountName",
						Separator: "!",
					},
					m.UntypedMove{
						From:      "annotations!kubernetes.io/service-account.uid",
						To:        "accountUid",
						Separator: "!",
					},
					m.UntypedMove{
						From: "data/ca.crt",
						To:   "caCrt",
					},
					m.UntypedMove{
						From: "data/token",
						To:   "token",
					},
					m.Base64{
						Field:            "caCrt",
						IgnoreDefinition: true,
					},
					m.Base64{
						Field:            "token",
						IgnoreDefinition: true,
					},
					m.SetValue{
						Field:            "type",
						Value:            "serviceAccountToken",
						IgnoreDefinition: true,
					},
				},
			},
		).
		AddMapperForType(&Version, v1.Secret{}, RegistryCredentialMapper{}).
		MustImportAndCustomize(&Version, v1.Secret{}, func(schema *types.Schema) {
			schema.MustCustomizeField("kind", func(f types.Field) types.Field {
				f.Options = []string{
					"Opaque",
					"serviceAccountToken",
					"dockerCredential",
					"basicAuth",
					"sshAuth",
					"certificate",
				}
				return f
			})
			schema.MustCustomizeField("name", func(field types.Field) types.Field {
				field.Type = "hostname"
				field.Nullable = false
				field.Required = true
				return field
			})
		}, projectOverride{}, struct {
			Description string `json:"description"`
		}{}).
		Init(func(schemas *types.Schemas) *types.Schemas {
			return addSecretSubtypes(schemas,
				v3.ServiceAccountToken{},
				v3.DockerCredential{},
				v3.Certificate{},
				v3.BasicAuth{},
				v3.SSHAuth{})
		})
}

func addSecretSubtypes(schemas *types.Schemas, objs ...interface{}) *types.Schemas {
	namespaced := []string{"secret"}

	for _, obj := range objs {
		schemas.MustImportAndCustomize(&Version, obj, func(schema *types.Schema) {
			schema.BaseType = "secret"
			schema.Mapper = schemas.Schema(&Version, "secret").Mapper
			namespaced = append(namespaced, schema.ID)
		}, projectOverride{})
	}

	for _, name := range namespaced {
		baseSchema := schemas.Schema(&Version, name)

		// Make non-namespaced have namespaceId not required
		newFields := map[string]types.Field{}
		for name, field := range baseSchema.ResourceFields {
			if name == "namespaceId" {
				field.Required = false
			}
			newFields[name] = field
		}

		schema := *baseSchema
		schema.ID = "namespaced" + convert.Capitalize(schema.ID)
		schema.PluralName = "namespaced" + convert.Capitalize(schema.PluralName)
		schema.CodeName = "Namespaced" + schema.CodeName
		schema.CodeNamePlural = "Namespaced" + schema.CodeNamePlural
		schema.BaseType = "namespacedSecret"
		schemas.AddSchema(schema)

		baseSchema.ResourceFields = newFields
	}

	return schemas
}
