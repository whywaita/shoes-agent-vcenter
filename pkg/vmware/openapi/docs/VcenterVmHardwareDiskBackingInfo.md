# VcenterVmHardwareDiskBackingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareDiskBackingType**](VcenterVmHardwareDiskBackingType.md) |  | 
**VmdkFile** | Pointer to **string** | Path of the VMDK file backing the virtual disk. This field is optional and it is only relevant when the value of Disk.BackingInfo.type is VMDK_FILE. | [optional] 

## Methods

### NewVcenterVmHardwareDiskBackingInfo

`func NewVcenterVmHardwareDiskBackingInfo(type_ VcenterVmHardwareDiskBackingType, ) *VcenterVmHardwareDiskBackingInfo`

NewVcenterVmHardwareDiskBackingInfo instantiates a new VcenterVmHardwareDiskBackingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareDiskBackingInfoWithDefaults

`func NewVcenterVmHardwareDiskBackingInfoWithDefaults() *VcenterVmHardwareDiskBackingInfo`

NewVcenterVmHardwareDiskBackingInfoWithDefaults instantiates a new VcenterVmHardwareDiskBackingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareDiskBackingInfo) GetType() VcenterVmHardwareDiskBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareDiskBackingInfo) GetTypeOk() (*VcenterVmHardwareDiskBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareDiskBackingInfo) SetType(v VcenterVmHardwareDiskBackingType)`

SetType sets Type field to given value.


### GetVmdkFile

`func (o *VcenterVmHardwareDiskBackingInfo) GetVmdkFile() string`

GetVmdkFile returns the VmdkFile field if non-nil, zero value otherwise.

### GetVmdkFileOk

`func (o *VcenterVmHardwareDiskBackingInfo) GetVmdkFileOk() (*string, bool)`

GetVmdkFileOk returns a tuple with the VmdkFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmdkFile

`func (o *VcenterVmHardwareDiskBackingInfo) SetVmdkFile(v string)`

SetVmdkFile sets VmdkFile field to given value.

### HasVmdkFile

`func (o *VcenterVmHardwareDiskBackingInfo) HasVmdkFile() bool`

HasVmdkFile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


