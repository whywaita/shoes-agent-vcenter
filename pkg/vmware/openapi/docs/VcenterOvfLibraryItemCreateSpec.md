# VcenterOvfLibraryItemCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name to use in the OVF descriptor stored in the library item. | [optional] 
**Description** | Pointer to **string** | Description to use in the OVF descriptor stored in the library item. | [optional] 
**Flags** | Pointer to **[]string** | Flags to use for OVF package creation. The supported flags can be obtained using {@link ExportFlag#list}. | [optional] 

## Methods

### NewVcenterOvfLibraryItemCreateSpec

`func NewVcenterOvfLibraryItemCreateSpec() *VcenterOvfLibraryItemCreateSpec`

NewVcenterOvfLibraryItemCreateSpec instantiates a new VcenterOvfLibraryItemCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterOvfLibraryItemCreateSpecWithDefaults

`func NewVcenterOvfLibraryItemCreateSpecWithDefaults() *VcenterOvfLibraryItemCreateSpec`

NewVcenterOvfLibraryItemCreateSpecWithDefaults instantiates a new VcenterOvfLibraryItemCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterOvfLibraryItemCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterOvfLibraryItemCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterOvfLibraryItemCreateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterOvfLibraryItemCreateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterOvfLibraryItemCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterOvfLibraryItemCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterOvfLibraryItemCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterOvfLibraryItemCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFlags

`func (o *VcenterOvfLibraryItemCreateSpec) GetFlags() []string`

GetFlags returns the Flags field if non-nil, zero value otherwise.

### GetFlagsOk

`func (o *VcenterOvfLibraryItemCreateSpec) GetFlagsOk() (*[]string, bool)`

GetFlagsOk returns a tuple with the Flags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlags

`func (o *VcenterOvfLibraryItemCreateSpec) SetFlags(v []string)`

SetFlags sets Flags field to given value.

### HasFlags

`func (o *VcenterOvfLibraryItemCreateSpec) HasFlags() bool`

HasFlags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


