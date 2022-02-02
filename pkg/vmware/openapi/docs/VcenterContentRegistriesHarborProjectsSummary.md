# VcenterContentRegistriesHarborProjectsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Project** | **string** | Identifier of the harbor project. | 
**Name** | **string** | Name of the Harbor project. Should be between 1-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and &#39;-&#39;. Must be starting with characters or numbers, with the &#39;-&#39; character allowed anywhere except the first or last character. | 
**Scope** | [**VcenterContentRegistriesHarborProjectsScope**](VcenterContentRegistriesHarborProjectsScope.md) |  | 

## Methods

### NewVcenterContentRegistriesHarborProjectsSummary

`func NewVcenterContentRegistriesHarborProjectsSummary(project string, name string, scope VcenterContentRegistriesHarborProjectsScope, ) *VcenterContentRegistriesHarborProjectsSummary`

NewVcenterContentRegistriesHarborProjectsSummary instantiates a new VcenterContentRegistriesHarborProjectsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborProjectsSummaryWithDefaults

`func NewVcenterContentRegistriesHarborProjectsSummaryWithDefaults() *VcenterContentRegistriesHarborProjectsSummary`

NewVcenterContentRegistriesHarborProjectsSummaryWithDefaults instantiates a new VcenterContentRegistriesHarborProjectsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProject

`func (o *VcenterContentRegistriesHarborProjectsSummary) GetProject() string`

GetProject returns the Project field if non-nil, zero value otherwise.

### GetProjectOk

`func (o *VcenterContentRegistriesHarborProjectsSummary) GetProjectOk() (*string, bool)`

GetProjectOk returns a tuple with the Project field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProject

`func (o *VcenterContentRegistriesHarborProjectsSummary) SetProject(v string)`

SetProject sets Project field to given value.


### GetName

`func (o *VcenterContentRegistriesHarborProjectsSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterContentRegistriesHarborProjectsSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterContentRegistriesHarborProjectsSummary) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *VcenterContentRegistriesHarborProjectsSummary) GetScope() VcenterContentRegistriesHarborProjectsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VcenterContentRegistriesHarborProjectsSummary) GetScopeOk() (*VcenterContentRegistriesHarborProjectsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VcenterContentRegistriesHarborProjectsSummary) SetScope(v VcenterContentRegistriesHarborProjectsScope)`

SetScope sets Scope field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


