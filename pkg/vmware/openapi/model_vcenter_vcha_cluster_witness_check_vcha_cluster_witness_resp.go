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

// VcenterVchaClusterWitnessCheckVchaClusterWitnessResp struct for VcenterVchaClusterWitnessCheckVchaClusterWitnessResp
type VcenterVchaClusterWitnessCheckVchaClusterWitnessResp struct {
	Value VcenterVchaClusterWitnessCheckResult `json:"value"`
}

// NewVcenterVchaClusterWitnessCheckVchaClusterWitnessResp instantiates a new VcenterVchaClusterWitnessCheckVchaClusterWitnessResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterWitnessCheckVchaClusterWitnessResp(value VcenterVchaClusterWitnessCheckResult) *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp {
	this := VcenterVchaClusterWitnessCheckVchaClusterWitnessResp{}
	this.Value = value
	return &this
}

// NewVcenterVchaClusterWitnessCheckVchaClusterWitnessRespWithDefaults instantiates a new VcenterVchaClusterWitnessCheckVchaClusterWitnessResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterWitnessCheckVchaClusterWitnessRespWithDefaults() *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp {
	this := VcenterVchaClusterWitnessCheckVchaClusterWitnessResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp) GetValue() VcenterVchaClusterWitnessCheckResult {
	if o == nil {
		var ret VcenterVchaClusterWitnessCheckResult
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp) GetValueOk() (*VcenterVchaClusterWitnessCheckResult, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp) SetValue(v VcenterVchaClusterWitnessCheckResult) {
	o.Value = v
}

func (o VcenterVchaClusterWitnessCheckVchaClusterWitnessResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp struct {
	value *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp
	isSet bool
}

func (v NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp) Get() *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp {
	return v.value
}

func (v *NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp) Set(val *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp(val *VcenterVchaClusterWitnessCheckVchaClusterWitnessResp) *NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp {
	return &NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterWitnessCheckVchaClusterWitnessResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


