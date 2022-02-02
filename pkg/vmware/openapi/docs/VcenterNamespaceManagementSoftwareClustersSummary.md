# VcenterNamespaceManagementSoftwareClustersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier for the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**ClusterName** | **string** | Name of the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.name. | 
**CurrentVersion** | **string** | Current version of the cluster. | 
**AvailableVersions** | **[]string** | Set of versions available for upgrade. | 
**LastUpgradedDate** | Pointer to **time.Time** | Date of last successful upgrade. If unset, the cluster has not yet been upgraded. | [optional] 
**DesiredVersion** | Pointer to **string** | Desired version the cluster will be upgraded to. If unset, the cluster upgrade is not in progress. | [optional] 
**State** | [**VcenterNamespaceManagementSoftwareClustersState**](VcenterNamespaceManagementSoftwareClustersState.md) |  | 

## Methods

### NewVcenterNamespaceManagementSoftwareClustersSummary

`func NewVcenterNamespaceManagementSoftwareClustersSummary(cluster string, clusterName string, currentVersion string, availableVersions []string, state VcenterNamespaceManagementSoftwareClustersState, ) *VcenterNamespaceManagementSoftwareClustersSummary`

NewVcenterNamespaceManagementSoftwareClustersSummary instantiates a new VcenterNamespaceManagementSoftwareClustersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementSoftwareClustersSummaryWithDefaults

`func NewVcenterNamespaceManagementSoftwareClustersSummaryWithDefaults() *VcenterNamespaceManagementSoftwareClustersSummary`

NewVcenterNamespaceManagementSoftwareClustersSummaryWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetClusterName

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetCurrentVersion

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.


### GetAvailableVersions

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetAvailableVersions() []string`

GetAvailableVersions returns the AvailableVersions field if non-nil, zero value otherwise.

### GetAvailableVersionsOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetAvailableVersionsOk() (*[]string, bool)`

GetAvailableVersionsOk returns a tuple with the AvailableVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableVersions

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetAvailableVersions(v []string)`

SetAvailableVersions sets AvailableVersions field to given value.


### GetLastUpgradedDate

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetLastUpgradedDate() time.Time`

GetLastUpgradedDate returns the LastUpgradedDate field if non-nil, zero value otherwise.

### GetLastUpgradedDateOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetLastUpgradedDateOk() (*time.Time, bool)`

GetLastUpgradedDateOk returns a tuple with the LastUpgradedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpgradedDate

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetLastUpgradedDate(v time.Time)`

SetLastUpgradedDate sets LastUpgradedDate field to given value.

### HasLastUpgradedDate

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) HasLastUpgradedDate() bool`

HasLastUpgradedDate returns a boolean if a field has been set.

### GetDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetDesiredVersion() string`

GetDesiredVersion returns the DesiredVersion field if non-nil, zero value otherwise.

### GetDesiredVersionOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetDesiredVersionOk() (*string, bool)`

GetDesiredVersionOk returns a tuple with the DesiredVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetDesiredVersion(v string)`

SetDesiredVersion sets DesiredVersion field to given value.

### HasDesiredVersion

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) HasDesiredVersion() bool`

HasDesiredVersion returns a boolean if a field has been set.

### GetState

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetState() VcenterNamespaceManagementSoftwareClustersState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) GetStateOk() (*VcenterNamespaceManagementSoftwareClustersState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterNamespaceManagementSoftwareClustersSummary) SetState(v VcenterNamespaceManagementSoftwareClustersState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


