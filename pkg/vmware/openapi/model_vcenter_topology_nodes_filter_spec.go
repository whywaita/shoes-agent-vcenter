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

// VcenterTopologyNodesFilterSpec struct for VcenterTopologyNodesFilterSpec
type VcenterTopologyNodesFilterSpec struct {
	// Types of the appliance that a vCenter and Platform Services Controller node must be to match the filter (see Nodes.ApplianceType. If unset or empty, node of any ApplianceType match the filter.
	Types *[]VcenterTopologyNodesApplianceType `json:"types,omitempty"`
}

// NewVcenterTopologyNodesFilterSpec instantiates a new VcenterTopologyNodesFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTopologyNodesFilterSpec() *VcenterTopologyNodesFilterSpec {
	this := VcenterTopologyNodesFilterSpec{}
	return &this
}

// NewVcenterTopologyNodesFilterSpecWithDefaults instantiates a new VcenterTopologyNodesFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTopologyNodesFilterSpecWithDefaults() *VcenterTopologyNodesFilterSpec {
	this := VcenterTopologyNodesFilterSpec{}
	return &this
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *VcenterTopologyNodesFilterSpec) GetTypes() []VcenterTopologyNodesApplianceType {
	if o == nil || o.Types == nil {
		var ret []VcenterTopologyNodesApplianceType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTopologyNodesFilterSpec) GetTypesOk() (*[]VcenterTopologyNodesApplianceType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *VcenterTopologyNodesFilterSpec) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []VcenterTopologyNodesApplianceType and assigns it to the Types field.
func (o *VcenterTopologyNodesFilterSpec) SetTypes(v []VcenterTopologyNodesApplianceType) {
	o.Types = &v
}

func (o VcenterTopologyNodesFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTopologyNodesFilterSpec struct {
	value *VcenterTopologyNodesFilterSpec
	isSet bool
}

func (v NullableVcenterTopologyNodesFilterSpec) Get() *VcenterTopologyNodesFilterSpec {
	return v.value
}

func (v *NullableVcenterTopologyNodesFilterSpec) Set(val *VcenterTopologyNodesFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTopologyNodesFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTopologyNodesFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTopologyNodesFilterSpec(val *VcenterTopologyNodesFilterSpec) *NullableVcenterTopologyNodesFilterSpec {
	return &NullableVcenterTopologyNodesFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterTopologyNodesFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTopologyNodesFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


