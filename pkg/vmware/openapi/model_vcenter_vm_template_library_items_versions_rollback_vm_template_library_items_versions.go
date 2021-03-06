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

// VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions struct for VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions
type VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions struct {
	Spec *VcenterVmTemplateLibraryItemsVersionsRollbackSpec `json:"spec,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions instantiates a new VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions() *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions {
	this := VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersionsWithDefaults instantiates a new VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersionsWithDefaults() *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions {
	this := VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) GetSpec() VcenterVmTemplateLibraryItemsVersionsRollbackSpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVmTemplateLibraryItemsVersionsRollbackSpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) GetSpecOk() (*VcenterVmTemplateLibraryItemsVersionsRollbackSpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVmTemplateLibraryItemsVersionsRollbackSpec and assigns it to the Spec field.
func (o *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) SetSpec(v VcenterVmTemplateLibraryItemsVersionsRollbackSpec) {
	o.Spec = &v
}

func (o VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions struct {
	value *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) Get() *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) Set(val *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions(val *VcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) *NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions {
	return &NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsVersionsRollbackVmTemplateLibraryItemsVersions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


