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

// VcenterNamespaceManagementSoftwareClustersMessage struct for VcenterNamespaceManagementSoftwareClustersMessage
type VcenterNamespaceManagementSoftwareClustersMessage struct {
	Severity VcenterNamespaceManagementSoftwareClustersMessageSeverity `json:"severity"`
	Details VapiStdLocalizableMessage `json:"details"`
}

// NewVcenterNamespaceManagementSoftwareClustersMessage instantiates a new VcenterNamespaceManagementSoftwareClustersMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementSoftwareClustersMessage(severity VcenterNamespaceManagementSoftwareClustersMessageSeverity, details VapiStdLocalizableMessage) *VcenterNamespaceManagementSoftwareClustersMessage {
	this := VcenterNamespaceManagementSoftwareClustersMessage{}
	this.Severity = severity
	this.Details = details
	return &this
}

// NewVcenterNamespaceManagementSoftwareClustersMessageWithDefaults instantiates a new VcenterNamespaceManagementSoftwareClustersMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementSoftwareClustersMessageWithDefaults() *VcenterNamespaceManagementSoftwareClustersMessage {
	this := VcenterNamespaceManagementSoftwareClustersMessage{}
	return &this
}

// GetSeverity returns the Severity field value
func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetSeverity() VcenterNamespaceManagementSoftwareClustersMessageSeverity {
	if o == nil {
		var ret VcenterNamespaceManagementSoftwareClustersMessageSeverity
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetSeverityOk() (*VcenterNamespaceManagementSoftwareClustersMessageSeverity, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *VcenterNamespaceManagementSoftwareClustersMessage) SetSeverity(v VcenterNamespaceManagementSoftwareClustersMessageSeverity) {
	o.Severity = v
}

// GetDetails returns the Details field value
func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetDetails() VapiStdLocalizableMessage {
	if o == nil {
		var ret VapiStdLocalizableMessage
		return ret
	}

	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementSoftwareClustersMessage) GetDetailsOk() (*VapiStdLocalizableMessage, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Details, true
}

// SetDetails sets field value
func (o *VcenterNamespaceManagementSoftwareClustersMessage) SetDetails(v VapiStdLocalizableMessage) {
	o.Details = v
}

func (o VcenterNamespaceManagementSoftwareClustersMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementSoftwareClustersMessage struct {
	value *VcenterNamespaceManagementSoftwareClustersMessage
	isSet bool
}

func (v NullableVcenterNamespaceManagementSoftwareClustersMessage) Get() *VcenterNamespaceManagementSoftwareClustersMessage {
	return v.value
}

func (v *NullableVcenterNamespaceManagementSoftwareClustersMessage) Set(val *VcenterNamespaceManagementSoftwareClustersMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementSoftwareClustersMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementSoftwareClustersMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementSoftwareClustersMessage(val *VcenterNamespaceManagementSoftwareClustersMessage) *NullableVcenterNamespaceManagementSoftwareClustersMessage {
	return &NullableVcenterNamespaceManagementSoftwareClustersMessage{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementSoftwareClustersMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementSoftwareClustersMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


