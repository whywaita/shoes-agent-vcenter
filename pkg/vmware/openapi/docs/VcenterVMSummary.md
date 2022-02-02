# VcenterVMSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vm** | **string** | Identifier of the virtual machine. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine. | 
**Name** | **string** | Name of the Virtual machine. | 
**PowerState** | [**VcenterVmPowerState**](VcenterVmPowerState.md) |  | 
**CpuCount** | Pointer to **int64** | Number of CPU cores. This field will be unset if the virtual machine configuration is not available. For example, the configuration information would be unavailable if the server is unable to access the virtual machine files on disk, and is often also unavailable during the intial phases of virtual machine creation. | [optional] 
**MemorySizeMiB** | Pointer to **int64** | Memory size in mebibytes. This field will be unset if the virtual machine configuration is not available. For example, the configuration information would be unavailable if the server is unable to access the virtual machine files on disk, and is often also unavailable during the intial phases of virtual machine creation. | [optional] 

## Methods

### NewVcenterVMSummary

`func NewVcenterVMSummary(vm string, name string, powerState VcenterVmPowerState, ) *VcenterVMSummary`

NewVcenterVMSummary instantiates a new VcenterVMSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMSummaryWithDefaults

`func NewVcenterVMSummaryWithDefaults() *VcenterVMSummary`

NewVcenterVMSummaryWithDefaults instantiates a new VcenterVMSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVm

`func (o *VcenterVMSummary) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VcenterVMSummary) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VcenterVMSummary) SetVm(v string)`

SetVm sets Vm field to given value.


### GetName

`func (o *VcenterVMSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVMSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVMSummary) SetName(v string)`

SetName sets Name field to given value.


### GetPowerState

`func (o *VcenterVMSummary) GetPowerState() VcenterVmPowerState`

GetPowerState returns the PowerState field if non-nil, zero value otherwise.

### GetPowerStateOk

`func (o *VcenterVMSummary) GetPowerStateOk() (*VcenterVmPowerState, bool)`

GetPowerStateOk returns a tuple with the PowerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerState

`func (o *VcenterVMSummary) SetPowerState(v VcenterVmPowerState)`

SetPowerState sets PowerState field to given value.


### GetCpuCount

`func (o *VcenterVMSummary) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *VcenterVMSummary) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *VcenterVMSummary) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *VcenterVMSummary) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetMemorySizeMiB

`func (o *VcenterVMSummary) GetMemorySizeMiB() int64`

GetMemorySizeMiB returns the MemorySizeMiB field if non-nil, zero value otherwise.

### GetMemorySizeMiBOk

`func (o *VcenterVMSummary) GetMemorySizeMiBOk() (*int64, bool)`

GetMemorySizeMiBOk returns a tuple with the MemorySizeMiB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeMiB

`func (o *VcenterVMSummary) SetMemorySizeMiB(v int64)`

SetMemorySizeMiB sets MemorySizeMiB field to given value.

### HasMemorySizeMiB

`func (o *VcenterVMSummary) HasMemorySizeMiB() bool`

HasMemorySizeMiB returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


