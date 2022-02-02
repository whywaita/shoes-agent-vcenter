# VcenterVmStoragePolicyUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VmHome** | Pointer to [**VcenterVmStoragePolicyVmHomePolicySpec**](VcenterVmStoragePolicyVmHomePolicySpec.md) |  | [optional] 
**Disks** | Pointer to [**[]VcenterVmStoragePolicyUpdateSpecDisks**](VcenterVmStoragePolicyUpdateSpecDisks.md) | Storage policy or policies to be used when reconfiguring virtual machine diks. if unset the current storage policy is retained. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | [optional] 

## Methods

### NewVcenterVmStoragePolicyUpdateSpec

`func NewVcenterVmStoragePolicyUpdateSpec() *VcenterVmStoragePolicyUpdateSpec`

NewVcenterVmStoragePolicyUpdateSpec instantiates a new VcenterVmStoragePolicyUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmStoragePolicyUpdateSpecWithDefaults

`func NewVcenterVmStoragePolicyUpdateSpecWithDefaults() *VcenterVmStoragePolicyUpdateSpec`

NewVcenterVmStoragePolicyUpdateSpecWithDefaults instantiates a new VcenterVmStoragePolicyUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVmHome

`func (o *VcenterVmStoragePolicyUpdateSpec) GetVmHome() VcenterVmStoragePolicyVmHomePolicySpec`

GetVmHome returns the VmHome field if non-nil, zero value otherwise.

### GetVmHomeOk

`func (o *VcenterVmStoragePolicyUpdateSpec) GetVmHomeOk() (*VcenterVmStoragePolicyVmHomePolicySpec, bool)`

GetVmHomeOk returns a tuple with the VmHome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHome

`func (o *VcenterVmStoragePolicyUpdateSpec) SetVmHome(v VcenterVmStoragePolicyVmHomePolicySpec)`

SetVmHome sets VmHome field to given value.

### HasVmHome

`func (o *VcenterVmStoragePolicyUpdateSpec) HasVmHome() bool`

HasVmHome returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterVmStoragePolicyUpdateSpec) GetDisks() []VcenterVmStoragePolicyUpdateSpecDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVmStoragePolicyUpdateSpec) GetDisksOk() (*[]VcenterVmStoragePolicyUpdateSpecDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVmStoragePolicyUpdateSpec) SetDisks(v []VcenterVmStoragePolicyUpdateSpecDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VcenterVmStoragePolicyUpdateSpec) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


