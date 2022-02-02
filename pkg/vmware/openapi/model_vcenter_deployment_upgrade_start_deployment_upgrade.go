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

// VcenterDeploymentUpgradeStartDeploymentUpgrade struct for VcenterDeploymentUpgradeStartDeploymentUpgrade
type VcenterDeploymentUpgradeStartDeploymentUpgrade struct {
	Spec *VcenterDeploymentUpgradeUpgradeSpec `json:"spec,omitempty"`
}

// NewVcenterDeploymentUpgradeStartDeploymentUpgrade instantiates a new VcenterDeploymentUpgradeStartDeploymentUpgrade object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentUpgradeStartDeploymentUpgrade() *VcenterDeploymentUpgradeStartDeploymentUpgrade {
	this := VcenterDeploymentUpgradeStartDeploymentUpgrade{}
	return &this
}

// NewVcenterDeploymentUpgradeStartDeploymentUpgradeWithDefaults instantiates a new VcenterDeploymentUpgradeStartDeploymentUpgrade object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentUpgradeStartDeploymentUpgradeWithDefaults() *VcenterDeploymentUpgradeStartDeploymentUpgrade {
	this := VcenterDeploymentUpgradeStartDeploymentUpgrade{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterDeploymentUpgradeStartDeploymentUpgrade) GetSpec() VcenterDeploymentUpgradeUpgradeSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterDeploymentUpgradeUpgradeSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentUpgradeStartDeploymentUpgrade) GetSpecOk() (*VcenterDeploymentUpgradeUpgradeSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterDeploymentUpgradeStartDeploymentUpgrade) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterDeploymentUpgradeUpgradeSpec and assigns it to the Spec field.
func (o *VcenterDeploymentUpgradeStartDeploymentUpgrade) SetSpec(v VcenterDeploymentUpgradeUpgradeSpec) {
	o.Spec = &v
}

func (o VcenterDeploymentUpgradeStartDeploymentUpgrade) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentUpgradeStartDeploymentUpgrade struct {
	value *VcenterDeploymentUpgradeStartDeploymentUpgrade
	isSet bool
}

func (v NullableVcenterDeploymentUpgradeStartDeploymentUpgrade) Get() *VcenterDeploymentUpgradeStartDeploymentUpgrade {
	return v.value
}

func (v *NullableVcenterDeploymentUpgradeStartDeploymentUpgrade) Set(val *VcenterDeploymentUpgradeStartDeploymentUpgrade) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentUpgradeStartDeploymentUpgrade) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentUpgradeStartDeploymentUpgrade) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentUpgradeStartDeploymentUpgrade(val *VcenterDeploymentUpgradeStartDeploymentUpgrade) *NullableVcenterDeploymentUpgradeStartDeploymentUpgrade {
	return &NullableVcenterDeploymentUpgradeStartDeploymentUpgrade{value: val, isSet: true}
}

func (v NullableVcenterDeploymentUpgradeStartDeploymentUpgrade) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentUpgradeStartDeploymentUpgrade) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

