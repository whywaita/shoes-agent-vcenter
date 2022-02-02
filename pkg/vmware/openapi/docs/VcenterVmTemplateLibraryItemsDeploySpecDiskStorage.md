# VcenterVmTemplateLibraryItemsDeploySpecDiskStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Identifier for the datastore associated the deployed virtual machine&#39;s disk. | [optional] 
**StoragePolicy** | Pointer to [**VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy**](VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorage

`func NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorage() *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage`

NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorage instantiates a new VcenterVmTemplateLibraryItemsDeploySpecDiskStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorageWithDefaults

`func NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorageWithDefaults() *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage`

NewVcenterVmTemplateLibraryItemsDeploySpecDiskStorageWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDeploySpecDiskStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetStoragePolicy() VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) GetStoragePolicyOk() (*VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) SetStoragePolicy(v VcenterVmTemplateLibraryItemsDeploySpecDiskStoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDeploySpecDiskStorage) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


