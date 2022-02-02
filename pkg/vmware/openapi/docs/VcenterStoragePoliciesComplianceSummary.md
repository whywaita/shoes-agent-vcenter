# VcenterStoragePoliciesComplianceSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vm** | **string** | Identifier of virtual machine When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine. | 
**VmHome** | Pointer to [**VcenterStoragePoliciesComplianceStatus**](VcenterStoragePoliciesComplianceStatus.md) |  | [optional] 
**Disks** | Pointer to [**[]VcenterStoragePoliciesComplianceSummaryDisks**](VcenterStoragePoliciesComplianceSummaryDisks.md) | List of the virtual hard disk. If unset or empty, virtual machine entity does not have any disks or its disks are not associated with a storage policy. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | [optional] 

## Methods

### NewVcenterStoragePoliciesComplianceSummary

`func NewVcenterStoragePoliciesComplianceSummary(vm string, ) *VcenterStoragePoliciesComplianceSummary`

NewVcenterStoragePoliciesComplianceSummary instantiates a new VcenterStoragePoliciesComplianceSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesComplianceSummaryWithDefaults

`func NewVcenterStoragePoliciesComplianceSummaryWithDefaults() *VcenterStoragePoliciesComplianceSummary`

NewVcenterStoragePoliciesComplianceSummaryWithDefaults instantiates a new VcenterStoragePoliciesComplianceSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVm

`func (o *VcenterStoragePoliciesComplianceSummary) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VcenterStoragePoliciesComplianceSummary) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VcenterStoragePoliciesComplianceSummary) SetVm(v string)`

SetVm sets Vm field to given value.


### GetVmHome

`func (o *VcenterStoragePoliciesComplianceSummary) GetVmHome() VcenterStoragePoliciesComplianceStatus`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterStoragePoliciesComplianceSummary) GetVmHomeOk() (*VcenterStoragePoliciesComplianceStatus, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterStoragePoliciesComplianceSummary) SetVmHome(v VcenterStoragePoliciesComplianceStatus)`

SetVmHome sets VmHome field to given value.

### HasVmHome

`func (o *VcenterStoragePoliciesComplianceSummary) HasVmHome() bool`

HasVmHome returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterStoragePoliciesComplianceSummary) GetDisks() []VcenterStoragePoliciesComplianceSummaryDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterStoragePoliciesComplianceSummary) GetDisksOk() (*[]VcenterStoragePoliciesComplianceSummaryDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterStoragePoliciesComplianceSummary) SetDisks(v []VcenterStoragePoliciesComplianceSummaryDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VcenterStoragePoliciesComplianceSummary) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


