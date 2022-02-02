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

// VcenterVmHardwareEthernetBackingType The Ethernet.BackingType enumerated type defines the valid backing types for a virtual Ethernet adapter.
type VcenterVmHardwareEthernetBackingType string

// List of vcenter.vm.hardware.ethernet.backing_type
const (
	VCENTERVMHARDWAREETHERNETBACKINGTYPE_STANDARD_PORTGROUP VcenterVmHardwareEthernetBackingType = "STANDARD_PORTGROUP"
	VCENTERVMHARDWAREETHERNETBACKINGTYPE_HOST_DEVICE VcenterVmHardwareEthernetBackingType = "HOST_DEVICE"
	VCENTERVMHARDWAREETHERNETBACKINGTYPE_DISTRIBUTED_PORTGROUP VcenterVmHardwareEthernetBackingType = "DISTRIBUTED_PORTGROUP"
	VCENTERVMHARDWAREETHERNETBACKINGTYPE_OPAQUE_NETWORK VcenterVmHardwareEthernetBackingType = "OPAQUE_NETWORK"
)

// All allowed values of VcenterVmHardwareEthernetBackingType enum
var AllowedVcenterVmHardwareEthernetBackingTypeEnumValues = []VcenterVmHardwareEthernetBackingType{
	"STANDARD_PORTGROUP",
	"HOST_DEVICE",
	"DISTRIBUTED_PORTGROUP",
	"OPAQUE_NETWORK",
}

func (v *VcenterVmHardwareEthernetBackingType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmHardwareEthernetBackingType(value)
	for _, existing := range AllowedVcenterVmHardwareEthernetBackingTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmHardwareEthernetBackingType", value)
}

// NewVcenterVmHardwareEthernetBackingTypeFromValue returns a pointer to a valid VcenterVmHardwareEthernetBackingType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmHardwareEthernetBackingTypeFromValue(v string) (*VcenterVmHardwareEthernetBackingType, error) {
	ev := VcenterVmHardwareEthernetBackingType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmHardwareEthernetBackingType: valid values are %v", v, AllowedVcenterVmHardwareEthernetBackingTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmHardwareEthernetBackingType) IsValid() bool {
	for _, existing := range AllowedVcenterVmHardwareEthernetBackingTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.hardware.ethernet.backing_type value
func (v VcenterVmHardwareEthernetBackingType) Ptr() *VcenterVmHardwareEthernetBackingType {
	return &v
}

type NullableVcenterVmHardwareEthernetBackingType struct {
	value *VcenterVmHardwareEthernetBackingType
	isSet bool
}

func (v NullableVcenterVmHardwareEthernetBackingType) Get() *VcenterVmHardwareEthernetBackingType {
	return v.value
}

func (v *NullableVcenterVmHardwareEthernetBackingType) Set(val *VcenterVmHardwareEthernetBackingType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareEthernetBackingType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareEthernetBackingType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareEthernetBackingType(val *VcenterVmHardwareEthernetBackingType) *NullableVcenterVmHardwareEthernetBackingType {
	return &NullableVcenterVmHardwareEthernetBackingType{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareEthernetBackingType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareEthernetBackingType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

