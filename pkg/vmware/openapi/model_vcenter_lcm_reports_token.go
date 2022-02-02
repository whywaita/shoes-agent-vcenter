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

// VcenterLcmReportsToken struct for VcenterLcmReportsToken
type VcenterLcmReportsToken struct {
	// A one-time, short-lived token required in the HTTP header of the request to the url. This token needs to be passed in as a header with the name \"session-id\".
	Token string `json:"token"`
	// Expiry time of the token
	Expiry time.Time `json:"expiry"`
}

// NewVcenterLcmReportsToken instantiates a new VcenterLcmReportsToken object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterLcmReportsToken(token string, expiry time.Time) *VcenterLcmReportsToken {
	this := VcenterLcmReportsToken{}
	this.Token = token
	this.Expiry = expiry
	return &this
}

// NewVcenterLcmReportsTokenWithDefaults instantiates a new VcenterLcmReportsToken object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterLcmReportsTokenWithDefaults() *VcenterLcmReportsToken {
	this := VcenterLcmReportsToken{}
	return &this
}

// GetToken returns the Token field value
func (o *VcenterLcmReportsToken) GetToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Token
}

// GetTokenOk returns a tuple with the Token field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmReportsToken) GetTokenOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Token, true
}

// SetToken sets field value
func (o *VcenterLcmReportsToken) SetToken(v string) {
	o.Token = v
}

// GetExpiry returns the Expiry field value
func (o *VcenterLcmReportsToken) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *VcenterLcmReportsToken) GetExpiryOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *VcenterLcmReportsToken) SetExpiry(v time.Time) {
	o.Expiry = v
}

func (o VcenterLcmReportsToken) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["token"] = o.Token
	}
	if true {
		toSerialize["expiry"] = o.Expiry
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterLcmReportsToken struct {
	value *VcenterLcmReportsToken
	isSet bool
}

func (v NullableVcenterLcmReportsToken) Get() *VcenterLcmReportsToken {
	return v.value
}

func (v *NullableVcenterLcmReportsToken) Set(val *VcenterLcmReportsToken) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterLcmReportsToken) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterLcmReportsToken) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterLcmReportsToken(val *VcenterLcmReportsToken) *NullableVcenterLcmReportsToken {
	return &NullableVcenterLcmReportsToken{value: val, isSet: true}
}

func (v NullableVcenterLcmReportsToken) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterLcmReportsToken) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


