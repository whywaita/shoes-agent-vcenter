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

// VcenterVMCloneVmResp struct for VcenterVMCloneVmResp
type VcenterVMCloneVmResp struct {
	Value string `json:"value"`
}

// NewVcenterVMCloneVmResp instantiates a new VcenterVMCloneVmResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMCloneVmResp(value string) *VcenterVMCloneVmResp {
	this := VcenterVMCloneVmResp{}
	this.Value = value
	return &this
}

// NewVcenterVMCloneVmRespWithDefaults instantiates a new VcenterVMCloneVmResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMCloneVmRespWithDefaults() *VcenterVMCloneVmResp {
	this := VcenterVMCloneVmResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVMCloneVmResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneVmResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVMCloneVmResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVMCloneVmResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMCloneVmResp struct {
	value *VcenterVMCloneVmResp
	isSet bool
}

func (v NullableVcenterVMCloneVmResp) Get() *VcenterVMCloneVmResp {
	return v.value
}

func (v *NullableVcenterVMCloneVmResp) Set(val *VcenterVMCloneVmResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMCloneVmResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMCloneVmResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMCloneVmResp(val *VcenterVMCloneVmResp) *NullableVcenterVMCloneVmResp {
	return &NullableVcenterVMCloneVmResp{value: val, isSet: true}
}

func (v NullableVcenterVMCloneVmResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMCloneVmResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

