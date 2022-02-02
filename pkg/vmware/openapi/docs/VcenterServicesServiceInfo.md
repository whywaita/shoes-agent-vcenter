# VcenterServicesServiceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NameKey** | **string** | Service name key. Can be used to lookup resource bundle | 
**DescriptionKey** | **string** | Service description key. Can be used to lookup resource bundle | 
**StartupType** | [**VcenterServicesServiceStartupType**](VcenterServicesServiceStartupType.md) |  | 
**State** | [**VcenterServicesServiceState**](VcenterServicesServiceState.md) |  | 
**Health** | Pointer to [**VcenterServicesServiceHealth**](VcenterServicesServiceHealth.md) |  | [optional] 
**HealthMessages** | Pointer to [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Localizable messages associated with the health of the service This field is optional and it is only relevant when the value of Service.Info.state is STARTED. | [optional] 

## Methods

### NewVcenterServicesServiceInfo

`func NewVcenterServicesServiceInfo(nameKey string, descriptionKey string, startupType VcenterServicesServiceStartupType, state VcenterServicesServiceState, ) *VcenterServicesServiceInfo`

NewVcenterServicesServiceInfo instantiates a new VcenterServicesServiceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterServicesServiceInfoWithDefaults

`func NewVcenterServicesServiceInfoWithDefaults() *VcenterServicesServiceInfo`

NewVcenterServicesServiceInfoWithDefaults instantiates a new VcenterServicesServiceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNameKey

`func (o *VcenterServicesServiceInfo) GetNameKey() string`

GetNameKey returns the NameKey field if non-nil, zero value otherwise.

### GetNameKeyOk

`func (o *VcenterServicesServiceInfo) GetNameKeyOk() (*string, bool)`

GetNameKeyOk returns a tuple with the NameKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameKey

`func (o *VcenterServicesServiceInfo) SetNameKey(v string)`

SetNameKey sets NameKey field to given value.


### GetDescriptionKey

`func (o *VcenterServicesServiceInfo) GetDescriptionKey() string`

GetDescriptionKey returns the DescriptionKey field if non-nil, zero value otherwise.

### GetDescriptionKeyOk

`func (o *VcenterServicesServiceInfo) GetDescriptionKeyOk() (*string, bool)`

GetDescriptionKeyOk returns a tuple with the DescriptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionKey

`func (o *VcenterServicesServiceInfo) SetDescriptionKey(v string)`

SetDescriptionKey sets DescriptionKey field to given value.


### GetStartupType

`func (o *VcenterServicesServiceInfo) GetStartupType() VcenterServicesServiceStartupType`

GetStartupType returns the StartupType field if non-nil, zero value otherwise.

### GetStartupTypeOk

`func (o *VcenterServicesServiceInfo) GetStartupTypeOk() (*VcenterServicesServiceStartupType, bool)`

GetStartupTypeOk returns a tuple with the StartupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartupType

`func (o *VcenterServicesServiceInfo) SetStartupType(v VcenterServicesServiceStartupType)`

SetStartupType sets StartupType field to given value.


### GetState

`func (o *VcenterServicesServiceInfo) GetState() VcenterServicesServiceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterServicesServiceInfo) GetStateOk() (*VcenterServicesServiceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterServicesServiceInfo) SetState(v VcenterServicesServiceState)`

SetState sets State field to given value.


### GetHealth

`func (o *VcenterServicesServiceInfo) GetHealth() VcenterServicesServiceHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterServicesServiceInfo) GetHealthOk() (*VcenterServicesServiceHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterServicesServiceInfo) SetHealth(v VcenterServicesServiceHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VcenterServicesServiceInfo) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetHealthMessages

`func (o *VcenterServicesServiceInfo) GetHealthMessages() []VapiStdLocalizableMessage`

GetHealthMessages returns the HealthMessages field if non-nil, zero value otherwise.

### GetHealthMessagesOk

`func (o *VcenterServicesServiceInfo) GetHealthMessagesOk() (*[]VapiStdLocalizableMessage, bool)`

GetHealthMessagesOk returns a tuple with the HealthMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthMessages

`func (o *VcenterServicesServiceInfo) SetHealthMessages(v []VapiStdLocalizableMessage)`

SetHealthMessages sets HealthMessages field to given value.

### HasHealthMessages

`func (o *VcenterServicesServiceInfo) HasHealthMessages() bool`

HasHealthMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


