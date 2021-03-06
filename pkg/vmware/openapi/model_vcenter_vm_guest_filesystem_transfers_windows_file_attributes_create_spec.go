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

// VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec struct for VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec
type VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec struct {
	// The file is hidden. If unset the file will not be hidden.
	Hidden *bool `json:"hidden,omitempty"`
	// The file is read-only. If unset the file will be writeable.
	ReadOnly *bool `json:"read_only,omitempty"`
}

// NewVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec instantiates a new VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec() *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec {
	this := VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec{}
	return &this
}

// NewVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpecWithDefaults instantiates a new VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpecWithDefaults() *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec {
	this := VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec{}
	return &this
}

// GetHidden returns the Hidden field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) GetHidden() bool {
	if o == nil || o.Hidden == nil {
		var ret bool
		return ret
	}
	return *o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) GetHiddenOk() (*bool, bool) {
	if o == nil || o.Hidden == nil {
		return nil, false
	}
	return o.Hidden, true
}

// HasHidden returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) HasHidden() bool {
	if o != nil && o.Hidden != nil {
		return true
	}

	return false
}

// SetHidden gets a reference to the given bool and assigns it to the Hidden field.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) SetHidden(v bool) {
	o.Hidden = &v
}

// GetReadOnly returns the ReadOnly field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) GetReadOnly() bool {
	if o == nil || o.ReadOnly == nil {
		var ret bool
		return ret
	}
	return *o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) GetReadOnlyOk() (*bool, bool) {
	if o == nil || o.ReadOnly == nil {
		return nil, false
	}
	return o.ReadOnly, true
}

// HasReadOnly returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) HasReadOnly() bool {
	if o != nil && o.ReadOnly != nil {
		return true
	}

	return false
}

// SetReadOnly gets a reference to the given bool and assigns it to the ReadOnly field.
func (o *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) SetReadOnly(v bool) {
	o.ReadOnly = &v
}

func (o VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hidden != nil {
		toSerialize["hidden"] = o.Hidden
	}
	if o.ReadOnly != nil {
		toSerialize["read_only"] = o.ReadOnly
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec struct {
	value *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec
	isSet bool
}

func (v NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) Get() *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec {
	return v.value
}

func (v *NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) Set(val *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec(val *VcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) *NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec {
	return &NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestFilesystemTransfersWindowsFileAttributesCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


