# VcenterResourcePoolResourceAllocationUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservation** | Pointer to **int64** | Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU. If unset or empty, ResourcePool.ResourceAllocationUpdateSpec.reservation will be set to 0. | [optional] 
**ExpandableReservation** | Pointer to **bool** | In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation. If unset or empty, ResourcePool.ResourceAllocationUpdateSpec.expandable-reservation will be set to true. | [optional] 
**Limit** | Pointer to **int64** | The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU. If unset or empty, ResourcePool.ResourceAllocationUpdateSpec.limit will be set to -1. | [optional] 
**Shares** | Pointer to [**VcenterResourcePoolSharesInfo**](VcenterResourcePoolSharesInfo.md) |  | [optional] 

## Methods

### NewVcenterResourcePoolResourceAllocationUpdateSpec

`func NewVcenterResourcePoolResourceAllocationUpdateSpec() *VcenterResourcePoolResourceAllocationUpdateSpec`

NewVcenterResourcePoolResourceAllocationUpdateSpec instantiates a new VcenterResourcePoolResourceAllocationUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolResourceAllocationUpdateSpecWithDefaults

`func NewVcenterResourcePoolResourceAllocationUpdateSpecWithDefaults() *VcenterResourcePoolResourceAllocationUpdateSpec`

NewVcenterResourcePoolResourceAllocationUpdateSpecWithDefaults instantiates a new VcenterResourcePoolResourceAllocationUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservation

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetReservation() int64`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetReservationOk() (*int64, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) SetReservation(v int64)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetExpandableReservation() bool`

GetExpandableReservation returns the ExpandableReservation field if non-nil, zero value otherwise.

### GetExpandableReservationOk

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetExpandableReservationOk() (*bool, bool)`

GetExpandableReservationOk returns a tuple with the ExpandableReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) SetExpandableReservation(v bool)`

SetExpandableReservation sets ExpandableReservation field to given value.

### HasExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) HasExpandableReservation() bool`

HasExpandableReservation returns a boolean if a field has been set.

### GetLimit

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetShares

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetShares() VcenterResourcePoolSharesInfo`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) GetSharesOk() (*VcenterResourcePoolSharesInfo, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) SetShares(v VcenterResourcePoolSharesInfo)`

SetShares sets Shares field to given value.

### HasShares

`func (o *VcenterResourcePoolResourceAllocationUpdateSpec) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


