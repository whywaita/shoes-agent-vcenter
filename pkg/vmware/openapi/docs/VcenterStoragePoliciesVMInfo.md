# VcenterStoragePoliciesVMInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmHome** | **bool** | Flag to indicate whether or not the virtual machine home is associated with the given storage policy. | 
**Disks** | **[]string** | List of the virtual disks that are associated with the given storage policy. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.vm.hardware.Disk. | 

## Methods

### NewVcenterStoragePoliciesVMInfo

`func NewVcenterStoragePoliciesVMInfo(vmHome bool, disks []string, ) *VcenterStoragePoliciesVMInfo`

NewVcenterStoragePoliciesVMInfo instantiates a new VcenterStoragePoliciesVMInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesVMInfoWithDefaults

`func NewVcenterStoragePoliciesVMInfoWithDefaults() *VcenterStoragePoliciesVMInfo`

NewVcenterStoragePoliciesVMInfoWithDefaults instantiates a new VcenterStoragePoliciesVMInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmHome

`func (o *VcenterStoragePoliciesVMInfo) GetVmHome() bool`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterStoragePoliciesVMInfo) GetVmHomeOk() (*bool, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterStoragePoliciesVMInfo) SetVmHome(v bool)`

SetVmHome sets VmHome field to given value.


### GetDisks

`func (o *VcenterStoragePoliciesVMInfo) GetDisks() []string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterStoragePoliciesVMInfo) GetDisksOk() (*[]string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterStoragePoliciesVMInfo) SetDisks(v []string)`

SetDisks sets Disks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


