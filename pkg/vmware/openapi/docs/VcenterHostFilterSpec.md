# VcenterHostFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** | Identifiers of hosts that can match the filter. If unset or empty, hosts with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: HostSystem. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: HostSystem. | [optional] 
**Names** | Pointer to **[]string** | Names that hosts must have to match the filter (see Host.Summary.name). If unset or empty, hosts with any name match the filter. | [optional] 
**Folders** | Pointer to **[]string** | Folders that must contain the hosts for the hosts to match the filter. If unset or empty, hosts in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the hosts for the hosts to match the filter. If unset or empty, hosts in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 
**Standalone** | Pointer to **bool** | If true, only hosts that are not part of a cluster can match the filter, and if false, only hosts that are are part of a cluster can match the filter. If unset Hosts can match filter independent of whether they are part of a cluster or not. If this field is true and Host.FilterSpec.clusters os not empty, no hosts will match the filter. | [optional] 
**Clusters** | Pointer to **[]string** | Clusters that must contain the hosts for the hosts to match the filter. If unset or empty, hosts in any cluster and hosts that are not in a cluster match the filter. If this field is not empty and Host.FilterSpec.standalone is true, no hosts will match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**ConnectionStates** | Pointer to [**[]VcenterHostConnectionState**](VcenterHostConnectionState.md) | Connection states that a host must be in to match the filter (see Host.Summary.connection-state. If unset or empty, hosts in any connection state match the filter. | [optional] 

## Methods

### NewVcenterHostFilterSpec

`func NewVcenterHostFilterSpec() *VcenterHostFilterSpec`

NewVcenterHostFilterSpec instantiates a new VcenterHostFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHostFilterSpecWithDefaults

`func NewVcenterHostFilterSpecWithDefaults() *VcenterHostFilterSpec`

NewVcenterHostFilterSpecWithDefaults instantiates a new VcenterHostFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *VcenterHostFilterSpec) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VcenterHostFilterSpec) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VcenterHostFilterSpec) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VcenterHostFilterSpec) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetNames

`func (o *VcenterHostFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterHostFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterHostFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterHostFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetFolders

`func (o *VcenterHostFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterHostFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterHostFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterHostFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterHostFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterHostFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterHostFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterHostFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.

### GetStandalone

`func (o *VcenterHostFilterSpec) GetStandalone() bool`

GetStandalone returns the Standalone field if non-nil, zero value otherwise.

### GetStandaloneOk

`func (o *VcenterHostFilterSpec) GetStandaloneOk() (*bool, bool)`

GetStandaloneOk returns a tuple with the Standalone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandalone

`func (o *VcenterHostFilterSpec) SetStandalone(v bool)`

SetStandalone sets Standalone field to given value.

### HasStandalone

`func (o *VcenterHostFilterSpec) HasStandalone() bool`

HasStandalone returns a boolean if a field has been set.

### GetClusters

`func (o *VcenterHostFilterSpec) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VcenterHostFilterSpec) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VcenterHostFilterSpec) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VcenterHostFilterSpec) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetConnectionStates

`func (o *VcenterHostFilterSpec) GetConnectionStates() []VcenterHostConnectionState`

GetConnectionStates returns the ConnectionStates field if non-nil, zero value otherwise.

### GetConnectionStatesOk

`func (o *VcenterHostFilterSpec) GetConnectionStatesOk() (*[]VcenterHostConnectionState, bool)`

GetConnectionStatesOk returns a tuple with the ConnectionStates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionStates

`func (o *VcenterHostFilterSpec) SetConnectionStates(v []VcenterHostConnectionState)`

SetConnectionStates sets ConnectionStates field to given value.

### HasConnectionStates

`func (o *VcenterHostFilterSpec) HasConnectionStates() bool`

HasConnectionStates returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


