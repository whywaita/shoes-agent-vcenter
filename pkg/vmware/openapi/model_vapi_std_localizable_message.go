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

// VapiStdLocalizableMessage struct for VapiStdLocalizableMessage
type VapiStdLocalizableMessage struct {
	// Unique identifier of the localizable string or message template. <p> This identifier is typically used to retrieve a locale-specific string or message template from a message catalog.
	Id string `json:"id"`
	// The value of this localizable string or message template in the {@code en_US} (English) locale.  If {@link #id} refers to a message template, the default message will contain the substituted arguments. This value can be used by clients that do not need to display strings and messages in the native language of the user.  It could also be used as a fallback if a client is unable to access the appropriate message catalog.
	DefaultMessage string `json:"default_message"`
	// Positional arguments to be substituted into the message template. This list will be empty if the message uses named arguments or has no arguments.
	Args []string `json:"args"`
	// Named arguments to be substituted into the message template.
	Params *[]VapiStdLocalizableMessageParams `json:"params,omitempty"`
	// Localized string value as per request requirements.
	Localized *string `json:"localized,omitempty"`
}

// NewVapiStdLocalizableMessage instantiates a new VapiStdLocalizableMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdLocalizableMessage(id string, defaultMessage string, args []string) *VapiStdLocalizableMessage {
	this := VapiStdLocalizableMessage{}
	this.Id = id
	this.DefaultMessage = defaultMessage
	this.Args = args
	return &this
}

// NewVapiStdLocalizableMessageWithDefaults instantiates a new VapiStdLocalizableMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdLocalizableMessageWithDefaults() *VapiStdLocalizableMessage {
	this := VapiStdLocalizableMessage{}
	return &this
}

// GetId returns the Id field value
func (o *VapiStdLocalizableMessage) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizableMessage) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *VapiStdLocalizableMessage) SetId(v string) {
	o.Id = v
}

// GetDefaultMessage returns the DefaultMessage field value
func (o *VapiStdLocalizableMessage) GetDefaultMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultMessage
}

// GetDefaultMessageOk returns a tuple with the DefaultMessage field value
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizableMessage) GetDefaultMessageOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultMessage, true
}

// SetDefaultMessage sets field value
func (o *VapiStdLocalizableMessage) SetDefaultMessage(v string) {
	o.DefaultMessage = v
}

// GetArgs returns the Args field value
func (o *VapiStdLocalizableMessage) GetArgs() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Args
}

// GetArgsOk returns a tuple with the Args field value
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizableMessage) GetArgsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Args, true
}

// SetArgs sets field value
func (o *VapiStdLocalizableMessage) SetArgs(v []string) {
	o.Args = v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *VapiStdLocalizableMessage) GetParams() []VapiStdLocalizableMessageParams {
	if o == nil || o.Params == nil {
		var ret []VapiStdLocalizableMessageParams
		return ret
	}
	return *o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizableMessage) GetParamsOk() (*[]VapiStdLocalizableMessageParams, bool) {
	if o == nil || o.Params == nil {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *VapiStdLocalizableMessage) HasParams() bool {
	if o != nil && o.Params != nil {
		return true
	}

	return false
}

// SetParams gets a reference to the given []VapiStdLocalizableMessageParams and assigns it to the Params field.
func (o *VapiStdLocalizableMessage) SetParams(v []VapiStdLocalizableMessageParams) {
	o.Params = &v
}

// GetLocalized returns the Localized field value if set, zero value otherwise.
func (o *VapiStdLocalizableMessage) GetLocalized() string {
	if o == nil || o.Localized == nil {
		var ret string
		return ret
	}
	return *o.Localized
}

// GetLocalizedOk returns a tuple with the Localized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdLocalizableMessage) GetLocalizedOk() (*string, bool) {
	if o == nil || o.Localized == nil {
		return nil, false
	}
	return o.Localized, true
}

// HasLocalized returns a boolean if a field has been set.
func (o *VapiStdLocalizableMessage) HasLocalized() bool {
	if o != nil && o.Localized != nil {
		return true
	}

	return false
}

// SetLocalized gets a reference to the given string and assigns it to the Localized field.
func (o *VapiStdLocalizableMessage) SetLocalized(v string) {
	o.Localized = &v
}

func (o VapiStdLocalizableMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["default_message"] = o.DefaultMessage
	}
	if true {
		toSerialize["args"] = o.Args
	}
	if o.Params != nil {
		toSerialize["params"] = o.Params
	}
	if o.Localized != nil {
		toSerialize["localized"] = o.Localized
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdLocalizableMessage struct {
	value *VapiStdLocalizableMessage
	isSet bool
}

func (v NullableVapiStdLocalizableMessage) Get() *VapiStdLocalizableMessage {
	return v.value
}

func (v *NullableVapiStdLocalizableMessage) Set(val *VapiStdLocalizableMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdLocalizableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdLocalizableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdLocalizableMessage(val *VapiStdLocalizableMessage) *NullableVapiStdLocalizableMessage {
	return &NullableVapiStdLocalizableMessage{value: val, isSet: true}
}

func (v NullableVapiStdLocalizableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdLocalizableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


