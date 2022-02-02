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

// VcenterIsoImageMountIsoImage struct for VcenterIsoImageMountIsoImage
type VcenterIsoImageMountIsoImage struct {
	// The identifier of the virtual machine where the specified ISO image will be mounted.
	Vm *string `json:"vm,omitempty"`
}

// NewVcenterIsoImageMountIsoImage instantiates a new VcenterIsoImageMountIsoImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIsoImageMountIsoImage() *VcenterIsoImageMountIsoImage {
	this := VcenterIsoImageMountIsoImage{}
	return &this
}

// NewVcenterIsoImageMountIsoImageWithDefaults instantiates a new VcenterIsoImageMountIsoImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIsoImageMountIsoImageWithDefaults() *VcenterIsoImageMountIsoImage {
	this := VcenterIsoImageMountIsoImage{}
	return &this
}

// GetVm returns the Vm field value if set, zero value otherwise.
func (o *VcenterIsoImageMountIsoImage) GetVm() string {
	if o == nil || o.Vm == nil {
		var ret string
		return ret
	}
	return *o.Vm
}

// GetVmOk returns a tuple with the Vm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterIsoImageMountIsoImage) GetVmOk() (*string, bool) {
	if o == nil || o.Vm == nil {
		return nil, false
	}
	return o.Vm, true
}

// HasVm returns a boolean if a field has been set.
func (o *VcenterIsoImageMountIsoImage) HasVm() bool {
	if o != nil && o.Vm != nil {
		return true
	}

	return false
}

// SetVm gets a reference to the given string and assigns it to the Vm field.
func (o *VcenterIsoImageMountIsoImage) SetVm(v string) {
	o.Vm = &v
}

func (o VcenterIsoImageMountIsoImage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Vm != nil {
		toSerialize["vm"] = o.Vm
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterIsoImageMountIsoImage struct {
	value *VcenterIsoImageMountIsoImage
	isSet bool
}

func (v NullableVcenterIsoImageMountIsoImage) Get() *VcenterIsoImageMountIsoImage {
	return v.value
}

func (v *NullableVcenterIsoImageMountIsoImage) Set(val *VcenterIsoImageMountIsoImage) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIsoImageMountIsoImage) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIsoImageMountIsoImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIsoImageMountIsoImage(val *VcenterIsoImageMountIsoImage) *NullableVcenterIsoImageMountIsoImage {
	return &NullableVcenterIsoImageMountIsoImage{value: val, isSet: true}
}

func (v NullableVcenterIsoImageMountIsoImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIsoImageMountIsoImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


