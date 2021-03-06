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

// VcenterGuestCustomizationSpecsCreateSpec struct for VcenterGuestCustomizationSpecsCreateSpec
type VcenterGuestCustomizationSpecsCreateSpec struct {
	Spec VcenterGuestCustomizationSpec `json:"spec"`
	// Description of the specification.
	Description string `json:"description"`
	// Name of the specification.
	Name string `json:"name"`
}

// NewVcenterGuestCustomizationSpecsCreateSpec instantiates a new VcenterGuestCustomizationSpecsCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCustomizationSpecsCreateSpec(spec VcenterGuestCustomizationSpec, description string, name string) *VcenterGuestCustomizationSpecsCreateSpec {
	this := VcenterGuestCustomizationSpecsCreateSpec{}
	this.Spec = spec
	this.Description = description
	this.Name = name
	return &this
}

// NewVcenterGuestCustomizationSpecsCreateSpecWithDefaults instantiates a new VcenterGuestCustomizationSpecsCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCustomizationSpecsCreateSpecWithDefaults() *VcenterGuestCustomizationSpecsCreateSpec {
	this := VcenterGuestCustomizationSpecsCreateSpec{}
	return &this
}

// GetSpec returns the Spec field value
func (o *VcenterGuestCustomizationSpecsCreateSpec) GetSpec() VcenterGuestCustomizationSpec {
	if o == nil {
		var ret VcenterGuestCustomizationSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsCreateSpec) GetSpecOk() (*VcenterGuestCustomizationSpec, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *VcenterGuestCustomizationSpecsCreateSpec) SetSpec(v VcenterGuestCustomizationSpec) {
	o.Spec = v
}

// GetDescription returns the Description field value
func (o *VcenterGuestCustomizationSpecsCreateSpec) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VcenterGuestCustomizationSpecsCreateSpec) SetDescription(v string) {
	o.Description = v
}

// GetName returns the Name field value
func (o *VcenterGuestCustomizationSpecsCreateSpec) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsCreateSpec) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterGuestCustomizationSpecsCreateSpec) SetName(v string) {
	o.Name = v
}

func (o VcenterGuestCustomizationSpecsCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["spec"] = o.Spec
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCustomizationSpecsCreateSpec struct {
	value *VcenterGuestCustomizationSpecsCreateSpec
	isSet bool
}

func (v NullableVcenterGuestCustomizationSpecsCreateSpec) Get() *VcenterGuestCustomizationSpecsCreateSpec {
	return v.value
}

func (v *NullableVcenterGuestCustomizationSpecsCreateSpec) Set(val *VcenterGuestCustomizationSpecsCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCustomizationSpecsCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCustomizationSpecsCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCustomizationSpecsCreateSpec(val *VcenterGuestCustomizationSpecsCreateSpec) *NullableVcenterGuestCustomizationSpecsCreateSpec {
	return &NullableVcenterGuestCustomizationSpecsCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterGuestCustomizationSpecsCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCustomizationSpecsCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


