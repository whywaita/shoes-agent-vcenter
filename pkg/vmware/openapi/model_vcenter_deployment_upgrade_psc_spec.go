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

// VcenterDeploymentUpgradePscSpec struct for VcenterDeploymentUpgradePscSpec
type VcenterDeploymentUpgradePscSpec struct {
	// Customer experience improvement program should be enabled or disabled for this Platform Services Controller upgrade.
	CeipEnabled bool `json:"ceip_enabled"`
}

// NewVcenterDeploymentUpgradePscSpec instantiates a new VcenterDeploymentUpgradePscSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentUpgradePscSpec(ceipEnabled bool) *VcenterDeploymentUpgradePscSpec {
	this := VcenterDeploymentUpgradePscSpec{}
	this.CeipEnabled = ceipEnabled
	return &this
}

// NewVcenterDeploymentUpgradePscSpecWithDefaults instantiates a new VcenterDeploymentUpgradePscSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentUpgradePscSpecWithDefaults() *VcenterDeploymentUpgradePscSpec {
	this := VcenterDeploymentUpgradePscSpec{}
	return &this
}

// GetCeipEnabled returns the CeipEnabled field value
func (o *VcenterDeploymentUpgradePscSpec) GetCeipEnabled() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CeipEnabled
}

// GetCeipEnabledOk returns a tuple with the CeipEnabled field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradePscSpec) GetCeipEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CeipEnabled, true
}

// SetCeipEnabled sets field value
func (o *VcenterDeploymentUpgradePscSpec) SetCeipEnabled(v bool) {
	o.CeipEnabled = v
}

func (o VcenterDeploymentUpgradePscSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ceip_enabled"] = o.CeipEnabled
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentUpgradePscSpec struct {
	value *VcenterDeploymentUpgradePscSpec
	isSet bool
}

func (v NullableVcenterDeploymentUpgradePscSpec) Get() *VcenterDeploymentUpgradePscSpec {
	return v.value
}

func (v *NullableVcenterDeploymentUpgradePscSpec) Set(val *VcenterDeploymentUpgradePscSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentUpgradePscSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentUpgradePscSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentUpgradePscSpec(val *VcenterDeploymentUpgradePscSpec) *NullableVcenterDeploymentUpgradePscSpec {
	return &NullableVcenterDeploymentUpgradePscSpec{value: val, isSet: true}
}

func (v NullableVcenterDeploymentUpgradePscSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentUpgradePscSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

