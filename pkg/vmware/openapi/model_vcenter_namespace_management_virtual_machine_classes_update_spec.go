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

// VcenterNamespaceManagementVirtualMachineClassesUpdateSpec struct for VcenterNamespaceManagementVirtualMachineClassesUpdateSpec
type VcenterNamespaceManagementVirtualMachineClassesUpdateSpec struct {
	// The number of CPUs configured for virtual machine of this class. If unset the current value the will not be modified.
	CpuCount *int64 `json:"cpu_count,omitempty"`
	// The percentage of total available CPUs reserved for a virtual machine. We multiply this percentage by the minimum frequency amongst all the cluster nodes to get the CPU reservation that is specified to vSphere in MHz. If unset, no CPU reservation is requested for the virtual machine.
	CpuReservation *int64 `json:"cpu_reservation,omitempty"`
	// The amount of memory in MB configured for virtual machine of this class. If unset the current value the will not be modified.
	MemoryMB *int64 `json:"memory_MB,omitempty"`
	// The percentage of available memory reserved for a virtual machine of this class. Memory reservation must be set to 100% for VM class with vGPU or Dynamic DirectPath I/O devices. If unset, no memory reservation is requested for virtual machine.
	MemoryReservation *int64 `json:"memory_reservation,omitempty"`
	// Description for the VM class. If unset, description is not updated.
	Description *string `json:"description,omitempty"`
}

// NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpec instantiates a new VcenterNamespaceManagementVirtualMachineClassesUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpec() *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec {
	this := VcenterNamespaceManagementVirtualMachineClassesUpdateSpec{}
	return &this
}

// NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementVirtualMachineClassesUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpecWithDefaults() *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec {
	this := VcenterNamespaceManagementVirtualMachineClassesUpdateSpec{}
	return &this
}

// GetCpuCount returns the CpuCount field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuCount() int64 {
	if o == nil || o.CpuCount == nil {
		var ret int64
		return ret
	}
	return *o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuCountOk() (*int64, bool) {
	if o == nil || o.CpuCount == nil {
		return nil, false
	}
	return o.CpuCount, true
}

// HasCpuCount returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasCpuCount() bool {
	if o != nil && o.CpuCount != nil {
		return true
	}

	return false
}

// SetCpuCount gets a reference to the given int64 and assigns it to the CpuCount field.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetCpuCount(v int64) {
	o.CpuCount = &v
}

// GetCpuReservation returns the CpuReservation field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuReservation() int64 {
	if o == nil || o.CpuReservation == nil {
		var ret int64
		return ret
	}
	return *o.CpuReservation
}

// GetCpuReservationOk returns a tuple with the CpuReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuReservationOk() (*int64, bool) {
	if o == nil || o.CpuReservation == nil {
		return nil, false
	}
	return o.CpuReservation, true
}

// HasCpuReservation returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasCpuReservation() bool {
	if o != nil && o.CpuReservation != nil {
		return true
	}

	return false
}

// SetCpuReservation gets a reference to the given int64 and assigns it to the CpuReservation field.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetCpuReservation(v int64) {
	o.CpuReservation = &v
}

// GetMemoryMB returns the MemoryMB field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryMB() int64 {
	if o == nil || o.MemoryMB == nil {
		var ret int64
		return ret
	}
	return *o.MemoryMB
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryMBOk() (*int64, bool) {
	if o == nil || o.MemoryMB == nil {
		return nil, false
	}
	return o.MemoryMB, true
}

// HasMemoryMB returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasMemoryMB() bool {
	if o != nil && o.MemoryMB != nil {
		return true
	}

	return false
}

// SetMemoryMB gets a reference to the given int64 and assigns it to the MemoryMB field.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetMemoryMB(v int64) {
	o.MemoryMB = &v
}

// GetMemoryReservation returns the MemoryReservation field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryReservation() int64 {
	if o == nil || o.MemoryReservation == nil {
		var ret int64
		return ret
	}
	return *o.MemoryReservation
}

// GetMemoryReservationOk returns a tuple with the MemoryReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryReservationOk() (*int64, bool) {
	if o == nil || o.MemoryReservation == nil {
		return nil, false
	}
	return o.MemoryReservation, true
}

// HasMemoryReservation returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasMemoryReservation() bool {
	if o != nil && o.MemoryReservation != nil {
		return true
	}

	return false
}

// SetMemoryReservation gets a reference to the given int64 and assigns it to the MemoryReservation field.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetMemoryReservation(v int64) {
	o.MemoryReservation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetDescription(v string) {
	o.Description = &v
}

func (o VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CpuCount != nil {
		toSerialize["cpu_count"] = o.CpuCount
	}
	if o.CpuReservation != nil {
		toSerialize["cpu_reservation"] = o.CpuReservation
	}
	if o.MemoryMB != nil {
		toSerialize["memory_MB"] = o.MemoryMB
	}
	if o.MemoryReservation != nil {
		toSerialize["memory_reservation"] = o.MemoryReservation
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec struct {
	value *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec) Get() *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec) Set(val *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec(val *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) *NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec {
	return &NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementVirtualMachineClassesUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


