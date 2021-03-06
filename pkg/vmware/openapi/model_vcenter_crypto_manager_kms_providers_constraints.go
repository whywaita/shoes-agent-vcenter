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

// VcenterCryptoManagerKmsProvidersConstraints struct for VcenterCryptoManagerKmsProvidersConstraints
type VcenterCryptoManagerKmsProvidersConstraints struct {
	// Determines if a provider is restricted to hosts with TPM 2.0 capability.
	TpmRequired bool `json:"tpm_required"`
}

// NewVcenterCryptoManagerKmsProvidersConstraints instantiates a new VcenterCryptoManagerKmsProvidersConstraints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterCryptoManagerKmsProvidersConstraints(tpmRequired bool) *VcenterCryptoManagerKmsProvidersConstraints {
	this := VcenterCryptoManagerKmsProvidersConstraints{}
	this.TpmRequired = tpmRequired
	return &this
}

// NewVcenterCryptoManagerKmsProvidersConstraintsWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersConstraints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterCryptoManagerKmsProvidersConstraintsWithDefaults() *VcenterCryptoManagerKmsProvidersConstraints {
	this := VcenterCryptoManagerKmsProvidersConstraints{}
	return &this
}

// GetTpmRequired returns the TpmRequired field value
func (o *VcenterCryptoManagerKmsProvidersConstraints) GetTpmRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.TpmRequired
}

// GetTpmRequiredOk returns a tuple with the TpmRequired field value
// and a boolean to check if the value has been set.
func (o *VcenterCryptoManagerKmsProvidersConstraints) GetTpmRequiredOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TpmRequired, true
}

// SetTpmRequired sets field value
func (o *VcenterCryptoManagerKmsProvidersConstraints) SetTpmRequired(v bool) {
	o.TpmRequired = v
}

func (o VcenterCryptoManagerKmsProvidersConstraints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["tpm_required"] = o.TpmRequired
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterCryptoManagerKmsProvidersConstraints struct {
	value *VcenterCryptoManagerKmsProvidersConstraints
	isSet bool
}

func (v NullableVcenterCryptoManagerKmsProvidersConstraints) Get() *VcenterCryptoManagerKmsProvidersConstraints {
	return v.value
}

func (v *NullableVcenterCryptoManagerKmsProvidersConstraints) Set(val *VcenterCryptoManagerKmsProvidersConstraints) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterCryptoManagerKmsProvidersConstraints) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterCryptoManagerKmsProvidersConstraints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterCryptoManagerKmsProvidersConstraints(val *VcenterCryptoManagerKmsProvidersConstraints) *NullableVcenterCryptoManagerKmsProvidersConstraints {
	return &NullableVcenterCryptoManagerKmsProvidersConstraints{value: val, isSet: true}
}

func (v NullableVcenterCryptoManagerKmsProvidersConstraints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterCryptoManagerKmsProvidersConstraints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


