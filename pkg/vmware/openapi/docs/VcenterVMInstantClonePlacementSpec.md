# VcenterVMInstantClonePlacementSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Folder** | Pointer to **string** | Virtual machine folder into which the InstantCloned virtual machine should be placed. If field is unset, the system will use the virtual machine folder of the source virtual machine. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Folder. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Folder. | [optional] 
**ResourcePool** | Pointer to **string** | Resource pool into which the InstantCloned virtual machine should be placed. If field is unset, the system will use the resource pool of the source virtual machine. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool. | [optional] 
**Datastore** | Pointer to **string** | Datastore on which the InstantCloned virtual machine&#39;s configuration state should be stored. This datastore will also be used for any virtual disks that are created as part of the virtual machine InstantClone operation. If field is unset, the system will use the datastore of the source virtual machine. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Datastore. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Datastore. | [optional] 

## Methods

### NewVcenterVMInstantClonePlacementSpec

`func NewVcenterVMInstantClonePlacementSpec() *VcenterVMInstantClonePlacementSpec`

NewVcenterVMInstantClonePlacementSpec instantiates a new VcenterVMInstantClonePlacementSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMInstantClonePlacementSpecWithDefaults

`func NewVcenterVMInstantClonePlacementSpecWithDefaults() *VcenterVMInstantClonePlacementSpec`

NewVcenterVMInstantClonePlacementSpecWithDefaults instantiates a new VcenterVMInstantClonePlacementSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFolder

`func (o *VcenterVMInstantClonePlacementSpec) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VcenterVMInstantClonePlacementSpec) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VcenterVMInstantClonePlacementSpec) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VcenterVMInstantClonePlacementSpec) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetResourcePool

`func (o *VcenterVMInstantClonePlacementSpec) GetResourcePool() string`

GetResourcePool returns the ResourcePool field if non-nil, zero value otherwise.

### GetResourcePoolOk

`func (o *VcenterVMInstantClonePlacementSpec) GetResourcePoolOk() (*string, bool)`

GetResourcePoolOk returns a tuple with the ResourcePool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePool

`func (o *VcenterVMInstantClonePlacementSpec) SetResourcePool(v string)`

SetResourcePool sets ResourcePool field to given value.

### HasResourcePool

`func (o *VcenterVMInstantClonePlacementSpec) HasResourcePool() bool`

HasResourcePool returns a boolean if a field has been set.

### GetDatastore

`func (o *VcenterVMInstantClonePlacementSpec) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVMInstantClonePlacementSpec) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVMInstantClonePlacementSpec) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVMInstantClonePlacementSpec) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


