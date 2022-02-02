# VcenterVmHardwareParallelBackingInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterVmHardwareParallelBackingType**](VcenterVmHardwareParallelBackingType.md) |  | 
**File** | Pointer to **string** | Path of the file backing the virtual parallel port. This field is optional and it is only relevant when the value of Parallel.BackingInfo.type is FILE. | [optional] 
**HostDevice** | Pointer to **string** | Name of the device backing the virtual parallel port.    This field will be unset if Parallel.BackingInfo.auto-detect is true and the virtual parallel port is not connected or no suitable device is available on the host. | [optional] 
**AutoDetect** | Pointer to **bool** | Flag indicating whether the virtual parallel port is configured to automatically detect a suitable host device. This field is optional and it is only relevant when the value of Parallel.BackingInfo.type is HOST_DEVICE. | [optional] 

## Methods

### NewVcenterVmHardwareParallelBackingInfo

`func NewVcenterVmHardwareParallelBackingInfo(type_ VcenterVmHardwareParallelBackingType, ) *VcenterVmHardwareParallelBackingInfo`

NewVcenterVmHardwareParallelBackingInfo instantiates a new VcenterVmHardwareParallelBackingInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareParallelBackingInfoWithDefaults

`func NewVcenterVmHardwareParallelBackingInfoWithDefaults() *VcenterVmHardwareParallelBackingInfo`

NewVcenterVmHardwareParallelBackingInfoWithDefaults instantiates a new VcenterVmHardwareParallelBackingInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterVmHardwareParallelBackingInfo) GetType() VcenterVmHardwareParallelBackingType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterVmHardwareParallelBackingInfo) GetTypeOk() (*VcenterVmHardwareParallelBackingType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterVmHardwareParallelBackingInfo) SetType(v VcenterVmHardwareParallelBackingType)`

SetType sets Type field to given value.


### GetFile

`func (o *VcenterVmHardwareParallelBackingInfo) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *VcenterVmHardwareParallelBackingInfo) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *VcenterVmHardwareParallelBackingInfo) SetFile(v string)`

SetFile sets File field to given value.

### HasFile

`func (o *VcenterVmHardwareParallelBackingInfo) HasFile() bool`

HasFile returns a boolean if a field has been set.

### GetHostDevice

`func (o *VcenterVmHardwareParallelBackingInfo) GetHostDevice() string`

GetHostDevice returns the HostDevice field if non-nil, zero value otherwise.

### GetHostDeviceOk

`func (o *VcenterVmHardwareParallelBackingInfo) GetHostDeviceOk() (*string, bool)`

GetHostDeviceOk returns a tuple with the HostDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostDevice

`func (o *VcenterVmHardwareParallelBackingInfo) SetHostDevice(v string)`

SetHostDevice sets HostDevice field to given value.

### HasHostDevice

`func (o *VcenterVmHardwareParallelBackingInfo) HasHostDevice() bool`

HasHostDevice returns a boolean if a field has been set.

### GetAutoDetect

`func (o *VcenterVmHardwareParallelBackingInfo) GetAutoDetect() bool`

GetAutoDetect returns the AutoDetect field if non-nil, zero value otherwise.

### GetAutoDetectOk

`func (o *VcenterVmHardwareParallelBackingInfo) GetAutoDetectOk() (*bool, bool)`

GetAutoDetectOk returns a tuple with the AutoDetect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoDetect

`func (o *VcenterVmHardwareParallelBackingInfo) SetAutoDetect(v bool)`

SetAutoDetect sets AutoDetect field to given value.

### HasAutoDetect

`func (o *VcenterVmHardwareParallelBackingInfo) HasAutoDetect() bool`

HasAutoDetect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


