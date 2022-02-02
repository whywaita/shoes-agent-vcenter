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

// VcenterNamespaceManagementSupervisorServicesVersionsContentType The Versions.ContentType enumerated type defines the type of content that describes the format of Supervisor Service version definition.
type VcenterNamespaceManagementSupervisorServicesVersionsContentType string

// List of VcenterNamespaceManagementSupervisorServicesVersionsContentType
const (
	VCENTERNAMESPACEMANAGEMENTSUPERVISORSERVICESVERSIONSCONTENTTYPE_VSPHERE_APPS_YAML VcenterNamespaceManagementSupervisorServicesVersionsContentType = "VSPHERE_APPS_YAML"
	VCENTERNAMESPACEMANAGEMENTSUPERVISORSERVICESVERSIONSCONTENTTYPE_CUSTOM_YAML VcenterNamespaceManagementSupervisorServicesVersionsContentType = "CUSTOM_YAML"
)

// All allowed values of VcenterNamespaceManagementSupervisorServicesVersionsContentType enum
var AllowedVcenterNamespaceManagementSupervisorServicesVersionsContentTypeEnumValues = []VcenterNamespaceManagementSupervisorServicesVersionsContentType{
	"VSPHERE_APPS_YAML",
	"CUSTOM_YAML",
}

func (v *VcenterNamespaceManagementSupervisorServicesVersionsContentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterNamespaceManagementSupervisorServicesVersionsContentType(value)
	for _, existing := range AllowedVcenterNamespaceManagementSupervisorServicesVersionsContentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterNamespaceManagementSupervisorServicesVersionsContentType", value)
}

// NewVcenterNamespaceManagementSupervisorServicesVersionsContentTypeFromValue returns a pointer to a valid VcenterNamespaceManagementSupervisorServicesVersionsContentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterNamespaceManagementSupervisorServicesVersionsContentTypeFromValue(v string) (*VcenterNamespaceManagementSupervisorServicesVersionsContentType, error) {
	ev := VcenterNamespaceManagementSupervisorServicesVersionsContentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterNamespaceManagementSupervisorServicesVersionsContentType: valid values are %v", v, AllowedVcenterNamespaceManagementSupervisorServicesVersionsContentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterNamespaceManagementSupervisorServicesVersionsContentType) IsValid() bool {
	for _, existing := range AllowedVcenterNamespaceManagementSupervisorServicesVersionsContentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterNamespaceManagementSupervisorServicesVersionsContentType value
func (v VcenterNamespaceManagementSupervisorServicesVersionsContentType) Ptr() *VcenterNamespaceManagementSupervisorServicesVersionsContentType {
	return &v
}

type NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType struct {
	value *VcenterNamespaceManagementSupervisorServicesVersionsContentType
	isSet bool
}

func (v NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType) Get() *VcenterNamespaceManagementSupervisorServicesVersionsContentType {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType) Set(val *VcenterNamespaceManagementSupervisorServicesVersionsContentType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSupervisorServicesVersionsContentType(val *VcenterNamespaceManagementSupervisorServicesVersionsContentType) *NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType {
	return &NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSupervisorServicesVersionsContentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
