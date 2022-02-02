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

// VcenterVmHardwareCpuInfo struct for VcenterVmHardwareCpuInfo
type VcenterVmHardwareCpuInfo struct {
	// Number of CPU cores.
	Count int64 `json:"count"`
	// Number of CPU cores per socket.
	CoresPerSocket int64 `json:"cores_per_socket"`
	// Flag indicating whether adding CPUs while the virtual machine is running is enabled.
	HotAddEnabled bool `json:"hot_add_enabled"`
	// Flag indicating whether removing CPUs while the virtual machine is running is enabled.
	HotRemoveEnabled bool `json:"hot_remove_enabled"`
}

// NewVcenterVmHardwareCpuInfo instantiates a new VcenterVmHardwareCpuInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareCpuInfo(count int64, coresPerSocket int64, hotAddEnabled bool, hotRemoveEnabled bool) *VcenterVmHardwareCpuInfo {
	this := VcenterVmHardwareCpuInfo{}
	this.Count = count
	this.CoresPerSocket = coresPerSocket
	this.HotAddEnabled = hotAddEnabled
	this.HotRemoveEnabled = hotRemoveEnabled
	return &this
}

// NewVcenterVmHardwareCpuInfoWithDefaults instantiates a new VcenterVmHardwareCpuInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareCpuInfoWithDefaults() *VcenterVmHardwareCpuInfo {
	this := VcenterVmHardwareCpuInfo{}
	return &this
}

// GetCount returns the Count field value
func (o *VcenterVmHardwareCpuInfo) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCpuInfo) GetCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *VcenterVmHardwareCpuInfo) SetCount(v int64) {
	o.Count = v
}

// GetCoresPerSocket returns the CoresPerSocket field value
func (o *VcenterVmHardwareCpuInfo) GetCoresPerSocket() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCpuInfo) GetCoresPerSocketOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CoresPerSocket, true
}

// SetCoresPerSocket sets field value
func (o *VcenterVmHardwareCpuInfo) SetCoresPerSocket(v int64) {
	o.CoresPerSocket = v
}

// GetHotAddEnabled returns the HotAddEnabled field value
func (o *VcenterVmHardwareCpuInfo) GetHotAddEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HotAddEnabled
}

// GetHotAddEnabledOk returns a tuple with the HotAddEnabled field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCpuInfo) GetHotAddEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HotAddEnabled, true
}

// SetHotAddEnabled sets field value
func (o *VcenterVmHardwareCpuInfo) SetHotAddEnabled(v bool) {
	o.HotAddEnabled = v
}

// GetHotRemoveEnabled returns the HotRemoveEnabled field value
func (o *VcenterVmHardwareCpuInfo) GetHotRemoveEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.HotRemoveEnabled
}

// GetHotRemoveEnabledOk returns a tuple with the HotRemoveEnabled field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCpuInfo) GetHotRemoveEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.HotRemoveEnabled, true
}

// SetHotRemoveEnabled sets field value
func (o *VcenterVmHardwareCpuInfo) SetHotRemoveEnabled(v bool) {
	o.HotRemoveEnabled = v
}

func (o VcenterVmHardwareCpuInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["cores_per_socket"] = o.CoresPerSocket
	}
	if true {
		toSerialize["hot_add_enabled"] = o.HotAddEnabled
	}
	if true {
		toSerialize["hot_remove_enabled"] = o.HotRemoveEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareCpuInfo struct {
	value *VcenterVmHardwareCpuInfo
	isSet bool
}

func (v NullableVcenterVmHardwareCpuInfo) Get() *VcenterVmHardwareCpuInfo {
	return v.value
}

func (v *NullableVcenterVmHardwareCpuInfo) Set(val *VcenterVmHardwareCpuInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareCpuInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareCpuInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareCpuInfo(val *VcenterVmHardwareCpuInfo) *NullableVcenterVmHardwareCpuInfo {
	return &NullableVcenterVmHardwareCpuInfo{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareCpuInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareCpuInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


