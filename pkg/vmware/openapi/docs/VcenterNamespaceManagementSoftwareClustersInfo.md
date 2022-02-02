# VcenterNamespaceManagementSoftwareClustersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentVersion** | **string** | Current version of the cluster. | 
**AvailableVersions** | **[]string** | Set of available versions can be upgraded to. | 
**LastUpgradedDate** | Pointer to **time.Time** | Date of last successful upgrade. If unset, the cluster has not yet been upgraded. | [optional] 
**Messages** | [**[]VcenterNamespaceManagementSoftwareClustersMessage**](VcenterNamespaceManagementSoftwareClustersMessage.md) | Current set of messages associated with the cluster version. | 
**State** | [**VcenterNamespaceManagementSoftwareClustersState**](VcenterNamespaceManagementSoftwareClustersState.md) |  | 
**UpgradeStatus** | Pointer to [**VcenterNamespaceManagementSoftwareClustersUpgradeStatus**](VcenterNamespaceManagementSoftwareClustersUpgradeStatus.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementSoftwareClustersInfo

`func NewVcenterNamespaceManagementSoftwareClustersInfo(currentVersion string, availableVersions []string, messages []VcenterNamespaceManagementSoftwareClustersMessage, state VcenterNamespaceManagementSoftwareClustersState, ) *VcenterNamespaceManagementSoftwareClustersInfo`

NewVcenterNamespaceManagementSoftwareClustersInfo instantiates a new VcenterNamespaceManagementSoftwareClustersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSoftwareClustersInfoWithDefaults

`func NewVcenterNamespaceManagementSoftwareClustersInfoWithDefaults() *VcenterNamespaceManagementSoftwareClustersInfo`

NewVcenterNamespaceManagementSoftwareClustersInfoWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentVersion

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetAvailableVersions

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetAvailableVersions() []string`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetAvailableVersionsOk() (*[]string, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) SetAvailableVersions(v []string)`

SetAvailableVersions sets AvailableVersions field to given value.


### GetLastUpgradedDate

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetLastUpgradedDate() time.Time`

GetLastUpgradedDate returns the LastUpgradedDate field if non-nil, zero value otherwise.

### GetLastUpgradedDateOk

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetLastUpgradedDateOk() (*time.Time, bool)`

GetLastUpgradedDateOk returns a tuple with the LastUpgradedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradedDate

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) SetLastUpgradedDate(v time.Time)`

SetLastUpgradedDate sets LastUpgradedDate field to given value.

### HasLastUpgradedDate

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) HasLastUpgradedDate() bool`

HasLastUpgradedDate returns a boolean if a field has been set.

### GetMessages

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetMessages() []VcenterNamespaceManagementSoftwareClustersMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetMessagesOk() (*[]VcenterNamespaceManagementSoftwareClustersMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) SetMessages(v []VcenterNamespaceManagementSoftwareClustersMessage)`

SetMessages sets Messages field to given value.


### GetState

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetState() VcenterNamespaceManagementSoftwareClustersState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetStateOk() (*VcenterNamespaceManagementSoftwareClustersState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) SetState(v VcenterNamespaceManagementSoftwareClustersState)`

SetState sets State field to given value.


### GetUpgradeStatus

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetUpgradeStatus() VcenterNamespaceManagementSoftwareClustersUpgradeStatus`

GetUpgradeStatus returns the UpgradeStatus field if non-nil, zero value otherwise.

### GetUpgradeStatusOk

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) GetUpgradeStatusOk() (*VcenterNamespaceManagementSoftwareClustersUpgradeStatus, bool)`

GetUpgradeStatusOk returns a tuple with the UpgradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpgradeStatus

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) SetUpgradeStatus(v VcenterNamespaceManagementSoftwareClustersUpgradeStatus)`

SetUpgradeStatus sets UpgradeStatus field to given value.

### HasUpgradeStatus

`func (o *VcenterNamespaceManagementSoftwareClustersInfo) HasUpgradeStatus() bool`

HasUpgradeStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


