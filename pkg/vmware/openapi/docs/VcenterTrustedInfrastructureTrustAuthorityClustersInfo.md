# VcenterTrustedInfrastructureTrustAuthorityClustersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifies the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**State** | [**VcenterTrustedInfrastructureTrustAuthorityClustersState**](VcenterTrustedInfrastructureTrustAuthorityClustersState.md) |  | 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersInfo

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersInfo(cluster string, state VcenterTrustedInfrastructureTrustAuthorityClustersState, ) *VcenterTrustedInfrastructureTrustAuthorityClustersInfo`

NewVcenterTrustedInfrastructureTrustAuthorityClustersInfo instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersInfoWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersInfoWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersInfo`

NewVcenterTrustedInfrastructureTrustAuthorityClustersInfoWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetState

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersInfo) GetState() VcenterTrustedInfrastructureTrustAuthorityClustersState`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersInfo) GetStateOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersState, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersInfo) SetState(v VcenterTrustedInfrastructureTrustAuthorityClustersState)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


