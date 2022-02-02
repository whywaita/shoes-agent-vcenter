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

// VcenterVmHardwareSerialSummary struct for VcenterVmHardwareSerialSummary
type VcenterVmHardwareSerialSummary struct {
	// Identifier of the virtual serial port. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.SerialPort. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.SerialPort.
	Port string `json:"port"`
}

// NewVcenterVmHardwareSerialSummary instantiates a new VcenterVmHardwareSerialSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareSerialSummary(port string) *VcenterVmHardwareSerialSummary {
	this := VcenterVmHardwareSerialSummary{}
	this.Port = port
	return &this
}

// NewVcenterVmHardwareSerialSummaryWithDefaults instantiates a new VcenterVmHardwareSerialSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareSerialSummaryWithDefaults() *VcenterVmHardwareSerialSummary {
	this := VcenterVmHardwareSerialSummary{}
	return &this
}

// GetPort returns the Port field value
func (o *VcenterVmHardwareSerialSummary) GetPort() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialSummary) GetPortOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *VcenterVmHardwareSerialSummary) SetPort(v string) {
	o.Port = v
}

func (o VcenterVmHardwareSerialSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareSerialSummary struct {
	value *VcenterVmHardwareSerialSummary
	isSet bool
}

func (v NullableVcenterVmHardwareSerialSummary) Get() *VcenterVmHardwareSerialSummary {
	return v.value
}

func (v *NullableVcenterVmHardwareSerialSummary) Set(val *VcenterVmHardwareSerialSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareSerialSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareSerialSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareSerialSummary(val *VcenterVmHardwareSerialSummary) *NullableVcenterVmHardwareSerialSummary {
	return &NullableVcenterVmHardwareSerialSummary{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareSerialSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareSerialSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


