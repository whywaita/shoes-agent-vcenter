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

// VcenterVchaClusterPassiveCheckVchaClusterPassive struct for VcenterVchaClusterPassiveCheckVchaClusterPassive
type VcenterVchaClusterPassiveCheckVchaClusterPassive struct {
	Spec *VcenterVchaClusterPassiveCheckSpec `json:"spec,omitempty"`
}

// NewVcenterVchaClusterPassiveCheckVchaClusterPassive instantiates a new VcenterVchaClusterPassiveCheckVchaClusterPassive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterPassiveCheckVchaClusterPassive() *VcenterVchaClusterPassiveCheckVchaClusterPassive {
	this := VcenterVchaClusterPassiveCheckVchaClusterPassive{}
	return &this
}

// NewVcenterVchaClusterPassiveCheckVchaClusterPassiveWithDefaults instantiates a new VcenterVchaClusterPassiveCheckVchaClusterPassive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterPassiveCheckVchaClusterPassiveWithDefaults() *VcenterVchaClusterPassiveCheckVchaClusterPassive {
	this := VcenterVchaClusterPassiveCheckVchaClusterPassive{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVchaClusterPassiveCheckVchaClusterPassive) GetSpec() VcenterVchaClusterPassiveCheckSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVchaClusterPassiveCheckSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterPassiveCheckVchaClusterPassive) GetSpecOk() (*VcenterVchaClusterPassiveCheckSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVchaClusterPassiveCheckVchaClusterPassive) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVchaClusterPassiveCheckSpec and assigns it to the Spec field.
func (o *VcenterVchaClusterPassiveCheckVchaClusterPassive) SetSpec(v VcenterVchaClusterPassiveCheckSpec) {
	o.Spec = &v
}

func (o VcenterVchaClusterPassiveCheckVchaClusterPassive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterPassiveCheckVchaClusterPassive struct {
	value *VcenterVchaClusterPassiveCheckVchaClusterPassive
	isSet bool
}

func (v NullableVcenterVchaClusterPassiveCheckVchaClusterPassive) Get() *VcenterVchaClusterPassiveCheckVchaClusterPassive {
	return v.value
}

func (v *NullableVcenterVchaClusterPassiveCheckVchaClusterPassive) Set(val *VcenterVchaClusterPassiveCheckVchaClusterPassive) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterPassiveCheckVchaClusterPassive) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterPassiveCheckVchaClusterPassive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterPassiveCheckVchaClusterPassive(val *VcenterVchaClusterPassiveCheckVchaClusterPassive) *NullableVcenterVchaClusterPassiveCheckVchaClusterPassive {
	return &NullableVcenterVchaClusterPassiveCheckVchaClusterPassive{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterPassiveCheckVchaClusterPassive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterPassiveCheckVchaClusterPassive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

