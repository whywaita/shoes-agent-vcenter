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

// VcenterVMInstantCloneVm struct for VcenterVMInstantCloneVm
type VcenterVMInstantCloneVm struct {
	Spec *VcenterVMInstantCloneSpec `json:"spec,omitempty"`
}

// NewVcenterVMInstantCloneVm instantiates a new VcenterVMInstantCloneVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMInstantCloneVm() *VcenterVMInstantCloneVm {
	this := VcenterVMInstantCloneVm{}
	return &this
}

// NewVcenterVMInstantCloneVmWithDefaults instantiates a new VcenterVMInstantCloneVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMInstantCloneVmWithDefaults() *VcenterVMInstantCloneVm {
	this := VcenterVMInstantCloneVm{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVMInstantCloneVm) GetSpec() VcenterVMInstantCloneSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVMInstantCloneSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInstantCloneVm) GetSpecOk() (*VcenterVMInstantCloneSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVMInstantCloneVm) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVMInstantCloneSpec and assigns it to the Spec field.
func (o *VcenterVMInstantCloneVm) SetSpec(v VcenterVMInstantCloneSpec) {
	o.Spec = &v
}

func (o VcenterVMInstantCloneVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMInstantCloneVm struct {
	value *VcenterVMInstantCloneVm
	isSet bool
}

func (v NullableVcenterVMInstantCloneVm) Get() *VcenterVMInstantCloneVm {
	return v.value
}

func (v *NullableVcenterVMInstantCloneVm) Set(val *VcenterVMInstantCloneVm) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMInstantCloneVm) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMInstantCloneVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMInstantCloneVm(val *VcenterVMInstantCloneVm) *NullableVcenterVMInstantCloneVm {
	return &NullableVcenterVMInstantCloneVm{value: val, isSet: true}
}

func (v NullableVcenterVMInstantCloneVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMInstantCloneVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


