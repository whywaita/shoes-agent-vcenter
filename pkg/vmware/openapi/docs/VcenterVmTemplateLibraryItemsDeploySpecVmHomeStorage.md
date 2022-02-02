# VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to **string** | Identifier of the datastore for the deployed virtual machine&#39;s configuration and log files. | [optional] 
**StoragePolicy** | Pointer to [**VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicy**](VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicy.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage

`func NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage() *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage`

NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage instantiates a new VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStorageWithDefaults

`func NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStorageWithDefaults() *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage`

NewVcenterVmTemplateLibraryItemsDeploySpecVmHomeStorageWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) GetStoragePolicy() VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicy`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) GetStoragePolicyOk() (*VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicy, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) SetStoragePolicy(v VcenterVmTemplateLibraryItemsDeploySpecVmHomeStoragePolicy)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


