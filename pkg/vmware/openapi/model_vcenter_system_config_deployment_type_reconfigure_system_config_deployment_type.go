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

// VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType struct for VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType
type VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType struct {
	Spec *VcenterSystemConfigDeploymentTypeReconfigureSpec `json:"spec,omitempty"`
}

// NewVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType instantiates a new VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType() *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType {
	this := VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType{}
	return &this
}

// NewVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentTypeWithDefaults instantiates a new VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentTypeWithDefaults() *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType {
	this := VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) GetSpec() VcenterSystemConfigDeploymentTypeReconfigureSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterSystemConfigDeploymentTypeReconfigureSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) GetSpecOk() (*VcenterSystemConfigDeploymentTypeReconfigureSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterSystemConfigDeploymentTypeReconfigureSpec and assigns it to the Spec field.
func (o *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) SetSpec(v VcenterSystemConfigDeploymentTypeReconfigureSpec) {
	o.Spec = &v
}

func (o VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType struct {
	value *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType
	isSet bool
}

func (v NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) Get() *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType {
	return v.value
}

func (v *NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) Set(val *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType(val *VcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) *NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType {
	return &NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType{value: val, isSet: true}
}

func (v NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterSystemConfigDeploymentTypeReconfigureSystemConfigDeploymentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

