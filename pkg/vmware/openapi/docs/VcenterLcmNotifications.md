# VcenterLcmNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**[]VcenterLcmNotification**](VcenterLcmNotification.md) | Info notification messages reported. Only set if an info was reported by the appliance task. | [optional] 
**Warnings** | Pointer to [**[]VcenterLcmNotification**](VcenterLcmNotification.md) | Warning notification messages reported. Only set if an warning was reported by the appliance task. | [optional] 
**Errors** | Pointer to [**[]VcenterLcmNotification**](VcenterLcmNotification.md) | Error notification messages reported. Only set if an error was reported by the appliance task. | [optional] 

## Methods

### NewVcenterLcmNotifications

`func NewVcenterLcmNotifications() *VcenterLcmNotifications`

NewVcenterLcmNotifications instantiates a new VcenterLcmNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmNotificationsWithDefaults

`func NewVcenterLcmNotificationsWithDefaults() *VcenterLcmNotifications`

NewVcenterLcmNotificationsWithDefaults instantiates a new VcenterLcmNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *VcenterLcmNotifications) GetInfo() []VcenterLcmNotification`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *VcenterLcmNotifications) GetInfoOk() (*[]VcenterLcmNotification, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *VcenterLcmNotifications) SetInfo(v []VcenterLcmNotification)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *VcenterLcmNotifications) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetWarnings

`func (o *VcenterLcmNotifications) GetWarnings() []VcenterLcmNotification`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VcenterLcmNotifications) GetWarningsOk() (*[]VcenterLcmNotification, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VcenterLcmNotifications) SetWarnings(v []VcenterLcmNotification)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VcenterLcmNotifications) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *VcenterLcmNotifications) GetErrors() []VcenterLcmNotification`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VcenterLcmNotifications) GetErrorsOk() (*[]VcenterLcmNotification, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VcenterLcmNotifications) SetErrors(v []VcenterLcmNotification)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VcenterLcmNotifications) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


