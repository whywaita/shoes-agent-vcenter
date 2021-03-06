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

// VcenterNamespaceManagementSupervisorServicesValidationStatus The SupervisorServices.ValidationStatus enumerated type defines the type of status for validating content of a Supervisor Service version.
type VcenterNamespaceManagementSupervisorServicesValidationStatus string

// List of VcenterNamespaceManagementSupervisorServicesValidationStatus
const (
	VCENTERNAMESPACEMANAGEMENTSUPERVISORSERVICESVALIDATIONSTATUS_VALID VcenterNamespaceManagementSupervisorServicesValidationStatus = "VALID"
	VCENTERNAMESPACEMANAGEMENTSUPERVISORSERVICESVALIDATIONSTATUS_VALID_WITH_WARNINGS VcenterNamespaceManagementSupervisorServicesValidationStatus = "VALID_WITH_WARNINGS"
	VCENTERNAMESPACEMANAGEMENTSUPERVISORSERVICESVALIDATIONSTATUS_INVALID VcenterNamespaceManagementSupervisorServicesValidationStatus = "INVALID"
)

// All allowed values of VcenterNamespaceManagementSupervisorServicesValidationStatus enum
var AllowedVcenterNamespaceManagementSupervisorServicesValidationStatusEnumValues = []VcenterNamespaceManagementSupervisorServicesValidationStatus{
	"VALID",
	"VALID_WITH_WARNINGS",
	"INVALID",
}

func (v *VcenterNamespaceManagementSupervisorServicesValidationStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterNamespaceManagementSupervisorServicesValidationStatus(value)
	for _, existing := range AllowedVcenterNamespaceManagementSupervisorServicesValidationStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterNamespaceManagementSupervisorServicesValidationStatus", value)
}

// NewVcenterNamespaceManagementSupervisorServicesValidationStatusFromValue returns a pointer to a valid VcenterNamespaceManagementSupervisorServicesValidationStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterNamespaceManagementSupervisorServicesValidationStatusFromValue(v string) (*VcenterNamespaceManagementSupervisorServicesValidationStatus, error) {
	ev := VcenterNamespaceManagementSupervisorServicesValidationStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterNamespaceManagementSupervisorServicesValidationStatus: valid values are %v", v, AllowedVcenterNamespaceManagementSupervisorServicesValidationStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterNamespaceManagementSupervisorServicesValidationStatus) IsValid() bool {
	for _, existing := range AllowedVcenterNamespaceManagementSupervisorServicesValidationStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterNamespaceManagementSupervisorServicesValidationStatus value
func (v VcenterNamespaceManagementSupervisorServicesValidationStatus) Ptr() *VcenterNamespaceManagementSupervisorServicesValidationStatus {
	return &v
}

type NullableVcenterNamespaceManagementSupervisorServicesValidationStatus struct {
	value *VcenterNamespaceManagementSupervisorServicesValidationStatus
	isSet bool
}

func (v NullableVcenterNamespaceManagementSupervisorServicesValidationStatus) Get() *VcenterNamespaceManagementSupervisorServicesValidationStatus {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesValidationStatus) Set(val *VcenterNamespaceManagementSupervisorServicesValidationStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSupervisorServicesValidationStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesValidationStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSupervisorServicesValidationStatus(val *VcenterNamespaceManagementSupervisorServicesValidationStatus) *NullableVcenterNamespaceManagementSupervisorServicesValidationStatus {
	return &NullableVcenterNamespaceManagementSupervisorServicesValidationStatus{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSupervisorServicesValidationStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesValidationStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

