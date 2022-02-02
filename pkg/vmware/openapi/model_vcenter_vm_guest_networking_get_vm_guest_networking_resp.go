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

// VcenterVmGuestNetworkingGetVmGuestNetworkingResp struct for VcenterVmGuestNetworkingGetVmGuestNetworkingResp
type VcenterVmGuestNetworkingGetVmGuestNetworkingResp struct {
	Value VcenterVmGuestNetworkingInfo `json:"value"`
}

// NewVcenterVmGuestNetworkingGetVmGuestNetworkingResp instantiates a new VcenterVmGuestNetworkingGetVmGuestNetworkingResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestNetworkingGetVmGuestNetworkingResp(value VcenterVmGuestNetworkingInfo) *VcenterVmGuestNetworkingGetVmGuestNetworkingResp {
	this := VcenterVmGuestNetworkingGetVmGuestNetworkingResp{}
	this.Value = value
	return &this
}

// NewVcenterVmGuestNetworkingGetVmGuestNetworkingRespWithDefaults instantiates a new VcenterVmGuestNetworkingGetVmGuestNetworkingResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestNetworkingGetVmGuestNetworkingRespWithDefaults() *VcenterVmGuestNetworkingGetVmGuestNetworkingResp {
	this := VcenterVmGuestNetworkingGetVmGuestNetworkingResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmGuestNetworkingGetVmGuestNetworkingResp) GetValue() VcenterVmGuestNetworkingInfo {
	if o == nil {
		var ret VcenterVmGuestNetworkingInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestNetworkingGetVmGuestNetworkingResp) GetValueOk() (*VcenterVmGuestNetworkingInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmGuestNetworkingGetVmGuestNetworkingResp) SetValue(v VcenterVmGuestNetworkingInfo) {
	o.Value = v
}

func (o VcenterVmGuestNetworkingGetVmGuestNetworkingResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp struct {
	value *VcenterVmGuestNetworkingGetVmGuestNetworkingResp
	isSet bool
}

func (v NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp) Get() *VcenterVmGuestNetworkingGetVmGuestNetworkingResp {
	return v.value
}

func (v *NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp) Set(val *VcenterVmGuestNetworkingGetVmGuestNetworkingResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp(val *VcenterVmGuestNetworkingGetVmGuestNetworkingResp) *NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp {
	return &NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp{value: val, isSet: true}
}

func (v NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestNetworkingGetVmGuestNetworkingResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


