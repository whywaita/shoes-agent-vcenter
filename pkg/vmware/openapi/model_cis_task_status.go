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

// CisTaskStatus The Status enumerated type defines the status values that can be reported for an operation.
type CisTaskStatus string

// List of cis.task.status
const (
	CISTASKSTATUS_PENDING CisTaskStatus = "PENDING"
	CISTASKSTATUS_RUNNING CisTaskStatus = "RUNNING"
	CISTASKSTATUS_BLOCKED CisTaskStatus = "BLOCKED"
	CISTASKSTATUS_SUCCEEDED CisTaskStatus = "SUCCEEDED"
	CISTASKSTATUS_FAILED CisTaskStatus = "FAILED"
)

// All allowed values of CisTaskStatus enum
var AllowedCisTaskStatusEnumValues = []CisTaskStatus{
	"PENDING",
	"RUNNING",
	"BLOCKED",
	"SUCCEEDED",
	"FAILED",
}

func (v *CisTaskStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CisTaskStatus(value)
	for _, existing := range AllowedCisTaskStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CisTaskStatus", value)
}

// NewCisTaskStatusFromValue returns a pointer to a valid CisTaskStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCisTaskStatusFromValue(v string) (*CisTaskStatus, error) {
	ev := CisTaskStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CisTaskStatus: valid values are %v", v, AllowedCisTaskStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CisTaskStatus) IsValid() bool {
	for _, existing := range AllowedCisTaskStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to cis.task.status value
func (v CisTaskStatus) Ptr() *CisTaskStatus {
	return &v
}

type NullableCisTaskStatus struct {
	value *CisTaskStatus
	isSet bool
}

func (v NullableCisTaskStatus) Get() *CisTaskStatus {
	return v.value
}

func (v *NullableCisTaskStatus) Set(val *CisTaskStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCisTaskStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCisTaskStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCisTaskStatus(val *CisTaskStatus) *NullableCisTaskStatus {
	return &NullableCisTaskStatus{value: val, isSet: true}
}

func (v NullableCisTaskStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCisTaskStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
