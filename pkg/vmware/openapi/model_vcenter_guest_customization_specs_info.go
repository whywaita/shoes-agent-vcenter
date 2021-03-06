/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VcenterGuestCustomizationSpecsInfo struct for VcenterGuestCustomizationSpecsInfo
type VcenterGuestCustomizationSpecsInfo struct {
	// Time when the specification was last modified.
	LastModified time.Time `json:"last_modified"`
	Spec VcenterGuestCustomizationSpecsSpec `json:"spec"`
}

// NewVcenterGuestCustomizationSpecsInfo instantiates a new VcenterGuestCustomizationSpecsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCustomizationSpecsInfo(lastModified time.Time, spec VcenterGuestCustomizationSpecsSpec) *VcenterGuestCustomizationSpecsInfo {
	this := VcenterGuestCustomizationSpecsInfo{}
	this.LastModified = lastModified
	this.Spec = spec
	return &this
}

// NewVcenterGuestCustomizationSpecsInfoWithDefaults instantiates a new VcenterGuestCustomizationSpecsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCustomizationSpecsInfoWithDefaults() *VcenterGuestCustomizationSpecsInfo {
	this := VcenterGuestCustomizationSpecsInfo{}
	return &this
}

// GetLastModified returns the LastModified field value
func (o *VcenterGuestCustomizationSpecsInfo) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsInfo) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *VcenterGuestCustomizationSpecsInfo) SetLastModified(v time.Time) {
	o.LastModified = v
}

// GetSpec returns the Spec field value
func (o *VcenterGuestCustomizationSpecsInfo) GetSpec() VcenterGuestCustomizationSpecsSpec {
	if o == nil {
		var ret VcenterGuestCustomizationSpecsSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsInfo) GetSpecOk() (*VcenterGuestCustomizationSpecsSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *VcenterGuestCustomizationSpecsInfo) SetSpec(v VcenterGuestCustomizationSpecsSpec) {
	o.Spec = v
}

func (o VcenterGuestCustomizationSpecsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["last_modified"] = o.LastModified
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCustomizationSpecsInfo struct {
	value *VcenterGuestCustomizationSpecsInfo
	isSet bool
}

func (v NullableVcenterGuestCustomizationSpecsInfo) Get() *VcenterGuestCustomizationSpecsInfo {
	return v.value
}

func (v *NullableVcenterGuestCustomizationSpecsInfo) Set(val *VcenterGuestCustomizationSpecsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCustomizationSpecsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCustomizationSpecsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCustomizationSpecsInfo(val *VcenterGuestCustomizationSpecsInfo) *NullableVcenterGuestCustomizationSpecsInfo {
	return &NullableVcenterGuestCustomizationSpecsInfo{value: val, isSet: true}
}

func (v NullableVcenterGuestCustomizationSpecsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCustomizationSpecsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


