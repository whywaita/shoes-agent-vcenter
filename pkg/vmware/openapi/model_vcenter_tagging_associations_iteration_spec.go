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

// VcenterTaggingAssociationsIterationSpec struct for VcenterTaggingAssociationsIterationSpec
type VcenterTaggingAssociationsIterationSpec struct {
	// Marker is an opaque token that allows the caller to request the next page of tag associations. If unset or empty, first page of tag associations will be returned. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.tagging.associations.Marker. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.tagging.associations.Marker.
	Marker *string `json:"marker,omitempty"`
}

// NewVcenterTaggingAssociationsIterationSpec instantiates a new VcenterTaggingAssociationsIterationSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTaggingAssociationsIterationSpec() *VcenterTaggingAssociationsIterationSpec {
	this := VcenterTaggingAssociationsIterationSpec{}
	return &this
}

// NewVcenterTaggingAssociationsIterationSpecWithDefaults instantiates a new VcenterTaggingAssociationsIterationSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTaggingAssociationsIterationSpecWithDefaults() *VcenterTaggingAssociationsIterationSpec {
	this := VcenterTaggingAssociationsIterationSpec{}
	return &this
}

// GetMarker returns the Marker field value if set, zero value otherwise.
func (o *VcenterTaggingAssociationsIterationSpec) GetMarker() string {
	if o == nil || o.Marker == nil {
		var ret string
		return ret
	}
	return *o.Marker
}

// GetMarkerOk returns a tuple with the Marker field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTaggingAssociationsIterationSpec) GetMarkerOk() (*string, bool) {
	if o == nil || o.Marker == nil {
		return nil, false
	}
	return o.Marker, true
}

// HasMarker returns a boolean if a field has been set.
func (o *VcenterTaggingAssociationsIterationSpec) HasMarker() bool {
	if o != nil && o.Marker != nil {
		return true
	}

	return false
}

// SetMarker gets a reference to the given string and assigns it to the Marker field.
func (o *VcenterTaggingAssociationsIterationSpec) SetMarker(v string) {
	o.Marker = &v
}

func (o VcenterTaggingAssociationsIterationSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Marker != nil {
		toSerialize["marker"] = o.Marker
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTaggingAssociationsIterationSpec struct {
	value *VcenterTaggingAssociationsIterationSpec
	isSet bool
}

func (v NullableVcenterTaggingAssociationsIterationSpec) Get() *VcenterTaggingAssociationsIterationSpec {
	return v.value
}

func (v *NullableVcenterTaggingAssociationsIterationSpec) Set(val *VcenterTaggingAssociationsIterationSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTaggingAssociationsIterationSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTaggingAssociationsIterationSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTaggingAssociationsIterationSpec(val *VcenterTaggingAssociationsIterationSpec) *NullableVcenterTaggingAssociationsIterationSpec {
	return &NullableVcenterTaggingAssociationsIterationSpec{value: val, isSet: true}
}

func (v NullableVcenterTaggingAssociationsIterationSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTaggingAssociationsIterationSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


