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

// VcenterVmHardwareFloppyCreateVmHardwareFloppyResp struct for VcenterVmHardwareFloppyCreateVmHardwareFloppyResp
type VcenterVmHardwareFloppyCreateVmHardwareFloppyResp struct {
	Value string `json:"value"`
}

// NewVcenterVmHardwareFloppyCreateVmHardwareFloppyResp instantiates a new VcenterVmHardwareFloppyCreateVmHardwareFloppyResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareFloppyCreateVmHardwareFloppyResp(value string) *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp {
	this := VcenterVmHardwareFloppyCreateVmHardwareFloppyResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareFloppyCreateVmHardwareFloppyRespWithDefaults instantiates a new VcenterVmHardwareFloppyCreateVmHardwareFloppyResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareFloppyCreateVmHardwareFloppyRespWithDefaults() *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp {
	this := VcenterVmHardwareFloppyCreateVmHardwareFloppyResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVmHardwareFloppyCreateVmHardwareFloppyResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp struct {
	value *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp
	isSet bool
}

func (v NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp) Get() *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp {
	return v.value
}

func (v *NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp) Set(val *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp(val *VcenterVmHardwareFloppyCreateVmHardwareFloppyResp) *NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp {
	return &NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareFloppyCreateVmHardwareFloppyResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

