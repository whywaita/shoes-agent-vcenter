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

// VcenterVmHardwareNvmeAddressInfo struct for VcenterVmHardwareNvmeAddressInfo
type VcenterVmHardwareNvmeAddressInfo struct {
	// Bus number of the adapter to which the device is attached.
	Bus int64 `json:"bus"`
	// Unit number of the device.
	Unit int64 `json:"unit"`
}

// NewVcenterVmHardwareNvmeAddressInfo instantiates a new VcenterVmHardwareNvmeAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareNvmeAddressInfo(bus int64, unit int64) *VcenterVmHardwareNvmeAddressInfo {
	this := VcenterVmHardwareNvmeAddressInfo{}
	this.Bus = bus
	this.Unit = unit
	return &this
}

// NewVcenterVmHardwareNvmeAddressInfoWithDefaults instantiates a new VcenterVmHardwareNvmeAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareNvmeAddressInfoWithDefaults() *VcenterVmHardwareNvmeAddressInfo {
	this := VcenterVmHardwareNvmeAddressInfo{}
	return &this
}

// GetBus returns the Bus field value
func (o *VcenterVmHardwareNvmeAddressInfo) GetBus() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Bus
}

// GetBusOk returns a tuple with the Bus field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareNvmeAddressInfo) GetBusOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bus, true
}

// SetBus sets field value
func (o *VcenterVmHardwareNvmeAddressInfo) SetBus(v int64) {
	o.Bus = v
}

// GetUnit returns the Unit field value
func (o *VcenterVmHardwareNvmeAddressInfo) GetUnit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareNvmeAddressInfo) GetUnitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *VcenterVmHardwareNvmeAddressInfo) SetUnit(v int64) {
	o.Unit = v
}

func (o VcenterVmHardwareNvmeAddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["bus"] = o.Bus
	}
	if true {
		toSerialize["unit"] = o.Unit
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareNvmeAddressInfo struct {
	value *VcenterVmHardwareNvmeAddressInfo
	isSet bool
}

func (v NullableVcenterVmHardwareNvmeAddressInfo) Get() *VcenterVmHardwareNvmeAddressInfo {
	return v.value
}

func (v *NullableVcenterVmHardwareNvmeAddressInfo) Set(val *VcenterVmHardwareNvmeAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareNvmeAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareNvmeAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareNvmeAddressInfo(val *VcenterVmHardwareNvmeAddressInfo) *NullableVcenterVmHardwareNvmeAddressInfo {
	return &NullableVcenterVmHardwareNvmeAddressInfo{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareNvmeAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareNvmeAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


