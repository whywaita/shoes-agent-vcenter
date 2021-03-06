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

// VcenterLcmUpdatePendingUpdateType The Pending.UpdateType enumerated type defines update type
type VcenterLcmUpdatePendingUpdateType string

// List of VcenterLcmUpdatePendingUpdateType
const (
	VCENTERLCMUPDATEPENDINGUPDATETYPE_PATCH VcenterLcmUpdatePendingUpdateType = "PATCH"
	VCENTERLCMUPDATEPENDINGUPDATETYPE_UPDATE VcenterLcmUpdatePendingUpdateType = "UPDATE"
	VCENTERLCMUPDATEPENDINGUPDATETYPE_UPGRADE VcenterLcmUpdatePendingUpdateType = "UPGRADE"
)

// All allowed values of VcenterLcmUpdatePendingUpdateType enum
var AllowedVcenterLcmUpdatePendingUpdateTypeEnumValues = []VcenterLcmUpdatePendingUpdateType{
	"PATCH",
	"UPDATE",
	"UPGRADE",
}

func (v *VcenterLcmUpdatePendingUpdateType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterLcmUpdatePendingUpdateType(value)
	for _, existing := range AllowedVcenterLcmUpdatePendingUpdateTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterLcmUpdatePendingUpdateType", value)
}

// NewVcenterLcmUpdatePendingUpdateTypeFromValue returns a pointer to a valid VcenterLcmUpdatePendingUpdateType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterLcmUpdatePendingUpdateTypeFromValue(v string) (*VcenterLcmUpdatePendingUpdateType, error) {
	ev := VcenterLcmUpdatePendingUpdateType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterLcmUpdatePendingUpdateType: valid values are %v", v, AllowedVcenterLcmUpdatePendingUpdateTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterLcmUpdatePendingUpdateType) IsValid() bool {
	for _, existing := range AllowedVcenterLcmUpdatePendingUpdateTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterLcmUpdatePendingUpdateType value
func (v VcenterLcmUpdatePendingUpdateType) Ptr() *VcenterLcmUpdatePendingUpdateType {
	return &v
}

type NullableVcenterLcmUpdatePendingUpdateType struct {
	value *VcenterLcmUpdatePendingUpdateType
	isSet bool
}

func (v NullableVcenterLcmUpdatePendingUpdateType) Get() *VcenterLcmUpdatePendingUpdateType {
	return v.value
}

func (v *NullableVcenterLcmUpdatePendingUpdateType) Set(val *VcenterLcmUpdatePendingUpdateType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterLcmUpdatePendingUpdateType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterLcmUpdatePendingUpdateType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterLcmUpdatePendingUpdateType(val *VcenterLcmUpdatePendingUpdateType) *NullableVcenterLcmUpdatePendingUpdateType {
	return &NullableVcenterLcmUpdatePendingUpdateType{value: val, isSet: true}
}

func (v NullableVcenterLcmUpdatePendingUpdateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterLcmUpdatePendingUpdateType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

