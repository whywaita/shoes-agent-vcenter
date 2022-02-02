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

// VcenterNamespacesInstancesStats struct for VcenterNamespacesInstancesStats
type VcenterNamespacesInstancesStats struct {
	// Overall CPU usage of the namespace, in MHz. This is the sum of CPU usage across all pods in the Kubernetes namespace.
	CpuUsed int64 `json:"cpu_used"`
	// Overall memory usage of the namespace (in mebibytes). This is the sum of memory usage across all pods.
	MemoryUsed int64 `json:"memory_used"`
	// Overall storage used by the namespace (in mebibytes). This is the sum of storage used by pods across all datastores in the cluster associated with storage policies configured for the namespace.
	StorageUsed int64 `json:"storage_used"`
}

// NewVcenterNamespacesInstancesStats instantiates a new VcenterNamespacesInstancesStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespacesInstancesStats(cpuUsed int64, memoryUsed int64, storageUsed int64) *VcenterNamespacesInstancesStats {
	this := VcenterNamespacesInstancesStats{}
	this.CpuUsed = cpuUsed
	this.MemoryUsed = memoryUsed
	this.StorageUsed = storageUsed
	return &this
}

// NewVcenterNamespacesInstancesStatsWithDefaults instantiates a new VcenterNamespacesInstancesStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespacesInstancesStatsWithDefaults() *VcenterNamespacesInstancesStats {
	this := VcenterNamespacesInstancesStats{}
	return &this
}

// GetCpuUsed returns the CpuUsed field value
func (o *VcenterNamespacesInstancesStats) GetCpuUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CpuUsed
}

// GetCpuUsedOk returns a tuple with the CpuUsed field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesInstancesStats) GetCpuUsedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CpuUsed, true
}

// SetCpuUsed sets field value
func (o *VcenterNamespacesInstancesStats) SetCpuUsed(v int64) {
	o.CpuUsed = v
}

// GetMemoryUsed returns the MemoryUsed field value
func (o *VcenterNamespacesInstancesStats) GetMemoryUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MemoryUsed
}

// GetMemoryUsedOk returns a tuple with the MemoryUsed field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesInstancesStats) GetMemoryUsedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MemoryUsed, true
}

// SetMemoryUsed sets field value
func (o *VcenterNamespacesInstancesStats) SetMemoryUsed(v int64) {
	o.MemoryUsed = v
}

// GetStorageUsed returns the StorageUsed field value
func (o *VcenterNamespacesInstancesStats) GetStorageUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.StorageUsed
}

// GetStorageUsedOk returns a tuple with the StorageUsed field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesInstancesStats) GetStorageUsedOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StorageUsed, true
}

// SetStorageUsed sets field value
func (o *VcenterNamespacesInstancesStats) SetStorageUsed(v int64) {
	o.StorageUsed = v
}

func (o VcenterNamespacesInstancesStats) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cpu_used"] = o.CpuUsed
	}
	if true {
		toSerialize["memory_used"] = o.MemoryUsed
	}
	if true {
		toSerialize["storage_used"] = o.StorageUsed
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespacesInstancesStats struct {
	value *VcenterNamespacesInstancesStats
	isSet bool
}

func (v NullableVcenterNamespacesInstancesStats) Get() *VcenterNamespacesInstancesStats {
	return v.value
}

func (v *NullableVcenterNamespacesInstancesStats) Set(val *VcenterNamespacesInstancesStats) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesInstancesStats) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesInstancesStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesInstancesStats(val *VcenterNamespacesInstancesStats) *NullableVcenterNamespacesInstancesStats {
	return &NullableVcenterNamespacesInstancesStats{value: val, isSet: true}
}

func (v NullableVcenterNamespacesInstancesStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesInstancesStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

