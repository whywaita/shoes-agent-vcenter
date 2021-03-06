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

// VcenterVmGuestFilesystemFilesListResult struct for VcenterVmGuestFilesystemFilesListResult
type VcenterVmGuestFilesystemFilesListResult struct {
	// A list of Files.Summary structures containing information for all the matching files.
	Files []VcenterVmGuestFilesystemFilesSummary `json:"files"`
	// The total number of results from the Files.list. This is a hint to the user of the iterator regarding how many items are available to be retrieved. The total could change if the inventory of items are being changed.
	Total int64 `json:"total"`
	// Positional index into the logical item list of the first item returned in the list of results. The first item in the logical item list has an index of 0. This is a hint to the user of the iterator regarding the logical position in the iteration. For example, this can be used to display to the user which page of the iteration is being shown. The total could change if the inventory of items are being changed. If unset no items were returned.
	StartIndex *int64 `json:"start_index,omitempty"`
	// Positional index into the logical item list of the last item returned in the list of results. The first item in the logical item list has an index of 0. This is a hint to the user of the iterator regarding the logical position in the iteration. For example, this can be used to display to the user which page of the iteration is being shown. The total could change if the inventory of items are being changed. If unset no items were returned.
	EndIndex *int64 `json:"end_index,omitempty"`
	Status VcenterVmGuestFilesystemFilesLastIterationStatus `json:"status"`
}

// NewVcenterVmGuestFilesystemFilesListResult instantiates a new VcenterVmGuestFilesystemFilesListResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestFilesystemFilesListResult(files []VcenterVmGuestFilesystemFilesSummary, total int64, status VcenterVmGuestFilesystemFilesLastIterationStatus) *VcenterVmGuestFilesystemFilesListResult {
	this := VcenterVmGuestFilesystemFilesListResult{}
	this.Files = files
	this.Total = total
	this.Status = status
	return &this
}

// NewVcenterVmGuestFilesystemFilesListResultWithDefaults instantiates a new VcenterVmGuestFilesystemFilesListResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestFilesystemFilesListResultWithDefaults() *VcenterVmGuestFilesystemFilesListResult {
	this := VcenterVmGuestFilesystemFilesListResult{}
	return &this
}

// GetFiles returns the Files field value
func (o *VcenterVmGuestFilesystemFilesListResult) GetFiles() []VcenterVmGuestFilesystemFilesSummary {
	if o == nil {
		var ret []VcenterVmGuestFilesystemFilesSummary
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) GetFilesOk() (*[]VcenterVmGuestFilesystemFilesSummary, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Files, true
}

// SetFiles sets field value
func (o *VcenterVmGuestFilesystemFilesListResult) SetFiles(v []VcenterVmGuestFilesystemFilesSummary) {
	o.Files = v
}

// GetTotal returns the Total field value
func (o *VcenterVmGuestFilesystemFilesListResult) GetTotal() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) GetTotalOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *VcenterVmGuestFilesystemFilesListResult) SetTotal(v int64) {
	o.Total = v
}

// GetStartIndex returns the StartIndex field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemFilesListResult) GetStartIndex() int64 {
	if o == nil || o.StartIndex == nil {
		var ret int64
		return ret
	}
	return *o.StartIndex
}

// GetStartIndexOk returns a tuple with the StartIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) GetStartIndexOk() (*int64, bool) {
	if o == nil || o.StartIndex == nil {
		return nil, false
	}
	return o.StartIndex, true
}

// HasStartIndex returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) HasStartIndex() bool {
	if o != nil && o.StartIndex != nil {
		return true
	}

	return false
}

// SetStartIndex gets a reference to the given int64 and assigns it to the StartIndex field.
func (o *VcenterVmGuestFilesystemFilesListResult) SetStartIndex(v int64) {
	o.StartIndex = &v
}

// GetEndIndex returns the EndIndex field value if set, zero value otherwise.
func (o *VcenterVmGuestFilesystemFilesListResult) GetEndIndex() int64 {
	if o == nil || o.EndIndex == nil {
		var ret int64
		return ret
	}
	return *o.EndIndex
}

// GetEndIndexOk returns a tuple with the EndIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) GetEndIndexOk() (*int64, bool) {
	if o == nil || o.EndIndex == nil {
		return nil, false
	}
	return o.EndIndex, true
}

// HasEndIndex returns a boolean if a field has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) HasEndIndex() bool {
	if o != nil && o.EndIndex != nil {
		return true
	}

	return false
}

// SetEndIndex gets a reference to the given int64 and assigns it to the EndIndex field.
func (o *VcenterVmGuestFilesystemFilesListResult) SetEndIndex(v int64) {
	o.EndIndex = &v
}

// GetStatus returns the Status field value
func (o *VcenterVmGuestFilesystemFilesListResult) GetStatus() VcenterVmGuestFilesystemFilesLastIterationStatus {
	if o == nil {
		var ret VcenterVmGuestFilesystemFilesLastIterationStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestFilesystemFilesListResult) GetStatusOk() (*VcenterVmGuestFilesystemFilesLastIterationStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VcenterVmGuestFilesystemFilesListResult) SetStatus(v VcenterVmGuestFilesystemFilesLastIterationStatus) {
	o.Status = v
}

func (o VcenterVmGuestFilesystemFilesListResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["files"] = o.Files
	}
	if true {
		toSerialize["total"] = o.Total
	}
	if o.StartIndex != nil {
		toSerialize["start_index"] = o.StartIndex
	}
	if o.EndIndex != nil {
		toSerialize["end_index"] = o.EndIndex
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestFilesystemFilesListResult struct {
	value *VcenterVmGuestFilesystemFilesListResult
	isSet bool
}

func (v NullableVcenterVmGuestFilesystemFilesListResult) Get() *VcenterVmGuestFilesystemFilesListResult {
	return v.value
}

func (v *NullableVcenterVmGuestFilesystemFilesListResult) Set(val *VcenterVmGuestFilesystemFilesListResult) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestFilesystemFilesListResult) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestFilesystemFilesListResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestFilesystemFilesListResult(val *VcenterVmGuestFilesystemFilesListResult) *NullableVcenterVmGuestFilesystemFilesListResult {
	return &NullableVcenterVmGuestFilesystemFilesListResult{value: val, isSet: true}
}

func (v NullableVcenterVmGuestFilesystemFilesListResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestFilesystemFilesListResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


