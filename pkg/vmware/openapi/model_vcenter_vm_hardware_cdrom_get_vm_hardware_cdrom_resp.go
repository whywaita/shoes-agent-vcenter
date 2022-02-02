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

// VcenterVmHardwareCdromGetVmHardwareCdromResp struct for VcenterVmHardwareCdromGetVmHardwareCdromResp
type VcenterVmHardwareCdromGetVmHardwareCdromResp struct {
	Value VcenterVmHardwareCdromInfo `json:"value"`
}

// NewVcenterVmHardwareCdromGetVmHardwareCdromResp instantiates a new VcenterVmHardwareCdromGetVmHardwareCdromResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareCdromGetVmHardwareCdromResp(value VcenterVmHardwareCdromInfo) *VcenterVmHardwareCdromGetVmHardwareCdromResp {
	this := VcenterVmHardwareCdromGetVmHardwareCdromResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareCdromGetVmHardwareCdromRespWithDefaults instantiates a new VcenterVmHardwareCdromGetVmHardwareCdromResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareCdromGetVmHardwareCdromRespWithDefaults() *VcenterVmHardwareCdromGetVmHardwareCdromResp {
	this := VcenterVmHardwareCdromGetVmHardwareCdromResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareCdromGetVmHardwareCdromResp) GetValue() VcenterVmHardwareCdromInfo {
	if o == nil {
		var ret VcenterVmHardwareCdromInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromGetVmHardwareCdromResp) GetValueOk() (*VcenterVmHardwareCdromInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareCdromGetVmHardwareCdromResp) SetValue(v VcenterVmHardwareCdromInfo) {
	o.Value = v
}

func (o VcenterVmHardwareCdromGetVmHardwareCdromResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareCdromGetVmHardwareCdromResp struct {
	value *VcenterVmHardwareCdromGetVmHardwareCdromResp
	isSet bool
}

func (v NullableVcenterVmHardwareCdromGetVmHardwareCdromResp) Get() *VcenterVmHardwareCdromGetVmHardwareCdromResp {
	return v.value
}

func (v *NullableVcenterVmHardwareCdromGetVmHardwareCdromResp) Set(val *VcenterVmHardwareCdromGetVmHardwareCdromResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareCdromGetVmHardwareCdromResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareCdromGetVmHardwareCdromResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareCdromGetVmHardwareCdromResp(val *VcenterVmHardwareCdromGetVmHardwareCdromResp) *NullableVcenterVmHardwareCdromGetVmHardwareCdromResp {
	return &NullableVcenterVmHardwareCdromGetVmHardwareCdromResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareCdromGetVmHardwareCdromResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareCdromGetVmHardwareCdromResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

