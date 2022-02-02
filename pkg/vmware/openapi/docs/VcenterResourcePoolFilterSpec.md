# VcenterResourcePoolFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcePools** | Pointer to **[]string** | Identifiers of resource pools that can match the filter. If unset or empty, resource pools with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool. | [optional] 
**Names** | Pointer to **[]string** | Names that resource pools must have to match the filter (see ResourcePool.Info.name). If unset or empty, resource pools with any name match the filter. | [optional] 
**ParentResourcePools** | Pointer to **[]string** | Resource pools that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any resource pool match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 
**Hosts** | Pointer to **[]string** | Hosts that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any host match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: HostSystem. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: HostSystem. | [optional] 
**Clusters** | Pointer to **[]string** | Clusters that must contain the resource pool for the resource pool to match the filter. If unset or empty, resource pools in any cluster match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 

## Methods

### NewVcenterResourcePoolFilterSpec

`func NewVcenterResourcePoolFilterSpec() *VcenterResourcePoolFilterSpec`

NewVcenterResourcePoolFilterSpec instantiates a new VcenterResourcePoolFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolFilterSpecWithDefaults

`func NewVcenterResourcePoolFilterSpecWithDefaults() *VcenterResourcePoolFilterSpec`

NewVcenterResourcePoolFilterSpecWithDefaults instantiates a new VcenterResourcePoolFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcePools

`func (o *VcenterResourcePoolFilterSpec) GetResourcePools() []string`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *VcenterResourcePoolFilterSpec) GetResourcePoolsOk() (*[]string, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *VcenterResourcePoolFilterSpec) SetResourcePools(v []string)`

SetResourcePools sets ResourcePools field to given value.

### HasResourcePools

`func (o *VcenterResourcePoolFilterSpec) HasResourcePools() bool`

HasResourcePools returns a boolean if a field has been set.

### GetNames

`func (o *VcenterResourcePoolFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterResourcePoolFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterResourcePoolFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterResourcePoolFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetParentResourcePools

`func (o *VcenterResourcePoolFilterSpec) GetParentResourcePools() []string`

GetParentResourcePools returns the ParentResourcePools field if non-nil, zero value otherwise.

### GetParentResourcePoolsOk

`func (o *VcenterResourcePoolFilterSpec) GetParentResourcePoolsOk() (*[]string, bool)`

GetParentResourcePoolsOk returns a tuple with the ParentResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourcePools

`func (o *VcenterResourcePoolFilterSpec) SetParentResourcePools(v []string)`

SetParentResourcePools sets ParentResourcePools field to given value.

### HasParentResourcePools

`func (o *VcenterResourcePoolFilterSpec) HasParentResourcePools() bool`

HasParentResourcePools returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterResourcePoolFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterResourcePoolFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterResourcePoolFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterResourcePoolFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetHosts

`func (o *VcenterResourcePoolFilterSpec) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VcenterResourcePoolFilterSpec) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VcenterResourcePoolFilterSpec) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VcenterResourcePoolFilterSpec) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetClusters

`func (o *VcenterResourcePoolFilterSpec) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VcenterResourcePoolFilterSpec) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VcenterResourcePoolFilterSpec) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VcenterResourcePoolFilterSpec) HasClusters() bool`

HasClusters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


