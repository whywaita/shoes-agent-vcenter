# VcenterDeploymentTask

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Progress** | Pointer to [**CisTaskProgress**](CisTaskProgress.md) |  | [optional] 
**Result** | Pointer to [**VcenterDeploymentNotifications**](VcenterDeploymentNotifications.md) |  | [optional] 
**Description** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 
**Service** | **string** | Identifier of the service containing the operation. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vapi.service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vapi.service. | 
**Operation** | **string** | Identifier of the operation associated with the task. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vapi.operation. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vapi.operation. | 
**Parent** | Pointer to **string** | Parent of the current task. This field will be unset if the task has no parent. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: cis.task. When operations return a value of this structure as a result, the field will be an identifier for the resource type: cis.task. | [optional] 
**Target** | Pointer to [**VapiStdDynamicID**](VapiStdDynamicID.md) |  | [optional] 
**Status** | [**CisTaskStatus**](CisTaskStatus.md) |  | 
**Cancelable** | **bool** | Flag to indicate whether or not the operation can be cancelled. The value may change as the operation progresses. | 
**Error** | Pointer to **string** | Description of the error if the operation status is \&quot;FAILED\&quot;. If unset the description of why the operation failed will be included in the result of the operation (see Info.result). | [optional] 
**StartTime** | Pointer to **time.Time** | Time when the operation is started. This field is optional and it is only relevant when the value of CommonInfo.status is one of RUNNING, BLOCKED, SUCCEEDED, or FAILED. | [optional] 
**EndTime** | Pointer to **time.Time** | Time when the operation is completed. This field is optional and it is only relevant when the value of CommonInfo.status is one of SUCCEEDED or FAILED. | [optional] 
**User** | Pointer to **string** | Name of the user who performed the operation. This field will be unset if the operation is performed by the system. | [optional] 

## Methods

### NewVcenterDeploymentTask

`func NewVcenterDeploymentTask(description VapiStdLocalizableMessage, service string, operation string, status CisTaskStatus, cancelable bool, ) *VcenterDeploymentTask`

NewVcenterDeploymentTask instantiates a new VcenterDeploymentTask object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentTaskWithDefaults

`func NewVcenterDeploymentTaskWithDefaults() *VcenterDeploymentTask`

NewVcenterDeploymentTaskWithDefaults instantiates a new VcenterDeploymentTask object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProgress

`func (o *VcenterDeploymentTask) GetProgress() CisTaskProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VcenterDeploymentTask) GetProgressOk() (*CisTaskProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VcenterDeploymentTask) SetProgress(v CisTaskProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VcenterDeploymentTask) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetResult

`func (o *VcenterDeploymentTask) GetResult() VcenterDeploymentNotifications`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *VcenterDeploymentTask) GetResultOk() (*VcenterDeploymentNotifications, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *VcenterDeploymentTask) SetResult(v VcenterDeploymentNotifications)`

SetResult sets Result field to given value.

### HasResult

`func (o *VcenterDeploymentTask) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterDeploymentTask) GetDescription() VapiStdLocalizableMessage`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterDeploymentTask) GetDescriptionOk() (*VapiStdLocalizableMessage, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterDeploymentTask) SetDescription(v VapiStdLocalizableMessage)`

SetDescription sets Description field to given value.


### GetService

`func (o *VcenterDeploymentTask) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterDeploymentTask) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterDeploymentTask) SetService(v string)`

SetService sets Service field to given value.


### GetOperation

`func (o *VcenterDeploymentTask) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VcenterDeploymentTask) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VcenterDeploymentTask) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetParent

`func (o *VcenterDeploymentTask) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VcenterDeploymentTask) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VcenterDeploymentTask) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VcenterDeploymentTask) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetTarget

`func (o *VcenterDeploymentTask) GetTarget() VapiStdDynamicID`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VcenterDeploymentTask) GetTargetOk() (*VapiStdDynamicID, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VcenterDeploymentTask) SetTarget(v VapiStdDynamicID)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VcenterDeploymentTask) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStatus

`func (o *VcenterDeploymentTask) GetStatus() CisTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterDeploymentTask) GetStatusOk() (*CisTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterDeploymentTask) SetStatus(v CisTaskStatus)`

SetStatus sets Status field to given value.


### GetCancelable

`func (o *VcenterDeploymentTask) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *VcenterDeploymentTask) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *VcenterDeploymentTask) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.


### GetError

`func (o *VcenterDeploymentTask) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterDeploymentTask) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterDeploymentTask) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VcenterDeploymentTask) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartTime

`func (o *VcenterDeploymentTask) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VcenterDeploymentTask) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VcenterDeploymentTask) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VcenterDeploymentTask) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *VcenterDeploymentTask) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VcenterDeploymentTask) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VcenterDeploymentTask) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *VcenterDeploymentTask) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUser

`func (o *VcenterDeploymentTask) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VcenterDeploymentTask) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VcenterDeploymentTask) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *VcenterDeploymentTask) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


