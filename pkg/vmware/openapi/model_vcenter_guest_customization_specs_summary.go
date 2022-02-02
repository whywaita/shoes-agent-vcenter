/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VcenterGuestCustomizationSpecsSummary struct for VcenterGuestCustomizationSpecsSummary
type VcenterGuestCustomizationSpecsSummary struct {
	// Name of the guest customization specification. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.guest.CustomizationSpec. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.guest.CustomizationSpec.
	Name string `json:"name"`
	// Description of the guest customization specification.
	Description string `json:"description"`
	OSType VcenterGuestCustomizationSpecsOsType `json:"OS_type"`
	// Date and tme when this guest customization specification was last modified.
	LastModified time.Time `json:"last_modified"`
}

// NewVcenterGuestCustomizationSpecsSummary instantiates a new VcenterGuestCustomizationSpecsSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestCustomizationSpecsSummary(name string, description string, oSType VcenterGuestCustomizationSpecsOsType, lastModified time.Time) *VcenterGuestCustomizationSpecsSummary {
	this := VcenterGuestCustomizationSpecsSummary{}
	this.Name = name
	this.Description = description
	this.OSType = oSType
	this.LastModified = lastModified
	return &this
}

// NewVcenterGuestCustomizationSpecsSummaryWithDefaults instantiates a new VcenterGuestCustomizationSpecsSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestCustomizationSpecsSummaryWithDefaults() *VcenterGuestCustomizationSpecsSummary {
	this := VcenterGuestCustomizationSpecsSummary{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterGuestCustomizationSpecsSummary) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsSummary) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterGuestCustomizationSpecsSummary) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *VcenterGuestCustomizationSpecsSummary) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsSummary) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VcenterGuestCustomizationSpecsSummary) SetDescription(v string) {
	o.Description = v
}

// GetOSType returns the OSType field value
func (o *VcenterGuestCustomizationSpecsSummary) GetOSType() VcenterGuestCustomizationSpecsOsType {
	if o == nil {
		var ret VcenterGuestCustomizationSpecsOsType
		return ret
	}

	return o.OSType
}

// GetOSTypeOk returns a tuple with the OSType field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsSummary) GetOSTypeOk() (*VcenterGuestCustomizationSpecsOsType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.OSType, true
}

// SetOSType sets field value
func (o *VcenterGuestCustomizationSpecsSummary) SetOSType(v VcenterGuestCustomizationSpecsOsType) {
	o.OSType = v
}

// GetLastModified returns the LastModified field value
func (o *VcenterGuestCustomizationSpecsSummary) GetLastModified() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.LastModified
}

// GetLastModifiedOk returns a tuple with the LastModified field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestCustomizationSpecsSummary) GetLastModifiedOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastModified, true
}

// SetLastModified sets field value
func (o *VcenterGuestCustomizationSpecsSummary) SetLastModified(v time.Time) {
	o.LastModified = v
}

func (o VcenterGuestCustomizationSpecsSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["OS_type"] = o.OSType
	}
	if true {
		toSerialize["last_modified"] = o.LastModified
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestCustomizationSpecsSummary struct {
	value *VcenterGuestCustomizationSpecsSummary
	isSet bool
}

func (v NullableVcenterGuestCustomizationSpecsSummary) Get() *VcenterGuestCustomizationSpecsSummary {
	return v.value
}

func (v *NullableVcenterGuestCustomizationSpecsSummary) Set(val *VcenterGuestCustomizationSpecsSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestCustomizationSpecsSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestCustomizationSpecsSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestCustomizationSpecsSummary(val *VcenterGuestCustomizationSpecsSummary) *NullableVcenterGuestCustomizationSpecsSummary {
	return &NullableVcenterGuestCustomizationSpecsSummary{value: val, isSet: true}
}

func (v NullableVcenterGuestCustomizationSpecsSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestCustomizationSpecsSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


