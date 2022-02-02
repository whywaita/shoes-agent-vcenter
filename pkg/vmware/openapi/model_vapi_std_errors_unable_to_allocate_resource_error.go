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

// VapiStdErrorsUnableToAllocateResourceError struct for VapiStdErrorsUnableToAllocateResourceError
type VapiStdErrorsUnableToAllocateResourceError struct {
	Type *string `json:"type,omitempty"`
	Value *VapiStdErrorsUnableToAllocateResource `json:"value,omitempty"`
}

// NewVapiStdErrorsUnableToAllocateResourceError instantiates a new VapiStdErrorsUnableToAllocateResourceError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdErrorsUnableToAllocateResourceError() *VapiStdErrorsUnableToAllocateResourceError {
	this := VapiStdErrorsUnableToAllocateResourceError{}
	return &this
}

// NewVapiStdErrorsUnableToAllocateResourceErrorWithDefaults instantiates a new VapiStdErrorsUnableToAllocateResourceError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdErrorsUnableToAllocateResourceErrorWithDefaults() *VapiStdErrorsUnableToAllocateResourceError {
	this := VapiStdErrorsUnableToAllocateResourceError{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VapiStdErrorsUnableToAllocateResourceError) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsUnableToAllocateResourceError) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VapiStdErrorsUnableToAllocateResourceError) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VapiStdErrorsUnableToAllocateResourceError) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *VapiStdErrorsUnableToAllocateResourceError) GetValue() VapiStdErrorsUnableToAllocateResource {
	if o == nil || o.Value == nil {
		var ret VapiStdErrorsUnableToAllocateResource
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsUnableToAllocateResourceError) GetValueOk() (*VapiStdErrorsUnableToAllocateResource, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *VapiStdErrorsUnableToAllocateResourceError) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given VapiStdErrorsUnableToAllocateResource and assigns it to the Value field.
func (o *VapiStdErrorsUnableToAllocateResourceError) SetValue(v VapiStdErrorsUnableToAllocateResource) {
	o.Value = &v
}

func (o VapiStdErrorsUnableToAllocateResourceError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdErrorsUnableToAllocateResourceError struct {
	value *VapiStdErrorsUnableToAllocateResourceError
	isSet bool
}

func (v NullableVapiStdErrorsUnableToAllocateResourceError) Get() *VapiStdErrorsUnableToAllocateResourceError {
	return v.value
}

func (v *NullableVapiStdErrorsUnableToAllocateResourceError) Set(val *VapiStdErrorsUnableToAllocateResourceError) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdErrorsUnableToAllocateResourceError) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdErrorsUnableToAllocateResourceError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdErrorsUnableToAllocateResourceError(val *VapiStdErrorsUnableToAllocateResourceError) *NullableVapiStdErrorsUnableToAllocateResourceError {
	return &NullableVapiStdErrorsUnableToAllocateResourceError{value: val, isSet: true}
}

func (v NullableVapiStdErrorsUnableToAllocateResourceError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdErrorsUnableToAllocateResourceError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

