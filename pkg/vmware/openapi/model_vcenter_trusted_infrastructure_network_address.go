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

// VcenterTrustedInfrastructureNetworkAddress struct for VcenterTrustedInfrastructureNetworkAddress
type VcenterTrustedInfrastructureNetworkAddress struct {
	// The IP address or DNS resolvable name of the service.
	Hostname string `json:"hostname"`
	// The port of the service. If unset, port 443 will be used.
	Port *int64 `json:"port,omitempty"`
}

// NewVcenterTrustedInfrastructureNetworkAddress instantiates a new VcenterTrustedInfrastructureNetworkAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureNetworkAddress(hostname string) *VcenterTrustedInfrastructureNetworkAddress {
	this := VcenterTrustedInfrastructureNetworkAddress{}
	this.Hostname = hostname
	return &this
}

// NewVcenterTrustedInfrastructureNetworkAddressWithDefaults instantiates a new VcenterTrustedInfrastructureNetworkAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureNetworkAddressWithDefaults() *VcenterTrustedInfrastructureNetworkAddress {
	this := VcenterTrustedInfrastructureNetworkAddress{}
	return &this
}

// GetHostname returns the Hostname field value
func (o *VcenterTrustedInfrastructureNetworkAddress) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureNetworkAddress) GetHostnameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *VcenterTrustedInfrastructureNetworkAddress) SetHostname(v string) {
	o.Hostname = v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureNetworkAddress) GetPort() int64 {
	if o == nil || o.Port == nil {
		var ret int64
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureNetworkAddress) GetPortOk() (*int64, bool) {
	if o == nil || o.Port == nil {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureNetworkAddress) HasPort() bool {
	if o != nil && o.Port != nil {
		return true
	}

	return false
}

// SetPort gets a reference to the given int64 and assigns it to the Port field.
func (o *VcenterTrustedInfrastructureNetworkAddress) SetPort(v int64) {
	o.Port = &v
}

func (o VcenterTrustedInfrastructureNetworkAddress) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hostname"] = o.Hostname
	}
	if o.Port != nil {
		toSerialize["port"] = o.Port
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureNetworkAddress struct {
	value *VcenterTrustedInfrastructureNetworkAddress
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureNetworkAddress) Get() *VcenterTrustedInfrastructureNetworkAddress {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureNetworkAddress) Set(val *VcenterTrustedInfrastructureNetworkAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureNetworkAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureNetworkAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureNetworkAddress(val *VcenterTrustedInfrastructureNetworkAddress) *NullableVcenterTrustedInfrastructureNetworkAddress {
	return &NullableVcenterTrustedInfrastructureNetworkAddress{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureNetworkAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureNetworkAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

