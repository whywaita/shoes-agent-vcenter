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

// VcenterVmGuestIdentityGetVmGuestIdentityResp struct for VcenterVmGuestIdentityGetVmGuestIdentityResp
type VcenterVmGuestIdentityGetVmGuestIdentityResp struct {
	Value VcenterVmGuestIdentityInfo `json:"value"`
}

// NewVcenterVmGuestIdentityGetVmGuestIdentityResp instantiates a new VcenterVmGuestIdentityGetVmGuestIdentityResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestIdentityGetVmGuestIdentityResp(value VcenterVmGuestIdentityInfo) *VcenterVmGuestIdentityGetVmGuestIdentityResp {
	this := VcenterVmGuestIdentityGetVmGuestIdentityResp{}
	this.Value = value
	return &this
}

// NewVcenterVmGuestIdentityGetVmGuestIdentityRespWithDefaults instantiates a new VcenterVmGuestIdentityGetVmGuestIdentityResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestIdentityGetVmGuestIdentityRespWithDefaults() *VcenterVmGuestIdentityGetVmGuestIdentityResp {
	this := VcenterVmGuestIdentityGetVmGuestIdentityResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmGuestIdentityGetVmGuestIdentityResp) GetValue() VcenterVmGuestIdentityInfo {
	if o == nil {
		var ret VcenterVmGuestIdentityInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestIdentityGetVmGuestIdentityResp) GetValueOk() (*VcenterVmGuestIdentityInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmGuestIdentityGetVmGuestIdentityResp) SetValue(v VcenterVmGuestIdentityInfo) {
	o.Value = v
}

func (o VcenterVmGuestIdentityGetVmGuestIdentityResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestIdentityGetVmGuestIdentityResp struct {
	value *VcenterVmGuestIdentityGetVmGuestIdentityResp
	isSet bool
}

func (v NullableVcenterVmGuestIdentityGetVmGuestIdentityResp) Get() *VcenterVmGuestIdentityGetVmGuestIdentityResp {
	return v.value
}

func (v *NullableVcenterVmGuestIdentityGetVmGuestIdentityResp) Set(val *VcenterVmGuestIdentityGetVmGuestIdentityResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestIdentityGetVmGuestIdentityResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestIdentityGetVmGuestIdentityResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestIdentityGetVmGuestIdentityResp(val *VcenterVmGuestIdentityGetVmGuestIdentityResp) *NullableVcenterVmGuestIdentityGetVmGuestIdentityResp {
	return &NullableVcenterVmGuestIdentityGetVmGuestIdentityResp{value: val, isSet: true}
}

func (v NullableVcenterVmGuestIdentityGetVmGuestIdentityResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestIdentityGetVmGuestIdentityResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


