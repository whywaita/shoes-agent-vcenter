/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VcenterNamespacesAccessInfo struct for VcenterNamespacesAccessInfo
type VcenterNamespacesAccessInfo struct {
	Role VcenterNamespacesAccessRole `json:"role"`
	// Flag to indicate if the Access.Info.role is direct or inherited. The value is set to true if the Access.Info.role is inherited from group membership This field is optional because it was added in a newer version than its parent node.
	Inherited *bool `json:"inherited,omitempty"`
}

// NewVcenterNamespacesAccessInfo instantiates a new VcenterNamespacesAccessInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespacesAccessInfo(role VcenterNamespacesAccessRole) *VcenterNamespacesAccessInfo {
	this := VcenterNamespacesAccessInfo{}
	this.Role = role
	return &this
}

// NewVcenterNamespacesAccessInfoWithDefaults instantiates a new VcenterNamespacesAccessInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespacesAccessInfoWithDefaults() *VcenterNamespacesAccessInfo {
	this := VcenterNamespacesAccessInfo{}
	return &this
}

// GetRole returns the Role field value
func (o *VcenterNamespacesAccessInfo) GetRole() VcenterNamespacesAccessRole {
	if o == nil {
		var ret VcenterNamespacesAccessRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesAccessInfo) GetRoleOk() (*VcenterNamespacesAccessRole, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *VcenterNamespacesAccessInfo) SetRole(v VcenterNamespacesAccessRole) {
	o.Role = v
}

// GetInherited returns the Inherited field value if set, zero value otherwise.
func (o *VcenterNamespacesAccessInfo) GetInherited() bool {
	if o == nil || o.Inherited == nil {
		var ret bool
		return ret
	}
	return *o.Inherited
}

// GetInheritedOk returns a tuple with the Inherited field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesAccessInfo) GetInheritedOk() (*bool, bool) {
	if o == nil || o.Inherited == nil {
		return nil, false
	}
	return o.Inherited, true
}

// HasInherited returns a boolean if a field has been set.
func (o *VcenterNamespacesAccessInfo) HasInherited() bool {
	if o != nil && o.Inherited != nil {
		return true
	}

	return false
}

// SetInherited gets a reference to the given bool and assigns it to the Inherited field.
func (o *VcenterNamespacesAccessInfo) SetInherited(v bool) {
	o.Inherited = &v
}

func (o VcenterNamespacesAccessInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["role"] = o.Role
	}
	if o.Inherited != nil {
		toSerialize["inherited"] = o.Inherited
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespacesAccessInfo struct {
	value *VcenterNamespacesAccessInfo
	isSet bool
}

func (v NullableVcenterNamespacesAccessInfo) Get() *VcenterNamespacesAccessInfo {
	return v.value
}

func (v *NullableVcenterNamespacesAccessInfo) Set(val *VcenterNamespacesAccessInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesAccessInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesAccessInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesAccessInfo(val *VcenterNamespacesAccessInfo) *NullableVcenterNamespacesAccessInfo {
	return &NullableVcenterNamespacesAccessInfo{value: val, isSet: true}
}

func (v NullableVcenterNamespacesAccessInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesAccessInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

