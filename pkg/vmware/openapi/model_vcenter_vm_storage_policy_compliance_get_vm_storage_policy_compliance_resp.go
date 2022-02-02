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

// VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp struct for VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp
type VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp struct {
	Value VcenterVmStoragePolicyComplianceInfo `json:"value"`
}

// NewVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp instantiates a new VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp(value VcenterVmStoragePolicyComplianceInfo) *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp {
	this := VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp{}
	this.Value = value
	return &this
}

// NewVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceRespWithDefaults instantiates a new VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceRespWithDefaults() *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp {
	this := VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) GetValue() VcenterVmStoragePolicyComplianceInfo {
	if o == nil {
		var ret VcenterVmStoragePolicyComplianceInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) GetValueOk() (*VcenterVmStoragePolicyComplianceInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) SetValue(v VcenterVmStoragePolicyComplianceInfo) {
	o.Value = v
}

func (o VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp struct {
	value *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp
	isSet bool
}

func (v NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) Get() *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp {
	return v.value
}

func (v *NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) Set(val *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp(val *VcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) *NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp {
	return &NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp{value: val, isSet: true}
}

func (v NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmStoragePolicyComplianceGetVmStoragePolicyComplianceResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

