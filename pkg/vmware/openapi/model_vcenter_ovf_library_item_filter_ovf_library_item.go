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

// VcenterOvfLibraryItemFilterOvfLibraryItem struct for VcenterOvfLibraryItemFilterOvfLibraryItem
type VcenterOvfLibraryItemFilterOvfLibraryItem struct {
	Target *VcenterOvfLibraryItemDeploymentTarget `json:"target,omitempty"`
}

// NewVcenterOvfLibraryItemFilterOvfLibraryItem instantiates a new VcenterOvfLibraryItemFilterOvfLibraryItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterOvfLibraryItemFilterOvfLibraryItem() *VcenterOvfLibraryItemFilterOvfLibraryItem {
	this := VcenterOvfLibraryItemFilterOvfLibraryItem{}
	return &this
}

// NewVcenterOvfLibraryItemFilterOvfLibraryItemWithDefaults instantiates a new VcenterOvfLibraryItemFilterOvfLibraryItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterOvfLibraryItemFilterOvfLibraryItemWithDefaults() *VcenterOvfLibraryItemFilterOvfLibraryItem {
	this := VcenterOvfLibraryItemFilterOvfLibraryItem{}
	return &this
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemFilterOvfLibraryItem) GetTarget() VcenterOvfLibraryItemDeploymentTarget {
	if o == nil || o.Target == nil {
		var ret VcenterOvfLibraryItemDeploymentTarget
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemFilterOvfLibraryItem) GetTargetOk() (*VcenterOvfLibraryItemDeploymentTarget, bool) {
	if o == nil || o.Target == nil {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemFilterOvfLibraryItem) HasTarget() bool {
	if o != nil && o.Target != nil {
		return true
	}

	return false
}

// SetTarget gets a reference to the given VcenterOvfLibraryItemDeploymentTarget and assigns it to the Target field.
func (o *VcenterOvfLibraryItemFilterOvfLibraryItem) SetTarget(v VcenterOvfLibraryItemDeploymentTarget) {
	o.Target = &v
}

func (o VcenterOvfLibraryItemFilterOvfLibraryItem) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Target != nil {
		toSerialize["target"] = o.Target
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterOvfLibraryItemFilterOvfLibraryItem struct {
	value *VcenterOvfLibraryItemFilterOvfLibraryItem
	isSet bool
}

func (v NullableVcenterOvfLibraryItemFilterOvfLibraryItem) Get() *VcenterOvfLibraryItemFilterOvfLibraryItem {
	return v.value
}

func (v *NullableVcenterOvfLibraryItemFilterOvfLibraryItem) Set(val *VcenterOvfLibraryItemFilterOvfLibraryItem) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfLibraryItemFilterOvfLibraryItem) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfLibraryItemFilterOvfLibraryItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfLibraryItemFilterOvfLibraryItem(val *VcenterOvfLibraryItemFilterOvfLibraryItem) *NullableVcenterOvfLibraryItemFilterOvfLibraryItem {
	return &NullableVcenterOvfLibraryItemFilterOvfLibraryItem{value: val, isSet: true}
}

func (v NullableVcenterOvfLibraryItemFilterOvfLibraryItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfLibraryItemFilterOvfLibraryItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


