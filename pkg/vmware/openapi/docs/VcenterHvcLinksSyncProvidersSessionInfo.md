# VcenterHvcLinksSyncProvidersSessionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Stage** | [**VcenterHvcLinksSyncProvidersSessionInfoStage**](VcenterHvcLinksSyncProvidersSessionInfoStage.md) |  | 
**CompletedWork** | **int64** | Completed work for the session. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**TotalWork** | **int64** | Total work for the session. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**CompletionTime** | Pointer to **time.Time** | Time at which forced sync session was completed. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | [optional] 
**StartTime** | **time.Time** | Time at which force sync was initiated. *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**Exception** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterHvcLinksSyncProvidersSessionInfo

`func NewVcenterHvcLinksSyncProvidersSessionInfo(stage VcenterHvcLinksSyncProvidersSessionInfoStage, completedWork int64, totalWork int64, startTime time.Time, ) *VcenterHvcLinksSyncProvidersSessionInfo`

NewVcenterHvcLinksSyncProvidersSessionInfo instantiates a new VcenterHvcLinksSyncProvidersSessionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHvcLinksSyncProvidersSessionInfoWithDefaults

`func NewVcenterHvcLinksSyncProvidersSessionInfoWithDefaults() *VcenterHvcLinksSyncProvidersSessionInfo`

NewVcenterHvcLinksSyncProvidersSessionInfoWithDefaults instantiates a new VcenterHvcLinksSyncProvidersSessionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStage

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetStage() VcenterHvcLinksSyncProvidersSessionInfoStage`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetStageOk() (*VcenterHvcLinksSyncProvidersSessionInfoStage, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) SetStage(v VcenterHvcLinksSyncProvidersSessionInfoStage)`

SetStage sets Stage field to given value.


### GetCompletedWork

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetCompletedWork() int64`

GetCompletedWork returns the CompletedWork field if non-nil, zero value otherwise.

### GetCompletedWorkOk

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetCompletedWorkOk() (*int64, bool)`

GetCompletedWorkOk returns a tuple with the CompletedWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedWork

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) SetCompletedWork(v int64)`

SetCompletedWork sets CompletedWork field to given value.


### GetTotalWork

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetTotalWork() int64`

GetTotalWork returns the TotalWork field if non-nil, zero value otherwise.

### GetTotalWorkOk

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetTotalWorkOk() (*int64, bool)`

GetTotalWorkOk returns a tuple with the TotalWork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWork

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) SetTotalWork(v int64)`

SetTotalWork sets TotalWork field to given value.


### GetCompletionTime

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetCompletionTime() time.Time`

GetCompletionTime returns the CompletionTime field if non-nil, zero value otherwise.

### GetCompletionTimeOk

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetCompletionTimeOk() (*time.Time, bool)`

GetCompletionTimeOk returns a tuple with the CompletionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTime

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) SetCompletionTime(v time.Time)`

SetCompletionTime sets CompletionTime field to given value.

### HasCompletionTime

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) HasCompletionTime() bool`

HasCompletionTime returns a boolean if a field has been set.

### GetStartTime

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.


### GetException

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetException() VapiStdLocalizableMessage`

GetException returns the Exception field if non-nil, zero value otherwise.

### GetExceptionOk

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) GetExceptionOk() (*VapiStdLocalizableMessage, bool)`

GetExceptionOk returns a tuple with the Exception field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetException

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) SetException(v VapiStdLocalizableMessage)`

SetException sets Exception field to given value.

### HasException

`func (o *VcenterHvcLinksSyncProvidersSessionInfo) HasException() bool`

HasException returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


