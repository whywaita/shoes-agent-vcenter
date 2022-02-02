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

// VcenterStoragePoliciesComplianceSummary struct for VcenterStoragePoliciesComplianceSummary
type VcenterStoragePoliciesComplianceSummary struct {
	// Identifier of virtual machine When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine.
	Vm string `json:"vm"`
	VmHome *VcenterStoragePoliciesComplianceStatus `json:"vm_home,omitempty"`
	// List of the virtual hard disk. If unset or empty, virtual machine entity does not have any disks or its disks are not associated with a storage policy. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk.
	Disks *[]VcenterStoragePoliciesComplianceSummaryDisks `json:"disks,omitempty"`
}

// NewVcenterStoragePoliciesComplianceSummary instantiates a new VcenterStoragePoliciesComplianceSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterStoragePoliciesComplianceSummary(vm string) *VcenterStoragePoliciesComplianceSummary {
	this := VcenterStoragePoliciesComplianceSummary{}
	this.Vm = vm
	return &this
}

// NewVcenterStoragePoliciesComplianceSummaryWithDefaults instantiates a new VcenterStoragePoliciesComplianceSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterStoragePoliciesComplianceSummaryWithDefaults() *VcenterStoragePoliciesComplianceSummary {
	this := VcenterStoragePoliciesComplianceSummary{}
	return &this
}

// GetVm returns the Vm field value
func (o *VcenterStoragePoliciesComplianceSummary) GetVm() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Vm
}

// GetVmOk returns a tuple with the Vm field value
// and a boolean to check if the value has been set.
func (o *VcenterStoragePoliciesComplianceSummary) GetVmOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Vm, true
}

// SetVm sets field value
func (o *VcenterStoragePoliciesComplianceSummary) SetVm(v string) {
	o.Vm = v
}

// GetVmHome returns the VmHome field value if set, zero value otherwise.
func (o *VcenterStoragePoliciesComplianceSummary) GetVmHome() VcenterStoragePoliciesComplianceStatus {
	if o == nil || o.VmHome == nil {
		var ret VcenterStoragePoliciesComplianceStatus
		return ret
	}
	return *o.VmHome
}

// GetVmHomeOk returns a tuple with the VmHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterStoragePoliciesComplianceSummary) GetVmHomeOk() (*VcenterStoragePoliciesComplianceStatus, bool) {
	if o == nil || o.VmHome == nil {
		return nil, false
	}
	return o.VmHome, true
}

// HasVmHome returns a boolean if a field has been set.
func (o *VcenterStoragePoliciesComplianceSummary) HasVmHome() bool {
	if o != nil && o.VmHome != nil {
		return true
	}

	return false
}

// SetVmHome gets a reference to the given VcenterStoragePoliciesComplianceStatus and assigns it to the VmHome field.
func (o *VcenterStoragePoliciesComplianceSummary) SetVmHome(v VcenterStoragePoliciesComplianceStatus) {
	o.VmHome = &v
}

// GetDisks returns the Disks field value if set, zero value otherwise.
func (o *VcenterStoragePoliciesComplianceSummary) GetDisks() []VcenterStoragePoliciesComplianceSummaryDisks {
	if o == nil || o.Disks == nil {
		var ret []VcenterStoragePoliciesComplianceSummaryDisks
		return ret
	}
	return *o.Disks
}

// GetDisksOk returns a tuple with the Disks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterStoragePoliciesComplianceSummary) GetDisksOk() (*[]VcenterStoragePoliciesComplianceSummaryDisks, bool) {
	if o == nil || o.Disks == nil {
		return nil, false
	}
	return o.Disks, true
}

// HasDisks returns a boolean if a field has been set.
func (o *VcenterStoragePoliciesComplianceSummary) HasDisks() bool {
	if o != nil && o.Disks != nil {
		return true
	}

	return false
}

// SetDisks gets a reference to the given []VcenterStoragePoliciesComplianceSummaryDisks and assigns it to the Disks field.
func (o *VcenterStoragePoliciesComplianceSummary) SetDisks(v []VcenterStoragePoliciesComplianceSummaryDisks) {
	o.Disks = &v
}

func (o VcenterStoragePoliciesComplianceSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vm"] = o.Vm
	}
	if o.VmHome != nil {
		toSerialize["vm_home"] = o.VmHome
	}
	if o.Disks != nil {
		toSerialize["disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterStoragePoliciesComplianceSummary struct {
	value *VcenterStoragePoliciesComplianceSummary
	isSet bool
}

func (v NullableVcenterStoragePoliciesComplianceSummary) Get() *VcenterStoragePoliciesComplianceSummary {
	return v.value
}

func (v *NullableVcenterStoragePoliciesComplianceSummary) Set(val *VcenterStoragePoliciesComplianceSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterStoragePoliciesComplianceSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterStoragePoliciesComplianceSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterStoragePoliciesComplianceSummary(val *VcenterStoragePoliciesComplianceSummary) *NullableVcenterStoragePoliciesComplianceSummary {
	return &NullableVcenterStoragePoliciesComplianceSummary{value: val, isSet: true}
}

func (v NullableVcenterStoragePoliciesComplianceSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterStoragePoliciesComplianceSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


