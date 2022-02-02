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

// VcenterVMRelocatePlacementSpec struct for VcenterVMRelocatePlacementSpec
type VcenterVMRelocatePlacementSpec struct {
	// Virtual machine folder into which the virtual machine should be placed. If this field is unset, the virtual machine will stay in the current folder. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder.
	Folder *string `json:"folder,omitempty"`
	// Resource pool into which the virtual machine should be placed. If this field is unset, the virtual machine will stay in the current resource pool. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool.
	ResourcePool *string `json:"resource_pool,omitempty"`
	// Host onto which the virtual machine should be placed.   If VM.RelocatePlacementSpec.host and VM.RelocatePlacementSpec.resource-pool are both specified, VM.RelocatePlacementSpec.resource-pool must belong to VM.RelocatePlacementSpec.host.    If VM.RelocatePlacementSpec.host and VM.RelocatePlacementSpec.cluster are both specified, VM.RelocatePlacementSpec.host must be a member of VM.RelocatePlacementSpec.cluster.  If this field is unset, if VM.RelocatePlacementSpec.resource-pool is unset, the virtual machine will remain on the current host. if VM.RelocatePlacementSpec.resource-pool is set, and the target is a standalone host, the host is used. if VM.RelocatePlacementSpec.resource-pool is set, and the target is a DRS cluster, a host will be picked by DRS. if VM.RelocatePlacementSpec.resource-pool is set, and the target is a cluster without DRS, InvalidArgument will be thrown. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem.
	Host *string `json:"host,omitempty"`
	// Cluster into which the virtual machine should be placed.   If VM.RelocatePlacementSpec.cluster and VM.RelocatePlacementSpec.resource-pool are both specified, VM.RelocatePlacementSpec.resource-pool must belong to VM.RelocatePlacementSpec.cluster.    If VM.RelocatePlacementSpec.cluster and VM.RelocatePlacementSpec.host are both specified, VM.RelocatePlacementSpec.host must be a member of VM.RelocatePlacementSpec.cluster.  If VM.RelocatePlacementSpec.resource-pool or VM.RelocatePlacementSpec.host is specified, it is recommended that this field be unset. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.
	Cluster *string `json:"cluster,omitempty"`
	// Datastore on which the virtual machine's configuration state should be stored. This datastore will also be used for any virtual disks that are associated with the virtual machine, unless individually overridden. If this field is unset, the virtual machine will remain on the current datastore. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore.
	Datastore *string `json:"datastore,omitempty"`
}

// NewVcenterVMRelocatePlacementSpec instantiates a new VcenterVMRelocatePlacementSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVMRelocatePlacementSpec() *VcenterVMRelocatePlacementSpec {
	this := VcenterVMRelocatePlacementSpec{}
	return &this
}

// NewVcenterVMRelocatePlacementSpecWithDefaults instantiates a new VcenterVMRelocatePlacementSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVMRelocatePlacementSpecWithDefaults() *VcenterVMRelocatePlacementSpec {
	this := VcenterVMRelocatePlacementSpec{}
	return &this
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *VcenterVMRelocatePlacementSpec) GetFolder() string {
	if o == nil || o.Folder == nil {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRelocatePlacementSpec) GetFolderOk() (*string, bool) {
	if o == nil || o.Folder == nil {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *VcenterVMRelocatePlacementSpec) HasFolder() bool {
	if o != nil && o.Folder != nil {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *VcenterVMRelocatePlacementSpec) SetFolder(v string) {
	o.Folder = &v
}

// GetResourcePool returns the ResourcePool field value if set, zero value otherwise.
func (o *VcenterVMRelocatePlacementSpec) GetResourcePool() string {
	if o == nil || o.ResourcePool == nil {
		var ret string
		return ret
	}
	return *o.ResourcePool
}

// GetResourcePoolOk returns a tuple with the ResourcePool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRelocatePlacementSpec) GetResourcePoolOk() (*string, bool) {
	if o == nil || o.ResourcePool == nil {
		return nil, false
	}
	return o.ResourcePool, true
}

// HasResourcePool returns a boolean if a field has been set.
func (o *VcenterVMRelocatePlacementSpec) HasResourcePool() bool {
	if o != nil && o.ResourcePool != nil {
		return true
	}

	return false
}

// SetResourcePool gets a reference to the given string and assigns it to the ResourcePool field.
func (o *VcenterVMRelocatePlacementSpec) SetResourcePool(v string) {
	o.ResourcePool = &v
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *VcenterVMRelocatePlacementSpec) GetHost() string {
	if o == nil || o.Host == nil {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRelocatePlacementSpec) GetHostOk() (*string, bool) {
	if o == nil || o.Host == nil {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *VcenterVMRelocatePlacementSpec) HasHost() bool {
	if o != nil && o.Host != nil {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the Host field.
func (o *VcenterVMRelocatePlacementSpec) SetHost(v string) {
	o.Host = &v
}

// GetCluster returns the Cluster field value if set, zero value otherwise.
func (o *VcenterVMRelocatePlacementSpec) GetCluster() string {
	if o == nil || o.Cluster == nil {
		var ret string
		return ret
	}
	return *o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRelocatePlacementSpec) GetClusterOk() (*string, bool) {
	if o == nil || o.Cluster == nil {
		return nil, false
	}
	return o.Cluster, true
}

// HasCluster returns a boolean if a field has been set.
func (o *VcenterVMRelocatePlacementSpec) HasCluster() bool {
	if o != nil && o.Cluster != nil {
		return true
	}

	return false
}

// SetCluster gets a reference to the given string and assigns it to the Cluster field.
func (o *VcenterVMRelocatePlacementSpec) SetCluster(v string) {
	o.Cluster = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *VcenterVMRelocatePlacementSpec) GetDatastore() string {
	if o == nil || o.Datastore == nil {
		var ret string
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVMRelocatePlacementSpec) GetDatastoreOk() (*string, bool) {
	if o == nil || o.Datastore == nil {
		return nil, false
	}
	return o.Datastore, true
}

// HasDatastore returns a boolean if a field has been set.
func (o *VcenterVMRelocatePlacementSpec) HasDatastore() bool {
	if o != nil && o.Datastore != nil {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given string and assigns it to the Datastore field.
func (o *VcenterVMRelocatePlacementSpec) SetDatastore(v string) {
	o.Datastore = &v
}

func (o VcenterVMRelocatePlacementSpec) MarshalJSON() ([]byte, error) {
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
	if o.Datastore != nil {
		toSerialize["datastore"] = o.Datastore
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVMRelocatePlacementSpec struct {
	value *VcenterVMRelocatePlacementSpec
	isSet bool
}

func (v NullableVcenterVMRelocatePlacementSpec) Get() *VcenterVMRelocatePlacementSpec {
	return v.value
}

func (v *NullableVcenterVMRelocatePlacementSpec) Set(val *VcenterVMRelocatePlacementSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVMRelocatePlacementSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVMRelocatePlacementSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVMRelocatePlacementSpec(val *VcenterVMRelocatePlacementSpec) *NullableVcenterVMRelocatePlacementSpec {
	return &NullableVcenterVMRelocatePlacementSpec{value: val, isSet: true}
}

func (v NullableVcenterVMRelocatePlacementSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVMRelocatePlacementSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


