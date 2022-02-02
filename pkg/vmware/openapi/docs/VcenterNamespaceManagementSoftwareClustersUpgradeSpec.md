# VcenterNamespaceManagementSoftwareClustersUpgradeSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DesiredVersion** | **string** | Version number the cluster is going to be upgraded to. | 
**IgnorePrecheckWarnings** | Pointer to **bool** | If true, the upgrade workflow will ignore any pre-check warnings and proceed with the upgrade. If unset, the upgrade workflow will not ignore pre-check warnings and fail the upgrade. It is equivalent to setting the value to false. The workflow adopts a conservative approach of failing the upgrade if unset to solely let the user decide whether to force the upgrade despite the warnings. | [optional] 

## Methods

### NewVcenterNamespaceManagementSoftwareClustersUpgradeSpec

`func NewVcenterNamespaceManagementSoftwareClustersUpgradeSpec(desiredVersion string, ) *VcenterNamespaceManagementSoftwareClustersUpgradeSpec`

NewVcenterNamespaceManagementSoftwareClustersUpgradeSpec instantiates a new VcenterNamespaceManagementSoftwareClustersUpgradeSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSoftwareClustersUpgradeSpecWithDefaults

`func NewVcenterNamespaceManagementSoftwareClustersUpgradeSpecWithDefaults() *VcenterNamespaceManagementSoftwareClustersUpgradeSpec`

NewVcenterNamespaceManagementSoftwareClustersUpgradeSpecWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersUpgradeSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.


### GetIgnorePrecheckWarnings

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) GetIgnorePrecheckWarnings() bool`

GetIgnorePrecheckWarnings returns the IgnorePrecheckWarnings field if non-nil, zero value otherwise.

### GetIgnorePrecheckWarningsOk

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) GetIgnorePrecheckWarningsOk() (*bool, bool)`

GetIgnorePrecheckWarningsOk returns a tuple with the IgnorePrecheckWarnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnorePrecheckWarnings

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) SetIgnorePrecheckWarnings(v bool)`

SetIgnorePrecheckWarnings sets IgnorePrecheckWarnings field to given value.

### HasIgnorePrecheckWarnings

`func (o *VcenterNamespaceManagementSoftwareClustersUpgradeSpec) HasIgnorePrecheckWarnings() bool`

HasIgnorePrecheckWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


