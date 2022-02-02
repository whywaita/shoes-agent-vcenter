# VcenterResourcePoolResourceAllocationCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservation** | Pointer to **int64** | Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU. If unset or empty, ResourcePool.ResourceAllocationCreateSpec.reservation will be set to 0. | [optional] 
**ExpandableReservation** | Pointer to **bool** | In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation. If unset or empty, ResourcePool.ResourceAllocationCreateSpec.expandable-reservation will be set to true. | [optional] 
**Limit** | Pointer to **int64** | The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU. If unset or empty, ResourcePool.ResourceAllocationCreateSpec.limit will be set to -1. | [optional] 
**Shares** | Pointer to [**VcenterResourcePoolSharesInfo**](VcenterResourcePoolSharesInfo.md) |  | [optional] 

## Methods

### NewVcenterResourcePoolResourceAllocationCreateSpec

`func NewVcenterResourcePoolResourceAllocationCreateSpec() *VcenterResourcePoolResourceAllocationCreateSpec`

NewVcenterResourcePoolResourceAllocationCreateSpec instantiates a new VcenterResourcePoolResourceAllocationCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolResourceAllocationCreateSpecWithDefaults

`func NewVcenterResourcePoolResourceAllocationCreateSpecWithDefaults() *VcenterResourcePoolResourceAllocationCreateSpec`

NewVcenterResourcePoolResourceAllocationCreateSpecWithDefaults instantiates a new VcenterResourcePoolResourceAllocationCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservation

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetReservation() int64`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetReservationOk() (*int64, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) SetReservation(v int64)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetExpandableReservation() bool`

GetExpandableReservation returns the ExpandableReservation field if non-nil, zero value otherwise.

### GetExpandableReservationOk

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetExpandableReservationOk() (*bool, bool)`

GetExpandableReservationOk returns a tuple with the ExpandableReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) SetExpandableReservation(v bool)`

SetExpandableReservation sets ExpandableReservation field to given value.

### HasExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) HasExpandableReservation() bool`

HasExpandableReservation returns a boolean if a field has been set.

### GetLimit

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetShares

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetShares() VcenterResourcePoolSharesInfo`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) GetSharesOk() (*VcenterResourcePoolSharesInfo, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) SetShares(v VcenterResourcePoolSharesInfo)`

SetShares sets Shares field to given value.

### HasShares

`func (o *VcenterResourcePoolResourceAllocationCreateSpec) HasShares() bool`

HasShares returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


