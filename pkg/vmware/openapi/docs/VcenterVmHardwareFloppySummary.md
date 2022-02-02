# VcenterVmHardwareFloppySummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Floppy** | **string** | Identifier of the virtual floppy drive. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Floppy. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Floppy. | 

## Methods

### NewVcenterVmHardwareFloppySummary

`func NewVcenterVmHardwareFloppySummary(floppy string, ) *VcenterVmHardwareFloppySummary`

NewVcenterVmHardwareFloppySummary instantiates a new VcenterVmHardwareFloppySummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareFloppySummaryWithDefaults

`func NewVcenterVmHardwareFloppySummaryWithDefaults() *VcenterVmHardwareFloppySummary`

NewVcenterVmHardwareFloppySummaryWithDefaults instantiates a new VcenterVmHardwareFloppySummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFloppy

`func (o *VcenterVmHardwareFloppySummary) GetFloppy() string`

GetFloppy returns the Floppy field if non-nil, zero value otherwise.

### GetFloppyOk

`func (o *VcenterVmHardwareFloppySummary) GetFloppyOk() (*string, bool)`

GetFloppyOk returns a tuple with the Floppy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloppy

`func (o *VcenterVmHardwareFloppySummary) SetFloppy(v string)`

SetFloppy sets Floppy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


