# VcenterResourcePoolCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource pool. | 
**Parent** | **string** | Parent of the created resource pool. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ResourcePool. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ResourcePool. | 
**CpuAllocation** | Pointer to [**VcenterResourcePoolResourceAllocationCreateSpec**](VcenterResourcePoolResourceAllocationCreateSpec.md) |  | [optional] 
**MemoryAllocation** | Pointer to [**VcenterResourcePoolResourceAllocationCreateSpec**](VcenterResourcePoolResourceAllocationCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterResourcePoolCreateSpec

`func NewVcenterResourcePoolCreateSpec(name string, parent string, ) *VcenterResourcePoolCreateSpec`

NewVcenterResourcePoolCreateSpec instantiates a new VcenterResourcePoolCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolCreateSpecWithDefaults

`func NewVcenterResourcePoolCreateSpecWithDefaults() *VcenterResourcePoolCreateSpec`

NewVcenterResourcePoolCreateSpecWithDefaults instantiates a new VcenterResourcePoolCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterResourcePoolCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterResourcePoolCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterResourcePoolCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetParent

`func (o *VcenterResourcePoolCreateSpec) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *VcenterResourcePoolCreateSpec) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *VcenterResourcePoolCreateSpec) SetParent(v string)`

SetParent sets Parent field to given value.


### GetCpuAllocation

`func (o *VcenterResourcePoolCreateSpec) GetCpuAllocation() VcenterResourcePoolResourceAllocationCreateSpec`

GetCpuAllocation returns the CpuAllocation field if non-nil, zero value otherwise.

### GetCpuAllocationOk

`func (o *VcenterResourcePoolCreateSpec) GetCpuAllocationOk() (*VcenterResourcePoolResourceAllocationCreateSpec, bool)`

GetCpuAllocationOk returns a tuple with the CpuAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAllocation

`func (o *VcenterResourcePoolCreateSpec) SetCpuAllocation(v VcenterResourcePoolResourceAllocationCreateSpec)`

SetCpuAllocation sets CpuAllocation field to given value.

### HasCpuAllocation

`func (o *VcenterResourcePoolCreateSpec) HasCpuAllocation() bool`

HasCpuAllocation returns a boolean if a field has been set.

### GetMemoryAllocation

`func (o *VcenterResourcePoolCreateSpec) GetMemoryAllocation() VcenterResourcePoolResourceAllocationCreateSpec`

GetMemoryAllocation returns the MemoryAllocation field if non-nil, zero value otherwise.

### GetMemoryAllocationOk

`func (o *VcenterResourcePoolCreateSpec) GetMemoryAllocationOk() (*VcenterResourcePoolResourceAllocationCreateSpec, bool)`

GetMemoryAllocationOk returns a tuple with the MemoryAllocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryAllocation

`func (o *VcenterResourcePoolCreateSpec) SetMemoryAllocation(v VcenterResourcePoolResourceAllocationCreateSpec)`

SetMemoryAllocation sets MemoryAllocation field to given value.

### HasMemoryAllocation

`func (o *VcenterResourcePoolCreateSpec) HasMemoryAllocation() bool`

HasMemoryAllocation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


