# VcenterVchaClusterNodeVmInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Vm** | **string** | The identifier of the virtual machine of the VCHA node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine:VCenter. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine:VCenter. | 
**BiosUuid** | **string** | BIOS UUID for the node. | 

## Methods

### NewVcenterVchaClusterNodeVmInfo

`func NewVcenterVchaClusterNodeVmInfo(vm string, biosUuid string, ) *VcenterVchaClusterNodeVmInfo`

NewVcenterVchaClusterNodeVmInfo instantiates a new VcenterVchaClusterNodeVmInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterNodeVmInfoWithDefaults

`func NewVcenterVchaClusterNodeVmInfoWithDefaults() *VcenterVchaClusterNodeVmInfo`

NewVcenterVchaClusterNodeVmInfoWithDefaults instantiates a new VcenterVchaClusterNodeVmInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVm

`func (o *VcenterVchaClusterNodeVmInfo) GetVm() string`

GetVm returns the Vm field if non-nil, zero value otherwise.

### GetVmOk

`func (o *VcenterVchaClusterNodeVmInfo) GetVmOk() (*string, bool)`

GetVmOk returns a tuple with the Vm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVm

`func (o *VcenterVchaClusterNodeVmInfo) SetVm(v string)`

SetVm sets Vm field to given value.


### GetBiosUuid

`func (o *VcenterVchaClusterNodeVmInfo) GetBiosUuid() string`

GetBiosUuid returns the BiosUuid field if non-nil, zero value otherwise.

### GetBiosUuidOk

`func (o *VcenterVchaClusterNodeVmInfo) GetBiosUuidOk() (*string, bool)`

GetBiosUuidOk returns a tuple with the BiosUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBiosUuid

`func (o *VcenterVchaClusterNodeVmInfo) SetBiosUuid(v string)`

SetBiosUuid sets BiosUuid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


