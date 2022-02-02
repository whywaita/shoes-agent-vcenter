# VcenterVmHardwareDiskSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disk** | **string** | Identifier of the virtual Disk. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.vm.hardware.Disk. | 

## Methods

### NewVcenterVmHardwareDiskSummary

`func NewVcenterVmHardwareDiskSummary(disk string, ) *VcenterVmHardwareDiskSummary`

NewVcenterVmHardwareDiskSummary instantiates a new VcenterVmHardwareDiskSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmHardwareDiskSummaryWithDefaults

`func NewVcenterVmHardwareDiskSummaryWithDefaults() *VcenterVmHardwareDiskSummary`

NewVcenterVmHardwareDiskSummaryWithDefaults instantiates a new VcenterVmHardwareDiskSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisk

`func (o *VcenterVmHardwareDiskSummary) GetDisk() string`

GetDisk returns the Disk field if non-nil, zero value otherwise.

### GetDiskOk

`func (o *VcenterVmHardwareDiskSummary) GetDiskOk() (*string, bool)`

GetDiskOk returns a tuple with the Disk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisk

`func (o *VcenterVmHardwareDiskSummary) SetDisk(v string)`

SetDisk sets Disk field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


