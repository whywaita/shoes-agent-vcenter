# VcenterVMCloneSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | **string** | Virtual machine to clone from. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: VirtualMachine. When operations return a value of this structure as a result, the field will be an identifier for the resource type: VirtualMachine. | 
**Name** | **string** | Virtual machine name. | 
**Placement** | Pointer to [**VcenterVMClonePlacementSpec**](VcenterVMClonePlacementSpec.md) |  | [optional] 
**DisksToRemove** | Pointer to **[]string** | Set of Disks to Remove. If unset, all disks will be copied. If the same identifier is in VM.CloneSpec.disks-to-update InvalidArgument fault will be returned. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.vm.hardware.Disk. | [optional] 
**DisksToUpdate** | Pointer to [**[]VcenterVMCloneSpecDisksToUpdate**](VcenterVMCloneSpecDisksToUpdate.md) | Map of Disks to Update. If unset, all disks will copied to the datastore specified in the VM.ClonePlacementSpec.datastore field of VM.CloneSpec.placement. If the same identifier is in VM.CloneSpec.disks-to-remove InvalidArgument fault will be thrown. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | [optional] 
**PowerOn** | Pointer to **bool** | Attempt to perform a VM.CloneSpec.power-on after clone. If unset, the virtual machine will not be powered on. | [optional] 
**GuestCustomizationSpec** | Pointer to [**VcenterVMGuestCustomizationSpec**](VcenterVMGuestCustomizationSpec.md) |  | [optional] 

## Methods

### NewVcenterVMCloneSpec

`func NewVcenterVMCloneSpec(source string, name string, ) *VcenterVMCloneSpec`

NewVcenterVMCloneSpec instantiates a new VcenterVMCloneSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMCloneSpecWithDefaults

`func NewVcenterVMCloneSpecWithDefaults() *VcenterVMCloneSpec`

NewVcenterVMCloneSpecWithDefaults instantiates a new VcenterVMCloneSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *VcenterVMCloneSpec) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *VcenterVMCloneSpec) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *VcenterVMCloneSpec) SetSource(v string)`

SetSource sets Source field to given value.


### GetName

`func (o *VcenterVMCloneSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVMCloneSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVMCloneSpec) SetName(v string)`

SetName sets Name field to given value.


### GetPlacement

`func (o *VcenterVMCloneSpec) GetPlacement() VcenterVMClonePlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVMCloneSpec) GetPlacementOk() (*VcenterVMClonePlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVMCloneSpec) SetPlacement(v VcenterVMClonePlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVMCloneSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetDisksToRemove

`func (o *VcenterVMCloneSpec) GetDisksToRemove() []string`

GetDisksToRemove returns the DisksToRemove field if non-nil, zero value otherwise.

### GetDisksToRemoveOk

`func (o *VcenterVMCloneSpec) GetDisksToRemoveOk() (*[]string, bool)`

GetDisksToRemoveOk returns a tuple with the DisksToRemove field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksToRemove

`func (o *VcenterVMCloneSpec) SetDisksToRemove(v []string)`

SetDisksToRemove sets DisksToRemove field to given value.

### HasDisksToRemove

`func (o *VcenterVMCloneSpec) HasDisksToRemove() bool`

HasDisksToRemove returns a boolean if a field has been set.

### GetDisksToUpdate

`func (o *VcenterVMCloneSpec) GetDisksToUpdate() []VcenterVMCloneSpecDisksToUpdate`

GetDisksToUpdate returns the DisksToUpdate field if non-nil, zero value otherwise.

### GetDisksToUpdateOk

`func (o *VcenterVMCloneSpec) GetDisksToUpdateOk() (*[]VcenterVMCloneSpecDisksToUpdate, bool)`

GetDisksToUpdateOk returns a tuple with the DisksToUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisksToUpdate

`func (o *VcenterVMCloneSpec) SetDisksToUpdate(v []VcenterVMCloneSpecDisksToUpdate)`

SetDisksToUpdate sets DisksToUpdate field to given value.

### HasDisksToUpdate

`func (o *VcenterVMCloneSpec) HasDisksToUpdate() bool`

HasDisksToUpdate returns a boolean if a field has been set.

### GetPowerOn

`func (o *VcenterVMCloneSpec) GetPowerOn() bool`

GetPowerOn returns the PowerOn field if non-nil, zero value otherwise.

### GetPowerOnOk

`func (o *VcenterVMCloneSpec) GetPowerOnOk() (*bool, bool)`

GetPowerOnOk returns a tuple with the PowerOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPowerOn

`func (o *VcenterVMCloneSpec) SetPowerOn(v bool)`

SetPowerOn sets PowerOn field to given value.

### HasPowerOn

`func (o *VcenterVMCloneSpec) HasPowerOn() bool`

HasPowerOn returns a boolean if a field has been set.

### GetGuestCustomizationSpec

`func (o *VcenterVMCloneSpec) GetGuestCustomizationSpec() VcenterVMGuestCustomizationSpec`

GetGuestCustomizationSpec returns the GuestCustomizationSpec field if non-nil, zero value otherwise.

### GetGuestCustomizationSpecOk

`func (o *VcenterVMCloneSpec) GetGuestCustomizationSpecOk() (*VcenterVMGuestCustomizationSpec, bool)`

GetGuestCustomizationSpecOk returns a tuple with the GuestCustomizationSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestCustomizationSpec

`func (o *VcenterVMCloneSpec) SetGuestCustomizationSpec(v VcenterVMGuestCustomizationSpec)`

SetGuestCustomizationSpec sets GuestCustomizationSpec field to given value.

### HasGuestCustomizationSpec

`func (o *VcenterVMCloneSpec) HasGuestCustomizationSpec() bool`

HasGuestCustomizationSpec returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


