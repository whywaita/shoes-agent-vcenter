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

// VcenterVmHardwareCdromBackingSpec struct for VcenterVmHardwareCdromBackingSpec
type VcenterVmHardwareCdromBackingSpec struct {
	Type VcenterVmHardwareCdromBackingType `json:"type"`
	// Path of the image file that should be used as the virtual CD-ROM device backing. This field is optional and it is only relevant when the value of Cdrom.BackingSpec.type is ISO_FILE.
	IsoFile *string `json:"iso_file,omitempty"`
	// Name of the device that should be used as the virtual CD-ROM device backing. If unset, the virtual CD-ROM device will be configured to automatically detect a suitable host device.
	HostDevice *string `json:"host_device,omitempty"`
	DeviceAccessType *VcenterVmHardwareCdromDeviceAccessType `json:"device_access_type,omitempty"`
}

// NewVcenterVmHardwareCdromBackingSpec instantiates a new VcenterVmHardwareCdromBackingSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareCdromBackingSpec(type_ VcenterVmHardwareCdromBackingType) *VcenterVmHardwareCdromBackingSpec {
	this := VcenterVmHardwareCdromBackingSpec{}
	this.Type = type_
	return &this
}

// NewVcenterVmHardwareCdromBackingSpecWithDefaults instantiates a new VcenterVmHardwareCdromBackingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareCdromBackingSpecWithDefaults() *VcenterVmHardwareCdromBackingSpec {
	this := VcenterVmHardwareCdromBackingSpec{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmHardwareCdromBackingSpec) GetType() VcenterVmHardwareCdromBackingType {
	if o == nil {
		var ret VcenterVmHardwareCdromBackingType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromBackingSpec) GetTypeOk() (*VcenterVmHardwareCdromBackingType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmHardwareCdromBackingSpec) SetType(v VcenterVmHardwareCdromBackingType) {
	o.Type = v
}

// GetIsoFile returns the IsoFile field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromBackingSpec) GetIsoFile() string {
	if o == nil || o.IsoFile == nil {
		var ret string
		return ret
	}
	return *o.IsoFile
}

// GetIsoFileOk returns a tuple with the IsoFile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromBackingSpec) GetIsoFileOk() (*string, bool) {
	if o == nil || o.IsoFile == nil {
		return nil, false
	}
	return o.IsoFile, true
}

// HasIsoFile returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromBackingSpec) HasIsoFile() bool {
	if o != nil && o.IsoFile != nil {
		return true
	}

	return false
}

// SetIsoFile gets a reference to the given string and assigns it to the IsoFile field.
func (o *VcenterVmHardwareCdromBackingSpec) SetIsoFile(v string) {
	o.IsoFile = &v
}

// GetHostDevice returns the HostDevice field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromBackingSpec) GetHostDevice() string {
	if o == nil || o.HostDevice == nil {
		var ret string
		return ret
	}
	return *o.HostDevice
}

// GetHostDeviceOk returns a tuple with the HostDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromBackingSpec) GetHostDeviceOk() (*string, bool) {
	if o == nil || o.HostDevice == nil {
		return nil, false
	}
	return o.HostDevice, true
}

// HasHostDevice returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromBackingSpec) HasHostDevice() bool {
	if o != nil && o.HostDevice != nil {
		return true
	}

	return false
}

// SetHostDevice gets a reference to the given string and assigns it to the HostDevice field.
func (o *VcenterVmHardwareCdromBackingSpec) SetHostDevice(v string) {
	o.HostDevice = &v
}

// GetDeviceAccessType returns the DeviceAccessType field value if set, zero value otherwise.
func (o *VcenterVmHardwareCdromBackingSpec) GetDeviceAccessType() VcenterVmHardwareCdromDeviceAccessType {
	if o == nil || o.DeviceAccessType == nil {
		var ret VcenterVmHardwareCdromDeviceAccessType
		return ret
	}
	return *o.DeviceAccessType
}

// GetDeviceAccessTypeOk returns a tuple with the DeviceAccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromBackingSpec) GetDeviceAccessTypeOk() (*VcenterVmHardwareCdromDeviceAccessType, bool) {
	if o == nil || o.DeviceAccessType == nil {
		return nil, false
	}
	return o.DeviceAccessType, true
}

// HasDeviceAccessType returns a boolean if a field has been set.
func (o *VcenterVmHardwareCdromBackingSpec) HasDeviceAccessType() bool {
	if o != nil && o.DeviceAccessType != nil {
		return true
	}

	return false
}

// SetDeviceAccessType gets a reference to the given VcenterVmHardwareCdromDeviceAccessType and assigns it to the DeviceAccessType field.
func (o *VcenterVmHardwareCdromBackingSpec) SetDeviceAccessType(v VcenterVmHardwareCdromDeviceAccessType) {
	o.DeviceAccessType = &v
}

func (o VcenterVmHardwareCdromBackingSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.IsoFile != nil {
		toSerialize["iso_file"] = o.IsoFile
	}
	if o.HostDevice != nil {
		toSerialize["host_device"] = o.HostDevice
	}
	if o.DeviceAccessType != nil {
		toSerialize["device_access_type"] = o.DeviceAccessType
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareCdromBackingSpec struct {
	value *VcenterVmHardwareCdromBackingSpec
	isSet bool
}

func (v NullableVcenterVmHardwareCdromBackingSpec) Get() *VcenterVmHardwareCdromBackingSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareCdromBackingSpec) Set(val *VcenterVmHardwareCdromBackingSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareCdromBackingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareCdromBackingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareCdromBackingSpec(val *VcenterVmHardwareCdromBackingSpec) *NullableVcenterVmHardwareCdromBackingSpec {
	return &NullableVcenterVmHardwareCdromBackingSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareCdromBackingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareCdromBackingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


