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

// VapiStdErrorsNotAllowedInCurrentState struct for VapiStdErrorsNotAllowedInCurrentState
type VapiStdErrorsNotAllowedInCurrentState struct {
	// Stack of one or more localizable messages for human {@term error} consumers. <p> The message at the top of the stack (first in the list) describes the {@term error} from the perspective of the {@term operation} the client invoked. Each subsequent message in the stack describes the \"cause\" of the prior message.
	Messages []VapiStdLocalizableMessage `json:"messages"`
	// Data to facilitate clients responding to the {@term operation} reporting a standard {@term error} to indicating that it was unable to complete successfully. <p> {@term Operations} may provide data that clients can use when responding to {@term errors}.  Since the data that clients need may be specific to the context of the {@term operation} reporting the {@term error}, different {@term operations} that report the same {@term error} may provide different data in the {@term error}.  The documentation for each each {@term operation} will describe what, if any, data it provides for each {@term error} it reports. The {@link ArgumentLocations}, {@link FileLocations}, and {@link TransientIndication} {@term structures} are intended as possible values for this {@term field}.  {@link vapi.std.DynamicID} may also be useful as a value for this {@term field} (although that is not its primary purpose).  Some {@term services} may provide their own specific {@term structures} for use as the value of this {@term field} when reporting {@term errors} from their {@term operations}.
	Data *map[string]interface{} `json:"data,omitempty"`
	ErrorType *VapiStdErrorsErrorType `json:"error_type,omitempty"`
}

// NewVapiStdErrorsNotAllowedInCurrentState instantiates a new VapiStdErrorsNotAllowedInCurrentState object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVapiStdErrorsNotAllowedInCurrentState(messages []VapiStdLocalizableMessage) *VapiStdErrorsNotAllowedInCurrentState {
	this := VapiStdErrorsNotAllowedInCurrentState{}
	this.Messages = messages
	return &this
}

// NewVapiStdErrorsNotAllowedInCurrentStateWithDefaults instantiates a new VapiStdErrorsNotAllowedInCurrentState object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVapiStdErrorsNotAllowedInCurrentStateWithDefaults() *VapiStdErrorsNotAllowedInCurrentState {
	this := VapiStdErrorsNotAllowedInCurrentState{}
	return &this
}

// GetMessages returns the Messages field value
func (o *VapiStdErrorsNotAllowedInCurrentState) GetMessages() []VapiStdLocalizableMessage {
	if o == nil {
		var ret []VapiStdLocalizableMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsNotAllowedInCurrentState) GetMessagesOk() (*[]VapiStdLocalizableMessage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *VapiStdErrorsNotAllowedInCurrentState) SetMessages(v []VapiStdLocalizableMessage) {
	o.Messages = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VapiStdErrorsNotAllowedInCurrentState) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsNotAllowedInCurrentState) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VapiStdErrorsNotAllowedInCurrentState) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *VapiStdErrorsNotAllowedInCurrentState) SetData(v map[string]interface{}) {
	o.Data = &v
}

// GetErrorType returns the ErrorType field value if set, zero value otherwise.
func (o *VapiStdErrorsNotAllowedInCurrentState) GetErrorType() VapiStdErrorsErrorType {
	if o == nil || o.ErrorType == nil {
		var ret VapiStdErrorsErrorType
		return ret
	}
	return *o.ErrorType
}

// GetErrorTypeOk returns a tuple with the ErrorType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VapiStdErrorsNotAllowedInCurrentState) GetErrorTypeOk() (*VapiStdErrorsErrorType, bool) {
	if o == nil || o.ErrorType == nil {
		return nil, false
	}
	return o.ErrorType, true
}

// HasErrorType returns a boolean if a field has been set.
func (o *VapiStdErrorsNotAllowedInCurrentState) HasErrorType() bool {
	if o != nil && o.ErrorType != nil {
		return true
	}

	return false
}

// SetErrorType gets a reference to the given VapiStdErrorsErrorType and assigns it to the ErrorType field.
func (o *VapiStdErrorsNotAllowedInCurrentState) SetErrorType(v VapiStdErrorsErrorType) {
	o.ErrorType = &v
}

func (o VapiStdErrorsNotAllowedInCurrentState) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["messages"] = o.Messages
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.ErrorType != nil {
		toSerialize["error_type"] = o.ErrorType
	}
	return json.Marshal(toSerialize)
}

type NullableVapiStdErrorsNotAllowedInCurrentState struct {
	value *VapiStdErrorsNotAllowedInCurrentState
	isSet bool
}

func (v NullableVapiStdErrorsNotAllowedInCurrentState) Get() *VapiStdErrorsNotAllowedInCurrentState {
	return v.value
}

func (v *NullableVapiStdErrorsNotAllowedInCurrentState) Set(val *VapiStdErrorsNotAllowedInCurrentState) {
	v.value = val
	v.isSet = true
}

func (v NullableVapiStdErrorsNotAllowedInCurrentState) IsSet() bool {
	return v.isSet
}

func (v *NullableVapiStdErrorsNotAllowedInCurrentState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVapiStdErrorsNotAllowedInCurrentState(val *VapiStdErrorsNotAllowedInCurrentState) *NullableVapiStdErrorsNotAllowedInCurrentState {
	return &NullableVapiStdErrorsNotAllowedInCurrentState{value: val, isSet: true}
}

func (v NullableVapiStdErrorsNotAllowedInCurrentState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVapiStdErrorsNotAllowedInCurrentState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

