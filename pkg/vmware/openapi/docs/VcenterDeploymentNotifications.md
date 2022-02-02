# VcenterDeploymentNotifications

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Info** | Pointer to [**[]VcenterDeploymentNotification**](VcenterDeploymentNotification.md) | Info notification messages reported. Only set if an info was reported by the task. | [optional] 
**Warnings** | Pointer to [**[]VcenterDeploymentNotification**](VcenterDeploymentNotification.md) | Warning notification messages reported. Only set if an warning was reported by the task. | [optional] 
**Errors** | Pointer to [**[]VcenterDeploymentNotification**](VcenterDeploymentNotification.md) | Error notification messages reported. Only set if an error was reported by the task. | [optional] 

## Methods

### NewVcenterDeploymentNotifications

`func NewVcenterDeploymentNotifications() *VcenterDeploymentNotifications`

NewVcenterDeploymentNotifications instantiates a new VcenterDeploymentNotifications object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentNotificationsWithDefaults

`func NewVcenterDeploymentNotificationsWithDefaults() *VcenterDeploymentNotifications`

NewVcenterDeploymentNotificationsWithDefaults instantiates a new VcenterDeploymentNotifications object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInfo

`func (o *VcenterDeploymentNotifications) GetInfo() []VcenterDeploymentNotification`

GetInfo returns the Info field if non-nil, zero value otherwise.

### GetInfoOk

`func (o *VcenterDeploymentNotifications) GetInfoOk() (*[]VcenterDeploymentNotification, bool)`

GetInfoOk returns a tuple with the Info field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInfo

`func (o *VcenterDeploymentNotifications) SetInfo(v []VcenterDeploymentNotification)`

SetInfo sets Info field to given value.

### HasInfo

`func (o *VcenterDeploymentNotifications) HasInfo() bool`

HasInfo returns a boolean if a field has been set.

### GetWarnings

`func (o *VcenterDeploymentNotifications) GetWarnings() []VcenterDeploymentNotification`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VcenterDeploymentNotifications) GetWarningsOk() (*[]VcenterDeploymentNotification, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VcenterDeploymentNotifications) SetWarnings(v []VcenterDeploymentNotification)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VcenterDeploymentNotifications) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *VcenterDeploymentNotifications) GetErrors() []VcenterDeploymentNotification`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *VcenterDeploymentNotifications) GetErrorsOk() (*[]VcenterDeploymentNotification, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *VcenterDeploymentNotifications) SetErrors(v []VcenterDeploymentNotification)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *VcenterDeploymentNotifications) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


