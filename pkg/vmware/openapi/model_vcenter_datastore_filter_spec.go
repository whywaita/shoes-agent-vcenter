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

// VcenterDatastoreFilterSpec struct for VcenterDatastoreFilterSpec
type VcenterDatastoreFilterSpec struct {
	// Identifiers of datastores that can match the filter. If unset or empty, datastores with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datastore. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datastore.
	Datastores *[]string `json:"datastores,omitempty"`
	// Names that datastores must have to match the filter (see Datastore.Info.name). If unset or empty, datastores with any name match the filter.
	Names *[]string `json:"names,omitempty"`
	// Types that datastores must have to match the filter (see Datastore.Summary.type). If unset or empty, datastores with any type match the filter.
	Types *[]VcenterDatastoreType `json:"types,omitempty"`
	// Folders that must contain the datastore for the datastore to match the filter. If unset or empty, datastores in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder.
	Folders *[]string `json:"folders,omitempty"`
	// Datacenters that must contain the datastore for the datastore to match the filter. If unset or empty, datastores in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter.
	Datacenters *[]string `json:"datacenters,omitempty"`
}

// NewVcenterDatastoreFilterSpec instantiates a new VcenterDatastoreFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDatastoreFilterSpec() *VcenterDatastoreFilterSpec {
	this := VcenterDatastoreFilterSpec{}
	return &this
}

// NewVcenterDatastoreFilterSpecWithDefaults instantiates a new VcenterDatastoreFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDatastoreFilterSpecWithDefaults() *VcenterDatastoreFilterSpec {
	this := VcenterDatastoreFilterSpec{}
	return &this
}

// GetDatastores returns the Datastores field value if set, zero value otherwise.
func (o *VcenterDatastoreFilterSpec) GetDatastores() []string {
	if o == nil || o.Datastores == nil {
		var ret []string
		return ret
	}
	return *o.Datastores
}

// GetDatastoresOk returns a tuple with the Datastores field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreFilterSpec) GetDatastoresOk() (*[]string, bool) {
	if o == nil || o.Datastores == nil {
		return nil, false
	}
	return o.Datastores, true
}

// HasDatastores returns a boolean if a field has been set.
func (o *VcenterDatastoreFilterSpec) HasDatastores() bool {
	if o != nil && o.Datastores != nil {
		return true
	}

	return false
}

// SetDatastores gets a reference to the given []string and assigns it to the Datastores field.
func (o *VcenterDatastoreFilterSpec) SetDatastores(v []string) {
	o.Datastores = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *VcenterDatastoreFilterSpec) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreFilterSpec) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *VcenterDatastoreFilterSpec) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *VcenterDatastoreFilterSpec) SetNames(v []string) {
	o.Names = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *VcenterDatastoreFilterSpec) GetTypes() []VcenterDatastoreType {
	if o == nil || o.Types == nil {
		var ret []VcenterDatastoreType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreFilterSpec) GetTypesOk() (*[]VcenterDatastoreType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *VcenterDatastoreFilterSpec) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []VcenterDatastoreType and assigns it to the Types field.
func (o *VcenterDatastoreFilterSpec) SetTypes(v []VcenterDatastoreType) {
	o.Types = &v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *VcenterDatastoreFilterSpec) GetFolders() []string {
	if o == nil || o.Folders == nil {
		var ret []string
		return ret
	}
	return *o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreFilterSpec) GetFoldersOk() (*[]string, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *VcenterDatastoreFilterSpec) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []string and assigns it to the Folders field.
func (o *VcenterDatastoreFilterSpec) SetFolders(v []string) {
	o.Folders = &v
}

// GetDatacenters returns the Datacenters field value if set, zero value otherwise.
func (o *VcenterDatastoreFilterSpec) GetDatacenters() []string {
	if o == nil || o.Datacenters == nil {
		var ret []string
		return ret
	}
	return *o.Datacenters
}

// GetDatacentersOk returns a tuple with the Datacenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterDatastoreFilterSpec) GetDatacentersOk() (*[]string, bool) {
	if o == nil || o.Datacenters == nil {
		return nil, false
	}
	return o.Datacenters, true
}

// HasDatacenters returns a boolean if a field has been set.
func (o *VcenterDatastoreFilterSpec) HasDatacenters() bool {
	if o != nil && o.Datacenters != nil {
		return true
	}

	return false
}

// SetDatacenters gets a reference to the given []string and assigns it to the Datacenters field.
func (o *VcenterDatastoreFilterSpec) SetDatacenters(v []string) {
	o.Datacenters = &v
}

func (o VcenterDatastoreFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Datastores != nil {
		toSerialize["datastores"] = o.Datastores
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.Types != nil {
		toSerialize["types"] = o.Types
	}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	if o.Datacenters != nil {
		toSerialize["datacenters"] = o.Datacenters
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDatastoreFilterSpec struct {
	value *VcenterDatastoreFilterSpec
	isSet bool
}

func (v NullableVcenterDatastoreFilterSpec) Get() *VcenterDatastoreFilterSpec {
	return v.value
}

func (v *NullableVcenterDatastoreFilterSpec) Set(val *VcenterDatastoreFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDatastoreFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDatastoreFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDatastoreFilterSpec(val *VcenterDatastoreFilterSpec) *NullableVcenterDatastoreFilterSpec {
	return &NullableVcenterDatastoreFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterDatastoreFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDatastoreFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


