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

// VcenterGuestIpv6Type The Ipv6.Type enumerated type specifies different types of the IPv6 configuration.
type VcenterGuestIpv6Type string

// List of vcenter.guest.ipv6.type
const (
	VCENTERGUESTIPV6TYPE_DHCP VcenterGuestIpv6Type = "DHCP"
	VCENTERGUESTIPV6TYPE_STATIC VcenterGuestIpv6Type = "STATIC"
	VCENTERGUESTIPV6TYPE_USER_INPUT_REQUIRED VcenterGuestIpv6Type = "USER_INPUT_REQUIRED"
)

// All allowed values of VcenterGuestIpv6Type enum
var AllowedVcenterGuestIpv6TypeEnumValues = []VcenterGuestIpv6Type{
	"DHCP",
	"STATIC",
	"USER_INPUT_REQUIRED",
}

func (v *VcenterGuestIpv6Type) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterGuestIpv6Type(value)
	for _, existing := range AllowedVcenterGuestIpv6TypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterGuestIpv6Type", value)
}

// NewVcenterGuestIpv6TypeFromValue returns a pointer to a valid VcenterGuestIpv6Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterGuestIpv6TypeFromValue(v string) (*VcenterGuestIpv6Type, error) {
	ev := VcenterGuestIpv6Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterGuestIpv6Type: valid values are %v", v, AllowedVcenterGuestIpv6TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterGuestIpv6Type) IsValid() bool {
	for _, existing := range AllowedVcenterGuestIpv6TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.guest.ipv6.type value
func (v VcenterGuestIpv6Type) Ptr() *VcenterGuestIpv6Type {
	return &v
}

type NullableVcenterGuestIpv6Type struct {
	value *VcenterGuestIpv6Type
	isSet bool
}

func (v NullableVcenterGuestIpv6Type) Get() *VcenterGuestIpv6Type {
	return v.value
}

func (v *NullableVcenterGuestIpv6Type) Set(val *VcenterGuestIpv6Type) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestIpv6Type) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestIpv6Type) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestIpv6Type(val *VcenterGuestIpv6Type) *NullableVcenterGuestIpv6Type {
	return &NullableVcenterGuestIpv6Type{value: val, isSet: true}
}

func (v NullableVcenterGuestIpv6Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestIpv6Type) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

