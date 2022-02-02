# VcenterDeploymentCheckInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**VcenterDeploymentCheckStatus**](VcenterDeploymentCheckStatus.md) |  | 
**Result** | Pointer to [**VcenterDeploymentNotifications**](VcenterDeploymentNotifications.md) |  | [optional] 
**SourceInfo** | Pointer to [**VcenterDeploymentSourceInfo**](VcenterDeploymentSourceInfo.md) |  | [optional] 

## Methods

### NewVcenterDeploymentCheckInfo

`func NewVcenterDeploymentCheckInfo(status VcenterDeploymentCheckStatus, ) *VcenterDeploymentCheckInfo`

NewVcenterDeploymentCheckInfo instantiates a new VcenterDeploymentCheckInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterDeploymentCheckInfoWithDefaults

`func NewVcenterDeploymentCheckInfoWithDefaults() *VcenterDeploymentCheckInfo`

NewVcenterDeploymentCheckInfoWithDefaults instantiates a new VcenterDeploymentCheckInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *VcenterDeploymentCheckInfo) GetStatus() VcenterDeploymentCheckStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterDeploymentCheckInfo) GetStatusOk() (*VcenterDeploymentCheckStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterDeploymentCheckInfo) SetStatus(v VcenterDeploymentCheckStatus)`

SetStatus sets Status field to given value.


### GetResult

`func (o *VcenterDeploymentCheckInfo) GetResult() VcenterDeploymentNotifications`

GetResult returns the Result field if non-nil, zero value otherwise.

### GetResultOk

`func (o *VcenterDeploymentCheckInfo) GetResultOk() (*VcenterDeploymentNotifications, bool)`

GetResultOk returns a tuple with the Result field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResult

`func (o *VcenterDeploymentCheckInfo) SetResult(v VcenterDeploymentNotifications)`

SetResult sets Result field to given value.

### HasResult

`func (o *VcenterDeploymentCheckInfo) HasResult() bool`

HasResult returns a boolean if a field has been set.

### GetSourceInfo

`func (o *VcenterDeploymentCheckInfo) GetSourceInfo() VcenterDeploymentSourceInfo`

GetSourceInfo returns the SourceInfo field if non-nil, zero value otherwise.

### GetSourceInfoOk

`func (o *VcenterDeploymentCheckInfo) GetSourceInfoOk() (*VcenterDeploymentSourceInfo, bool)`

GetSourceInfoOk returns a tuple with the SourceInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceInfo

`func (o *VcenterDeploymentCheckInfo) SetSourceInfo(v VcenterDeploymentSourceInfo)`

SetSourceInfo sets SourceInfo field to given value.

### HasSourceInfo

`func (o *VcenterDeploymentCheckInfo) HasSourceInfo() bool`

HasSourceInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


