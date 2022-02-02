# VcenterVMClonePlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | Pointer to **string** | Virtual machine folder into which the cloned virtual machine should be placed. If field is unset, the system will use the virtual machine folder of the source virtual machine. If this results in a conflict due to other placement parameters, the virtual machine clone operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | [optional] 
**ResourcePool** | Pointer to **string** | Resource pool into which the cloned virtual machine should be placed. If field is unset, the system will use the resource pool of the source virtual machine. If this results in a conflict due to other placement parameters, the virtual machine clone operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool. | [optional] 
**Host** | Pointer to **string** | Host onto which the cloned virtual machine should be placed.   If VM.ClonePlacementSpec.host and VM.ClonePlacementSpec.resource-pool are both specified, VM.ClonePlacementSpec.resource-pool must belong to VM.ClonePlacementSpec.host.    If VM.ClonePlacementSpec.host and VM.ClonePlacementSpec.cluster are both specified, VM.ClonePlacementSpec.host must be a member of VM.ClonePlacementSpec.cluster.  If this field is unset, if VM.ClonePlacementSpec.resource-pool is unset, the cloned virtual machine will use the host of the source virtual machine. if VM.ClonePlacementSpec.resource-pool is set, and the target is a standalone host, the host is used. if VM.ClonePlacementSpec.resource-pool is set, and the target is a DRS cluster, a host will be picked by DRS. if VM.ClonePlacementSpec.resource-pool is set, and the target is a cluster without DRS, InvalidArgument will be thrown. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | [optional] 
**Cluster** | Pointer to **string** | Cluster into which the cloned virtual machine should be placed.   If VM.ClonePlacementSpec.cluster and VM.ClonePlacementSpec.resource-pool are both specified, VM.ClonePlacementSpec.resource-pool must belong to VM.ClonePlacementSpec.cluster.    If VM.ClonePlacementSpec.cluster and VM.ClonePlacementSpec.host are both specified, VM.ClonePlacementSpec.host must be a member of VM.ClonePlacementSpec.cluster.  If VM.ClonePlacementSpec.resource-pool or VM.ClonePlacementSpec.host is specified, it is recommended that this field be unset. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | [optional] 
**Datastore** | Pointer to **string** | Datastore on which the cloned virtual machine&#39;s configuration state should be stored. This datastore will also be used for any virtual disks that are created as part of the virtual machine clone operation unless individually overridden. If field is unset, the system will use the datastore of the source virtual machine. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | [optional] 

## Methods

### NewVcenterVMClonePlacementSpec

`func NewVcenterVMClonePlacementSpec() *VcenterVMClonePlacementSpec`

NewVcenterVMClonePlacementSpec instantiates a new VcenterVMClonePlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMClonePlacementSpecWithDefaults

`func NewVcenterVMClonePlacementSpecWithDefaults() *VcenterVMClonePlacementSpec`

NewVcenterVMClonePlacementSpecWithDefaults instantiates a new VcenterVMClonePlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *VcenterVMClonePlacementSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterVMClonePlacementSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterVMClonePlacementSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterVMClonePlacementSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetResourcePool

`func (o *VcenterVMClonePlacementSpec) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterVMClonePlacementSpec) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterVMClonePlacementSpec) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VcenterVMClonePlacementSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHost

`func (o *VcenterVMClonePlacementSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterVMClonePlacementSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterVMClonePlacementSpec) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterVMClonePlacementSpec) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterVMClonePlacementSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterVMClonePlacementSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterVMClonePlacementSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterVMClonePlacementSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatastore

`func (o *VcenterVMClonePlacementSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVMClonePlacementSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVMClonePlacementSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVMClonePlacementSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


