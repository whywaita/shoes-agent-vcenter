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

// VcenterVmHardwareAdapterSataSummary struct for VcenterVmHardwareAdapterSataSummary
type VcenterVmHardwareAdapterSataSummary struct {
	// Identifier of the virtual SATA adapter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.SataAdapter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.SataAdapter.
	Adapter string `json:"adapter"`
}

// NewVcenterVmHardwareAdapterSataSummary instantiates a new VcenterVmHardwareAdapterSataSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareAdapterSataSummary(adapter string) *VcenterVmHardwareAdapterSataSummary {
	this := VcenterVmHardwareAdapterSataSummary{}
	this.Adapter = adapter
	return &this
}

// NewVcenterVmHardwareAdapterSataSummaryWithDefaults instantiates a new VcenterVmHardwareAdapterSataSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareAdapterSataSummaryWithDefaults() *VcenterVmHardwareAdapterSataSummary {
	this := VcenterVmHardwareAdapterSataSummary{}
	return &this
}

// GetAdapter returns the Adapter field value
func (o *VcenterVmHardwareAdapterSataSummary) GetAdapter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Adapter
}

// GetAdapterOk returns a tuple with the Adapter field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareAdapterSataSummary) GetAdapterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Adapter, true
}

// SetAdapter sets field value
func (o *VcenterVmHardwareAdapterSataSummary) SetAdapter(v string) {
	o.Adapter = v
}

func (o VcenterVmHardwareAdapterSataSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["adapter"] = o.Adapter
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareAdapterSataSummary struct {
	value *VcenterVmHardwareAdapterSataSummary
	isSet bool
}

func (v NullableVcenterVmHardwareAdapterSataSummary) Get() *VcenterVmHardwareAdapterSataSummary {
	return v.value
}

func (v *NullableVcenterVmHardwareAdapterSataSummary) Set(val *VcenterVmHardwareAdapterSataSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareAdapterSataSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareAdapterSataSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareAdapterSataSummary(val *VcenterVmHardwareAdapterSataSummary) *NullableVcenterVmHardwareAdapterSataSummary {
	return &NullableVcenterVmHardwareAdapterSataSummary{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareAdapterSataSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareAdapterSataSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


