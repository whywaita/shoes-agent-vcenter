# VcenterVchaClusterFailoverTaskVchaCluster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Planned** | Pointer to **bool** | If false, a failover is initiated immediately and may result in data loss.  If true, a failover is initated after the Active node flushes its state to Passive and there is no data loss. | [optional] 

## Methods

### NewVcenterVchaClusterFailoverTaskVchaCluster

`func NewVcenterVchaClusterFailoverTaskVchaCluster() *VcenterVchaClusterFailoverTaskVchaCluster`

NewVcenterVchaClusterFailoverTaskVchaCluster instantiates a new VcenterVchaClusterFailoverTaskVchaCluster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterVchaClusterFailoverTaskVchaClusterWithDefaults

`func NewVcenterVchaClusterFailoverTaskVchaClusterWithDefaults() *VcenterVchaClusterFailoverTaskVchaCluster`

NewVcenterVchaClusterFailoverTaskVchaClusterWithDefaults instantiates a new VcenterVchaClusterFailoverTaskVchaCluster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlanned

`func (o *VcenterVchaClusterFailoverTaskVchaCluster) GetPlanned() bool`

GetPlanned returns the Planned field if non-nil, zero value otherwise.

### GetPlannedOk

`func (o *VcenterVchaClusterFailoverTaskVchaCluster) GetPlannedOk() (*bool, bool)`

GetPlannedOk returns a tuple with the Planned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanned

`func (o *VcenterVchaClusterFailoverTaskVchaCluster) SetPlanned(v bool)`

SetPlanned sets Planned field to given value.

### HasPlanned

`func (o *VcenterVchaClusterFailoverTaskVchaCluster) HasPlanned() bool`

HasPlanned returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


