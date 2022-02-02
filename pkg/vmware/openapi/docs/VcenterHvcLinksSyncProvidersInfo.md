# VcenterHvcLinksSyncProvidersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastSyncTime** | Pointer to **time.Time** | Last sync time for the provider. This indicates the last time that either a background sync or a force sync was started for the provider *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | [optional] 
**Status** | [**VcenterHvcLinksSyncProvidersStatus**](VcenterHvcLinksSyncProvidersStatus.md) |  | 
**PollingIntervalInSeconds** | **int64** | Sync Polling interval between local and remote replicas for the provider *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments. | 
**CurrentSessionInfo** | Pointer to [**VcenterHvcLinksSyncProvidersSessionInfo**](VcenterHvcLinksSyncProvidersSessionInfo.md) |  | [optional] 
**StatusMessage** | Pointer to [**VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) |  | [optional] 

## Methods

### NewVcenterHvcLinksSyncProvidersInfo

`func NewVcenterHvcLinksSyncProvidersInfo(status VcenterHvcLinksSyncProvidersStatus, pollingIntervalInSeconds int64, ) *VcenterHvcLinksSyncProvidersInfo`

NewVcenterHvcLinksSyncProvidersInfo instantiates a new VcenterHvcLinksSyncProvidersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterHvcLinksSyncProvidersInfoWithDefaults

`func NewVcenterHvcLinksSyncProvidersInfoWithDefaults() *VcenterHvcLinksSyncProvidersInfo`

NewVcenterHvcLinksSyncProvidersInfoWithDefaults instantiates a new VcenterHvcLinksSyncProvidersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastSyncTime

`func (o *VcenterHvcLinksSyncProvidersInfo) GetLastSyncTime() time.Time`

GetLastSyncTime returns the LastSyncTime field if non-nil, zero value otherwise.

### GetLastSyncTimeOk

`func (o *VcenterHvcLinksSyncProvidersInfo) GetLastSyncTimeOk() (*time.Time, bool)`

GetLastSyncTimeOk returns a tuple with the LastSyncTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncTime

`func (o *VcenterHvcLinksSyncProvidersInfo) SetLastSyncTime(v time.Time)`

SetLastSyncTime sets LastSyncTime field to given value.

### HasLastSyncTime

`func (o *VcenterHvcLinksSyncProvidersInfo) HasLastSyncTime() bool`

HasLastSyncTime returns a boolean if a field has been set.

### GetStatus

`func (o *VcenterHvcLinksSyncProvidersInfo) GetStatus() VcenterHvcLinksSyncProvidersStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VcenterHvcLinksSyncProvidersInfo) GetStatusOk() (*VcenterHvcLinksSyncProvidersStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VcenterHvcLinksSyncProvidersInfo) SetStatus(v VcenterHvcLinksSyncProvidersStatus)`

SetStatus sets Status field to given value.


### GetPollingIntervalInSeconds

`func (o *VcenterHvcLinksSyncProvidersInfo) GetPollingIntervalInSeconds() int64`

GetPollingIntervalInSeconds returns the PollingIntervalInSeconds field if non-nil, zero value otherwise.

### GetPollingIntervalInSecondsOk

`func (o *VcenterHvcLinksSyncProvidersInfo) GetPollingIntervalInSecondsOk() (*int64, bool)`

GetPollingIntervalInSecondsOk returns a tuple with the PollingIntervalInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPollingIntervalInSeconds

`func (o *VcenterHvcLinksSyncProvidersInfo) SetPollingIntervalInSeconds(v int64)`

SetPollingIntervalInSeconds sets PollingIntervalInSeconds field to given value.


### GetCurrentSessionInfo

`func (o *VcenterHvcLinksSyncProvidersInfo) GetCurrentSessionInfo() VcenterHvcLinksSyncProvidersSessionInfo`

GetCurrentSessionInfo returns the CurrentSessionInfo field if non-nil, zero value otherwise.

### GetCurrentSessionInfoOk

`func (o *VcenterHvcLinksSyncProvidersInfo) GetCurrentSessionInfoOk() (*VcenterHvcLinksSyncProvidersSessionInfo, bool)`

GetCurrentSessionInfoOk returns a tuple with the CurrentSessionInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentSessionInfo

`func (o *VcenterHvcLinksSyncProvidersInfo) SetCurrentSessionInfo(v VcenterHvcLinksSyncProvidersSessionInfo)`

SetCurrentSessionInfo sets CurrentSessionInfo field to given value.

### HasCurrentSessionInfo

`func (o *VcenterHvcLinksSyncProvidersInfo) HasCurrentSessionInfo() bool`

HasCurrentSessionInfo returns a boolean if a field has been set.

### GetStatusMessage

`func (o *VcenterHvcLinksSyncProvidersInfo) GetStatusMessage() VapiStdLocalizableMessage`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *VcenterHvcLinksSyncProvidersInfo) GetStatusMessageOk() (*VapiStdLocalizableMessage, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *VcenterHvcLinksSyncProvidersInfo) SetStatusMessage(v VapiStdLocalizableMessage)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *VcenterHvcLinksSyncProvidersInfo) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


