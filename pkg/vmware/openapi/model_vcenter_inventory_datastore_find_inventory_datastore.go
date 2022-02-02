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

// VcenterInventoryDatastoreFindInventoryDatastore struct for VcenterInventoryDatastoreFindInventoryDatastore
type VcenterInventoryDatastoreFindInventoryDatastore struct {
	// Identifiers of the datastores for which information will be returned.
	Datastores *[]string `json:"datastores,omitempty"`
}

// NewVcenterInventoryDatastoreFindInventoryDatastore instantiates a new VcenterInventoryDatastoreFindInventoryDatastore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterInventoryDatastoreFindInventoryDatastore() *VcenterInventoryDatastoreFindInventoryDatastore {
	this := VcenterInventoryDatastoreFindInventoryDatastore{}
	return &this
}

// NewVcenterInventoryDatastoreFindInventoryDatastoreWithDefaults instantiates a new VcenterInventoryDatastoreFindInventoryDatastore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterInventoryDatastoreFindInventoryDatastoreWithDefaults() *VcenterInventoryDatastoreFindInventoryDatastore {
	this := VcenterInventoryDatastoreFindInventoryDatastore{}
	return &this
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *VcenterInventoryDatastoreFindInventoryDatastore) GetDatastores() []string {
	if o == nil || o.Datastores == nil {
		var ret []string
		return ret
	}
	return *o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterInventoryDatastoreFindInventoryDatastore) GetDatastoresOk() (*[]string, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *VcenterInventoryDatastoreFindInventoryDatastore) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []string and assigns it to the Datastores field.
func (o *VcenterInventoryDatastoreFindInventoryDatastore) SetDatastores(v []string) {
	o.Datastores = &v
}

func (o VcenterInventoryDatastoreFindInventoryDatastore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastores != nil {
		toSerialize["datastores"] = o.Datastores
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterInventoryDatastoreFindInventoryDatastore struct {
	value *VcenterInventoryDatastoreFindInventoryDatastore
	isSet bool
}

func (v NullableVcenterInventoryDatastoreFindInventoryDatastore) Get() *VcenterInventoryDatastoreFindInventoryDatastore {
	return v.value
}

func (v *NullableVcenterInventoryDatastoreFindInventoryDatastore) Set(val *VcenterInventoryDatastoreFindInventoryDatastore) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterInventoryDatastoreFindInventoryDatastore) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterInventoryDatastoreFindInventoryDatastore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterInventoryDatastoreFindInventoryDatastore(val *VcenterInventoryDatastoreFindInventoryDatastore) *NullableVcenterInventoryDatastoreFindInventoryDatastore {
	return &NullableVcenterInventoryDatastoreFindInventoryDatastore{value: val, isSet: true}
}

func (v NullableVcenterInventoryDatastoreFindInventoryDatastore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterInventoryDatastoreFindInventoryDatastore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


