# VcenterVmTemplateLibraryItemsCreateSpecDiskStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Identifier for the datastore associated with a virtual machine template&#39;s disk. | [optional] 
**StoragePolicy** | Pointer to [**VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy**](VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorage

`func NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorage() *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage`

NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorage instantiates a new VcenterVmTemplateLibraryItemsCreateSpecDiskStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageWithDefaults

`func NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageWithDefaults() *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage`

NewVcenterVmTemplateLibraryItemsCreateSpecDiskStorageWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreateSpecDiskStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) GetStoragePolicy() VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) GetStoragePolicyOk() (*VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) SetStoragePolicy(v VcenterVmTemplateLibraryItemsCreateSpecDiskStoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsCreateSpecDiskStorage) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


