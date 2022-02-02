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

// VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs struct for VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs
type VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs struct {
	// content to be converted to the spec.
	CustomizationSpec *string `json:"customization_spec,omitempty"`
}

// NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs instantiates a new VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs() *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs {
	this := VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs{}
	return &this
}

// NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsWithDefaults instantiates a new VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecsWithDefaults() *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs {
	this := VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs{}
	return &this
}

// GetCustomizationSpec returns the CustomizationSpec field value if set, zero value otherwise.
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) GetCustomizationSpec() string {
	if o == nil || o.CustomizationSpec == nil {
		var ret string
		return ret
	}
	return *o.CustomizationSpec
}

// GetCustomizationSpecOk returns a tuple with the CustomizationSpec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) GetCustomizationSpecOk() (*string, bool) {
	if o == nil || o.CustomizationSpec == nil {
		return nil, false
	}
	return o.CustomizationSpec, true
}

// HasCustomizationSpec returns a boolean if a field has been set.
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) HasCustomizationSpec() bool {
	if o != nil && o.CustomizationSpec != nil {
		return true
	}

	return false
}

// SetCustomizationSpec gets a reference to the given string and assigns it to the CustomizationSpec field.
func (o *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) SetCustomizationSpec(v string) {
	o.CustomizationSpec = &v
}

func (o VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CustomizationSpec != nil {
		toSerialize["customization_spec"] = o.CustomizationSpec
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs struct {
	value *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs
	isSet bool
}

func (v NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) Get() *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs {
	return v.value
}

func (v *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) Set(val *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs(val *VcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs {
	return &NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs{value: val, isSet: true}
}

func (v NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCustomizationSpecsImportSpecificationGuestCustomizationSpecs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


