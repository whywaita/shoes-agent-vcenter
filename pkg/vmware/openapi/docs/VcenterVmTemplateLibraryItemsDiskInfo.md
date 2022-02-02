# VcenterVmTemplateLibraryItemsDiskInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | Pointer to **int64** | Capacity of the virtual disk in bytes. | [optional] 
**DiskStorage** | [**VcenterVmTemplateLibraryItemsDiskStorageInfo**](VcenterVmTemplateLibraryItemsDiskStorageInfo.md) |  | 

## Methods

### NewVcenterVmTemplateLibraryItemsDiskInfo

`func NewVcenterVmTemplateLibraryItemsDiskInfo(diskStorage VcenterVmTemplateLibraryItemsDiskStorageInfo, ) *VcenterVmTemplateLibraryItemsDiskInfo`

NewVcenterVmTemplateLibraryItemsDiskInfo instantiates a new VcenterVmTemplateLibraryItemsDiskInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDiskInfoWithDefaults

`func NewVcenterVmTemplateLibraryItemsDiskInfoWithDefaults() *VcenterVmTemplateLibraryItemsDiskInfo`

NewVcenterVmTemplateLibraryItemsDiskInfoWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDiskInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.

### HasCapacity

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) HasCapacity() bool`

HasCapacity returns a boolean if a field has been set.

### GetDiskStorage

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) GetDiskStorage() VcenterVmTemplateLibraryItemsDiskStorageInfo`

GetDiskStorage returns the DiskStorage field if non-nil, zero value otherwise.

### GetDiskStorageOk

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) GetDiskStorageOk() (*VcenterVmTemplateLibraryItemsDiskStorageInfo, bool)`

GetDiskStorageOk returns a tuple with the DiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorage

`func (o *VcenterVmTemplateLibraryItemsDiskInfo) SetDiskStorage(v VcenterVmTemplateLibraryItemsDiskStorageInfo)`

SetDiskStorage sets DiskStorage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


