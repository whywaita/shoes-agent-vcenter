# VcenterResourcePoolUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the resource pool. if unset or empty, the name of the resource pool will not be changed. | [optional] 
**CpuAllocation** | Pointer to [**VcenterResourcePoolResourceAllocationUpdateSpec**](VcenterResourcePoolResourceAllocationUpdateSpec.md) |  | [optional] 
**MemoryAllocation** | Pointer to [**VcenterResourcePoolResourceAllocationUpdateSpec**](VcenterResourcePoolResourceAllocationUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterResourcePoolUpdateSpec

`func NewVcenterResourcePoolUpdateSpec() *VcenterResourcePoolUpdateSpec`

NewVcenterResourcePoolUpdateSpec instantiates a new VcenterResourcePoolUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolUpdateSpecWithDefaults

`func NewVcenterResourcePoolUpdateSpecWithDefaults() *VcenterResourcePoolUpdateSpec`

NewVcenterResourcePoolUpdateSpecWithDefaults instantiates a new VcenterResourcePoolUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterResourcePoolUpdateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterResourcePoolUpdateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterResourcePoolUpdateSpec) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VcenterResourcePoolUpdateSpec) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCpuAllocation

`func (o *VcenterResourcePoolUpdateSpec) GetCpuAllocation() VcenterResourcePoolResourceAllocationUpdateSpec`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *VcenterResourcePoolUpdateSpec) GetCpuAllocationOk() (*VcenterResourcePoolResourceAllocationUpdateSpec, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *VcenterResourcePoolUpdateSpec) SetCpuAllocation(v VcenterResourcePoolResourceAllocationUpdateSpec)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *VcenterResourcePoolUpdateSpec) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *VcenterResourcePoolUpdateSpec) GetMemoryAllocation() VcenterResourcePoolResourceAllocationUpdateSpec`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *VcenterResourcePoolUpdateSpec) GetMemoryAllocationOk() (*VcenterResourcePoolResourceAllocationUpdateSpec, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *VcenterResourcePoolUpdateSpec) SetMemoryAllocation(v VcenterResourcePoolResourceAllocationUpdateSpec)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *VcenterResourcePoolUpdateSpec) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


