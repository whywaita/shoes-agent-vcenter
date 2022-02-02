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

// VcenterVchaClusterNodeRole The Cluster.NodeRole enumerated type defines the role node can be in a VCHA Cluster.
type VcenterVchaClusterNodeRole string

// List of vcenter.vcha.cluster.node_role
const (
	VCENTERVCHACLUSTERNODEROLE_ACTIVE VcenterVchaClusterNodeRole = "ACTIVE"
	VCENTERVCHACLUSTERNODEROLE_PASSIVE VcenterVchaClusterNodeRole = "PASSIVE"
	VCENTERVCHACLUSTERNODEROLE_WITNESS VcenterVchaClusterNodeRole = "WITNESS"
)

// All allowed values of VcenterVchaClusterNodeRole enum
var AllowedVcenterVchaClusterNodeRoleEnumValues = []VcenterVchaClusterNodeRole{
	"ACTIVE",
	"PASSIVE",
	"WITNESS",
}

func (v *VcenterVchaClusterNodeRole) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VcenterVchaClusterNodeRole(value)
	for _, existing := range AllowedVcenterVchaClusterNodeRoleEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VcenterVchaClusterNodeRole", value)
}

// NewVcenterVchaClusterNodeRoleFromValue returns a pointer to a valid VcenterVchaClusterNodeRole
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVcenterVchaClusterNodeRoleFromValue(v string) (*VcenterVchaClusterNodeRole, error) {
	ev := VcenterVchaClusterNodeRole(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VcenterVchaClusterNodeRole: valid values are %v", v, AllowedVcenterVchaClusterNodeRoleEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VcenterVchaClusterNodeRole) IsValid() bool {
	for _, existing := range AllowedVcenterVchaClusterNodeRoleEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to vcenter.vcha.cluster.node_role value
func (v VcenterVchaClusterNodeRole) Ptr() *VcenterVchaClusterNodeRole {
	return &v
}

type NullableVcenterVchaClusterNodeRole struct {
	value *VcenterVchaClusterNodeRole
	isSet bool
}

func (v NullableVcenterVchaClusterNodeRole) Get() *VcenterVchaClusterNodeRole {
	return v.value
}

func (v *NullableVcenterVchaClusterNodeRole) Set(val *VcenterVchaClusterNodeRole) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterNodeRole) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterNodeRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterNodeRole(val *VcenterVchaClusterNodeRole) *NullableVcenterVchaClusterNodeRole {
	return &NullableVcenterVchaClusterNodeRole{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterNodeRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterNodeRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

