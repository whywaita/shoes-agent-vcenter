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

// VcenterVMCloneTaskVm struct for VcenterVMCloneTaskVm
type VcenterVMCloneTaskVm struct {
	Spec *VcenterVMCloneSpec `json:"spec,omitempty"`
}

// NewVcenterVMCloneTaskVm instantiates a new VcenterVMCloneTaskVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMCloneTaskVm() *VcenterVMCloneTaskVm {
	this := VcenterVMCloneTaskVm{}
	return &this
}

// NewVcenterVMCloneTaskVmWithDefaults instantiates a new VcenterVMCloneTaskVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMCloneTaskVmWithDefaults() *VcenterVMCloneTaskVm {
	this := VcenterVMCloneTaskVm{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVMCloneTaskVm) GetSpec() VcenterVMCloneSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVMCloneSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMCloneTaskVm) GetSpecOk() (*VcenterVMCloneSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVMCloneTaskVm) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVMCloneSpec and assigns it to the Spec field.
func (o *VcenterVMCloneTaskVm) SetSpec(v VcenterVMCloneSpec) {
	o.Spec = &v
}

func (o VcenterVMCloneTaskVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMCloneTaskVm struct {
	value *VcenterVMCloneTaskVm
	isSet bool
}

func (v NullableVcenterVMCloneTaskVm) Get() *VcenterVMCloneTaskVm {
	return v.value
}

func (v *NullableVcenterVMCloneTaskVm) Set(val *VcenterVMCloneTaskVm) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMCloneTaskVm) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMCloneTaskVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMCloneTaskVm(val *VcenterVMCloneTaskVm) *NullableVcenterVMCloneTaskVm {
	return &NullableVcenterVMCloneTaskVm{value: val, isSet: true}
}

func (v NullableVcenterVMCloneTaskVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMCloneTaskVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

