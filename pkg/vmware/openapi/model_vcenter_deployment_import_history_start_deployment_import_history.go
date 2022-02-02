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

// VcenterDeploymentImportHistoryStartDeploymentImportHistory struct for VcenterDeploymentImportHistoryStartDeploymentImportHistory
type VcenterDeploymentImportHistoryStartDeploymentImportHistory struct {
	Spec *VcenterDeploymentImportHistoryCreateSpec `json:"spec,omitempty"`
}

// NewVcenterDeploymentImportHistoryStartDeploymentImportHistory instantiates a new VcenterDeploymentImportHistoryStartDeploymentImportHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentImportHistoryStartDeploymentImportHistory() *VcenterDeploymentImportHistoryStartDeploymentImportHistory {
	this := VcenterDeploymentImportHistoryStartDeploymentImportHistory{}
	return &this
}

// NewVcenterDeploymentImportHistoryStartDeploymentImportHistoryWithDefaults instantiates a new VcenterDeploymentImportHistoryStartDeploymentImportHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentImportHistoryStartDeploymentImportHistoryWithDefaults() *VcenterDeploymentImportHistoryStartDeploymentImportHistory {
	this := VcenterDeploymentImportHistoryStartDeploymentImportHistory{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterDeploymentImportHistoryStartDeploymentImportHistory) GetSpec() VcenterDeploymentImportHistoryCreateSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterDeploymentImportHistoryCreateSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentImportHistoryStartDeploymentImportHistory) GetSpecOk() (*VcenterDeploymentImportHistoryCreateSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterDeploymentImportHistoryStartDeploymentImportHistory) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterDeploymentImportHistoryCreateSpec and assigns it to the Spec field.
func (o *VcenterDeploymentImportHistoryStartDeploymentImportHistory) SetSpec(v VcenterDeploymentImportHistoryCreateSpec) {
	o.Spec = &v
}

func (o VcenterDeploymentImportHistoryStartDeploymentImportHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory struct {
	value *VcenterDeploymentImportHistoryStartDeploymentImportHistory
	isSet bool
}

func (v NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory) Get() *VcenterDeploymentImportHistoryStartDeploymentImportHistory {
	return v.value
}

func (v *NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory) Set(val *VcenterDeploymentImportHistoryStartDeploymentImportHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentImportHistoryStartDeploymentImportHistory(val *VcenterDeploymentImportHistoryStartDeploymentImportHistory) *NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory {
	return &NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory{value: val, isSet: true}
}

func (v NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentImportHistoryStartDeploymentImportHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


