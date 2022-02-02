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

// VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp struct for VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp
type VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp struct {
	Value VcenterContentRegistriesHarborProjectsInfo `json:"value"`
}

// NewVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp instantiates a new VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp(value VcenterContentRegistriesHarborProjectsInfo) *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp {
	this := VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp{}
	this.Value = value
	return &this
}

// NewVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsRespWithDefaults instantiates a new VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsRespWithDefaults() *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp {
	this := VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) GetValue() VcenterContentRegistriesHarborProjectsInfo {
	if o == nil {
		var ret VcenterContentRegistriesHarborProjectsInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) GetValueOk() (*VcenterContentRegistriesHarborProjectsInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) SetValue(v VcenterContentRegistriesHarborProjectsInfo) {
	o.Value = v
}

func (o VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp struct {
	value *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp
	isSet bool
}

func (v NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) Get() *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp {
	return v.value
}

func (v *NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) Set(val *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp(val *VcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) *NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp {
	return &NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp{value: val, isSet: true}
}

func (v NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterContentRegistriesHarborProjectsGetContentRegistriesHarborProjectsResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


