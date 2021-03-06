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

// VcenterVmConsoleTicketsCreateSpec struct for VcenterVmConsoleTicketsCreateSpec
type VcenterVmConsoleTicketsCreateSpec struct {
	Type VcenterVmConsoleTicketsType `json:"type"`
}

// NewVcenterVmConsoleTicketsCreateSpec instantiates a new VcenterVmConsoleTicketsCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmConsoleTicketsCreateSpec(type_ VcenterVmConsoleTicketsType) *VcenterVmConsoleTicketsCreateSpec {
	this := VcenterVmConsoleTicketsCreateSpec{}
	this.Type = type_
	return &this
}

// NewVcenterVmConsoleTicketsCreateSpecWithDefaults instantiates a new VcenterVmConsoleTicketsCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmConsoleTicketsCreateSpecWithDefaults() *VcenterVmConsoleTicketsCreateSpec {
	this := VcenterVmConsoleTicketsCreateSpec{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmConsoleTicketsCreateSpec) GetType() VcenterVmConsoleTicketsType {
	if o == nil {
		var ret VcenterVmConsoleTicketsType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmConsoleTicketsCreateSpec) GetTypeOk() (*VcenterVmConsoleTicketsType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmConsoleTicketsCreateSpec) SetType(v VcenterVmConsoleTicketsType) {
	o.Type = v
}

func (o VcenterVmConsoleTicketsCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmConsoleTicketsCreateSpec struct {
	value *VcenterVmConsoleTicketsCreateSpec
	isSet bool
}

func (v NullableVcenterVmConsoleTicketsCreateSpec) Get() *VcenterVmConsoleTicketsCreateSpec {
	return v.value
}

func (v *NullableVcenterVmConsoleTicketsCreateSpec) Set(val *VcenterVmConsoleTicketsCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmConsoleTicketsCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmConsoleTicketsCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmConsoleTicketsCreateSpec(val *VcenterVmConsoleTicketsCreateSpec) *NullableVcenterVmConsoleTicketsCreateSpec {
	return &NullableVcenterVmConsoleTicketsCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmConsoleTicketsCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmConsoleTicketsCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


