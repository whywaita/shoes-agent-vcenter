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

// VcenterHvcLinksSyncProvidersInfo struct for VcenterHvcLinksSyncProvidersInfo
type VcenterHvcLinksSyncProvidersInfo struct {
	// Last sync time for the provider. This indicates the last time that either a background sync or a force sync was started for the provider *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	LastSyncTime *time.Time `json:"last_sync_time,omitempty"`
	Status VcenterHvcLinksSyncProvidersStatus `json:"status"`
	// Sync Polling interval between local and remote replicas for the provider *Warning:* This attribute is available as Technology Preview. These are early access APIs provided to test, automate and provide feedback on the feature. Since this can change based on feedback, VMware does not guarantee backwards compatibility and recommends against using them in production environments. Some Technology Preview APIs might only be applicable to specific environments.
	PollingIntervalInSeconds int64 `json:"polling_interval_in_seconds"`
	CurrentSessionInfo *VcenterHvcLinksSyncProvidersSessionInfo `json:"current_session_info,omitempty"`
	StatusMessage *VapiStdLocalizableMessage `json:"status_message,omitempty"`
}

// NewVcenterHvcLinksSyncProvidersInfo instantiates a new VcenterHvcLinksSyncProvidersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterHvcLinksSyncProvidersInfo(status VcenterHvcLinksSyncProvidersStatus, pollingIntervalInSeconds int64) *VcenterHvcLinksSyncProvidersInfo {
	this := VcenterHvcLinksSyncProvidersInfo{}
	this.Status = status
	this.PollingIntervalInSeconds = pollingIntervalInSeconds
	return &this
}

// NewVcenterHvcLinksSyncProvidersInfoWithDefaults instantiates a new VcenterHvcLinksSyncProvidersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterHvcLinksSyncProvidersInfoWithDefaults() *VcenterHvcLinksSyncProvidersInfo {
	this := VcenterHvcLinksSyncProvidersInfo{}
	return &this
}

// GetLastSyncTime returns the LastSyncTime field value if set, zero value otherwise.
func (o *VcenterHvcLinksSyncProvidersInfo) GetLastSyncTime() time.Time {
	if o == nil || o.LastSyncTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastSyncTime
}

// GetLastSyncTimeOk returns a tuple with the LastSyncTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) GetLastSyncTimeOk() (*time.Time, bool) {
	if o == nil || o.LastSyncTime == nil {
		return nil, false
	}
	return o.LastSyncTime, true
}

// HasLastSyncTime returns a boolean if a field has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) HasLastSyncTime() bool {
	if o != nil && o.LastSyncTime != nil {
		return true
	}

	return false
}

// SetLastSyncTime gets a reference to the given time.Time and assigns it to the LastSyncTime field.
func (o *VcenterHvcLinksSyncProvidersInfo) SetLastSyncTime(v time.Time) {
	o.LastSyncTime = &v
}

// GetStatus returns the Status field value
func (o *VcenterHvcLinksSyncProvidersInfo) GetStatus() VcenterHvcLinksSyncProvidersStatus {
	if o == nil {
		var ret VcenterHvcLinksSyncProvidersStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) GetStatusOk() (*VcenterHvcLinksSyncProvidersStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *VcenterHvcLinksSyncProvidersInfo) SetStatus(v VcenterHvcLinksSyncProvidersStatus) {
	o.Status = v
}

// GetPollingIntervalInSeconds returns the PollingIntervalInSeconds field value
func (o *VcenterHvcLinksSyncProvidersInfo) GetPollingIntervalInSeconds() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PollingIntervalInSeconds
}

// GetPollingIntervalInSecondsOk returns a tuple with the PollingIntervalInSeconds field value
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) GetPollingIntervalInSecondsOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PollingIntervalInSeconds, true
}

// SetPollingIntervalInSeconds sets field value
func (o *VcenterHvcLinksSyncProvidersInfo) SetPollingIntervalInSeconds(v int64) {
	o.PollingIntervalInSeconds = v
}

// GetCurrentSessionInfo returns the CurrentSessionInfo field value if set, zero value otherwise.
func (o *VcenterHvcLinksSyncProvidersInfo) GetCurrentSessionInfo() VcenterHvcLinksSyncProvidersSessionInfo {
	if o == nil || o.CurrentSessionInfo == nil {
		var ret VcenterHvcLinksSyncProvidersSessionInfo
		return ret
	}
	return *o.CurrentSessionInfo
}

// GetCurrentSessionInfoOk returns a tuple with the CurrentSessionInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) GetCurrentSessionInfoOk() (*VcenterHvcLinksSyncProvidersSessionInfo, bool) {
	if o == nil || o.CurrentSessionInfo == nil {
		return nil, false
	}
	return o.CurrentSessionInfo, true
}

// HasCurrentSessionInfo returns a boolean if a field has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) HasCurrentSessionInfo() bool {
	if o != nil && o.CurrentSessionInfo != nil {
		return true
	}

	return false
}

// SetCurrentSessionInfo gets a reference to the given VcenterHvcLinksSyncProvidersSessionInfo and assigns it to the CurrentSessionInfo field.
func (o *VcenterHvcLinksSyncProvidersInfo) SetCurrentSessionInfo(v VcenterHvcLinksSyncProvidersSessionInfo) {
	o.CurrentSessionInfo = &v
}

// GetStatusMessage returns the StatusMessage field value if set, zero value otherwise.
func (o *VcenterHvcLinksSyncProvidersInfo) GetStatusMessage() VapiStdLocalizableMessage {
	if o == nil || o.StatusMessage == nil {
		var ret VapiStdLocalizableMessage
		return ret
	}
	return *o.StatusMessage
}

// GetStatusMessageOk returns a tuple with the StatusMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) GetStatusMessageOk() (*VapiStdLocalizableMessage, bool) {
	if o == nil || o.StatusMessage == nil {
		return nil, false
	}
	return o.StatusMessage, true
}

// HasStatusMessage returns a boolean if a field has been set.
func (o *VcenterHvcLinksSyncProvidersInfo) HasStatusMessage() bool {
	if o != nil && o.StatusMessage != nil {
		return true
	}

	return false
}

// SetStatusMessage gets a reference to the given VapiStdLocalizableMessage and assigns it to the StatusMessage field.
func (o *VcenterHvcLinksSyncProvidersInfo) SetStatusMessage(v VapiStdLocalizableMessage) {
	o.StatusMessage = &v
}

func (o VcenterHvcLinksSyncProvidersInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LastSyncTime != nil {
		toSerialize["last_sync_time"] = o.LastSyncTime
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["polling_interval_in_seconds"] = o.PollingIntervalInSeconds
	}
	if o.CurrentSessionInfo != nil {
		toSerialize["current_session_info"] = o.CurrentSessionInfo
	}
	if o.StatusMessage != nil {
		toSerialize["status_message"] = o.StatusMessage
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterHvcLinksSyncProvidersInfo struct {
	value *VcenterHvcLinksSyncProvidersInfo
	isSet bool
}

func (v NullableVcenterHvcLinksSyncProvidersInfo) Get() *VcenterHvcLinksSyncProvidersInfo {
	return v.value
}

func (v *NullableVcenterHvcLinksSyncProvidersInfo) Set(val *VcenterHvcLinksSyncProvidersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterHvcLinksSyncProvidersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterHvcLinksSyncProvidersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterHvcLinksSyncProvidersInfo(val *VcenterHvcLinksSyncProvidersInfo) *NullableVcenterHvcLinksSyncProvidersInfo {
	return &NullableVcenterHvcLinksSyncProvidersInfo{value: val, isSet: true}
}

func (v NullableVcenterHvcLinksSyncProvidersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterHvcLinksSyncProvidersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


