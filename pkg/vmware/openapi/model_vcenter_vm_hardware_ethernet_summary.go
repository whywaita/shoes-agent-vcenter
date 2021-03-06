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

// VcenterVmHardwareEthernetSummary struct for VcenterVmHardwareEthernetSummary
type VcenterVmHardwareEthernetSummary struct {
	// Identifier of the virtual Ethernet adapter. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Ethernet. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Ethernet.
	Nic string `json:"nic"`
}

// NewVcenterVmHardwareEthernetSummary instantiates a new VcenterVmHardwareEthernetSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareEthernetSummary(nic string) *VcenterVmHardwareEthernetSummary {
	this := VcenterVmHardwareEthernetSummary{}
	this.Nic = nic
	return &this
}

// NewVcenterVmHardwareEthernetSummaryWithDefaults instantiates a new VcenterVmHardwareEthernetSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareEthernetSummaryWithDefaults() *VcenterVmHardwareEthernetSummary {
	this := VcenterVmHardwareEthernetSummary{}
	return &this
}

// GetNic returns the Nic field value
func (o *VcenterVmHardwareEthernetSummary) GetNic() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Nic
}

// GetNicOk returns a tuple with the Nic field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareEthernetSummary) GetNicOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Nic, true
}

// SetNic sets field value
func (o *VcenterVmHardwareEthernetSummary) SetNic(v string) {
	o.Nic = v
}

func (o VcenterVmHardwareEthernetSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["nic"] = o.Nic
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareEthernetSummary struct {
	value *VcenterVmHardwareEthernetSummary
	isSet bool
}

func (v NullableVcenterVmHardwareEthernetSummary) Get() *VcenterVmHardwareEthernetSummary {
	return v.value
}

func (v *NullableVcenterVmHardwareEthernetSummary) Set(val *VcenterVmHardwareEthernetSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareEthernetSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareEthernetSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareEthernetSummary(val *VcenterVmHardwareEthernetSummary) *NullableVcenterVmHardwareEthernetSummary {
	return &NullableVcenterVmHardwareEthernetSummary{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareEthernetSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareEthernetSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


