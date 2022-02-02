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

// VcenterVmHardwareDiskBackingSpec struct for VcenterVmHardwareDiskBackingSpec
type VcenterVmHardwareDiskBackingSpec struct {
	Type VcenterVmHardwareDiskBackingType `json:"type"`
	// Path of the VMDK file backing the virtual disk. This field is optional and it is only relevant when the value of Disk.BackingSpec.type is VMDK_FILE.
	VmdkFile *string `json:"vmdk_file,omitempty"`
}

// NewVcenterVmHardwareDiskBackingSpec instantiates a new VcenterVmHardwareDiskBackingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareDiskBackingSpec(type_ VcenterVmHardwareDiskBackingType) *VcenterVmHardwareDiskBackingSpec {
	this := VcenterVmHardwareDiskBackingSpec{}
	this.Type = type_
	return &this
}

// NewVcenterVmHardwareDiskBackingSpecWithDefaults instantiates a new VcenterVmHardwareDiskBackingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareDiskBackingSpecWithDefaults() *VcenterVmHardwareDiskBackingSpec {
	this := VcenterVmHardwareDiskBackingSpec{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmHardwareDiskBackingSpec) GetType() VcenterVmHardwareDiskBackingType {
	if o == nil {
		var ret VcenterVmHardwareDiskBackingType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskBackingSpec) GetTypeOk() (*VcenterVmHardwareDiskBackingType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareDiskBackingSpec) SetType(v VcenterVmHardwareDiskBackingType) {
	o.Type = v
}

// GetVmdkFile returns the VmdkFile field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskBackingSpec) GetVmdkFile() string {
	if o == nil || o.VmdkFile == nil {
		var ret string
		return ret
	}
	return *o.VmdkFile
}

// GetVmdkFileOk returns a tuple with the VmdkFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskBackingSpec) GetVmdkFileOk() (*string, bool) {
	if o == nil || o.VmdkFile == nil {
		return nil, false
	}
	return o.VmdkFile, true
}

// HasVmdkFile returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskBackingSpec) HasVmdkFile() bool {
	if o != nil && o.VmdkFile != nil {
		return true
	}

	return false
}

// SetVmdkFile gets a reference to the given string and assigns it to the VmdkFile field.
func (o *VcenterVmHardwareDiskBackingSpec) SetVmdkFile(v string) {
	o.VmdkFile = &v
}

func (o VcenterVmHardwareDiskBackingSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.VmdkFile != nil {
		toSerialize["vmdk_file"] = o.VmdkFile
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareDiskBackingSpec struct {
	value *VcenterVmHardwareDiskBackingSpec
	isSet bool
}

func (v NullableVcenterVmHardwareDiskBackingSpec) Get() *VcenterVmHardwareDiskBackingSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareDiskBackingSpec) Set(val *VcenterVmHardwareDiskBackingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareDiskBackingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareDiskBackingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareDiskBackingSpec(val *VcenterVmHardwareDiskBackingSpec) *NullableVcenterVmHardwareDiskBackingSpec {
	return &NullableVcenterVmHardwareDiskBackingSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareDiskBackingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareDiskBackingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

