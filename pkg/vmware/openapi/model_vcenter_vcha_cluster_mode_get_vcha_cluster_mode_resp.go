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

// VcenterVchaClusterModeGetVchaClusterModeResp struct for VcenterVchaClusterModeGetVchaClusterModeResp
type VcenterVchaClusterModeGetVchaClusterModeResp struct {
	Value VcenterVchaClusterModeInfo `json:"value"`
}

// NewVcenterVchaClusterModeGetVchaClusterModeResp instantiates a new VcenterVchaClusterModeGetVchaClusterModeResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterModeGetVchaClusterModeResp(value VcenterVchaClusterModeInfo) *VcenterVchaClusterModeGetVchaClusterModeResp {
	this := VcenterVchaClusterModeGetVchaClusterModeResp{}
	this.Value = value
	return &this
}

// NewVcenterVchaClusterModeGetVchaClusterModeRespWithDefaults instantiates a new VcenterVchaClusterModeGetVchaClusterModeResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterModeGetVchaClusterModeRespWithDefaults() *VcenterVchaClusterModeGetVchaClusterModeResp {
	this := VcenterVchaClusterModeGetVchaClusterModeResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVchaClusterModeGetVchaClusterModeResp) GetValue() VcenterVchaClusterModeInfo {
	if o == nil {
		var ret VcenterVchaClusterModeInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterModeGetVchaClusterModeResp) GetValueOk() (*VcenterVchaClusterModeInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVchaClusterModeGetVchaClusterModeResp) SetValue(v VcenterVchaClusterModeInfo) {
	o.Value = v
}

func (o VcenterVchaClusterModeGetVchaClusterModeResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterModeGetVchaClusterModeResp struct {
	value *VcenterVchaClusterModeGetVchaClusterModeResp
	isSet bool
}

func (v NullableVcenterVchaClusterModeGetVchaClusterModeResp) Get() *VcenterVchaClusterModeGetVchaClusterModeResp {
	return v.value
}

func (v *NullableVcenterVchaClusterModeGetVchaClusterModeResp) Set(val *VcenterVchaClusterModeGetVchaClusterModeResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterModeGetVchaClusterModeResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterModeGetVchaClusterModeResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterModeGetVchaClusterModeResp(val *VcenterVchaClusterModeGetVchaClusterModeResp) *NullableVcenterVchaClusterModeGetVchaClusterModeResp {
	return &NullableVcenterVchaClusterModeGetVchaClusterModeResp{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterModeGetVchaClusterModeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterModeGetVchaClusterModeResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


