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

// VcenterServicesServiceInfo struct for VcenterServicesServiceInfo
type VcenterServicesServiceInfo struct {
	// Service name key. Can be used to lookup resource bundle
	NameKey string `json:"name_key"`
	// Service description key. Can be used to lookup resource bundle
	DescriptionKey string `json:"description_key"`
	StartupType VcenterServicesServiceStartupType `json:"startup_type"`
	State VcenterServicesServiceState `json:"state"`
	Health *VcenterServicesServiceHealth `json:"health,omitempty"`
	// Localizable messages associated with the health of the service This field is optional and it is only relevant when the value of Service.Info.state is STARTED.
	HealthMessages *[]VapiStdLocalizableMessage `json:"health_messages,omitempty"`
}

// NewVcenterServicesServiceInfo instantiates a new VcenterServicesServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterServicesServiceInfo(nameKey string, descriptionKey string, startupType VcenterServicesServiceStartupType, state VcenterServicesServiceState) *VcenterServicesServiceInfo {
	this := VcenterServicesServiceInfo{}
	this.NameKey = nameKey
	this.DescriptionKey = descriptionKey
	this.StartupType = startupType
	this.State = state
	return &this
}

// NewVcenterServicesServiceInfoWithDefaults instantiates a new VcenterServicesServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterServicesServiceInfoWithDefaults() *VcenterServicesServiceInfo {
	this := VcenterServicesServiceInfo{}
	return &this
}

// GetNameKey returns the NameKey field value
func (o *VcenterServicesServiceInfo) GetNameKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NameKey
}

// GetNameKeyOk returns a tuple with the NameKey field value
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceInfo) GetNameKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NameKey, true
}

// SetNameKey sets field value
func (o *VcenterServicesServiceInfo) SetNameKey(v string) {
	o.NameKey = v
}

// GetDescriptionKey returns the DescriptionKey field value
func (o *VcenterServicesServiceInfo) GetDescriptionKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DescriptionKey
}

// GetDescriptionKeyOk returns a tuple with the DescriptionKey field value
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceInfo) GetDescriptionKeyOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DescriptionKey, true
}

// SetDescriptionKey sets field value
func (o *VcenterServicesServiceInfo) SetDescriptionKey(v string) {
	o.DescriptionKey = v
}

// GetStartupType returns the StartupType field value
func (o *VcenterServicesServiceInfo) GetStartupType() VcenterServicesServiceStartupType {
	if o == nil {
		var ret VcenterServicesServiceStartupType
		return ret
	}

	return o.StartupType
}

// GetStartupTypeOk returns a tuple with the StartupType field value
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceInfo) GetStartupTypeOk() (*VcenterServicesServiceStartupType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.StartupType, true
}

// SetStartupType sets field value
func (o *VcenterServicesServiceInfo) SetStartupType(v VcenterServicesServiceStartupType) {
	o.StartupType = v
}

// GetState returns the State field value
func (o *VcenterServicesServiceInfo) GetState() VcenterServicesServiceState {
	if o == nil {
		var ret VcenterServicesServiceState
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceInfo) GetStateOk() (*VcenterServicesServiceState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VcenterServicesServiceInfo) SetState(v VcenterServicesServiceState) {
	o.State = v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *VcenterServicesServiceInfo) GetHealth() VcenterServicesServiceHealth {
	if o == nil || o.Health == nil {
		var ret VcenterServicesServiceHealth
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceInfo) GetHealthOk() (*VcenterServicesServiceHealth, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *VcenterServicesServiceInfo) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given VcenterServicesServiceHealth and assigns it to the Health field.
func (o *VcenterServicesServiceInfo) SetHealth(v VcenterServicesServiceHealth) {
	o.Health = &v
}

// GetHealthMessages returns the HealthMessages field value if set, zero value otherwise.
func (o *VcenterServicesServiceInfo) GetHealthMessages() []VapiStdLocalizableMessage {
	if o == nil || o.HealthMessages == nil {
		var ret []VapiStdLocalizableMessage
		return ret
	}
	return *o.HealthMessages
}

// GetHealthMessagesOk returns a tuple with the HealthMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterServicesServiceInfo) GetHealthMessagesOk() (*[]VapiStdLocalizableMessage, bool) {
	if o == nil || o.HealthMessages == nil {
		return nil, false
	}
	return o.HealthMessages, true
}

// HasHealthMessages returns a boolean if a field has been set.
func (o *VcenterServicesServiceInfo) HasHealthMessages() bool {
	if o != nil && o.HealthMessages != nil {
		return true
	}

	return false
}

// SetHealthMessages gets a reference to the given []VapiStdLocalizableMessage and assigns it to the HealthMessages field.
func (o *VcenterServicesServiceInfo) SetHealthMessages(v []VapiStdLocalizableMessage) {
	o.HealthMessages = &v
}

func (o VcenterServicesServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name_key"] = o.NameKey
	}
	if true {
		toSerialize["description_key"] = o.DescriptionKey
	}
	if true {
		toSerialize["startup_type"] = o.StartupType
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.Health != nil {
		toSerialize["health"] = o.Health
	}
	if o.HealthMessages != nil {
		toSerialize["health_messages"] = o.HealthMessages
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterServicesServiceInfo struct {
	value *VcenterServicesServiceInfo
	isSet bool
}

func (v NullableVcenterServicesServiceInfo) Get() *VcenterServicesServiceInfo {
	return v.value
}

func (v *NullableVcenterServicesServiceInfo) Set(val *VcenterServicesServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterServicesServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterServicesServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterServicesServiceInfo(val *VcenterServicesServiceInfo) *NullableVcenterServicesServiceInfo {
	return &NullableVcenterServicesServiceInfo{value: val, isSet: true}
}

func (v NullableVcenterServicesServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterServicesServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


