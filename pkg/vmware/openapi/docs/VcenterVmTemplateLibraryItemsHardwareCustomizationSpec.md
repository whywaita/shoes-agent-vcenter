# VcenterVmTemplateLibraryItemsHardwareCustomizationSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nics** | Pointer to [**[]VcenterVmTemplateLibraryItemsHardwareCustomizationSpecNics**](VcenterVmTemplateLibraryItemsHardwareCustomizationSpecNics.md) | Map of Ethernet network adapters to update. | [optional] 
**DisksToRemove** | Pointer to **[]string** | Idenfiers of disks to remove from the deployed virtual machine. | [optional] 
**DisksToUpdate** | Pointer to [**[]VcenterVmTemplateLibraryItemsHardwareCustomizationSpecDisksToUpdate**](VcenterVmTemplateLibraryItemsHardwareCustomizationSpecDisksToUpdate.md) | Disk update specification for individual disks in the deployed virtual machine. | [optional] 
**CpuUpdate** | Pointer to [**VcenterVmTemplateLibraryItemsCpuUpdateSpec**](VcenterVmTemplateLibraryItemsCpuUpdateSpec.md) |  | [optional] 
**MemoryUpdate** | Pointer to [**VcenterVmTemplateLibraryItemsMemoryUpdateSpec**](VcenterVmTemplateLibraryItemsMemoryUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsHardwareCustomizationSpec

`func NewVcenterVmTemplateLibraryItemsHardwareCustomizationSpec() *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec`

NewVcenterVmTemplateLibraryItemsHardwareCustomizationSpec instantiates a new VcenterVmTemplateLibraryItemsHardwareCustomizationSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsHardwareCustomizationSpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsHardwareCustomizationSpecWithDefaults() *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec`

NewVcenterVmTemplateLibraryItemsHardwareCustomizationSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsHardwareCustomizationSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNics

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetNics() []VcenterVmTemplateLibraryItemsHardwareCustomizationSpecNics`

GetNics returns the Nics field if non-nil, zero value otherwise.

### GetNicsOk

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetNicsOk() (*[]VcenterVmTemplateLibraryItemsHardwareCustomizationSpecNics, bool)`

GetNicsOk returns a tuple with the Nics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNics

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) SetNics(v []VcenterVmTemplateLibraryItemsHardwareCustomizationSpecNics)`

SetNics sets Nics field to given value.

### HasNics

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) HasNics() bool`

HasNics returns a boolean if a field has been set.

### GetDisksToRemove

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetDisksToRemove() []string`

GetDisksToRemove returns the DisksToRemove field if non-nil, zero value otherwise.

### GetDisksToRemoveOk

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetDisksToRemoveOk() (*[]string, bool)`

GetDisksToRemoveOk returns a tuple with the DisksToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksToRemove

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) SetDisksToRemove(v []string)`

SetDisksToRemove sets DisksToRemove field to given value.

### HasDisksToRemove

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) HasDisksToRemove() bool`

HasDisksToRemove returns a boolean if a field has been set.

### GetDisksToUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetDisksToUpdate() []VcenterVmTemplateLibraryItemsHardwareCustomizationSpecDisksToUpdate`

GetDisksToUpdate returns the DisksToUpdate field if non-nil, zero value otherwise.

### GetDisksToUpdateOk

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetDisksToUpdateOk() (*[]VcenterVmTemplateLibraryItemsHardwareCustomizationSpecDisksToUpdate, bool)`

GetDisksToUpdateOk returns a tuple with the DisksToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksToUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) SetDisksToUpdate(v []VcenterVmTemplateLibraryItemsHardwareCustomizationSpecDisksToUpdate)`

SetDisksToUpdate sets DisksToUpdate field to given value.

### HasDisksToUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) HasDisksToUpdate() bool`

HasDisksToUpdate returns a boolean if a field has been set.

### GetCpuUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetCpuUpdate() VcenterVmTemplateLibraryItemsCpuUpdateSpec`

GetCpuUpdate returns the CpuUpdate field if non-nil, zero value otherwise.

### GetCpuUpdateOk

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetCpuUpdateOk() (*VcenterVmTemplateLibraryItemsCpuUpdateSpec, bool)`

GetCpuUpdateOk returns a tuple with the CpuUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) SetCpuUpdate(v VcenterVmTemplateLibraryItemsCpuUpdateSpec)`

SetCpuUpdate sets CpuUpdate field to given value.

### HasCpuUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) HasCpuUpdate() bool`

HasCpuUpdate returns a boolean if a field has been set.

### GetMemoryUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetMemoryUpdate() VcenterVmTemplateLibraryItemsMemoryUpdateSpec`

GetMemoryUpdate returns the MemoryUpdate field if non-nil, zero value otherwise.

### GetMemoryUpdateOk

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) GetMemoryUpdateOk() (*VcenterVmTemplateLibraryItemsMemoryUpdateSpec, bool)`

GetMemoryUpdateOk returns a tuple with the MemoryUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) SetMemoryUpdate(v VcenterVmTemplateLibraryItemsMemoryUpdateSpec)`

SetMemoryUpdate sets MemoryUpdate field to given value.

### HasMemoryUpdate

`func (o *VcenterVmTemplateLibraryItemsHardwareCustomizationSpec) HasMemoryUpdate() bool`

HasMemoryUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


