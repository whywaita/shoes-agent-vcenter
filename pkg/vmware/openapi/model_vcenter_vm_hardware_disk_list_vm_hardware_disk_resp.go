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

// VcenterVmHardwareDiskListVmHardwareDiskResp struct for VcenterVmHardwareDiskListVmHardwareDiskResp
type VcenterVmHardwareDiskListVmHardwareDiskResp struct {
	Value []VcenterVmHardwareDiskSummary `json:"value"`
}

// NewVcenterVmHardwareDiskListVmHardwareDiskResp instantiates a new VcenterVmHardwareDiskListVmHardwareDiskResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareDiskListVmHardwareDiskResp(value []VcenterVmHardwareDiskSummary) *VcenterVmHardwareDiskListVmHardwareDiskResp {
	this := VcenterVmHardwareDiskListVmHardwareDiskResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareDiskListVmHardwareDiskRespWithDefaults instantiates a new VcenterVmHardwareDiskListVmHardwareDiskResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareDiskListVmHardwareDiskRespWithDefaults() *VcenterVmHardwareDiskListVmHardwareDiskResp {
	this := VcenterVmHardwareDiskListVmHardwareDiskResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareDiskListVmHardwareDiskResp) GetValue() []VcenterVmHardwareDiskSummary {
	if o == nil {
		var ret []VcenterVmHardwareDiskSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskListVmHardwareDiskResp) GetValueOk() (*[]VcenterVmHardwareDiskSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareDiskListVmHardwareDiskResp) SetValue(v []VcenterVmHardwareDiskSummary) {
	o.Value = v
}

func (o VcenterVmHardwareDiskListVmHardwareDiskResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareDiskListVmHardwareDiskResp struct {
	value *VcenterVmHardwareDiskListVmHardwareDiskResp
	isSet bool
}

func (v NullableVcenterVmHardwareDiskListVmHardwareDiskResp) Get() *VcenterVmHardwareDiskListVmHardwareDiskResp {
	return v.value
}

func (v *NullableVcenterVmHardwareDiskListVmHardwareDiskResp) Set(val *VcenterVmHardwareDiskListVmHardwareDiskResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareDiskListVmHardwareDiskResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareDiskListVmHardwareDiskResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareDiskListVmHardwareDiskResp(val *VcenterVmHardwareDiskListVmHardwareDiskResp) *NullableVcenterVmHardwareDiskListVmHardwareDiskResp {
	return &NullableVcenterVmHardwareDiskListVmHardwareDiskResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareDiskListVmHardwareDiskResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareDiskListVmHardwareDiskResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


