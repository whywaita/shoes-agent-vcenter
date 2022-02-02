# VcenterDeploymentInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | [**VcenterDeploymentApplianceState**](VcenterDeploymentApplianceState.md) |  | 
**Progress** | Pointer to [**CisTaskProgress**](CisTaskProgress.md) |  | [optional] 
**SubtaskOrder** | Pointer to **[]string** | The ordered list of subtasks for this deployment operation. Only set when the appliance state is RUNNING_IN_PROGRESS, FAILED, CANCELLED and SUCCEEDED. | [optional] 
**Subtasks** | Pointer to [**[]VcenterDeploymentInfoSubtasks**](VcenterDeploymentInfoSubtasks.md) | The map of the deployment subtasks and their status infomation. Only set when the appliance state is RUNNING_IN_PROGRESS, FAILED, CANCELLED and SUCCEEDED. | [optional] 
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

### NewVcenterDeploymentInfo

`func NewVcenterDeploymentInfo(state VcenterDeploymentApplianceState, description VapiStdLocalizableMessage, service string, operation string, status CisTaskStatus, cancelable bool, ) *VcenterDeploymentInfo`

NewVcenterDeploymentInfo instantiates a new VcenterDeploymentInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentInfoWithDefaults

`func NewVcenterDeploymentInfoWithDefaults() *VcenterDeploymentInfo`

NewVcenterDeploymentInfoWithDefaults instantiates a new VcenterDeploymentInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *VcenterDeploymentInfo) GetState() VcenterDeploymentApplianceState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterDeploymentInfo) GetStateOk() (*VcenterDeploymentApplianceState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterDeploymentInfo) SetState(v VcenterDeploymentApplianceState)`

SetState sets State field to given value.


### GetProgress

`func (o *VcenterDeploymentInfo) GetProgress() CisTaskProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VcenterDeploymentInfo) GetProgressOk() (*CisTaskProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VcenterDeploymentInfo) SetProgress(v CisTaskProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VcenterDeploymentInfo) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetSubtaskOrder

`func (o *VcenterDeploymentInfo) GetSubtaskOrder() []string`

GetSubtaskOrder returns the SubtaskOrder field if non-nil, zero value otherwise.

### GetSubtaskOrderOk

`func (o *VcenterDeploymentInfo) GetSubtaskOrderOk() (*[]string, bool)`

GetSubtaskOrderOk returns a tuple with the SubtaskOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtaskOrder

`func (o *VcenterDeploymentInfo) SetSubtaskOrder(v []string)`

SetSubtaskOrder sets SubtaskOrder field to given value.

### HasSubtaskOrder

`func (o *VcenterDeploymentInfo) HasSubtaskOrder() bool`

HasSubtaskOrder returns a boolean if a field has been set.

### GetSubtasks

`func (o *VcenterDeploymentInfo) GetSubtasks() []VcenterDeploymentInfoSubtasks`

GetSubtasks returns the Subtasks field if non-nil, zero value otherwise.

### GetSubtasksOk

`func (o *VcenterDeploymentInfo) GetSubtasksOk() (*[]VcenterDeploymentInfoSubtasks, bool)`

GetSubtasksOk returns a tuple with the Subtasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtasks

`func (o *VcenterDeploymentInfo) SetSubtasks(v []VcenterDeploymentInfoSubtasks)`

SetSubtasks sets Subtasks field to given value.

### HasSubtasks

`func (o *VcenterDeploymentInfo) HasSubtasks() bool`

HasSubtasks returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterDeploymentInfo) GetDescription() VapiStdLocalizableMessage`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterDeploymentInfo) GetDescriptionOk() (*VapiStdLocalizableMessage, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterDeploymentInfo) SetDescription(v VapiStdLocalizableMessage)`

SetDescription sets Description field to given value.


### GetService

`func (o *VcenterDeploymentInfo) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterDeploymentInfo) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterDeploymentInfo) SetService(v string)`

SetService sets Service field to given value.


### GetOperation

`func (o *VcenterDeploymentInfo) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *VcenterDeploymentInfo) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *VcenterDeploymentInfo) SetOperation(v string)`

SetOperation sets Operation field to given value.


### GetParent

`func (o *VcenterDeploymentInfo) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VcenterDeploymentInfo) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VcenterDeploymentInfo) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *VcenterDeploymentInfo) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetTarget

`func (o *VcenterDeploymentInfo) GetTarget() VapiStdDynamicID`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *VcenterDeploymentInfo) GetTargetOk() (*VapiStdDynamicID, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *VcenterDeploymentInfo) SetTarget(v VapiStdDynamicID)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *VcenterDeploymentInfo) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetStatus

`func (o *VcenterDeploymentInfo) GetStatus() CisTaskStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterDeploymentInfo) GetStatusOk() (*CisTaskStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterDeploymentInfo) SetStatus(v CisTaskStatus)`

SetStatus sets Status field to given value.


### GetCancelable

`func (o *VcenterDeploymentInfo) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *VcenterDeploymentInfo) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *VcenterDeploymentInfo) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.


### GetError

`func (o *VcenterDeploymentInfo) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *VcenterDeploymentInfo) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *VcenterDeploymentInfo) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *VcenterDeploymentInfo) HasError() bool`

HasError returns a boolean if a field has been set.

### GetStartTime

`func (o *VcenterDeploymentInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VcenterDeploymentInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VcenterDeploymentInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *VcenterDeploymentInfo) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetEndTime

`func (o *VcenterDeploymentInfo) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *VcenterDeploymentInfo) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *VcenterDeploymentInfo) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *VcenterDeploymentInfo) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetUser

`func (o *VcenterDeploymentInfo) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VcenterDeploymentInfo) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VcenterDeploymentInfo) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *VcenterDeploymentInfo) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


