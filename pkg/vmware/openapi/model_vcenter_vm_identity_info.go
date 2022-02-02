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

// VcenterVmIdentityInfo struct for VcenterVmIdentityInfo
type VcenterVmIdentityInfo struct {
	// Virtual machine name.
	Name string `json:"name"`
	// 128-bit SMBIOS UUID of a virtual machine represented as a hexadecimal string in \"12345678-abcd-1234-cdef-123456789abc\" format.
	BiosUuid string `json:"bios_uuid"`
	// VirtualCenter-specific 128-bit UUID of a virtual machine, represented as a hexademical string. This identifier is used by VirtualCenter to uniquely identify all virtual machine instances, including those that may share the same SMBIOS UUID.
	InstanceUuid string `json:"instance_uuid"`
}

// NewVcenterVmIdentityInfo instantiates a new VcenterVmIdentityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmIdentityInfo(name string, biosUuid string, instanceUuid string) *VcenterVmIdentityInfo {
	this := VcenterVmIdentityInfo{}
	this.Name = name
	this.BiosUuid = biosUuid
	this.InstanceUuid = instanceUuid
	return &this
}

// NewVcenterVmIdentityInfoWithDefaults instantiates a new VcenterVmIdentityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmIdentityInfoWithDefaults() *VcenterVmIdentityInfo {
	this := VcenterVmIdentityInfo{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterVmIdentityInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterVmIdentityInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterVmIdentityInfo) SetName(v string) {
	o.Name = v
}

// GetBiosUuid returns the BiosUuid field value
func (o *VcenterVmIdentityInfo) GetBiosUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BiosUuid
}

// GetBiosUuidOk returns a tuple with the BiosUuid field value
// and a boolean to check if the value has been set.
func (o *VcenterVmIdentityInfo) GetBiosUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BiosUuid, true
}

// SetBiosUuid sets field value
func (o *VcenterVmIdentityInfo) SetBiosUuid(v string) {
	o.BiosUuid = v
}

// GetInstanceUuid returns the InstanceUuid field value
func (o *VcenterVmIdentityInfo) GetInstanceUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InstanceUuid
}

// GetInstanceUuidOk returns a tuple with the InstanceUuid field value
// and a boolean to check if the value has been set.
func (o *VcenterVmIdentityInfo) GetInstanceUuidOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InstanceUuid, true
}

// SetInstanceUuid sets field value
func (o *VcenterVmIdentityInfo) SetInstanceUuid(v string) {
	o.InstanceUuid = v
}

func (o VcenterVmIdentityInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["bios_uuid"] = o.BiosUuid
	}
	if true {
		toSerialize["instance_uuid"] = o.InstanceUuid
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmIdentityInfo struct {
	value *VcenterVmIdentityInfo
	isSet bool
}

func (v NullableVcenterVmIdentityInfo) Get() *VcenterVmIdentityInfo {
	return v.value
}

func (v *NullableVcenterVmIdentityInfo) Set(val *VcenterVmIdentityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmIdentityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmIdentityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmIdentityInfo(val *VcenterVmIdentityInfo) *NullableVcenterVmIdentityInfo {
	return &NullableVcenterVmIdentityInfo{value: val, isSet: true}
}

func (v NullableVcenterVmIdentityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmIdentityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

