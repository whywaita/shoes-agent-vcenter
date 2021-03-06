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

// VcenterStoragePoliciesVMListStoragePoliciesVmResp struct for VcenterStoragePoliciesVMListStoragePoliciesVmResp
type VcenterStoragePoliciesVMListStoragePoliciesVmResp struct {
	Value []VcenterStoragePoliciesVMListStoragePoliciesVmRespValue `json:"value"`
}

// NewVcenterStoragePoliciesVMListStoragePoliciesVmResp instantiates a new VcenterStoragePoliciesVMListStoragePoliciesVmResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterStoragePoliciesVMListStoragePoliciesVmResp(value []VcenterStoragePoliciesVMListStoragePoliciesVmRespValue) *VcenterStoragePoliciesVMListStoragePoliciesVmResp {
	this := VcenterStoragePoliciesVMListStoragePoliciesVmResp{}
	this.Value = value
	return &this
}

// NewVcenterStoragePoliciesVMListStoragePoliciesVmRespWithDefaults instantiates a new VcenterStoragePoliciesVMListStoragePoliciesVmResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterStoragePoliciesVMListStoragePoliciesVmRespWithDefaults() *VcenterStoragePoliciesVMListStoragePoliciesVmResp {
	this := VcenterStoragePoliciesVMListStoragePoliciesVmResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterStoragePoliciesVMListStoragePoliciesVmResp) GetValue() []VcenterStoragePoliciesVMListStoragePoliciesVmRespValue {
	if o == nil {
		var ret []VcenterStoragePoliciesVMListStoragePoliciesVmRespValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterStoragePoliciesVMListStoragePoliciesVmResp) GetValueOk() (*[]VcenterStoragePoliciesVMListStoragePoliciesVmRespValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterStoragePoliciesVMListStoragePoliciesVmResp) SetValue(v []VcenterStoragePoliciesVMListStoragePoliciesVmRespValue) {
	o.Value = v
}

func (o VcenterStoragePoliciesVMListStoragePoliciesVmResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp struct {
	value *VcenterStoragePoliciesVMListStoragePoliciesVmResp
	isSet bool
}

func (v NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp) Get() *VcenterStoragePoliciesVMListStoragePoliciesVmResp {
	return v.value
}

func (v *NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp) Set(val *VcenterStoragePoliciesVMListStoragePoliciesVmResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterStoragePoliciesVMListStoragePoliciesVmResp(val *VcenterStoragePoliciesVMListStoragePoliciesVmResp) *NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp {
	return &NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp{value: val, isSet: true}
}

func (v NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterStoragePoliciesVMListStoragePoliciesVmResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


