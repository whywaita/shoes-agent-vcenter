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

// VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp struct for VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp
type VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp struct {
	Value []VcenterStoragePoliciesComplianceSummary `json:"value"`
}

// NewVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp instantiates a new VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp(value []VcenterStoragePoliciesComplianceSummary) *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp {
	this := VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp{}
	this.Value = value
	return &this
}

// NewVcenterStoragePoliciesComplianceListStoragePoliciesComplianceRespWithDefaults instantiates a new VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterStoragePoliciesComplianceListStoragePoliciesComplianceRespWithDefaults() *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp {
	this := VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) GetValue() []VcenterStoragePoliciesComplianceSummary {
	if o == nil {
		var ret []VcenterStoragePoliciesComplianceSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) GetValueOk() (*[]VcenterStoragePoliciesComplianceSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) SetValue(v []VcenterStoragePoliciesComplianceSummary) {
	o.Value = v
}

func (o VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp struct {
	value *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp
	isSet bool
}

func (v NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) Get() *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp {
	return v.value
}

func (v *NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) Set(val *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp(val *VcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) *NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp {
	return &NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp{value: val, isSet: true}
}

func (v NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterStoragePoliciesComplianceListStoragePoliciesComplianceResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


