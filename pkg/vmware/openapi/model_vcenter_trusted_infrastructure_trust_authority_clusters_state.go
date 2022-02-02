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

// VcenterTrustedInfrastructureTrustAuthorityClustersState the model 'VcenterTrustedInfrastructureTrustAuthorityClustersState'
type VcenterTrustedInfrastructureTrustAuthorityClustersState string

// List of VcenterTrustedInfrastructureTrustAuthorityClustersState
const (
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSSTATE_ENABLE VcenterTrustedInfrastructureTrustAuthorityClustersState = "ENABLE"
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSSTATE_DISABLE VcenterTrustedInfrastructureTrustAuthorityClustersState = "DISABLE"
)

// All allowed values of VcenterTrustedInfrastructureTrustAuthorityClustersState enum
var AllowedVcenterTrustedInfrastructureTrustAuthorityClustersStateEnumValues = []VcenterTrustedInfrastructureTrustAuthorityClustersState{
	"ENABLE",
	"DISABLE",
}

func (v *VcenterTrustedInfrastructureTrustAuthorityClustersState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterTrustedInfrastructureTrustAuthorityClustersState(value)
	for _, existing := range AllowedVcenterTrustedInfrastructureTrustAuthorityClustersStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterTrustedInfrastructureTrustAuthorityClustersState", value)
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersStateFromValue returns a pointer to a valid VcenterTrustedInfrastructureTrustAuthorityClustersState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterTrustedInfrastructureTrustAuthorityClustersStateFromValue(v string) (*VcenterTrustedInfrastructureTrustAuthorityClustersState, error) {
	ev := VcenterTrustedInfrastructureTrustAuthorityClustersState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterTrustedInfrastructureTrustAuthorityClustersState: valid values are %v", v, AllowedVcenterTrustedInfrastructureTrustAuthorityClustersStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterTrustedInfrastructureTrustAuthorityClustersState) IsValid() bool {
	for _, existing := range AllowedVcenterTrustedInfrastructureTrustAuthorityClustersStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterTrustedInfrastructureTrustAuthorityClustersState value
func (v VcenterTrustedInfrastructureTrustAuthorityClustersState) Ptr() *VcenterTrustedInfrastructureTrustAuthorityClustersState {
	return &v
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersState struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersState
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersState) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersState {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersState) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersState) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersState) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersState(val *VcenterTrustedInfrastructureTrustAuthorityClustersState) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersState {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersState{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

