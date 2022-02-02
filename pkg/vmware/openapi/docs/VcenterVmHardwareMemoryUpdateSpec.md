# VcenterVmHardwareMemoryUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeMiB** | Pointer to **int64** | New memory size in mebibytes.   The supported range of memory sizes is constrained by the configured guest operating system and virtual hardware version of the virtual machine.    If the virtual machine is running, this value may only be changed if Memory.Info.hot-add-enabled is true, and the new memory size must satisfy the constraints specified by Memory.Info.hot-add-increment-size-mib and Memory.Info.hot-add-limit-mib.  If unset, the value is unchanged. | [optional] 
**HotAddEnabled** | Pointer to **bool** | Flag indicating whether adding memory while the virtual machine is running should be enabled.   Some guest operating systems may consume more resources or perform less efficiently when they run on hardware that supports adding memory while the machine is running.    This field may only be modified if the virtual machine is not powered on.  If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareMemoryUpdateSpec

`func NewVcenterVmHardwareMemoryUpdateSpec() *VcenterVmHardwareMemoryUpdateSpec`

NewVcenterVmHardwareMemoryUpdateSpec instantiates a new VcenterVmHardwareMemoryUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareMemoryUpdateSpecWithDefaults

`func NewVcenterVmHardwareMemoryUpdateSpecWithDefaults() *VcenterVmHardwareMemoryUpdateSpec`

NewVcenterVmHardwareMemoryUpdateSpecWithDefaults instantiates a new VcenterVmHardwareMemoryUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeMiB

`func (o *VcenterVmHardwareMemoryUpdateSpec) GetSizeMiB() int64`

GetSizeMiB returns the SizeMiB field if non-nil, zero value otherwise.

### GetSizeMiBOk

`func (o *VcenterVmHardwareMemoryUpdateSpec) GetSizeMiBOk() (*int64, bool)`

GetSizeMiBOk returns a tuple with the SizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeMiB

`func (o *VcenterVmHardwareMemoryUpdateSpec) SetSizeMiB(v int64)`

SetSizeMiB sets SizeMiB field to given value.

### HasSizeMiB

`func (o *VcenterVmHardwareMemoryUpdateSpec) HasSizeMiB() bool`

HasSizeMiB returns a boolean if a field has been set.

### GetHotAddEnabled

`func (o *VcenterVmHardwareMemoryUpdateSpec) GetHotAddEnabled() bool`

GetHotAddEnabled returns the HotAddEnabled field if non-nil, zero value otherwise.

### GetHotAddEnabledOk

`func (o *VcenterVmHardwareMemoryUpdateSpec) GetHotAddEnabledOk() (*bool, bool)`

GetHotAddEnabledOk returns a tuple with the HotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotAddEnabled

`func (o *VcenterVmHardwareMemoryUpdateSpec) SetHotAddEnabled(v bool)`

SetHotAddEnabled sets HotAddEnabled field to given value.

### HasHotAddEnabled

`func (o *VcenterVmHardwareMemoryUpdateSpec) HasHotAddEnabled() bool`

HasHotAddEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


