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

// VcenterNamespaceManagementSoftwareClustersMessageSeverity The Clusters.Message.Severity enumerated type represents the severity of the message.
type VcenterNamespaceManagementSoftwareClustersMessageSeverity string

// List of VcenterNamespaceManagementSoftwareClustersMessageSeverity
const (
	VCENTERNAMESPACEMANAGEMENTSOFTWARECLUSTERSMESSAGESEVERITY_INFO VcenterNamespaceManagementSoftwareClustersMessageSeverity = "INFO"
	VCENTERNAMESPACEMANAGEMENTSOFTWARECLUSTERSMESSAGESEVERITY_WARNING VcenterNamespaceManagementSoftwareClustersMessageSeverity = "WARNING"
	VCENTERNAMESPACEMANAGEMENTSOFTWARECLUSTERSMESSAGESEVERITY_ERROR VcenterNamespaceManagementSoftwareClustersMessageSeverity = "ERROR"
)

// All allowed values of VcenterNamespaceManagementSoftwareClustersMessageSeverity enum
var AllowedVcenterNamespaceManagementSoftwareClustersMessageSeverityEnumValues = []VcenterNamespaceManagementSoftwareClustersMessageSeverity{
	"INFO",
	"WARNING",
	"ERROR",
}

func (v *VcenterNamespaceManagementSoftwareClustersMessageSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterNamespaceManagementSoftwareClustersMessageSeverity(value)
	for _, existing := range AllowedVcenterNamespaceManagementSoftwareClustersMessageSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterNamespaceManagementSoftwareClustersMessageSeverity", value)
}

// NewVcenterNamespaceManagementSoftwareClustersMessageSeverityFromValue returns a pointer to a valid VcenterNamespaceManagementSoftwareClustersMessageSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterNamespaceManagementSoftwareClustersMessageSeverityFromValue(v string) (*VcenterNamespaceManagementSoftwareClustersMessageSeverity, error) {
	ev := VcenterNamespaceManagementSoftwareClustersMessageSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterNamespaceManagementSoftwareClustersMessageSeverity: valid values are %v", v, AllowedVcenterNamespaceManagementSoftwareClustersMessageSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterNamespaceManagementSoftwareClustersMessageSeverity) IsValid() bool {
	for _, existing := range AllowedVcenterNamespaceManagementSoftwareClustersMessageSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterNamespaceManagementSoftwareClustersMessageSeverity value
func (v VcenterNamespaceManagementSoftwareClustersMessageSeverity) Ptr() *VcenterNamespaceManagementSoftwareClustersMessageSeverity {
	return &v
}

type NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity struct {
	value *VcenterNamespaceManagementSoftwareClustersMessageSeverity
	isSet bool
}

func (v NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity) Get() *VcenterNamespaceManagementSoftwareClustersMessageSeverity {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity) Set(val *VcenterNamespaceManagementSoftwareClustersMessageSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSoftwareClustersMessageSeverity(val *VcenterNamespaceManagementSoftwareClustersMessageSeverity) *NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity {
	return &NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSoftwareClustersMessageSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

