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

// VcenterVMCreateVmResp struct for VcenterVMCreateVmResp
type VcenterVMCreateVmResp struct {
	Value string `json:"value"`
}

// NewVcenterVMCreateVmResp instantiates a new VcenterVMCreateVmResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMCreateVmResp(value string) *VcenterVMCreateVmResp {
	this := VcenterVMCreateVmResp{}
	this.Value = value
	return &this
}

// NewVcenterVMCreateVmRespWithDefaults instantiates a new VcenterVMCreateVmResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMCreateVmRespWithDefaults() *VcenterVMCreateVmResp {
	this := VcenterVMCreateVmResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVMCreateVmResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVMCreateVmResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVMCreateVmResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVMCreateVmResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMCreateVmResp struct {
	value *VcenterVMCreateVmResp
	isSet bool
}

func (v NullableVcenterVMCreateVmResp) Get() *VcenterVMCreateVmResp {
	return v.value
}

func (v *NullableVcenterVMCreateVmResp) Set(val *VcenterVMCreateVmResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMCreateVmResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMCreateVmResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMCreateVmResp(val *VcenterVMCreateVmResp) *NullableVcenterVMCreateVmResp {
	return &NullableVcenterVMCreateVmResp{value: val, isSet: true}
}

func (v NullableVcenterVMCreateVmResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMCreateVmResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


