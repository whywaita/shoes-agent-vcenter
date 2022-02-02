# VcenterResourcePoolResourceAllocationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reservation** | **int64** | Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU. | 
**ExpandableReservation** | **bool** | In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation. | 
**Limit** | **int64** | The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU. | 
**Shares** | [**VcenterResourcePoolSharesInfo**](VcenterResourcePoolSharesInfo.md) |  | 

## Methods

### NewVcenterResourcePoolResourceAllocationInfo

`func NewVcenterResourcePoolResourceAllocationInfo(reservation int64, expandableReservation bool, limit int64, shares VcenterResourcePoolSharesInfo, ) *VcenterResourcePoolResourceAllocationInfo`

NewVcenterResourcePoolResourceAllocationInfo instantiates a new VcenterResourcePoolResourceAllocationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterResourcePoolResourceAllocationInfoWithDefaults

`func NewVcenterResourcePoolResourceAllocationInfoWithDefaults() *VcenterResourcePoolResourceAllocationInfo`

NewVcenterResourcePoolResourceAllocationInfoWithDefaults instantiates a new VcenterResourcePoolResourceAllocationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReservation

`func (o *VcenterResourcePoolResourceAllocationInfo) GetReservation() int64`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *VcenterResourcePoolResourceAllocationInfo) GetReservationOk() (*int64, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *VcenterResourcePoolResourceAllocationInfo) SetReservation(v int64)`

SetReservation sets Reservation field to given value.


### GetExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationInfo) GetExpandableReservation() bool`

GetExpandableReservation returns the ExpandableReservation field if non-nil, zero value otherwise.

### GetExpandableReservationOk

`func (o *VcenterResourcePoolResourceAllocationInfo) GetExpandableReservationOk() (*bool, bool)`

GetExpandableReservationOk returns a tuple with the ExpandableReservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpandableReservation

`func (o *VcenterResourcePoolResourceAllocationInfo) SetExpandableReservation(v bool)`

SetExpandableReservation sets ExpandableReservation field to given value.


### GetLimit

`func (o *VcenterResourcePoolResourceAllocationInfo) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *VcenterResourcePoolResourceAllocationInfo) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *VcenterResourcePoolResourceAllocationInfo) SetLimit(v int64)`

SetLimit sets Limit field to given value.


### GetShares

`func (o *VcenterResourcePoolResourceAllocationInfo) GetShares() VcenterResourcePoolSharesInfo`

GetShares returns the Shares field if non-nil, zero value otherwise.

### GetSharesOk

`func (o *VcenterResourcePoolResourceAllocationInfo) GetSharesOk() (*VcenterResourcePoolSharesInfo, bool)`

GetSharesOk returns a tuple with the Shares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShares

`func (o *VcenterResourcePoolResourceAllocationInfo) SetShares(v VcenterResourcePoolSharesInfo)`

SetShares sets Shares field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


