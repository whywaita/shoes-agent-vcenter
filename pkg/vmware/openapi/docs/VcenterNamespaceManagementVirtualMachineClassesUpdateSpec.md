# VcenterNamespaceManagementVirtualMachineClassesUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuCount** | Pointer to **int64** | The number of CPUs configured for virtual machine of this class. If unset the current value the will not be modified. | [optional] 
**CpuReservation** | Pointer to **int64** | The percentage of total available CPUs reserved for a virtual machine. We multiply this percentage by the minimum frequency amongst all the cluster nodes to get the CPU reservation that is specified to vSphere in MHz. If unset, no CPU reservation is requested for the virtual machine. | [optional] 
**MemoryMB** | Pointer to **int64** | The amount of memory in MB configured for virtual machine of this class. If unset the current value the will not be modified. | [optional] 
**MemoryReservation** | Pointer to **int64** | The percentage of available memory reserved for a virtual machine of this class. Memory reservation must be set to 100% for VM class with vGPU or Dynamic DirectPath I/O devices. If unset, no memory reservation is requested for virtual machine. | [optional] 
**Description** | Pointer to **string** | Description for the VM class. If unset, description is not updated. | [optional] 

## Methods

### NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpec

`func NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpec() *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec`

NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpec instantiates a new VcenterNamespaceManagementVirtualMachineClassesUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpecWithDefaults

`func NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpecWithDefaults() *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec`

NewVcenterNamespaceManagementVirtualMachineClassesUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementVirtualMachineClassesUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.

### HasCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasCpuCount() bool`

HasCpuCount returns a boolean if a field has been set.

### GetCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuReservation() int64`

GetCpuReservation returns the CpuReservation field if non-nil, zero value otherwise.

### GetCpuReservationOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetCpuReservationOk() (*int64, bool)`

GetCpuReservationOk returns a tuple with the CpuReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetCpuReservation(v int64)`

SetCpuReservation sets CpuReservation field to given value.

### HasCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasCpuReservation() bool`

HasCpuReservation returns a boolean if a field has been set.

### GetMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryMB() int64`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryMBOk() (*int64, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetMemoryMB(v int64)`

SetMemoryMB sets MemoryMB field to given value.

### HasMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasMemoryMB() bool`

HasMemoryMB returns a boolean if a field has been set.

### GetMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryReservation() int64`

GetMemoryReservation returns the MemoryReservation field if non-nil, zero value otherwise.

### GetMemoryReservationOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetMemoryReservationOk() (*int64, bool)`

GetMemoryReservationOk returns a tuple with the MemoryReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetMemoryReservation(v int64)`

SetMemoryReservation sets MemoryReservation field to given value.

### HasMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasMemoryReservation() bool`

HasMemoryReservation returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


