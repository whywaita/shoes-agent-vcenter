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

// VcenterGuestIpv6Address struct for VcenterGuestIpv6Address
type VcenterGuestIpv6Address struct {
	// Static IPv6 Address.
	IpAddress string `json:"ip_address"`
	// The CIDR prefix for the interface.
	Prefix int64 `json:"prefix"`
}

// NewVcenterGuestIpv6Address instantiates a new VcenterGuestIpv6Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterGuestIpv6Address(ipAddress string, prefix int64) *VcenterGuestIpv6Address {
	this := VcenterGuestIpv6Address{}
	this.IpAddress = ipAddress
	this.Prefix = prefix
	return &this
}

// NewVcenterGuestIpv6AddressWithDefaults instantiates a new VcenterGuestIpv6Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterGuestIpv6AddressWithDefaults() *VcenterGuestIpv6Address {
	this := VcenterGuestIpv6Address{}
	return &this
}

// GetIpAddress returns the IpAddress field value
func (o *VcenterGuestIpv6Address) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestIpv6Address) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *VcenterGuestIpv6Address) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetPrefix returns the Prefix field value
func (o *VcenterGuestIpv6Address) GetPrefix() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *VcenterGuestIpv6Address) GetPrefixOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *VcenterGuestIpv6Address) SetPrefix(v int64) {
	o.Prefix = v
}

func (o VcenterGuestIpv6Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip_address"] = o.IpAddress
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterGuestIpv6Address struct {
	value *VcenterGuestIpv6Address
	isSet bool
}

func (v NullableVcenterGuestIpv6Address) Get() *VcenterGuestIpv6Address {
	return v.value
}

func (v *NullableVcenterGuestIpv6Address) Set(val *VcenterGuestIpv6Address) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterGuestIpv6Address) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterGuestIpv6Address) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterGuestIpv6Address(val *VcenterGuestIpv6Address) *NullableVcenterGuestIpv6Address {
	return &NullableVcenterGuestIpv6Address{value: val, isSet: true}
}

func (v NullableVcenterGuestIpv6Address) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterGuestIpv6Address) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


