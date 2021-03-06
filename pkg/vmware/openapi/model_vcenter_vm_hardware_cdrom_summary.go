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

// VcenterVmHardwareCdromSummary struct for VcenterVmHardwareCdromSummary
type VcenterVmHardwareCdromSummary struct {
	// Identifier of the virtual CD-ROM device. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Cdrom. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Cdrom.
	Cdrom string `json:"cdrom"`
}

// NewVcenterVmHardwareCdromSummary instantiates a new VcenterVmHardwareCdromSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareCdromSummary(cdrom string) *VcenterVmHardwareCdromSummary {
	this := VcenterVmHardwareCdromSummary{}
	this.Cdrom = cdrom
	return &this
}

// NewVcenterVmHardwareCdromSummaryWithDefaults instantiates a new VcenterVmHardwareCdromSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareCdromSummaryWithDefaults() *VcenterVmHardwareCdromSummary {
	this := VcenterVmHardwareCdromSummary{}
	return &this
}

// GetCdrom returns the Cdrom field value
func (o *VcenterVmHardwareCdromSummary) GetCdrom() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cdrom
}

// GetCdromOk returns a tuple with the Cdrom field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareCdromSummary) GetCdromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cdrom, true
}

// SetCdrom sets field value
func (o *VcenterVmHardwareCdromSummary) SetCdrom(v string) {
	o.Cdrom = v
}

func (o VcenterVmHardwareCdromSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cdrom"] = o.Cdrom
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareCdromSummary struct {
	value *VcenterVmHardwareCdromSummary
	isSet bool
}

func (v NullableVcenterVmHardwareCdromSummary) Get() *VcenterVmHardwareCdromSummary {
	return v.value
}

func (v *NullableVcenterVmHardwareCdromSummary) Set(val *VcenterVmHardwareCdromSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareCdromSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareCdromSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareCdromSummary(val *VcenterVmHardwareCdromSummary) *NullableVcenterVmHardwareCdromSummary {
	return &NullableVcenterVmHardwareCdromSummary{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareCdromSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareCdromSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


