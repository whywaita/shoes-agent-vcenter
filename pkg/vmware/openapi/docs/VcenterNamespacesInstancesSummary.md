# VcenterNamespacesInstancesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** | Identifier of the namespace. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.Instance. | 
**Description** | **string** | Description of the namespace. | 
**Cluster** | **string** | Identifier for the cluster hosting the namespace. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**ConfigStatus** | [**VcenterNamespacesInstancesConfigStatus**](VcenterNamespacesInstancesConfigStatus.md) |  | 
**Stats** | [**VcenterNamespacesInstancesStats**](VcenterNamespacesInstancesStats.md) |  | 
**SelfServiceNamespace** | Pointer to **bool** | Flag to indicate the self service namespace. If unset, the namespace is not marked as self service namespace. | [optional] 

## Methods

### NewVcenterNamespacesInstancesSummary

`func NewVcenterNamespacesInstancesSummary(namespace string, description string, cluster string, configStatus VcenterNamespacesInstancesConfigStatus, stats VcenterNamespacesInstancesStats, ) *VcenterNamespacesInstancesSummary`

NewVcenterNamespacesInstancesSummary instantiates a new VcenterNamespacesInstancesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesSummaryWithDefaults

`func NewVcenterNamespacesInstancesSummaryWithDefaults() *VcenterNamespacesInstancesSummary`

NewVcenterNamespacesInstancesSummaryWithDefaults instantiates a new VcenterNamespacesInstancesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *VcenterNamespacesInstancesSummary) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VcenterNamespacesInstancesSummary) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VcenterNamespacesInstancesSummary) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetDescription

`func (o *VcenterNamespacesInstancesSummary) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterNamespacesInstancesSummary) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterNamespacesInstancesSummary) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetCluster

`func (o *VcenterNamespacesInstancesSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespacesInstancesSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespacesInstancesSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetConfigStatus

`func (o *VcenterNamespacesInstancesSummary) GetConfigStatus() VcenterNamespacesInstancesConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespacesInstancesSummary) GetConfigStatusOk() (*VcenterNamespacesInstancesConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespacesInstancesSummary) SetConfigStatus(v VcenterNamespacesInstancesConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetStats

`func (o *VcenterNamespacesInstancesSummary) GetStats() VcenterNamespacesInstancesStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *VcenterNamespacesInstancesSummary) GetStatsOk() (*VcenterNamespacesInstancesStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *VcenterNamespacesInstancesSummary) SetStats(v VcenterNamespacesInstancesStats)`

SetStats sets Stats field to given value.


### GetSelfServiceNamespace

`func (o *VcenterNamespacesInstancesSummary) GetSelfServiceNamespace() bool`

GetSelfServiceNamespace returns the SelfServiceNamespace field if non-nil, zero value otherwise.

### GetSelfServiceNamespaceOk

`func (o *VcenterNamespacesInstancesSummary) GetSelfServiceNamespaceOk() (*bool, bool)`

GetSelfServiceNamespaceOk returns a tuple with the SelfServiceNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfServiceNamespace

`func (o *VcenterNamespacesInstancesSummary) SetSelfServiceNamespace(v bool)`

SetSelfServiceNamespace sets SelfServiceNamespace field to given value.

### HasSelfServiceNamespace

`func (o *VcenterNamespacesInstancesSummary) HasSelfServiceNamespace() bool`

HasSelfServiceNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


