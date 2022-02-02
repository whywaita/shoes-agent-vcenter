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

// VcenterContentRegistriesHarborProjectsInfo struct for VcenterContentRegistriesHarborProjectsInfo
type VcenterContentRegistriesHarborProjectsInfo struct {
	// Name of the Harbor project. Should be between 1-63 characters long alphanumeric string and may contain the following characters: a-z,0-9, and '-'. Must be starting with characters or numbers, with the '-' character allowed anywhere except the first or last character.
	Name string `json:"name"`
	ConfigStatus VcenterContentRegistriesHarborProjectsConfigStatus `json:"config_status"`
	Scope VcenterContentRegistriesHarborProjectsScope `json:"scope"`
	// The date and time when the harbor project creation API was triggered and project identifier generated.
	CreationTime time.Time `json:"creation_time"`
	// The date and time when the harbor project purge API was triggered. In case no purge was triggered, {@name #updateTime} is same as {@name #creationTime}.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// URL to access the harbor project through docker client.
	AccessUrl *string `json:"access_url,omitempty"`
	Message *VapiStdLocalizableMessage `json:"message,omitempty"`
}

// NewVcenterContentRegistriesHarborProjectsInfo instantiates a new VcenterContentRegistriesHarborProjectsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterContentRegistriesHarborProjectsInfo(name string, configStatus VcenterContentRegistriesHarborProjectsConfigStatus, scope VcenterContentRegistriesHarborProjectsScope, creationTime time.Time) *VcenterContentRegistriesHarborProjectsInfo {
	this := VcenterContentRegistriesHarborProjectsInfo{}
	this.Name = name
	this.ConfigStatus = configStatus
	this.Scope = scope
	this.CreationTime = creationTime
	return &this
}

// NewVcenterContentRegistriesHarborProjectsInfoWithDefaults instantiates a new VcenterContentRegistriesHarborProjectsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterContentRegistriesHarborProjectsInfoWithDefaults() *VcenterContentRegistriesHarborProjectsInfo {
	this := VcenterContentRegistriesHarborProjectsInfo{}
	return &this
}

// GetName returns the Name field value
func (o *VcenterContentRegistriesHarborProjectsInfo) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *VcenterContentRegistriesHarborProjectsInfo) SetName(v string) {
	o.Name = v
}

// GetConfigStatus returns the ConfigStatus field value
func (o *VcenterContentRegistriesHarborProjectsInfo) GetConfigStatus() VcenterContentRegistriesHarborProjectsConfigStatus {
	if o == nil {
		var ret VcenterContentRegistriesHarborProjectsConfigStatus
		return ret
	}

	return o.ConfigStatus
}

// GetConfigStatusOk returns a tuple with the ConfigStatus field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetConfigStatusOk() (*VcenterContentRegistriesHarborProjectsConfigStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigStatus, true
}

// SetConfigStatus sets field value
func (o *VcenterContentRegistriesHarborProjectsInfo) SetConfigStatus(v VcenterContentRegistriesHarborProjectsConfigStatus) {
	o.ConfigStatus = v
}

// GetScope returns the Scope field value
func (o *VcenterContentRegistriesHarborProjectsInfo) GetScope() VcenterContentRegistriesHarborProjectsScope {
	if o == nil {
		var ret VcenterContentRegistriesHarborProjectsScope
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetScopeOk() (*VcenterContentRegistriesHarborProjectsScope, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *VcenterContentRegistriesHarborProjectsInfo) SetScope(v VcenterContentRegistriesHarborProjectsScope) {
	o.Scope = v
}

// GetCreationTime returns the CreationTime field value
func (o *VcenterContentRegistriesHarborProjectsInfo) GetCreationTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CreationTime, true
}

// SetCreationTime sets field value
func (o *VcenterContentRegistriesHarborProjectsInfo) SetCreationTime(v time.Time) {
	o.CreationTime = v
}

// GetUpdateTime returns the UpdateTime field value if set, zero value otherwise.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetUpdateTime() time.Time {
	if o == nil || o.UpdateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdateTime
}

// GetUpdateTimeOk returns a tuple with the UpdateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetUpdateTimeOk() (*time.Time, bool) {
	if o == nil || o.UpdateTime == nil {
		return nil, false
	}
	return o.UpdateTime, true
}

// HasUpdateTime returns a boolean if a field has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) HasUpdateTime() bool {
	if o != nil && o.UpdateTime != nil {
		return true
	}

	return false
}

// SetUpdateTime gets a reference to the given time.Time and assigns it to the UpdateTime field.
func (o *VcenterContentRegistriesHarborProjectsInfo) SetUpdateTime(v time.Time) {
	o.UpdateTime = &v
}

// GetAccessUrl returns the AccessUrl field value if set, zero value otherwise.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetAccessUrl() string {
	if o == nil || o.AccessUrl == nil {
		var ret string
		return ret
	}
	return *o.AccessUrl
}

// GetAccessUrlOk returns a tuple with the AccessUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetAccessUrlOk() (*string, bool) {
	if o == nil || o.AccessUrl == nil {
		return nil, false
	}
	return o.AccessUrl, true
}

// HasAccessUrl returns a boolean if a field has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) HasAccessUrl() bool {
	if o != nil && o.AccessUrl != nil {
		return true
	}

	return false
}

// SetAccessUrl gets a reference to the given string and assigns it to the AccessUrl field.
func (o *VcenterContentRegistriesHarborProjectsInfo) SetAccessUrl(v string) {
	o.AccessUrl = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetMessage() VapiStdLocalizableMessage {
	if o == nil || o.Message == nil {
		var ret VapiStdLocalizableMessage
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) GetMessageOk() (*VapiStdLocalizableMessage, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *VcenterContentRegistriesHarborProjectsInfo) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given VapiStdLocalizableMessage and assigns it to the Message field.
func (o *VcenterContentRegistriesHarborProjectsInfo) SetMessage(v VapiStdLocalizableMessage) {
	o.Message = &v
}

func (o VcenterContentRegistriesHarborProjectsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["config_status"] = o.ConfigStatus
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["creation_time"] = o.CreationTime
	}
	if o.UpdateTime != nil {
		toSerialize["update_time"] = o.UpdateTime
	}
	if o.AccessUrl != nil {
		toSerialize["access_url"] = o.AccessUrl
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterContentRegistriesHarborProjectsInfo struct {
	value *VcenterContentRegistriesHarborProjectsInfo
	isSet bool
}

func (v NullableVcenterContentRegistriesHarborProjectsInfo) Get() *VcenterContentRegistriesHarborProjectsInfo {
	return v.value
}

func (v *NullableVcenterContentRegistriesHarborProjectsInfo) Set(val *VcenterContentRegistriesHarborProjectsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterContentRegistriesHarborProjectsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterContentRegistriesHarborProjectsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterContentRegistriesHarborProjectsInfo(val *VcenterContentRegistriesHarborProjectsInfo) *NullableVcenterContentRegistriesHarborProjectsInfo {
	return &NullableVcenterContentRegistriesHarborProjectsInfo{value: val, isSet: true}
}

func (v NullableVcenterContentRegistriesHarborProjectsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterContentRegistriesHarborProjectsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

