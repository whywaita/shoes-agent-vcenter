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

// VcenterVmHardwareParallelUpdateSpec struct for VcenterVmHardwareParallelUpdateSpec
type VcenterVmHardwareParallelUpdateSpec struct {
	Backing *VcenterVmHardwareParallelBackingSpec `json:"backing,omitempty"`
	// Flag indicating whether the virtual device should be connected whenever the virtual machine is powered on. If unset, the value is unchanged.
	StartConnected *bool `json:"start_connected,omitempty"`
	// Flag indicating whether the guest can connect and disconnect the device. If unset, the value is unchanged.
	AllowGuestControl *bool `json:"allow_guest_control,omitempty"`
}

// NewVcenterVmHardwareParallelUpdateSpec instantiates a new VcenterVmHardwareParallelUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareParallelUpdateSpec() *VcenterVmHardwareParallelUpdateSpec {
	this := VcenterVmHardwareParallelUpdateSpec{}
	return &this
}

// NewVcenterVmHardwareParallelUpdateSpecWithDefaults instantiates a new VcenterVmHardwareParallelUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareParallelUpdateSpecWithDefaults() *VcenterVmHardwareParallelUpdateSpec {
	this := VcenterVmHardwareParallelUpdateSpec{}
	return &this
}

// GetBacking returns the Backing field value if set, zero value otherwise.
func (o *VcenterVmHardwareParallelUpdateSpec) GetBacking() VcenterVmHardwareParallelBackingSpec {
	if o == nil || o.Backing == nil {
		var ret VcenterVmHardwareParallelBackingSpec
		return ret
	}
	return *o.Backing
}

// GetBackingOk returns a tuple with the Backing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareParallelUpdateSpec) GetBackingOk() (*VcenterVmHardwareParallelBackingSpec, bool) {
	if o == nil || o.Backing == nil {
		return nil, false
	}
	return o.Backing, true
}

// HasBacking returns a boolean if a field has been set.
func (o *VcenterVmHardwareParallelUpdateSpec) HasBacking() bool {
	if o != nil && o.Backing != nil {
		return true
	}

	return false
}

// SetBacking gets a reference to the given VcenterVmHardwareParallelBackingSpec and assigns it to the Backing field.
func (o *VcenterVmHardwareParallelUpdateSpec) SetBacking(v VcenterVmHardwareParallelBackingSpec) {
	o.Backing = &v
}

// GetStartConnected returns the StartConnected field value if set, zero value otherwise.
func (o *VcenterVmHardwareParallelUpdateSpec) GetStartConnected() bool {
	if o == nil || o.StartConnected == nil {
		var ret bool
		return ret
	}
	return *o.StartConnected
}

// GetStartConnectedOk returns a tuple with the StartConnected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareParallelUpdateSpec) GetStartConnectedOk() (*bool, bool) {
	if o == nil || o.StartConnected == nil {
		return nil, false
	}
	return o.StartConnected, true
}

// HasStartConnected returns a boolean if a field has been set.
func (o *VcenterVmHardwareParallelUpdateSpec) HasStartConnected() bool {
	if o != nil && o.StartConnected != nil {
		return true
	}

	return false
}

// SetStartConnected gets a reference to the given bool and assigns it to the StartConnected field.
func (o *VcenterVmHardwareParallelUpdateSpec) SetStartConnected(v bool) {
	o.StartConnected = &v
}

// GetAllowGuestControl returns the AllowGuestControl field value if set, zero value otherwise.
func (o *VcenterVmHardwareParallelUpdateSpec) GetAllowGuestControl() bool {
	if o == nil || o.AllowGuestControl == nil {
		var ret bool
		return ret
	}
	return *o.AllowGuestControl
}

// GetAllowGuestControlOk returns a tuple with the AllowGuestControl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareParallelUpdateSpec) GetAllowGuestControlOk() (*bool, bool) {
	if o == nil || o.AllowGuestControl == nil {
		return nil, false
	}
	return o.AllowGuestControl, true
}

// HasAllowGuestControl returns a boolean if a field has been set.
func (o *VcenterVmHardwareParallelUpdateSpec) HasAllowGuestControl() bool {
	if o != nil && o.AllowGuestControl != nil {
		return true
	}

	return false
}

// SetAllowGuestControl gets a reference to the given bool and assigns it to the AllowGuestControl field.
func (o *VcenterVmHardwareParallelUpdateSpec) SetAllowGuestControl(v bool) {
	o.AllowGuestControl = &v
}

func (o VcenterVmHardwareParallelUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Backing != nil {
		toSerialize["backing"] = o.Backing
	}
	if o.StartConnected != nil {
		toSerialize["start_connected"] = o.StartConnected
	}
	if o.AllowGuestControl != nil {
		toSerialize["allow_guest_control"] = o.AllowGuestControl
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareParallelUpdateSpec struct {
	value *VcenterVmHardwareParallelUpdateSpec
	isSet bool
}

func (v NullableVcenterVmHardwareParallelUpdateSpec) Get() *VcenterVmHardwareParallelUpdateSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareParallelUpdateSpec) Set(val *VcenterVmHardwareParallelUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareParallelUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareParallelUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareParallelUpdateSpec(val *VcenterVmHardwareParallelUpdateSpec) *NullableVcenterVmHardwareParallelUpdateSpec {
	return &NullableVcenterVmHardwareParallelUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareParallelUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareParallelUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


