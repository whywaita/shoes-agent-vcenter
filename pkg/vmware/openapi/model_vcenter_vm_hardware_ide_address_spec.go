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

// VcenterVmHardwareIdeAddressSpec struct for VcenterVmHardwareIdeAddressSpec
type VcenterVmHardwareIdeAddressSpec struct {
	// Flag specifying whether the device should be attached to the primary or secondary IDE adapter of the virtual machine. If unset, the server will choose a adapter with an available connection. If no IDE connections are available, the request will be rejected.
	Primary *bool `json:"primary,omitempty"`
	// Flag specifying whether the device should be the master or slave device on the IDE adapter. If unset, the server will choose an available connection type. If no IDE connections are available, the request will be rejected.
	Master *bool `json:"master,omitempty"`
}

// NewVcenterVmHardwareIdeAddressSpec instantiates a new VcenterVmHardwareIdeAddressSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareIdeAddressSpec() *VcenterVmHardwareIdeAddressSpec {
	this := VcenterVmHardwareIdeAddressSpec{}
	return &this
}

// NewVcenterVmHardwareIdeAddressSpecWithDefaults instantiates a new VcenterVmHardwareIdeAddressSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareIdeAddressSpecWithDefaults() *VcenterVmHardwareIdeAddressSpec {
	this := VcenterVmHardwareIdeAddressSpec{}
	return &this
}

// GetPrimary returns the Primary field value if set, zero value otherwise.
func (o *VcenterVmHardwareIdeAddressSpec) GetPrimary() bool {
	if o == nil || o.Primary == nil {
		var ret bool
		return ret
	}
	return *o.Primary
}

// GetPrimaryOk returns a tuple with the Primary field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareIdeAddressSpec) GetPrimaryOk() (*bool, bool) {
	if o == nil || o.Primary == nil {
		return nil, false
	}
	return o.Primary, true
}

// HasPrimary returns a boolean if a field has been set.
func (o *VcenterVmHardwareIdeAddressSpec) HasPrimary() bool {
	if o != nil && o.Primary != nil {
		return true
	}

	return false
}

// SetPrimary gets a reference to the given bool and assigns it to the Primary field.
func (o *VcenterVmHardwareIdeAddressSpec) SetPrimary(v bool) {
	o.Primary = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *VcenterVmHardwareIdeAddressSpec) GetMaster() bool {
	if o == nil || o.Master == nil {
		var ret bool
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareIdeAddressSpec) GetMasterOk() (*bool, bool) {
	if o == nil || o.Master == nil {
		return nil, false
	}
	return o.Master, true
}

// HasMaster returns a boolean if a field has been set.
func (o *VcenterVmHardwareIdeAddressSpec) HasMaster() bool {
	if o != nil && o.Master != nil {
		return true
	}

	return false
}

// SetMaster gets a reference to the given bool and assigns it to the Master field.
func (o *VcenterVmHardwareIdeAddressSpec) SetMaster(v bool) {
	o.Master = &v
}

func (o VcenterVmHardwareIdeAddressSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Primary != nil {
		toSerialize["primary"] = o.Primary
	}
	if o.Master != nil {
		toSerialize["master"] = o.Master
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareIdeAddressSpec struct {
	value *VcenterVmHardwareIdeAddressSpec
	isSet bool
}

func (v NullableVcenterVmHardwareIdeAddressSpec) Get() *VcenterVmHardwareIdeAddressSpec {
	return v.value
}

func (v *NullableVcenterVmHardwareIdeAddressSpec) Set(val *VcenterVmHardwareIdeAddressSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareIdeAddressSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareIdeAddressSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareIdeAddressSpec(val *VcenterVmHardwareIdeAddressSpec) *NullableVcenterVmHardwareIdeAddressSpec {
	return &NullableVcenterVmHardwareIdeAddressSpec{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareIdeAddressSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareIdeAddressSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


