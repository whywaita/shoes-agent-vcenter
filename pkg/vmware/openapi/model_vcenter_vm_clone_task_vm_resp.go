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

// VcenterVMCloneTaskVmResp struct for VcenterVMCloneTaskVmResp
type VcenterVMCloneTaskVmResp struct {
	Value string `json:"value"`
}

// NewVcenterVMCloneTaskVmResp instantiates a new VcenterVMCloneTaskVmResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMCloneTaskVmResp(value string) *VcenterVMCloneTaskVmResp {
	this := VcenterVMCloneTaskVmResp{}
	this.Value = value
	return &this
}

// NewVcenterVMCloneTaskVmRespWithDefaults instantiates a new VcenterVMCloneTaskVmResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMCloneTaskVmRespWithDefaults() *VcenterVMCloneTaskVmResp {
	this := VcenterVMCloneTaskVmResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVMCloneTaskVmResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneTaskVmResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVMCloneTaskVmResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVMCloneTaskVmResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMCloneTaskVmResp struct {
	value *VcenterVMCloneTaskVmResp
	isSet bool
}

func (v NullableVcenterVMCloneTaskVmResp) Get() *VcenterVMCloneTaskVmResp {
	return v.value
}

func (v *NullableVcenterVMCloneTaskVmResp) Set(val *VcenterVMCloneTaskVmResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMCloneTaskVmResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMCloneTaskVmResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMCloneTaskVmResp(val *VcenterVMCloneTaskVmResp) *NullableVcenterVMCloneTaskVmResp {
	return &NullableVcenterVMCloneTaskVmResp{value: val, isSet: true}
}

func (v NullableVcenterVMCloneTaskVmResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMCloneTaskVmResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


