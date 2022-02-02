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

// VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth the model 'VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth'
type VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth string

// List of VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth
const (
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSATTESTATIONOSESXBASEIMAGESHEALTH_NONE VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth = "NONE"
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSATTESTATIONOSESXBASEIMAGESHEALTH_OK VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth = "OK"
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSATTESTATIONOSESXBASEIMAGESHEALTH_WARNING VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth = "WARNING"
	VCENTERTRUSTEDINFRASTRUCTURETRUSTAUTHORITYCLUSTERSATTESTATIONOSESXBASEIMAGESHEALTH_ERROR VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth = "ERROR"
)

// All allowed values of VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth enum
var AllowedVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealthEnumValues = []VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth{
	"NONE",
	"OK",
	"WARNING",
	"ERROR",
}

func (v *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth(value)
	for _, existing := range AllowedVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealthEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth", value)
}

// NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealthFromValue returns a pointer to a valid VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealthFromValue(v string) (*VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth, error) {
	ev := VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth: valid values are %v", v, AllowedVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealthEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) IsValid() bool {
	for _, existing := range AllowedVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealthEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth value
func (v VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) Ptr() *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth {
	return &v
}

type NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth struct {
	value *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) Get() *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) Set(val *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth(val *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) *NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
