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

// VcenterVmTemplateLibraryItemsCreatePlacementSpec struct for VcenterVmTemplateLibraryItemsCreatePlacementSpec
type VcenterVmTemplateLibraryItemsCreatePlacementSpec struct {
	// Virtual machine folder into which the virtual machine template should be placed.
	Folder *string `json:"folder,omitempty"`
	// Resource pool into which the virtual machine template should be placed.
	ResourcePool *string `json:"resource_pool,omitempty"`
	// Host onto which the virtual machine template should be placed. If {@name #host} and {@name #resourcePool} are both specified, {@name #resourcePool} must belong to {@name #host}. If {@name #host} and {@name #cluster} are both specified, {@name #host} must be a member of {@name #cluster}.
	Host *string `json:"host,omitempty"`
	// Cluster onto which the virtual machine template should be placed. If {@name #cluster} and {@name #resourcePool} are both specified, {@name #resourcePool} must belong to {@name #cluster}. If {@name #cluster} and {@name #host} are both specified, {@name #host} must be a member of {@name #cluster}.
	Cluster *string `json:"cluster,omitempty"`
}

// NewVcenterVmTemplateLibraryItemsCreatePlacementSpec instantiates a new VcenterVmTemplateLibraryItemsCreatePlacementSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmTemplateLibraryItemsCreatePlacementSpec() *VcenterVmTemplateLibraryItemsCreatePlacementSpec {
	this := VcenterVmTemplateLibraryItemsCreatePlacementSpec{}
	return &this
}

// NewVcenterVmTemplateLibraryItemsCreatePlacementSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreatePlacementSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmTemplateLibraryItemsCreatePlacementSpecWithDefaults() *VcenterVmTemplateLibraryItemsCreatePlacementSpec {
	this := VcenterVmTemplateLibraryItemsCreatePlacementSpec{}
	return &this
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) SetFolder(v string) {
	o.Folder = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetResourcePool() string {
	if o == nil || o.ResourcePool == nil {
		var ret string
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetResourcePoolOk() (*string, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given string and assigns it to the ResourcePool field.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) SetResourcePool(v string) {
	o.ResourcePool = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) SetHost(v string) {
	o.Host = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *VcenterVmTemplateLibraryItemsCreatePlacementSpec) SetCluster(v string) {
	o.Cluster = &v
}

func (o VcenterVmTemplateLibraryItemsCreatePlacementSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Folder != nil {
		toSerialize["folder"] = o.Folder
	}
	if o.ResourcePool != nil {
		toSerialize["resource_pool"] = o.ResourcePool
	}
	if o.Host != nil {
		toSerialize["host"] = o.Host
	}
	if o.Cluster != nil {
		toSerialize["cluster"] = o.Cluster
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec struct {
	value *VcenterVmTemplateLibraryItemsCreatePlacementSpec
	isSet bool
}

func (v NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec) Get() *VcenterVmTemplateLibraryItemsCreatePlacementSpec {
	return v.value
}

func (v *NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec) Set(val *VcenterVmTemplateLibraryItemsCreatePlacementSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmTemplateLibraryItemsCreatePlacementSpec(val *VcenterVmTemplateLibraryItemsCreatePlacementSpec) *NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec {
	return &NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec{value: val, isSet: true}
}

func (v NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmTemplateLibraryItemsCreatePlacementSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

