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

// VcenterNamespaceManagementIpv4Cidr struct for VcenterNamespaceManagementIpv4Cidr
type VcenterNamespaceManagementIpv4Cidr struct {
	// The IPv4 address.
	Address string `json:"address"`
	// The CIDR prefix.
	Prefix int64 `json:"prefix"`
}

// NewVcenterNamespaceManagementIpv4Cidr instantiates a new VcenterNamespaceManagementIpv4Cidr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementIpv4Cidr(address string, prefix int64) *VcenterNamespaceManagementIpv4Cidr {
	this := VcenterNamespaceManagementIpv4Cidr{}
	this.Address = address
	this.Prefix = prefix
	return &this
}

// NewVcenterNamespaceManagementIpv4CidrWithDefaults instantiates a new VcenterNamespaceManagementIpv4Cidr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementIpv4CidrWithDefaults() *VcenterNamespaceManagementIpv4Cidr {
	this := VcenterNamespaceManagementIpv4Cidr{}
	return &this
}

// GetAddress returns the Address field value
func (o *VcenterNamespaceManagementIpv4Cidr) GetAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementIpv4Cidr) GetAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VcenterNamespaceManagementIpv4Cidr) SetAddress(v string) {
	o.Address = v
}

// GetPrefix returns the Prefix field value
func (o *VcenterNamespaceManagementIpv4Cidr) GetPrefix() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementIpv4Cidr) GetPrefixOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *VcenterNamespaceManagementIpv4Cidr) SetPrefix(v int64) {
	o.Prefix = v
}

func (o VcenterNamespaceManagementIpv4Cidr) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["prefix"] = o.Prefix
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementIpv4Cidr struct {
	value *VcenterNamespaceManagementIpv4Cidr
	isSet bool
}

func (v NullableVcenterNamespaceManagementIpv4Cidr) Get() *VcenterNamespaceManagementIpv4Cidr {
	return v.value
}

func (v *NullableVcenterNamespaceManagementIpv4Cidr) Set(val *VcenterNamespaceManagementIpv4Cidr) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementIpv4Cidr) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementIpv4Cidr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementIpv4Cidr(val *VcenterNamespaceManagementIpv4Cidr) *NullableVcenterNamespaceManagementIpv4Cidr {
	return &NullableVcenterNamespaceManagementIpv4Cidr{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementIpv4Cidr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementIpv4Cidr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

