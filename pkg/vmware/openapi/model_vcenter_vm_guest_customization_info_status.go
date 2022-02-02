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

// VcenterVmGuestCustomizationInfoStatus The Customization.Info.Status enumerated type defines the status values that can be reported for the customization operation.
type VcenterVmGuestCustomizationInfoStatus string

// List of vcenter.vm.guest.customization.info.status
const (
	VCENTERVMGUESTCUSTOMIZATIONINFOSTATUS_IDLE VcenterVmGuestCustomizationInfoStatus = "IDLE"
	VCENTERVMGUESTCUSTOMIZATIONINFOSTATUS_PENDING VcenterVmGuestCustomizationInfoStatus = "PENDING"
	VCENTERVMGUESTCUSTOMIZATIONINFOSTATUS_RUNNING VcenterVmGuestCustomizationInfoStatus = "RUNNING"
	VCENTERVMGUESTCUSTOMIZATIONINFOSTATUS_SUCCEEDED VcenterVmGuestCustomizationInfoStatus = "SUCCEEDED"
	VCENTERVMGUESTCUSTOMIZATIONINFOSTATUS_FAILED VcenterVmGuestCustomizationInfoStatus = "FAILED"
)

// All allowed values of VcenterVmGuestCustomizationInfoStatus enum
var AllowedVcenterVmGuestCustomizationInfoStatusEnumValues = []VcenterVmGuestCustomizationInfoStatus{
	"IDLE",
	"PENDING",
	"RUNNING",
	"SUCCEEDED",
	"FAILED",
}

func (v *VcenterVmGuestCustomizationInfoStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVmGuestCustomizationInfoStatus(value)
	for _, existing := range AllowedVcenterVmGuestCustomizationInfoStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVmGuestCustomizationInfoStatus", value)
}

// NewVcenterVmGuestCustomizationInfoStatusFromValue returns a pointer to a valid VcenterVmGuestCustomizationInfoStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVmGuestCustomizationInfoStatusFromValue(v string) (*VcenterVmGuestCustomizationInfoStatus, error) {
	ev := VcenterVmGuestCustomizationInfoStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVmGuestCustomizationInfoStatus: valid values are %v", v, AllowedVcenterVmGuestCustomizationInfoStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVmGuestCustomizationInfoStatus) IsValid() bool {
	for _, existing := range AllowedVcenterVmGuestCustomizationInfoStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vm.guest.customization.info.status value
func (v VcenterVmGuestCustomizationInfoStatus) Ptr() *VcenterVmGuestCustomizationInfoStatus {
	return &v
}

type NullableVcenterVmGuestCustomizationInfoStatus struct {
	value *VcenterVmGuestCustomizationInfoStatus
	isSet bool
}

func (v NullableVcenterVmGuestCustomizationInfoStatus) Get() *VcenterVmGuestCustomizationInfoStatus {
	return v.value
}

func (v *NullableVcenterVmGuestCustomizationInfoStatus) Set(val *VcenterVmGuestCustomizationInfoStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestCustomizationInfoStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestCustomizationInfoStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestCustomizationInfoStatus(val *VcenterVmGuestCustomizationInfoStatus) *NullableVcenterVmGuestCustomizationInfoStatus {
	return &NullableVcenterVmGuestCustomizationInfoStatus{value: val, isSet: true}
}

func (v NullableVcenterVmGuestCustomizationInfoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestCustomizationInfoStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
