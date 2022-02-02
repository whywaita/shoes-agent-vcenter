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

// VcenterVchaClusterDeployTaskVchaClusterResp struct for VcenterVchaClusterDeployTaskVchaClusterResp
type VcenterVchaClusterDeployTaskVchaClusterResp struct {
	Value string `json:"value"`
}

// NewVcenterVchaClusterDeployTaskVchaClusterResp instantiates a new VcenterVchaClusterDeployTaskVchaClusterResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVchaClusterDeployTaskVchaClusterResp(value string) *VcenterVchaClusterDeployTaskVchaClusterResp {
	this := VcenterVchaClusterDeployTaskVchaClusterResp{}
	this.Value = value
	return &this
}

// NewVcenterVchaClusterDeployTaskVchaClusterRespWithDefaults instantiates a new VcenterVchaClusterDeployTaskVchaClusterResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVchaClusterDeployTaskVchaClusterRespWithDefaults() *VcenterVchaClusterDeployTaskVchaClusterResp {
	this := VcenterVchaClusterDeployTaskVchaClusterResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterVchaClusterDeployTaskVchaClusterResp) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterVchaClusterDeployTaskVchaClusterResp) GetValueOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterVchaClusterDeployTaskVchaClusterResp) SetValue(v string) {
	o.Value = v
}

func (o VcenterVchaClusterDeployTaskVchaClusterResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVchaClusterDeployTaskVchaClusterResp struct {
	value *VcenterVchaClusterDeployTaskVchaClusterResp
	isSet bool
}

func (v NullableVcenterVchaClusterDeployTaskVchaClusterResp) Get() *VcenterVchaClusterDeployTaskVchaClusterResp {
	return v.value
}

func (v *NullableVcenterVchaClusterDeployTaskVchaClusterResp) Set(val *VcenterVchaClusterDeployTaskVchaClusterResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVchaClusterDeployTaskVchaClusterResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVchaClusterDeployTaskVchaClusterResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVchaClusterDeployTaskVchaClusterResp(val *VcenterVchaClusterDeployTaskVchaClusterResp) *NullableVcenterVchaClusterDeployTaskVchaClusterResp {
	return &NullableVcenterVchaClusterDeployTaskVchaClusterResp{value: val, isSet: true}
}

func (v NullableVcenterVchaClusterDeployTaskVchaClusterResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVchaClusterDeployTaskVchaClusterResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


