# VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | Pointer to **[]string** | Identifies the cluster. cluster If {@term.unset} return all Trust Authority Clusters. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**State** | Pointer to [**[]VcenterTrustedInfrastructureTrustAuthorityClustersState**](VcenterTrustedInfrastructureTrustAuthorityClustersState.md) | The state of the TrustAuthorityClusters. state If {@term.unset} return all Trust Authority Clusters. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) GetCluster() []string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) GetClusterOk() (*[]string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) SetCluster(v []string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetState

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) GetState() []VcenterTrustedInfrastructureTrustAuthorityClustersState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) GetStateOk() (*[]VcenterTrustedInfrastructureTrustAuthorityClustersState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) SetState(v []VcenterTrustedInfrastructureTrustAuthorityClustersState)`

SetState sets State field to given value.

### HasState

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersFilterSpec) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


