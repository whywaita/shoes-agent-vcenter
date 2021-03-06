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

// VcenterVmTemplateLibraryItemsDeploySpecDiskStorage struct for VcenterVmTemplateLibraryItemsDeploySpecDiskStorage
type VcenterVmTemplateLibraryItemsDeploySpecDiskStorage struct {
	// Identifier for the datastore associated the deployed virtual machine's disk.
	Datastore *string `json:"datastore,omitempty"`
	StoragePolicy *VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy `json:"storage_policy,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorage instantiates a new VcenterVmTemplateLibraryItemsDeploySpecDiskStorage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorage() *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage {
	this := VcenterVmTemplateLibraryItemsDeploySpecDiskStorage{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorageWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDeploySpecDiskStorage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorageWithDefaults() *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage {
	this := VcenterVmTemplateLibraryItemsDeploySpecDiskStorage{}
	return &this
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) SetDatastore(v string) {
	o.Datastore = &v
}

// GetStoragePolicy returns the StoragePolicy field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetStoragePolicy() VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy {
	if o == nil || o.StoragePolicy == nil {
		var ret VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy
		return ret
	}
	return *o.StoragePolicy
}

// GetStoragePolicyOk returns a tuple with the StoragePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetStoragePolicyOk() (*VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy, bool) {
	if o == nil || o.StoragePolicy == nil {
		return nil, false
	}
	return o.StoragePolicy, true
}

// HasStoragePolicy returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) HasStoragePolicy() bool {
	if o != nil && o.StoragePolicy != nil {
		return true
	}

	return false
}

// SetStoragePolicy gets a reference to the given VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy and assigns it to the StoragePolicy field.
func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) SetStoragePolicy(v VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy) {
	o.StoragePolicy = &v
}

func (o VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	if o.StoragePolicy != nil {
		toSerialize["storage_policy"] = o.StoragePolicy
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage struct {
	value *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage) Get() *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage) Set(val *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage(val *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) *NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage {
	return &NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsDeploySpecDiskStorage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


