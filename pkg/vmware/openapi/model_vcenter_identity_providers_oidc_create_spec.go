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

// VcenterIdentityProvidersOidcCreateSpec struct for VcenterIdentityProvidersOidcCreateSpec
type VcenterIdentityProvidersOidcCreateSpec struct {
	// Endpoint to retrieve the provider metadata
	DiscoveryEndpoint string `json:"discovery_endpoint"`
	// Client identifier to connect to the provider
	ClientId string `json:"client_id"`
	// The secret shared between the client and the provider
	ClientSecret string `json:"client_secret"`
	// The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key \"perms\" is supported. The key \"perms\" is used for mapping the \"perms\" claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value.
	ClaimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap `json:"claim_map"`
}

// NewVcenterIdentityProvidersOidcCreateSpec instantiates a new VcenterIdentityProvidersOidcCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIdentityProvidersOidcCreateSpec(discoveryEndpoint string, clientId string, clientSecret string, claimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap) *VcenterIdentityProvidersOidcCreateSpec {
	this := VcenterIdentityProvidersOidcCreateSpec{}
	this.DiscoveryEndpoint = discoveryEndpoint
	this.ClientId = clientId
	this.ClientSecret = clientSecret
	this.ClaimMap = claimMap
	return &this
}

// NewVcenterIdentityProvidersOidcCreateSpecWithDefaults instantiates a new VcenterIdentityProvidersOidcCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIdentityProvidersOidcCreateSpecWithDefaults() *VcenterIdentityProvidersOidcCreateSpec {
	this := VcenterIdentityProvidersOidcCreateSpec{}
	return &this
}

// GetDiscoveryEndpoint returns the DiscoveryEndpoint field value
func (o *VcenterIdentityProvidersOidcCreateSpec) GetDiscoveryEndpoint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DiscoveryEndpoint
}

// GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOidcCreateSpec) GetDiscoveryEndpointOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DiscoveryEndpoint, true
}

// SetDiscoveryEndpoint sets field value
func (o *VcenterIdentityProvidersOidcCreateSpec) SetDiscoveryEndpoint(v string) {
	o.DiscoveryEndpoint = v
}

// GetClientId returns the ClientId field value
func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *VcenterIdentityProvidersOidcCreateSpec) SetClientId(v string) {
	o.ClientId = v
}

// GetClientSecret returns the ClientSecret field value
func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientSecret() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOidcCreateSpec) GetClientSecretOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClientSecret, true
}

// SetClientSecret sets field value
func (o *VcenterIdentityProvidersOidcCreateSpec) SetClientSecret(v string) {
	o.ClientSecret = v
}

// GetClaimMap returns the ClaimMap field value
func (o *VcenterIdentityProvidersOidcCreateSpec) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap {
	if o == nil {
		var ret []VcenterIdentityProvidersOauth2CreateSpecClaimMap
		return ret
	}

	return o.ClaimMap
}

// GetClaimMapOk returns a tuple with the ClaimMap field value
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOidcCreateSpec) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClaimMap, true
}

// SetClaimMap sets field value
func (o *VcenterIdentityProvidersOidcCreateSpec) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap) {
	o.ClaimMap = v
}

func (o VcenterIdentityProvidersOidcCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["discovery_endpoint"] = o.DiscoveryEndpoint
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
	return json.Marshal(toSerialize)
}

type NullableVcenterIdentityProvidersOidcCreateSpec struct {
	value *VcenterIdentityProvidersOidcCreateSpec
	isSet bool
}

func (v NullableVcenterIdentityProvidersOidcCreateSpec) Get() *VcenterIdentityProvidersOidcCreateSpec {
	return v.value
}

func (v *NullableVcenterIdentityProvidersOidcCreateSpec) Set(val *VcenterIdentityProvidersOidcCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIdentityProvidersOidcCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIdentityProvidersOidcCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIdentityProvidersOidcCreateSpec(val *VcenterIdentityProvidersOidcCreateSpec) *NullableVcenterIdentityProvidersOidcCreateSpec {
	return &NullableVcenterIdentityProvidersOidcCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterIdentityProvidersOidcCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIdentityProvidersOidcCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

