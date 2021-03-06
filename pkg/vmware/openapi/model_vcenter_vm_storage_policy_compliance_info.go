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

// VcenterVmStoragePolicyComplianceInfo struct for VcenterVmStoragePolicyComplianceInfo
type VcenterVmStoragePolicyComplianceInfo struct {
	OverallCompliance VcenterVmStoragePolicyComplianceStatus `json:"overall_compliance"`
	VmHome *VcenterVmStoragePolicyComplianceVmComplianceInfo `json:"vm_home,omitempty"`
	// The compliance information Compliance.VmComplianceInfo for the virtual machine's virtual disks that are currently associated with a storage policy. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk.
	Disks []VcenterVmStoragePolicyComplianceInfoDisks `json:"disks"`
}

// NewVcenterVmStoragePolicyComplianceInfo instantiates a new VcenterVmStoragePolicyComplianceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmStoragePolicyComplianceInfo(overallCompliance VcenterVmStoragePolicyComplianceStatus, disks []VcenterVmStoragePolicyComplianceInfoDisks) *VcenterVmStoragePolicyComplianceInfo {
	this := VcenterVmStoragePolicyComplianceInfo{}
	this.OverallCompliance = overallCompliance
	this.Disks = disks
	return &this
}

// NewVcenterVmStoragePolicyComplianceInfoWithDefaults instantiates a new VcenterVmStoragePolicyComplianceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmStoragePolicyComplianceInfoWithDefaults() *VcenterVmStoragePolicyComplianceInfo {
	this := VcenterVmStoragePolicyComplianceInfo{}
	return &this
}

// GetOverallCompliance returns the OverallCompliance field value
func (o *VcenterVmStoragePolicyComplianceInfo) GetOverallCompliance() VcenterVmStoragePolicyComplianceStatus {
	if o == nil {
		var ret VcenterVmStoragePolicyComplianceStatus
		return ret
	}

	return o.OverallCompliance
}

// GetOverallComplianceOk returns a tuple with the OverallCompliance field value
// and a boolean to check if the value has been set.
func (o *VcenterVmStoragePolicyComplianceInfo) GetOverallComplianceOk() (*VcenterVmStoragePolicyComplianceStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OverallCompliance, true
}

// SetOverallCompliance sets field value
func (o *VcenterVmStoragePolicyComplianceInfo) SetOverallCompliance(v VcenterVmStoragePolicyComplianceStatus) {
	o.OverallCompliance = v
}

// GetVmHome returns the VmHome field value if set, zero value otherwise.
func (o *VcenterVmStoragePolicyComplianceInfo) GetVmHome() VcenterVmStoragePolicyComplianceVmComplianceInfo {
	if o == nil || o.VmHome == nil {
		var ret VcenterVmStoragePolicyComplianceVmComplianceInfo
		return ret
	}
	return *o.VmHome
}

// GetVmHomeOk returns a tuple with the VmHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmStoragePolicyComplianceInfo) GetVmHomeOk() (*VcenterVmStoragePolicyComplianceVmComplianceInfo, bool) {
	if o == nil || o.VmHome == nil {
		return nil, false
	}
	return o.VmHome, true
}

// HasVmHome returns a boolean if a field has been set.
func (o *VcenterVmStoragePolicyComplianceInfo) HasVmHome() bool {
	if o != nil && o.VmHome != nil {
		return true
	}

	return false
}

// SetVmHome gets a reference to the given VcenterVmStoragePolicyComplianceVmComplianceInfo and assigns it to the VmHome field.
func (o *VcenterVmStoragePolicyComplianceInfo) SetVmHome(v VcenterVmStoragePolicyComplianceVmComplianceInfo) {
	o.VmHome = &v
}

// GetDisks returns the Disks field value
func (o *VcenterVmStoragePolicyComplianceInfo) GetDisks() []VcenterVmStoragePolicyComplianceInfoDisks {
	if o == nil {
		var ret []VcenterVmStoragePolicyComplianceInfoDisks
		return ret
	}

	return o.Disks
}

// GetDisksOk returns a tuple with the Disks field value
// and a boolean to check if the value has been set.
func (o *VcenterVmStoragePolicyComplianceInfo) GetDisksOk() (*[]VcenterVmStoragePolicyComplianceInfoDisks, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Disks, true
}

// SetDisks sets field value
func (o *VcenterVmStoragePolicyComplianceInfo) SetDisks(v []VcenterVmStoragePolicyComplianceInfoDisks) {
	o.Disks = v
}

func (o VcenterVmStoragePolicyComplianceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["overall_compliance"] = o.OverallCompliance
	}
	if o.VmHome != nil {
		toSerialize["vm_home"] = o.VmHome
	}
	if true {
		toSerialize["disks"] = o.Disks
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmStoragePolicyComplianceInfo struct {
	value *VcenterVmStoragePolicyComplianceInfo
	isSet bool
}

func (v NullableVcenterVmStoragePolicyComplianceInfo) Get() *VcenterVmStoragePolicyComplianceInfo {
	return v.value
}

func (v *NullableVcenterVmStoragePolicyComplianceInfo) Set(val *VcenterVmStoragePolicyComplianceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmStoragePolicyComplianceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmStoragePolicyComplianceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmStoragePolicyComplianceInfo(val *VcenterVmStoragePolicyComplianceInfo) *NullableVcenterVmStoragePolicyComplianceInfo {
	return &NullableVcenterVmStoragePolicyComplianceInfo{value: val, isSet: true}
}

func (v NullableVcenterVmStoragePolicyComplianceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmStoragePolicyComplianceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


