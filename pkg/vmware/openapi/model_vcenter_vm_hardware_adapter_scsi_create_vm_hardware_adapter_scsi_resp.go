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

// VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp struct for VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp
type VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp struct {
	Value string `json:"value"`
}

// NewVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp instantiates a new VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp(value string) *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp {
	this := VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiRespWithDefaults instantiates a new VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiRespWithDefaults() *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp {
	this := VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp struct {
	value *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp
	isSet bool
}

func (v NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) Get() *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp {
	return v.value
}

func (v *NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) Set(val *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp(val *VcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) *NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp {
	return &NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareAdapterScsiCreateVmHardwareAdapterScsiResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

