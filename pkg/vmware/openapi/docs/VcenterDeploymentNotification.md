# VcenterDeploymentNotification

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The notification id. | 
**Time** | Pointer to **time.Time** | The time the notification was raised/found. Only set if the time information is available. | [optional] 
**Message** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 
**Resolution** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterDeploymentNotification

`func NewVcenterDeploymentNotification(id string, message VapiStdLocalizableMessage, ) *VcenterDeploymentNotification`

NewVcenterDeploymentNotification instantiates a new VcenterDeploymentNotification object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentNotificationWithDefaults

`func NewVcenterDeploymentNotificationWithDefaults() *VcenterDeploymentNotification`

NewVcenterDeploymentNotificationWithDefaults instantiates a new VcenterDeploymentNotification object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VcenterDeploymentNotification) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VcenterDeploymentNotification) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VcenterDeploymentNotification) SetId(v string)`

SetId sets Id field to given value.


### GetTime

`func (o *VcenterDeploymentNotification) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *VcenterDeploymentNotification) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *VcenterDeploymentNotification) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *VcenterDeploymentNotification) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetMessage

`func (o *VcenterDeploymentNotification) GetMessage() VapiStdLocalizableMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *VcenterDeploymentNotification) GetMessageOk() (*VapiStdLocalizableMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *VcenterDeploymentNotification) SetMessage(v VapiStdLocalizableMessage)`

SetMessage sets Message field to given value.


### GetResolution

`func (o *VcenterDeploymentNotification) GetResolution() VapiStdLocalizableMessage`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VcenterDeploymentNotification) GetResolutionOk() (*VapiStdLocalizableMessage, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VcenterDeploymentNotification) SetResolution(v VapiStdLocalizableMessage)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VcenterDeploymentNotification) HasResolution() bool`

HasResolution returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


