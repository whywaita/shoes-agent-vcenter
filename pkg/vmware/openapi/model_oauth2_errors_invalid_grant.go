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

// Oauth2ErrorsInvalidGrant struct for Oauth2ErrorsInvalidGrant
type Oauth2ErrorsInvalidGrant struct {
	Error Oauth2ErrorsErrorType `json:"error"`
	// Human-readable ASCII text providing additional information, used to assist the client developer in understanding the error that occurred. Values for the \"error_description\" parameter MUST NOT include characters outside the set %x20-21 / %x23-5B / %x5D-7E. if no additional information is available.
	ErrorDescription *string `json:"error_description,omitempty"`
	// A URI identifying a human-readable web page with information about the error, used to provide the client developer with additional information about the error. if no such web-page is available.
	ErrorUri *string `json:"error_uri,omitempty"`
}

// NewOauth2ErrorsInvalidGrant instantiates a new Oauth2ErrorsInvalidGrant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOauth2ErrorsInvalidGrant(error_ Oauth2ErrorsErrorType) *Oauth2ErrorsInvalidGrant {
	this := Oauth2ErrorsInvalidGrant{}
	this.Error = error_
	return &this
}

// NewOauth2ErrorsInvalidGrantWithDefaults instantiates a new Oauth2ErrorsInvalidGrant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOauth2ErrorsInvalidGrantWithDefaults() *Oauth2ErrorsInvalidGrant {
	this := Oauth2ErrorsInvalidGrant{}
	return &this
}

// GetError returns the Error field value
func (o *Oauth2ErrorsInvalidGrant) GetError() Oauth2ErrorsErrorType {
	if o == nil {
		var ret Oauth2ErrorsErrorType
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *Oauth2ErrorsInvalidGrant) GetErrorOk() (*Oauth2ErrorsErrorType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *Oauth2ErrorsInvalidGrant) SetError(v Oauth2ErrorsErrorType) {
	o.Error = v
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *Oauth2ErrorsInvalidGrant) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2ErrorsInvalidGrant) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *Oauth2ErrorsInvalidGrant) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *Oauth2ErrorsInvalidGrant) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetErrorUri returns the ErrorUri field value if set, zero value otherwise.
func (o *Oauth2ErrorsInvalidGrant) GetErrorUri() string {
	if o == nil || o.ErrorUri == nil {
		var ret string
		return ret
	}
	return *o.ErrorUri
}

// GetErrorUriOk returns a tuple with the ErrorUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Oauth2ErrorsInvalidGrant) GetErrorUriOk() (*string, bool) {
	if o == nil || o.ErrorUri == nil {
		return nil, false
	}
	return o.ErrorUri, true
}

// HasErrorUri returns a boolean if a field has been set.
func (o *Oauth2ErrorsInvalidGrant) HasErrorUri() bool {
	if o != nil && o.ErrorUri != nil {
		return true
	}

	return false
}

// SetErrorUri gets a reference to the given string and assigns it to the ErrorUri field.
func (o *Oauth2ErrorsInvalidGrant) SetErrorUri(v string) {
	o.ErrorUri = &v
}

func (o Oauth2ErrorsInvalidGrant) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["error"] = o.Error
	}
	if o.ErrorDescription != nil {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.ErrorUri != nil {
		toSerialize["error_uri"] = o.ErrorUri
	}
	return json.Marshal(toSerialize)
}

type NullableOauth2ErrorsInvalidGrant struct {
	value *Oauth2ErrorsInvalidGrant
	isSet bool
}

func (v NullableOauth2ErrorsInvalidGrant) Get() *Oauth2ErrorsInvalidGrant {
	return v.value
}

func (v *NullableOauth2ErrorsInvalidGrant) Set(val *Oauth2ErrorsInvalidGrant) {
	v.value = val
	v.isSet = true
}

func (v NullableOauth2ErrorsInvalidGrant) IsSet() bool {
	return v.isSet
}

func (v *NullableOauth2ErrorsInvalidGrant) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOauth2ErrorsInvalidGrant(val *Oauth2ErrorsInvalidGrant) *NullableOauth2ErrorsInvalidGrant {
	return &NullableOauth2ErrorsInvalidGrant{value: val, isSet: true}
}

func (v NullableOauth2ErrorsInvalidGrant) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOauth2ErrorsInvalidGrant) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


