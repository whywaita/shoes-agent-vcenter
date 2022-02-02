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

// VcenterDatastoreListDatastoreResp struct for VcenterDatastoreListDatastoreResp
type VcenterDatastoreListDatastoreResp struct {
	Value []VcenterDatastoreSummary `json:"value"`
}

// NewVcenterDatastoreListDatastoreResp instantiates a new VcenterDatastoreListDatastoreResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatastoreListDatastoreResp(value []VcenterDatastoreSummary) *VcenterDatastoreListDatastoreResp {
	this := VcenterDatastoreListDatastoreResp{}
	this.Value = value
	return &this
}

// NewVcenterDatastoreListDatastoreRespWithDefaults instantiates a new VcenterDatastoreListDatastoreResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatastoreListDatastoreRespWithDefaults() *VcenterDatastoreListDatastoreResp {
	this := VcenterDatastoreListDatastoreResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDatastoreListDatastoreResp) GetValue() []VcenterDatastoreSummary {
	if o == nil {
		var ret []VcenterDatastoreSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreListDatastoreResp) GetValueOk() (*[]VcenterDatastoreSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDatastoreListDatastoreResp) SetValue(v []VcenterDatastoreSummary) {
	o.Value = v
}

func (o VcenterDatastoreListDatastoreResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatastoreListDatastoreResp struct {
	value *VcenterDatastoreListDatastoreResp
	isSet bool
}

func (v NullableVcenterDatastoreListDatastoreResp) Get() *VcenterDatastoreListDatastoreResp {
	return v.value
}

func (v *NullableVcenterDatastoreListDatastoreResp) Set(val *VcenterDatastoreListDatastoreResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatastoreListDatastoreResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatastoreListDatastoreResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatastoreListDatastoreResp(val *VcenterDatastoreListDatastoreResp) *NullableVcenterDatastoreListDatastoreResp {
	return &NullableVcenterDatastoreListDatastoreResp{value: val, isSet: true}
}

func (v NullableVcenterDatastoreListDatastoreResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatastoreListDatastoreResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


