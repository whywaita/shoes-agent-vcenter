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

// VcenterVMInfoSataAdapters struct for VcenterVMInfoSataAdapters
type VcenterVMInfoSataAdapters struct {
	Key *string `json:"key,omitempty"`
	Value *VcenterVmHardwareAdapterSataInfo `json:"value,omitempty"`
}

// NewVcenterVMInfoSataAdapters instantiates a new VcenterVMInfoSataAdapters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMInfoSataAdapters() *VcenterVMInfoSataAdapters {
	this := VcenterVMInfoSataAdapters{}
	return &this
}

// NewVcenterVMInfoSataAdaptersWithDefaults instantiates a new VcenterVMInfoSataAdapters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMInfoSataAdaptersWithDefaults() *VcenterVMInfoSataAdapters {
	this := VcenterVMInfoSataAdapters{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterVMInfoSataAdapters) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInfoSataAdapters) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterVMInfoSataAdapters) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterVMInfoSataAdapters) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterVMInfoSataAdapters) GetValue() VcenterVmHardwareAdapterSataInfo {
	if o == nil || o.Value == nil {
		var ret VcenterVmHardwareAdapterSataInfo
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInfoSataAdapters) GetValueOk() (*VcenterVmHardwareAdapterSataInfo, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterVMInfoSataAdapters) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VcenterVmHardwareAdapterSataInfo and assigns it to the Value field.
func (o *VcenterVMInfoSataAdapters) SetValue(v VcenterVmHardwareAdapterSataInfo) {
	o.Value = &v
}

func (o VcenterVMInfoSataAdapters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMInfoSataAdapters struct {
	value *VcenterVMInfoSataAdapters
	isSet bool
}

func (v NullableVcenterVMInfoSataAdapters) Get() *VcenterVMInfoSataAdapters {
	return v.value
}

func (v *NullableVcenterVMInfoSataAdapters) Set(val *VcenterVMInfoSataAdapters) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMInfoSataAdapters) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMInfoSataAdapters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMInfoSataAdapters(val *VcenterVMInfoSataAdapters) *NullableVcenterVMInfoSataAdapters {
	return &NullableVcenterVMInfoSataAdapters{value: val, isSet: true}
}

func (v NullableVcenterVMInfoSataAdapters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMInfoSataAdapters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

