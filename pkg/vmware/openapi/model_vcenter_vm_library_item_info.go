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

// VcenterVmLibraryItemInfo struct for VcenterVmLibraryItemInfo
type VcenterVmLibraryItemInfo struct {
	CheckOut *VcenterVmLibraryItemCheckOutInfo `json:"check_out,omitempty"`
}

// NewVcenterVmLibraryItemInfo instantiates a new VcenterVmLibraryItemInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmLibraryItemInfo() *VcenterVmLibraryItemInfo {
	this := VcenterVmLibraryItemInfo{}
	return &this
}

// NewVcenterVmLibraryItemInfoWithDefaults instantiates a new VcenterVmLibraryItemInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmLibraryItemInfoWithDefaults() *VcenterVmLibraryItemInfo {
	this := VcenterVmLibraryItemInfo{}
	return &this
}

// GetCheckOut returns the CheckOut field value if set, zero value otherwise.
func (o *VcenterVmLibraryItemInfo) GetCheckOut() VcenterVmLibraryItemCheckOutInfo {
	if o == nil || o.CheckOut == nil {
		var ret VcenterVmLibraryItemCheckOutInfo
		return ret
	}
	return *o.CheckOut
}

// GetCheckOutOk returns a tuple with the CheckOut field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmLibraryItemInfo) GetCheckOutOk() (*VcenterVmLibraryItemCheckOutInfo, bool) {
	if o == nil || o.CheckOut == nil {
		return nil, false
	}
	return o.CheckOut, true
}

// HasCheckOut returns a boolean if a field has been set.
func (o *VcenterVmLibraryItemInfo) HasCheckOut() bool {
	if o != nil && o.CheckOut != nil {
		return true
	}

	return false
}

// SetCheckOut gets a reference to the given VcenterVmLibraryItemCheckOutInfo and assigns it to the CheckOut field.
func (o *VcenterVmLibraryItemInfo) SetCheckOut(v VcenterVmLibraryItemCheckOutInfo) {
	o.CheckOut = &v
}

func (o VcenterVmLibraryItemInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CheckOut != nil {
		toSerialize["check_out"] = o.CheckOut
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmLibraryItemInfo struct {
	value *VcenterVmLibraryItemInfo
	isSet bool
}

func (v NullableVcenterVmLibraryItemInfo) Get() *VcenterVmLibraryItemInfo {
	return v.value
}

func (v *NullableVcenterVmLibraryItemInfo) Set(val *VcenterVmLibraryItemInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmLibraryItemInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmLibraryItemInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmLibraryItemInfo(val *VcenterVmLibraryItemInfo) *NullableVcenterVmLibraryItemInfo {
	return &NullableVcenterVmLibraryItemInfo{value: val, isSet: true}
}

func (v NullableVcenterVmLibraryItemInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmLibraryItemInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

