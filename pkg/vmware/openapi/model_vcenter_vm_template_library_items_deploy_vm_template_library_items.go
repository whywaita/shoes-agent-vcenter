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

// VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems struct for VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems
type VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems struct {
	Spec *VcenterVmTemplateLibraryItemsDeploySpec `json:"spec,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems instantiates a new VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems() *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems {
	this := VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItemsWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItemsWithDefaults() *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems {
	this := VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems{}
	return &this
}

// GetSpec returns the Spec field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) GetSpec() VcenterVmTemplateLibraryItemsDeploySpec {
	if o == nil || o.Spec == nil {
		var ret VcenterVmTemplateLibraryItemsDeploySpec
		return ret
	}
	return *o.Spec
}

// GetSpecOk returns a tuple with the Spec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) GetSpecOk() (*VcenterVmTemplateLibraryItemsDeploySpec, bool) {
	if o == nil || o.Spec == nil {
		return nil, false
	}
	return o.Spec, true
}

// HasSpec returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) HasSpec() bool {
	if o != nil && o.Spec != nil {
		return true
	}

	return false
}

// SetSpec gets a reference to the given VcenterVmTemplateLibraryItemsDeploySpec and assigns it to the Spec field.
func (o *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) SetSpec(v VcenterVmTemplateLibraryItemsDeploySpec) {
	o.Spec = &v
}

func (o VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Spec != nil {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems struct {
	value *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) Get() *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) Set(val *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems(val *VcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) *NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems {
	return &NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsDeployVmTemplateLibraryItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


