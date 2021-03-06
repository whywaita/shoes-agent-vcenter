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

// VcenterVmHardwareSerialCreateVmHardwareSerial struct for VcenterVmHardwareSerialCreateVmHardwareSerial
type VcenterVmHardwareSerialCreateVmHardwareSerial struct {
	Spec *VcenterVmHardwareSerialCreateSpec `json:"spec,omitempty"`
}

// NewVcenterVmHardwareSerialCreateVmHardwareSerial instantiates a new VcenterVmHardwareSerialCreateVmHardwareSerial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareSerialCreateVmHardwareSerial() *VcenterVmHardwareSerialCreateVmHardwareSerial {
	this := VcenterVmHardwareSerialCreateVmHardwareSerial{}
	return &this
}

// NewVcenterVmHardwareSerialCreateVmHardwareSerialWithDefaults instantiates a new VcenterVmHardwareSerialCreateVmHardwareSerial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareSerialCreateVmHardwareSerialWithDefaults() *VcenterVmHardwareSerialCreateVmHardwareSerial {
	this := VcenterVmHardwareSerialCreateVmHardwareSerial{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialCreateVmHardwareSerial) GetSpec() VcenterVmHardwareSerialCreateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVmHardwareSerialCreateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialCreateVmHardwareSerial) GetSpecOk() (*VcenterVmHardwareSerialCreateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialCreateVmHardwareSerial) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVmHardwareSerialCreateSpec and assigns it to the Spec field.
func (o *VcenterVmHardwareSerialCreateVmHardwareSerial) SetSpec(v VcenterVmHardwareSerialCreateSpec) {
	o.Spec = &v
}

func (o VcenterVmHardwareSerialCreateVmHardwareSerial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareSerialCreateVmHardwareSerial struct {
	value *VcenterVmHardwareSerialCreateVmHardwareSerial
	isSet bool
}

func (v NullableVcenterVmHardwareSerialCreateVmHardwareSerial) Get() *VcenterVmHardwareSerialCreateVmHardwareSerial {
	return v.value
}

func (v *NullableVcenterVmHardwareSerialCreateVmHardwareSerial) Set(val *VcenterVmHardwareSerialCreateVmHardwareSerial) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareSerialCreateVmHardwareSerial) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareSerialCreateVmHardwareSerial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareSerialCreateVmHardwareSerial(val *VcenterVmHardwareSerialCreateVmHardwareSerial) *NullableVcenterVmHardwareSerialCreateVmHardwareSerial {
	return &NullableVcenterVmHardwareSerialCreateVmHardwareSerial{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareSerialCreateVmHardwareSerial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareSerialCreateVmHardwareSerial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


