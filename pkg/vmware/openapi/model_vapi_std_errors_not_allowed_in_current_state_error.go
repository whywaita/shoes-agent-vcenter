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

// VapiStdErrorsNotAllowedInCurrentStateError struct for VapiStdErrorsNotAllowedInCurrentStateError
type VapiStdErrorsNotAllowedInCurrentStateError struct {
	Type *string `json:"type,omitempty"`
	Value *VapiStdErrorsNotAllowedInCurrentState `json:"value,omitempty"`
}

// NewVapiStdErrorsNotAllowedInCurrentStateError instantiates a new VapiStdErrorsNotAllowedInCurrentStateError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdErrorsNotAllowedInCurrentStateError() *VapiStdErrorsNotAllowedInCurrentStateError {
	this := VapiStdErrorsNotAllowedInCurrentStateError{}
	return &this
}

// NewVapiStdErrorsNotAllowedInCurrentStateErrorWithDefaults instantiates a new VapiStdErrorsNotAllowedInCurrentStateError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdErrorsNotAllowedInCurrentStateErrorWithDefaults() *VapiStdErrorsNotAllowedInCurrentStateError {
	this := VapiStdErrorsNotAllowedInCurrentStateError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) GetValue() VapiStdErrorsNotAllowedInCurrentState {
	if o == nil || o.Value == nil {
		var ret VapiStdErrorsNotAllowedInCurrentState
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) GetValueOk() (*VapiStdErrorsNotAllowedInCurrentState, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VapiStdErrorsNotAllowedInCurrentState and assigns it to the Value field.
func (o *VapiStdErrorsNotAllowedInCurrentStateError) SetValue(v VapiStdErrorsNotAllowedInCurrentState) {
	o.Value = &v
}

func (o VapiStdErrorsNotAllowedInCurrentStateError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdErrorsNotAllowedInCurrentStateError struct {
	value *VapiStdErrorsNotAllowedInCurrentStateError
	isSet bool
}

func (v NullableVapiStdErrorsNotAllowedInCurrentStateError) Get() *VapiStdErrorsNotAllowedInCurrentStateError {
	return v.value
}

func (v *NullableVapiStdErrorsNotAllowedInCurrentStateError) Set(val *VapiStdErrorsNotAllowedInCurrentStateError) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdErrorsNotAllowedInCurrentStateError) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdErrorsNotAllowedInCurrentStateError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdErrorsNotAllowedInCurrentStateError(val *VapiStdErrorsNotAllowedInCurrentStateError) *NullableVapiStdErrorsNotAllowedInCurrentStateError {
	return &NullableVapiStdErrorsNotAllowedInCurrentStateError{value: val, isSet: true}
}

func (v NullableVapiStdErrorsNotAllowedInCurrentStateError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdErrorsNotAllowedInCurrentStateError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


