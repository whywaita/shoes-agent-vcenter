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

// VcenterTopologyReplicationStatusListTopologyReplicationStatusResp struct for VcenterTopologyReplicationStatusListTopologyReplicationStatusResp
type VcenterTopologyReplicationStatusListTopologyReplicationStatusResp struct {
	Value []VcenterTopologyReplicationStatusSummary `json:"value"`
}

// NewVcenterTopologyReplicationStatusListTopologyReplicationStatusResp instantiates a new VcenterTopologyReplicationStatusListTopologyReplicationStatusResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTopologyReplicationStatusListTopologyReplicationStatusResp(value []VcenterTopologyReplicationStatusSummary) *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp {
	this := VcenterTopologyReplicationStatusListTopologyReplicationStatusResp{}
	this.Value = value
	return &this
}

// NewVcenterTopologyReplicationStatusListTopologyReplicationStatusRespWithDefaults instantiates a new VcenterTopologyReplicationStatusListTopologyReplicationStatusResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTopologyReplicationStatusListTopologyReplicationStatusRespWithDefaults() *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp {
	this := VcenterTopologyReplicationStatusListTopologyReplicationStatusResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp) GetValue() []VcenterTopologyReplicationStatusSummary {
	if o == nil {
		var ret []VcenterTopologyReplicationStatusSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp) GetValueOk() (*[]VcenterTopologyReplicationStatusSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp) SetValue(v []VcenterTopologyReplicationStatusSummary) {
	o.Value = v
}

func (o VcenterTopologyReplicationStatusListTopologyReplicationStatusResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp struct {
	value *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp
	isSet bool
}

func (v NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp) Get() *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp {
	return v.value
}

func (v *NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp) Set(val *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp(val *VcenterTopologyReplicationStatusListTopologyReplicationStatusResp) *NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp {
	return &NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp{value: val, isSet: true}
}

func (v NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTopologyReplicationStatusListTopologyReplicationStatusResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

