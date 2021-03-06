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

// VcenterVmHardwareSerialUpdateVmHardwareSerial struct for VcenterVmHardwareSerialUpdateVmHardwareSerial
type VcenterVmHardwareSerialUpdateVmHardwareSerial struct {
	Spec *VcenterVmHardwareSerialUpdateSpec `json:"spec,omitempty"`
}

// NewVcenterVmHardwareSerialUpdateVmHardwareSerial instantiates a new VcenterVmHardwareSerialUpdateVmHardwareSerial object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmHardwareSerialUpdateVmHardwareSerial() *VcenterVmHardwareSerialUpdateVmHardwareSerial {
	this := VcenterVmHardwareSerialUpdateVmHardwareSerial{}
	return &this
}

// NewVcenterVmHardwareSerialUpdateVmHardwareSerialWithDefaults instantiates a new VcenterVmHardwareSerialUpdateVmHardwareSerial object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmHardwareSerialUpdateVmHardwareSerialWithDefaults() *VcenterVmHardwareSerialUpdateVmHardwareSerial {
	this := VcenterVmHardwareSerialUpdateVmHardwareSerial{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVmHardwareSerialUpdateVmHardwareSerial) GetSpec() VcenterVmHardwareSerialUpdateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVmHardwareSerialUpdateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmHardwareSerialUpdateVmHardwareSerial) GetSpecOk() (*VcenterVmHardwareSerialUpdateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVmHardwareSerialUpdateVmHardwareSerial) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVmHardwareSerialUpdateSpec and assigns it to the Spec field.
func (o *VcenterVmHardwareSerialUpdateVmHardwareSerial) SetSpec(v VcenterVmHardwareSerialUpdateSpec) {
	o.Spec = &v
}

func (o VcenterVmHardwareSerialUpdateVmHardwareSerial) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmHardwareSerialUpdateVmHardwareSerial struct {
	value *VcenterVmHardwareSerialUpdateVmHardwareSerial
	isSet bool
}

func (v NullableVcenterVmHardwareSerialUpdateVmHardwareSerial) Get() *VcenterVmHardwareSerialUpdateVmHardwareSerial {
	return v.value
}

func (v *NullableVcenterVmHardwareSerialUpdateVmHardwareSerial) Set(val *VcenterVmHardwareSerialUpdateVmHardwareSerial) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmHardwareSerialUpdateVmHardwareSerial) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmHardwareSerialUpdateVmHardwareSerial) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmHardwareSerialUpdateVmHardwareSerial(val *VcenterVmHardwareSerialUpdateVmHardwareSerial) *NullableVcenterVmHardwareSerialUpdateVmHardwareSerial {
	return &NullableVcenterVmHardwareSerialUpdateVmHardwareSerial{value: val, isSet: true}
}

func (v NullableVcenterVmHardwareSerialUpdateVmHardwareSerial) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmHardwareSerialUpdateVmHardwareSerial) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


