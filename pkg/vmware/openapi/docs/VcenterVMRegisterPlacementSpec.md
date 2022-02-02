# VcenterVMRegisterPlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | Pointer to **string** | Virtual machine folder into which the virtual machine should be placed. This field is currently required. In the future, if this field is unset, the system will attempt to choose a suitable folder for the virtual machine; if a folder cannot be chosen, the virtual machine creation operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | [optional] 
**ResourcePool** | Pointer to **string** | Resource pool into which the virtual machine should be placed. This field is currently required if both VM.ComputePlacementSpec.host and VM.ComputePlacementSpec.cluster are unset. In the future, if this field is unset, the system will attempt to choose a suitable resource pool for the virtual machine; if a resource pool cannot be chosen, the virtual machine creation operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool. | [optional] 
**Host** | Pointer to **string** | Host onto which the virtual machine should be placed.   If VM.ComputePlacementSpec.host and VM.ComputePlacementSpec.resource-pool are both specified, VM.ComputePlacementSpec.resource-pool must belong to VM.ComputePlacementSpec.host.    If VM.ComputePlacementSpec.host and VM.ComputePlacementSpec.cluster are both specified, VM.ComputePlacementSpec.host must be a member of VM.ComputePlacementSpec.cluster.  This field may be unset if VM.ComputePlacementSpec.resource-pool or VM.ComputePlacementSpec.cluster is specified. If unset, the system will attempt to choose a suitable host for the virtual machine; if a host cannot be chosen, the virtual machine creation operation will fail. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | [optional] 
**Cluster** | Pointer to **string** | Cluster into which the virtual machine should be placed.   If VM.ComputePlacementSpec.cluster and VM.ComputePlacementSpec.resource-pool are both specified, VM.ComputePlacementSpec.resource-pool must belong to VM.ComputePlacementSpec.cluster.    If VM.ComputePlacementSpec.cluster and VM.ComputePlacementSpec.host are both specified, VM.ComputePlacementSpec.host must be a member of VM.ComputePlacementSpec.cluster.  If VM.ComputePlacementSpec.resource-pool or VM.ComputePlacementSpec.host is specified, it is recommended that this field be unset. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | [optional] 

## Methods

### NewVcenterVMRegisterPlacementSpec

`func NewVcenterVMRegisterPlacementSpec() *VcenterVMRegisterPlacementSpec`

NewVcenterVMRegisterPlacementSpec instantiates a new VcenterVMRegisterPlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMRegisterPlacementSpecWithDefaults

`func NewVcenterVMRegisterPlacementSpecWithDefaults() *VcenterVMRegisterPlacementSpec`

NewVcenterVMRegisterPlacementSpecWithDefaults instantiates a new VcenterVMRegisterPlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *VcenterVMRegisterPlacementSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterVMRegisterPlacementSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterVMRegisterPlacementSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterVMRegisterPlacementSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetResourcePool

`func (o *VcenterVMRegisterPlacementSpec) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterVMRegisterPlacementSpec) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterVMRegisterPlacementSpec) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VcenterVMRegisterPlacementSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHost

`func (o *VcenterVMRegisterPlacementSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterVMRegisterPlacementSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterVMRegisterPlacementSpec) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterVMRegisterPlacementSpec) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterVMRegisterPlacementSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterVMRegisterPlacementSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterVMRegisterPlacementSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterVMRegisterPlacementSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


