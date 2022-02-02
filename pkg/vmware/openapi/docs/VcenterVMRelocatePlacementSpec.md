# VcenterVMRelocatePlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | Pointer to **string** | Virtual machine folder into which the virtual machine should be placed. If this field is unset, the virtual machine will stay in the current folder. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | [optional] 
**ResourcePool** | Pointer to **string** | Resource pool into which the virtual machine should be placed. If this field is unset, the virtual machine will stay in the current resource pool. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool. | [optional] 
**Host** | Pointer to **string** | Host onto which the virtual machine should be placed.   If VM.RelocatePlacementSpec.host and VM.RelocatePlacementSpec.resource-pool are both specified, VM.RelocatePlacementSpec.resource-pool must belong to VM.RelocatePlacementSpec.host.    If VM.RelocatePlacementSpec.host and VM.RelocatePlacementSpec.cluster are both specified, VM.RelocatePlacementSpec.host must be a member of VM.RelocatePlacementSpec.cluster.  If this field is unset, if VM.RelocatePlacementSpec.resource-pool is unset, the virtual machine will remain on the current host. if VM.RelocatePlacementSpec.resource-pool is set, and the target is a standalone host, the host is used. if VM.RelocatePlacementSpec.resource-pool is set, and the target is a DRS cluster, a host will be picked by DRS. if VM.RelocatePlacementSpec.resource-pool is set, and the target is a cluster without DRS, InvalidArgument will be thrown. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | [optional] 
**Cluster** | Pointer to **string** | Cluster into which the virtual machine should be placed.   If VM.RelocatePlacementSpec.cluster and VM.RelocatePlacementSpec.resource-pool are both specified, VM.RelocatePlacementSpec.resource-pool must belong to VM.RelocatePlacementSpec.cluster.    If VM.RelocatePlacementSpec.cluster and VM.RelocatePlacementSpec.host are both specified, VM.RelocatePlacementSpec.host must be a member of VM.RelocatePlacementSpec.cluster.  If VM.RelocatePlacementSpec.resource-pool or VM.RelocatePlacementSpec.host is specified, it is recommended that this field be unset. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | [optional] 
**Datastore** | Pointer to **string** | Datastore on which the virtual machine&#39;s configuration state should be stored. This datastore will also be used for any virtual disks that are associated with the virtual machine, unless individually overridden. If this field is unset, the virtual machine will remain on the current datastore. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | [optional] 

## Methods

### NewVcenterVMRelocatePlacementSpec

`func NewVcenterVMRelocatePlacementSpec() *VcenterVMRelocatePlacementSpec`

NewVcenterVMRelocatePlacementSpec instantiates a new VcenterVMRelocatePlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMRelocatePlacementSpecWithDefaults

`func NewVcenterVMRelocatePlacementSpecWithDefaults() *VcenterVMRelocatePlacementSpec`

NewVcenterVMRelocatePlacementSpecWithDefaults instantiates a new VcenterVMRelocatePlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *VcenterVMRelocatePlacementSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterVMRelocatePlacementSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterVMRelocatePlacementSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterVMRelocatePlacementSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetResourcePool

`func (o *VcenterVMRelocatePlacementSpec) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterVMRelocatePlacementSpec) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterVMRelocatePlacementSpec) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VcenterVMRelocatePlacementSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHost

`func (o *VcenterVMRelocatePlacementSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterVMRelocatePlacementSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterVMRelocatePlacementSpec) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterVMRelocatePlacementSpec) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterVMRelocatePlacementSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterVMRelocatePlacementSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterVMRelocatePlacementSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterVMRelocatePlacementSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetDatastore

`func (o *VcenterVMRelocatePlacementSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVMRelocatePlacementSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVMRelocatePlacementSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVMRelocatePlacementSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


