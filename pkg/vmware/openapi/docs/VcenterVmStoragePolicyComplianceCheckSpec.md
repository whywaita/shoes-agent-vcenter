# VcenterVmStoragePolicyComplianceCheckSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmHome** | **bool** | Invoke compliance check on the virtual machine home directory if set to true. | 
**Disks** | Pointer to **[]string** | Identifiers of the virtual machine&#39;s virtual disks for which compliance should be checked. If unset or empty, compliance check is invoked on all the associated disks. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.vm.hardware.Disk. | [optional] 

## Methods

### NewVcenterVmStoragePolicyComplianceCheckSpec

`func NewVcenterVmStoragePolicyComplianceCheckSpec(vmHome bool, ) *VcenterVmStoragePolicyComplianceCheckSpec`

NewVcenterVmStoragePolicyComplianceCheckSpec instantiates a new VcenterVmStoragePolicyComplianceCheckSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmStoragePolicyComplianceCheckSpecWithDefaults

`func NewVcenterVmStoragePolicyComplianceCheckSpecWithDefaults() *VcenterVmStoragePolicyComplianceCheckSpec`

NewVcenterVmStoragePolicyComplianceCheckSpecWithDefaults instantiates a new VcenterVmStoragePolicyComplianceCheckSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmHome

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) GetVmHome() bool`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) GetVmHomeOk() (*bool, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) SetVmHome(v bool)`

SetVmHome sets VmHome field to given value.


### GetDisks

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) GetDisks() []string`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) GetDisksOk() (*[]string, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) SetDisks(v []string)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VcenterVmStoragePolicyComplianceCheckSpec) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


