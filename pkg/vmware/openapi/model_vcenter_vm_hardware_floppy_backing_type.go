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

// VcenterVmHardwareFloppyBackingType The Floppy.BackingType enumerated type defines the valid backing types for a virtual floppy drive.
type VcenterVmHardwareFloppyBackingType string

// List of vcenter.vm.hardware.floppy.backing_type
const (
	VCENTERVMHARDWAREFLOPPYBACKINGTYPE_IMAGE_FILE VcenterVmHardwareFloppyBackingType = "IMAGE_FILE"
	VCENTERVMHARDWAREFLOPPYBACKINGTYPE_HOST_DEVICE VcenterVmHardwareFloppyBackingType = "HOST_DEVICE"
	VCENTERVMHARDWAREFLOPPYBACKINGTYPE_CLIENT_DEVICE VcenterVmHardwareFloppyBackingType = "CLIENT_DEVICE"
)

// All allowed values of VcenterVmHardwareFloppyBackingType enum
var AllowedVcenterVmHardwareFloppyBackingTypeEnumValues = []VcenterVmHardwareFloppyBackingType{
	"IMAGE_FILE",
	"HOST_DEVICE",
	"CLIENT_DEVICE",
}

func (v *VcenterVmHardwareFloppyBackingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmHardwareFloppyBackingType(value)
	for _, existing := range AllowedVcenterVmHardwareFloppyBackingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmHardwareFloppyBackingType", value)
}

// NewVcenterVmHardwareFloppyBackingTypeFromValue returns a pointer to a valid VcenterVmHardwareFloppyBackingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmHardwareFloppyBackingTypeFromValue(v string) (*VcenterVmHardwareFloppyBackingType, error) {
	ev := VcenterVmHardwareFloppyBackingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmHardwareFloppyBackingType: valid values are %v", v, AllowedVcenterVmHardwareFloppyBackingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmHardwareFloppyBackingType) IsValid() bool {
	for _, existing := range AllowedVcenterVmHardwareFloppyBackingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.hardware.floppy.backing_type value
func (v VcenterVmHardwareFloppyBackingType) Ptr() *VcenterVmHardwareFloppyBackingType {
	return &v
}

type NullableVcenterVmHardwareFloppyBackingType struct {
	value *VcenterVmHardwareFloppyBackingType
	isSet bool
}

func (v NullableVcenterVmHardwareFloppyBackingType) Get() *VcenterVmHardwareFloppyBackingType {
	return v.value
}

func (v *NullableVcenterVmHardwareFloppyBackingType) Set(val *VcenterVmHardwareFloppyBackingType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareFloppyBackingType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareFloppyBackingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareFloppyBackingType(val *VcenterVmHardwareFloppyBackingType) *NullableVcenterVmHardwareFloppyBackingType {
	return &NullableVcenterVmHardwareFloppyBackingType{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareFloppyBackingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareFloppyBackingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

