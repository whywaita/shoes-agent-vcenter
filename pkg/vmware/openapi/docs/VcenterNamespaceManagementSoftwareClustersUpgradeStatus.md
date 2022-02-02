# VcenterNamespaceManagementSoftwareClustersUpgradeStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredVersion** | Pointer to **string** | Desired version the cluster will be upgraded to. If unset, the cluster upgrade is not in progress. | [optional] 
**Messages** | [**[]VcenterNamespaceManagementSoftwareClustersMessage**](VcenterNamespaceManagementSoftwareClustersMessage.md) | Current set of messages associated with the upgrade state. | 
**Progress** | Pointer to [**VcenterNamespaceManagementSoftwareClustersUpgradeProgress**](VcenterNamespaceManagementSoftwareClustersUpgradeProgress.md) |  | [optional] 

## Methods

### NewVcenterNamespaceManagementSoftwareClustersUpgradeStatus

`func NewVcenterNamespaceManagementSoftwareClustersUpgradeStatus(messages []VcenterNamespaceManagementSoftwareClustersMessage, ) *VcenterNamespaceManagementSoftwareClustersUpgradeStatus`

NewVcenterNamespaceManagementSoftwareClustersUpgradeStatus instantiates a new VcenterNamespaceManagementSoftwareClustersUpgradeStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSoftwareClustersUpgradeStatusWithDefaults

`func NewVcenterNamespaceManagementSoftwareClustersUpgradeStatusWithDefaults() *VcenterNamespaceManagementSoftwareClustersUpgradeStatus`

NewVcenterNamespaceManagementSoftwareClustersUpgradeStatusWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersUpgradeStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetMessages

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) GetMessages() []VcenterNamespaceManagementSoftwareClustersMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) GetMessagesOk() (*[]VcenterNamespaceManagementSoftwareClustersMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) SetMessages(v []VcenterNamespaceManagementSoftwareClustersMessage)`

SetMessages sets Messages field to given value.


### GetProgress

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) GetProgress() VcenterNamespaceManagementSoftwareClustersUpgradeProgress`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) GetProgressOk() (*VcenterNamespaceManagementSoftwareClustersUpgradeProgress, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) SetProgress(v VcenterNamespaceManagementSoftwareClustersUpgradeProgress)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeStatus) HasProgress() bool`

HasProgress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


