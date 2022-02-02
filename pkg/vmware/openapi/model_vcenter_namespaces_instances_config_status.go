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

// VcenterNamespacesInstancesConfigStatus The Instances.ConfigStatus enumerated type describes the status of reaching the desired state configuration for the namespace.
type VcenterNamespacesInstancesConfigStatus string

// List of VcenterNamespacesInstancesConfigStatus
const (
	VCENTERNAMESPACESINSTANCESCONFIGSTATUS_CONFIGURING VcenterNamespacesInstancesConfigStatus = "CONFIGURING"
	VCENTERNAMESPACESINSTANCESCONFIGSTATUS_REMOVING VcenterNamespacesInstancesConfigStatus = "REMOVING"
	VCENTERNAMESPACESINSTANCESCONFIGSTATUS_RUNNING VcenterNamespacesInstancesConfigStatus = "RUNNING"
	VCENTERNAMESPACESINSTANCESCONFIGSTATUS_ERROR VcenterNamespacesInstancesConfigStatus = "ERROR"
)

// All allowed values of VcenterNamespacesInstancesConfigStatus enum
var AllowedVcenterNamespacesInstancesConfigStatusEnumValues = []VcenterNamespacesInstancesConfigStatus{
	"CONFIGURING",
	"REMOVING",
	"RUNNING",
	"ERROR",
}

func (v *VcenterNamespacesInstancesConfigStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterNamespacesInstancesConfigStatus(value)
	for _, existing := range AllowedVcenterNamespacesInstancesConfigStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterNamespacesInstancesConfigStatus", value)
}

// NewVcenterNamespacesInstancesConfigStatusFromValue returns a pointer to a valid VcenterNamespacesInstancesConfigStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterNamespacesInstancesConfigStatusFromValue(v string) (*VcenterNamespacesInstancesConfigStatus, error) {
	ev := VcenterNamespacesInstancesConfigStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterNamespacesInstancesConfigStatus: valid values are %v", v, AllowedVcenterNamespacesInstancesConfigStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterNamespacesInstancesConfigStatus) IsValid() bool {
	for _, existing := range AllowedVcenterNamespacesInstancesConfigStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterNamespacesInstancesConfigStatus value
func (v VcenterNamespacesInstancesConfigStatus) Ptr() *VcenterNamespacesInstancesConfigStatus {
	return &v
}

type NullableVcenterNamespacesInstancesConfigStatus struct {
	value *VcenterNamespacesInstancesConfigStatus
	isSet bool
}

func (v NullableVcenterNamespacesInstancesConfigStatus) Get() *VcenterNamespacesInstancesConfigStatus {
	return v.value
}

func (v *NullableVcenterNamespacesInstancesConfigStatus) Set(val *VcenterNamespacesInstancesConfigStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesInstancesConfigStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesInstancesConfigStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesInstancesConfigStatus(val *VcenterNamespacesInstancesConfigStatus) *NullableVcenterNamespacesInstancesConfigStatus {
	return &NullableVcenterNamespacesInstancesConfigStatus{value: val, isSet: true}
}

func (v NullableVcenterNamespacesInstancesConfigStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesInstancesConfigStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

