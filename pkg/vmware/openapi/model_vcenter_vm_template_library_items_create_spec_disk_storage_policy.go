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

// VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy struct for VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy
type VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy struct {
	Type VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyType `json:"type"`
	// Identifier for the storage policy to use.
	Policy *string `json:"policy,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy instantiates a new VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy(type_ VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyType) *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy {
	this := VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy{}
	this.Type = type_
	return &this
}

// NewVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyWithDefaults() *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy {
	this := VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy{}
	return &this
}

// GetType returns the Type field value
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) GetType() VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyType {
	if o == nil {
		var ret VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) GetTypeOk() (*VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) SetType(v VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicyType) {
	o.Type = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) GetPolicy() string {
	if o == nil || o.Policy == nil {
		var ret string
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) GetPolicyOk() (*string, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given string and assigns it to the Policy field.
func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) SetPolicy(v string) {
	o.Policy = &v
}

func (o VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy struct {
	value *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) Get() *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) Set(val *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy(val *VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy {
	return &NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


