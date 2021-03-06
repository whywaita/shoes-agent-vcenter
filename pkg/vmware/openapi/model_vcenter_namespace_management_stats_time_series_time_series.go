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

// VcenterNamespaceManagementStatsTimeSeriesTimeSeries struct for VcenterNamespaceManagementStatsTimeSeriesTimeSeries
type VcenterNamespaceManagementStatsTimeSeriesTimeSeries struct {
	// Counter identifier.
	Counter string `json:"counter"`
	// Sequence of UNIX timestamp values at which statistical values were sampled. https://en.wikipedia.org/wiki/Unix_time
	TimeStamps []int64 `json:"time_stamps"`
	// Sequence of sampled values corresponding to the timestamps in tss.
	Values []int64 `json:"values"`
}

// NewVcenterNamespaceManagementStatsTimeSeriesTimeSeries instantiates a new VcenterNamespaceManagementStatsTimeSeriesTimeSeries object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementStatsTimeSeriesTimeSeries(counter string, timeStamps []int64, values []int64) *VcenterNamespaceManagementStatsTimeSeriesTimeSeries {
	this := VcenterNamespaceManagementStatsTimeSeriesTimeSeries{}
	this.Counter = counter
	this.TimeStamps = timeStamps
	this.Values = values
	return &this
}

// NewVcenterNamespaceManagementStatsTimeSeriesTimeSeriesWithDefaults instantiates a new VcenterNamespaceManagementStatsTimeSeriesTimeSeries object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementStatsTimeSeriesTimeSeriesWithDefaults() *VcenterNamespaceManagementStatsTimeSeriesTimeSeries {
	this := VcenterNamespaceManagementStatsTimeSeriesTimeSeries{}
	return &this
}

// GetCounter returns the Counter field value
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetCounter() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Counter
}

// GetCounterOk returns a tuple with the Counter field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetCounterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Counter, true
}

// SetCounter sets field value
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) SetCounter(v string) {
	o.Counter = v
}

// GetTimeStamps returns the TimeStamps field value
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetTimeStamps() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.TimeStamps
}

// GetTimeStampsOk returns a tuple with the TimeStamps field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetTimeStampsOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TimeStamps, true
}

// SetTimeStamps sets field value
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) SetTimeStamps(v []int64) {
	o.TimeStamps = v
}

// GetValues returns the Values field value
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetValues() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.Values
}

// GetValuesOk returns a tuple with the Values field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) GetValuesOk() (*[]int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Values, true
}

// SetValues sets field value
func (o *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) SetValues(v []int64) {
	o.Values = v
}

func (o VcenterNamespaceManagementStatsTimeSeriesTimeSeries) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counter"] = o.Counter
	}
	if true {
		toSerialize["time_stamps"] = o.TimeStamps
	}
	if true {
		toSerialize["values"] = o.Values
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries struct {
	value *VcenterNamespaceManagementStatsTimeSeriesTimeSeries
	isSet bool
}

func (v NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries) Get() *VcenterNamespaceManagementStatsTimeSeriesTimeSeries {
	return v.value
}

func (v *NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries) Set(val *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries(val *VcenterNamespaceManagementStatsTimeSeriesTimeSeries) *NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries {
	return &NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementStatsTimeSeriesTimeSeries) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


