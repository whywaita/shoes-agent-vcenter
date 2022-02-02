# VcenterResourcePoolInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the vCenter Server resource pool. | 
**ResourcePools** | **[]string** | Identifiers of the child resource pools contained in this resource pool. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ResourcePool. | 
**CpuAllocation** | Pointer to [**VcenterResourcePoolResourceAllocationInfo**](VcenterResourcePoolResourceAllocationInfo.md) |  | [optional] 
**MemoryAllocation** | Pointer to [**VcenterResourcePoolResourceAllocationInfo**](VcenterResourcePoolResourceAllocationInfo.md) |  | [optional] 

## Methods

### NewVcenterResourcePoolInfo

`func NewVcenterResourcePoolInfo(name string, resourcePools []string, ) *VcenterResourcePoolInfo`

NewVcenterResourcePoolInfo instantiates a new VcenterResourcePoolInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolInfoWithDefaults

`func NewVcenterResourcePoolInfoWithDefaults() *VcenterResourcePoolInfo`

NewVcenterResourcePoolInfoWithDefaults instantiates a new VcenterResourcePoolInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterResourcePoolInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterResourcePoolInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterResourcePoolInfo) SetName(v string)`

SetName sets Name field to given value.


### GetResourcePools

`func (o *VcenterResourcePoolInfo) GetResourcePools() []string`

GetResourcePools returns the ResourcePools field if non-nil, zero value otherwise.

### GetResourcePoolsOk

`func (o *VcenterResourcePoolInfo) GetResourcePoolsOk() (*[]string, bool)`

GetResourcePoolsOk returns a tuple with the ResourcePools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcePools

`func (o *VcenterResourcePoolInfo) SetResourcePools(v []string)`

SetResourcePools sets ResourcePools field to given value.


### GetCpuAllocation

`func (o *VcenterResourcePoolInfo) GetCpuAllocation() VcenterResourcePoolResourceAllocationInfo`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *VcenterResourcePoolInfo) GetCpuAllocationOk() (*VcenterResourcePoolResourceAllocationInfo, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *VcenterResourcePoolInfo) SetCpuAllocation(v VcenterResourcePoolResourceAllocationInfo)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *VcenterResourcePoolInfo) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *VcenterResourcePoolInfo) GetMemoryAllocation() VcenterResourcePoolResourceAllocationInfo`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *VcenterResourcePoolInfo) GetMemoryAllocationOk() (*VcenterResourcePoolResourceAllocationInfo, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *VcenterResourcePoolInfo) SetMemoryAllocation(v VcenterResourcePoolResourceAllocationInfo)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *VcenterResourcePoolInfo) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


