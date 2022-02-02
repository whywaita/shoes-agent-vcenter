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

// VapiStdErrorsAlreadyExistsError struct for VapiStdErrorsAlreadyExistsError
type VapiStdErrorsAlreadyExistsError struct {
	Type *string `json:"type,omitempty"`
	Value *VapiStdErrorsAlreadyExists `json:"value,omitempty"`
}

// NewVapiStdErrorsAlreadyExistsError instantiates a new VapiStdErrorsAlreadyExistsError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdErrorsAlreadyExistsError() *VapiStdErrorsAlreadyExistsError {
	this := VapiStdErrorsAlreadyExistsError{}
	return &this
}

// NewVapiStdErrorsAlreadyExistsErrorWithDefaults instantiates a new VapiStdErrorsAlreadyExistsError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdErrorsAlreadyExistsErrorWithDefaults() *VapiStdErrorsAlreadyExistsError {
	this := VapiStdErrorsAlreadyExistsError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VapiStdErrorsAlreadyExistsError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsAlreadyExistsError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VapiStdErrorsAlreadyExistsError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VapiStdErrorsAlreadyExistsError) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VapiStdErrorsAlreadyExistsError) GetValue() VapiStdErrorsAlreadyExists {
	if o == nil || o.Value == nil {
		var ret VapiStdErrorsAlreadyExists
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsAlreadyExistsError) GetValueOk() (*VapiStdErrorsAlreadyExists, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VapiStdErrorsAlreadyExistsError) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VapiStdErrorsAlreadyExists and assigns it to the Value field.
func (o *VapiStdErrorsAlreadyExistsError) SetValue(v VapiStdErrorsAlreadyExists) {
	o.Value = &v
}

func (o VapiStdErrorsAlreadyExistsError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdErrorsAlreadyExistsError struct {
	value *VapiStdErrorsAlreadyExistsError
	isSet bool
}

func (v NullableVapiStdErrorsAlreadyExistsError) Get() *VapiStdErrorsAlreadyExistsError {
	return v.value
}

func (v *NullableVapiStdErrorsAlreadyExistsError) Set(val *VapiStdErrorsAlreadyExistsError) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdErrorsAlreadyExistsError) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdErrorsAlreadyExistsError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdErrorsAlreadyExistsError(val *VapiStdErrorsAlreadyExistsError) *NullableVapiStdErrorsAlreadyExistsError {
	return &NullableVapiStdErrorsAlreadyExistsError{value: val, isSet: true}
}

func (v NullableVapiStdErrorsAlreadyExistsError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdErrorsAlreadyExistsError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


