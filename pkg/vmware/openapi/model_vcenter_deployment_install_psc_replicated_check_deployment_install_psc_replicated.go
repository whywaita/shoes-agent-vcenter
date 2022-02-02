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

// VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated struct for VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated
type VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated struct {
	Spec *VcenterDeploymentReplicatedPscSpec `json:"spec,omitempty"`
}

// NewVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated instantiates a new VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated() *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated {
	this := VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated{}
	return &this
}

// NewVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicatedWithDefaults instantiates a new VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicatedWithDefaults() *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated {
	this := VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) GetSpec() VcenterDeploymentReplicatedPscSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterDeploymentReplicatedPscSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) GetSpecOk() (*VcenterDeploymentReplicatedPscSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterDeploymentReplicatedPscSpec and assigns it to the Spec field.
func (o *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) SetSpec(v VcenterDeploymentReplicatedPscSpec) {
	o.Spec = &v
}

func (o VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated struct {
	value *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated
	isSet bool
}

func (v NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) Get() *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated {
	return v.value
}

func (v *NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) Set(val *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated(val *VcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) *NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated {
	return &NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated{value: val, isSet: true}
}

func (v NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentInstallPscReplicatedCheckDeploymentInstallPscReplicated) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


