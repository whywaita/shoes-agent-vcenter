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

// VcenterOvfLibraryItemCreateTarget struct for VcenterOvfLibraryItemCreateTarget
type VcenterOvfLibraryItemCreateTarget struct {
	// Identifier of the library in which a new library item should be created. This {@term field} is not used if the {@name #libraryItemId} {@term field} is specified.
	LibraryId *string `json:"library_id,omitempty"`
	// Identifier of the library item that should be should be updated.
	LibraryItemId *string `json:"library_item_id,omitempty"`
}

// NewVcenterOvfLibraryItemCreateTarget instantiates a new VcenterOvfLibraryItemCreateTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterOvfLibraryItemCreateTarget() *VcenterOvfLibraryItemCreateTarget {
	this := VcenterOvfLibraryItemCreateTarget{}
	return &this
}

// NewVcenterOvfLibraryItemCreateTargetWithDefaults instantiates a new VcenterOvfLibraryItemCreateTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterOvfLibraryItemCreateTargetWithDefaults() *VcenterOvfLibraryItemCreateTarget {
	this := VcenterOvfLibraryItemCreateTarget{}
	return &this
}

// GetLibraryId returns the LibraryId field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryId() string {
	if o == nil || o.LibraryId == nil {
		var ret string
		return ret
	}
	return *o.LibraryId
}

// GetLibraryIdOk returns a tuple with the LibraryId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryIdOk() (*string, bool) {
	if o == nil || o.LibraryId == nil {
		return nil, false
	}
	return o.LibraryId, true
}

// HasLibraryId returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemCreateTarget) HasLibraryId() bool {
	if o != nil && o.LibraryId != nil {
		return true
	}

	return false
}

// SetLibraryId gets a reference to the given string and assigns it to the LibraryId field.
func (o *VcenterOvfLibraryItemCreateTarget) SetLibraryId(v string) {
	o.LibraryId = &v
}

// GetLibraryItemId returns the LibraryItemId field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryItemId() string {
	if o == nil || o.LibraryItemId == nil {
		var ret string
		return ret
	}
	return *o.LibraryItemId
}

// GetLibraryItemIdOk returns a tuple with the LibraryItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemCreateTarget) GetLibraryItemIdOk() (*string, bool) {
	if o == nil || o.LibraryItemId == nil {
		return nil, false
	}
	return o.LibraryItemId, true
}

// HasLibraryItemId returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemCreateTarget) HasLibraryItemId() bool {
	if o != nil && o.LibraryItemId != nil {
		return true
	}

	return false
}

// SetLibraryItemId gets a reference to the given string and assigns it to the LibraryItemId field.
func (o *VcenterOvfLibraryItemCreateTarget) SetLibraryItemId(v string) {
	o.LibraryItemId = &v
}

func (o VcenterOvfLibraryItemCreateTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LibraryId != nil {
		toSerialize["library_id"] = o.LibraryId
	}
	if o.LibraryItemId != nil {
		toSerialize["library_item_id"] = o.LibraryItemId
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterOvfLibraryItemCreateTarget struct {
	value *VcenterOvfLibraryItemCreateTarget
	isSet bool
}

func (v NullableVcenterOvfLibraryItemCreateTarget) Get() *VcenterOvfLibraryItemCreateTarget {
	return v.value
}

func (v *NullableVcenterOvfLibraryItemCreateTarget) Set(val *VcenterOvfLibraryItemCreateTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfLibraryItemCreateTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfLibraryItemCreateTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfLibraryItemCreateTarget(val *VcenterOvfLibraryItemCreateTarget) *NullableVcenterOvfLibraryItemCreateTarget {
	return &NullableVcenterOvfLibraryItemCreateTarget{value: val, isSet: true}
}

func (v NullableVcenterOvfLibraryItemCreateTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfLibraryItemCreateTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

