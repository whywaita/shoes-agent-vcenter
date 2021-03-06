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

// VcenterVmTemplateLibraryItemsCpuInfo struct for VcenterVmTemplateLibraryItemsCpuInfo
type VcenterVmTemplateLibraryItemsCpuInfo struct {
	// Number of CPU cores.
	Count int64 `json:"count"`
	// Number of CPU cores per socket.
	CoresPerSocket int64 `json:"cores_per_socket"`
}

// NewVcenterVmTemplateLibraryItemsCpuInfo instantiates a new VcenterVmTemplateLibraryItemsCpuInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCpuInfo(count int64, coresPerSocket int64) *VcenterVmTemplateLibraryItemsCpuInfo {
	this := VcenterVmTemplateLibraryItemsCpuInfo{}
	this.Count = count
	this.CoresPerSocket = coresPerSocket
	return &this
}

// NewVcenterVmTemplateLibraryItemsCpuInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCpuInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCpuInfoWithDefaults() *VcenterVmTemplateLibraryItemsCpuInfo {
	this := VcenterVmTemplateLibraryItemsCpuInfo{}
	return &this
}

// GetCount returns the Count field value
func (o *VcenterVmTemplateLibraryItemsCpuInfo) GetCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCpuInfo) GetCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Count, true
}

// SetCount sets field value
func (o *VcenterVmTemplateLibraryItemsCpuInfo) SetCount(v int64) {
	o.Count = v
}

// GetCoresPerSocket returns the CoresPerSocket field value
func (o *VcenterVmTemplateLibraryItemsCpuInfo) GetCoresPerSocket() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CoresPerSocket
}

// GetCoresPerSocketOk returns a tuple with the CoresPerSocket field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCpuInfo) GetCoresPerSocketOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CoresPerSocket, true
}

// SetCoresPerSocket sets field value
func (o *VcenterVmTemplateLibraryItemsCpuInfo) SetCoresPerSocket(v int64) {
	o.CoresPerSocket = v
}

func (o VcenterVmTemplateLibraryItemsCpuInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["count"] = o.Count
	}
	if true {
		toSerialize["cores_per_socket"] = o.CoresPerSocket
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCpuInfo struct {
	value *VcenterVmTemplateLibraryItemsCpuInfo
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCpuInfo) Get() *VcenterVmTemplateLibraryItemsCpuInfo {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCpuInfo) Set(val *VcenterVmTemplateLibraryItemsCpuInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCpuInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCpuInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCpuInfo(val *VcenterVmTemplateLibraryItemsCpuInfo) *NullableVcenterVmTemplateLibraryItemsCpuInfo {
	return &NullableVcenterVmTemplateLibraryItemsCpuInfo{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCpuInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCpuInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


