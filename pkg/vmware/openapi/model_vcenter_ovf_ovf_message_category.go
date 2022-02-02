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

// VcenterOvfOvfMessageCategory The {@name Category} {@term enumerated type} defines the categories of messages (see {@link OvfMessage}).
type VcenterOvfOvfMessageCategory string

// List of vcenter.ovf.ovf_message.category
const (
	VCENTEROVFOVFMESSAGECATEGORY_VALIDATION VcenterOvfOvfMessageCategory = "VALIDATION"
	VCENTEROVFOVFMESSAGECATEGORY_INPUT VcenterOvfOvfMessageCategory = "INPUT"
	VCENTEROVFOVFMESSAGECATEGORY_SERVER VcenterOvfOvfMessageCategory = "SERVER"
)

// All allowed values of VcenterOvfOvfMessageCategory enum
var AllowedVcenterOvfOvfMessageCategoryEnumValues = []VcenterOvfOvfMessageCategory{
	"VALIDATION",
	"INPUT",
	"SERVER",
}

func (v *VcenterOvfOvfMessageCategory) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterOvfOvfMessageCategory(value)
	for _, existing := range AllowedVcenterOvfOvfMessageCategoryEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterOvfOvfMessageCategory", value)
}

// NewVcenterOvfOvfMessageCategoryFromValue returns a pointer to a valid VcenterOvfOvfMessageCategory
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterOvfOvfMessageCategoryFromValue(v string) (*VcenterOvfOvfMessageCategory, error) {
	ev := VcenterOvfOvfMessageCategory(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterOvfOvfMessageCategory: valid values are %v", v, AllowedVcenterOvfOvfMessageCategoryEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterOvfOvfMessageCategory) IsValid() bool {
	for _, existing := range AllowedVcenterOvfOvfMessageCategoryEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.ovf.ovf_message.category value
func (v VcenterOvfOvfMessageCategory) Ptr() *VcenterOvfOvfMessageCategory {
	return &v
}

type NullableVcenterOvfOvfMessageCategory struct {
	value *VcenterOvfOvfMessageCategory
	isSet bool
}

func (v NullableVcenterOvfOvfMessageCategory) Get() *VcenterOvfOvfMessageCategory {
	return v.value
}

func (v *NullableVcenterOvfOvfMessageCategory) Set(val *VcenterOvfOvfMessageCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfOvfMessageCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfOvfMessageCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfOvfMessageCategory(val *VcenterOvfOvfMessageCategory) *NullableVcenterOvfOvfMessageCategory {
	return &NullableVcenterOvfOvfMessageCategory{value: val, isSet: true}
}

func (v NullableVcenterOvfOvfMessageCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfOvfMessageCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

