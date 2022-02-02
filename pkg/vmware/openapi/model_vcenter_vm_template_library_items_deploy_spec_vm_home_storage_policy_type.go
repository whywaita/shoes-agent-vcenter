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

// VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType Policy type for the deployed virtual machine's configuration and log files.
type VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType string

// List of vcenter.vm_template.library_items.deploy_spec_vm_home_storage_policy.type
const (
	VCENTERVMTEMPLATELIBRARYITEMSDEPLOYSPECVMHOMESTORAGEPOLICYTYPE_SPECIFIED_POLICY VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType = "USE_SPECIFIED_POLICY"
	VCENTERVMTEMPLATELIBRARYITEMSDEPLOYSPECVMHOMESTORAGEPOLICYTYPE_SOURCE_POLICY VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType = "USE_SOURCE_POLICY"
)

// All allowed values of VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType enum
var AllowedVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyTypeEnumValues = []VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType{
	"USE_SPECIFIED_POLICY",
	"USE_SOURCE_POLICY",
}

func (v *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType(value)
	for _, existing := range AllowedVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType", value)
}

// NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyTypeFromValue returns a pointer to a valid VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyTypeFromValue(v string) (*VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType, error) {
	ev := VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType: valid values are %v", v, AllowedVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) IsValid() bool {
	for _, existing := range AllowedVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm_template.library_items.deploy_spec_vm_home_storage_policy.type value
func (v VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) Ptr() *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType {
	return &v
}

type NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType struct {
	value *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) Get() *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) Set(val *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType(val *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) *NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType {
	return &NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
