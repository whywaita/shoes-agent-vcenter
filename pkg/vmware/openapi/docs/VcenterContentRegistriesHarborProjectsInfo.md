# VcenterContentRegistriesHarborProjectsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the Harbor project. Should be between 1-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and &#39;-&#39;. Must be starting with characters or numbers, with the &#39;-&#39; character allowed anywhere except the first or last character. | 
**ConfigStatus** | [**VcenterContentRegistriesHarborProjectsConfigStatus**](VcenterContentRegistriesHarborProjectsConfigStatus.md) |  | 
**Scope** | [**VcenterContentRegistriesHarborProjectsScope**](VcenterContentRegistriesHarborProjectsScope.md) |  | 
**CreationTime** | **time.Time** | The date and time when the harbor project creation API was triggered and project identifier generated. | 
**UpdateTime** | Pointer to **time.Time** | The date and time when the harbor project purge API was triggered. In case no purge was triggered, {@name #updateTime} is same as {@name #creationTime}. | [optional] 
**AccessUrl** | Pointer to **string** | URL to access the harbor project through docker client. | [optional] 
**Message** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterContentRegistriesHarborProjectsInfo

`func NewVcenterContentRegistriesHarborProjectsInfo(name string, configStatus VcenterContentRegistriesHarborProjectsConfigStatus, scope VcenterContentRegistriesHarborProjectsScope, creationTime time.Time, ) *VcenterContentRegistriesHarborProjectsInfo`

NewVcenterContentRegistriesHarborProjectsInfo instantiates a new VcenterContentRegistriesHarborProjectsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterContentRegistriesHarborProjectsInfoWithDefaults

`func NewVcenterContentRegistriesHarborProjectsInfoWithDefaults() *VcenterContentRegistriesHarborProjectsInfo`

NewVcenterContentRegistriesHarborProjectsInfoWithDefaults instantiates a new VcenterContentRegistriesHarborProjectsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetName(v string)`

SetName sets Name field to given value.


### GetConfigStatus

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetConfigStatus() VcenterContentRegistriesHarborProjectsConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetConfigStatusOk() (*VcenterContentRegistriesHarborProjectsConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetConfigStatus(v VcenterContentRegistriesHarborProjectsConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetScope

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetScope() VcenterContentRegistriesHarborProjectsScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetScopeOk() (*VcenterContentRegistriesHarborProjectsScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetScope(v VcenterContentRegistriesHarborProjectsScope)`

SetScope sets Scope field to given value.


### GetCreationTime

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.


### GetUpdateTime

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetUpdateTime() time.Time`

GetUpdateTime returns the UpdateTime field if non-nil, zero value otherwise.

### GetUpdateTimeOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetUpdateTimeOk() (*time.Time, bool)`

GetUpdateTimeOk returns a tuple with the UpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTime

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetUpdateTime(v time.Time)`

SetUpdateTime sets UpdateTime field to given value.

### HasUpdateTime

`func (o *VcenterContentRegistriesHarborProjectsInfo) HasUpdateTime() bool`

HasUpdateTime returns a boolean if a field has been set.

### GetAccessUrl

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetAccessUrl() string`

GetAccessUrl returns the AccessUrl field if non-nil, zero value otherwise.

### GetAccessUrlOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetAccessUrlOk() (*string, bool)`

GetAccessUrlOk returns a tuple with the AccessUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessUrl

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetAccessUrl(v string)`

SetAccessUrl sets AccessUrl field to given value.

### HasAccessUrl

`func (o *VcenterContentRegistriesHarborProjectsInfo) HasAccessUrl() bool`

HasAccessUrl returns a boolean if a field has been set.

### GetMessage

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetMessage() VapiStdLocalizableMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VcenterContentRegistriesHarborProjectsInfo) GetMessageOk() (*VapiStdLocalizableMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VcenterContentRegistriesHarborProjectsInfo) SetMessage(v VapiStdLocalizableMessage)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *VcenterContentRegistriesHarborProjectsInfo) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


