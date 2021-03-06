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

// VcenterVMInstantCloneVmResp struct for VcenterVMInstantCloneVmResp
type VcenterVMInstantCloneVmResp struct {
	Value string `json:"value"`
}

// NewVcenterVMInstantCloneVmResp instantiates a new VcenterVMInstantCloneVmResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMInstantCloneVmResp(value string) *VcenterVMInstantCloneVmResp {
	this := VcenterVMInstantCloneVmResp{}
	this.Value = value
	return &this
}

// NewVcenterVMInstantCloneVmRespWithDefaults instantiates a new VcenterVMInstantCloneVmResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMInstantCloneVmRespWithDefaults() *VcenterVMInstantCloneVmResp {
	this := VcenterVMInstantCloneVmResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVMInstantCloneVmResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVMInstantCloneVmResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVMInstantCloneVmResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVMInstantCloneVmResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMInstantCloneVmResp struct {
	value *VcenterVMInstantCloneVmResp
	isSet bool
}

func (v NullableVcenterVMInstantCloneVmResp) Get() *VcenterVMInstantCloneVmResp {
	return v.value
}

func (v *NullableVcenterVMInstantCloneVmResp) Set(val *VcenterVMInstantCloneVmResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMInstantCloneVmResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMInstantCloneVmResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMInstantCloneVmResp(val *VcenterVMInstantCloneVmResp) *NullableVcenterVMInstantCloneVmResp {
	return &NullableVcenterVMInstantCloneVmResp{value: val, isSet: true}
}

func (v NullableVcenterVMInstantCloneVmResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMInstantCloneVmResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


