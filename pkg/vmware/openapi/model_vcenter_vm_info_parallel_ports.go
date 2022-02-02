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

// VcenterVMInfoParallelPorts struct for VcenterVMInfoParallelPorts
type VcenterVMInfoParallelPorts struct {
	Key *string `json:"key,omitempty"`
	Value *VcenterVmHardwareParallelInfo `json:"value,omitempty"`
}

// NewVcenterVMInfoParallelPorts instantiates a new VcenterVMInfoParallelPorts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMInfoParallelPorts() *VcenterVMInfoParallelPorts {
	this := VcenterVMInfoParallelPorts{}
	return &this
}

// NewVcenterVMInfoParallelPortsWithDefaults instantiates a new VcenterVMInfoParallelPorts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMInfoParallelPortsWithDefaults() *VcenterVMInfoParallelPorts {
	this := VcenterVMInfoParallelPorts{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *VcenterVMInfoParallelPorts) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInfoParallelPorts) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *VcenterVMInfoParallelPorts) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *VcenterVMInfoParallelPorts) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VcenterVMInfoParallelPorts) GetValue() VcenterVmHardwareParallelInfo {
	if o == nil || o.Value == nil {
		var ret VcenterVmHardwareParallelInfo
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMInfoParallelPorts) GetValueOk() (*VcenterVmHardwareParallelInfo, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VcenterVMInfoParallelPorts) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VcenterVmHardwareParallelInfo and assigns it to the Value field.
func (o *VcenterVMInfoParallelPorts) SetValue(v VcenterVmHardwareParallelInfo) {
	o.Value = &v
}

func (o VcenterVMInfoParallelPorts) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["key"] = o.Key
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMInfoParallelPorts struct {
	value *VcenterVMInfoParallelPorts
	isSet bool
}

func (v NullableVcenterVMInfoParallelPorts) Get() *VcenterVMInfoParallelPorts {
	return v.value
}

func (v *NullableVcenterVMInfoParallelPorts) Set(val *VcenterVMInfoParallelPorts) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMInfoParallelPorts) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMInfoParallelPorts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMInfoParallelPorts(val *VcenterVMInfoParallelPorts) *NullableVcenterVMInfoParallelPorts {
	return &NullableVcenterVMInfoParallelPorts{value: val, isSet: true}
}

func (v NullableVcenterVMInfoParallelPorts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMInfoParallelPorts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


