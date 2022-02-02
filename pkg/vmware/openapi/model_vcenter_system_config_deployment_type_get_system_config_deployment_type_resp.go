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

// VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp struct for VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp
type VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp struct {
	Value VcenterSystemConfigDeploymentTypeInfo `json:"value"`
}

// NewVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp instantiates a new VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp(value VcenterSystemConfigDeploymentTypeInfo) *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp {
	this := VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp{}
	this.Value = value
	return &this
}

// NewVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeRespWithDefaults instantiates a new VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeRespWithDefaults() *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp {
	this := VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) GetValue() VcenterSystemConfigDeploymentTypeInfo {
	if o == nil {
		var ret VcenterSystemConfigDeploymentTypeInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) GetValueOk() (*VcenterSystemConfigDeploymentTypeInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) SetValue(v VcenterSystemConfigDeploymentTypeInfo) {
	o.Value = v
}

func (o VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp struct {
	value *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp
	isSet bool
}

func (v NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) Get() *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp {
	return v.value
}

func (v *NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) Set(val *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp(val *VcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) *NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp {
	return &NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp{value: val, isSet: true}
}

func (v NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterSystemConfigDeploymentTypeGetSystemConfigDeploymentTypeResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


