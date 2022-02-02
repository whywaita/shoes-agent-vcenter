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

// VcenterVmStoragePolicyVmHomePolicySpecPolicyType The Policy.VmHomePolicySpec.PolicyType enumerated type defines the choices for how to specify the policy to be associated with the virtual machine home's directory.
type VcenterVmStoragePolicyVmHomePolicySpecPolicyType string

// List of vcenter.vm.storage.policy.vm_home_policy_spec.policy_type
const (
	VCENTERVMSTORAGEPOLICYVMHOMEPOLICYSPECPOLICYTYPE_SPECIFIED_POLICY VcenterVmStoragePolicyVmHomePolicySpecPolicyType = "USE_SPECIFIED_POLICY"
	VCENTERVMSTORAGEPOLICYVMHOMEPOLICYSPECPOLICYTYPE_DEFAULT_POLICY VcenterVmStoragePolicyVmHomePolicySpecPolicyType = "USE_DEFAULT_POLICY"
)

// All allowed values of VcenterVmStoragePolicyVmHomePolicySpecPolicyType enum
var AllowedVcenterVmStoragePolicyVmHomePolicySpecPolicyTypeEnumValues = []VcenterVmStoragePolicyVmHomePolicySpecPolicyType{
	"USE_SPECIFIED_POLICY",
	"USE_DEFAULT_POLICY",
}

func (v *VcenterVmStoragePolicyVmHomePolicySpecPolicyType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmStoragePolicyVmHomePolicySpecPolicyType(value)
	for _, existing := range AllowedVcenterVmStoragePolicyVmHomePolicySpecPolicyTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmStoragePolicyVmHomePolicySpecPolicyType", value)
}

// NewVcenterVmStoragePolicyVmHomePolicySpecPolicyTypeFromValue returns a pointer to a valid VcenterVmStoragePolicyVmHomePolicySpecPolicyType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmStoragePolicyVmHomePolicySpecPolicyTypeFromValue(v string) (*VcenterVmStoragePolicyVmHomePolicySpecPolicyType, error) {
	ev := VcenterVmStoragePolicyVmHomePolicySpecPolicyType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmStoragePolicyVmHomePolicySpecPolicyType: valid values are %v", v, AllowedVcenterVmStoragePolicyVmHomePolicySpecPolicyTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmStoragePolicyVmHomePolicySpecPolicyType) IsValid() bool {
	for _, existing := range AllowedVcenterVmStoragePolicyVmHomePolicySpecPolicyTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.storage.policy.vm_home_policy_spec.policy_type value
func (v VcenterVmStoragePolicyVmHomePolicySpecPolicyType) Ptr() *VcenterVmStoragePolicyVmHomePolicySpecPolicyType {
	return &v
}

type NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType struct {
	value *VcenterVmStoragePolicyVmHomePolicySpecPolicyType
	isSet bool
}

func (v NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType) Get() *VcenterVmStoragePolicyVmHomePolicySpecPolicyType {
	return v.value
}

func (v *NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType) Set(val *VcenterVmStoragePolicyVmHomePolicySpecPolicyType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType(val *VcenterVmStoragePolicyVmHomePolicySpecPolicyType) *NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType {
	return &NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType{value: val, isSet: true}
}

func (v NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmStoragePolicyVmHomePolicySpecPolicyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
