# VcenterNamespacesInstancesStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUsed** | **int64** | Overall CPU usage of the namespace, in MHz. This is the sum of CPU usage across all pods in the Kubernetes namespace. | 
**MemoryUsed** | **int64** | Overall memory usage of the namespace (in mebibytes). This is the sum of memory usage across all pods. | 
**StorageUsed** | **int64** | Overall storage used by the namespace (in mebibytes). This is the sum of storage used by pods across all datastores in the cluster associated with storage policies configured for the namespace. | 

## Methods

### NewVcenterNamespacesInstancesStats

`func NewVcenterNamespacesInstancesStats(cpuUsed int64, memoryUsed int64, storageUsed int64, ) *VcenterNamespacesInstancesStats`

NewVcenterNamespacesInstancesStats instantiates a new VcenterNamespacesInstancesStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespacesInstancesStatsWithDefaults

`func NewVcenterNamespacesInstancesStatsWithDefaults() *VcenterNamespacesInstancesStats`

NewVcenterNamespacesInstancesStatsWithDefaults instantiates a new VcenterNamespacesInstancesStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUsed

`func (o *VcenterNamespacesInstancesStats) GetCpuUsed() int64`

GetCpuUsed returns the CpuUsed field if non-nil, zero value otherwise.

### GetCpuUsedOk

`func (o *VcenterNamespacesInstancesStats) GetCpuUsedOk() (*int64, bool)`

GetCpuUsedOk returns a tuple with the CpuUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsed

`func (o *VcenterNamespacesInstancesStats) SetCpuUsed(v int64)`

SetCpuUsed sets CpuUsed field to given value.


### GetMemoryUsed

`func (o *VcenterNamespacesInstancesStats) GetMemoryUsed() int64`

GetMemoryUsed returns the MemoryUsed field if non-nil, zero value otherwise.

### GetMemoryUsedOk

`func (o *VcenterNamespacesInstancesStats) GetMemoryUsedOk() (*int64, bool)`

GetMemoryUsedOk returns a tuple with the MemoryUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsed

`func (o *VcenterNamespacesInstancesStats) SetMemoryUsed(v int64)`

SetMemoryUsed sets MemoryUsed field to given value.


### GetStorageUsed

`func (o *VcenterNamespacesInstancesStats) GetStorageUsed() int64`

GetStorageUsed returns the StorageUsed field if non-nil, zero value otherwise.

### GetStorageUsedOk

`func (o *VcenterNamespacesInstancesStats) GetStorageUsedOk() (*int64, bool)`

GetStorageUsedOk returns a tuple with the StorageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsed

`func (o *VcenterNamespacesInstancesStats) SetStorageUsed(v int64)`

SetStorageUsed sets StorageUsed field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


