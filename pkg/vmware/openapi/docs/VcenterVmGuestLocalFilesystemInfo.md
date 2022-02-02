# VcenterVmGuestLocalFilesystemInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capacity** | **int64** | Total capacity of the file system, in bytes. | 
**FreeSpace** | **int64** | Free space on the file system, in bytes. | 
**Filesystem** | Pointer to **string** | Filesystem type, if known. For example, ext3 or NTFS. set if VMware Tools reports a value. | [optional] 
**Mappings** | Pointer to [**[]VcenterVmGuestLocalFilesystemVirtualDiskMapping**](VcenterVmGuestLocalFilesystemVirtualDiskMapping.md) | VirtualDisks backing the guest partition, if known. This field is optional because it was added in a newer version than its parent node. | [optional] 

## Methods

### NewVcenterVmGuestLocalFilesystemInfo

`func NewVcenterVmGuestLocalFilesystemInfo(capacity int64, freeSpace int64, ) *VcenterVmGuestLocalFilesystemInfo`

NewVcenterVmGuestLocalFilesystemInfo instantiates a new VcenterVmGuestLocalFilesystemInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestLocalFilesystemInfoWithDefaults

`func NewVcenterVmGuestLocalFilesystemInfoWithDefaults() *VcenterVmGuestLocalFilesystemInfo`

NewVcenterVmGuestLocalFilesystemInfoWithDefaults instantiates a new VcenterVmGuestLocalFilesystemInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapacity

`func (o *VcenterVmGuestLocalFilesystemInfo) GetCapacity() int64`

GetCapacity returns the Capacity field if non-nil, zero value otherwise.

### GetCapacityOk

`func (o *VcenterVmGuestLocalFilesystemInfo) GetCapacityOk() (*int64, bool)`

GetCapacityOk returns a tuple with the Capacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapacity

`func (o *VcenterVmGuestLocalFilesystemInfo) SetCapacity(v int64)`

SetCapacity sets Capacity field to given value.


### GetFreeSpace

`func (o *VcenterVmGuestLocalFilesystemInfo) GetFreeSpace() int64`

GetFreeSpace returns the FreeSpace field if non-nil, zero value otherwise.

### GetFreeSpaceOk

`func (o *VcenterVmGuestLocalFilesystemInfo) GetFreeSpaceOk() (*int64, bool)`

GetFreeSpaceOk returns a tuple with the FreeSpace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeSpace

`func (o *VcenterVmGuestLocalFilesystemInfo) SetFreeSpace(v int64)`

SetFreeSpace sets FreeSpace field to given value.


### GetFilesystem

`func (o *VcenterVmGuestLocalFilesystemInfo) GetFilesystem() string`

GetFilesystem returns the Filesystem field if non-nil, zero value otherwise.

### GetFilesystemOk

`func (o *VcenterVmGuestLocalFilesystemInfo) GetFilesystemOk() (*string, bool)`

GetFilesystemOk returns a tuple with the Filesystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilesystem

`func (o *VcenterVmGuestLocalFilesystemInfo) SetFilesystem(v string)`

SetFilesystem sets Filesystem field to given value.

### HasFilesystem

`func (o *VcenterVmGuestLocalFilesystemInfo) HasFilesystem() bool`

HasFilesystem returns a boolean if a field has been set.

### GetMappings

`func (o *VcenterVmGuestLocalFilesystemInfo) GetMappings() []VcenterVmGuestLocalFilesystemVirtualDiskMapping`

GetMappings returns the Mappings field if non-nil, zero value otherwise.

### GetMappingsOk

`func (o *VcenterVmGuestLocalFilesystemInfo) GetMappingsOk() (*[]VcenterVmGuestLocalFilesystemVirtualDiskMapping, bool)`

GetMappingsOk returns a tuple with the Mappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMappings

`func (o *VcenterVmGuestLocalFilesystemInfo) SetMappings(v []VcenterVmGuestLocalFilesystemVirtualDiskMapping)`

SetMappings sets Mappings field to given value.

### HasMappings

`func (o *VcenterVmGuestLocalFilesystemInfo) HasMappings() bool`

HasMappings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


