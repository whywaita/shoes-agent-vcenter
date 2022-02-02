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

// VcenterComputePoliciesCapabilitiesInfo struct for VcenterComputePoliciesCapabilitiesInfo
type VcenterComputePoliciesCapabilitiesInfo struct {
	// Name of the capability.
	Name string `json:"name"`
	// Description of the capability.
	Description string `json:"description"`
	// Identifier of the {@term structure} used to create a policy based on this capability. See {@link vcenter.compute.Policies#create}.
	CreateSpecType string `json:"create_spec_type"`
	// Identifier of the {@term structure} returned when retrieving information about a policy based on this capability. See {@link vcenter.compute.Policies#get}.
	InfoType string `json:"info_type"`
}

// NewVcenterComputePoliciesCapabilitiesInfo instantiates a new VcenterComputePoliciesCapabilitiesInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterComputePoliciesCapabilitiesInfo(name string, description string, createSpecType string, infoType string) *VcenterComputePoliciesCapabilitiesInfo {
	this := VcenterComputePoliciesCapabilitiesInfo{}
	this.Name = name
	this.Description = description
	this.CreateSpecType = createSpecType
	this.InfoType = infoType
	return &this
}

// NewVcenterComputePoliciesCapabilitiesInfoWithDefaults instantiates a new VcenterComputePoliciesCapabilitiesInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterComputePoliciesCapabilitiesInfoWithDefaults() *VcenterComputePoliciesCapabilitiesInfo {
	this := VcenterComputePoliciesCapabilitiesInfo{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterComputePoliciesCapabilitiesInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterComputePoliciesCapabilitiesInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterComputePoliciesCapabilitiesInfo) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *VcenterComputePoliciesCapabilitiesInfo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VcenterComputePoliciesCapabilitiesInfo) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VcenterComputePoliciesCapabilitiesInfo) SetDescription(v string) {
	o.Description = v
}

// GetCreateSpecType returns the CreateSpecType field value
func (o *VcenterComputePoliciesCapabilitiesInfo) GetCreateSpecType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CreateSpecType
}

// GetCreateSpecTypeOk returns a tuple with the CreateSpecType field value
// and a boolean to check if the value has been set.
func (o *VcenterComputePoliciesCapabilitiesInfo) GetCreateSpecTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreateSpecType, true
}

// SetCreateSpecType sets field value
func (o *VcenterComputePoliciesCapabilitiesInfo) SetCreateSpecType(v string) {
	o.CreateSpecType = v
}

// GetInfoType returns the InfoType field value
func (o *VcenterComputePoliciesCapabilitiesInfo) GetInfoType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InfoType
}

// GetInfoTypeOk returns a tuple with the InfoType field value
// and a boolean to check if the value has been set.
func (o *VcenterComputePoliciesCapabilitiesInfo) GetInfoTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.InfoType, true
}

// SetInfoType sets field value
func (o *VcenterComputePoliciesCapabilitiesInfo) SetInfoType(v string) {
	o.InfoType = v
}

func (o VcenterComputePoliciesCapabilitiesInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["create_spec_type"] = o.CreateSpecType
	}
	if true {
		toSerialize["info_type"] = o.InfoType
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterComputePoliciesCapabilitiesInfo struct {
	value *VcenterComputePoliciesCapabilitiesInfo
	isSet bool
}

func (v NullableVcenterComputePoliciesCapabilitiesInfo) Get() *VcenterComputePoliciesCapabilitiesInfo {
	return v.value
}

func (v *NullableVcenterComputePoliciesCapabilitiesInfo) Set(val *VcenterComputePoliciesCapabilitiesInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterComputePoliciesCapabilitiesInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterComputePoliciesCapabilitiesInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterComputePoliciesCapabilitiesInfo(val *VcenterComputePoliciesCapabilitiesInfo) *NullableVcenterComputePoliciesCapabilitiesInfo {
	return &NullableVcenterComputePoliciesCapabilitiesInfo{value: val, isSet: true}
}

func (v NullableVcenterComputePoliciesCapabilitiesInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterComputePoliciesCapabilitiesInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


