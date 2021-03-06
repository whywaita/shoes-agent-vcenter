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

// VcenterResourcePoolInfo struct for VcenterResourcePoolInfo
type VcenterResourcePoolInfo struct {
	// Name of the vCenter Server resource pool.
	Name string `json:"name"`
	// Identifiers of the child resource pools contained in this resource pool. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool.
	ResourcePools []string `json:"resource_pools"`
	CpuAllocation *VcenterResourcePoolResourceAllocationInfo `json:"cpu_allocation,omitempty"`
	MemoryAllocation *VcenterResourcePoolResourceAllocationInfo `json:"memory_allocation,omitempty"`
}

// NewVcenterResourcePoolInfo instantiates a new VcenterResourcePoolInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterResourcePoolInfo(name string, resourcePools []string) *VcenterResourcePoolInfo {
	this := VcenterResourcePoolInfo{}
	this.Name = name
	this.ResourcePools = resourcePools
	return &this
}

// NewVcenterResourcePoolInfoWithDefaults instantiates a new VcenterResourcePoolInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterResourcePoolInfoWithDefaults() *VcenterResourcePoolInfo {
	this := VcenterResourcePoolInfo{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterResourcePoolInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterResourcePoolInfo) SetName(v string) {
	o.Name = v
}

// GetResourcePools returns the ResourcePools field value
func (o *VcenterResourcePoolInfo) GetResourcePools() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResourcePools
}

// GetResourcePoolsOk returns a tuple with the ResourcePools field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolInfo) GetResourcePoolsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ResourcePools, true
}

// SetResourcePools sets field value
func (o *VcenterResourcePoolInfo) SetResourcePools(v []string) {
	o.ResourcePools = v
}

// GetCpuAllocation returns the CpuAllocation field value if set, zero value otherwise.
func (o *VcenterResourcePoolInfo) GetCpuAllocation() VcenterResourcePoolResourceAllocationInfo {
	if o == nil || o.CpuAllocation == nil {
		var ret VcenterResourcePoolResourceAllocationInfo
		return ret
	}
	return *o.CpuAllocation
}

// GetCpuAllocationOk returns a tuple with the CpuAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolInfo) GetCpuAllocationOk() (*VcenterResourcePoolResourceAllocationInfo, bool) {
	if o == nil || o.CpuAllocation == nil {
		return nil, false
	}
	return o.CpuAllocation, true
}

// HasCpuAllocation returns a boolean if a field has been set.
func (o *VcenterResourcePoolInfo) HasCpuAllocation() bool {
	if o != nil && o.CpuAllocation != nil {
		return true
	}

	return false
}

// SetCpuAllocation gets a reference to the given VcenterResourcePoolResourceAllocationInfo and assigns it to the CpuAllocation field.
func (o *VcenterResourcePoolInfo) SetCpuAllocation(v VcenterResourcePoolResourceAllocationInfo) {
	o.CpuAllocation = &v
}

// GetMemoryAllocation returns the MemoryAllocation field value if set, zero value otherwise.
func (o *VcenterResourcePoolInfo) GetMemoryAllocation() VcenterResourcePoolResourceAllocationInfo {
	if o == nil || o.MemoryAllocation == nil {
		var ret VcenterResourcePoolResourceAllocationInfo
		return ret
	}
	return *o.MemoryAllocation
}

// GetMemoryAllocationOk returns a tuple with the MemoryAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolInfo) GetMemoryAllocationOk() (*VcenterResourcePoolResourceAllocationInfo, bool) {
	if o == nil || o.MemoryAllocation == nil {
		return nil, false
	}
	return o.MemoryAllocation, true
}

// HasMemoryAllocation returns a boolean if a field has been set.
func (o *VcenterResourcePoolInfo) HasMemoryAllocation() bool {
	if o != nil && o.MemoryAllocation != nil {
		return true
	}

	return false
}

// SetMemoryAllocation gets a reference to the given VcenterResourcePoolResourceAllocationInfo and assigns it to the MemoryAllocation field.
func (o *VcenterResourcePoolInfo) SetMemoryAllocation(v VcenterResourcePoolResourceAllocationInfo) {
	o.MemoryAllocation = &v
}

func (o VcenterResourcePoolInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["resource_pools"] = o.ResourcePools
	}
	if o.CpuAllocation != nil {
		toSerialize["cpu_allocation"] = o.CpuAllocation
	}
	if o.MemoryAllocation != nil {
		toSerialize["memory_allocation"] = o.MemoryAllocation
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterResourcePoolInfo struct {
	value *VcenterResourcePoolInfo
	isSet bool
}

func (v NullableVcenterResourcePoolInfo) Get() *VcenterResourcePoolInfo {
	return v.value
}

func (v *NullableVcenterResourcePoolInfo) Set(val *VcenterResourcePoolInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterResourcePoolInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterResourcePoolInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterResourcePoolInfo(val *VcenterResourcePoolInfo) *NullableVcenterResourcePoolInfo {
	return &NullableVcenterResourcePoolInfo{value: val, isSet: true}
}

func (v NullableVcenterResourcePoolInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterResourcePoolInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


