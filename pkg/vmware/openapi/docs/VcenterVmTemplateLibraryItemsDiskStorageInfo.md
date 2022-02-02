# VcenterVmTemplateLibraryItemsDiskStorageInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | **string** | Identifier of the datastore where the disk is stored. | 
**StoragePolicy** | Pointer to **string** | Identifier of the storage policy associated with the virtual disk. | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsDiskStorageInfo

`func NewVcenterVmTemplateLibraryItemsDiskStorageInfo(datastore string, ) *VcenterVmTemplateLibraryItemsDiskStorageInfo`

NewVcenterVmTemplateLibraryItemsDiskStorageInfo instantiates a new VcenterVmTemplateLibraryItemsDiskStorageInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDiskStorageInfoWithDefaults

`func NewVcenterVmTemplateLibraryItemsDiskStorageInfoWithDefaults() *VcenterVmTemplateLibraryItemsDiskStorageInfo`

NewVcenterVmTemplateLibraryItemsDiskStorageInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDiskStorageInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.


### GetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) GetStoragePolicy() string`

GetStoragePolicy returns the StoragePolicy field if non-nil, zero value otherwise.

### GetStoragePolicyOk

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) GetStoragePolicyOk() (*string, bool)`

GetStoragePolicyOk returns a tuple with the StoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) SetStoragePolicy(v string)`

SetStoragePolicy sets StoragePolicy field to given value.

### HasStoragePolicy

`func (o *VcenterVmTemplateLibraryItemsDiskStorageInfo) HasStoragePolicy() bool`

HasStoragePolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


