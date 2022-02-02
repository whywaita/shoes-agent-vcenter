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

// VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp struct for VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp
type VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp struct {
	Value VcenterTokenserviceTokenExchangeInfo `json:"value"`
}

// NewVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp instantiates a new VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp(value VcenterTokenserviceTokenExchangeInfo) *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp {
	this := VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp{}
	this.Value = value
	return &this
}

// NewVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeRespWithDefaults instantiates a new VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeRespWithDefaults() *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp {
	this := VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) GetValue() VcenterTokenserviceTokenExchangeInfo {
	if o == nil {
		var ret VcenterTokenserviceTokenExchangeInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) GetValueOk() (*VcenterTokenserviceTokenExchangeInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) SetValue(v VcenterTokenserviceTokenExchangeInfo) {
	o.Value = v
}

func (o VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp struct {
	value *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp
	isSet bool
}

func (v NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) Get() *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp {
	return v.value
}

func (v *NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) Set(val *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp(val *VcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) *NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp {
	return &NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp{value: val, isSet: true}
}

func (v NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTokenserviceTokenExchangeExchangeTokenserviceTokenExchangeResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


