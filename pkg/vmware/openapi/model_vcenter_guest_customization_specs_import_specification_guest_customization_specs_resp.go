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

// VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp struct for VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp
type VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp struct {
	Value VcenterGuestCustomizationSpecsCreateSpec `json:"value"`
}

// NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp instantiates a new VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp(value VcenterGuestCustomizationSpecsCreateSpec) *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp {
	this := VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp{}
	this.Value = value
	return &this
}

// NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsRespWithDefaults instantiates a new VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsRespWithDefaults() *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp {
	this := VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) GetValue() VcenterGuestCustomizationSpecsCreateSpec {
	if o == nil {
		var ret VcenterGuestCustomizationSpecsCreateSpec
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) GetValueOk() (*VcenterGuestCustomizationSpecsCreateSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) SetValue(v VcenterGuestCustomizationSpecsCreateSpec) {
	o.Value = v
}

func (o VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp struct {
	value *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp
	isSet bool
}

func (v NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) Get() *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp {
	return v.value
}

func (v *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) Set(val *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp(val *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp {
	return &NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp{value: val, isSet: true}
}

func (v NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


