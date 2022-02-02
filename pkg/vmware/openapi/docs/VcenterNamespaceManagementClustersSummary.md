# VcenterNamespaceManagementClustersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier for the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**ClusterName** | **string** | Name of the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.name. | 
**Stats** | [**VcenterNamespaceManagementClustersStats**](VcenterNamespaceManagementClustersStats.md) |  | 
**ConfigStatus** | [**VcenterNamespaceManagementClustersConfigStatus**](VcenterNamespaceManagementClustersConfigStatus.md) |  | 
**KubernetesStatus** | [**VcenterNamespaceManagementClustersKubernetesStatus**](VcenterNamespaceManagementClustersKubernetesStatus.md) |  | 

## Methods

### NewVcenterNamespaceManagementClustersSummary

`func NewVcenterNamespaceManagementClustersSummary(cluster string, clusterName string, stats VcenterNamespaceManagementClustersStats, configStatus VcenterNamespaceManagementClustersConfigStatus, kubernetesStatus VcenterNamespaceManagementClustersKubernetesStatus, ) *VcenterNamespaceManagementClustersSummary`

NewVcenterNamespaceManagementClustersSummary instantiates a new VcenterNamespaceManagementClustersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersSummaryWithDefaults

`func NewVcenterNamespaceManagementClustersSummaryWithDefaults() *VcenterNamespaceManagementClustersSummary`

NewVcenterNamespaceManagementClustersSummaryWithDefaults instantiates a new VcenterNamespaceManagementClustersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterNamespaceManagementClustersSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespaceManagementClustersSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespaceManagementClustersSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetClusterName

`func (o *VcenterNamespaceManagementClustersSummary) GetClusterName() string`

GetClusterName returns the ClusterName field if non-nil, zero value otherwise.

### GetClusterNameOk

`func (o *VcenterNamespaceManagementClustersSummary) GetClusterNameOk() (*string, bool)`

GetClusterNameOk returns a tuple with the ClusterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterName

`func (o *VcenterNamespaceManagementClustersSummary) SetClusterName(v string)`

SetClusterName sets ClusterName field to given value.


### GetStats

`func (o *VcenterNamespaceManagementClustersSummary) GetStats() VcenterNamespaceManagementClustersStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *VcenterNamespaceManagementClustersSummary) GetStatsOk() (*VcenterNamespaceManagementClustersStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *VcenterNamespaceManagementClustersSummary) SetStats(v VcenterNamespaceManagementClustersStats)`

SetStats sets Stats field to given value.


### GetConfigStatus

`func (o *VcenterNamespaceManagementClustersSummary) GetConfigStatus() VcenterNamespaceManagementClustersConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespaceManagementClustersSummary) GetConfigStatusOk() (*VcenterNamespaceManagementClustersConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespaceManagementClustersSummary) SetConfigStatus(v VcenterNamespaceManagementClustersConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetKubernetesStatus

`func (o *VcenterNamespaceManagementClustersSummary) GetKubernetesStatus() VcenterNamespaceManagementClustersKubernetesStatus`

GetKubernetesStatus returns the KubernetesStatus field if non-nil, zero value otherwise.

### GetKubernetesStatusOk

`func (o *VcenterNamespaceManagementClustersSummary) GetKubernetesStatusOk() (*VcenterNamespaceManagementClustersKubernetesStatus, bool)`

GetKubernetesStatusOk returns a tuple with the KubernetesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesStatus

`func (o *VcenterNamespaceManagementClustersSummary) SetKubernetesStatus(v VcenterNamespaceManagementClustersKubernetesStatus)`

SetKubernetesStatus sets KubernetesStatus field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


