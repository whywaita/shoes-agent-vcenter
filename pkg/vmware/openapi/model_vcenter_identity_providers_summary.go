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

// VcenterIdentityProvidersSummary struct for VcenterIdentityProvidersSummary
type VcenterIdentityProvidersSummary struct {
	// The identifier of the provider When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.identity.Providers. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.identity.Providers.
	Provider string `json:"provider"`
	// The user friendly name for the provider This field is optional because it was added in a newer version than its parent node.
	Name *string `json:"name,omitempty"`
	ConfigTag VcenterIdentityProvidersConfigType `json:"config_tag"`
	Oauth2 *VcenterIdentityProvidersOauth2Summary `json:"oauth2,omitempty"`
	Oidc *VcenterIdentityProvidersOidcSummary `json:"oidc,omitempty"`
	// Specifies whether the provider is the default provider.
	IsDefault bool `json:"is_default"`
	// Set of fully qualified domain names to trust when federating with this identity provider. Tokens from this identity provider will only be validated if the user belongs to one of these domains, and any domain-qualified groups in the tokens will be filtered to include only those groups that belong to one of these domains. If domainNames is an empty set, domain validation behavior at login with this identity provider will be as follows: the user's domain will be parsed from the User Principal Name (UPN) value that is found in the tokens returned by the identity provider. This domain will then be implicitly trusted and used to filter any groups that are also provided in the tokens. This field is optional because it was added in a newer version than its parent node.
	DomainNames *[]string `json:"domain_names,omitempty"`
	//  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \"?\" is added to the endpoint URL, and combination of each k and each string in the v is added with an \"&\" delimiter. Details:    - If the value contains only one string, then the key is added with \"k=v\".    - If the value is an empty list, then the key is added without a \"=v\".    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  This field is optional because it was added in a newer version than its parent node.
	AuthQueryParams *[]VcenterIdentityProvidersCreateSpecAuthQueryParams `json:"auth_query_params,omitempty"`
}

// NewVcenterIdentityProvidersSummary instantiates a new VcenterIdentityProvidersSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIdentityProvidersSummary(provider string, configTag VcenterIdentityProvidersConfigType, isDefault bool) *VcenterIdentityProvidersSummary {
	this := VcenterIdentityProvidersSummary{}
	this.Provider = provider
	this.ConfigTag = configTag
	this.IsDefault = isDefault
	return &this
}

// NewVcenterIdentityProvidersSummaryWithDefaults instantiates a new VcenterIdentityProvidersSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIdentityProvidersSummaryWithDefaults() *VcenterIdentityProvidersSummary {
	this := VcenterIdentityProvidersSummary{}
	return &this
}

// GetProvider returns the Provider field value
func (o *VcenterIdentityProvidersSummary) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *VcenterIdentityProvidersSummary) SetProvider(v string) {
	o.Provider = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersSummary) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersSummary) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VcenterIdentityProvidersSummary) SetName(v string) {
	o.Name = &v
}

// GetConfigTag returns the ConfigTag field value
func (o *VcenterIdentityProvidersSummary) GetConfigTag() VcenterIdentityProvidersConfigType {
	if o == nil {
		var ret VcenterIdentityProvidersConfigType
		return ret
	}

	return o.ConfigTag
}

// GetConfigTagOk returns a tuple with the ConfigTag field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetConfigTagOk() (*VcenterIdentityProvidersConfigType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigTag, true
}

// SetConfigTag sets field value
func (o *VcenterIdentityProvidersSummary) SetConfigTag(v VcenterIdentityProvidersConfigType) {
	o.ConfigTag = v
}

// GetOauth2 returns the Oauth2 field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersSummary) GetOauth2() VcenterIdentityProvidersOauth2Summary {
	if o == nil || o.Oauth2 == nil {
		var ret VcenterIdentityProvidersOauth2Summary
		return ret
	}
	return *o.Oauth2
}

// GetOauth2Ok returns a tuple with the Oauth2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetOauth2Ok() (*VcenterIdentityProvidersOauth2Summary, bool) {
	if o == nil || o.Oauth2 == nil {
		return nil, false
	}
	return o.Oauth2, true
}

// HasOauth2 returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersSummary) HasOauth2() bool {
	if o != nil && o.Oauth2 != nil {
		return true
	}

	return false
}

// SetOauth2 gets a reference to the given VcenterIdentityProvidersOauth2Summary and assigns it to the Oauth2 field.
func (o *VcenterIdentityProvidersSummary) SetOauth2(v VcenterIdentityProvidersOauth2Summary) {
	o.Oauth2 = &v
}

// GetOidc returns the Oidc field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersSummary) GetOidc() VcenterIdentityProvidersOidcSummary {
	if o == nil || o.Oidc == nil {
		var ret VcenterIdentityProvidersOidcSummary
		return ret
	}
	return *o.Oidc
}

// GetOidcOk returns a tuple with the Oidc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetOidcOk() (*VcenterIdentityProvidersOidcSummary, bool) {
	if o == nil || o.Oidc == nil {
		return nil, false
	}
	return o.Oidc, true
}

// HasOidc returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersSummary) HasOidc() bool {
	if o != nil && o.Oidc != nil {
		return true
	}

	return false
}

// SetOidc gets a reference to the given VcenterIdentityProvidersOidcSummary and assigns it to the Oidc field.
func (o *VcenterIdentityProvidersSummary) SetOidc(v VcenterIdentityProvidersOidcSummary) {
	o.Oidc = &v
}

// GetIsDefault returns the IsDefault field value
func (o *VcenterIdentityProvidersSummary) GetIsDefault() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsDefault
}

// GetIsDefaultOk returns a tuple with the IsDefault field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetIsDefaultOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsDefault, true
}

// SetIsDefault sets field value
func (o *VcenterIdentityProvidersSummary) SetIsDefault(v bool) {
	o.IsDefault = v
}

// GetDomainNames returns the DomainNames field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersSummary) GetDomainNames() []string {
	if o == nil || o.DomainNames == nil {
		var ret []string
		return ret
	}
	return *o.DomainNames
}

// GetDomainNamesOk returns a tuple with the DomainNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetDomainNamesOk() (*[]string, bool) {
	if o == nil || o.DomainNames == nil {
		return nil, false
	}
	return o.DomainNames, true
}

// HasDomainNames returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersSummary) HasDomainNames() bool {
	if o != nil && o.DomainNames != nil {
		return true
	}

	return false
}

// SetDomainNames gets a reference to the given []string and assigns it to the DomainNames field.
func (o *VcenterIdentityProvidersSummary) SetDomainNames(v []string) {
	o.DomainNames = &v
}

// GetAuthQueryParams returns the AuthQueryParams field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersSummary) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams {
	if o == nil || o.AuthQueryParams == nil {
		var ret []VcenterIdentityProvidersCreateSpecAuthQueryParams
		return ret
	}
	return *o.AuthQueryParams
}

// GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersSummary) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool) {
	if o == nil || o.AuthQueryParams == nil {
		return nil, false
	}
	return o.AuthQueryParams, true
}

// HasAuthQueryParams returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersSummary) HasAuthQueryParams() bool {
	if o != nil && o.AuthQueryParams != nil {
		return true
	}

	return false
}

// SetAuthQueryParams gets a reference to the given []VcenterIdentityProvidersCreateSpecAuthQueryParams and assigns it to the AuthQueryParams field.
func (o *VcenterIdentityProvidersSummary) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams) {
	o.AuthQueryParams = &v
}

func (o VcenterIdentityProvidersSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["config_tag"] = o.ConfigTag
	}
	if o.Oauth2 != nil {
		toSerialize["oauth2"] = o.Oauth2
	}
	if o.Oidc != nil {
		toSerialize["oidc"] = o.Oidc
	}
	if true {
		toSerialize["is_default"] = o.IsDefault
	}
	if o.DomainNames != nil {
		toSerialize["domain_names"] = o.DomainNames
	}
	if o.AuthQueryParams != nil {
		toSerialize["auth_query_params"] = o.AuthQueryParams
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterIdentityProvidersSummary struct {
	value *VcenterIdentityProvidersSummary
	isSet bool
}

func (v NullableVcenterIdentityProvidersSummary) Get() *VcenterIdentityProvidersSummary {
	return v.value
}

func (v *NullableVcenterIdentityProvidersSummary) Set(val *VcenterIdentityProvidersSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIdentityProvidersSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIdentityProvidersSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIdentityProvidersSummary(val *VcenterIdentityProvidersSummary) *NullableVcenterIdentityProvidersSummary {
	return &NullableVcenterIdentityProvidersSummary{value: val, isSet: true}
}

func (v NullableVcenterIdentityProvidersSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIdentityProvidersSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

