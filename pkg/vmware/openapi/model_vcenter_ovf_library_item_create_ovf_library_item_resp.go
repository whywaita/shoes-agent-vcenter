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

// VcenterOvfLibraryItemCreateOvfLibraryItemResp struct for VcenterOvfLibraryItemCreateOvfLibraryItemResp
type VcenterOvfLibraryItemCreateOvfLibraryItemResp struct {
	Value VcenterOvfLibraryItemCreateResult `json:"value"`
}

// NewVcenterOvfLibraryItemCreateOvfLibraryItemResp instantiates a new VcenterOvfLibraryItemCreateOvfLibraryItemResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterOvfLibraryItemCreateOvfLibraryItemResp(value VcenterOvfLibraryItemCreateResult) *VcenterOvfLibraryItemCreateOvfLibraryItemResp {
	this := VcenterOvfLibraryItemCreateOvfLibraryItemResp{}
	this.Value = value
	return &this
}

// NewVcenterOvfLibraryItemCreateOvfLibraryItemRespWithDefaults instantiates a new VcenterOvfLibraryItemCreateOvfLibraryItemResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterOvfLibraryItemCreateOvfLibraryItemRespWithDefaults() *VcenterOvfLibraryItemCreateOvfLibraryItemResp {
	this := VcenterOvfLibraryItemCreateOvfLibraryItemResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterOvfLibraryItemCreateOvfLibraryItemResp) GetValue() VcenterOvfLibraryItemCreateResult {
	if o == nil {
		var ret VcenterOvfLibraryItemCreateResult
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemCreateOvfLibraryItemResp) GetValueOk() (*VcenterOvfLibraryItemCreateResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterOvfLibraryItemCreateOvfLibraryItemResp) SetValue(v VcenterOvfLibraryItemCreateResult) {
	o.Value = v
}

func (o VcenterOvfLibraryItemCreateOvfLibraryItemResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp struct {
	value *VcenterOvfLibraryItemCreateOvfLibraryItemResp
	isSet bool
}

func (v NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp) Get() *VcenterOvfLibraryItemCreateOvfLibraryItemResp {
	return v.value
}

func (v *NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp) Set(val *VcenterOvfLibraryItemCreateOvfLibraryItemResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfLibraryItemCreateOvfLibraryItemResp(val *VcenterOvfLibraryItemCreateOvfLibraryItemResp) *NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp {
	return &NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp{value: val, isSet: true}
}

func (v NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfLibraryItemCreateOvfLibraryItemResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


