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

// VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp struct for VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp
type VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp struct {
	Value VcenterNamespacesNamespaceTemplatesInfo `json:"value"`
}

// NewVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp instantiates a new VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp(value VcenterNamespacesNamespaceTemplatesInfo) *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp {
	this := VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp{}
	this.Value = value
	return &this
}

// NewVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesRespWithDefaults instantiates a new VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesRespWithDefaults() *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp {
	this := VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) GetValue() VcenterNamespacesNamespaceTemplatesInfo {
	if o == nil {
		var ret VcenterNamespacesNamespaceTemplatesInfo
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) GetValueOk() (*VcenterNamespacesNamespaceTemplatesInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) SetValue(v VcenterNamespacesNamespaceTemplatesInfo) {
	o.Value = v
}

func (o VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp struct {
	value *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp
	isSet bool
}

func (v NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) Get() *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp {
	return v.value
}

func (v *NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) Set(val *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp(val *VcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) *NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp {
	return &NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp{value: val, isSet: true}
}

func (v NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespacesNamespaceTemplatesGetNamespacesNamespaceTemplatesResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

