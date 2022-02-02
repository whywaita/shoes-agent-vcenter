/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// VcenterResourcePoolResourceAllocationInfo struct for VcenterResourcePoolResourceAllocationInfo
type VcenterResourcePoolResourceAllocationInfo struct {
	// Amount of resource that is guaranteed available to a resource pool. Reserved resources are not wasted if they are not used. If the utilization is less than the reservation, the resources can be utilized by other running virtual machines. Units are MB fo memory, and MHz for CPU.
	Reservation int64 `json:"reservation"`
	// In a resource pool with an expandable reservation, the reservation can grow beyond the specified value, if the parent resource pool has unreserved resources. A non-expandable reservation is called a fixed reservation.
	ExpandableReservation bool `json:"expandable_reservation"`
	// The utilization of a resource pool will not exceed this limit, even if there are available resources. This is typically used to ensure a consistent performance of resource pools independent of available resources. If set to -1, then there is no fixed limit on resource usage (only bounded by available resources and shares). Units are MB for memory, and MHz for CPU.
	Limit int64 `json:"limit"`
	Shares VcenterResourcePoolSharesInfo `json:"shares"`
}

// NewVcenterResourcePoolResourceAllocationInfo instantiates a new VcenterResourcePoolResourceAllocationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterResourcePoolResourceAllocationInfo(reservation int64, expandableReservation bool, limit int64, shares VcenterResourcePoolSharesInfo) *VcenterResourcePoolResourceAllocationInfo {
	this := VcenterResourcePoolResourceAllocationInfo{}
	this.Reservation = reservation
	this.ExpandableReservation = expandableReservation
	this.Limit = limit
	this.Shares = shares
	return &this
}

// NewVcenterResourcePoolResourceAllocationInfoWithDefaults instantiates a new VcenterResourcePoolResourceAllocationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterResourcePoolResourceAllocationInfoWithDefaults() *VcenterResourcePoolResourceAllocationInfo {
	this := VcenterResourcePoolResourceAllocationInfo{}
	return &this
}

// GetReservation returns the Reservation field value
func (o *VcenterResourcePoolResourceAllocationInfo) GetReservation() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Reservation
}

// GetReservationOk returns a tuple with the Reservation field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolResourceAllocationInfo) GetReservationOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Reservation, true
}

// SetReservation sets field value
func (o *VcenterResourcePoolResourceAllocationInfo) SetReservation(v int64) {
	o.Reservation = v
}

// GetExpandableReservation returns the ExpandableReservation field value
func (o *VcenterResourcePoolResourceAllocationInfo) GetExpandableReservation() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ExpandableReservation
}

// GetExpandableReservationOk returns a tuple with the ExpandableReservation field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolResourceAllocationInfo) GetExpandableReservationOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExpandableReservation, true
}

// SetExpandableReservation sets field value
func (o *VcenterResourcePoolResourceAllocationInfo) SetExpandableReservation(v bool) {
	o.ExpandableReservation = v
}

// GetLimit returns the Limit field value
func (o *VcenterResourcePoolResourceAllocationInfo) GetLimit() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Limit
}

// GetLimitOk returns a tuple with the Limit field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolResourceAllocationInfo) GetLimitOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Limit, true
}

// SetLimit sets field value
func (o *VcenterResourcePoolResourceAllocationInfo) SetLimit(v int64) {
	o.Limit = v
}

// GetShares returns the Shares field value
func (o *VcenterResourcePoolResourceAllocationInfo) GetShares() VcenterResourcePoolSharesInfo {
	if o == nil {
		var ret VcenterResourcePoolSharesInfo
		return ret
	}

	return o.Shares
}

// GetSharesOk returns a tuple with the Shares field value
// and a boolean to check if the value has been set.
func (o *VcenterResourcePoolResourceAllocationInfo) GetSharesOk() (*VcenterResourcePoolSharesInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Shares, true
}

// SetShares sets field value
func (o *VcenterResourcePoolResourceAllocationInfo) SetShares(v VcenterResourcePoolSharesInfo) {
	o.Shares = v
}

func (o VcenterResourcePoolResourceAllocationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reservation"] = o.Reservation
	}
	if true {
		toSerialize["expandable_reservation"] = o.ExpandableReservation
	}
	if true {
		toSerialize["limit"] = o.Limit
	}
	if true {
		toSerialize["shares"] = o.Shares
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterResourcePoolResourceAllocationInfo struct {
	value *VcenterResourcePoolResourceAllocationInfo
	isSet bool
}

func (v NullableVcenterResourcePoolResourceAllocationInfo) Get() *VcenterResourcePoolResourceAllocationInfo {
	return v.value
}

func (v *NullableVcenterResourcePoolResourceAllocationInfo) Set(val *VcenterResourcePoolResourceAllocationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterResourcePoolResourceAllocationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterResourcePoolResourceAllocationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterResourcePoolResourceAllocationInfo(val *VcenterResourcePoolResourceAllocationInfo) *NullableVcenterResourcePoolResourceAllocationInfo {
	return &NullableVcenterResourcePoolResourceAllocationInfo{value: val, isSet: true}
}

func (v NullableVcenterResourcePoolResourceAllocationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterResourcePoolResourceAllocationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

