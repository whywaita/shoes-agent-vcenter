# VcenterNamespaceManagementCNSFileConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VsanClusters** | **[]string** | CNSFileConfig.vsan-clusters is a list of clusters to be used for provisioning file volumes. As a prerequisite these clusters must have vSAN and vSAN file services enabled, and must be in the same vCenter as the Supervisor Cluster. Currently this list must have a single entry which is the cluster identifier of the current cluster. This cluster must be a vSAN cluster and must have vSAN File Service enabled. If a cluster in the list is not a vSAN cluster or does not have vSAN File Service enabled, an InvalidArgument error will be thrown from Clusters.enable, Clusters.update and Clusters.set APIs. For disabling file volume support on Supervisor Cluster, an empty list is expected. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | 

## Methods

### NewVcenterNamespaceManagementCNSFileConfig

`func NewVcenterNamespaceManagementCNSFileConfig(vsanClusters []string, ) *VcenterNamespaceManagementCNSFileConfig`

NewVcenterNamespaceManagementCNSFileConfig instantiates a new VcenterNamespaceManagementCNSFileConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementCNSFileConfigWithDefaults

`func NewVcenterNamespaceManagementCNSFileConfigWithDefaults() *VcenterNamespaceManagementCNSFileConfig`

NewVcenterNamespaceManagementCNSFileConfigWithDefaults instantiates a new VcenterNamespaceManagementCNSFileConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVsanClusters

`func (o *VcenterNamespaceManagementCNSFileConfig) GetVsanClusters() []string`

GetVsanClusters returns the VsanClusters field if non-nil, zero value otherwise.

### GetVsanClustersOk

`func (o *VcenterNamespaceManagementCNSFileConfig) GetVsanClustersOk() (*[]string, bool)`

GetVsanClustersOk returns a tuple with the VsanClusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVsanClusters

`func (o *VcenterNamespaceManagementCNSFileConfig) SetVsanClusters(v []string)`

SetVsanClusters sets VsanClusters field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


