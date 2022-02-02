# VcenterVmHardwareParallelSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | **string** | Identifier of the virtual parallel port. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.ParallelPort. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.ParallelPort. | 

## Methods

### NewVcenterVmHardwareParallelSummary

`func NewVcenterVmHardwareParallelSummary(port string, ) *VcenterVmHardwareParallelSummary`

NewVcenterVmHardwareParallelSummary instantiates a new VcenterVmHardwareParallelSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareParallelSummaryWithDefaults

`func NewVcenterVmHardwareParallelSummaryWithDefaults() *VcenterVmHardwareParallelSummary`

NewVcenterVmHardwareParallelSummaryWithDefaults instantiates a new VcenterVmHardwareParallelSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *VcenterVmHardwareParallelSummary) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *VcenterVmHardwareParallelSummary) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *VcenterVmHardwareParallelSummary) SetPort(v string)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


