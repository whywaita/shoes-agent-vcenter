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

// VcenterNetworkListNetworkResp struct for VcenterNetworkListNetworkResp
type VcenterNetworkListNetworkResp struct {
	Value []VcenterNetworkSummary `json:"value"`
}

// NewVcenterNetworkListNetworkResp instantiates a new VcenterNetworkListNetworkResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNetworkListNetworkResp(value []VcenterNetworkSummary) *VcenterNetworkListNetworkResp {
	this := VcenterNetworkListNetworkResp{}
	this.Value = value
	return &this
}

// NewVcenterNetworkListNetworkRespWithDefaults instantiates a new VcenterNetworkListNetworkResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNetworkListNetworkRespWithDefaults() *VcenterNetworkListNetworkResp {
	this := VcenterNetworkListNetworkResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterNetworkListNetworkResp) GetValue() []VcenterNetworkSummary {
	if o == nil {
		var ret []VcenterNetworkSummary
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterNetworkListNetworkResp) GetValueOk() (*[]VcenterNetworkSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterNetworkListNetworkResp) SetValue(v []VcenterNetworkSummary) {
	o.Value = v
}

func (o VcenterNetworkListNetworkResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNetworkListNetworkResp struct {
	value *VcenterNetworkListNetworkResp
	isSet bool
}

func (v NullableVcenterNetworkListNetworkResp) Get() *VcenterNetworkListNetworkResp {
	return v.value
}

func (v *NullableVcenterNetworkListNetworkResp) Set(val *VcenterNetworkListNetworkResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNetworkListNetworkResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNetworkListNetworkResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNetworkListNetworkResp(val *VcenterNetworkListNetworkResp) *NullableVcenterNetworkListNetworkResp {
	return &NullableVcenterNetworkListNetworkResp{value: val, isSet: true}
}

func (v NullableVcenterNetworkListNetworkResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNetworkListNetworkResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


