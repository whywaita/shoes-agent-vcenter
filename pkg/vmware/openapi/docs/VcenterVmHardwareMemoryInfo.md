# VcenterVmHardwareMemoryInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeMiB** | **int64** | Memory size in mebibytes. | 
**HotAddEnabled** | **bool** | Flag indicating whether adding memory while the virtual machine is running is enabled.   Some guest operating systems may consume more resources or perform less efficiently when they run on hardware that supports adding memory while the machine is running.  | 
**HotAddIncrementSizeMiB** | Pointer to **int64** | The granularity, in mebibytes, at which memory can be added to a running virtual machine.   When adding memory to a running virtual machine, the amount of memory added must be at least Memory.Info.hot-add-increment-size-mib and the total memory size of the virtual machine must be a multiple of {@link&gt;hotAddIncrementSize}.  Only set when Memory.Info.hot-add-enabled is true and the virtual machine is running. | [optional] 
**HotAddLimitMiB** | Pointer to **int64** | The maximum amount of memory, in mebibytes, that can be added to a running virtual machine. Only set when Memory.Info.hot-add-enabled is true and the virtual machine is running. | [optional] 

## Methods

### NewVcenterVmHardwareMemoryInfo

`func NewVcenterVmHardwareMemoryInfo(sizeMiB int64, hotAddEnabled bool, ) *VcenterVmHardwareMemoryInfo`

NewVcenterVmHardwareMemoryInfo instantiates a new VcenterVmHardwareMemoryInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareMemoryInfoWithDefaults

`func NewVcenterVmHardwareMemoryInfoWithDefaults() *VcenterVmHardwareMemoryInfo`

NewVcenterVmHardwareMemoryInfoWithDefaults instantiates a new VcenterVmHardwareMemoryInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeMiB

`func (o *VcenterVmHardwareMemoryInfo) GetSizeMiB() int64`

GetSizeMiB returns the SizeMiB field if non-nil, zero value otherwise.

### GetSizeMiBOk

`func (o *VcenterVmHardwareMemoryInfo) GetSizeMiBOk() (*int64, bool)`

GetSizeMiBOk returns a tuple with the SizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMiB

`func (o *VcenterVmHardwareMemoryInfo) SetSizeMiB(v int64)`

SetSizeMiB sets SizeMiB field to given value.


### GetHotAddEnabled

`func (o *VcenterVmHardwareMemoryInfo) GetHotAddEnabled() bool`

GetHotAddEnabled returns the HotAddEnabled field if non-nil, zero value otherwise.

### GetHotAddEnabledOk

`func (o *VcenterVmHardwareMemoryInfo) GetHotAddEnabledOk() (*bool, bool)`

GetHotAddEnabledOk returns a tuple with the HotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotAddEnabled

`func (o *VcenterVmHardwareMemoryInfo) SetHotAddEnabled(v bool)`

SetHotAddEnabled sets HotAddEnabled field to given value.


### GetHotAddIncrementSizeMiB

`func (o *VcenterVmHardwareMemoryInfo) GetHotAddIncrementSizeMiB() int64`

GetHotAddIncrementSizeMiB returns the HotAddIncrementSizeMiB field if non-nil, zero value otherwise.

### GetHotAddIncrementSizeMiBOk

`func (o *VcenterVmHardwareMemoryInfo) GetHotAddIncrementSizeMiBOk() (*int64, bool)`

GetHotAddIncrementSizeMiBOk returns a tuple with the HotAddIncrementSizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotAddIncrementSizeMiB

`func (o *VcenterVmHardwareMemoryInfo) SetHotAddIncrementSizeMiB(v int64)`

SetHotAddIncrementSizeMiB sets HotAddIncrementSizeMiB field to given value.

### HasHotAddIncrementSizeMiB

`func (o *VcenterVmHardwareMemoryInfo) HasHotAddIncrementSizeMiB() bool`

HasHotAddIncrementSizeMiB returns a boolean if a field has been set.

### GetHotAddLimitMiB

`func (o *VcenterVmHardwareMemoryInfo) GetHotAddLimitMiB() int64`

GetHotAddLimitMiB returns the HotAddLimitMiB field if non-nil, zero value otherwise.

### GetHotAddLimitMiBOk

`func (o *VcenterVmHardwareMemoryInfo) GetHotAddLimitMiBOk() (*int64, bool)`

GetHotAddLimitMiBOk returns a tuple with the HotAddLimitMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotAddLimitMiB

`func (o *VcenterVmHardwareMemoryInfo) SetHotAddLimitMiB(v int64)`

SetHotAddLimitMiB sets HotAddLimitMiB field to given value.

### HasHotAddLimitMiB

`func (o *VcenterVmHardwareMemoryInfo) HasHotAddLimitMiB() bool`

HasHotAddLimitMiB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


