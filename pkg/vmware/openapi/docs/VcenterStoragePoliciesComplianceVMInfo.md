# VcenterStoragePoliciesComplianceVMInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmHome** | Pointer to [**VcenterStoragePoliciesComplianceVMStatus**](VcenterStoragePoliciesComplianceVMStatus.md) |  | [optional] 
**Disks** | [**[]VcenterStoragePoliciesComplianceVMInfoDisks**](VcenterStoragePoliciesComplianceVMInfoDisks.md) | A Map of virtual disks and their compliance status If empty, the virtual machine does not have any disks or its disks are not associated with a storage policy. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | 

## Methods

### NewVcenterStoragePoliciesComplianceVMInfo

`func NewVcenterStoragePoliciesComplianceVMInfo(disks []VcenterStoragePoliciesComplianceVMInfoDisks, ) *VcenterStoragePoliciesComplianceVMInfo`

NewVcenterStoragePoliciesComplianceVMInfo instantiates a new VcenterStoragePoliciesComplianceVMInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesComplianceVMInfoWithDefaults

`func NewVcenterStoragePoliciesComplianceVMInfoWithDefaults() *VcenterStoragePoliciesComplianceVMInfo`

NewVcenterStoragePoliciesComplianceVMInfoWithDefaults instantiates a new VcenterStoragePoliciesComplianceVMInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmHome

`func (o *VcenterStoragePoliciesComplianceVMInfo) GetVmHome() VcenterStoragePoliciesComplianceVMStatus`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterStoragePoliciesComplianceVMInfo) GetVmHomeOk() (*VcenterStoragePoliciesComplianceVMStatus, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterStoragePoliciesComplianceVMInfo) SetVmHome(v VcenterStoragePoliciesComplianceVMStatus)`

SetVmHome sets VmHome field to given value.

### HasVmHome

`func (o *VcenterStoragePoliciesComplianceVMInfo) HasVmHome() bool`

HasVmHome returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterStoragePoliciesComplianceVMInfo) GetDisks() []VcenterStoragePoliciesComplianceVMInfoDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterStoragePoliciesComplianceVMInfo) GetDisksOk() (*[]VcenterStoragePoliciesComplianceVMInfoDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterStoragePoliciesComplianceVMInfo) SetDisks(v []VcenterStoragePoliciesComplianceVMInfoDisks)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


