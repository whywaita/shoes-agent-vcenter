# VcenterVmGuestLocalFilesystemVirtualDiskMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | **string** | The virtual disk. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Disk. | 

## Methods

### NewVcenterVmGuestLocalFilesystemVirtualDiskMapping

`func NewVcenterVmGuestLocalFilesystemVirtualDiskMapping(disk string, ) *VcenterVmGuestLocalFilesystemVirtualDiskMapping`

NewVcenterVmGuestLocalFilesystemVirtualDiskMapping instantiates a new VcenterVmGuestLocalFilesystemVirtualDiskMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmGuestLocalFilesystemVirtualDiskMappingWithDefaults

`func NewVcenterVmGuestLocalFilesystemVirtualDiskMappingWithDefaults() *VcenterVmGuestLocalFilesystemVirtualDiskMapping`

NewVcenterVmGuestLocalFilesystemVirtualDiskMappingWithDefaults instantiates a new VcenterVmGuestLocalFilesystemVirtualDiskMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *VcenterVmGuestLocalFilesystemVirtualDiskMapping) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VcenterVmGuestLocalFilesystemVirtualDiskMapping) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VcenterVmGuestLocalFilesystemVirtualDiskMapping) SetDisk(v string)`

SetDisk sets Disk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


