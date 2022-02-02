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

// VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides struct for VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides
type VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides struct {
	Key *string `json:"key,omitempty"`
	Value *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage `json:"value,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides instantiates a new VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides() *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides {
	this := VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverridesWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverridesWithDefaults() *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides {
	this := VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) GetValue() VcenterVmTemplateLibraryItemsCreateSpecDiskStorage {
	if o == nil || o.Value == nil {
		var ret VcenterVmTemplateLibraryItemsCreateSpecDiskStorage
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) GetValueOk() (*VcenterVmTemplateLibraryItemsCreateSpecDiskStorage, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VcenterVmTemplateLibraryItemsCreateSpecDiskStorage and assigns it to the Value field.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) SetValue(v VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) {
	o.Value = &v
}

func (o VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides struct {
	value *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) Get() *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) Set(val *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides(val *VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides {
	return &NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

