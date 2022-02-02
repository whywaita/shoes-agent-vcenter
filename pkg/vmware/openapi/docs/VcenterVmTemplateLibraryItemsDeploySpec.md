# VcenterVmTemplateLibraryItemsDeploySpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the deployed virtual machine. | 
**Description** | Pointer to **string** | Description of the deployed virtual machine. | [optional] 
**VmHomeStorage** | Pointer to [**VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage**](VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage.md) |  | [optional] 
**DiskStorage** | Pointer to [**VcenterVmTemplateLibraryItemsDeploySpecDiskStorage**](VcenterVmTemplateLibraryItemsDeploySpecDiskStorage.md) |  | [optional] 
**DiskStorageOverrides** | Pointer to [**[]VcenterVmTemplateLibraryItemsDeploySpecDiskStorageOverrides**](VcenterVmTemplateLibraryItemsDeploySpecDiskStorageOverrides.md) | Storage specification for individual disks in the deployed virtual machine. This is specified as a mapping between disk identifiers in the source virtual machine template contained in the library item and their storage specifications. | [optional] 
**Placement** | Pointer to [**VcenterVmTemplateLibraryItemsDeployPlacementSpec**](VcenterVmTemplateLibraryItemsDeployPlacementSpec.md) |  | [optional] 
**PoweredOn** | Pointer to **bool** | Specifies whether the deployed virtual machine should be powered on after deployment. | [optional] 
**GuestCustomization** | Pointer to [**VcenterVmTemplateLibraryItemsGuestCustomizationSpec**](VcenterVmTemplateLibraryItemsGuestCustomizationSpec.md) |  | [optional] 
**HardwareCustomization** | Pointer to [**VcenterVmTemplateLibraryItemsHardwareCustomizationSpec**](VcenterVmTemplateLibraryItemsHardwareCustomizationSpec.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsDeploySpec

`func NewVcenterVmTemplateLibraryItemsDeploySpec(name string, ) *VcenterVmTemplateLibraryItemsDeploySpec`

NewVcenterVmTemplateLibraryItemsDeploySpec instantiates a new VcenterVmTemplateLibraryItemsDeploySpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsDeploySpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsDeploySpecWithDefaults() *VcenterVmTemplateLibraryItemsDeploySpec`

NewVcenterVmTemplateLibraryItemsDeploySpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsDeploySpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetVmHomeStorage() VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage`

GetVmHomeStorage returns the VmHomeStorage field if non-nil, zero value otherwise.

### GetVmHomeStorageOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetVmHomeStorageOk() (*VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage, bool)`

GetVmHomeStorageOk returns a tuple with the VmHomeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetVmHomeStorage(v VcenterVmTemplateLibraryItemsDeploySpecVmHomeStorage)`

SetVmHomeStorage sets VmHomeStorage field to given value.

### HasVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasVmHomeStorage() bool`

HasVmHomeStorage returns a boolean if a field has been set.

### GetDiskStorage

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetDiskStorage() VcenterVmTemplateLibraryItemsDeploySpecDiskStorage`

GetDiskStorage returns the DiskStorage field if non-nil, zero value otherwise.

### GetDiskStorageOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetDiskStorageOk() (*VcenterVmTemplateLibraryItemsDeploySpecDiskStorage, bool)`

GetDiskStorageOk returns a tuple with the DiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorage

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetDiskStorage(v VcenterVmTemplateLibraryItemsDeploySpecDiskStorage)`

SetDiskStorage sets DiskStorage field to given value.

### HasDiskStorage

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasDiskStorage() bool`

HasDiskStorage returns a boolean if a field has been set.

### GetDiskStorageOverrides

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetDiskStorageOverrides() []VcenterVmTemplateLibraryItemsDeploySpecDiskStorageOverrides`

GetDiskStorageOverrides returns the DiskStorageOverrides field if non-nil, zero value otherwise.

### GetDiskStorageOverridesOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetDiskStorageOverridesOk() (*[]VcenterVmTemplateLibraryItemsDeploySpecDiskStorageOverrides, bool)`

GetDiskStorageOverridesOk returns a tuple with the DiskStorageOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageOverrides

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetDiskStorageOverrides(v []VcenterVmTemplateLibraryItemsDeploySpecDiskStorageOverrides)`

SetDiskStorageOverrides sets DiskStorageOverrides field to given value.

### HasDiskStorageOverrides

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasDiskStorageOverrides() bool`

HasDiskStorageOverrides returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetPlacement() VcenterVmTemplateLibraryItemsDeployPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetPlacementOk() (*VcenterVmTemplateLibraryItemsDeployPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetPlacement(v VcenterVmTemplateLibraryItemsDeployPlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetPoweredOn

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetPoweredOn() bool`

GetPoweredOn returns the PoweredOn field if non-nil, zero value otherwise.

### GetPoweredOnOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetPoweredOnOk() (*bool, bool)`

GetPoweredOnOk returns a tuple with the PoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredOn

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetPoweredOn(v bool)`

SetPoweredOn sets PoweredOn field to given value.

### HasPoweredOn

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasPoweredOn() bool`

HasPoweredOn returns a boolean if a field has been set.

### GetGuestCustomization

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetGuestCustomization() VcenterVmTemplateLibraryItemsGuestCustomizationSpec`

GetGuestCustomization returns the GuestCustomization field if non-nil, zero value otherwise.

### GetGuestCustomizationOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetGuestCustomizationOk() (*VcenterVmTemplateLibraryItemsGuestCustomizationSpec, bool)`

GetGuestCustomizationOk returns a tuple with the GuestCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomization

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetGuestCustomization(v VcenterVmTemplateLibraryItemsGuestCustomizationSpec)`

SetGuestCustomization sets GuestCustomization field to given value.

### HasGuestCustomization

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasGuestCustomization() bool`

HasGuestCustomization returns a boolean if a field has been set.

### GetHardwareCustomization

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetHardwareCustomization() VcenterVmTemplateLibraryItemsHardwareCustomizationSpec`

GetHardwareCustomization returns the HardwareCustomization field if non-nil, zero value otherwise.

### GetHardwareCustomizationOk

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) GetHardwareCustomizationOk() (*VcenterVmTemplateLibraryItemsHardwareCustomizationSpec, bool)`

GetHardwareCustomizationOk returns a tuple with the HardwareCustomization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardwareCustomization

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) SetHardwareCustomization(v VcenterVmTemplateLibraryItemsHardwareCustomizationSpec)`

SetHardwareCustomization sets HardwareCustomization field to given value.

### HasHardwareCustomization

`func (o *VcenterVmTemplateLibraryItemsDeploySpec) HasHardwareCustomization() bool`

HasHardwareCustomization returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


