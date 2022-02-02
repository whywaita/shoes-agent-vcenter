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

// VcenterServicesServiceListDetailsServicesServiceResp struct for VcenterServicesServiceListDetailsServicesServiceResp
type VcenterServicesServiceListDetailsServicesServiceResp struct {
	Value []VcenterServicesServiceListDetailsServicesServiceRespValue `json:"value"`
}

// NewVcenterServicesServiceListDetailsServicesServiceResp instantiates a new VcenterServicesServiceListDetailsServicesServiceResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterServicesServiceListDetailsServicesServiceResp(value []VcenterServicesServiceListDetailsServicesServiceRespValue) *VcenterServicesServiceListDetailsServicesServiceResp {
	this := VcenterServicesServiceListDetailsServicesServiceResp{}
	this.Value = value
	return &this
}

// NewVcenterServicesServiceListDetailsServicesServiceRespWithDefaults instantiates a new VcenterServicesServiceListDetailsServicesServiceResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterServicesServiceListDetailsServicesServiceRespWithDefaults() *VcenterServicesServiceListDetailsServicesServiceResp {
	this := VcenterServicesServiceListDetailsServicesServiceResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterServicesServiceListDetailsServicesServiceResp) GetValue() []VcenterServicesServiceListDetailsServicesServiceRespValue {
	if o == nil {
		var ret []VcenterServicesServiceListDetailsServicesServiceRespValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceListDetailsServicesServiceResp) GetValueOk() (*[]VcenterServicesServiceListDetailsServicesServiceRespValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterServicesServiceListDetailsServicesServiceResp) SetValue(v []VcenterServicesServiceListDetailsServicesServiceRespValue) {
	o.Value = v
}

func (o VcenterServicesServiceListDetailsServicesServiceResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterServicesServiceListDetailsServicesServiceResp struct {
	value *VcenterServicesServiceListDetailsServicesServiceResp
	isSet bool
}

func (v NullableVcenterServicesServiceListDetailsServicesServiceResp) Get() *VcenterServicesServiceListDetailsServicesServiceResp {
	return v.value
}

func (v *NullableVcenterServicesServiceListDetailsServicesServiceResp) Set(val *VcenterServicesServiceListDetailsServicesServiceResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterServicesServiceListDetailsServicesServiceResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterServicesServiceListDetailsServicesServiceResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterServicesServiceListDetailsServicesServiceResp(val *VcenterServicesServiceListDetailsServicesServiceResp) *NullableVcenterServicesServiceListDetailsServicesServiceResp {
	return &NullableVcenterServicesServiceListDetailsServicesServiceResp{value: val, isSet: true}
}

func (v NullableVcenterServicesServiceListDetailsServicesServiceResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterServicesServiceListDetailsServicesServiceResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


