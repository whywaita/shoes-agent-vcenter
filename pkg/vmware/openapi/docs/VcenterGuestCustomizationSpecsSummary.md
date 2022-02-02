# VcenterGuestCustomizationSpecsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the guest customization specification. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.guest.CustomizationSpec. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.guest.CustomizationSpec. | 
**Description** | **string** | Description of the guest customization specification. | 
**OSType** | [**VcenterGuestCustomizationSpecsOsType**](VcenterGuestCustomizationSpecsOsType.md) |  | 
**LastModified** | **time.Time** | Date and tme when this guest customization specification was last modified. | 

## Methods

### NewVcenterGuestCustomizationSpecsSummary

`func NewVcenterGuestCustomizationSpecsSummary(name string, description string, oSType VcenterGuestCustomizationSpecsOsType, lastModified time.Time, ) *VcenterGuestCustomizationSpecsSummary`

NewVcenterGuestCustomizationSpecsSummary instantiates a new VcenterGuestCustomizationSpecsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterGuestCustomizationSpecsSummaryWithDefaults

`func NewVcenterGuestCustomizationSpecsSummaryWithDefaults() *VcenterGuestCustomizationSpecsSummary`

NewVcenterGuestCustomizationSpecsSummaryWithDefaults instantiates a new VcenterGuestCustomizationSpecsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterGuestCustomizationSpecsSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterGuestCustomizationSpecsSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterGuestCustomizationSpecsSummary) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *VcenterGuestCustomizationSpecsSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterGuestCustomizationSpecsSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterGuestCustomizationSpecsSummary) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetOSType

`func (o *VcenterGuestCustomizationSpecsSummary) GetOSType() VcenterGuestCustomizationSpecsOsType`

GetOSType returns the OSType field if non-nil, zero value otherwise.

### GetOSTypeOk

`func (o *VcenterGuestCustomizationSpecsSummary) GetOSTypeOk() (*VcenterGuestCustomizationSpecsOsType, bool)`

GetOSTypeOk returns a tuple with the OSType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOSType

`func (o *VcenterGuestCustomizationSpecsSummary) SetOSType(v VcenterGuestCustomizationSpecsOsType)`

SetOSType sets OSType field to given value.


### GetLastModified

`func (o *VcenterGuestCustomizationSpecsSummary) GetLastModified() time.Time`

GetLastModified returns the LastModified field if non-nil, zero value otherwise.

### GetLastModifiedOk

`func (o *VcenterGuestCustomizationSpecsSummary) GetLastModifiedOk() (*time.Time, bool)`

GetLastModifiedOk returns a tuple with the LastModified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastModified

`func (o *VcenterGuestCustomizationSpecsSummary) SetLastModified(v time.Time)`

SetLastModified sets LastModified field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


