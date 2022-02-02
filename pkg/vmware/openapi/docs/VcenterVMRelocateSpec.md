# VcenterVMRelocateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Placement** | Pointer to [**VcenterVMRelocatePlacementSpec**](VcenterVMRelocatePlacementSpec.md) |  | [optional] 
**Disks** | Pointer to [**[]VcenterVMRelocateSpecDisks**](VcenterVMRelocateSpecDisks.md) | Individual disk relocation map. If unset, all disks will migrate to the datastore specified in the VM.RelocatePlacementSpec.datastore field of VM.RelocateSpec.placement. When clients pass a value of this structure as a parameter, the key in the field map must be an identifier for the resource type: vcenter.vm.hardware.Disk. When operations return a value of this structure as a result, the key in the field map will be an identifier for the resource type: vcenter.vm.hardware.Disk. | [optional] 

## Methods

### NewVcenterVMRelocateSpec

`func NewVcenterVMRelocateSpec() *VcenterVMRelocateSpec`

NewVcenterVMRelocateSpec instantiates a new VcenterVMRelocateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVMRelocateSpecWithDefaults

`func NewVcenterVMRelocateSpecWithDefaults() *VcenterVMRelocateSpec`

NewVcenterVMRelocateSpecWithDefaults instantiates a new VcenterVMRelocateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlacement

`func (o *VcenterVMRelocateSpec) GetPlacement() VcenterVMRelocatePlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVMRelocateSpec) GetPlacementOk() (*VcenterVMRelocatePlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVMRelocateSpec) SetPlacement(v VcenterVMRelocatePlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVMRelocateSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetDisks

`func (o *VcenterVMRelocateSpec) GetDisks() []VcenterVMRelocateSpecDisks`

GetDisks returns the Disks field if non-nil, zero value otherwise.

### GetDisksOk

`func (o *VcenterVMRelocateSpec) GetDisksOk() (*[]VcenterVMRelocateSpecDisks, bool)`

GetDisksOk returns a tuple with the Disks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisks

`func (o *VcenterVMRelocateSpec) SetDisks(v []VcenterVMRelocateSpecDisks)`

SetDisks sets Disks field to given value.

### HasDisks

`func (o *VcenterVMRelocateSpec) HasDisks() bool`

HasDisks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


