# VcenterNamespaceManagementVirtualMachineClassesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier of the virtual machine class. This has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the &#39;-&#39; character allowed anywhere except the first or last character. This name is unique in this vCenter server. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.VirtualMachineClass. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.VirtualMachineClass. | 
**CpuCount** | **int64** | The number of CPUs configured for virtual machine of this class. | 
**CpuReservation** | Pointer to **int64** | The percentage of total available CPUs reserved for a virtual machine. We multiply this percentage by the minimum frequency amongst all the cluster nodes to get the CPU reservation that is specified to vSphere in MHz. If unset, no CPU reservation is requested for the virtual machine. | [optional] 
**MemoryMB** | **int64** | The amount of memory in MB configured for virtual machine of this class. | 
**MemoryReservation** | Pointer to **int64** | The percentage of available memory reserved for a virtual machine of this class. Memory reservation must be set to 100% for VM class with vGPU or Dynamic DirectPath I/O devices. If unset, no memory reservation is requested for virtual machine. | [optional] 
**Description** | Pointer to **string** | Description for the VM class. If unset, no description is added to the VM class. | [optional] 

## Methods

### NewVcenterNamespaceManagementVirtualMachineClassesCreateSpec

`func NewVcenterNamespaceManagementVirtualMachineClassesCreateSpec(id string, cpuCount int64, memoryMB int64, ) *VcenterNamespaceManagementVirtualMachineClassesCreateSpec`

NewVcenterNamespaceManagementVirtualMachineClassesCreateSpec instantiates a new VcenterNamespaceManagementVirtualMachineClassesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementVirtualMachineClassesCreateSpecWithDefaults

`func NewVcenterNamespaceManagementVirtualMachineClassesCreateSpecWithDefaults() *VcenterNamespaceManagementVirtualMachineClassesCreateSpec`

NewVcenterNamespaceManagementVirtualMachineClassesCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementVirtualMachineClassesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetId(v string)`

SetId sets Id field to given value.


### GetCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.


### GetCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuReservation() int64`

GetCpuReservation returns the CpuReservation field if non-nil, zero value otherwise.

### GetCpuReservationOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetCpuReservationOk() (*int64, bool)`

GetCpuReservationOk returns a tuple with the CpuReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetCpuReservation(v int64)`

SetCpuReservation sets CpuReservation field to given value.

### HasCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) HasCpuReservation() bool`

HasCpuReservation returns a boolean if a field has been set.

### GetMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryMB() int64`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryMBOk() (*int64, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetMemoryMB(v int64)`

SetMemoryMB sets MemoryMB field to given value.


### GetMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryReservation() int64`

GetMemoryReservation returns the MemoryReservation field if non-nil, zero value otherwise.

### GetMemoryReservationOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetMemoryReservationOk() (*int64, bool)`

GetMemoryReservationOk returns a tuple with the MemoryReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetMemoryReservation(v int64)`

SetMemoryReservation sets MemoryReservation field to given value.

### HasMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) HasMemoryReservation() bool`

HasMemoryReservation returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


