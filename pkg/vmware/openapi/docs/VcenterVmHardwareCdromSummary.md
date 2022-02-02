# VcenterVmHardwareCdromSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cdrom** | **string** | Identifier of the virtual CD-ROM device. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Cdrom. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Cdrom. | 

## Methods

### NewVcenterVmHardwareCdromSummary

`func NewVcenterVmHardwareCdromSummary(cdrom string, ) *VcenterVmHardwareCdromSummary`

NewVcenterVmHardwareCdromSummary instantiates a new VcenterVmHardwareCdromSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareCdromSummaryWithDefaults

`func NewVcenterVmHardwareCdromSummaryWithDefaults() *VcenterVmHardwareCdromSummary`

NewVcenterVmHardwareCdromSummaryWithDefaults instantiates a new VcenterVmHardwareCdromSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCdrom

`func (o *VcenterVmHardwareCdromSummary) GetCdrom() string`

GetCdrom returns the Cdrom field if non-nil, zero value otherwise.

### GetCdromOk

`func (o *VcenterVmHardwareCdromSummary) GetCdromOk() (*string, bool)`

GetCdromOk returns a tuple with the Cdrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCdrom

`func (o *VcenterVmHardwareCdromSummary) SetCdrom(v string)`

SetCdrom sets Cdrom field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


