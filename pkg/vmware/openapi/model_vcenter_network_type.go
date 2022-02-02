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

// VcenterNetworkType the model 'VcenterNetworkType'
type VcenterNetworkType string

// List of vcenter.network.type
const (
	VCENTERNETWORKTYPE_STANDARD_PORTGROUP VcenterNetworkType = "STANDARD_PORTGROUP"
	VCENTERNETWORKTYPE_DISTRIBUTED_PORTGROUP VcenterNetworkType = "DISTRIBUTED_PORTGROUP"
	VCENTERNETWORKTYPE_OPAQUE_NETWORK VcenterNetworkType = "OPAQUE_NETWORK"
)

// All allowed values of VcenterNetworkType enum
var AllowedVcenterNetworkTypeEnumValues = []VcenterNetworkType{
	"STANDARD_PORTGROUP",
	"DISTRIBUTED_PORTGROUP",
	"OPAQUE_NETWORK",
}

func (v *VcenterNetworkType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterNetworkType(value)
	for _, existing := range AllowedVcenterNetworkTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterNetworkType", value)
}

// NewVcenterNetworkTypeFromValue returns a pointer to a valid VcenterNetworkType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterNetworkTypeFromValue(v string) (*VcenterNetworkType, error) {
	ev := VcenterNetworkType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterNetworkType: valid values are %v", v, AllowedVcenterNetworkTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterNetworkType) IsValid() bool {
	for _, existing := range AllowedVcenterNetworkTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.network.type value
func (v VcenterNetworkType) Ptr() *VcenterNetworkType {
	return &v
}

type NullableVcenterNetworkType struct {
	value *VcenterNetworkType
	isSet bool
}

func (v NullableVcenterNetworkType) Get() *VcenterNetworkType {
	return v.value
}

func (v *NullableVcenterNetworkType) Set(val *VcenterNetworkType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNetworkType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNetworkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNetworkType(val *VcenterNetworkType) *NullableVcenterNetworkType {
	return &NullableVcenterNetworkType{value: val, isSet: true}
}

func (v NullableVcenterNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNetworkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

