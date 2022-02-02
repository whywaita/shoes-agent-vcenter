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

// VcenterVmToolsToolsInstallType The Tools.ToolsInstallType enumerated type defines the installation type of the Tools in the guest operating system.
type VcenterVmToolsToolsInstallType string

// List of vcenter.vm.tools.tools_install_type
const (
	VCENTERVMTOOLSTOOLSINSTALLTYPE_UNKNOWN VcenterVmToolsToolsInstallType = "UNKNOWN"
	VCENTERVMTOOLSTOOLSINSTALLTYPE_MSI VcenterVmToolsToolsInstallType = "MSI"
	VCENTERVMTOOLSTOOLSINSTALLTYPE_TAR VcenterVmToolsToolsInstallType = "TAR"
	VCENTERVMTOOLSTOOLSINSTALLTYPE_OSP VcenterVmToolsToolsInstallType = "OSP"
	VCENTERVMTOOLSTOOLSINSTALLTYPE_OPEN_VM_TOOLS VcenterVmToolsToolsInstallType = "OPEN_VM_TOOLS"
)

// All allowed values of VcenterVmToolsToolsInstallType enum
var AllowedVcenterVmToolsToolsInstallTypeEnumValues = []VcenterVmToolsToolsInstallType{
	"UNKNOWN",
	"MSI",
	"TAR",
	"OSP",
	"OPEN_VM_TOOLS",
}

func (v *VcenterVmToolsToolsInstallType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmToolsToolsInstallType(value)
	for _, existing := range AllowedVcenterVmToolsToolsInstallTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmToolsToolsInstallType", value)
}

// NewVcenterVmToolsToolsInstallTypeFromValue returns a pointer to a valid VcenterVmToolsToolsInstallType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmToolsToolsInstallTypeFromValue(v string) (*VcenterVmToolsToolsInstallType, error) {
	ev := VcenterVmToolsToolsInstallType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmToolsToolsInstallType: valid values are %v", v, AllowedVcenterVmToolsToolsInstallTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmToolsToolsInstallType) IsValid() bool {
	for _, existing := range AllowedVcenterVmToolsToolsInstallTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.tools.tools_install_type value
func (v VcenterVmToolsToolsInstallType) Ptr() *VcenterVmToolsToolsInstallType {
	return &v
}

type NullableVcenterVmToolsToolsInstallType struct {
	value *VcenterVmToolsToolsInstallType
	isSet bool
}

func (v NullableVcenterVmToolsToolsInstallType) Get() *VcenterVmToolsToolsInstallType {
	return v.value
}

func (v *NullableVcenterVmToolsToolsInstallType) Set(val *VcenterVmToolsToolsInstallType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmToolsToolsInstallType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmToolsToolsInstallType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmToolsToolsInstallType(val *VcenterVmToolsToolsInstallType) *NullableVcenterVmToolsToolsInstallType {
	return &NullableVcenterVmToolsToolsInstallType{value: val, isSet: true}
}

func (v NullableVcenterVmToolsToolsInstallType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmToolsToolsInstallType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
