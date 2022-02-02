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

// VcenterContentRegistriesHarborProjectsScope The {@name Scope} {@term enumerated type} in a project defines access levels of the project.
type VcenterContentRegistriesHarborProjectsScope string

// List of vcenter.content.registries.harbor.projects.scope
const (
	VCENTERCONTENTREGISTRIESHARBORPROJECTSSCOPE_PUBLIC VcenterContentRegistriesHarborProjectsScope = "PUBLIC"
	VCENTERCONTENTREGISTRIESHARBORPROJECTSSCOPE_PRIVATE VcenterContentRegistriesHarborProjectsScope = "PRIVATE"
)

// All allowed values of VcenterContentRegistriesHarborProjectsScope enum
var AllowedVcenterContentRegistriesHarborProjectsScopeEnumValues = []VcenterContentRegistriesHarborProjectsScope{
	"PUBLIC",
	"PRIVATE",
}

func (v *VcenterContentRegistriesHarborProjectsScope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterContentRegistriesHarborProjectsScope(value)
	for _, existing := range AllowedVcenterContentRegistriesHarborProjectsScopeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterContentRegistriesHarborProjectsScope", value)
}

// NewVcenterContentRegistriesHarborProjectsScopeFromValue returns a pointer to a valid VcenterContentRegistriesHarborProjectsScope
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterContentRegistriesHarborProjectsScopeFromValue(v string) (*VcenterContentRegistriesHarborProjectsScope, error) {
	ev := VcenterContentRegistriesHarborProjectsScope(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterContentRegistriesHarborProjectsScope: valid values are %v", v, AllowedVcenterContentRegistriesHarborProjectsScopeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterContentRegistriesHarborProjectsScope) IsValid() bool {
	for _, existing := range AllowedVcenterContentRegistriesHarborProjectsScopeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.content.registries.harbor.projects.scope value
func (v VcenterContentRegistriesHarborProjectsScope) Ptr() *VcenterContentRegistriesHarborProjectsScope {
	return &v
}

type NullableVcenterContentRegistriesHarborProjectsScope struct {
	value *VcenterContentRegistriesHarborProjectsScope
	isSet bool
}

func (v NullableVcenterContentRegistriesHarborProjectsScope) Get() *VcenterContentRegistriesHarborProjectsScope {
	return v.value
}

func (v *NullableVcenterContentRegistriesHarborProjectsScope) Set(val *VcenterContentRegistriesHarborProjectsScope) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterContentRegistriesHarborProjectsScope) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterContentRegistriesHarborProjectsScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterContentRegistriesHarborProjectsScope(val *VcenterContentRegistriesHarborProjectsScope) *NullableVcenterContentRegistriesHarborProjectsScope {
	return &NullableVcenterContentRegistriesHarborProjectsScope{value: val, isSet: true}
}

func (v NullableVcenterContentRegistriesHarborProjectsScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterContentRegistriesHarborProjectsScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

