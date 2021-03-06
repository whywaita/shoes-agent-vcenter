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

// VcenterDeploymentMigrateGetDeploymentMigrateResp struct for VcenterDeploymentMigrateGetDeploymentMigrateResp
type VcenterDeploymentMigrateGetDeploymentMigrateResp struct {
	Value VcenterDeploymentMigrateMigrateSpec `json:"value"`
}

// NewVcenterDeploymentMigrateGetDeploymentMigrateResp instantiates a new VcenterDeploymentMigrateGetDeploymentMigrateResp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentMigrateGetDeploymentMigrateResp(value VcenterDeploymentMigrateMigrateSpec) *VcenterDeploymentMigrateGetDeploymentMigrateResp {
	this := VcenterDeploymentMigrateGetDeploymentMigrateResp{}
	this.Value = value
	return &this
}

// NewVcenterDeploymentMigrateGetDeploymentMigrateRespWithDefaults instantiates a new VcenterDeploymentMigrateGetDeploymentMigrateResp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentMigrateGetDeploymentMigrateRespWithDefaults() *VcenterDeploymentMigrateGetDeploymentMigrateResp {
	this := VcenterDeploymentMigrateGetDeploymentMigrateResp{}
	return &this
}

// GetValue returns the Value field value
func (o *VcenterDeploymentMigrateGetDeploymentMigrateResp) GetValue() VcenterDeploymentMigrateMigrateSpec {
	if o == nil {
		var ret VcenterDeploymentMigrateMigrateSpec
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentMigrateGetDeploymentMigrateResp) GetValueOk() (*VcenterDeploymentMigrateMigrateSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *VcenterDeploymentMigrateGetDeploymentMigrateResp) SetValue(v VcenterDeploymentMigrateMigrateSpec) {
	o.Value = v
}

func (o VcenterDeploymentMigrateGetDeploymentMigrateResp) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentMigrateGetDeploymentMigrateResp struct {
	value *VcenterDeploymentMigrateGetDeploymentMigrateResp
	isSet bool
}

func (v NullableVcenterDeploymentMigrateGetDeploymentMigrateResp) Get() *VcenterDeploymentMigrateGetDeploymentMigrateResp {
	return v.value
}

func (v *NullableVcenterDeploymentMigrateGetDeploymentMigrateResp) Set(val *VcenterDeploymentMigrateGetDeploymentMigrateResp) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentMigrateGetDeploymentMigrateResp) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentMigrateGetDeploymentMigrateResp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentMigrateGetDeploymentMigrateResp(val *VcenterDeploymentMigrateGetDeploymentMigrateResp) *NullableVcenterDeploymentMigrateGetDeploymentMigrateResp {
	return &NullableVcenterDeploymentMigrateGetDeploymentMigrateResp{value: val, isSet: true}
}

func (v NullableVcenterDeploymentMigrateGetDeploymentMigrateResp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentMigrateGetDeploymentMigrateResp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


