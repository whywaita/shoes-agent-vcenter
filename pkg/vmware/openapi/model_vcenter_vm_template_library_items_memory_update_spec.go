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

// VcenterVmTemplateLibraryItemsMemoryUpdateSpec struct for VcenterVmTemplateLibraryItemsMemoryUpdateSpec
type VcenterVmTemplateLibraryItemsMemoryUpdateSpec struct {
	// Size of a virtual machine's memory in MB.
	Memory *int64 `json:"memory,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsMemoryUpdateSpec instantiates a new VcenterVmTemplateLibraryItemsMemoryUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsMemoryUpdateSpec() *VcenterVmTemplateLibraryItemsMemoryUpdateSpec {
	this := VcenterVmTemplateLibraryItemsMemoryUpdateSpec{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsMemoryUpdateSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsMemoryUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsMemoryUpdateSpecWithDefaults() *VcenterVmTemplateLibraryItemsMemoryUpdateSpec {
	this := VcenterVmTemplateLibraryItemsMemoryUpdateSpec{}
	return &this
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsMemoryUpdateSpec) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsMemoryUpdateSpec) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsMemoryUpdateSpec) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *VcenterVmTemplateLibraryItemsMemoryUpdateSpec) SetMemory(v int64) {
	o.Memory = &v
}

func (o VcenterVmTemplateLibraryItemsMemoryUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Memory != nil {
		toSerialize["memory"] = o.Memory
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec struct {
	value *VcenterVmTemplateLibraryItemsMemoryUpdateSpec
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec) Get() *VcenterVmTemplateLibraryItemsMemoryUpdateSpec {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec) Set(val *VcenterVmTemplateLibraryItemsMemoryUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec(val *VcenterVmTemplateLibraryItemsMemoryUpdateSpec) *NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec {
	return &NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsMemoryUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


