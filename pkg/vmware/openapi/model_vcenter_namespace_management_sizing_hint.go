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

// VcenterNamespaceManagementSizingHint The SizingHint enumerated type determines the configuration of Kubernetes API server and the worker nodes. It also determines the default values associated with the maximum number of pods and services. Use ClusterSizeInfo.get to get information associated with a SizingHint.
type VcenterNamespaceManagementSizingHint string

// List of VcenterNamespaceManagementSizingHint
const (
	VCENTERNAMESPACEMANAGEMENTSIZINGHINT_TINY VcenterNamespaceManagementSizingHint = "TINY"
	VCENTERNAMESPACEMANAGEMENTSIZINGHINT_SMALL VcenterNamespaceManagementSizingHint = "SMALL"
	VCENTERNAMESPACEMANAGEMENTSIZINGHINT_MEDIUM VcenterNamespaceManagementSizingHint = "MEDIUM"
	VCENTERNAMESPACEMANAGEMENTSIZINGHINT_LARGE VcenterNamespaceManagementSizingHint = "LARGE"
)

// All allowed values of VcenterNamespaceManagementSizingHint enum
var AllowedVcenterNamespaceManagementSizingHintEnumValues = []VcenterNamespaceManagementSizingHint{
	"TINY",
	"SMALL",
	"MEDIUM",
	"LARGE",
}

func (v *VcenterNamespaceManagementSizingHint) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterNamespaceManagementSizingHint(value)
	for _, existing := range AllowedVcenterNamespaceManagementSizingHintEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterNamespaceManagementSizingHint", value)
}

// NewVcenterNamespaceManagementSizingHintFromValue returns a pointer to a valid VcenterNamespaceManagementSizingHint
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterNamespaceManagementSizingHintFromValue(v string) (*VcenterNamespaceManagementSizingHint, error) {
	ev := VcenterNamespaceManagementSizingHint(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterNamespaceManagementSizingHint: valid values are %v", v, AllowedVcenterNamespaceManagementSizingHintEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterNamespaceManagementSizingHint) IsValid() bool {
	for _, existing := range AllowedVcenterNamespaceManagementSizingHintEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterNamespaceManagementSizingHint value
func (v VcenterNamespaceManagementSizingHint) Ptr() *VcenterNamespaceManagementSizingHint {
	return &v
}

type NullableVcenterNamespaceManagementSizingHint struct {
	value *VcenterNamespaceManagementSizingHint
	isSet bool
}

func (v NullableVcenterNamespaceManagementSizingHint) Get() *VcenterNamespaceManagementSizingHint {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSizingHint) Set(val *VcenterNamespaceManagementSizingHint) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSizingHint) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSizingHint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSizingHint(val *VcenterNamespaceManagementSizingHint) *NullableVcenterNamespaceManagementSizingHint {
	return &NullableVcenterNamespaceManagementSizingHint{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSizingHint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSizingHint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
