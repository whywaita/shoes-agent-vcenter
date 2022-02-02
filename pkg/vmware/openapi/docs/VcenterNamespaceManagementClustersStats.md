# VcenterNamespaceManagementClustersStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuUsed** | **int64** | Overall CPU usage of the cluster, in MHz. This is the sum of CPU usage across all worker nodes in the cluster. | 
**CpuCapacity** | **int64** | Total CPU capacity in the cluster available for vSphere Namespaces, in MHz. This is the sum of CPU capacities from all worker nodes in the cluster. | 
**MemoryUsed** | **int64** | Overall memory usage of the cluster, in mebibytes. This is the sum of memory usage across all worker nodes in the cluster. | 
**MemoryCapacity** | **int64** | Total memory capacity of the cluster available for vSphere Namespaces, in mebibytes. This is the sum of memory capacities from all worker nodesin the cluster. | 
**StorageUsed** | **int64** | Overall storage used by the cluster, in mebibytes. This is the sum of storage used across all worker nodes in the cluster. | 
**StorageCapacity** | **int64** | Overall storage capacity of the cluster available for vSphere Namespaces, in mebibytes. This is the sum of total storage available from all worker nodes in the cluster. | 

## Methods

### NewVcenterNamespaceManagementClustersStats

`func NewVcenterNamespaceManagementClustersStats(cpuUsed int64, cpuCapacity int64, memoryUsed int64, memoryCapacity int64, storageUsed int64, storageCapacity int64, ) *VcenterNamespaceManagementClustersStats`

NewVcenterNamespaceManagementClustersStats instantiates a new VcenterNamespaceManagementClustersStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersStatsWithDefaults

`func NewVcenterNamespaceManagementClustersStatsWithDefaults() *VcenterNamespaceManagementClustersStats`

NewVcenterNamespaceManagementClustersStatsWithDefaults instantiates a new VcenterNamespaceManagementClustersStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuUsed

`func (o *VcenterNamespaceManagementClustersStats) GetCpuUsed() int64`

GetCpuUsed returns the CpuUsed field if non-nil, zero value otherwise.

### GetCpuUsedOk

`func (o *VcenterNamespaceManagementClustersStats) GetCpuUsedOk() (*int64, bool)`

GetCpuUsedOk returns a tuple with the CpuUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsed

`func (o *VcenterNamespaceManagementClustersStats) SetCpuUsed(v int64)`

SetCpuUsed sets CpuUsed field to given value.


### GetCpuCapacity

`func (o *VcenterNamespaceManagementClustersStats) GetCpuCapacity() int64`

GetCpuCapacity returns the CpuCapacity field if non-nil, zero value otherwise.

### GetCpuCapacityOk

`func (o *VcenterNamespaceManagementClustersStats) GetCpuCapacityOk() (*int64, bool)`

GetCpuCapacityOk returns a tuple with the CpuCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuCapacity

`func (o *VcenterNamespaceManagementClustersStats) SetCpuCapacity(v int64)`

SetCpuCapacity sets CpuCapacity field to given value.


### GetMemoryUsed

`func (o *VcenterNamespaceManagementClustersStats) GetMemoryUsed() int64`

GetMemoryUsed returns the MemoryUsed field if non-nil, zero value otherwise.

### GetMemoryUsedOk

`func (o *VcenterNamespaceManagementClustersStats) GetMemoryUsedOk() (*int64, bool)`

GetMemoryUsedOk returns a tuple with the MemoryUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUsed

`func (o *VcenterNamespaceManagementClustersStats) SetMemoryUsed(v int64)`

SetMemoryUsed sets MemoryUsed field to given value.


### GetMemoryCapacity

`func (o *VcenterNamespaceManagementClustersStats) GetMemoryCapacity() int64`

GetMemoryCapacity returns the MemoryCapacity field if non-nil, zero value otherwise.

### GetMemoryCapacityOk

`func (o *VcenterNamespaceManagementClustersStats) GetMemoryCapacityOk() (*int64, bool)`

GetMemoryCapacityOk returns a tuple with the MemoryCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryCapacity

`func (o *VcenterNamespaceManagementClustersStats) SetMemoryCapacity(v int64)`

SetMemoryCapacity sets MemoryCapacity field to given value.


### GetStorageUsed

`func (o *VcenterNamespaceManagementClustersStats) GetStorageUsed() int64`

GetStorageUsed returns the StorageUsed field if non-nil, zero value otherwise.

### GetStorageUsedOk

`func (o *VcenterNamespaceManagementClustersStats) GetStorageUsedOk() (*int64, bool)`

GetStorageUsedOk returns a tuple with the StorageUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageUsed

`func (o *VcenterNamespaceManagementClustersStats) SetStorageUsed(v int64)`

SetStorageUsed sets StorageUsed field to given value.


### GetStorageCapacity

`func (o *VcenterNamespaceManagementClustersStats) GetStorageCapacity() int64`

GetStorageCapacity returns the StorageCapacity field if non-nil, zero value otherwise.

### GetStorageCapacityOk

`func (o *VcenterNamespaceManagementClustersStats) GetStorageCapacityOk() (*int64, bool)`

GetStorageCapacityOk returns a tuple with the StorageCapacity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageCapacity

`func (o *VcenterNamespaceManagementClustersStats) SetStorageCapacity(v int64)`

SetStorageCapacity sets StorageCapacity field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


