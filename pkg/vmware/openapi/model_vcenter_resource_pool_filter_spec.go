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

// VcenterResourcePoolFilterSpec struct for VcenterResourcePoolFilterSpec
type VcenterResourcePoolFilterSpec struct {
	// Identifiers of resource pools that can match the filter. If unset or empty, resource pools with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool.
	ResourcePools *[]string `json:"resource_pools,omitempty"`
	// Names that resource pools must have to match the filter (see ResourcePool.Info.name). If unset or empty, resource pools with any name match the filter.
	Names *[]string `json:"names,omitempty"`
	// Resource pools that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any resource pool match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool.
	ParentResourcePools *[]string `json:"parent_resource_pools,omitempty"`
	// Datacenters that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter.
	Datacenters *[]string `json:"datacenters,omitempty"`
	// Hosts that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any host match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: HostSystem. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: HostSystem.
	Hosts *[]string `json:"hosts,omitempty"`
	// Clusters that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any cluster match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource.
	Clusters *[]string `json:"clusters,omitempty"`
}

// NewVcenterResourcePoolFilterSpec instantiates a new VcenterResourcePoolFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterResourcePoolFilterSpec() *VcenterResourcePoolFilterSpec {
	this := VcenterResourcePoolFilterSpec{}
	return &this
}

// NewVcenterResourcePoolFilterSpecWithDefaults instantiates a new VcenterResourcePoolFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterResourcePoolFilterSpecWithDefaults() *VcenterResourcePoolFilterSpec {
	this := VcenterResourcePoolFilterSpec{}
	return &this
}

// GetResourcePools returns the ResourcePools field value if set, zero value otherwise.
func (o *VcenterResourcePoolFilterSpec) GetResourcePools() []string {
	if o == nil || o.ResourcePools == nil {
		var ret []string
		return ret
	}
	return *o.ResourcePools
}

// GetResourcePoolsOk returns a tuple with the ResourcePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolFilterSpec) GetResourcePoolsOk() (*[]string, bool) {
	if o == nil || o.ResourcePools == nil {
		return nil, false
	}
	return o.ResourcePools, true
}

// HasResourcePools returns a boolean if a field has been set.
func (o *VcenterResourcePoolFilterSpec) HasResourcePools() bool {
	if o != nil && o.ResourcePools != nil {
		return true
	}

	return false
}

// SetResourcePools gets a reference to the given []string and assigns it to the ResourcePools field.
func (o *VcenterResourcePoolFilterSpec) SetResourcePools(v []string) {
	o.ResourcePools = &v
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *VcenterResourcePoolFilterSpec) GetNames() []string {
	if o == nil || o.Names == nil {
		var ret []string
		return ret
	}
	return *o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolFilterSpec) GetNamesOk() (*[]string, bool) {
	if o == nil || o.Names == nil {
		return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *VcenterResourcePoolFilterSpec) HasNames() bool {
	if o != nil && o.Names != nil {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *VcenterResourcePoolFilterSpec) SetNames(v []string) {
	o.Names = &v
}

// GetParentResourcePools returns the ParentResourcePools field value if set, zero value otherwise.
func (o *VcenterResourcePoolFilterSpec) GetParentResourcePools() []string {
	if o == nil || o.ParentResourcePools == nil {
		var ret []string
		return ret
	}
	return *o.ParentResourcePools
}

// GetParentResourcePoolsOk returns a tuple with the ParentResourcePools field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolFilterSpec) GetParentResourcePoolsOk() (*[]string, bool) {
	if o == nil || o.ParentResourcePools == nil {
		return nil, false
	}
	return o.ParentResourcePools, true
}

// HasParentResourcePools returns a boolean if a field has been set.
func (o *VcenterResourcePoolFilterSpec) HasParentResourcePools() bool {
	if o != nil && o.ParentResourcePools != nil {
		return true
	}

	return false
}

// SetParentResourcePools gets a reference to the given []string and assigns it to the ParentResourcePools field.
func (o *VcenterResourcePoolFilterSpec) SetParentResourcePools(v []string) {
	o.ParentResourcePools = &v
}

// GetDatacenters returns the Datacenters field value if set, zero value otherwise.
func (o *VcenterResourcePoolFilterSpec) GetDatacenters() []string {
	if o == nil || o.Datacenters == nil {
		var ret []string
		return ret
	}
	return *o.Datacenters
}

// GetDatacentersOk returns a tuple with the Datacenters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolFilterSpec) GetDatacentersOk() (*[]string, bool) {
	if o == nil || o.Datacenters == nil {
		return nil, false
	}
	return o.Datacenters, true
}

// HasDatacenters returns a boolean if a field has been set.
func (o *VcenterResourcePoolFilterSpec) HasDatacenters() bool {
	if o != nil && o.Datacenters != nil {
		return true
	}

	return false
}

// SetDatacenters gets a reference to the given []string and assigns it to the Datacenters field.
func (o *VcenterResourcePoolFilterSpec) SetDatacenters(v []string) {
	o.Datacenters = &v
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *VcenterResourcePoolFilterSpec) GetHosts() []string {
	if o == nil || o.Hosts == nil {
		var ret []string
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolFilterSpec) GetHostsOk() (*[]string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VcenterResourcePoolFilterSpec) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *VcenterResourcePoolFilterSpec) SetHosts(v []string) {
	o.Hosts = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *VcenterResourcePoolFilterSpec) GetClusters() []string {
	if o == nil || o.Clusters == nil {
		var ret []string
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolFilterSpec) GetClustersOk() (*[]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *VcenterResourcePoolFilterSpec) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *VcenterResourcePoolFilterSpec) SetClusters(v []string) {
	o.Clusters = &v
}

func (o VcenterResourcePoolFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourcePools != nil {
		toSerialize["resource_pools"] = o.ResourcePools
	}
	if o.Names != nil {
		toSerialize["names"] = o.Names
	}
	if o.ParentResourcePools != nil {
		toSerialize["parent_resource_pools"] = o.ParentResourcePools
	}
	if o.Datacenters != nil {
		toSerialize["datacenters"] = o.Datacenters
	}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterResourcePoolFilterSpec struct {
	value *VcenterResourcePoolFilterSpec
	isSet bool
}

func (v NullableVcenterResourcePoolFilterSpec) Get() *VcenterResourcePoolFilterSpec {
	return v.value
}

func (v *NullableVcenterResourcePoolFilterSpec) Set(val *VcenterResourcePoolFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterResourcePoolFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterResourcePoolFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterResourcePoolFilterSpec(val *VcenterResourcePoolFilterSpec) *NullableVcenterResourcePoolFilterSpec {
	return &NullableVcenterResourcePoolFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterResourcePoolFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterResourcePoolFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


