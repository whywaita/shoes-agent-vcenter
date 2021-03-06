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

// VcenterDeploymentDataMigrationEstimate struct for VcenterDeploymentDataMigrationEstimate
type VcenterDeploymentDataMigrationEstimate struct {
	// The time estimated to export data from the source vCenter Server.
	EstimatedExportTime int64 `json:"estimated_export_time"`
	// The time estimated to import data to the upgraded vCenter Server.
	EstimatedImportTime int64 `json:"estimated_import_time"`
	// The extra free space required on source vCenter Server.
	RequiredFreeDiskSpaceOnSource float64 `json:"required_free_disk_space_on_source"`
}

// NewVcenterDeploymentDataMigrationEstimate instantiates a new VcenterDeploymentDataMigrationEstimate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterDeploymentDataMigrationEstimate(estimatedExportTime int64, estimatedImportTime int64, requiredFreeDiskSpaceOnSource float64) *VcenterDeploymentDataMigrationEstimate {
	this := VcenterDeploymentDataMigrationEstimate{}
	this.EstimatedExportTime = estimatedExportTime
	this.EstimatedImportTime = estimatedImportTime
	this.RequiredFreeDiskSpaceOnSource = requiredFreeDiskSpaceOnSource
	return &this
}

// NewVcenterDeploymentDataMigrationEstimateWithDefaults instantiates a new VcenterDeploymentDataMigrationEstimate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterDeploymentDataMigrationEstimateWithDefaults() *VcenterDeploymentDataMigrationEstimate {
	this := VcenterDeploymentDataMigrationEstimate{}
	return &this
}

// GetEstimatedExportTime returns the EstimatedExportTime field value
func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedExportTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EstimatedExportTime
}

// GetEstimatedExportTimeOk returns a tuple with the EstimatedExportTime field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedExportTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EstimatedExportTime, true
}

// SetEstimatedExportTime sets field value
func (o *VcenterDeploymentDataMigrationEstimate) SetEstimatedExportTime(v int64) {
	o.EstimatedExportTime = v
}

// GetEstimatedImportTime returns the EstimatedImportTime field value
func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedImportTime() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.EstimatedImportTime
}

// GetEstimatedImportTimeOk returns a tuple with the EstimatedImportTime field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentDataMigrationEstimate) GetEstimatedImportTimeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EstimatedImportTime, true
}

// SetEstimatedImportTime sets field value
func (o *VcenterDeploymentDataMigrationEstimate) SetEstimatedImportTime(v int64) {
	o.EstimatedImportTime = v
}

// GetRequiredFreeDiskSpaceOnSource returns the RequiredFreeDiskSpaceOnSource field value
func (o *VcenterDeploymentDataMigrationEstimate) GetRequiredFreeDiskSpaceOnSource() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.RequiredFreeDiskSpaceOnSource
}

// GetRequiredFreeDiskSpaceOnSourceOk returns a tuple with the RequiredFreeDiskSpaceOnSource field value
// and a boolean to check if the value has been set.
func (o *VcenterDeploymentDataMigrationEstimate) GetRequiredFreeDiskSpaceOnSourceOk() (*float64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequiredFreeDiskSpaceOnSource, true
}

// SetRequiredFreeDiskSpaceOnSource sets field value
func (o *VcenterDeploymentDataMigrationEstimate) SetRequiredFreeDiskSpaceOnSource(v float64) {
	o.RequiredFreeDiskSpaceOnSource = v
}

func (o VcenterDeploymentDataMigrationEstimate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["estimated_export_time"] = o.EstimatedExportTime
	}
	if true {
		toSerialize["estimated_import_time"] = o.EstimatedImportTime
	}
	if true {
		toSerialize["required_free_disk_space_on_source"] = o.RequiredFreeDiskSpaceOnSource
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterDeploymentDataMigrationEstimate struct {
	value *VcenterDeploymentDataMigrationEstimate
	isSet bool
}

func (v NullableVcenterDeploymentDataMigrationEstimate) Get() *VcenterDeploymentDataMigrationEstimate {
	return v.value
}

func (v *NullableVcenterDeploymentDataMigrationEstimate) Set(val *VcenterDeploymentDataMigrationEstimate) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterDeploymentDataMigrationEstimate) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterDeploymentDataMigrationEstimate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterDeploymentDataMigrationEstimate(val *VcenterDeploymentDataMigrationEstimate) *NullableVcenterDeploymentDataMigrationEstimate {
	return &NullableVcenterDeploymentDataMigrationEstimate{value: val, isSet: true}
}

func (v NullableVcenterDeploymentDataMigrationEstimate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterDeploymentDataMigrationEstimate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


