# VcenterVmTemplateLibraryItemsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SourceVm** | **string** | Identifier of the source virtual machine to create the library item from. | 
**Name** | **string** | Name of the library item. | 
**Description** | Pointer to **string** | Description of the library item. | [optional] 
**Library** | **string** | Identifier of the library in which the new library item should be created. | 
**VmHomeStorage** | Pointer to [**VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage**](VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage.md) |  | [optional] 
**DiskStorage** | Pointer to [**VcenterVmTemplateLibraryItemsCreateSpecDiskStorage**](VcenterVmTemplateLibraryItemsCreateSpecDiskStorage.md) |  | [optional] 
**DiskStorageOverrides** | Pointer to [**[]VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides**](VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides.md) | Storage specification for individual disks in the virtual machine template. This is specified as a mapping between disk identifiers in the source virtual machine and their respective storage specifications. | [optional] 
**Placement** | Pointer to [**VcenterVmTemplateLibraryItemsCreatePlacementSpec**](VcenterVmTemplateLibraryItemsCreatePlacementSpec.md) |  | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsCreateSpec

`func NewVcenterVmTemplateLibraryItemsCreateSpec(sourceVm string, name string, library string, ) *VcenterVmTemplateLibraryItemsCreateSpec`

NewVcenterVmTemplateLibraryItemsCreateSpec instantiates a new VcenterVmTemplateLibraryItemsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsCreateSpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsCreateSpecWithDefaults() *VcenterVmTemplateLibraryItemsCreateSpec`

NewVcenterVmTemplateLibraryItemsCreateSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSourceVm

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetSourceVm() string`

GetSourceVm returns the SourceVm field if non-nil, zero value otherwise.

### GetSourceVmOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetSourceVmOk() (*string, bool)`

GetSourceVmOk returns a tuple with the SourceVm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceVm

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetSourceVm(v string)`

SetSourceVm sets SourceVm field to given value.


### GetName

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLibrary

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetLibrary() string`

GetLibrary returns the Library field if non-nil, zero value otherwise.

### GetLibraryOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetLibraryOk() (*string, bool)`

GetLibraryOk returns a tuple with the Library field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLibrary

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetLibrary(v string)`

SetLibrary sets Library field to given value.


### GetVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetVmHomeStorage() VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage`

GetVmHomeStorage returns the VmHomeStorage field if non-nil, zero value otherwise.

### GetVmHomeStorageOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetVmHomeStorageOk() (*VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage, bool)`

GetVmHomeStorageOk returns a tuple with the VmHomeStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetVmHomeStorage(v VcenterVmTemplateLibraryItemsCreateSpecVmHomeStorage)`

SetVmHomeStorage sets VmHomeStorage field to given value.

### HasVmHomeStorage

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) HasVmHomeStorage() bool`

HasVmHomeStorage returns a boolean if a field has been set.

### GetDiskStorage

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetDiskStorage() VcenterVmTemplateLibraryItemsCreateSpecDiskStorage`

GetDiskStorage returns the DiskStorage field if non-nil, zero value otherwise.

### GetDiskStorageOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetDiskStorageOk() (*VcenterVmTemplateLibraryItemsCreateSpecDiskStorage, bool)`

GetDiskStorageOk returns a tuple with the DiskStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorage

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetDiskStorage(v VcenterVmTemplateLibraryItemsCreateSpecDiskStorage)`

SetDiskStorage sets DiskStorage field to given value.

### HasDiskStorage

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) HasDiskStorage() bool`

HasDiskStorage returns a boolean if a field has been set.

### GetDiskStorageOverrides

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetDiskStorageOverrides() []VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides`

GetDiskStorageOverrides returns the DiskStorageOverrides field if non-nil, zero value otherwise.

### GetDiskStorageOverridesOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetDiskStorageOverridesOk() (*[]VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides, bool)`

GetDiskStorageOverridesOk returns a tuple with the DiskStorageOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiskStorageOverrides

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetDiskStorageOverrides(v []VcenterVmTemplateLibraryItemsCreateSpecDiskStorageOverrides)`

SetDiskStorageOverrides sets DiskStorageOverrides field to given value.

### HasDiskStorageOverrides

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) HasDiskStorageOverrides() bool`

HasDiskStorageOverrides returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetPlacement() VcenterVmTemplateLibraryItemsCreatePlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) GetPlacementOk() (*VcenterVmTemplateLibraryItemsCreatePlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) SetPlacement(v VcenterVmTemplateLibraryItemsCreatePlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVmTemplateLibraryItemsCreateSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


