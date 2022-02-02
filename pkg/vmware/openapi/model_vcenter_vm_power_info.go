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

// VcenterVmPowerInfo struct for VcenterVmPowerInfo
type VcenterVmPowerInfo struct {
	State VcenterVmPowerState `json:"state"`
	// Flag indicating whether the virtual machine was powered off cleanly. This field may be used to detect that the virtual machine crashed unexpectedly and should be restarted. This field is optional and it is only relevant when the value of Power.Info.state is POWERED_OFF.
	CleanPowerOff *bool `json:"clean_power_off,omitempty"`
}

// NewVcenterVmPowerInfo instantiates a new VcenterVmPowerInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmPowerInfo(state VcenterVmPowerState) *VcenterVmPowerInfo {
	this := VcenterVmPowerInfo{}
	this.State = state
	return &this
}

// NewVcenterVmPowerInfoWithDefaults instantiates a new VcenterVmPowerInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmPowerInfoWithDefaults() *VcenterVmPowerInfo {
	this := VcenterVmPowerInfo{}
	return &this
}

// GetState returns the State field value
func (o *VcenterVmPowerInfo) GetState() VcenterVmPowerState {
	if o == nil {
		var ret VcenterVmPowerState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VcenterVmPowerInfo) GetStateOk() (*VcenterVmPowerState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VcenterVmPowerInfo) SetState(v VcenterVmPowerState) {
	o.State = v
}

// GetCleanPowerOff returns the CleanPowerOff field value if set, zero value otherwise.
func (o *VcenterVmPowerInfo) GetCleanPowerOff() bool {
	if o == nil || o.CleanPowerOff == nil {
		var ret bool
		return ret
	}
	return *o.CleanPowerOff
}

// GetCleanPowerOffOk returns a tuple with the CleanPowerOff field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmPowerInfo) GetCleanPowerOffOk() (*bool, bool) {
	if o == nil || o.CleanPowerOff == nil {
		return nil, false
	}
	return o.CleanPowerOff, true
}

// HasCleanPowerOff returns a boolean if a field has been set.
func (o *VcenterVmPowerInfo) HasCleanPowerOff() bool {
	if o != nil && o.CleanPowerOff != nil {
		return true
	}

	return false
}

// SetCleanPowerOff gets a reference to the given bool and assigns it to the CleanPowerOff field.
func (o *VcenterVmPowerInfo) SetCleanPowerOff(v bool) {
	o.CleanPowerOff = &v
}

func (o VcenterVmPowerInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["state"] = o.State
	}
	if o.CleanPowerOff != nil {
		toSerialize["clean_power_off"] = o.CleanPowerOff
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmPowerInfo struct {
	value *VcenterVmPowerInfo
	isSet bool
}

func (v NullableVcenterVmPowerInfo) Get() *VcenterVmPowerInfo {
	return v.value
}

func (v *NullableVcenterVmPowerInfo) Set(val *VcenterVmPowerInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmPowerInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmPowerInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmPowerInfo(val *VcenterVmPowerInfo) *NullableVcenterVmPowerInfo {
	return &NullableVcenterVmPowerInfo{value: val, isSet: true}
}

func (v NullableVcenterVmPowerInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmPowerInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


