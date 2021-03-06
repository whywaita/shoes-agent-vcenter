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

// VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp struct for VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp
type VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp struct {
	Value []VcenterVmHardwareAdapterNvmeSummary `json:"value"`
}

// NewVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp instantiates a new VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp(value []VcenterVmHardwareAdapterNvmeSummary) *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp {
	this := VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeRespWithDefaults instantiates a new VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeRespWithDefaults() *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp {
	this := VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) GetValue() []VcenterVmHardwareAdapterNvmeSummary {
	if o == nil {
		var ret []VcenterVmHardwareAdapterNvmeSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) GetValueOk() (*[]VcenterVmHardwareAdapterNvmeSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) SetValue(v []VcenterVmHardwareAdapterNvmeSummary) {
	o.Value = v
}

func (o VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp struct {
	value *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp
	isSet bool
}

func (v NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) Get() *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp {
	return v.value
}

func (v *NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) Set(val *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp(val *VcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) *NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp {
	return &NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareAdapterNvmeListVmHardwareAdapterNvmeResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


