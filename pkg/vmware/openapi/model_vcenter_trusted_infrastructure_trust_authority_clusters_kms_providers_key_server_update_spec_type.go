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

// VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType The Providers.KeyServerUpdateSpec.Type enumerated type list the key server types.
type VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType string

// List of VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType
const (
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSKMSPROVIDERSKEYSERVERUPDATESPECTYPE_KMIP VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType = "KMIP"
)

// All allowed values of VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType enum
var AllowedVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecTypeEnumValues = []VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType{
	"KMIP",
}

func (v *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType(value)
	for _, existing := range AllowedVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType", value)
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecTypeFromValue returns a pointer to a valid VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecTypeFromValue(v string) (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType, error) {
	ev := VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType: valid values are %v", v, AllowedVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) IsValid() bool {
	for _, existing := range AllowedVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType value
func (v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) Ptr() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType {
	return &v
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType(val *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

