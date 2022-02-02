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

// VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp struct for VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp
type VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp struct {
	Value VcenterGuestCustomizationSpecsInfo `json:"value"`
}

// NewVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp instantiates a new VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp(value VcenterGuestCustomizationSpecsInfo) *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp {
	this := VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp{}
	this.Value = value
	return &this
}

// NewVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsRespWithDefaults instantiates a new VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsRespWithDefaults() *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp {
	this := VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) GetValue() VcenterGuestCustomizationSpecsInfo {
	if o == nil {
		var ret VcenterGuestCustomizationSpecsInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) GetValueOk() (*VcenterGuestCustomizationSpecsInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) SetValue(v VcenterGuestCustomizationSpecsInfo) {
	o.Value = v
}

func (o VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp struct {
	value *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp
	isSet bool
}

func (v NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) Get() *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp {
	return v.value
}

func (v *NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) Set(val *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp(val *VcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) *NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp {
	return &NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp{value: val, isSet: true}
}

func (v NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCustomizationSpecsGetGuestCustomizationSpecsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

