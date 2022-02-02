# VcenterVmTemplateLibraryItemsDeployPlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | Pointer to **string** | Virtual machine folder into which the deployed virtual machine should be placed. | [optional] 
**ResourcePool** | Pointer to **string** | Resource pool into which the deployed virtual machine should be placed. | [optional] 
**Host** | Pointer to **string** | Host onto which the virtual machine should be placed. If {@name #host} and {@name #resourcePool} are both specified, {@name #resourcePool} must belong to {@name #host}. If {@name #host} and {@name #cluster} are both specified, {@name #host} must be a member of {@name #cluster}. | [optional] 
**Cluster** | Pointer to **string** | Cluster onto which the deployed virtual machine should be placed. If {@name #cluster} and {@name #resourcePool} are both specified, {@name #resourcePool} must belong to {@name #cluster}. If {@name #cluster} and {@name #host} are both specified, {@name #host} must be a member of {@name #cluster}. | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsDeployPlacementSpec

`func NewVcenterVmTemplateLibraryItemsDeployPlacementSpec() *VcenterVmTemplateLibraryItemsDeployPlacementSpec`

NewVcenterVmTemplateLibraryItemsDeployPlacementSpec instantiates a new VcenterVmTemplateLibraryItemsDeployPlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDeployPlacementSpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsDeployPlacementSpecWithDefaults() *VcenterVmTemplateLibraryItemsDeployPlacementSpec`

NewVcenterVmTemplateLibraryItemsDeployPlacementSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDeployPlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetResourcePool

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetHost

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterVmTemplateLibraryItemsDeployPlacementSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


