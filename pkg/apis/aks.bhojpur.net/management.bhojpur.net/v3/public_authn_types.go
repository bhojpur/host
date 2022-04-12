package v3

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Type string `json:"type"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Token     string `json:"token"`
	ExpiresAt string `json:"expiresAt"`
}

type GenericLogin struct {
	TTLMillis    int64  `json:"ttl,omitempty"`
	Description  string `json:"description,omitempty" bhojpur:"type=string,required"`
	ResponseType string `json:"responseType,omitempty" bhojpur:"type=string,required"` //json or cookie
}

type BasicLogin struct {
	GenericLogin `json:",inline"`
	Username     string `json:"username" bhojpur:"type=string,required"`
	Password     string `json:"password" bhojpur:"type=string,required"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LocalProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GithubProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`

	RedirectURL string `json:"redirectUrl"`
}

type GithubLogin struct {
	GenericLogin `json:",inline"`
	Code         string `json:"code" bhojpur:"type=string,required"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleOAuthProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`

	RedirectURL string `json:"redirectUrl"`
}

type GoogleOauthLogin struct {
	GenericLogin `json:",inline"`
	Code         string `json:"code" bhojpur:"type=string,required"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ActiveDirectoryProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`

	DefaultLoginDomain string `json:"defaultLoginDomain,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureADProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`

	RedirectURL string `json:"redirectUrl"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SamlProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`

	RedirectURL string `json:"redirectUrl"`
}

type AzureADLogin struct {
	GenericLogin `json:",inline"`
	Code         string `json:"code" bhojpur:"type=string,required"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OpenLdapProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type FreeIpaProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`
}

type PingProvider struct {
	SamlProvider `json:",inline"`
}

type ShibbolethProvider struct {
	SamlProvider `json:",inline"`
}

type ADFSProvider struct {
	SamlProvider `json:",inline"`
}

type KeyCloakProvider struct {
	SamlProvider `json:",inline"`
}

type OKTAProvider struct {
	SamlProvider `json:",inline"`
}

type SamlLoginInput struct {
	FinalRedirectURL string `json:"finalRedirectUrl"`
	RequestID        string `json:"requestId"`
	PublicKey        string `json:"publicKey"`
	ResponseType     string `json:"responseType"`
}

type SamlLoginOutput struct {
	IdpRedirectURL string `json:"idpRedirectUrl"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OIDCProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	AuthProvider      `json:",inline"`

	RedirectURL string `json:"redirectUrl"`
}

type OIDCLogin struct {
	GenericLogin `json:",inline"`
	Code         string `json:"code" bhojpur:"type=string,required"`
}

type KeyCloakOIDCProvider struct {
	OIDCProvider `json:",inline"`
}
