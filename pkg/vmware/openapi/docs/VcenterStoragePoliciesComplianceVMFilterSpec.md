# VcenterStoragePoliciesComplianceVMFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**[]VcenterStoragePoliciesComplianceVMStatus**](VcenterStoragePoliciesComplianceVMStatus.md) | Compliance Status that a virtual machine must have to match the filter. Atleast one status must be specified. | 
**Vms** | Pointer to **[]string** | Identifiers of virtual machines that can match the filter If unset or empty, virtual machines with any identifier matches the filter When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: VirtualMachine. | [optional] 

## Methods

### NewVcenterStoragePoliciesComplianceVMFilterSpec

`func NewVcenterStoragePoliciesComplianceVMFilterSpec(status []VcenterStoragePoliciesComplianceVMStatus, ) *VcenterStoragePoliciesComplianceVMFilterSpec`

NewVcenterStoragePoliciesComplianceVMFilterSpec instantiates a new VcenterStoragePoliciesComplianceVMFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterStoragePoliciesComplianceVMFilterSpecWithDefaults

`func NewVcenterStoragePoliciesComplianceVMFilterSpecWithDefaults() *VcenterStoragePoliciesComplianceVMFilterSpec`

NewVcenterStoragePoliciesComplianceVMFilterSpecWithDefaults instantiates a new VcenterStoragePoliciesComplianceVMFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) GetStatus() []VcenterStoragePoliciesComplianceVMStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) GetStatusOk() (*[]VcenterStoragePoliciesComplianceVMStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) SetStatus(v []VcenterStoragePoliciesComplianceVMStatus)`

SetStatus sets Status field to given value.


### GetVms

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) GetVms() []string`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) GetVmsOk() (*[]string, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) SetVms(v []string)`

SetVms sets Vms field to given value.

### HasVms

`func (o *VcenterStoragePoliciesComplianceVMFilterSpec) HasVms() bool`

HasVms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


