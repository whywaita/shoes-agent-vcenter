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

// VcenterVmHardwareBootDeviceEntry struct for VcenterVmHardwareBootDeviceEntry
type VcenterVmHardwareBootDeviceEntry struct {
	Type VcenterVmHardwareBootDeviceType `json:"type"`
	// Virtual Ethernet device. Ethernet device to use as boot device for this entry. This field is optional and it is only relevant when the value of Device.Entry.type is ETHERNET. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Ethernet.
	Nic *string `json:"nic,omitempty"`
	// Virtual disk device. List of virtual disks in boot order. This field is optional and it is only relevant when the value of Device.Entry.type is DISK. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.vm.hardware.Disk.
	Disks *[]string `json:"disks,omitempty"`
}

// NewVcenterVmHardwareBootDeviceEntry instantiates a new VcenterVmHardwareBootDeviceEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareBootDeviceEntry(type_ VcenterVmHardwareBootDeviceType) *VcenterVmHardwareBootDeviceEntry {
	this := VcenterVmHardwareBootDeviceEntry{}
	this.Type = type_
	return &this
}

// NewVcenterVmHardwareBootDeviceEntryWithDefaults instantiates a new VcenterVmHardwareBootDeviceEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareBootDeviceEntryWithDefaults() *VcenterVmHardwareBootDeviceEntry {
	this := VcenterVmHardwareBootDeviceEntry{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmHardwareBootDeviceEntry) GetType() VcenterVmHardwareBootDeviceType {
	if o == nil {
		var ret VcenterVmHardwareBootDeviceType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootDeviceEntry) GetTypeOk() (*VcenterVmHardwareBootDeviceType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareBootDeviceEntry) SetType(v VcenterVmHardwareBootDeviceType) {
	o.Type = v
}

// GetNic returns the Nic field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootDeviceEntry) GetNic() string {
	if o == nil || o.Nic == nil {
		var ret string
		return ret
	}
	return *o.Nic
}

// GetNicOk returns a tuple with the Nic field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootDeviceEntry) GetNicOk() (*string, bool) {
	if o == nil || o.Nic == nil {
		return nil, false
	}
	return o.Nic, true
}

// HasNic returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootDeviceEntry) HasNic() bool {
	if o != nil && o.Nic != nil {
		return true
	}

	return false
}

// SetNic gets a reference to the given string and assigns it to the Nic field.
func (o *VcenterVmHardwareBootDeviceEntry) SetNic(v string) {
	o.Nic = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *VcenterVmHardwareBootDeviceEntry) GetDisks() []string {
	if o == nil || o.Disks == nil {
		var ret []string
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareBootDeviceEntry) GetDisksOk() (*[]string, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *VcenterVmHardwareBootDeviceEntry) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []string and assigns it to the Disks field.
func (o *VcenterVmHardwareBootDeviceEntry) SetDisks(v []string) {
	o.Disks = &v
}

func (o VcenterVmHardwareBootDeviceEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Nic != nil {
		toSerialize["nic"] = o.Nic
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareBootDeviceEntry struct {
	value *VcenterVmHardwareBootDeviceEntry
	isSet bool
}

func (v NullableVcenterVmHardwareBootDeviceEntry) Get() *VcenterVmHardwareBootDeviceEntry {
	return v.value
}

func (v *NullableVcenterVmHardwareBootDeviceEntry) Set(val *VcenterVmHardwareBootDeviceEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareBootDeviceEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareBootDeviceEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareBootDeviceEntry(val *VcenterVmHardwareBootDeviceEntry) *NullableVcenterVmHardwareBootDeviceEntry {
	return &NullableVcenterVmHardwareBootDeviceEntry{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareBootDeviceEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareBootDeviceEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


