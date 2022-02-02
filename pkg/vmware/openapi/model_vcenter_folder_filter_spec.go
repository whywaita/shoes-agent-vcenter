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

// VcenterFolderFilterSpec struct for VcenterFolderFilterSpec
type VcenterFolderFilterSpec struct {
	// Identifiers of folders that can match the filter. If unset or empty, folders with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder.
	Folders *[]string `json:"folders,omitempty"`
	// Names that folders must have to match the filter (see Folder.Summary.name). If unset or empty, folders with any name match the filter.
	Names *[]string `json:"names,omitempty"`
	Type *VcenterFolderType `json:"type,omitempty"`
	// Folders that must contain the folder for the folder to match the filter. If unset or empty, folder in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder.
	ParentFolders *[]string `json:"parent_folders,omitempty"`
	// Datacenters that must contain the folder for the folder to match the filter. If unset or empty, folder in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter.
	Datacenters *[]string `json:"datacenters,omitempty"`
}

// NewVcenterFolderFilterSpec instantiates a new VcenterFolderFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterFolderFilterSpec() *VcenterFolderFilterSpec {
	this := VcenterFolderFilterSpec{}
	return &this
}

// NewVcenterFolderFilterSpecWithDefaults instantiates a new VcenterFolderFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterFolderFilterSpecWithDefaults() *VcenterFolderFilterSpec {
	this := VcenterFolderFilterSpec{}
	return &this
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *VcenterFolderFilterSpec) GetFolders() []string {
	if o == nil || o.Folders == nil {
		var ret []string
		return ret
	}
	return *o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterFolderFilterSpec) GetFoldersOk() (*[]string, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *VcenterFolderFilterSpec) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []string and assigns it to the Folders field.
func (o *VcenterFolderFilterSpec) SetFolders(v []string) {
	o.Folders = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *VcenterFolderFilterSpec) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterFolderFilterSpec) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *VcenterFolderFilterSpec) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *VcenterFolderFilterSpec) SetNames(v []string) {
	o.Names = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VcenterFolderFilterSpec) GetType() VcenterFolderType {
	if o == nil || o.Type == nil {
		var ret VcenterFolderType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterFolderFilterSpec) GetTypeOk() (*VcenterFolderType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VcenterFolderFilterSpec) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given VcenterFolderType and assigns it to the Type field.
func (o *VcenterFolderFilterSpec) SetType(v VcenterFolderType) {
	o.Type = &v
}

// GetParentFolders returns the ParentFolders field value if set, zero value otherwise.
func (o *VcenterFolderFilterSpec) GetParentFolders() []string {
	if o == nil || o.ParentFolders == nil {
		var ret []string
		return ret
	}
	return *o.ParentFolders
}

// GetParentFoldersOk returns a tuple with the ParentFolders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterFolderFilterSpec) GetParentFoldersOk() (*[]string, bool) {
	if o == nil || o.ParentFolders == nil {
		return nil, false
	}
	return o.ParentFolders, true
}

// HasParentFolders returns a boolean if a field has been set.
func (o *VcenterFolderFilterSpec) HasParentFolders() bool {
	if o != nil && o.ParentFolders != nil {
		return true
	}

	return false
}

// SetParentFolders gets a reference to the given []string and assigns it to the ParentFolders field.
func (o *VcenterFolderFilterSpec) SetParentFolders(v []string) {
	o.ParentFolders = &v
}

// GetDatacenters returns the Datacenters field value if set, zero value otherwise.
func (o *VcenterFolderFilterSpec) GetDatacenters() []string {
	if o == nil || o.Datacenters == nil {
		var ret []string
		return ret
	}
	return *o.Datacenters
}

// GetDatacentersOk returns a tuple with the Datacenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterFolderFilterSpec) GetDatacentersOk() (*[]string, bool) {
	if o == nil || o.Datacenters == nil {
		return nil, false
	}
	return o.Datacenters, true
}

// HasDatacenters returns a boolean if a field has been set.
func (o *VcenterFolderFilterSpec) HasDatacenters() bool {
	if o != nil && o.Datacenters != nil {
		return true
	}

	return false
}

// SetDatacenters gets a reference to the given []string and assigns it to the Datacenters field.
func (o *VcenterFolderFilterSpec) SetDatacenters(v []string) {
	o.Datacenters = &v
}

func (o VcenterFolderFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Folders != nil {
		toSerialize["folders"] = o.Folders
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.ParentFolders != nil {
		toSerialize["parent_folders"] = o.ParentFolders
	}
	if o.Datacenters != nil {
		toSerialize["datacenters"] = o.Datacenters
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterFolderFilterSpec struct {
	value *VcenterFolderFilterSpec
	isSet bool
}

func (v NullableVcenterFolderFilterSpec) Get() *VcenterFolderFilterSpec {
	return v.value
}

func (v *NullableVcenterFolderFilterSpec) Set(val *VcenterFolderFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterFolderFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterFolderFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterFolderFilterSpec(val *VcenterFolderFilterSpec) *NullableVcenterFolderFilterSpec {
	return &NullableVcenterFolderFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterFolderFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterFolderFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

