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

// VcenterVMInstantCloneSpecParallelPortsToUpdate struct for VcenterVMInstantCloneSpecParallelPortsToUpdate
type VcenterVMInstantCloneSpecParallelPortsToUpdate struct {
	Key *string `json:"key,omitempty"`
	Value *VcenterVmHardwareParallelUpdateSpec `json:"value,omitempty"`
}

// NewVcenterVMInstantCloneSpecParallelPortsToUpdate instantiates a new VcenterVMInstantCloneSpecParallelPortsToUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMInstantCloneSpecParallelPortsToUpdate() *VcenterVMInstantCloneSpecParallelPortsToUpdate {
	this := VcenterVMInstantCloneSpecParallelPortsToUpdate{}
	return &this
}

// NewVcenterVMInstantCloneSpecParallelPortsToUpdateWithDefaults instantiates a new VcenterVMInstantCloneSpecParallelPortsToUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMInstantCloneSpecParallelPortsToUpdateWithDefaults() *VcenterVMInstantCloneSpecParallelPortsToUpdate {
	this := VcenterVMInstantCloneSpecParallelPortsToUpdate{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) GetValue() VcenterVmHardwareParallelUpdateSpec {
	if o == nil || o.Value == nil {
		var ret VcenterVmHardwareParallelUpdateSpec
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) GetValueOk() (*VcenterVmHardwareParallelUpdateSpec, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VcenterVmHardwareParallelUpdateSpec and assigns it to the Value field.
func (o *VcenterVMInstantCloneSpecParallelPortsToUpdate) SetValue(v VcenterVmHardwareParallelUpdateSpec) {
	o.Value = &v
}

func (o VcenterVMInstantCloneSpecParallelPortsToUpdate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMInstantCloneSpecParallelPortsToUpdate struct {
	value *VcenterVMInstantCloneSpecParallelPortsToUpdate
	isSet bool
}

func (v NullableVcenterVMInstantCloneSpecParallelPortsToUpdate) Get() *VcenterVMInstantCloneSpecParallelPortsToUpdate {
	return v.value
}

func (v *NullableVcenterVMInstantCloneSpecParallelPortsToUpdate) Set(val *VcenterVMInstantCloneSpecParallelPortsToUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMInstantCloneSpecParallelPortsToUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMInstantCloneSpecParallelPortsToUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMInstantCloneSpecParallelPortsToUpdate(val *VcenterVMInstantCloneSpecParallelPortsToUpdate) *NullableVcenterVMInstantCloneSpecParallelPortsToUpdate {
	return &NullableVcenterVMInstantCloneSpecParallelPortsToUpdate{value: val, isSet: true}
}

func (v NullableVcenterVMInstantCloneSpecParallelPortsToUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMInstantCloneSpecParallelPortsToUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

