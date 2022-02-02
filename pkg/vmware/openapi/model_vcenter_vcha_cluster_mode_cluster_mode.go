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

// VcenterVchaClusterModeClusterMode The Mode.ClusterMode enumerated type defines the possible modes for a VCHA Cluster.
type VcenterVchaClusterModeClusterMode string

// List of vcenter.vcha.cluster.mode.cluster_mode
const (
	VCENTERVCHACLUSTERMODECLUSTERMODE_ENABLED VcenterVchaClusterModeClusterMode = "ENABLED"
	VCENTERVCHACLUSTERMODECLUSTERMODE_DISABLED VcenterVchaClusterModeClusterMode = "DISABLED"
	VCENTERVCHACLUSTERMODECLUSTERMODE_MAINTENANCE VcenterVchaClusterModeClusterMode = "MAINTENANCE"
)

// All allowed values of VcenterVchaClusterModeClusterMode enum
var AllowedVcenterVchaClusterModeClusterModeEnumValues = []VcenterVchaClusterModeClusterMode{
	"ENABLED",
	"DISABLED",
	"MAINTENANCE",
}

func (v *VcenterVchaClusterModeClusterMode) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVchaClusterModeClusterMode(value)
	for _, existing := range AllowedVcenterVchaClusterModeClusterModeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVchaClusterModeClusterMode", value)
}

// NewVcenterVchaClusterModeClusterModeFromValue returns a pointer to a valid VcenterVchaClusterModeClusterMode
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVchaClusterModeClusterModeFromValue(v string) (*VcenterVchaClusterModeClusterMode, error) {
	ev := VcenterVchaClusterModeClusterMode(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVchaClusterModeClusterMode: valid values are %v", v, AllowedVcenterVchaClusterModeClusterModeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVchaClusterModeClusterMode) IsValid() bool {
	for _, existing := range AllowedVcenterVchaClusterModeClusterModeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vcha.cluster.mode.cluster_mode value
func (v VcenterVchaClusterModeClusterMode) Ptr() *VcenterVchaClusterModeClusterMode {
	return &v
}

type NullableVcenterVchaClusterModeClusterMode struct {
	value *VcenterVchaClusterModeClusterMode
	isSet bool
}

func (v NullableVcenterVchaClusterModeClusterMode) Get() *VcenterVchaClusterModeClusterMode {
	return v.value
}

func (v *NullableVcenterVchaClusterModeClusterMode) Set(val *VcenterVchaClusterModeClusterMode) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterModeClusterMode) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterModeClusterMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterModeClusterMode(val *VcenterVchaClusterModeClusterMode) *NullableVcenterVchaClusterModeClusterMode {
	return &NullableVcenterVchaClusterModeClusterMode{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterModeClusterMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterModeClusterMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
