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

// VcenterVmToolsVersionStatus The Tools.VersionStatus enumerated type defines the version status types of VMware Tools installed in the guest operating system.
type VcenterVmToolsVersionStatus string

// List of vcenter.vm.tools.version_status
const (
	VCENTERVMTOOLSVERSIONSTATUS_NOT_INSTALLED VcenterVmToolsVersionStatus = "NOT_INSTALLED"
	VCENTERVMTOOLSVERSIONSTATUS_CURRENT VcenterVmToolsVersionStatus = "CURRENT"
	VCENTERVMTOOLSVERSIONSTATUS_UNMANAGED VcenterVmToolsVersionStatus = "UNMANAGED"
	VCENTERVMTOOLSVERSIONSTATUS_TOO_OLD_UNSUPPORTED VcenterVmToolsVersionStatus = "TOO_OLD_UNSUPPORTED"
	VCENTERVMTOOLSVERSIONSTATUS_SUPPORTED_OLD VcenterVmToolsVersionStatus = "SUPPORTED_OLD"
	VCENTERVMTOOLSVERSIONSTATUS_SUPPORTED_NEW VcenterVmToolsVersionStatus = "SUPPORTED_NEW"
	VCENTERVMTOOLSVERSIONSTATUS_TOO_NEW VcenterVmToolsVersionStatus = "TOO_NEW"
	VCENTERVMTOOLSVERSIONSTATUS_BLACKLISTED VcenterVmToolsVersionStatus = "BLACKLISTED"
)

// All allowed values of VcenterVmToolsVersionStatus enum
var AllowedVcenterVmToolsVersionStatusEnumValues = []VcenterVmToolsVersionStatus{
	"NOT_INSTALLED",
	"CURRENT",
	"UNMANAGED",
	"TOO_OLD_UNSUPPORTED",
	"SUPPORTED_OLD",
	"SUPPORTED_NEW",
	"TOO_NEW",
	"BLACKLISTED",
}

func (v *VcenterVmToolsVersionStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmToolsVersionStatus(value)
	for _, existing := range AllowedVcenterVmToolsVersionStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmToolsVersionStatus", value)
}

// NewVcenterVmToolsVersionStatusFromValue returns a pointer to a valid VcenterVmToolsVersionStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmToolsVersionStatusFromValue(v string) (*VcenterVmToolsVersionStatus, error) {
	ev := VcenterVmToolsVersionStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmToolsVersionStatus: valid values are %v", v, AllowedVcenterVmToolsVersionStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmToolsVersionStatus) IsValid() bool {
	for _, existing := range AllowedVcenterVmToolsVersionStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.tools.version_status value
func (v VcenterVmToolsVersionStatus) Ptr() *VcenterVmToolsVersionStatus {
	return &v
}

type NullableVcenterVmToolsVersionStatus struct {
	value *VcenterVmToolsVersionStatus
	isSet bool
}

func (v NullableVcenterVmToolsVersionStatus) Get() *VcenterVmToolsVersionStatus {
	return v.value
}

func (v *NullableVcenterVmToolsVersionStatus) Set(val *VcenterVmToolsVersionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmToolsVersionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmToolsVersionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmToolsVersionStatus(val *VcenterVmToolsVersionStatus) *NullableVcenterVmToolsVersionStatus {
	return &NullableVcenterVmToolsVersionStatus{value: val, isSet: true}
}

func (v NullableVcenterVmToolsVersionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmToolsVersionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
