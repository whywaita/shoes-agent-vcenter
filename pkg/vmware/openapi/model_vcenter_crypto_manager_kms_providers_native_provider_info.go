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

// VcenterCryptoManagerKmsProvidersNativeProviderInfo struct for VcenterCryptoManagerKmsProvidersNativeProviderInfo
type VcenterCryptoManagerKmsProvidersNativeProviderInfo struct {
	// Key identifier for the provider
	KeyId string `json:"key_id"`
}

// NewVcenterCryptoManagerKmsProvidersNativeProviderInfo instantiates a new VcenterCryptoManagerKmsProvidersNativeProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCryptoManagerKmsProvidersNativeProviderInfo(keyId string) *VcenterCryptoManagerKmsProvidersNativeProviderInfo {
	this := VcenterCryptoManagerKmsProvidersNativeProviderInfo{}
	this.KeyId = keyId
	return &this
}

// NewVcenterCryptoManagerKmsProvidersNativeProviderInfoWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersNativeProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCryptoManagerKmsProvidersNativeProviderInfoWithDefaults() *VcenterCryptoManagerKmsProvidersNativeProviderInfo {
	this := VcenterCryptoManagerKmsProvidersNativeProviderInfo{}
	return &this
}

// GetKeyId returns the KeyId field value
func (o *VcenterCryptoManagerKmsProvidersNativeProviderInfo) GetKeyId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersNativeProviderInfo) GetKeyIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *VcenterCryptoManagerKmsProvidersNativeProviderInfo) SetKeyId(v string) {
	o.KeyId = v
}

func (o VcenterCryptoManagerKmsProvidersNativeProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["key_id"] = o.KeyId
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo struct {
	value *VcenterCryptoManagerKmsProvidersNativeProviderInfo
	isSet bool
}

func (v NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo) Get() *VcenterCryptoManagerKmsProvidersNativeProviderInfo {
	return v.value
}

func (v *NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo) Set(val *VcenterCryptoManagerKmsProvidersNativeProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCryptoManagerKmsProvidersNativeProviderInfo(val *VcenterCryptoManagerKmsProvidersNativeProviderInfo) *NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo {
	return &NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo{value: val, isSet: true}
}

func (v NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCryptoManagerKmsProvidersNativeProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

