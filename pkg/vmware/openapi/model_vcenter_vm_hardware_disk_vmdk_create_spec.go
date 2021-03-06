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

// VcenterVmHardwareDiskVmdkCreateSpec struct for VcenterVmHardwareDiskVmdkCreateSpec
type VcenterVmHardwareDiskVmdkCreateSpec struct {
	// Base name of the VMDK file. The name should not include the '.vmdk' file extension. If unset, a name (derived from the name of the virtual machine) will be chosen by the server.
	Name *string `json:"name,omitempty"`
	// Capacity of the virtual disk backing in bytes. If unset, defaults to a guest-specific capacity.
	Capacity *int64 `json:"capacity,omitempty"`
	StoragePolicy *VcenterVmHardwareDiskStoragePolicySpec `json:"storage_policy,omitempty"`
}

// NewVcenterVmHardwareDiskVmdkCreateSpec instantiates a new VcenterVmHardwareDiskVmdkCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareDiskVmdkCreateSpec() *VcenterVmHardwareDiskVmdkCreateSpec {
	this := VcenterVmHardwareDiskVmdkCreateSpec{}
	return &this
}

// NewVcenterVmHardwareDiskVmdkCreateSpecWithDefaults instantiates a new VcenterVmHardwareDiskVmdkCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareDiskVmdkCreateSpecWithDefaults() *VcenterVmHardwareDiskVmdkCreateSpec {
	this := VcenterVmHardwareDiskVmdkCreateSpec{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) SetName(v string) {
	o.Name = &v
}

// GetCapacity returns the Capacity field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetCapacity() int64 {
	if o == nil || o.Capacity == nil {
		var ret int64
		return ret
	}
	return *o.Capacity
}

// GetCapacityOk returns a tuple with the Capacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetCapacityOk() (*int64, bool) {
	if o == nil || o.Capacity == nil {
		return nil, false
	}
	return o.Capacity, true
}

// HasCapacity returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) HasCapacity() bool {
	if o != nil && o.Capacity != nil {
		return true
	}

	return false
}

// SetCapacity gets a reference to the given int64 and assigns it to the Capacity field.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) SetCapacity(v int64) {
	o.Capacity = &v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetStoragePolicy() VcenterVmHardwareDiskStoragePolicySpec {
	if o == nil || o.StoragePolicy == nil {
		var ret VcenterVmHardwareDiskStoragePolicySpec
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) GetStoragePolicyOk() (*VcenterVmHardwareDiskStoragePolicySpec, bool) {
	if o == nil || o.StoragePolicy == nil {
		return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy != nil {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given VcenterVmHardwareDiskStoragePolicySpec and assigns it to the StoragePolicy field.
func (o *VcenterVmHardwareDiskVmdkCreateSpec) SetStoragePolicy(v VcenterVmHardwareDiskStoragePolicySpec) {
	o.StoragePolicy = &v
}

func (o VcenterVmHardwareDiskVmdkCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Capacity != nil {
		toSerialize["capacity"] = o.Capacity
	}
	if o.StoragePolicy != nil {
		toSerialize["storage_policy"] = o.StoragePolicy
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareDiskVmdkCreateSpec struct {
	value *VcenterVmHardwareDiskVmdkCreateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareDiskVmdkCreateSpec) Get() *VcenterVmHardwareDiskVmdkCreateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareDiskVmdkCreateSpec) Set(val *VcenterVmHardwareDiskVmdkCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareDiskVmdkCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareDiskVmdkCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareDiskVmdkCreateSpec(val *VcenterVmHardwareDiskVmdkCreateSpec) *NullableVcenterVmHardwareDiskVmdkCreateSpec {
	return &NullableVcenterVmHardwareDiskVmdkCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareDiskVmdkCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareDiskVmdkCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


