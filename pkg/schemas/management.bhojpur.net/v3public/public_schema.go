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
	"net/http"

	v3 "github.com/bhojpur/host/pkg/apis/management.bhojpur.net/v3"
	"github.com/bhojpur/host/pkg/core/types"
	"github.com/bhojpur/host/pkg/schemas/factory"
)

var (
	PublicVersion = types.APIVersion{
		Version: "v3public",
		Group:   "management.bhojpur.net",
		Path:    "/v3-public",
	}

	PublicSchemas = factory.Schemas(&PublicVersion).
			Init(authProvidersTypes)
)

func authProvidersTypes(schemas *types.Schemas) *types.Schemas {
	return schemas.
		MustImportAndCustomize(&PublicVersion, v3.Token{}, func(schema *types.Schema) {
			// No collection methods causes the store to not create a CRD for it
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{}
		}).
		MustImportAndCustomize(&PublicVersion, v3.AuthToken{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet, http.MethodDelete}
			schema.ResourceMethods = []string{http.MethodGet, http.MethodDelete}
		}).
		MustImportAndCustomize(&PublicVersion, v3.AuthProvider{}, func(schema *types.Schema) {
			schema.CollectionMethods = []string{http.MethodGet}
		}).
		// Local provider
		MustImportAndCustomize(&PublicVersion, v3.LocalProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.BasicLogin{}).
		// Github provider
		MustImportAndCustomize(&PublicVersion, v3.GithubProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "githubLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.GithubLogin{}).
		// Google OAuth provider
		MustImportAndCustomize(&PublicVersion, v3.GoogleOAuthProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "googleOauthLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.GoogleOauthLogin{}).
		// Active Directory provider
		MustImportAndCustomize(&PublicVersion, v3.ActiveDirectoryProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		// Azure AD provider
		MustImportAndCustomize(&PublicVersion, v3.AzureADProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "azureADLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.AzureADLogin{}).
		// Saml provider
		MustImportAndCustomize(&PublicVersion, v3.PingProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "samlLoginInput",
					Output: "samlLoginOutput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImportAndCustomize(&PublicVersion, v3.ADFSProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "samlLoginInput",
					Output: "samlLoginOutput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImportAndCustomize(&PublicVersion, v3.KeyCloakProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "samlLoginInput",
					Output: "samlLoginOutput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImportAndCustomize(&PublicVersion, v3.OKTAProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "samlLoginInput",
					Output: "samlLoginOutput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImportAndCustomize(&PublicVersion, v3.ShibbolethProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "samlLoginInput",
					Output: "samlLoginOutput",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.SamlLoginInput{}).
		MustImport(&PublicVersion, v3.SamlLoginOutput{}).
		// OpenLdap provider
		MustImportAndCustomize(&PublicVersion, v3.OpenLdapProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		// FreeIpa provider
		MustImportAndCustomize(&PublicVersion, v3.FreeIpaProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "basicLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		// OIDC provider
		MustImportAndCustomize(&PublicVersion, v3.OIDCProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "oidcLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.OIDCLogin{}).
		// KeyCloak OIDC provider
		MustImportAndCustomize(&PublicVersion, v3.KeyCloakOIDCProvider{}, func(schema *types.Schema) {
			schema.BaseType = "authProvider"
			schema.ResourceActions = map[string]types.Action{
				"login": {
					Input:  "keyCloakOidcLogin",
					Output: "token",
				},
			}
			schema.CollectionMethods = []string{}
			schema.ResourceMethods = []string{http.MethodGet}
		}).
		MustImport(&PublicVersion, v3.OIDCLogin{})
}
