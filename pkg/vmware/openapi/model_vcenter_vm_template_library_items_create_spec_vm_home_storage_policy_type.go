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

// VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType Policy type for the virtual machine template's configuration and log files.
type VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType string

// List of vcenter.vm_template.library_items.create_spec_vm_home_storage_policy.type
const (
	VCENTERVMTEMPLATELIBRARYITEMSCREATESPECVMHOMESTORAGEPOLICYTYPE_USE_SPECIFIED_POLICY VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType = "USE_SPECIFIED_POLICY"
)

// All allowed values of VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType enum
var AllowedVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyTypeEnumValues = []VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType{
	"USE_SPECIFIED_POLICY",
}

func (v *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType(value)
	for _, existing := range AllowedVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType", value)
}

// NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyTypeFromValue returns a pointer to a valid VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyTypeFromValue(v string) (*VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType, error) {
	ev := VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType: valid values are %v", v, AllowedVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) IsValid() bool {
	for _, existing := range AllowedVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm_template.library_items.create_spec_vm_home_storage_policy.type value
func (v VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) Ptr() *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType {
	return &v
}

type NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType struct {
	value *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) Get() *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) Set(val *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType(val *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType {
	return &NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

