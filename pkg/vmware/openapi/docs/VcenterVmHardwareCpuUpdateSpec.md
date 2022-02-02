# VcenterVmHardwareCpuUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | Pointer to **int64** | New number of CPU cores. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket.   The supported range of CPU counts is constrained by the configured guest operating system and virtual hardware version of the virtual machine.    If the virtual machine is running, the number of CPU cores may only be increased if Cpu.Info.hot-add-enabled is true, and may only be decreased if Cpu.Info.hot-remove-enabled is true.  If unset, the value is unchanged. | [optional] 
**CoresPerSocket** | Pointer to **int64** | New number of CPU cores per socket. The number of CPU cores in the virtual machine must be a multiple of the number of cores per socket. If unset, the value is unchanged. | [optional] 
**HotAddEnabled** | Pointer to **bool** | Flag indicating whether adding CPUs while the virtual machine is running is enabled.   This field may only be modified if the virtual machine is powered off.  If unset, the value is unchanged. | [optional] 
**HotRemoveEnabled** | Pointer to **bool** | Flag indicating whether removing CPUs while the virtual machine is running is enabled.   This field may only be modified if the virtual machine is powered off.  If unset, the value is unchanged. | [optional] 

## Methods

### NewVcenterVmHardwareCpuUpdateSpec

`func NewVcenterVmHardwareCpuUpdateSpec() *VcenterVmHardwareCpuUpdateSpec`

NewVcenterVmHardwareCpuUpdateSpec instantiates a new VcenterVmHardwareCpuUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCpuUpdateSpecWithDefaults

`func NewVcenterVmHardwareCpuUpdateSpecWithDefaults() *VcenterVmHardwareCpuUpdateSpec`

NewVcenterVmHardwareCpuUpdateSpecWithDefaults instantiates a new VcenterVmHardwareCpuUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VcenterVmHardwareCpuUpdateSpec) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VcenterVmHardwareCpuUpdateSpec) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VcenterVmHardwareCpuUpdateSpec) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *VcenterVmHardwareCpuUpdateSpec) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *VcenterVmHardwareCpuUpdateSpec) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *VcenterVmHardwareCpuUpdateSpec) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *VcenterVmHardwareCpuUpdateSpec) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *VcenterVmHardwareCpuUpdateSpec) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetHotAddEnabled

`func (o *VcenterVmHardwareCpuUpdateSpec) GetHotAddEnabled() bool`

GetHotAddEnabled returns the HotAddEnabled field if non-nil, zero value otherwise.

### GetHotAddEnabledOk

`func (o *VcenterVmHardwareCpuUpdateSpec) GetHotAddEnabledOk() (*bool, bool)`

GetHotAddEnabledOk returns a tuple with the HotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotAddEnabled

`func (o *VcenterVmHardwareCpuUpdateSpec) SetHotAddEnabled(v bool)`

SetHotAddEnabled sets HotAddEnabled field to given value.

### HasHotAddEnabled

`func (o *VcenterVmHardwareCpuUpdateSpec) HasHotAddEnabled() bool`

HasHotAddEnabled returns a boolean if a field has been set.

### GetHotRemoveEnabled

`func (o *VcenterVmHardwareCpuUpdateSpec) GetHotRemoveEnabled() bool`

GetHotRemoveEnabled returns the HotRemoveEnabled field if non-nil, zero value otherwise.

### GetHotRemoveEnabledOk

`func (o *VcenterVmHardwareCpuUpdateSpec) GetHotRemoveEnabledOk() (*bool, bool)`

GetHotRemoveEnabledOk returns a tuple with the HotRemoveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotRemoveEnabled

`func (o *VcenterVmHardwareCpuUpdateSpec) SetHotRemoveEnabled(v bool)`

SetHotRemoveEnabled sets HotRemoveEnabled field to given value.

### HasHotRemoveEnabled

`func (o *VcenterVmHardwareCpuUpdateSpec) HasHotRemoveEnabled() bool`

HasHotRemoveEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


