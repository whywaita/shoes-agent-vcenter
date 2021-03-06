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

// VcenterIdentityProvidersOauth2CreateSpecClaimMap struct for VcenterIdentityProvidersOauth2CreateSpecClaimMap
type VcenterIdentityProvidersOauth2CreateSpecClaimMap struct {
	Key *string `json:"key,omitempty"`
	Value *[]VcenterIdentityProvidersCreateSpecAuthQueryParams `json:"value,omitempty"`
}

// NewVcenterIdentityProvidersOauth2CreateSpecClaimMap instantiates a new VcenterIdentityProvidersOauth2CreateSpecClaimMap object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIdentityProvidersOauth2CreateSpecClaimMap() *VcenterIdentityProvidersOauth2CreateSpecClaimMap {
	this := VcenterIdentityProvidersOauth2CreateSpecClaimMap{}
	return &this
}

// NewVcenterIdentityProvidersOauth2CreateSpecClaimMapWithDefaults instantiates a new VcenterIdentityProvidersOauth2CreateSpecClaimMap object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIdentityProvidersOauth2CreateSpecClaimMapWithDefaults() *VcenterIdentityProvidersOauth2CreateSpecClaimMap {
	this := VcenterIdentityProvidersOauth2CreateSpecClaimMap{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) GetValue() []VcenterIdentityProvidersCreateSpecAuthQueryParams {
	if o == nil || o.Value == nil {
		var ret []VcenterIdentityProvidersCreateSpecAuthQueryParams
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) GetValueOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []VcenterIdentityProvidersCreateSpecAuthQueryParams and assigns it to the Value field.
func (o *VcenterIdentityProvidersOauth2CreateSpecClaimMap) SetValue(v []VcenterIdentityProvidersCreateSpecAuthQueryParams) {
	o.Value = &v
}

func (o VcenterIdentityProvidersOauth2CreateSpecClaimMap) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap struct {
	value *VcenterIdentityProvidersOauth2CreateSpecClaimMap
	isSet bool
}

func (v NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap) Get() *VcenterIdentityProvidersOauth2CreateSpecClaimMap {
	return v.value
}

func (v *NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap) Set(val *VcenterIdentityProvidersOauth2CreateSpecClaimMap) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIdentityProvidersOauth2CreateSpecClaimMap(val *VcenterIdentityProvidersOauth2CreateSpecClaimMap) *NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap {
	return &NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap{value: val, isSet: true}
}

func (v NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIdentityProvidersOauth2CreateSpecClaimMap) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


