# VcenterVmHardwareCpuInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int64** | Number of CPU cores. | 
**CoresPerSocket** | **int64** | Number of CPU cores per socket. | 
**HotAddEnabled** | **bool** | Flag indicating whether adding CPUs while the virtual machine is running is enabled. | 
**HotRemoveEnabled** | **bool** | Flag indicating whether removing CPUs while the virtual machine is running is enabled. | 

## Methods

### NewVcenterVmHardwareCpuInfo

`func NewVcenterVmHardwareCpuInfo(count int64, coresPerSocket int64, hotAddEnabled bool, hotRemoveEnabled bool, ) *VcenterVmHardwareCpuInfo`

NewVcenterVmHardwareCpuInfo instantiates a new VcenterVmHardwareCpuInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCpuInfoWithDefaults

`func NewVcenterVmHardwareCpuInfoWithDefaults() *VcenterVmHardwareCpuInfo`

NewVcenterVmHardwareCpuInfoWithDefaults instantiates a new VcenterVmHardwareCpuInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *VcenterVmHardwareCpuInfo) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *VcenterVmHardwareCpuInfo) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *VcenterVmHardwareCpuInfo) SetCount(v int64)`

SetCount sets Count field to given value.


### GetCoresPerSocket

`func (o *VcenterVmHardwareCpuInfo) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *VcenterVmHardwareCpuInfo) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *VcenterVmHardwareCpuInfo) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.


### GetHotAddEnabled

`func (o *VcenterVmHardwareCpuInfo) GetHotAddEnabled() bool`

GetHotAddEnabled returns the HotAddEnabled field if non-nil, zero value otherwise.

### GetHotAddEnabledOk

`func (o *VcenterVmHardwareCpuInfo) GetHotAddEnabledOk() (*bool, bool)`

GetHotAddEnabledOk returns a tuple with the HotAddEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotAddEnabled

`func (o *VcenterVmHardwareCpuInfo) SetHotAddEnabled(v bool)`

SetHotAddEnabled sets HotAddEnabled field to given value.


### GetHotRemoveEnabled

`func (o *VcenterVmHardwareCpuInfo) GetHotRemoveEnabled() bool`

GetHotRemoveEnabled returns the HotRemoveEnabled field if non-nil, zero value otherwise.

### GetHotRemoveEnabledOk

`func (o *VcenterVmHardwareCpuInfo) GetHotRemoveEnabledOk() (*bool, bool)`

GetHotRemoveEnabledOk returns a tuple with the HotRemoveEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHotRemoveEnabled

`func (o *VcenterVmHardwareCpuInfo) SetHotRemoveEnabled(v bool)`

SetHotRemoveEnabled sets HotRemoveEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


