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
	"github.com/bhojpur/host/pkg/core/condition"
	"github.com/bhojpur/host/pkg/core/types"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const UserConditionInitialRolesPopulated condition.Cond = "InitialRolesPopulated"

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Token           string            `json:"token" bhojpur:"writeOnly,noupdate"`
	UserPrincipal   Principal         `json:"userPrincipal" bhojpur:"type=reference[principal]"`
	GroupPrincipals []Principal       `json:"groupPrincipals" bhojpur:"type=array[reference[principal]]"`
	ProviderInfo    map[string]string `json:"providerInfo,omitempty"`
	UserID          string            `json:"userId" bhojpur:"type=reference[user]"`
	AuthProvider    string            `json:"authProvider"`
	TTLMillis       int64             `json:"ttl"`
	LastUpdateTime  string            `json:"lastUpdateTime"`
	IsDerived       bool              `json:"isDerived"`
	Description     string            `json:"description"`
	Expired         bool              `json:"expired"`
	ExpiresAt       string            `json:"expiresAt"`
	Current         bool              `json:"current"`
	ClusterName     string            `json:"clusterName,omitempty" bhojpur:"noupdate,type=reference[cluster]"`
	Enabled         *bool             `json:"enabled,omitempty" bhojpur:"default=true"`
}

func (t *Token) ObjClusterName() string {
	return t.ClusterName
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName        string     `json:"displayName,omitempty"`
	Description        string     `json:"description"`
	Username           string     `json:"username,omitempty"`
	Password           string     `json:"password,omitempty" bhojpur:"writeOnly,noupdate"`
	MustChangePassword bool       `json:"mustChangePassword,omitempty"`
	PrincipalIDs       []string   `json:"principalIds,omitempty" bhojpur:"type=array[reference[principal]]"`
	Me                 bool       `json:"me,omitempty" bhojpur:"nocreate,noupdate"`
	Enabled            *bool      `json:"enabled,omitempty" bhojpur:"default=true"`
	Spec               UserSpec   `json:"spec,omitempty"`
	Status             UserStatus `json:"status"`
}

type UserStatus struct {
	Conditions []UserCondition `json:"conditions"`
}

type UserCondition struct {
	// Type of user condition.
	Type string `json:"type"`
	// Status of the condition, one of True, False, Unknown.
	Status v1.ConditionStatus `json:"status"`
	// The last time this condition was updated.
	LastUpdateTime string `json:"lastUpdateTime,omitempty"`
	// Last time the condition transitioned from one status to another.
	LastTransitionTime string `json:"lastTransitionTime,omitempty"`
	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`
	// Human-readable message indicating details about last transition
	Message string `json:"message,omitempty"`
}

type UserSpec struct{}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// UserAttribute will have a CRD (and controller) generated for it, but will not be exposed in the API.
type UserAttribute struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	UserName        string
	GroupPrincipals map[string]Principals // the value is a []Principal, but code generator cannot handle slice as a value
	LastRefresh     string
	NeedsRefresh    bool
	ExtraByProvider map[string]map[string][]string // extra information for the user to print in audit logs, stored per authProvider. example: map[openldap:map[principalid:[openldap_user://uid=testuser1,ou=dev,dc=us-west-2,dc=compute,dc=internal]]]
}

type Principals struct {
	Items []Principal
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Group struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName string `json:"displayName,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GroupMember struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	GroupName   string `json:"groupName,omitempty" bhojpur:"type=reference[group]"`
	PrincipalID string `json:"principalId,omitempty" bhojpur:"type=reference[principal]"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Principal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	DisplayName    string            `json:"displayName,omitempty"`
	LoginName      string            `json:"loginName,omitempty"`
	ProfilePicture string            `json:"profilePicture,omitempty"`
	ProfileURL     string            `json:"profileURL,omitempty"`
	PrincipalType  string            `json:"principalType,omitempty"`
	Me             bool              `json:"me,omitempty"`
	MemberOf       bool              `json:"memberOf,omitempty"`
	Provider       string            `json:"provider,omitempty"`
	ExtraInfo      map[string]string `json:"extraInfo,omitempty"`
}

type SearchPrincipalsInput struct {
	Name          string `json:"name" bhojpur:"type=string,required,notnullable"`
	PrincipalType string `json:"principalType,omitempty" bhojpur:"type=enum,options=user|group"`
}

type ChangePasswordInput struct {
	CurrentPassword string `json:"currentPassword" bhojpur:"type=string,required"`
	NewPassword     string `json:"newPassword" bhojpur:"type=string,required"`
}

type SetPasswordInput struct {
	NewPassword string `json:"newPassword" bhojpur:"type=string,required"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AuthConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Type                string   `json:"type" bhojpur:"noupdate"`
	Enabled             bool     `json:"enabled,omitempty"`
	AccessMode          string   `json:"accessMode,omitempty" bhojpur:"required,notnullable,type=enum,options=required|restricted|unrestricted"`
	AllowedPrincipalIDs []string `json:"allowedPrincipalIds,omitempty" bhojpur:"type=array[reference[principal]]"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SamlToken struct {
	types.Namespaced
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Token     string `json:"token" bhojpur:"writeOnly,noupdate"`
	ExpiresAt string `json:"expiresAt"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LocalConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`
}

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GithubConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`

	Hostname     string `json:"hostname,omitempty" bhojpur:"default=github.com" bhojpur:"required"`
	TLS          bool   `json:"tls,omitempty" bhojpur:"notnullable,default=true" bhojpur:"required"`
	ClientID     string `json:"clientId,omitempty" bhojpur:"required"`
	ClientSecret string `json:"clientSecret,omitempty" bhojpur:"required,type=password"`

	// AdditionalClientIDs is a map of clientID to client secrets
	AdditionalClientIDs map[string]string `json:"additionalClientIds,omitempty" bhojpur:"nocreate,noupdate"`
	HostnameToClientID  map[string]string `json:"hostnameToClientId,omitempty" bhojpur:"nocreate,noupdate"`
}

type GithubConfigTestOutput struct {
	RedirectURL string `json:"redirectUrl"`
}

type GithubConfigApplyInput struct {
	GithubConfig GithubConfig `json:"githubConfig,omitempty"`
	Code         string       `json:"code,omitempty"`
	Enabled      bool         `json:"enabled,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type GoogleOauthConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`

	OauthCredential              string `json:"oauthCredential,omitempty" bhojpur:"required,type=password,notnullable"`
	ServiceAccountCredential     string `json:"serviceAccountCredential,omitempty" bhojpur:"required,type=password,notnullable"`
	AdminEmail                   string `json:"adminEmail,omitempty" bhojpur:"required,notnullable"`
	Hostname                     string `json:"hostname,omitempty" bhojpur:"required,notnullable,noupdate"`
	UserInfoEndpoint             string `json:"userInfoEndpoint" bhojpur:"default=https://openidconnect.googleapis.com/v1/userinfo,required,notnullable"`
	NestedGroupMembershipEnabled bool   `json:"nestedGroupMembershipEnabled"    bhojpur:"default=false"`
}

type GoogleOauthConfigTestOutput struct {
	RedirectURL string `json:"redirectUrl"`
}

type GoogleOauthConfigApplyInput struct {
	GoogleOauthConfig GoogleOauthConfig `json:"googleOauthConfig,omitempty"`
	Code              string            `json:"code,omitempty"`
	Enabled           bool              `json:"enabled,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzureADConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`

	Endpoint          string `json:"endpoint,omitempty" bhojpur:"default=https://login.microsoftonline.com/,required,notnullable"`
	GraphEndpoint     string `json:"graphEndpoint,omitempty" bhojpur:"required,notnullable"`
	TokenEndpoint     string `json:"tokenEndpoint,omitempty" bhojpur:"required,notnullable"`
	AuthEndpoint      string `json:"authEndpoint,omitempty" bhojpur:"required,notnullable"`
	TenantID          string `json:"tenantId,omitempty" bhojpur:"required,notnullable"`
	ApplicationID     string `json:"applicationId,omitempty" bhojpur:"required,notnullable"`
	ApplicationSecret string `json:"applicationSecret,omitempty" bhojpur:"required,type=password"`
	BhojpurURL        string `json:"bhojpurUrl,omitempty" bhojpur:"required,notnullable"`
}

type AzureADConfigTestOutput struct {
	RedirectURL string `json:"redirectUrl"`
}

type AzureADConfigApplyInput struct {
	Config AzureADConfig `json:"config,omitempty"`
	Code   string        `json:"code,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ActiveDirectoryConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`

	Servers                      []string `json:"servers,omitempty"                     bhojpur:"type=array[string],required"`
	Port                         int64    `json:"port,omitempty"                        bhojpur:"default=389"`
	TLS                          bool     `json:"tls,omitempty"                         bhojpur:"default=false"`
	StartTLS                     bool     `json:"starttls,omitempty"                    bhojpur:"default=false"`
	Certificate                  string   `json:"certificate,omitempty"`
	DefaultLoginDomain           string   `json:"defaultLoginDomain,omitempty"`
	ServiceAccountUsername       string   `json:"serviceAccountUsername,omitempty"      bhojpur:"required"`
	ServiceAccountPassword       string   `json:"serviceAccountPassword,omitempty"      bhojpur:"type=password,required"`
	UserDisabledBitMask          int64    `json:"userDisabledBitMask,omitempty"         bhojpur:"default=2"`
	UserSearchBase               string   `json:"userSearchBase,omitempty"              bhojpur:"required"`
	UserSearchAttribute          string   `json:"userSearchAttribute,omitempty"         bhojpur:"default=sAMAccountName|sn|givenName,required"`
	UserSearchFilter             string   `json:"userSearchFilter,omitempty"`
	UserLoginAttribute           string   `json:"userLoginAttribute,omitempty"          bhojpur:"default=sAMAccountName,required"`
	UserObjectClass              string   `json:"userObjectClass,omitempty"             bhojpur:"default=person,required"`
	UserNameAttribute            string   `json:"userNameAttribute,omitempty"           bhojpur:"default=name,required"`
	UserEnabledAttribute         string   `json:"userEnabledAttribute,omitempty"        bhojpur:"default=userAccountControl,required"`
	GroupSearchBase              string   `json:"groupSearchBase,omitempty"`
	GroupSearchAttribute         string   `json:"groupSearchAttribute,omitempty"        bhojpur:"default=sAMAccountName,required"`
	GroupSearchFilter            string   `json:"groupSearchFilter,omitempty"`
	GroupObjectClass             string   `json:"groupObjectClass,omitempty"            bhojpur:"default=group,required"`
	GroupNameAttribute           string   `json:"groupNameAttribute,omitempty"          bhojpur:"default=name,required"`
	GroupDNAttribute             string   `json:"groupDNAttribute,omitempty"            bhojpur:"default=distinguishedName,required"`
	GroupMemberUserAttribute     string   `json:"groupMemberUserAttribute,omitempty"    bhojpur:"default=distinguishedName,required"`
	GroupMemberMappingAttribute  string   `json:"groupMemberMappingAttribute,omitempty" bhojpur:"default=member,required"`
	ConnectionTimeout            int64    `json:"connectionTimeout,omitempty"           bhojpur:"default=5000,notnullable,required"`
	NestedGroupMembershipEnabled *bool    `json:"nestedGroupMembershipEnabled,omitempty" bhojpur:"default=false"`
}

type ActiveDirectoryTestAndApplyInput struct {
	ActiveDirectoryConfig ActiveDirectoryConfig `json:"activeDirectoryConfig,omitempty"`
	Username              string                `json:"username"`
	Password              string                `json:"password"`
	Enabled               bool                  `json:"enabled,omitempty"`
}

type LdapFields struct {
	Servers                         []string `json:"servers,omitempty"                         bhojpur:"type=array[string],notnullable,required"`
	Port                            int64    `json:"port,omitempty"                            bhojpur:"default=389,notnullable,required"`
	TLS                             bool     `json:"tls,omitempty"                             bhojpur:"default=false,notnullable,required"`
	StartTLS                        bool     `json:"starttls,omitempty"                        bhojpur:"default=false"`
	Certificate                     string   `json:"certificate,omitempty"`
	ServiceAccountDistinguishedName string   `json:"serviceAccountDistinguishedName,omitempty" bhojpur:"required"`
	ServiceAccountPassword          string   `json:"serviceAccountPassword,omitempty"          bhojpur:"type=password,required"`
	UserDisabledBitMask             int64    `json:"userDisabledBitMask,omitempty"`
	UserSearchBase                  string   `json:"userSearchBase,omitempty"                  bhojpur:"notnullable,required"`
	UserSearchAttribute             string   `json:"userSearchAttribute,omitempty"             bhojpur:"default=uid|sn|givenName,notnullable,required"`
	UserSearchFilter                string   `json:"userSearchFilter,omitempty"`
	UserLoginAttribute              string   `json:"userLoginAttribute,omitempty"              bhojpur:"default=uid,notnullable,required"`
	UserObjectClass                 string   `json:"userObjectClass,omitempty"                 bhojpur:"default=inetOrgPerson,notnullable,required"`
	UserNameAttribute               string   `json:"userNameAttribute,omitempty"               bhojpur:"default=cn,notnullable,required"`
	UserMemberAttribute             string   `json:"userMemberAttribute,omitempty"             bhojpur:"default=memberOf,notnullable,required"`
	UserEnabledAttribute            string   `json:"userEnabledAttribute,omitempty"`
	GroupSearchBase                 string   `json:"groupSearchBase,omitempty"`
	GroupSearchAttribute            string   `json:"groupSearchAttribute,omitempty"            bhojpur:"default=cn,notnullable,required"`
	GroupSearchFilter               string   `json:"groupSearchFilter,omitempty"`
	GroupObjectClass                string   `json:"groupObjectClass,omitempty"                bhojpur:"default=groupOfNames,notnullable,required"`
	GroupNameAttribute              string   `json:"groupNameAttribute,omitempty"              bhojpur:"default=cn,notnullable,required"`
	GroupDNAttribute                string   `json:"groupDNAttribute,omitempty"                bhojpur:"default=entryDN,notnullable"`
	GroupMemberUserAttribute        string   `json:"groupMemberUserAttribute,omitempty"        bhojpur:"default=entryDN,notnullable"`
	GroupMemberMappingAttribute     string   `json:"groupMemberMappingAttribute,omitempty"     bhojpur:"default=member,notnullable,required"`
	ConnectionTimeout               int64    `json:"connectionTimeout,omitempty"               bhojpur:"default=5000,notnullable,required"`
	NestedGroupMembershipEnabled    bool     `json:"nestedGroupMembershipEnabled"              bhojpur:"default=false"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type LdapConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`
	LdapFields `json:",inline" mapstructure:",squash"`
}

type LdapTestAndApplyInput struct {
	LdapConfig `json:"ldapConfig,omitempty"`
	Username   string `json:"username"`
	Password   string `json:"password" bhojpur:"type=password,required"`
}

type OpenLdapConfig struct {
	LdapConfig `json:",inline" mapstructure:",squash"`
}

type OpenLdapTestAndApplyInput struct {
	LdapTestAndApplyInput `json:",inline" mapstructure:",squash"`
}

type FreeIpaConfig struct {
	LdapConfig `json:",inline" mapstructure:",squash"`
}

type FreeIpaTestAndApplyInput struct {
	LdapTestAndApplyInput `json:",inline" mapstructure:",squash"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type SamlConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`

	IDPMetadataContent string `json:"idpMetadataContent" bhojpur:"required"`
	SpCert             string `json:"spCert"             bhojpur:"required"`
	SpKey              string `json:"spKey"              bhojpur:"required,type=password"`
	GroupsField        string `json:"groupsField"        bhojpur:"required"`
	DisplayNameField   string `json:"displayNameField"   bhojpur:"required"`
	UserNameField      string `json:"userNameField"      bhojpur:"required"`
	UIDField           string `json:"uidField"           bhojpur:"required"`
	BhojpurAPIHost     string `json:"bhojpurApiHost"     bhojpur:"required"`
	EntityID           string `json:"entityID"`
}

type SamlConfigTestInput struct {
	FinalRedirectURL string `json:"finalRedirectUrl"`
}

type SamlConfigTestOutput struct {
	IdpRedirectURL string `json:"idpRedirectUrl"`
}

type PingConfig struct {
	SamlConfig `json:",inline" mapstructure:",squash"`
}

type ADFSConfig struct {
	SamlConfig `json:",inline" mapstructure:",squash"`
}

type KeyCloakConfig struct {
	SamlConfig `json:",inline" mapstructure:",squash"`
}

type OKTAConfig struct {
	SamlConfig `json:",inline" mapstructure:",squash"`
}

type ShibbolethConfig struct {
	SamlConfig     `json:",inline" mapstructure:",squash"`
	OpenLdapConfig LdapFields `json:"openLdapConfig" mapstructure:",squash"`
}

type AuthSystemImages struct {
	KubeAPIAuth string `json:"kubeAPIAuth,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type OIDCConfig struct {
	AuthConfig `json:",inline" mapstructure:",squash"`

	ClientID           string `json:"clientId" bhojpur:"required"`
	ClientSecret       string `json:"clientSecret,omitempty" bhojpur:"required,type=password"`
	Scopes             string `json:"scope", bhojpur:"required,notnullable"`
	AuthEndpoint       string `json:"authEndpoint,omitempty" bhojpur:"required,notnullable"`
	Issuer             string `json:"issuer" bhojpur:"required,notnullable"`
	Certificate        string `json:"certificate,omitempty"`
	PrivateKey         string `json:"privateKey" bhojpur:"type=password"`
	BhojpurURL         string `json:"bhojpurUrl" bhojpur:"required,notnullable"`
	GroupSearchEnabled *bool  `json:"groupSearchEnabled"`
}

type OIDCTestOutput struct {
	RedirectURL string `json:"redirectUrl"`
}

type OIDCApplyInput struct {
	OIDCConfig OIDCConfig `json:"oidcConfig,omitempty"`
	Code       string     `json:"code,omitempty"`
	Enabled    bool       `json:"enabled,omitempty"`
}

type KeyCloakOIDCConfig struct {
	OIDCConfig `json:",inline" mapstructure:",squash"`
}
