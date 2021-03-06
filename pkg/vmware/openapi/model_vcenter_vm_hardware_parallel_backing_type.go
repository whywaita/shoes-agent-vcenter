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

// VcenterVmHardwareParallelBackingType The Parallel.BackingType enumerated type defines the valid backing types for a virtual parallel port.
type VcenterVmHardwareParallelBackingType string

// List of vcenter.vm.hardware.parallel.backing_type
const (
	VCENTERVMHARDWAREPARALLELBACKINGTYPE_FILE VcenterVmHardwareParallelBackingType = "FILE"
	VCENTERVMHARDWAREPARALLELBACKINGTYPE_HOST_DEVICE VcenterVmHardwareParallelBackingType = "HOST_DEVICE"
)

// All allowed values of VcenterVmHardwareParallelBackingType enum
var AllowedVcenterVmHardwareParallelBackingTypeEnumValues = []VcenterVmHardwareParallelBackingType{
	"FILE",
	"HOST_DEVICE",
}

func (v *VcenterVmHardwareParallelBackingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmHardwareParallelBackingType(value)
	for _, existing := range AllowedVcenterVmHardwareParallelBackingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmHardwareParallelBackingType", value)
}

// NewVcenterVmHardwareParallelBackingTypeFromValue returns a pointer to a valid VcenterVmHardwareParallelBackingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmHardwareParallelBackingTypeFromValue(v string) (*VcenterVmHardwareParallelBackingType, error) {
	ev := VcenterVmHardwareParallelBackingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmHardwareParallelBackingType: valid values are %v", v, AllowedVcenterVmHardwareParallelBackingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmHardwareParallelBackingType) IsValid() bool {
	for _, existing := range AllowedVcenterVmHardwareParallelBackingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.hardware.parallel.backing_type value
func (v VcenterVmHardwareParallelBackingType) Ptr() *VcenterVmHardwareParallelBackingType {
	return &v
}

type NullableVcenterVmHardwareParallelBackingType struct {
	value *VcenterVmHardwareParallelBackingType
	isSet bool
}

func (v NullableVcenterVmHardwareParallelBackingType) Get() *VcenterVmHardwareParallelBackingType {
	return v.value
}

func (v *NullableVcenterVmHardwareParallelBackingType) Set(val *VcenterVmHardwareParallelBackingType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareParallelBackingType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareParallelBackingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareParallelBackingType(val *VcenterVmHardwareParallelBackingType) *NullableVcenterVmHardwareParallelBackingType {
	return &NullableVcenterVmHardwareParallelBackingType{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareParallelBackingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareParallelBackingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

