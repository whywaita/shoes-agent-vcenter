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

// VcenterResourcePoolCreateSpec struct for VcenterResourcePoolCreateSpec
type VcenterResourcePoolCreateSpec struct {
	// Name of the resource pool.
	Name string `json:"name"`
	// Parent of the created resource pool. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool.
	Parent string `json:"parent"`
	CpuAllocation *VcenterResourcePoolResourceAllocationCreateSpec `json:"cpu_allocation,omitempty"`
	MemoryAllocation *VcenterResourcePoolResourceAllocationCreateSpec `json:"memory_allocation,omitempty"`
}

// NewVcenterResourcePoolCreateSpec instantiates a new VcenterResourcePoolCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterResourcePoolCreateSpec(name string, parent string) *VcenterResourcePoolCreateSpec {
	this := VcenterResourcePoolCreateSpec{}
	this.Name = name
	this.Parent = parent
	return &this
}

// NewVcenterResourcePoolCreateSpecWithDefaults instantiates a new VcenterResourcePoolCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterResourcePoolCreateSpecWithDefaults() *VcenterResourcePoolCreateSpec {
	this := VcenterResourcePoolCreateSpec{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterResourcePoolCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolCreateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterResourcePoolCreateSpec) SetName(v string) {
	o.Name = v
}

// GetParent returns the Parent field value
func (o *VcenterResourcePoolCreateSpec) GetParent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Parent
}

// GetParentOk returns a tuple with the Parent field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolCreateSpec) GetParentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Parent, true
}

// SetParent sets field value
func (o *VcenterResourcePoolCreateSpec) SetParent(v string) {
	o.Parent = v
}

// GetCpuAllocation returns the CpuAllocation field value if set, zero value otherwise.
func (o *VcenterResourcePoolCreateSpec) GetCpuAllocation() VcenterResourcePoolResourceAllocationCreateSpec {
	if o == nil || o.CpuAllocation == nil {
		var ret VcenterResourcePoolResourceAllocationCreateSpec
		return ret
	}
	return *o.CpuAllocation
}

// GetCpuAllocationOk returns a tuple with the CpuAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolCreateSpec) GetCpuAllocationOk() (*VcenterResourcePoolResourceAllocationCreateSpec, bool) {
	if o == nil || o.CpuAllocation == nil {
		return nil, false
	}
	return o.CpuAllocation, true
}

// HasCpuAllocation returns a boolean if a field has been set.
func (o *VcenterResourcePoolCreateSpec) HasCpuAllocation() bool {
	if o != nil && o.CpuAllocation != nil {
		return true
	}

	return false
}

// SetCpuAllocation gets a reference to the given VcenterResourcePoolResourceAllocationCreateSpec and assigns it to the CpuAllocation field.
func (o *VcenterResourcePoolCreateSpec) SetCpuAllocation(v VcenterResourcePoolResourceAllocationCreateSpec) {
	o.CpuAllocation = &v
}

// GetMemoryAllocation returns the MemoryAllocation field value if set, zero value otherwise.
func (o *VcenterResourcePoolCreateSpec) GetMemoryAllocation() VcenterResourcePoolResourceAllocationCreateSpec {
	if o == nil || o.MemoryAllocation == nil {
		var ret VcenterResourcePoolResourceAllocationCreateSpec
		return ret
	}
	return *o.MemoryAllocation
}

// GetMemoryAllocationOk returns a tuple with the MemoryAllocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolCreateSpec) GetMemoryAllocationOk() (*VcenterResourcePoolResourceAllocationCreateSpec, bool) {
	if o == nil || o.MemoryAllocation == nil {
		return nil, false
	}
	return o.MemoryAllocation, true
}

// HasMemoryAllocation returns a boolean if a field has been set.
func (o *VcenterResourcePoolCreateSpec) HasMemoryAllocation() bool {
	if o != nil && o.MemoryAllocation != nil {
		return true
	}

	return false
}

// SetMemoryAllocation gets a reference to the given VcenterResourcePoolResourceAllocationCreateSpec and assigns it to the MemoryAllocation field.
func (o *VcenterResourcePoolCreateSpec) SetMemoryAllocation(v VcenterResourcePoolResourceAllocationCreateSpec) {
	o.MemoryAllocation = &v
}

func (o VcenterResourcePoolCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["parent"] = o.Parent
	}
	if o.CpuAllocation != nil {
		toSerialize["cpu_allocation"] = o.CpuAllocation
	}
	if o.MemoryAllocation != nil {
		toSerialize["memory_allocation"] = o.MemoryAllocation
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterResourcePoolCreateSpec struct {
	value *VcenterResourcePoolCreateSpec
	isSet bool
}

func (v NullableVcenterResourcePoolCreateSpec) Get() *VcenterResourcePoolCreateSpec {
	return v.value
}

func (v *NullableVcenterResourcePoolCreateSpec) Set(val *VcenterResourcePoolCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterResourcePoolCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterResourcePoolCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterResourcePoolCreateSpec(val *VcenterResourcePoolCreateSpec) *NullableVcenterResourcePoolCreateSpec {
	return &NullableVcenterResourcePoolCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterResourcePoolCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterResourcePoolCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


