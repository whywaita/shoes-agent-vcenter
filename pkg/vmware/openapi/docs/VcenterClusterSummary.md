# VcenterClusterSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cluster** | **string** | Identifier of the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource. | 
**Name** | **string** | Name of the cluster. | 
**HaEnabled** | **bool** | Flag indicating whether the vSphere HA feature is enabled for the cluster. | 
**DrsEnabled** | **bool** | Flag indicating whether the vSphere DRS service is enabled for the cluster. | 

## Methods

### NewVcenterClusterSummary

`func NewVcenterClusterSummary(cluster string, name string, haEnabled bool, drsEnabled bool, ) *VcenterClusterSummary`

NewVcenterClusterSummary instantiates a new VcenterClusterSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterClusterSummaryWithDefaults

`func NewVcenterClusterSummaryWithDefaults() *VcenterClusterSummary`

NewVcenterClusterSummaryWithDefaults instantiates a new VcenterClusterSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCluster

`func (o *VcenterClusterSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterClusterSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterClusterSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetName

`func (o *VcenterClusterSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterClusterSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterClusterSummary) SetName(v string)`

SetName sets Name field to given value.


### GetHaEnabled

`func (o *VcenterClusterSummary) GetHaEnabled() bool`

GetHaEnabled returns the HaEnabled field if non-nil, zero value otherwise.

### GetHaEnabledOk

`func (o *VcenterClusterSummary) GetHaEnabledOk() (*bool, bool)`

GetHaEnabledOk returns a tuple with the HaEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHaEnabled

`func (o *VcenterClusterSummary) SetHaEnabled(v bool)`

SetHaEnabled sets HaEnabled field to given value.


### GetDrsEnabled

`func (o *VcenterClusterSummary) GetDrsEnabled() bool`

GetDrsEnabled returns the DrsEnabled field if non-nil, zero value otherwise.

### GetDrsEnabledOk

`func (o *VcenterClusterSummary) GetDrsEnabledOk() (*bool, bool)`

GetDrsEnabledOk returns a tuple with the DrsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrsEnabled

`func (o *VcenterClusterSummary) SetDrsEnabled(v bool)`

SetDrsEnabled sets DrsEnabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


