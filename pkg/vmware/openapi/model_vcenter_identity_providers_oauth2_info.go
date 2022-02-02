/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VcenterIdentityProvidersOauth2Info struct for VcenterIdentityProvidersOauth2Info
type VcenterIdentityProvidersOauth2Info struct {
	// Authentication/authorization endpoint of the provider
	AuthEndpoint string `json:"auth_endpoint"`
	// Token endpoint of the provider
	TokenEndpoint string `json:"token_endpoint"`
	// Endpoint to retrieve the provider public key for validation
	PublicKeyUri string `json:"public_key_uri"`
	// Client identifier to connect to the provider
	ClientId string `json:"client_id"`
	// The secret shared between the client and the provider
	ClientSecret string `json:"client_secret"`
	// The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key \"perms\" is supported. The key \"perms\" is used for mapping the \"perms\" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value.
	ClaimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap `json:"claim_map"`
	// The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token
	Issuer string `json:"issuer"`
	AuthenticationMethod VcenterIdentityProvidersOauth2AuthenticationMethod `json:"authentication_method"`
	//  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \"?\" is added to the endpoint URL, and combination of each k and each string in the v is added with an \"&\" delimiter. Details:    - If the value contains only one string, then the key is added with \"k=v\".    - If the value is an empty list, then the key is added without a \"=v\".    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value. 
	AuthQueryParams []VcenterIdentityProvidersCreateSpecAuthQueryParams `json:"auth_query_params"`
}

// NewVcenterIdentityProvidersOauth2Info instantiates a new VcenterIdentityProvidersOauth2Info object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIdentityProvidersOauth2Info(authEndpoint string, tokenEndpoint string, publicKeyUri string, clientId string, clientSecret string, claimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap, issuer string, authenticationMethod VcenterIdentityProvidersOauth2AuthenticationMethod, authQueryParams []VcenterIdentityProvidersCreateSpecAuthQueryParams) *VcenterIdentityProvidersOauth2Info {
	this := VcenterIdentityProvidersOauth2Info{}
	this.AuthEndpoint = authEndpoint
	this.TokenEndpoint = tokenEndpoint
	this.PublicKeyUri = publicKeyUri
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.ClaimMap = claimMap
	this.Issuer = issuer
	this.AuthenticationMethod = authenticationMethod
	this.AuthQueryParams = authQueryParams
	return &this
}

// NewVcenterIdentityProvidersOauth2InfoWithDefaults instantiates a new VcenterIdentityProvidersOauth2Info object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIdentityProvidersOauth2InfoWithDefaults() *VcenterIdentityProvidersOauth2Info {
	this := VcenterIdentityProvidersOauth2Info{}
	return &this
}

// GetAuthEndpoint returns the AuthEndpoint field value
func (o *VcenterIdentityProvidersOauth2Info) GetAuthEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AuthEndpoint
}

// GetAuthEndpointOk returns a tuple with the AuthEndpoint field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetAuthEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthEndpoint, true
}

// SetAuthEndpoint sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetAuthEndpoint(v string) {
	o.AuthEndpoint = v
}

// GetTokenEndpoint returns the TokenEndpoint field value
func (o *VcenterIdentityProvidersOauth2Info) GetTokenEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenEndpoint
}

// GetTokenEndpointOk returns a tuple with the TokenEndpoint field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetTokenEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TokenEndpoint, true
}

// SetTokenEndpoint sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetTokenEndpoint(v string) {
	o.TokenEndpoint = v
}

// GetPublicKeyUri returns the PublicKeyUri field value
func (o *VcenterIdentityProvidersOauth2Info) GetPublicKeyUri() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKeyUri
}

// GetPublicKeyUriOk returns a tuple with the PublicKeyUri field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetPublicKeyUriOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PublicKeyUri, true
}

// SetPublicKeyUri sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetPublicKeyUri(v string) {
	o.PublicKeyUri = v
}

// GetClientId returns the ClientId field value
func (o *VcenterIdentityProvidersOauth2Info) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *VcenterIdentityProvidersOauth2Info) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetClientSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetClaimMap returns the ClaimMap field value
func (o *VcenterIdentityProvidersOauth2Info) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap {
	if o == nil {
		var ret []VcenterIdentityProvidersOauth2CreateSpecClaimMap
		return ret
	}

	return o.ClaimMap
}

// GetClaimMapOk returns a tuple with the ClaimMap field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClaimMap, true
}

// SetClaimMap sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap) {
	o.ClaimMap = v
}

// GetIssuer returns the Issuer field value
func (o *VcenterIdentityProvidersOauth2Info) GetIssuer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Issuer
}

// GetIssuerOk returns a tuple with the Issuer field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetIssuerOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Issuer, true
}

// SetIssuer sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetIssuer(v string) {
	o.Issuer = v
}

// GetAuthenticationMethod returns the AuthenticationMethod field value
func (o *VcenterIdentityProvidersOauth2Info) GetAuthenticationMethod() VcenterIdentityProvidersOauth2AuthenticationMethod {
	if o == nil {
		var ret VcenterIdentityProvidersOauth2AuthenticationMethod
		return ret
	}

	return o.AuthenticationMethod
}

// GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetAuthenticationMethodOk() (*VcenterIdentityProvidersOauth2AuthenticationMethod, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthenticationMethod, true
}

// SetAuthenticationMethod sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetAuthenticationMethod(v VcenterIdentityProvidersOauth2AuthenticationMethod) {
	o.AuthenticationMethod = v
}

// GetAuthQueryParams returns the AuthQueryParams field value
func (o *VcenterIdentityProvidersOauth2Info) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams {
	if o == nil {
		var ret []VcenterIdentityProvidersCreateSpecAuthQueryParams
		return ret
	}

	return o.AuthQueryParams
}

// GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2Info) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AuthQueryParams, true
}

// SetAuthQueryParams sets field value
func (o *VcenterIdentityProvidersOauth2Info) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams) {
	o.AuthQueryParams = v
}

func (o VcenterIdentityProvidersOauth2Info) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["auth_endpoint"] = o.AuthEndpoint
	}
	if true {
		toSerialize["token_endpoint"] = o.TokenEndpoint
	}
	if true {
		toSerialize["public_key_uri"] = o.PublicKeyUri
	}
	if true {
		toSerialize["client_id"] = o.ClientId
	}
	if true {
		toSerialize["client_secret"] = o.ClientSecret
	}
	if true {
		toSerialize["claim_map"] = o.ClaimMap
	}
	if true {
		toSerialize["issuer"] = o.Issuer
	}
	if true {
		toSerialize["authentication_method"] = o.AuthenticationMethod
	}
	if true {
		toSerialize["auth_query_params"] = o.AuthQueryParams
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterIdentityProvidersOauth2Info struct {
	value *VcenterIdentityProvidersOauth2Info
	isSet bool
}

func (v NullableVcenterIdentityProvidersOauth2Info) Get() *VcenterIdentityProvidersOauth2Info {
	return v.value
}

func (v *NullableVcenterIdentityProvidersOauth2Info) Set(val *VcenterIdentityProvidersOauth2Info) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIdentityProvidersOauth2Info) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIdentityProvidersOauth2Info) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIdentityProvidersOauth2Info(val *VcenterIdentityProvidersOauth2Info) *NullableVcenterIdentityProvidersOauth2Info {
	return &NullableVcenterIdentityProvidersOauth2Info{value: val, isSet: true}
}

func (v NullableVcenterIdentityProvidersOauth2Info) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIdentityProvidersOauth2Info) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


