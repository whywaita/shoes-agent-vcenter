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

// VcenterDeploymentGetDeploymentResp struct for VcenterDeploymentGetDeploymentResp
type VcenterDeploymentGetDeploymentResp struct {
	Value VcenterDeploymentInfo `json:"value"`
}

// NewVcenterDeploymentGetDeploymentResp instantiates a new VcenterDeploymentGetDeploymentResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentGetDeploymentResp(value VcenterDeploymentInfo) *VcenterDeploymentGetDeploymentResp {
	this := VcenterDeploymentGetDeploymentResp{}
	this.Value = value
	return &this
}

// NewVcenterDeploymentGetDeploymentRespWithDefaults instantiates a new VcenterDeploymentGetDeploymentResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentGetDeploymentRespWithDefaults() *VcenterDeploymentGetDeploymentResp {
	this := VcenterDeploymentGetDeploymentResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDeploymentGetDeploymentResp) GetValue() VcenterDeploymentInfo {
	if o == nil {
		var ret VcenterDeploymentInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentGetDeploymentResp) GetValueOk() (*VcenterDeploymentInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDeploymentGetDeploymentResp) SetValue(v VcenterDeploymentInfo) {
	o.Value = v
}

func (o VcenterDeploymentGetDeploymentResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentGetDeploymentResp struct {
	value *VcenterDeploymentGetDeploymentResp
	isSet bool
}

func (v NullableVcenterDeploymentGetDeploymentResp) Get() *VcenterDeploymentGetDeploymentResp {
	return v.value
}

func (v *NullableVcenterDeploymentGetDeploymentResp) Set(val *VcenterDeploymentGetDeploymentResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentGetDeploymentResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentGetDeploymentResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentGetDeploymentResp(val *VcenterDeploymentGetDeploymentResp) *NullableVcenterDeploymentGetDeploymentResp {
	return &NullableVcenterDeploymentGetDeploymentResp{value: val, isSet: true}
}

func (v NullableVcenterDeploymentGetDeploymentResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentGetDeploymentResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

