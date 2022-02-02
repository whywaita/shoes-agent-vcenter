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

// VcenterDatacenterGetDatacenterResp struct for VcenterDatacenterGetDatacenterResp
type VcenterDatacenterGetDatacenterResp struct {
	Value VcenterDatacenterInfo `json:"value"`
}

// NewVcenterDatacenterGetDatacenterResp instantiates a new VcenterDatacenterGetDatacenterResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatacenterGetDatacenterResp(value VcenterDatacenterInfo) *VcenterDatacenterGetDatacenterResp {
	this := VcenterDatacenterGetDatacenterResp{}
	this.Value = value
	return &this
}

// NewVcenterDatacenterGetDatacenterRespWithDefaults instantiates a new VcenterDatacenterGetDatacenterResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatacenterGetDatacenterRespWithDefaults() *VcenterDatacenterGetDatacenterResp {
	this := VcenterDatacenterGetDatacenterResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDatacenterGetDatacenterResp) GetValue() VcenterDatacenterInfo {
	if o == nil {
		var ret VcenterDatacenterInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDatacenterGetDatacenterResp) GetValueOk() (*VcenterDatacenterInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDatacenterGetDatacenterResp) SetValue(v VcenterDatacenterInfo) {
	o.Value = v
}

func (o VcenterDatacenterGetDatacenterResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatacenterGetDatacenterResp struct {
	value *VcenterDatacenterGetDatacenterResp
	isSet bool
}

func (v NullableVcenterDatacenterGetDatacenterResp) Get() *VcenterDatacenterGetDatacenterResp {
	return v.value
}

func (v *NullableVcenterDatacenterGetDatacenterResp) Set(val *VcenterDatacenterGetDatacenterResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatacenterGetDatacenterResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatacenterGetDatacenterResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatacenterGetDatacenterResp(val *VcenterDatacenterGetDatacenterResp) *NullableVcenterDatacenterGetDatacenterResp {
	return &NullableVcenterDatacenterGetDatacenterResp{value: val, isSet: true}
}

func (v NullableVcenterDatacenterGetDatacenterResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatacenterGetDatacenterResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


