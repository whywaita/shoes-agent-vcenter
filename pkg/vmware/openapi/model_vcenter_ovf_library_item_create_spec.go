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

// VcenterOvfLibraryItemCreateSpec struct for VcenterOvfLibraryItemCreateSpec
type VcenterOvfLibraryItemCreateSpec struct {
	// Name to use in the OVF descriptor stored in the library item.
	Name *string `json:"name,omitempty"`
	// Description to use in the OVF descriptor stored in the library item.
	Description *string `json:"description,omitempty"`
	// Flags to use for OVF package creation. The supported flags can be obtained using {@link ExportFlag#list}.
	Flags *[]string `json:"flags,omitempty"`
}

// NewVcenterOvfLibraryItemCreateSpec instantiates a new VcenterOvfLibraryItemCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterOvfLibraryItemCreateSpec() *VcenterOvfLibraryItemCreateSpec {
	this := VcenterOvfLibraryItemCreateSpec{}
	return &this
}

// NewVcenterOvfLibraryItemCreateSpecWithDefaults instantiates a new VcenterOvfLibraryItemCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterOvfLibraryItemCreateSpecWithDefaults() *VcenterOvfLibraryItemCreateSpec {
	this := VcenterOvfLibraryItemCreateSpec{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemCreateSpec) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemCreateSpec) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemCreateSpec) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VcenterOvfLibraryItemCreateSpec) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemCreateSpec) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemCreateSpec) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemCreateSpec) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *VcenterOvfLibraryItemCreateSpec) SetDescription(v string) {
	o.Description = &v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *VcenterOvfLibraryItemCreateSpec) GetFlags() []string {
	if o == nil || o.Flags == nil {
		var ret []string
		return ret
	}
	return *o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterOvfLibraryItemCreateSpec) GetFlagsOk() (*[]string, bool) {
	if o == nil || o.Flags == nil {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *VcenterOvfLibraryItemCreateSpec) HasFlags() bool {
	if o != nil && o.Flags != nil {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *VcenterOvfLibraryItemCreateSpec) SetFlags(v []string) {
	o.Flags = &v
}

func (o VcenterOvfLibraryItemCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Flags != nil {
		toSerialize["flags"] = o.Flags
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterOvfLibraryItemCreateSpec struct {
	value *VcenterOvfLibraryItemCreateSpec
	isSet bool
}

func (v NullableVcenterOvfLibraryItemCreateSpec) Get() *VcenterOvfLibraryItemCreateSpec {
	return v.value
}

func (v *NullableVcenterOvfLibraryItemCreateSpec) Set(val *VcenterOvfLibraryItemCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterOvfLibraryItemCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterOvfLibraryItemCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterOvfLibraryItemCreateSpec(val *VcenterOvfLibraryItemCreateSpec) *NullableVcenterOvfLibraryItemCreateSpec {
	return &NullableVcenterOvfLibraryItemCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterOvfLibraryItemCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterOvfLibraryItemCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


