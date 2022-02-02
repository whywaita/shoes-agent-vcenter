# VcenterVmTemplateLibraryItemsVmHomeStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | **string** | Identifier of the datastore where the virtual machine template&#39;s configuration and log files are stored. | 
**StoragePolicy** | Pointer to **string** | Identifier of the storage policy associated with the virtual machine template&#39;s configuration and log files. | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsVmHomeStorageInfo

`func NewVcenterVmTemplateLibraryItemsVmHomeStorageInfo(datastore string, ) *VcenterVmTemplateLibraryItemsVmHomeStorageInfo`

NewVcenterVmTemplateLibraryItemsVmHomeStorageInfo instantiates a new VcenterVmTemplateLibraryItemsVmHomeStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsVmHomeStorageInfoWithDefaults

`func NewVcenterVmTemplateLibraryItemsVmHomeStorageInfoWithDefaults() *VcenterVmTemplateLibraryItemsVmHomeStorageInfo`

NewVcenterVmTemplateLibraryItemsVmHomeStorageInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsVmHomeStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.


### GetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) GetStoragePolicy() string`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) GetStoragePolicyOk() (*string, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) SetStoragePolicy(v string)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsVmHomeStorageInfo) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


