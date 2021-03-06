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

// VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone struct for VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone
type VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone struct {
	Spec *VcenterDeploymentStandalonePscSpec `json:"spec,omitempty"`
}

// NewVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone instantiates a new VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone() *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone {
	this := VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone{}
	return &this
}

// NewVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandaloneWithDefaults instantiates a new VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandaloneWithDefaults() *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone {
	this := VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) GetSpec() VcenterDeploymentStandalonePscSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterDeploymentStandalonePscSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) GetSpecOk() (*VcenterDeploymentStandalonePscSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterDeploymentStandalonePscSpec and assigns it to the Spec field.
func (o *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) SetSpec(v VcenterDeploymentStandalonePscSpec) {
	o.Spec = &v
}

func (o VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone struct {
	value *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone
	isSet bool
}

func (v NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) Get() *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone {
	return v.value
}

func (v *NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) Set(val *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone(val *VcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) *NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone {
	return &NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone{value: val, isSet: true}
}

func (v NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentInstallPscStandaloneCheckDeploymentInstallPscStandalone) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


