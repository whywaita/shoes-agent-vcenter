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

// VcenterIdentityProvidersCreateSpecAuthQueryParams struct for VcenterIdentityProvidersCreateSpecAuthQueryParams
type VcenterIdentityProvidersCreateSpecAuthQueryParams struct {
	Key *string `json:"key,omitempty"`
	Value *[]string `json:"value,omitempty"`
}

// NewVcenterIdentityProvidersCreateSpecAuthQueryParams instantiates a new VcenterIdentityProvidersCreateSpecAuthQueryParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIdentityProvidersCreateSpecAuthQueryParams() *VcenterIdentityProvidersCreateSpecAuthQueryParams {
	this := VcenterIdentityProvidersCreateSpecAuthQueryParams{}
	return &this
}

// NewVcenterIdentityProvidersCreateSpecAuthQueryParamsWithDefaults instantiates a new VcenterIdentityProvidersCreateSpecAuthQueryParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIdentityProvidersCreateSpecAuthQueryParamsWithDefaults() *VcenterIdentityProvidersCreateSpecAuthQueryParams {
	this := VcenterIdentityProvidersCreateSpecAuthQueryParams{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) GetValue() []string {
	if o == nil || o.Value == nil {
		var ret []string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) GetValueOk() (*[]string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given []string and assigns it to the Value field.
func (o *VcenterIdentityProvidersCreateSpecAuthQueryParams) SetValue(v []string) {
	o.Value = &v
}

func (o VcenterIdentityProvidersCreateSpecAuthQueryParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterIdentityProvidersCreateSpecAuthQueryParams struct {
	value *VcenterIdentityProvidersCreateSpecAuthQueryParams
	isSet bool
}

func (v NullableVcenterIdentityProvidersCreateSpecAuthQueryParams) Get() *VcenterIdentityProvidersCreateSpecAuthQueryParams {
	return v.value
}

func (v *NullableVcenterIdentityProvidersCreateSpecAuthQueryParams) Set(val *VcenterIdentityProvidersCreateSpecAuthQueryParams) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIdentityProvidersCreateSpecAuthQueryParams) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIdentityProvidersCreateSpecAuthQueryParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIdentityProvidersCreateSpecAuthQueryParams(val *VcenterIdentityProvidersCreateSpecAuthQueryParams) *NullableVcenterIdentityProvidersCreateSpecAuthQueryParams {
	return &NullableVcenterIdentityProvidersCreateSpecAuthQueryParams{value: val, isSet: true}
}

func (v NullableVcenterIdentityProvidersCreateSpecAuthQueryParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIdentityProvidersCreateSpecAuthQueryParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


