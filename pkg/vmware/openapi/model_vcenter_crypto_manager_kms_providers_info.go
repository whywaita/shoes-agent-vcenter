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

// VcenterCryptoManagerKmsProvidersInfo struct for VcenterCryptoManagerKmsProvidersInfo
type VcenterCryptoManagerKmsProvidersInfo struct {
	Health VcenterCryptoManagerKmsProvidersHealth `json:"health"`
	// Details regarding the health status of the provider.   When the provider Providers.Health is not NONE or OK, this field will provide actionable descriptions of the issues. 
	Details []VapiStdLocalizableMessage `json:"details"`
	Constraints *VcenterCryptoManagerKmsProvidersConstraints `json:"constraints,omitempty"`
	Type VcenterCryptoManagerKmsProvidersType `json:"type"`
	NativeInfo *VcenterCryptoManagerKmsProvidersNativeProviderInfo `json:"native_info,omitempty"`
}

// NewVcenterCryptoManagerKmsProvidersInfo instantiates a new VcenterCryptoManagerKmsProvidersInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCryptoManagerKmsProvidersInfo(health VcenterCryptoManagerKmsProvidersHealth, details []VapiStdLocalizableMessage, type_ VcenterCryptoManagerKmsProvidersType) *VcenterCryptoManagerKmsProvidersInfo {
	this := VcenterCryptoManagerKmsProvidersInfo{}
	this.Health = health
	this.Details = details
	this.Type = type_
	return &this
}

// NewVcenterCryptoManagerKmsProvidersInfoWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCryptoManagerKmsProvidersInfoWithDefaults() *VcenterCryptoManagerKmsProvidersInfo {
	this := VcenterCryptoManagerKmsProvidersInfo{}
	return &this
}

// GetHealth returns the Health field value
func (o *VcenterCryptoManagerKmsProvidersInfo) GetHealth() VcenterCryptoManagerKmsProvidersHealth {
	if o == nil {
		var ret VcenterCryptoManagerKmsProvidersHealth
		return ret
	}

	return o.Health
}

// GetHealthOk returns a tuple with the Health field value
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetHealthOk() (*VcenterCryptoManagerKmsProvidersHealth, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Health, true
}

// SetHealth sets field value
func (o *VcenterCryptoManagerKmsProvidersInfo) SetHealth(v VcenterCryptoManagerKmsProvidersHealth) {
	o.Health = v
}

// GetDetails returns the Details field value
func (o *VcenterCryptoManagerKmsProvidersInfo) GetDetails() []VapiStdLocalizableMessage {
	if o == nil {
		var ret []VapiStdLocalizableMessage
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetDetailsOk() (*[]VapiStdLocalizableMessage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *VcenterCryptoManagerKmsProvidersInfo) SetDetails(v []VapiStdLocalizableMessage) {
	o.Details = v
}

// GetConstraints returns the Constraints field value if set, zero value otherwise.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetConstraints() VcenterCryptoManagerKmsProvidersConstraints {
	if o == nil || o.Constraints == nil {
		var ret VcenterCryptoManagerKmsProvidersConstraints
		return ret
	}
	return *o.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetConstraintsOk() (*VcenterCryptoManagerKmsProvidersConstraints, bool) {
	if o == nil || o.Constraints == nil {
		return nil, false
	}
	return o.Constraints, true
}

// HasConstraints returns a boolean if a field has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) HasConstraints() bool {
	if o != nil && o.Constraints != nil {
		return true
	}

	return false
}

// SetConstraints gets a reference to the given VcenterCryptoManagerKmsProvidersConstraints and assigns it to the Constraints field.
func (o *VcenterCryptoManagerKmsProvidersInfo) SetConstraints(v VcenterCryptoManagerKmsProvidersConstraints) {
	o.Constraints = &v
}

// GetType returns the Type field value
func (o *VcenterCryptoManagerKmsProvidersInfo) GetType() VcenterCryptoManagerKmsProvidersType {
	if o == nil {
		var ret VcenterCryptoManagerKmsProvidersType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetTypeOk() (*VcenterCryptoManagerKmsProvidersType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *VcenterCryptoManagerKmsProvidersInfo) SetType(v VcenterCryptoManagerKmsProvidersType) {
	o.Type = v
}

// GetNativeInfo returns the NativeInfo field value if set, zero value otherwise.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetNativeInfo() VcenterCryptoManagerKmsProvidersNativeProviderInfo {
	if o == nil || o.NativeInfo == nil {
		var ret VcenterCryptoManagerKmsProvidersNativeProviderInfo
		return ret
	}
	return *o.NativeInfo
}

// GetNativeInfoOk returns a tuple with the NativeInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) GetNativeInfoOk() (*VcenterCryptoManagerKmsProvidersNativeProviderInfo, bool) {
	if o == nil || o.NativeInfo == nil {
		return nil, false
	}
	return o.NativeInfo, true
}

// HasNativeInfo returns a boolean if a field has been set.
func (o *VcenterCryptoManagerKmsProvidersInfo) HasNativeInfo() bool {
	if o != nil && o.NativeInfo != nil {
		return true
	}

	return false
}

// SetNativeInfo gets a reference to the given VcenterCryptoManagerKmsProvidersNativeProviderInfo and assigns it to the NativeInfo field.
func (o *VcenterCryptoManagerKmsProvidersInfo) SetNativeInfo(v VcenterCryptoManagerKmsProvidersNativeProviderInfo) {
	o.NativeInfo = &v
}

func (o VcenterCryptoManagerKmsProvidersInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["health"] = o.Health
	}
	if true {
		toSerialize["details"] = o.Details
	}
	if o.Constraints != nil {
		toSerialize["constraints"] = o.Constraints
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.NativeInfo != nil {
		toSerialize["native_info"] = o.NativeInfo
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCryptoManagerKmsProvidersInfo struct {
	value *VcenterCryptoManagerKmsProvidersInfo
	isSet bool
}

func (v NullableVcenterCryptoManagerKmsProvidersInfo) Get() *VcenterCryptoManagerKmsProvidersInfo {
	return v.value
}

func (v *NullableVcenterCryptoManagerKmsProvidersInfo) Set(val *VcenterCryptoManagerKmsProvidersInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCryptoManagerKmsProvidersInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCryptoManagerKmsProvidersInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCryptoManagerKmsProvidersInfo(val *VcenterCryptoManagerKmsProvidersInfo) *NullableVcenterCryptoManagerKmsProvidersInfo {
	return &NullableVcenterCryptoManagerKmsProvidersInfo{value: val, isSet: true}
}

func (v NullableVcenterCryptoManagerKmsProvidersInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCryptoManagerKmsProvidersInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


