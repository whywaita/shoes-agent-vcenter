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

// VcenterDeploymentInstallStartDeploymentInstall struct for VcenterDeploymentInstallStartDeploymentInstall
type VcenterDeploymentInstallStartDeploymentInstall struct {
	Spec *VcenterDeploymentInstallInstallSpec `json:"spec,omitempty"`
}

// NewVcenterDeploymentInstallStartDeploymentInstall instantiates a new VcenterDeploymentInstallStartDeploymentInstall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentInstallStartDeploymentInstall() *VcenterDeploymentInstallStartDeploymentInstall {
	this := VcenterDeploymentInstallStartDeploymentInstall{}
	return &this
}

// NewVcenterDeploymentInstallStartDeploymentInstallWithDefaults instantiates a new VcenterDeploymentInstallStartDeploymentInstall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentInstallStartDeploymentInstallWithDefaults() *VcenterDeploymentInstallStartDeploymentInstall {
	this := VcenterDeploymentInstallStartDeploymentInstall{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterDeploymentInstallStartDeploymentInstall) GetSpec() VcenterDeploymentInstallInstallSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterDeploymentInstallInstallSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentInstallStartDeploymentInstall) GetSpecOk() (*VcenterDeploymentInstallInstallSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterDeploymentInstallStartDeploymentInstall) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterDeploymentInstallInstallSpec and assigns it to the Spec field.
func (o *VcenterDeploymentInstallStartDeploymentInstall) SetSpec(v VcenterDeploymentInstallInstallSpec) {
	o.Spec = &v
}

func (o VcenterDeploymentInstallStartDeploymentInstall) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentInstallStartDeploymentInstall struct {
	value *VcenterDeploymentInstallStartDeploymentInstall
	isSet bool
}

func (v NullableVcenterDeploymentInstallStartDeploymentInstall) Get() *VcenterDeploymentInstallStartDeploymentInstall {
	return v.value
}

func (v *NullableVcenterDeploymentInstallStartDeploymentInstall) Set(val *VcenterDeploymentInstallStartDeploymentInstall) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentInstallStartDeploymentInstall) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentInstallStartDeploymentInstall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentInstallStartDeploymentInstall(val *VcenterDeploymentInstallStartDeploymentInstall) *NullableVcenterDeploymentInstallStartDeploymentInstall {
	return &NullableVcenterDeploymentInstallStartDeploymentInstall{value: val, isSet: true}
}

func (v NullableVcenterDeploymentInstallStartDeploymentInstall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentInstallStartDeploymentInstall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


