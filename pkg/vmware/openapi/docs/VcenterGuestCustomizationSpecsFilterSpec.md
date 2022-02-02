# VcenterGuestCustomizationSpecsFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Names** | Pointer to **[]string** | Names that guest customization specifications must have to match the filter (see CustomizationSpecs.Summary.name). If unset or empty, guest customization specifications with any name match the filter. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.guest.CustomizationSpec. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.guest.CustomizationSpec. | [optional] 
**OSType** | Pointer to [**VcenterGuestCustomizationSpecsOsType**](VcenterGuestCustomizationSpecsOsType.md) |  | [optional] 

## Methods

### NewVcenterGuestCustomizationSpecsFilterSpec

`func NewVcenterGuestCustomizationSpecsFilterSpec() *VcenterGuestCustomizationSpecsFilterSpec`

NewVcenterGuestCustomizationSpecsFilterSpec instantiates a new VcenterGuestCustomizationSpecsFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCustomizationSpecsFilterSpecWithDefaults

`func NewVcenterGuestCustomizationSpecsFilterSpecWithDefaults() *VcenterGuestCustomizationSpecsFilterSpec`

NewVcenterGuestCustomizationSpecsFilterSpecWithDefaults instantiates a new VcenterGuestCustomizationSpecsFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNames

`func (o *VcenterGuestCustomizationSpecsFilterSpec) GetNames() []string`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *VcenterGuestCustomizationSpecsFilterSpec) GetNamesOk() (*[]string, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *VcenterGuestCustomizationSpecsFilterSpec) SetNames(v []string)`

SetNames sets Names field to given value.

### HasNames

`func (o *VcenterGuestCustomizationSpecsFilterSpec) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetOSType

`func (o *VcenterGuestCustomizationSpecsFilterSpec) GetOSType() VcenterGuestCustomizationSpecsOsType`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *VcenterGuestCustomizationSpecsFilterSpec) GetOSTypeOk() (*VcenterGuestCustomizationSpecsOsType, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *VcenterGuestCustomizationSpecsFilterSpec) SetOSType(v VcenterGuestCustomizationSpecsOsType)`

SetOSType sets OSType field to given value.

### HasOSType

`func (o *VcenterGuestCustomizationSpecsFilterSpec) HasOSType() bool`

HasOSType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


