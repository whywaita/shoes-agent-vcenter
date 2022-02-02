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

// VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy struct for VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy
type VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy struct {
	Type VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType `json:"type"`
	// Identifier for the storage policy to use.
	Policy *string `json:"policy,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy instantiates a new VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy(type_ VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy {
	this := VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy{}
	this.Type = type_
	return &this
}

// NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyWithDefaults() *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy {
	this := VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) GetType() VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType {
	if o == nil {
		var ret VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) GetTypeOk() (*VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) SetType(v VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicyType) {
	o.Type = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) GetPolicy() string {
	if o == nil || o.Policy == nil {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) GetPolicyOk() (*string, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) SetPolicy(v string) {
	o.Policy = &v
}

func (o VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy struct {
	value *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) Get() *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) Set(val *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy(val *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy {
	return &NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

