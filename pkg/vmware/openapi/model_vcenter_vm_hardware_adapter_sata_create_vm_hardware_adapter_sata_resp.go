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

// VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp struct for VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp
type VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp struct {
	Value string `json:"value"`
}

// NewVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp instantiates a new VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp(value string) *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp {
	this := VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataRespWithDefaults instantiates a new VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataRespWithDefaults() *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp {
	this := VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp struct {
	value *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp
	isSet bool
}

func (v NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) Get() *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp {
	return v.value
}

func (v *NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) Set(val *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp(val *VcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) *NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp {
	return &NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareAdapterSataCreateVmHardwareAdapterSataResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

