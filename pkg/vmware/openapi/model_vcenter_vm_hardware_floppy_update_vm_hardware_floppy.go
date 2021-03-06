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

// VcenterVmHardwareFloppyUpdateVmHardwareFloppy struct for VcenterVmHardwareFloppyUpdateVmHardwareFloppy
type VcenterVmHardwareFloppyUpdateVmHardwareFloppy struct {
	Spec *VcenterVmHardwareFloppyUpdateSpec `json:"spec,omitempty"`
}

// NewVcenterVmHardwareFloppyUpdateVmHardwareFloppy instantiates a new VcenterVmHardwareFloppyUpdateVmHardwareFloppy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareFloppyUpdateVmHardwareFloppy() *VcenterVmHardwareFloppyUpdateVmHardwareFloppy {
	this := VcenterVmHardwareFloppyUpdateVmHardwareFloppy{}
	return &this
}

// NewVcenterVmHardwareFloppyUpdateVmHardwareFloppyWithDefaults instantiates a new VcenterVmHardwareFloppyUpdateVmHardwareFloppy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareFloppyUpdateVmHardwareFloppyWithDefaults() *VcenterVmHardwareFloppyUpdateVmHardwareFloppy {
	this := VcenterVmHardwareFloppyUpdateVmHardwareFloppy{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVmHardwareFloppyUpdateVmHardwareFloppy) GetSpec() VcenterVmHardwareFloppyUpdateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVmHardwareFloppyUpdateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareFloppyUpdateVmHardwareFloppy) GetSpecOk() (*VcenterVmHardwareFloppyUpdateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVmHardwareFloppyUpdateVmHardwareFloppy) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVmHardwareFloppyUpdateSpec and assigns it to the Spec field.
func (o *VcenterVmHardwareFloppyUpdateVmHardwareFloppy) SetSpec(v VcenterVmHardwareFloppyUpdateSpec) {
	o.Spec = &v
}

func (o VcenterVmHardwareFloppyUpdateVmHardwareFloppy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy struct {
	value *VcenterVmHardwareFloppyUpdateVmHardwareFloppy
	isSet bool
}

func (v NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy) Get() *VcenterVmHardwareFloppyUpdateVmHardwareFloppy {
	return v.value
}

func (v *NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy) Set(val *VcenterVmHardwareFloppyUpdateVmHardwareFloppy) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy(val *VcenterVmHardwareFloppyUpdateVmHardwareFloppy) *NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy {
	return &NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareFloppyUpdateVmHardwareFloppy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


