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

// VcenterTokenserviceInvalidScope struct for VcenterTokenserviceInvalidScope
type VcenterTokenserviceInvalidScope struct {
	// Stack of one or more localizable messages for human {@term error} consumers. <p> The message at the top of the stack (first in the list) describes the {@term error} from the perspective of the {@term operation} the client invoked. Each subsequent message in the stack describes the \"cause\" of the prior message.
	Messages []VapiStdLocalizableMessage `json:"messages"`
	// Data to facilitate clients responding to the {@term operation} reporting a standard {@term error} to indicating that it was unable to complete successfully. <p> {@term Operations} may provide data that clients can use when responding to {@term errors}.  Since the data that clients need may be specific to the context of the {@term operation} reporting the {@term error}, different {@term operations} that report the same {@term error} may provide different data in the {@term error}.  The documentation for each each {@term operation} will describe what, if any, data it provides for each {@term error} it reports. The {@link ArgumentLocations}, {@link FileLocations}, and {@link TransientIndication} {@term structures} are intended as possible values for this {@term field}.  {@link vapi.std.DynamicID} may also be useful as a value for this {@term field} (although that is not its primary purpose).  Some {@term services} may provide their own specific {@term structures} for use as the value of this {@term field} when reporting {@term errors} from their {@term operations}.
	Data *map[string]interface{} `json:"data,omitempty"`
}

// NewVcenterTokenserviceInvalidScope instantiates a new VcenterTokenserviceInvalidScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTokenserviceInvalidScope(messages []VapiStdLocalizableMessage) *VcenterTokenserviceInvalidScope {
	this := VcenterTokenserviceInvalidScope{}
	this.Messages = messages
	return &this
}

// NewVcenterTokenserviceInvalidScopeWithDefaults instantiates a new VcenterTokenserviceInvalidScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTokenserviceInvalidScopeWithDefaults() *VcenterTokenserviceInvalidScope {
	this := VcenterTokenserviceInvalidScope{}
	return &this
}

// GetMessages returns the Messages field value
func (o *VcenterTokenserviceInvalidScope) GetMessages() []VapiStdLocalizableMessage {
	if o == nil {
		var ret []VapiStdLocalizableMessage
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *VcenterTokenserviceInvalidScope) GetMessagesOk() (*[]VapiStdLocalizableMessage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *VcenterTokenserviceInvalidScope) SetMessages(v []VapiStdLocalizableMessage) {
	o.Messages = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *VcenterTokenserviceInvalidScope) GetData() map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTokenserviceInvalidScope) GetDataOk() (*map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *VcenterTokenserviceInvalidScope) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]interface{} and assigns it to the Data field.
func (o *VcenterTokenserviceInvalidScope) SetData(v map[string]interface{}) {
	o.Data = &v
}

func (o VcenterTokenserviceInvalidScope) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["messages"] = o.Messages
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTokenserviceInvalidScope struct {
	value *VcenterTokenserviceInvalidScope
	isSet bool
}

func (v NullableVcenterTokenserviceInvalidScope) Get() *VcenterTokenserviceInvalidScope {
	return v.value
}

func (v *NullableVcenterTokenserviceInvalidScope) Set(val *VcenterTokenserviceInvalidScope) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTokenserviceInvalidScope) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTokenserviceInvalidScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTokenserviceInvalidScope(val *VcenterTokenserviceInvalidScope) *NullableVcenterTokenserviceInvalidScope {
	return &NullableVcenterTokenserviceInvalidScope{value: val, isSet: true}
}

func (v NullableVcenterTokenserviceInvalidScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTokenserviceInvalidScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


