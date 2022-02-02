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

// VcenterVchaClusterIpFamily The Cluster.IpFamily enumerated type defines the IP address family.
type VcenterVchaClusterIpFamily string

// List of vcenter.vcha.cluster.ip_family
const (
	VCENTERVCHACLUSTERIPFAMILY_IPV4 VcenterVchaClusterIpFamily = "IPV4"
	VCENTERVCHACLUSTERIPFAMILY_IPV6 VcenterVchaClusterIpFamily = "IPV6"
)

// All allowed values of VcenterVchaClusterIpFamily enum
var AllowedVcenterVchaClusterIpFamilyEnumValues = []VcenterVchaClusterIpFamily{
	"IPV4",
	"IPV6",
}

func (v *VcenterVchaClusterIpFamily) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVchaClusterIpFamily(value)
	for _, existing := range AllowedVcenterVchaClusterIpFamilyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVchaClusterIpFamily", value)
}

// NewVcenterVchaClusterIpFamilyFromValue returns a pointer to a valid VcenterVchaClusterIpFamily
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVchaClusterIpFamilyFromValue(v string) (*VcenterVchaClusterIpFamily, error) {
	ev := VcenterVchaClusterIpFamily(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVchaClusterIpFamily: valid values are %v", v, AllowedVcenterVchaClusterIpFamilyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVchaClusterIpFamily) IsValid() bool {
	for _, existing := range AllowedVcenterVchaClusterIpFamilyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vcha.cluster.ip_family value
func (v VcenterVchaClusterIpFamily) Ptr() *VcenterVchaClusterIpFamily {
	return &v
}

type NullableVcenterVchaClusterIpFamily struct {
	value *VcenterVchaClusterIpFamily
	isSet bool
}

func (v NullableVcenterVchaClusterIpFamily) Get() *VcenterVchaClusterIpFamily {
	return v.value
}

func (v *NullableVcenterVchaClusterIpFamily) Set(val *VcenterVchaClusterIpFamily) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterIpFamily) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterIpFamily) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterIpFamily(val *VcenterVchaClusterIpFamily) *NullableVcenterVchaClusterIpFamily {
	return &NullableVcenterVchaClusterIpFamily{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterIpFamily) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterIpFamily) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

