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

// VcenterHvcLinksListHvcLinksResp struct for VcenterHvcLinksListHvcLinksResp
type VcenterHvcLinksListHvcLinksResp struct {
	Value []VcenterHvcLinksSummary `json:"value"`
}

// NewVcenterHvcLinksListHvcLinksResp instantiates a new VcenterHvcLinksListHvcLinksResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterHvcLinksListHvcLinksResp(value []VcenterHvcLinksSummary) *VcenterHvcLinksListHvcLinksResp {
	this := VcenterHvcLinksListHvcLinksResp{}
	this.Value = value
	return &this
}

// NewVcenterHvcLinksListHvcLinksRespWithDefaults instantiates a new VcenterHvcLinksListHvcLinksResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterHvcLinksListHvcLinksRespWithDefaults() *VcenterHvcLinksListHvcLinksResp {
	this := VcenterHvcLinksListHvcLinksResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterHvcLinksListHvcLinksResp) GetValue() []VcenterHvcLinksSummary {
	if o == nil {
		var ret []VcenterHvcLinksSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksListHvcLinksResp) GetValueOk() (*[]VcenterHvcLinksSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterHvcLinksListHvcLinksResp) SetValue(v []VcenterHvcLinksSummary) {
	o.Value = v
}

func (o VcenterHvcLinksListHvcLinksResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterHvcLinksListHvcLinksResp struct {
	value *VcenterHvcLinksListHvcLinksResp
	isSet bool
}

func (v NullableVcenterHvcLinksListHvcLinksResp) Get() *VcenterHvcLinksListHvcLinksResp {
	return v.value
}

func (v *NullableVcenterHvcLinksListHvcLinksResp) Set(val *VcenterHvcLinksListHvcLinksResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterHvcLinksListHvcLinksResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterHvcLinksListHvcLinksResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterHvcLinksListHvcLinksResp(val *VcenterHvcLinksListHvcLinksResp) *NullableVcenterHvcLinksListHvcLinksResp {
	return &NullableVcenterHvcLinksListHvcLinksResp{value: val, isSet: true}
}

func (v NullableVcenterHvcLinksListHvcLinksResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterHvcLinksListHvcLinksResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


