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

// VcenterVmToolsUpdateSpec struct for VcenterVmToolsUpdateSpec
type VcenterVmToolsUpdateSpec struct {
	UpgradePolicy *VcenterVmToolsUpgradePolicy `json:"upgrade_policy,omitempty"`
}

// NewVcenterVmToolsUpdateSpec instantiates a new VcenterVmToolsUpdateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmToolsUpdateSpec() *VcenterVmToolsUpdateSpec {
	this := VcenterVmToolsUpdateSpec{}
	return &this
}

// NewVcenterVmToolsUpdateSpecWithDefaults instantiates a new VcenterVmToolsUpdateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmToolsUpdateSpecWithDefaults() *VcenterVmToolsUpdateSpec {
	this := VcenterVmToolsUpdateSpec{}
	return &this
}

// GetUpgradePolicy returns the UpgradePolicy field value if set, zero value otherwise.
func (o *VcenterVmToolsUpdateSpec) GetUpgradePolicy() VcenterVmToolsUpgradePolicy {
	if o == nil || o.UpgradePolicy == nil {
		var ret VcenterVmToolsUpgradePolicy
		return ret
	}
	return *o.UpgradePolicy
}

// GetUpgradePolicyOk returns a tuple with the UpgradePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmToolsUpdateSpec) GetUpgradePolicyOk() (*VcenterVmToolsUpgradePolicy, bool) {
	if o == nil || o.UpgradePolicy == nil {
		return nil, false
	}
	return o.UpgradePolicy, true
}

// HasUpgradePolicy returns a boolean if a field has been set.
func (o *VcenterVmToolsUpdateSpec) HasUpgradePolicy() bool {
	if o != nil && o.UpgradePolicy != nil {
		return true
	}

	return false
}

// SetUpgradePolicy gets a reference to the given VcenterVmToolsUpgradePolicy and assigns it to the UpgradePolicy field.
func (o *VcenterVmToolsUpdateSpec) SetUpgradePolicy(v VcenterVmToolsUpgradePolicy) {
	o.UpgradePolicy = &v
}

func (o VcenterVmToolsUpdateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.UpgradePolicy != nil {
		toSerialize["upgrade_policy"] = o.UpgradePolicy
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmToolsUpdateSpec struct {
	value *VcenterVmToolsUpdateSpec
	isSet bool
}

func (v NullableVcenterVmToolsUpdateSpec) Get() *VcenterVmToolsUpdateSpec {
	return v.value
}

func (v *NullableVcenterVmToolsUpdateSpec) Set(val *VcenterVmToolsUpdateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmToolsUpdateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmToolsUpdateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmToolsUpdateSpec(val *VcenterVmToolsUpdateSpec) *NullableVcenterVmToolsUpdateSpec {
	return &NullableVcenterVmToolsUpdateSpec{value: val, isSet: true}
}

func (v NullableVcenterVmToolsUpdateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmToolsUpdateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


