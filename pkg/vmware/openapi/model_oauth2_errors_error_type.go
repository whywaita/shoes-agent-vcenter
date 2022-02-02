/*
vcenter

VMware vCenter Server provides a centralized platform for managing your VMware vSphere environments

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// Oauth2ErrorsErrorType Enumeration of OAuth 2.0 errors.
type Oauth2ErrorsErrorType string

// List of Oauth2ErrorsErrorType
const (
	OAUTH2ERRORSERRORTYPE_REQUEST Oauth2ErrorsErrorType = "invalid_request"
	OAUTH2ERRORSERRORTYPE_SCOPE Oauth2ErrorsErrorType = "invalid_scope"
	OAUTH2ERRORSERRORTYPE_GRANT Oauth2ErrorsErrorType = "invalid_grant"
)

// All allowed values of Oauth2ErrorsErrorType enum
var AllowedOauth2ErrorsErrorTypeEnumValues = []Oauth2ErrorsErrorType{
	"invalid_request",
	"invalid_scope",
	"invalid_grant",
}

func (v *Oauth2ErrorsErrorType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Oauth2ErrorsErrorType(value)
	for _, existing := range AllowedOauth2ErrorsErrorTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Oauth2ErrorsErrorType", value)
}

// NewOauth2ErrorsErrorTypeFromValue returns a pointer to a valid Oauth2ErrorsErrorType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOauth2ErrorsErrorTypeFromValue(v string) (*Oauth2ErrorsErrorType, error) {
	ev := Oauth2ErrorsErrorType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Oauth2ErrorsErrorType: valid values are %v", v, AllowedOauth2ErrorsErrorTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Oauth2ErrorsErrorType) IsValid() bool {
	for _, existing := range AllowedOauth2ErrorsErrorTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Oauth2ErrorsErrorType value
func (v Oauth2ErrorsErrorType) Ptr() *Oauth2ErrorsErrorType {
	return &v
}

type NullableOauth2ErrorsErrorType struct {
	value *Oauth2ErrorsErrorType
	isSet bool
}

func (v NullableOauth2ErrorsErrorType) Get() *Oauth2ErrorsErrorType {
	return v.value
}

func (v *NullableOauth2ErrorsErrorType) Set(val *Oauth2ErrorsErrorType) {
	v.value = val
	v.isSet = true
}

func (v NullableOauth2ErrorsErrorType) IsSet() bool {
	return v.isSet
}

func (v *NullableOauth2ErrorsErrorType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauth2ErrorsErrorType(val *Oauth2ErrorsErrorType) *NullableOauth2ErrorsErrorType {
	return &NullableOauth2ErrorsErrorType{value: val, isSet: true}
}

func (v NullableOauth2ErrorsErrorType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauth2ErrorsErrorType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

