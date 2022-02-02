# VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Identifier of the datastore for the virtual machine template&#39;s configuration and log files. | [optional] 
**StoragePolicy** | Pointer to [**VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy**](VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage

`func NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage() *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage`

NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage instantiates a new VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStorageWithDefaults

`func NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStorageWithDefaults() *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage`

NewVcenterVmTemplateLibraryItemsCreateSpecVmHomeStorageWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) GetStoragePolicy() VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) GetStoragePolicyOk() (*VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) SetStoragePolicy(v VcenterVmTemplateLibraryItemsCreateSpecVmHomeStoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


