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

// VcenterVchaClusterActiveGetVchaClusterActive struct for VcenterVchaClusterActiveGetVchaClusterActive
type VcenterVchaClusterActiveGetVchaClusterActive struct {
	VcSpec *VcenterVchaCredentialsSpec `json:"vc_spec,omitempty"`
	// If true, then return only the information that does not require connecting to the Active vCenter Server.   If false or unset, then return all the information. If unset, then return all the information.
	Partial *bool `json:"partial,omitempty"`
}

// NewVcenterVchaClusterActiveGetVchaClusterActive instantiates a new VcenterVchaClusterActiveGetVchaClusterActive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterActiveGetVchaClusterActive() *VcenterVchaClusterActiveGetVchaClusterActive {
	this := VcenterVchaClusterActiveGetVchaClusterActive{}
	return &this
}

// NewVcenterVchaClusterActiveGetVchaClusterActiveWithDefaults instantiates a new VcenterVchaClusterActiveGetVchaClusterActive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterActiveGetVchaClusterActiveWithDefaults() *VcenterVchaClusterActiveGetVchaClusterActive {
	this := VcenterVchaClusterActiveGetVchaClusterActive{}
	return &this
}

// GetVcSpec returns the VcSpec field value if set, zero value otherwise.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetVcSpec() VcenterVchaCredentialsSpec {
	if o == nil || o.VcSpec == nil {
		var ret VcenterVchaCredentialsSpec
		return ret
	}
	return *o.VcSpec
}

// GetVcSpecOk returns a tuple with the VcSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetVcSpecOk() (*VcenterVchaCredentialsSpec, bool) {
	if o == nil || o.VcSpec == nil {
		return nil, false
	}
	return o.VcSpec, true
}

// HasVcSpec returns a boolean if a field has been set.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) HasVcSpec() bool {
	if o != nil && o.VcSpec != nil {
		return true
	}

	return false
}

// SetVcSpec gets a reference to the given VcenterVchaCredentialsSpec and assigns it to the VcSpec field.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) SetVcSpec(v VcenterVchaCredentialsSpec) {
	o.VcSpec = &v
}

// GetPartial returns the Partial field value if set, zero value otherwise.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetPartial() bool {
	if o == nil || o.Partial == nil {
		var ret bool
		return ret
	}
	return *o.Partial
}

// GetPartialOk returns a tuple with the Partial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) GetPartialOk() (*bool, bool) {
	if o == nil || o.Partial == nil {
		return nil, false
	}
	return o.Partial, true
}

// HasPartial returns a boolean if a field has been set.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) HasPartial() bool {
	if o != nil && o.Partial != nil {
		return true
	}

	return false
}

// SetPartial gets a reference to the given bool and assigns it to the Partial field.
func (o *VcenterVchaClusterActiveGetVchaClusterActive) SetPartial(v bool) {
	o.Partial = &v
}

func (o VcenterVchaClusterActiveGetVchaClusterActive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VcSpec != nil {
		toSerialize["vc_spec"] = o.VcSpec
	}
	if o.Partial != nil {
		toSerialize["partial"] = o.Partial
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterActiveGetVchaClusterActive struct {
	value *VcenterVchaClusterActiveGetVchaClusterActive
	isSet bool
}

func (v NullableVcenterVchaClusterActiveGetVchaClusterActive) Get() *VcenterVchaClusterActiveGetVchaClusterActive {
	return v.value
}

func (v *NullableVcenterVchaClusterActiveGetVchaClusterActive) Set(val *VcenterVchaClusterActiveGetVchaClusterActive) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterActiveGetVchaClusterActive) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterActiveGetVchaClusterActive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterActiveGetVchaClusterActive(val *VcenterVchaClusterActiveGetVchaClusterActive) *NullableVcenterVchaClusterActiveGetVchaClusterActive {
	return &NullableVcenterVchaClusterActiveGetVchaClusterActive{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterActiveGetVchaClusterActive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterActiveGetVchaClusterActive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


