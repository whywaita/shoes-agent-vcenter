# VcenterClusterFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clusters** | Pointer to **[]string** | Identifiers of clusters that can match the filter. If unset or empty, clusters with any identifier match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**Names** | Pointer to **[]string** | Names that clusters must have to match the filter (see Cluster.Info.name). If unset or empty, clusters with any name match the filter. | [optional] 
**Folders** | Pointer to **[]string** | Folders that must contain the cluster for the cluster to match the filter. If unset or empty, clusters in any folder match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Folder. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Folder. | [optional] 
**Datacenters** | Pointer to **[]string** | Datacenters that must contain the cluster for the cluster to match the filter. If unset or empty, clusters in any datacenter match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: Datacenter. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: Datacenter. | [optional] 

## Methods

### NewVcenterClusterFilterSpec

`func NewVcenterClusterFilterSpec() *VcenterClusterFilterSpec`

NewVcenterClusterFilterSpec instantiates a new VcenterClusterFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterClusterFilterSpecWithDefaults

`func NewVcenterClusterFilterSpecWithDefaults() *VcenterClusterFilterSpec`

NewVcenterClusterFilterSpecWithDefaults instantiates a new VcenterClusterFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusters

`func (o *VcenterClusterFilterSpec) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VcenterClusterFilterSpec) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VcenterClusterFilterSpec) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VcenterClusterFilterSpec) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetNames

`func (o *VcenterClusterFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterClusterFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterClusterFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterClusterFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetFolders

`func (o *VcenterClusterFilterSpec) GetFolders() []string`

GetFolders returns the Folders field if non-nil, zero value otherwise.

### GetFoldersOk

`func (o *VcenterClusterFilterSpec) GetFoldersOk() (*[]string, bool)`

GetFoldersOk returns a tuple with the Folders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolders

`func (o *VcenterClusterFilterSpec) SetFolders(v []string)`

SetFolders sets Folders field to given value.

### HasFolders

`func (o *VcenterClusterFilterSpec) HasFolders() bool`

HasFolders returns a boolean if a field has been set.

### GetDatacenters

`func (o *VcenterClusterFilterSpec) GetDatacenters() []string`

GetDatacenters returns the Datacenters field if non-nil, zero value otherwise.

### GetDatacentersOk

`func (o *VcenterClusterFilterSpec) GetDatacentersOk() (*[]string, bool)`

GetDatacentersOk returns a tuple with the Datacenters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatacenters

`func (o *VcenterClusterFilterSpec) SetDatacenters(v []string)`

SetDatacenters sets Datacenters field to given value.

### HasDatacenters

`func (o *VcenterClusterFilterSpec) HasDatacenters() bool`

HasDatacenters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


