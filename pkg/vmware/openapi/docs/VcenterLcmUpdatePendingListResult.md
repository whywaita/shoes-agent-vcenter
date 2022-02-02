# VcenterLcmUpdatePendingListResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastCheckTime** | **time.Time** | Time when the software depo was last checked. | 
**UpdateCount** | Pointer to **int64** | Number of pending updates Only set if there are available updates | [optional] 
**UpgradeCount** | Pointer to **int64** | Number of pending upgrades Only set if there are available upgrades | [optional] 
**Updates** | [**[]VcenterLcmUpdatePendingSummary**](VcenterLcmUpdatePendingSummary.md) | List of pending update details | 
**Issues** | Pointer to [**VcenterLcmNotifications**](VcenterLcmNotifications.md) |  | [optional] 

## Methods

### NewVcenterLcmUpdatePendingListResult

`func NewVcenterLcmUpdatePendingListResult(lastCheckTime time.Time, updates []VcenterLcmUpdatePendingSummary, ) *VcenterLcmUpdatePendingListResult`

NewVcenterLcmUpdatePendingListResult instantiates a new VcenterLcmUpdatePendingListResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterLcmUpdatePendingListResultWithDefaults

`func NewVcenterLcmUpdatePendingListResultWithDefaults() *VcenterLcmUpdatePendingListResult`

NewVcenterLcmUpdatePendingListResultWithDefaults instantiates a new VcenterLcmUpdatePendingListResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastCheckTime

`func (o *VcenterLcmUpdatePendingListResult) GetLastCheckTime() time.Time`

GetLastCheckTime returns the LastCheckTime field if non-nil, zero value otherwise.

### GetLastCheckTimeOk

`func (o *VcenterLcmUpdatePendingListResult) GetLastCheckTimeOk() (*time.Time, bool)`

GetLastCheckTimeOk returns a tuple with the LastCheckTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckTime

`func (o *VcenterLcmUpdatePendingListResult) SetLastCheckTime(v time.Time)`

SetLastCheckTime sets LastCheckTime field to given value.


### GetUpdateCount

`func (o *VcenterLcmUpdatePendingListResult) GetUpdateCount() int64`

GetUpdateCount returns the UpdateCount field if non-nil, zero value otherwise.

### GetUpdateCountOk

`func (o *VcenterLcmUpdatePendingListResult) GetUpdateCountOk() (*int64, bool)`

GetUpdateCountOk returns a tuple with the UpdateCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateCount

`func (o *VcenterLcmUpdatePendingListResult) SetUpdateCount(v int64)`

SetUpdateCount sets UpdateCount field to given value.

### HasUpdateCount

`func (o *VcenterLcmUpdatePendingListResult) HasUpdateCount() bool`

HasUpdateCount returns a boolean if a field has been set.

### GetUpgradeCount

`func (o *VcenterLcmUpdatePendingListResult) GetUpgradeCount() int64`

GetUpgradeCount returns the UpgradeCount field if non-nil, zero value otherwise.

### GetUpgradeCountOk

`func (o *VcenterLcmUpdatePendingListResult) GetUpgradeCountOk() (*int64, bool)`

GetUpgradeCountOk returns a tuple with the UpgradeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeCount

`func (o *VcenterLcmUpdatePendingListResult) SetUpgradeCount(v int64)`

SetUpgradeCount sets UpgradeCount field to given value.

### HasUpgradeCount

`func (o *VcenterLcmUpdatePendingListResult) HasUpgradeCount() bool`

HasUpgradeCount returns a boolean if a field has been set.

### GetUpdates

`func (o *VcenterLcmUpdatePendingListResult) GetUpdates() []VcenterLcmUpdatePendingSummary`

GetUpdates returns the Updates field if non-nil, zero value otherwise.

### GetUpdatesOk

`func (o *VcenterLcmUpdatePendingListResult) GetUpdatesOk() (*[]VcenterLcmUpdatePendingSummary, bool)`

GetUpdatesOk returns a tuple with the Updates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdates

`func (o *VcenterLcmUpdatePendingListResult) SetUpdates(v []VcenterLcmUpdatePendingSummary)`

SetUpdates sets Updates field to given value.


### GetIssues

`func (o *VcenterLcmUpdatePendingListResult) GetIssues() VcenterLcmNotifications`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *VcenterLcmUpdatePendingListResult) GetIssuesOk() (*VcenterLcmNotifications, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *VcenterLcmUpdatePendingListResult) SetIssues(v VcenterLcmNotifications)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *VcenterLcmUpdatePendingListResult) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


