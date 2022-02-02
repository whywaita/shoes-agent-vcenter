# VcenterContentRegistriesHarborProjectsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Harbor project. Should be between 1-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and &#39;-&#39;. Must be starting with characters or numbers, with the &#39;-&#39; character allowed anywhere except the first or last character. | 
**Scope** | [**VcenterContentRegistriesHarborProjectsScope**](VcenterContentRegistriesHarborProjectsScope.md) |  | 

## Methods

### NewVcenterContentRegistriesHarborProjectsCreateSpec

`func NewVcenterContentRegistriesHarborProjectsCreateSpec(name string, scope VcenterContentRegistriesHarborProjectsScope, ) *VcenterContentRegistriesHarborProjectsCreateSpec`

NewVcenterContentRegistriesHarborProjectsCreateSpec instantiates a new VcenterContentRegistriesHarborProjectsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborProjectsCreateSpecWithDefaults

`func NewVcenterContentRegistriesHarborProjectsCreateSpecWithDefaults() *VcenterContentRegistriesHarborProjectsCreateSpec`

NewVcenterContentRegistriesHarborProjectsCreateSpecWithDefaults instantiates a new VcenterContentRegistriesHarborProjectsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterContentRegistriesHarborProjectsCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterContentRegistriesHarborProjectsCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterContentRegistriesHarborProjectsCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *VcenterContentRegistriesHarborProjectsCreateSpec) GetScope() VcenterContentRegistriesHarborProjectsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VcenterContentRegistriesHarborProjectsCreateSpec) GetScopeOk() (*VcenterContentRegistriesHarborProjectsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VcenterContentRegistriesHarborProjectsCreateSpec) SetScope(v VcenterContentRegistriesHarborProjectsScope)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


