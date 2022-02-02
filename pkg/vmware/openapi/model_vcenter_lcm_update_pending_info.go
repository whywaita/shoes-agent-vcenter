/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// VcenterLcmUpdatePendingInfo struct for VcenterLcmUpdatePendingInfo
type VcenterLcmUpdatePendingInfo struct {
	// Description of the vSphere update
	Description string `json:"description"`
	// Identifier of the given vSphere update When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.lcm.update.pending. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.lcm.update.pending.
	PendingUpdate string `json:"pending_update"`
	// Version of the vSphere update or patch
	Version string `json:"version"`
	// Release date of the vSphere update or patch
	ReleaseDate time.Time `json:"release_date"`
	Severity VcenterLcmUpdatePendingSeverityType `json:"severity"`
	// Build number of the vCenter Release
	Build string `json:"build"`
	UpdateType VcenterLcmUpdatePendingUpdateType `json:"update_type"`
	Category VcenterLcmUpdatePendingCategory `json:"category"`
	// Flag to suggest a reboot after the release is applied
	RebootRequired bool `json:"reboot_required"`
	// VAMI or ISO URL for update or upgrade execute phase redirection
	ExecuteURL string `json:"execute_URL"`
	// List of URI pointing to patch or update release notes
	ReleaseNotes []string `json:"release_notes"`
}

// NewVcenterLcmUpdatePendingInfo instantiates a new VcenterLcmUpdatePendingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterLcmUpdatePendingInfo(description string, pendingUpdate string, version string, releaseDate time.Time, severity VcenterLcmUpdatePendingSeverityType, build string, updateType VcenterLcmUpdatePendingUpdateType, category VcenterLcmUpdatePendingCategory, rebootRequired bool, executeURL string, releaseNotes []string) *VcenterLcmUpdatePendingInfo {
	this := VcenterLcmUpdatePendingInfo{}
	this.Description = description
	this.PendingUpdate = pendingUpdate
	this.Version = version
	this.ReleaseDate = releaseDate
	this.Severity = severity
	this.Build = build
	this.UpdateType = updateType
	this.Category = category
	this.RebootRequired = rebootRequired
	this.ExecuteURL = executeURL
	this.ReleaseNotes = releaseNotes
	return &this
}

// NewVcenterLcmUpdatePendingInfoWithDefaults instantiates a new VcenterLcmUpdatePendingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterLcmUpdatePendingInfoWithDefaults() *VcenterLcmUpdatePendingInfo {
	this := VcenterLcmUpdatePendingInfo{}
	return &this
}

// GetDescription returns the Description field value
func (o *VcenterLcmUpdatePendingInfo) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetDescriptionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *VcenterLcmUpdatePendingInfo) SetDescription(v string) {
	o.Description = v
}

// GetPendingUpdate returns the PendingUpdate field value
func (o *VcenterLcmUpdatePendingInfo) GetPendingUpdate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PendingUpdate
}

// GetPendingUpdateOk returns a tuple with the PendingUpdate field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetPendingUpdateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PendingUpdate, true
}

// SetPendingUpdate sets field value
func (o *VcenterLcmUpdatePendingInfo) SetPendingUpdate(v string) {
	o.PendingUpdate = v
}

// GetVersion returns the Version field value
func (o *VcenterLcmUpdatePendingInfo) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetVersionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *VcenterLcmUpdatePendingInfo) SetVersion(v string) {
	o.Version = v
}

// GetReleaseDate returns the ReleaseDate field value
func (o *VcenterLcmUpdatePendingInfo) GetReleaseDate() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReleaseDate
}

// GetReleaseDateOk returns a tuple with the ReleaseDate field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetReleaseDateOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseDate, true
}

// SetReleaseDate sets field value
func (o *VcenterLcmUpdatePendingInfo) SetReleaseDate(v time.Time) {
	o.ReleaseDate = v
}

// GetSeverity returns the Severity field value
func (o *VcenterLcmUpdatePendingInfo) GetSeverity() VcenterLcmUpdatePendingSeverityType {
	if o == nil {
		var ret VcenterLcmUpdatePendingSeverityType
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetSeverityOk() (*VcenterLcmUpdatePendingSeverityType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *VcenterLcmUpdatePendingInfo) SetSeverity(v VcenterLcmUpdatePendingSeverityType) {
	o.Severity = v
}

// GetBuild returns the Build field value
func (o *VcenterLcmUpdatePendingInfo) GetBuild() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Build
}

// GetBuildOk returns a tuple with the Build field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetBuildOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Build, true
}

// SetBuild sets field value
func (o *VcenterLcmUpdatePendingInfo) SetBuild(v string) {
	o.Build = v
}

// GetUpdateType returns the UpdateType field value
func (o *VcenterLcmUpdatePendingInfo) GetUpdateType() VcenterLcmUpdatePendingUpdateType {
	if o == nil {
		var ret VcenterLcmUpdatePendingUpdateType
		return ret
	}

	return o.UpdateType
}

// GetUpdateTypeOk returns a tuple with the UpdateType field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetUpdateTypeOk() (*VcenterLcmUpdatePendingUpdateType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.UpdateType, true
}

// SetUpdateType sets field value
func (o *VcenterLcmUpdatePendingInfo) SetUpdateType(v VcenterLcmUpdatePendingUpdateType) {
	o.UpdateType = v
}

// GetCategory returns the Category field value
func (o *VcenterLcmUpdatePendingInfo) GetCategory() VcenterLcmUpdatePendingCategory {
	if o == nil {
		var ret VcenterLcmUpdatePendingCategory
		return ret
	}

	return o.Category
}

// GetCategoryOk returns a tuple with the Category field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetCategoryOk() (*VcenterLcmUpdatePendingCategory, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Category, true
}

// SetCategory sets field value
func (o *VcenterLcmUpdatePendingInfo) SetCategory(v VcenterLcmUpdatePendingCategory) {
	o.Category = v
}

// GetRebootRequired returns the RebootRequired field value
func (o *VcenterLcmUpdatePendingInfo) GetRebootRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.RebootRequired
}

// GetRebootRequiredOk returns a tuple with the RebootRequired field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetRebootRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RebootRequired, true
}

// SetRebootRequired sets field value
func (o *VcenterLcmUpdatePendingInfo) SetRebootRequired(v bool) {
	o.RebootRequired = v
}

// GetExecuteURL returns the ExecuteURL field value
func (o *VcenterLcmUpdatePendingInfo) GetExecuteURL() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExecuteURL
}

// GetExecuteURLOk returns a tuple with the ExecuteURL field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetExecuteURLOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ExecuteURL, true
}

// SetExecuteURL sets field value
func (o *VcenterLcmUpdatePendingInfo) SetExecuteURL(v string) {
	o.ExecuteURL = v
}

// GetReleaseNotes returns the ReleaseNotes field value
func (o *VcenterLcmUpdatePendingInfo) GetReleaseNotes() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ReleaseNotes
}

// GetReleaseNotesOk returns a tuple with the ReleaseNotes field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmUpdatePendingInfo) GetReleaseNotesOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ReleaseNotes, true
}

// SetReleaseNotes sets field value
func (o *VcenterLcmUpdatePendingInfo) SetReleaseNotes(v []string) {
	o.ReleaseNotes = v
}

func (o VcenterLcmUpdatePendingInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["pending_update"] = o.PendingUpdate
	}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["release_date"] = o.ReleaseDate
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["build"] = o.Build
	}
	if true {
		toSerialize["update_type"] = o.UpdateType
	}
	if true {
		toSerialize["category"] = o.Category
	}
	if true {
		toSerialize["reboot_required"] = o.RebootRequired
	}
	if true {
		toSerialize["execute_URL"] = o.ExecuteURL
	}
	if true {
		toSerialize["release_notes"] = o.ReleaseNotes
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterLcmUpdatePendingInfo struct {
	value *VcenterLcmUpdatePendingInfo
	isSet bool
}

func (v NullableVcenterLcmUpdatePendingInfo) Get() *VcenterLcmUpdatePendingInfo {
	return v.value
}

func (v *NullableVcenterLcmUpdatePendingInfo) Set(val *VcenterLcmUpdatePendingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterLcmUpdatePendingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterLcmUpdatePendingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterLcmUpdatePendingInfo(val *VcenterLcmUpdatePendingInfo) *NullableVcenterLcmUpdatePendingInfo {
	return &NullableVcenterLcmUpdatePendingInfo{value: val, isSet: true}
}

func (v NullableVcenterLcmUpdatePendingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterLcmUpdatePendingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


