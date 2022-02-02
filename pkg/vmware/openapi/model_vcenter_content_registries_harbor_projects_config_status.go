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

// VcenterContentRegistriesHarborProjectsConfigStatus The {@name ConfigStatus} {@term enumerated type} describes the status of reaching the desired configuration state for the Harbor project.
type VcenterContentRegistriesHarborProjectsConfigStatus string

// List of vcenter.content.registries.harbor.projects.config_status
const (
	VCENTERCONTENTREGISTRIESHARBORPROJECTSCONFIGSTATUS_PENDING VcenterContentRegistriesHarborProjectsConfigStatus = "PENDING"
	VCENTERCONTENTREGISTRIESHARBORPROJECTSCONFIGSTATUS_REMOVING VcenterContentRegistriesHarborProjectsConfigStatus = "REMOVING"
	VCENTERCONTENTREGISTRIESHARBORPROJECTSCONFIGSTATUS_ERROR VcenterContentRegistriesHarborProjectsConfigStatus = "ERROR"
	VCENTERCONTENTREGISTRIESHARBORPROJECTSCONFIGSTATUS_READY VcenterContentRegistriesHarborProjectsConfigStatus = "READY"
)

// All allowed values of VcenterContentRegistriesHarborProjectsConfigStatus enum
var AllowedVcenterContentRegistriesHarborProjectsConfigStatusEnumValues = []VcenterContentRegistriesHarborProjectsConfigStatus{
	"PENDING",
	"REMOVING",
	"ERROR",
	"READY",
}

func (v *VcenterContentRegistriesHarborProjectsConfigStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterContentRegistriesHarborProjectsConfigStatus(value)
	for _, existing := range AllowedVcenterContentRegistriesHarborProjectsConfigStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterContentRegistriesHarborProjectsConfigStatus", value)
}

// NewVcenterContentRegistriesHarborProjectsConfigStatusFromValue returns a pointer to a valid VcenterContentRegistriesHarborProjectsConfigStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterContentRegistriesHarborProjectsConfigStatusFromValue(v string) (*VcenterContentRegistriesHarborProjectsConfigStatus, error) {
	ev := VcenterContentRegistriesHarborProjectsConfigStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterContentRegistriesHarborProjectsConfigStatus: valid values are %v", v, AllowedVcenterContentRegistriesHarborProjectsConfigStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterContentRegistriesHarborProjectsConfigStatus) IsValid() bool {
	for _, existing := range AllowedVcenterContentRegistriesHarborProjectsConfigStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.content.registries.harbor.projects.config_status value
func (v VcenterContentRegistriesHarborProjectsConfigStatus) Ptr() *VcenterContentRegistriesHarborProjectsConfigStatus {
	return &v
}

type NullableVcenterContentRegistriesHarborProjectsConfigStatus struct {
	value *VcenterContentRegistriesHarborProjectsConfigStatus
	isSet bool
}

func (v NullableVcenterContentRegistriesHarborProjectsConfigStatus) Get() *VcenterContentRegistriesHarborProjectsConfigStatus {
	return v.value
}

func (v *NullableVcenterContentRegistriesHarborProjectsConfigStatus) Set(val *VcenterContentRegistriesHarborProjectsConfigStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterContentRegistriesHarborProjectsConfigStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterContentRegistriesHarborProjectsConfigStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterContentRegistriesHarborProjectsConfigStatus(val *VcenterContentRegistriesHarborProjectsConfigStatus) *NullableVcenterContentRegistriesHarborProjectsConfigStatus {
	return &NullableVcenterContentRegistriesHarborProjectsConfigStatus{value: val, isSet: true}
}

func (v NullableVcenterContentRegistriesHarborProjectsConfigStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterContentRegistriesHarborProjectsConfigStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
