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

// VcenterIsoImageMountIsoImageResp struct for VcenterIsoImageMountIsoImageResp
type VcenterIsoImageMountIsoImageResp struct {
	Value string `json:"value"`
}

// NewVcenterIsoImageMountIsoImageResp instantiates a new VcenterIsoImageMountIsoImageResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterIsoImageMountIsoImageResp(value string) *VcenterIsoImageMountIsoImageResp {
	this := VcenterIsoImageMountIsoImageResp{}
	this.Value = value
	return &this
}

// NewVcenterIsoImageMountIsoImageRespWithDefaults instantiates a new VcenterIsoImageMountIsoImageResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterIsoImageMountIsoImageRespWithDefaults() *VcenterIsoImageMountIsoImageResp {
	this := VcenterIsoImageMountIsoImageResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterIsoImageMountIsoImageResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterIsoImageMountIsoImageResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterIsoImageMountIsoImageResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterIsoImageMountIsoImageResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterIsoImageMountIsoImageResp struct {
	value *VcenterIsoImageMountIsoImageResp
	isSet bool
}

func (v NullableVcenterIsoImageMountIsoImageResp) Get() *VcenterIsoImageMountIsoImageResp {
	return v.value
}

func (v *NullableVcenterIsoImageMountIsoImageResp) Set(val *VcenterIsoImageMountIsoImageResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterIsoImageMountIsoImageResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterIsoImageMountIsoImageResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterIsoImageMountIsoImageResp(val *VcenterIsoImageMountIsoImageResp) *NullableVcenterIsoImageMountIsoImageResp {
	return &NullableVcenterIsoImageMountIsoImageResp{value: val, isSet: true}
}

func (v NullableVcenterIsoImageMountIsoImageResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterIsoImageMountIsoImageResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


