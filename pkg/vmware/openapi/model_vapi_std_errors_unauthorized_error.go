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

// VapiStdErrorsUnauthorizedError struct for VapiStdErrorsUnauthorizedError
type VapiStdErrorsUnauthorizedError struct {
	Type *string `json:"type,omitempty"`
	Value *VapiStdErrorsUnauthorized `json:"value,omitempty"`
}

// NewVapiStdErrorsUnauthorizedError instantiates a new VapiStdErrorsUnauthorizedError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdErrorsUnauthorizedError() *VapiStdErrorsUnauthorizedError {
	this := VapiStdErrorsUnauthorizedError{}
	return &this
}

// NewVapiStdErrorsUnauthorizedErrorWithDefaults instantiates a new VapiStdErrorsUnauthorizedError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdErrorsUnauthorizedErrorWithDefaults() *VapiStdErrorsUnauthorizedError {
	this := VapiStdErrorsUnauthorizedError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VapiStdErrorsUnauthorizedError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsUnauthorizedError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VapiStdErrorsUnauthorizedError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VapiStdErrorsUnauthorizedError) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VapiStdErrorsUnauthorizedError) GetValue() VapiStdErrorsUnauthorized {
	if o == nil || o.Value == nil {
		var ret VapiStdErrorsUnauthorized
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsUnauthorizedError) GetValueOk() (*VapiStdErrorsUnauthorized, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VapiStdErrorsUnauthorizedError) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VapiStdErrorsUnauthorized and assigns it to the Value field.
func (o *VapiStdErrorsUnauthorizedError) SetValue(v VapiStdErrorsUnauthorized) {
	o.Value = &v
}

func (o VapiStdErrorsUnauthorizedError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdErrorsUnauthorizedError struct {
	value *VapiStdErrorsUnauthorizedError
	isSet bool
}

func (v NullableVapiStdErrorsUnauthorizedError) Get() *VapiStdErrorsUnauthorizedError {
	return v.value
}

func (v *NullableVapiStdErrorsUnauthorizedError) Set(val *VapiStdErrorsUnauthorizedError) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdErrorsUnauthorizedError) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdErrorsUnauthorizedError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdErrorsUnauthorizedError(val *VapiStdErrorsUnauthorizedError) *NullableVapiStdErrorsUnauthorizedError {
	return &NullableVapiStdErrorsUnauthorizedError{value: val, isSet: true}
}

func (v NullableVapiStdErrorsUnauthorizedError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdErrorsUnauthorizedError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


