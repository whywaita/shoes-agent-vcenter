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

// VcenterVmHardwareSataAddressInfo struct for VcenterVmHardwareSataAddressInfo
type VcenterVmHardwareSataAddressInfo struct {
	// Bus number of the adapter to which the device is attached.
	Bus int64 `json:"bus"`
	// Unit number of the device.
	Unit int64 `json:"unit"`
}

// NewVcenterVmHardwareSataAddressInfo instantiates a new VcenterVmHardwareSataAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareSataAddressInfo(bus int64, unit int64) *VcenterVmHardwareSataAddressInfo {
	this := VcenterVmHardwareSataAddressInfo{}
	this.Bus = bus
	this.Unit = unit
	return &this
}

// NewVcenterVmHardwareSataAddressInfoWithDefaults instantiates a new VcenterVmHardwareSataAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareSataAddressInfoWithDefaults() *VcenterVmHardwareSataAddressInfo {
	this := VcenterVmHardwareSataAddressInfo{}
	return &this
}

// GetBus returns the Bus field value
func (o *VcenterVmHardwareSataAddressInfo) GetBus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Bus
}

// GetBusOk returns a tuple with the Bus field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSataAddressInfo) GetBusOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bus, true
}

// SetBus sets field value
func (o *VcenterVmHardwareSataAddressInfo) SetBus(v int64) {
	o.Bus = v
}

// GetUnit returns the Unit field value
func (o *VcenterVmHardwareSataAddressInfo) GetUnit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSataAddressInfo) GetUnitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *VcenterVmHardwareSataAddressInfo) SetUnit(v int64) {
	o.Unit = v
}

func (o VcenterVmHardwareSataAddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bus"] = o.Bus
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareSataAddressInfo struct {
	value *VcenterVmHardwareSataAddressInfo
	isSet bool
}

func (v NullableVcenterVmHardwareSataAddressInfo) Get() *VcenterVmHardwareSataAddressInfo {
	return v.value
}

func (v *NullableVcenterVmHardwareSataAddressInfo) Set(val *VcenterVmHardwareSataAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareSataAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareSataAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareSataAddressInfo(val *VcenterVmHardwareSataAddressInfo) *NullableVcenterVmHardwareSataAddressInfo {
	return &NullableVcenterVmHardwareSataAddressInfo{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareSataAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareSataAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

