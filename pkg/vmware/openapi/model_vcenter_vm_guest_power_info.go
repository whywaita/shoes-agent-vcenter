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

// VcenterVmGuestPowerInfo struct for VcenterVmGuestPowerInfo
type VcenterVmGuestPowerInfo struct {
	State VcenterVmGuestPowerState `json:"state"`
	// Flag indicating if the virtual machine is ready to process soft power operations.
	OperationsReady bool `json:"operations_ready"`
}

// NewVcenterVmGuestPowerInfo instantiates a new VcenterVmGuestPowerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestPowerInfo(state VcenterVmGuestPowerState, operationsReady bool) *VcenterVmGuestPowerInfo {
	this := VcenterVmGuestPowerInfo{}
	this.State = state
	this.OperationsReady = operationsReady
	return &this
}

// NewVcenterVmGuestPowerInfoWithDefaults instantiates a new VcenterVmGuestPowerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestPowerInfoWithDefaults() *VcenterVmGuestPowerInfo {
	this := VcenterVmGuestPowerInfo{}
	return &this
}

// GetState returns the State field value
func (o *VcenterVmGuestPowerInfo) GetState() VcenterVmGuestPowerState {
	if o == nil {
		var ret VcenterVmGuestPowerState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestPowerInfo) GetStateOk() (*VcenterVmGuestPowerState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VcenterVmGuestPowerInfo) SetState(v VcenterVmGuestPowerState) {
	o.State = v
}

// GetOperationsReady returns the OperationsReady field value
func (o *VcenterVmGuestPowerInfo) GetOperationsReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.OperationsReady
}

// GetOperationsReadyOk returns a tuple with the OperationsReady field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestPowerInfo) GetOperationsReadyOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OperationsReady, true
}

// SetOperationsReady sets field value
func (o *VcenterVmGuestPowerInfo) SetOperationsReady(v bool) {
	o.OperationsReady = v
}

func (o VcenterVmGuestPowerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["operations_ready"] = o.OperationsReady
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestPowerInfo struct {
	value *VcenterVmGuestPowerInfo
	isSet bool
}

func (v NullableVcenterVmGuestPowerInfo) Get() *VcenterVmGuestPowerInfo {
	return v.value
}

func (v *NullableVcenterVmGuestPowerInfo) Set(val *VcenterVmGuestPowerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestPowerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestPowerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestPowerInfo(val *VcenterVmGuestPowerInfo) *NullableVcenterVmGuestPowerInfo {
	return &NullableVcenterVmGuestPowerInfo{value: val, isSet: true}
}

func (v NullableVcenterVmGuestPowerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestPowerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


