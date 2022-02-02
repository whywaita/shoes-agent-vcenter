# CisTaskProgress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | **int64** | Total amount of the work for the operation. | 
**Completed** | **int64** | The amount of work completed for the operation. The value can only be incremented. | 
**Message** | [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | 

## Methods

### NewCisTaskProgress

`func NewCisTaskProgress(total int64, completed int64, message VapiStdLocalizableMessage, ) *CisTaskProgress`

NewCisTaskProgress instantiates a new CisTaskProgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCisTaskProgressWithDefaults

`func NewCisTaskProgressWithDefaults() *CisTaskProgress`

NewCisTaskProgressWithDefaults instantiates a new CisTaskProgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *CisTaskProgress) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CisTaskProgress) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CisTaskProgress) SetTotal(v int64)`

SetTotal sets Total field to given value.


### GetCompleted

`func (o *CisTaskProgress) GetCompleted() int64`

GetCompleted returns the Completed field if non-nil, zero value otherwise.

### GetCompletedOk

`func (o *CisTaskProgress) GetCompletedOk() (*int64, bool)`

GetCompletedOk returns a tuple with the Completed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleted

`func (o *CisTaskProgress) SetCompleted(v int64)`

SetCompleted sets Completed field to given value.


### GetMessage

`func (o *CisTaskProgress) GetMessage() VapiStdLocalizableMessage`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CisTaskProgress) GetMessageOk() (*VapiStdLocalizableMessage, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CisTaskProgress) SetMessage(v VapiStdLocalizableMessage)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


