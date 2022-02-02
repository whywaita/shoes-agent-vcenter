/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// VcenterVmHardwareBootDeviceType The Device.Type enumerated type defines the valid device types that may be used as bootable devices.
type VcenterVmHardwareBootDeviceType string

// List of vcenter.vm.hardware.boot.device.type
const (
	VCENTERVMHARDWAREBOOTDEVICETYPE_CDROM VcenterVmHardwareBootDeviceType = "CDROM"
	VCENTERVMHARDWAREBOOTDEVICETYPE_DISK VcenterVmHardwareBootDeviceType = "DISK"
	VCENTERVMHARDWAREBOOTDEVICETYPE_ETHERNET VcenterVmHardwareBootDeviceType = "ETHERNET"
	VCENTERVMHARDWAREBOOTDEVICETYPE_FLOPPY VcenterVmHardwareBootDeviceType = "FLOPPY"
)

// All allowed values of VcenterVmHardwareBootDeviceType enum
var AllowedVcenterVmHardwareBootDeviceTypeEnumValues = []VcenterVmHardwareBootDeviceType{
	"CDROM",
	"DISK",
	"ETHERNET",
	"FLOPPY",
}

func (v *VcenterVmHardwareBootDeviceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmHardwareBootDeviceType(value)
	for _, existing := range AllowedVcenterVmHardwareBootDeviceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmHardwareBootDeviceType", value)
}

// NewVcenterVmHardwareBootDeviceTypeFromValue returns a pointer to a valid VcenterVmHardwareBootDeviceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmHardwareBootDeviceTypeFromValue(v string) (*VcenterVmHardwareBootDeviceType, error) {
	ev := VcenterVmHardwareBootDeviceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmHardwareBootDeviceType: valid values are %v", v, AllowedVcenterVmHardwareBootDeviceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmHardwareBootDeviceType) IsValid() bool {
	for _, existing := range AllowedVcenterVmHardwareBootDeviceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.hardware.boot.device.type value
func (v VcenterVmHardwareBootDeviceType) Ptr() *VcenterVmHardwareBootDeviceType {
	return &v
}

type NullableVcenterVmHardwareBootDeviceType struct {
	value *VcenterVmHardwareBootDeviceType
	isSet bool
}

func (v NullableVcenterVmHardwareBootDeviceType) Get() *VcenterVmHardwareBootDeviceType {
	return v.value
}

func (v *NullableVcenterVmHardwareBootDeviceType) Set(val *VcenterVmHardwareBootDeviceType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareBootDeviceType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareBootDeviceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareBootDeviceType(val *VcenterVmHardwareBootDeviceType) *NullableVcenterVmHardwareBootDeviceType {
	return &NullableVcenterVmHardwareBootDeviceType{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareBootDeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareBootDeviceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

