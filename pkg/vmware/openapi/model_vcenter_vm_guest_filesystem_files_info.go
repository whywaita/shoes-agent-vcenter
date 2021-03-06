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

// VcenterVmGuestFilesystemFilesInfo struct for VcenterVmGuestFilesystemFilesInfo
type VcenterVmGuestFilesystemFilesInfo struct {
	Type VcenterVmGuestFilesystemFilesType `json:"type"`
	// The file size in bytes.
	Size int64 `json:"size"`
	Attributes VcenterVmGuestFilesystemFilesFileAttributesInfo `json:"attributes"`
}

// NewVcenterVmGuestFilesystemFilesInfo instantiates a new VcenterVmGuestFilesystemFilesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestFilesystemFilesInfo(type_ VcenterVmGuestFilesystemFilesType, size int64, attributes VcenterVmGuestFilesystemFilesFileAttributesInfo) *VcenterVmGuestFilesystemFilesInfo {
	this := VcenterVmGuestFilesystemFilesInfo{}
	this.Type = type_
	this.Size = size
	this.Attributes = attributes
	return &this
}

// NewVcenterVmGuestFilesystemFilesInfoWithDefaults instantiates a new VcenterVmGuestFilesystemFilesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestFilesystemFilesInfoWithDefaults() *VcenterVmGuestFilesystemFilesInfo {
	this := VcenterVmGuestFilesystemFilesInfo{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmGuestFilesystemFilesInfo) GetType() VcenterVmGuestFilesystemFilesType {
	if o == nil {
		var ret VcenterVmGuestFilesystemFilesType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesInfo) GetTypeOk() (*VcenterVmGuestFilesystemFilesType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmGuestFilesystemFilesInfo) SetType(v VcenterVmGuestFilesystemFilesType) {
	o.Type = v
}

// GetSize returns the Size field value
func (o *VcenterVmGuestFilesystemFilesInfo) GetSize() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesInfo) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *VcenterVmGuestFilesystemFilesInfo) SetSize(v int64) {
	o.Size = v
}

// GetAttributes returns the Attributes field value
func (o *VcenterVmGuestFilesystemFilesInfo) GetAttributes() VcenterVmGuestFilesystemFilesFileAttributesInfo {
	if o == nil {
		var ret VcenterVmGuestFilesystemFilesFileAttributesInfo
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesInfo) GetAttributesOk() (*VcenterVmGuestFilesystemFilesFileAttributesInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *VcenterVmGuestFilesystemFilesInfo) SetAttributes(v VcenterVmGuestFilesystemFilesFileAttributesInfo) {
	o.Attributes = v
}

func (o VcenterVmGuestFilesystemFilesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestFilesystemFilesInfo struct {
	value *VcenterVmGuestFilesystemFilesInfo
	isSet bool
}

func (v NullableVcenterVmGuestFilesystemFilesInfo) Get() *VcenterVmGuestFilesystemFilesInfo {
	return v.value
}

func (v *NullableVcenterVmGuestFilesystemFilesInfo) Set(val *VcenterVmGuestFilesystemFilesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestFilesystemFilesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestFilesystemFilesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestFilesystemFilesInfo(val *VcenterVmGuestFilesystemFilesInfo) *NullableVcenterVmGuestFilesystemFilesInfo {
	return &NullableVcenterVmGuestFilesystemFilesInfo{value: val, isSet: true}
}

func (v NullableVcenterVmGuestFilesystemFilesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestFilesystemFilesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


