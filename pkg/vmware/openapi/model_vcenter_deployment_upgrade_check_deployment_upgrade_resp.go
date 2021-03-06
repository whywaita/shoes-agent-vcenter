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

// VcenterDeploymentUpgradeCheckDeploymentUpgradeResp struct for VcenterDeploymentUpgradeCheckDeploymentUpgradeResp
type VcenterDeploymentUpgradeCheckDeploymentUpgradeResp struct {
	Value VcenterDeploymentCheckInfo `json:"value"`
}

// NewVcenterDeploymentUpgradeCheckDeploymentUpgradeResp instantiates a new VcenterDeploymentUpgradeCheckDeploymentUpgradeResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentUpgradeCheckDeploymentUpgradeResp(value VcenterDeploymentCheckInfo) *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp {
	this := VcenterDeploymentUpgradeCheckDeploymentUpgradeResp{}
	this.Value = value
	return &this
}

// NewVcenterDeploymentUpgradeCheckDeploymentUpgradeRespWithDefaults instantiates a new VcenterDeploymentUpgradeCheckDeploymentUpgradeResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentUpgradeCheckDeploymentUpgradeRespWithDefaults() *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp {
	this := VcenterDeploymentUpgradeCheckDeploymentUpgradeResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp) GetValue() VcenterDeploymentCheckInfo {
	if o == nil {
		var ret VcenterDeploymentCheckInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp) GetValueOk() (*VcenterDeploymentCheckInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp) SetValue(v VcenterDeploymentCheckInfo) {
	o.Value = v
}

func (o VcenterDeploymentUpgradeCheckDeploymentUpgradeResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp struct {
	value *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp
	isSet bool
}

func (v NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp) Get() *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp {
	return v.value
}

func (v *NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp) Set(val *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp(val *VcenterDeploymentUpgradeCheckDeploymentUpgradeResp) *NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp {
	return &NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp{value: val, isSet: true}
}

func (v NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentUpgradeCheckDeploymentUpgradeResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


