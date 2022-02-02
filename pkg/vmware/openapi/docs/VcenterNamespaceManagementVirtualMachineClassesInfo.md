# VcenterNamespaceManagementVirtualMachineClassesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Identifier for the VM class. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.VirtualMachineClass. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.VirtualMachineClass. | 
**CpuCount** | **int64** | The number of CPUs configured for virtual machine of this class. | 
**CpuReservation** | Pointer to **int64** | The percentage of total available CPUs reserved for a virtual machine. We multiply this percentage by the minimum frequency amongst all the cluster nodes to get the CPU reservation that is specified to vSphere in MHz. If unset, no CPU reservation is requested for the virtual machine. | [optional] 
**MemoryMB** | **int64** | The amount of memory in MB configured for virtual machine of this class. | 
**MemoryReservation** | Pointer to **int64** | The percentage of available memory reserved for a virtual machine of this class. If unset, no memory reservation is requested for virtual machine. | [optional] 
**Description** | **string** | Description of the VM class. | 
**Namespaces** | **[]string** | Set of Namespaces associated with this VM class. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.namespaces.Instance. | 
**Vms** | **[]string** | Set of virtual machines deployed for VM class. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: VirtualMachine. | 
**ConfigStatus** | [**VcenterNamespaceManagementVirtualMachineClassesConfigStatus**](VcenterNamespaceManagementVirtualMachineClassesConfigStatus.md) |  | 
**Messages** | [**[]VcenterNamespaceManagementVirtualMachineClassesMessage**](VcenterNamespaceManagementVirtualMachineClassesMessage.md) | Current set of messages associated with the object. | 

## Methods

### NewVcenterNamespaceManagementVirtualMachineClassesInfo

`func NewVcenterNamespaceManagementVirtualMachineClassesInfo(id string, cpuCount int64, memoryMB int64, description string, namespaces []string, vms []string, configStatus VcenterNamespaceManagementVirtualMachineClassesConfigStatus, messages []VcenterNamespaceManagementVirtualMachineClassesMessage, ) *VcenterNamespaceManagementVirtualMachineClassesInfo`

NewVcenterNamespaceManagementVirtualMachineClassesInfo instantiates a new VcenterNamespaceManagementVirtualMachineClassesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementVirtualMachineClassesInfoWithDefaults

`func NewVcenterNamespaceManagementVirtualMachineClassesInfoWithDefaults() *VcenterNamespaceManagementVirtualMachineClassesInfo`

NewVcenterNamespaceManagementVirtualMachineClassesInfoWithDefaults instantiates a new VcenterNamespaceManagementVirtualMachineClassesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetId(v string)`

SetId sets Id field to given value.


### GetCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetCpuCount() int64`

GetCpuCount returns the CpuCount field if non-nil, zero value otherwise.

### GetCpuCountOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetCpuCountOk() (*int64, bool)`

GetCpuCountOk returns a tuple with the CpuCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCount

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetCpuCount(v int64)`

SetCpuCount sets CpuCount field to given value.


### GetCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetCpuReservation() int64`

GetCpuReservation returns the CpuReservation field if non-nil, zero value otherwise.

### GetCpuReservationOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetCpuReservationOk() (*int64, bool)`

GetCpuReservationOk returns a tuple with the CpuReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetCpuReservation(v int64)`

SetCpuReservation sets CpuReservation field to given value.

### HasCpuReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) HasCpuReservation() bool`

HasCpuReservation returns a boolean if a field has been set.

### GetMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetMemoryMB() int64`

GetMemoryMB returns the MemoryMB field if non-nil, zero value otherwise.

### GetMemoryMBOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetMemoryMBOk() (*int64, bool)`

GetMemoryMBOk returns a tuple with the MemoryMB field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryMB

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetMemoryMB(v int64)`

SetMemoryMB sets MemoryMB field to given value.


### GetMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetMemoryReservation() int64`

GetMemoryReservation returns the MemoryReservation field if non-nil, zero value otherwise.

### GetMemoryReservationOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetMemoryReservationOk() (*int64, bool)`

GetMemoryReservationOk returns a tuple with the MemoryReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetMemoryReservation(v int64)`

SetMemoryReservation sets MemoryReservation field to given value.

### HasMemoryReservation

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) HasMemoryReservation() bool`

HasMemoryReservation returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetNamespaces

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetNamespaces() []string`

GetNamespaces returns the Namespaces field if non-nil, zero value otherwise.

### GetNamespacesOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetNamespacesOk() (*[]string, bool)`

GetNamespacesOk returns a tuple with the Namespaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaces

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetNamespaces(v []string)`

SetNamespaces sets Namespaces field to given value.


### GetVms

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetVms() []string`

GetVms returns the Vms field if non-nil, zero value otherwise.

### GetVmsOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetVmsOk() (*[]string, bool)`

GetVmsOk returns a tuple with the Vms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVms

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetVms(v []string)`

SetVms sets Vms field to given value.


### GetConfigStatus

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetConfigStatus() VcenterNamespaceManagementVirtualMachineClassesConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetConfigStatusOk() (*VcenterNamespaceManagementVirtualMachineClassesConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetConfigStatus(v VcenterNamespaceManagementVirtualMachineClassesConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetMessages

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetMessages() []VcenterNamespaceManagementVirtualMachineClassesMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) GetMessagesOk() (*[]VcenterNamespaceManagementVirtualMachineClassesMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespaceManagementVirtualMachineClassesInfo) SetMessages(v []VcenterNamespaceManagementVirtualMachineClassesMessage)`

SetMessages sets Messages field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


