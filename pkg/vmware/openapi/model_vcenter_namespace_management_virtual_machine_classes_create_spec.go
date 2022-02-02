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

// VcenterNamespaceManagementVirtualMachineClassesCreateSpec struct for VcenterNamespaceManagementVirtualMachineClassesCreateSpec
type VcenterNamespaceManagementVirtualMachineClassesCreateSpec struct {
	// Identifier of the virtual machine class. This has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the '-' character allowed anywhere except the first or last character. This name is unique in this vCenter server. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.VirtualMachineClass. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.VirtualMachineClass.
	Id string `json:"id"`
	// The number of CPUs configured for virtual machine of this class.
	CpuCount int64 `json:"cpu_count"`
	// The percentage of total available CPUs reserved for a virtual machine. We multiply this percentage by the minimum frequency amongst all the cluster nodes to get the CPU reservation that is specified to vSphere in MHz. If unset, no CPU reservation is requested for the virtual machine.
	CpuReservation *int64 `json:"cpu_reservation,omitempty"`
	// The amount of memory in MB configured for virtual machine of this class.
	MemoryMB int64 `json:"memory_MB"`
	// The percentage of available memory reserved for a virtual machine of this class. Memory reservation must be set to 100% for VM class with vGPU or Dynamic DirectPath I/O devices. If unset, no memory reservation is requested for virtual machine.
	MemoryReservation *int64 `json:"memory_reservation,omitempty"`
	// Description for the VM class. If unset, no description is added to the VM class.
	Description *string `json:"description,omitempty"`
}

// NewVcenterNamespaceManagementVirtualMachineClassesCreateSpec instantiates a new VcenterNamespaceManagementVirtualMachineClassesCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementVirtualMachineClassesCreateSpec(id string, cpuCount int64, memoryMB int64) *VcenterNamespaceManagementVirtualMachineClassesCreateSpec {
	this := VcenterNamespaceManagementVirtualMachineClassesCreateSpec{}
	this.Id = id
	this.CpuCount = cpuCount
	this.MemoryMB = memoryMB
	return &this
}

// NewVcenterNamespaceManagementVirtualMachineClassesCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementVirtualMachineClassesCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementVirtualMachineClassesCreateSpecWithDefaults() *VcenterNamespaceManagementVirtualMachineClassesCreateSpec {
	this := VcenterNamespaceManagementVirtualMachineClassesCreateSpec{}
	return &this
}

// GetId returns the Id field value
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetId(v string) {
	o.Id = v
}

// GetCpuCount returns the CpuCount field value
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CpuCount
}

// GetCpuCountOk returns a tuple with the CpuCount field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CpuCount, true
}

// SetCpuCount sets field value
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetCpuCount(v int64) {
	o.CpuCount = v
}

// GetCpuReservation returns the CpuReservation field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuReservation() int64 {
	if o == nil || o.CpuReservation == nil {
		var ret int64
		return ret
	}
	return *o.CpuReservation
}

// GetCpuReservationOk returns a tuple with the CpuReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuReservationOk() (*int64, bool) {
	if o == nil || o.CpuReservation == nil {
		return nil, false
	}
	return o.CpuReservation, true
}

// HasCpuReservation returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) HasCpuReservation() bool {
	if o != nil && o.CpuReservation != nil {
		return true
	}

	return false
}

// SetCpuReservation gets a reference to the given int64 and assigns it to the CpuReservation field.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetCpuReservation(v int64) {
	o.CpuReservation = &v
}

// GetMemoryMB returns the MemoryMB field value
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryMB() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.MemoryMB
}

// GetMemoryMBOk returns a tuple with the MemoryMB field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryMBOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.MemoryMB, true
}

// SetMemoryMB sets field value
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetMemoryMB(v int64) {
	o.MemoryMB = v
}

// GetMemoryReservation returns the MemoryReservation field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryReservation() int64 {
	if o == nil || o.MemoryReservation == nil {
		var ret int64
		return ret
	}
	return *o.MemoryReservation
}

// GetMemoryReservationOk returns a tuple with the MemoryReservation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryReservationOk() (*int64, bool) {
	if o == nil || o.MemoryReservation == nil {
		return nil, false
	}
	return o.MemoryReservation, true
}

// HasMemoryReservation returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) HasMemoryReservation() bool {
	if o != nil && o.MemoryReservation != nil {
		return true
	}

	return false
}

// SetMemoryReservation gets a reference to the given int64 and assigns it to the MemoryReservation field.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetMemoryReservation(v int64) {
	o.MemoryReservation = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetDescription(v string) {
	o.Description = &v
}

func (o VcenterNamespaceManagementVirtualMachineClassesCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["cpu_count"] = o.CpuCount
	}
	if o.CpuReservation != nil {
		toSerialize["cpu_reservation"] = o.CpuReservation
	}
	if true {
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

type NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec struct {
	value *VcenterNamespaceManagementVirtualMachineClassesCreateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec) Get() *VcenterNamespaceManagementVirtualMachineClassesCreateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec) Set(val *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec(val *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) *NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec {
	return &NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementVirtualMachineClassesCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


