# VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the virtual machine to check out of the library item. | [optional] 
**Placement** | Pointer to [**VcenterVmTemplateLibraryItemsCheckOutsPlacementSpec**](VcenterVmTemplateLibraryItemsCheckOutsPlacementSpec.md) |  | [optional] 
**PoweredOn** | Pointer to **bool** | Specifies whether the virtual machine should be powered on after check out. | [optional] 

## Methods

### NewVcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec

`func NewVcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec() *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec`

NewVcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec instantiates a new VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVmTemplateLibraryItemsCheckOutsCheckOutSpecWithDefaults

`func NewVcenterVmTemplateLibraryItemsCheckOutsCheckOutSpecWithDefaults() *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec`

NewVcenterVmTemplateLibraryItemsCheckOutsCheckOutSpecWithDefaults instantiates a new VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlacement

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) GetPlacement() VcenterVmTemplateLibraryItemsCheckOutsPlacementSpec`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) GetPlacementOk() (*VcenterVmTemplateLibraryItemsCheckOutsPlacementSpec, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) SetPlacement(v VcenterVmTemplateLibraryItemsCheckOutsPlacementSpec)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.

### GetPoweredOn

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) GetPoweredOn() bool`

GetPoweredOn returns the PoweredOn field if non-nil, zero value otherwise.

### GetPoweredOnOk

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) GetPoweredOnOk() (*bool, bool)`

GetPoweredOnOk returns a tuple with the PoweredOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoweredOn

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) SetPoweredOn(v bool)`

SetPoweredOn sets PoweredOn field to given value.

### HasPoweredOn

`func (o *VcenterVmTemplateLibraryItemsCheckOutsCheckOutSpec) HasPoweredOn() bool`

HasPoweredOn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


