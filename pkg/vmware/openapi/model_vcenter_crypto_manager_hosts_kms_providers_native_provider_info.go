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

// VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo struct for VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo
type VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo struct {
	// Key identifier for the provider
	KeyId string `json:"key_id"`
}

// NewVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo instantiates a new VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo(keyId string) *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo {
	this := VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo{}
	this.KeyId = keyId
	return &this
}

// NewVcenterCryptoManagerHostsKmsProvidersNativeProviderInfoWithDefaults instantiates a new VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCryptoManagerHostsKmsProvidersNativeProviderInfoWithDefaults() *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo {
	this := VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo{}
	return &this
}

// GetKeyId returns the KeyId field value
func (o *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) GetKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) SetKeyId(v string) {
	o.KeyId = v
}

func (o VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo struct {
	value *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo
	isSet bool
}

func (v NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) Get() *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo {
	return v.value
}

func (v *NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) Set(val *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo(val *VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) *NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo {
	return &NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo{value: val, isSet: true}
}

func (v NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCryptoManagerHostsKmsProvidersNativeProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


