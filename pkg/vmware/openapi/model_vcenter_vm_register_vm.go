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

// VcenterVMRegisterVm struct for VcenterVMRegisterVm
type VcenterVMRegisterVm struct {
	Spec *VcenterVMRegisterSpec `json:"spec,omitempty"`
}

// NewVcenterVMRegisterVm instantiates a new VcenterVMRegisterVm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMRegisterVm() *VcenterVMRegisterVm {
	this := VcenterVMRegisterVm{}
	return &this
}

// NewVcenterVMRegisterVmWithDefaults instantiates a new VcenterVMRegisterVm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMRegisterVmWithDefaults() *VcenterVMRegisterVm {
	this := VcenterVMRegisterVm{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVMRegisterVm) GetSpec() VcenterVMRegisterSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVMRegisterSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRegisterVm) GetSpecOk() (*VcenterVMRegisterSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVMRegisterVm) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVMRegisterSpec and assigns it to the Spec field.
func (o *VcenterVMRegisterVm) SetSpec(v VcenterVMRegisterSpec) {
	o.Spec = &v
}

func (o VcenterVMRegisterVm) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMRegisterVm struct {
	value *VcenterVMRegisterVm
	isSet bool
}

func (v NullableVcenterVMRegisterVm) Get() *VcenterVMRegisterVm {
	return v.value
}

func (v *NullableVcenterVMRegisterVm) Set(val *VcenterVMRegisterVm) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMRegisterVm) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMRegisterVm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMRegisterVm(val *VcenterVMRegisterVm) *NullableVcenterVMRegisterVm {
	return &NullableVcenterVMRegisterVm{value: val, isSet: true}
}

func (v NullableVcenterVMRegisterVm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMRegisterVm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


