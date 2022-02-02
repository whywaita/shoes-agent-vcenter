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

// VcenterVmLibraryItemGetVmLibraryItemResp struct for VcenterVmLibraryItemGetVmLibraryItemResp
type VcenterVmLibraryItemGetVmLibraryItemResp struct {
	Value VcenterVmLibraryItemInfo `json:"value"`
}

// NewVcenterVmLibraryItemGetVmLibraryItemResp instantiates a new VcenterVmLibraryItemGetVmLibraryItemResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmLibraryItemGetVmLibraryItemResp(value VcenterVmLibraryItemInfo) *VcenterVmLibraryItemGetVmLibraryItemResp {
	this := VcenterVmLibraryItemGetVmLibraryItemResp{}
	this.Value = value
	return &this
}

// NewVcenterVmLibraryItemGetVmLibraryItemRespWithDefaults instantiates a new VcenterVmLibraryItemGetVmLibraryItemResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmLibraryItemGetVmLibraryItemRespWithDefaults() *VcenterVmLibraryItemGetVmLibraryItemResp {
	this := VcenterVmLibraryItemGetVmLibraryItemResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVmLibraryItemGetVmLibraryItemResp) GetValue() VcenterVmLibraryItemInfo {
	if o == nil {
		var ret VcenterVmLibraryItemInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVmLibraryItemGetVmLibraryItemResp) GetValueOk() (*VcenterVmLibraryItemInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVmLibraryItemGetVmLibraryItemResp) SetValue(v VcenterVmLibraryItemInfo) {
	o.Value = v
}

func (o VcenterVmLibraryItemGetVmLibraryItemResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmLibraryItemGetVmLibraryItemResp struct {
	value *VcenterVmLibraryItemGetVmLibraryItemResp
	isSet bool
}

func (v NullableVcenterVmLibraryItemGetVmLibraryItemResp) Get() *VcenterVmLibraryItemGetVmLibraryItemResp {
	return v.value
}

func (v *NullableVcenterVmLibraryItemGetVmLibraryItemResp) Set(val *VcenterVmLibraryItemGetVmLibraryItemResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmLibraryItemGetVmLibraryItemResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmLibraryItemGetVmLibraryItemResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmLibraryItemGetVmLibraryItemResp(val *VcenterVmLibraryItemGetVmLibraryItemResp) *NullableVcenterVmLibraryItemGetVmLibraryItemResp {
	return &NullableVcenterVmLibraryItemGetVmLibraryItemResp{value: val, isSet: true}
}

func (v NullableVcenterVmLibraryItemGetVmLibraryItemResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmLibraryItemGetVmLibraryItemResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

