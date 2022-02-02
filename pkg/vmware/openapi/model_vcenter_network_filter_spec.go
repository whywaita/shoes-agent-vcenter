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

// VcenterNetworkFilterSpec struct for VcenterNetworkFilterSpec
type VcenterNetworkFilterSpec struct {
	// Identifiers of networks that can match the filter. If unset or empty, networks with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Network. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Network.
	Networks *[]string `json:"networks,omitempty"`
	// Names that networks must have to match the filter (see Network.Summary.name). If unset or empty, networks with any name match the filter.
	Names *[]string `json:"names,omitempty"`
	// Types that networks must have to match the filter (see Network.Summary.type). If unset, networks with any type match the filter.
	Types *[]VcenterNetworkType `json:"types,omitempty"`
	// Folders that must contain the network for the network to match the filter. If unset or empty, networks in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder.
	Folders *[]string `json:"folders,omitempty"`
	// Datacenters that must contain the network for the network to match the filter. If unset or empty, networks in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter.
	Datacenters *[]string `json:"datacenters,omitempty"`
}

// NewVcenterNetworkFilterSpec instantiates a new VcenterNetworkFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNetworkFilterSpec() *VcenterNetworkFilterSpec {
	this := VcenterNetworkFilterSpec{}
	return &this
}

// NewVcenterNetworkFilterSpecWithDefaults instantiates a new VcenterNetworkFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNetworkFilterSpecWithDefaults() *VcenterNetworkFilterSpec {
	this := VcenterNetworkFilterSpec{}
	return &this
}

// GetNetworks returns the Networks field value if set, zero value otherwise.
func (o *VcenterNetworkFilterSpec) GetNetworks() []string {
	if o == nil || o.Networks == nil {
		var ret []string
		return ret
	}
	return *o.Networks
}

// GetNetworksOk returns a tuple with the Networks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNetworkFilterSpec) GetNetworksOk() (*[]string, bool) {
	if o == nil || o.Networks == nil {
		return nil, false
	}
	return o.Networks, true
}

// HasNetworks returns a boolean if a field has been set.
func (o *VcenterNetworkFilterSpec) HasNetworks() bool {
	if o != nil && o.Networks != nil {
		return true
	}

	return false
}

// SetNetworks gets a reference to the given []string and assigns it to the Networks field.
func (o *VcenterNetworkFilterSpec) SetNetworks(v []string) {
	o.Networks = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *VcenterNetworkFilterSpec) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNetworkFilterSpec) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *VcenterNetworkFilterSpec) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *VcenterNetworkFilterSpec) SetNames(v []string) {
	o.Names = &v
}

// GetTypes returns the Types field value if set, zero value otherwise.
func (o *VcenterNetworkFilterSpec) GetTypes() []VcenterNetworkType {
	if o == nil || o.Types == nil {
		var ret []VcenterNetworkType
		return ret
	}
	return *o.Types
}

// GetTypesOk returns a tuple with the Types field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNetworkFilterSpec) GetTypesOk() (*[]VcenterNetworkType, bool) {
	if o == nil || o.Types == nil {
		return nil, false
	}
	return o.Types, true
}

// HasTypes returns a boolean if a field has been set.
func (o *VcenterNetworkFilterSpec) HasTypes() bool {
	if o != nil && o.Types != nil {
		return true
	}

	return false
}

// SetTypes gets a reference to the given []VcenterNetworkType and assigns it to the Types field.
func (o *VcenterNetworkFilterSpec) SetTypes(v []VcenterNetworkType) {
	o.Types = &v
}

// GetFolders returns the Folders field value if set, zero value otherwise.
func (o *VcenterNetworkFilterSpec) GetFolders() []string {
	if o == nil || o.Folders == nil {
		var ret []string
		return ret
	}
	return *o.Folders
}

// GetFoldersOk returns a tuple with the Folders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNetworkFilterSpec) GetFoldersOk() (*[]string, bool) {
	if o == nil || o.Folders == nil {
		return nil, false
	}
	return o.Folders, true
}

// HasFolders returns a boolean if a field has been set.
func (o *VcenterNetworkFilterSpec) HasFolders() bool {
	if o != nil && o.Folders != nil {
		return true
	}

	return false
}

// SetFolders gets a reference to the given []string and assigns it to the Folders field.
func (o *VcenterNetworkFilterSpec) SetFolders(v []string) {
	o.Folders = &v
}

// GetDatacenters returns the Datacenters field value if set, zero value otherwise.
func (o *VcenterNetworkFilterSpec) GetDatacenters() []string {
	if o == nil || o.Datacenters == nil {
		var ret []string
		return ret
	}
	return *o.Datacenters
}

// GetDatacentersOk returns a tuple with the Datacenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNetworkFilterSpec) GetDatacentersOk() (*[]string, bool) {
	if o == nil || o.Datacenters == nil {
		return nil, false
	}
	return o.Datacenters, true
}

// HasDatacenters returns a boolean if a field has been set.
func (o *VcenterNetworkFilterSpec) HasDatacenters() bool {
	if o != nil && o.Datacenters != nil {
		return true
	}

	return false
}

// SetDatacenters gets a reference to the given []string and assigns it to the Datacenters field.
func (o *VcenterNetworkFilterSpec) SetDatacenters(v []string) {
	o.Datacenters = &v
}

func (o VcenterNetworkFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Networks != nil {
		toSerialize["networks"] = o.Networks
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

type NullableVcenterNetworkFilterSpec struct {
	value *VcenterNetworkFilterSpec
	isSet bool
}

func (v NullableVcenterNetworkFilterSpec) Get() *VcenterNetworkFilterSpec {
	return v.value
}

func (v *NullableVcenterNetworkFilterSpec) Set(val *VcenterNetworkFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNetworkFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNetworkFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNetworkFilterSpec(val *VcenterNetworkFilterSpec) *NullableVcenterNetworkFilterSpec {
	return &NullableVcenterNetworkFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterNetworkFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNetworkFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

