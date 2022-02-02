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

// VcenterVmGuestCredentialsType Types of guest credentials
type VcenterVmGuestCredentialsType string

// List of VcenterVmGuestCredentialsType
const (
	VCENTERVMGUESTCREDENTIALSTYPE_USERNAME_PASSWORD VcenterVmGuestCredentialsType = "USERNAME_PASSWORD"
	VCENTERVMGUESTCREDENTIALSTYPE_SAML_BEARER_TOKEN VcenterVmGuestCredentialsType = "SAML_BEARER_TOKEN"
)

// All allowed values of VcenterVmGuestCredentialsType enum
var AllowedVcenterVmGuestCredentialsTypeEnumValues = []VcenterVmGuestCredentialsType{
	"USERNAME_PASSWORD",
	"SAML_BEARER_TOKEN",
}

func (v *VcenterVmGuestCredentialsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmGuestCredentialsType(value)
	for _, existing := range AllowedVcenterVmGuestCredentialsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmGuestCredentialsType", value)
}

// NewVcenterVmGuestCredentialsTypeFromValue returns a pointer to a valid VcenterVmGuestCredentialsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmGuestCredentialsTypeFromValue(v string) (*VcenterVmGuestCredentialsType, error) {
	ev := VcenterVmGuestCredentialsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmGuestCredentialsType: valid values are %v", v, AllowedVcenterVmGuestCredentialsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmGuestCredentialsType) IsValid() bool {
	for _, existing := range AllowedVcenterVmGuestCredentialsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterVmGuestCredentialsType value
func (v VcenterVmGuestCredentialsType) Ptr() *VcenterVmGuestCredentialsType {
	return &v
}

type NullableVcenterVmGuestCredentialsType struct {
	value *VcenterVmGuestCredentialsType
	isSet bool
}

func (v NullableVcenterVmGuestCredentialsType) Get() *VcenterVmGuestCredentialsType {
	return v.value
}

func (v *NullableVcenterVmGuestCredentialsType) Set(val *VcenterVmGuestCredentialsType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestCredentialsType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestCredentialsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestCredentialsType(val *VcenterVmGuestCredentialsType) *NullableVcenterVmGuestCredentialsType {
	return &NullableVcenterVmGuestCredentialsType{value: val, isSet: true}
}

func (v NullableVcenterVmGuestCredentialsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestCredentialsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
