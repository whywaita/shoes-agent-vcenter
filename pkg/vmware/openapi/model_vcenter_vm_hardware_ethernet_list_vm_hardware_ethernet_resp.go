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

// VcenterVmHardwareEthernetListVmHardwareEthernetResp struct for VcenterVmHardwareEthernetListVmHardwareEthernetResp
type VcenterVmHardwareEthernetListVmHardwareEthernetResp struct {
	Value []VcenterVmHardwareEthernetSummary `json:"value"`
}

// NewVcenterVmHardwareEthernetListVmHardwareEthernetResp instantiates a new VcenterVmHardwareEthernetListVmHardwareEthernetResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareEthernetListVmHardwareEthernetResp(value []VcenterVmHardwareEthernetSummary) *VcenterVmHardwareEthernetListVmHardwareEthernetResp {
	this := VcenterVmHardwareEthernetListVmHardwareEthernetResp{}
	this.Value = value
	return &this
}

// NewVcenterVmHardwareEthernetListVmHardwareEthernetRespWithDefaults instantiates a new VcenterVmHardwareEthernetListVmHardwareEthernetResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareEthernetListVmHardwareEthernetRespWithDefaults() *VcenterVmHardwareEthernetListVmHardwareEthernetResp {
	this := VcenterVmHardwareEthernetListVmHardwareEthernetResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmHardwareEthernetListVmHardwareEthernetResp) GetValue() []VcenterVmHardwareEthernetSummary {
	if o == nil {
		var ret []VcenterVmHardwareEthernetSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetListVmHardwareEthernetResp) GetValueOk() (*[]VcenterVmHardwareEthernetSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmHardwareEthernetListVmHardwareEthernetResp) SetValue(v []VcenterVmHardwareEthernetSummary) {
	o.Value = v
}

func (o VcenterVmHardwareEthernetListVmHardwareEthernetResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp struct {
	value *VcenterVmHardwareEthernetListVmHardwareEthernetResp
	isSet bool
}

func (v NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp) Get() *VcenterVmHardwareEthernetListVmHardwareEthernetResp {
	return v.value
}

func (v *NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp) Set(val *VcenterVmHardwareEthernetListVmHardwareEthernetResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareEthernetListVmHardwareEthernetResp(val *VcenterVmHardwareEthernetListVmHardwareEthernetResp) *NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp {
	return &NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareEthernetListVmHardwareEthernetResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


