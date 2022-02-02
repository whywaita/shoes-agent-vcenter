# VcenterNamespaceManagementStatsTimeSeriesTimeSeries

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Counter** | **string** | Counter identifier. | 
**TimeStamps** | **[]int64** | Sequence of UNIX timestamp values at which statistical values were sampled. https://en.wikipedia.org/wiki/Unix_time | 
**Values** | **[]int64** | Sequence of sampled values corresponding to the timestamps in tss. | 

## Methods

### NewVcenterNamespaceManagementStatsTimeSeriesTimeSeries

`func NewVcenterNamespaceManagementStatsTimeSeriesTimeSeries(counter string, timeStamps []int64, values []int64, ) *VcenterNamespaceManagementStatsTimeSeriesTimeSeries`

NewVcenterNamespaceManagementStatsTimeSeriesTimeSeries instantiates a new VcenterNamespaceManagementStatsTimeSeriesTimeSeries object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementStatsTimeSeriesTimeSeriesWithDefaults

`func NewVcenterNamespaceManagementStatsTimeSeriesTimeSeriesWithDefaults() *VcenterNamespaceManagementStatsTimeSeriesTimeSeries`

NewVcenterNamespaceManagementStatsTimeSeriesTimeSeriesWithDefaults instantiates a new VcenterNamespaceManagementStatsTimeSeriesTimeSeries object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCounter

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetCounter() string`

GetCounter returns the Counter field if non-nil, zero value otherwise.

### GetCounterOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetCounterOk() (*string, bool)`

GetCounterOk returns a tuple with the Counter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounter

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) SetCounter(v string)`

SetCounter sets Counter field to given value.


### GetTimeStamps

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetTimeStamps() []int64`

GetTimeStamps returns the TimeStamps field if non-nil, zero value otherwise.

### GetTimeStampsOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetTimeStampsOk() (*[]int64, bool)`

GetTimeStampsOk returns a tuple with the TimeStamps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeStamps

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) SetTimeStamps(v []int64)`

SetTimeStamps sets TimeStamps field to given value.


### GetValues

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetValues() []int64`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetValuesOk() (*[]int64, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) SetValues(v []int64)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


