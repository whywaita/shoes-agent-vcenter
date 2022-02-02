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

// VcenterStoragePoliciesCheckCompatibilityStoragePolicies struct for VcenterStoragePoliciesCheckCompatibilityStoragePolicies
type VcenterStoragePoliciesCheckCompatibilityStoragePolicies struct {
	// Datastores used to check compatibility against a storage policy. The number of datastores is limited to 1024. The parameter must contain identifiers for the resource type: Datastore.
	Datastores *[]string `json:"datastores,omitempty"`
}

// NewVcenterStoragePoliciesCheckCompatibilityStoragePolicies instantiates a new VcenterStoragePoliciesCheckCompatibilityStoragePolicies object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterStoragePoliciesCheckCompatibilityStoragePolicies() *VcenterStoragePoliciesCheckCompatibilityStoragePolicies {
	this := VcenterStoragePoliciesCheckCompatibilityStoragePolicies{}
	return &this
}

// NewVcenterStoragePoliciesCheckCompatibilityStoragePoliciesWithDefaults instantiates a new VcenterStoragePoliciesCheckCompatibilityStoragePolicies object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterStoragePoliciesCheckCompatibilityStoragePoliciesWithDefaults() *VcenterStoragePoliciesCheckCompatibilityStoragePolicies {
	this := VcenterStoragePoliciesCheckCompatibilityStoragePolicies{}
	return &this
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) GetDatastores() []string {
	if o == nil || o.Datastores == nil {
		var ret []string
		return ret
	}
	return *o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) GetDatastoresOk() (*[]string, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []string and assigns it to the Datastores field.
func (o *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) SetDatastores(v []string) {
	o.Datastores = &v
}

func (o VcenterStoragePoliciesCheckCompatibilityStoragePolicies) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastores != nil {
		toSerialize["datastores"] = o.Datastores
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies struct {
	value *VcenterStoragePoliciesCheckCompatibilityStoragePolicies
	isSet bool
}

func (v NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies) Get() *VcenterStoragePoliciesCheckCompatibilityStoragePolicies {
	return v.value
}

func (v *NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies) Set(val *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies(val *VcenterStoragePoliciesCheckCompatibilityStoragePolicies) *NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies {
	return &NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies{value: val, isSet: true}
}

func (v NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterStoragePoliciesCheckCompatibilityStoragePolicies) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


