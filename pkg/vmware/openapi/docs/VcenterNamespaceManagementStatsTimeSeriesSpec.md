# VcenterNamespaceManagementStatsTimeSeriesSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjType** | [**VcenterNamespaceManagementStatsTimeSeriesSpecObjType**](VcenterNamespaceManagementStatsTimeSeriesSpecObjType.md) |  | 
**Pod** | Pointer to [**VcenterNamespaceManagementStatsTimeSeriesPodIdentifier**](VcenterNamespaceManagementStatsTimeSeriesPodIdentifier.md) |  | [optional] 
**Namespace** | Pointer to **string** | Namespace name for queries for a namespace. This field is optional and it is only relevant when the value of TimeSeries.Spec.obj-type is NAMESPACE. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespaces.Instance. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespaces.Instance. | [optional] 
**Cluster** | Pointer to **string** | Cluster identifier for queries for a cluster. This field is optional and it is only relevant when the value of TimeSeries.Spec.obj-type is CLUSTER. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | [optional] 
**Start** | **int64** | UNIX timestamp value indicating when the requested series of statistical samples should begin. https://en.wikipedia.org/wiki/Unix_time | 
**End** | **int64** | UNIX timestamp value indicating when the requested series of statistical samples should end. https://en.wikipedia.org/wiki/Unix_time | 

## Methods

### NewVcenterNamespaceManagementStatsTimeSeriesSpec

`func NewVcenterNamespaceManagementStatsTimeSeriesSpec(objType VcenterNamespaceManagementStatsTimeSeriesSpecObjType, start int64, end int64, ) *VcenterNamespaceManagementStatsTimeSeriesSpec`

NewVcenterNamespaceManagementStatsTimeSeriesSpec instantiates a new VcenterNamespaceManagementStatsTimeSeriesSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementStatsTimeSeriesSpecWithDefaults

`func NewVcenterNamespaceManagementStatsTimeSeriesSpecWithDefaults() *VcenterNamespaceManagementStatsTimeSeriesSpec`

NewVcenterNamespaceManagementStatsTimeSeriesSpecWithDefaults instantiates a new VcenterNamespaceManagementStatsTimeSeriesSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjType

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetObjType() VcenterNamespaceManagementStatsTimeSeriesSpecObjType`

GetObjType returns the ObjType field if non-nil, zero value otherwise.

### GetObjTypeOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetObjTypeOk() (*VcenterNamespaceManagementStatsTimeSeriesSpecObjType, bool)`

GetObjTypeOk returns a tuple with the ObjType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjType

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) SetObjType(v VcenterNamespaceManagementStatsTimeSeriesSpecObjType)`

SetObjType sets ObjType field to given value.


### GetPod

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetPod() VcenterNamespaceManagementStatsTimeSeriesPodIdentifier`

GetPod returns the Pod field if non-nil, zero value otherwise.

### GetPodOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetPodOk() (*VcenterNamespaceManagementStatsTimeSeriesPodIdentifier, bool)`

GetPodOk returns a tuple with the Pod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPod

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) SetPod(v VcenterNamespaceManagementStatsTimeSeriesPodIdentifier)`

SetPod sets Pod field to given value.

### HasPod

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) HasPod() bool`

HasPod returns a boolean if a field has been set.

### GetNamespace

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetStart

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetStart() int64`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetStartOk() (*int64, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) SetStart(v int64)`

SetStart sets Start field to given value.


### GetEnd

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetEnd() int64`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) GetEndOk() (*int64, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *VcenterNamespaceManagementStatsTimeSeriesSpec) SetEnd(v int64)`

SetEnd sets End field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


