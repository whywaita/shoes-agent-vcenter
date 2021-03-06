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

// VcenterCryptoManagerHostsKmsProvidersFilterSpec struct for VcenterCryptoManagerHostsKmsProvidersFilterSpec
type VcenterCryptoManagerHostsKmsProvidersFilterSpec struct {
	// Provider identifiers. If unset or empty, the result will not be filtered by provider identifier. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.crypto_manager.kms.provider.
	Providers *[]string `json:"providers,omitempty"`
	// Provider health status. If unset or empty, the result will not be filtered by provider health status.
	Health *[]VcenterCryptoManagerHostsKmsProvidersHealth `json:"health,omitempty"`
	// Provider types. If unset or empty, the result will not be filtered by provider type.
	Types *[]VcenterCryptoManagerHostsKmsProvidersType `json:"types,omitempty"`
}

// NewVcenterCryptoManagerHostsKmsProvidersFilterSpec instantiates a new VcenterCryptoManagerHostsKmsProvidersFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCryptoManagerHostsKmsProvidersFilterSpec() *VcenterCryptoManagerHostsKmsProvidersFilterSpec {
	this := VcenterCryptoManagerHostsKmsProvidersFilterSpec{}
	return &this
}

// NewVcenterCryptoManagerHostsKmsProvidersFilterSpecWithDefaults instantiates a new VcenterCryptoManagerHostsKmsProvidersFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCryptoManagerHostsKmsProvidersFilterSpecWithDefaults() *VcenterCryptoManagerHostsKmsProvidersFilterSpec {
	this := VcenterCryptoManagerHostsKmsProvidersFilterSpec{}
	return &this
}

// GetProviders returns the Providers field value if set, zero value otherwise.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetProviders() []string {
	if o == nil || o.Providers == nil {
		var ret []string
		return ret
	}
	return *o.Providers
}

// GetProvidersOk returns a tuple with the Providers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetProvidersOk() (*[]string, bool) {
	if o == nil || o.Providers == nil {
		return nil, false
	}
	return o.Providers, true
}

// HasProviders returns a boolean if a field has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) HasProviders() bool {
	if o != nil && o.Providers != nil {
		return true
	}

	return false
}

// SetProviders gets a reference to the given []string and assigns it to the Providers field.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) SetProviders(v []string) {
	o.Providers = &v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetHealth() []VcenterCryptoManagerHostsKmsProvidersHealth {
	if o == nil || o.Health == nil {
		var ret []VcenterCryptoManagerHostsKmsProvidersHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetHealthOk() (*[]VcenterCryptoManagerHostsKmsProvidersHealth, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given []VcenterCryptoManagerHostsKmsProvidersHealth and assigns it to the Health field.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) SetHealth(v []VcenterCryptoManagerHostsKmsProvidersHealth) {
	o.Health = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetTypes() []VcenterCryptoManagerHostsKmsProvidersType {
	if o == nil || o.Types == nil {
		var ret []VcenterCryptoManagerHostsKmsProvidersType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetTypesOk() (*[]VcenterCryptoManagerHostsKmsProvidersType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []VcenterCryptoManagerHostsKmsProvidersType and assigns it to the Types field.
func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) SetTypes(v []VcenterCryptoManagerHostsKmsProvidersType) {
	o.Types = &v
}

func (o VcenterCryptoManagerHostsKmsProvidersFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Providers != nil {
		toSerialize["providers"] = o.Providers
	}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec struct {
	value *VcenterCryptoManagerHostsKmsProvidersFilterSpec
	isSet bool
}

func (v NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec) Get() *VcenterCryptoManagerHostsKmsProvidersFilterSpec {
	return v.value
}

func (v *NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec) Set(val *VcenterCryptoManagerHostsKmsProvidersFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCryptoManagerHostsKmsProvidersFilterSpec(val *VcenterCryptoManagerHostsKmsProvidersFilterSpec) *NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec {
	return &NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCryptoManagerHostsKmsProvidersFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


