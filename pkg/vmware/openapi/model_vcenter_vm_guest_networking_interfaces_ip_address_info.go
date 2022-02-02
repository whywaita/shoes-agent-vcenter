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

// VcenterVmGuestNetworkingInterfacesIpAddressInfo struct for VcenterVmGuestNetworkingInterfacesIpAddressInfo
type VcenterVmGuestNetworkingInterfacesIpAddressInfo struct {
	// IPv4 address is specified using dotted decimal notation. For example, \"192.0.2.1\". IPv6 addresses are 128-bit addresses specified using eight fields of up to four hexadecimal digits. A colon separates each field (:). For example, 2001:DB8:101::230:6eff:fe04:d9ff. The address can also consist of the symbol '::' to represent multiple 16-bit groups of contiguous 0's only once in an address as described in RFC 2373.
	IpAddress string `json:"ip_address"`
	// Denotes the length of a generic Internet network address prefix. Prefix length: the valid range of values is 0-32 for IPv4, and 0-128 for IPv6. A value of n corresponds to an IP address mask that has n contiguous 1-bits from the most significant bit (MSB), with all other bits set to 0. A value of zero is valid only if the calling context defines it.
	PrefixLength int64 `json:"prefix_length"`
	Origin *VcenterVmGuestNetworkingInterfacesIpAddressOrigin `json:"origin,omitempty"`
	State VcenterVmGuestNetworkingInterfacesIpAddressStatus `json:"state"`
}

// NewVcenterVmGuestNetworkingInterfacesIpAddressInfo instantiates a new VcenterVmGuestNetworkingInterfacesIpAddressInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterVmGuestNetworkingInterfacesIpAddressInfo(ipAddress string, prefixLength int64, state VcenterVmGuestNetworkingInterfacesIpAddressStatus) *VcenterVmGuestNetworkingInterfacesIpAddressInfo {
	this := VcenterVmGuestNetworkingInterfacesIpAddressInfo{}
	this.IpAddress = ipAddress
	this.PrefixLength = prefixLength
	this.State = state
	return &this
}

// NewVcenterVmGuestNetworkingInterfacesIpAddressInfoWithDefaults instantiates a new VcenterVmGuestNetworkingInterfacesIpAddressInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterVmGuestNetworkingInterfacesIpAddressInfoWithDefaults() *VcenterVmGuestNetworkingInterfacesIpAddressInfo {
	this := VcenterVmGuestNetworkingInterfacesIpAddressInfo{}
	return &this
}

// GetIpAddress returns the IpAddress field value
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetIpAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetIpAddressOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IpAddress, true
}

// SetIpAddress sets field value
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetIpAddress(v string) {
	o.IpAddress = v
}

// GetPrefixLength returns the PrefixLength field value
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetPrefixLength() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PrefixLength
}

// GetPrefixLengthOk returns a tuple with the PrefixLength field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetPrefixLengthOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PrefixLength, true
}

// SetPrefixLength sets field value
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetPrefixLength(v int64) {
	o.PrefixLength = v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetOrigin() VcenterVmGuestNetworkingInterfacesIpAddressOrigin {
	if o == nil || o.Origin == nil {
		var ret VcenterVmGuestNetworkingInterfacesIpAddressOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetOriginOk() (*VcenterVmGuestNetworkingInterfacesIpAddressOrigin, bool) {
	if o == nil || o.Origin == nil {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) HasOrigin() bool {
	if o != nil && o.Origin != nil {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given VcenterVmGuestNetworkingInterfacesIpAddressOrigin and assigns it to the Origin field.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetOrigin(v VcenterVmGuestNetworkingInterfacesIpAddressOrigin) {
	o.Origin = &v
}

// GetState returns the State field value
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetState() VcenterVmGuestNetworkingInterfacesIpAddressStatus {
	if o == nil {
		var ret VcenterVmGuestNetworkingInterfacesIpAddressStatus
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) GetStateOk() (*VcenterVmGuestNetworkingInterfacesIpAddressStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *VcenterVmGuestNetworkingInterfacesIpAddressInfo) SetState(v VcenterVmGuestNetworkingInterfacesIpAddressStatus) {
	o.State = v
}

func (o VcenterVmGuestNetworkingInterfacesIpAddressInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ip_address"] = o.IpAddress
	}
	if true {
		toSerialize["prefix_length"] = o.PrefixLength
	}
	if o.Origin != nil {
		toSerialize["origin"] = o.Origin
	}
	if true {
		toSerialize["state"] = o.State
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo struct {
	value *VcenterVmGuestNetworkingInterfacesIpAddressInfo
	isSet bool
}

func (v NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo) Get() *VcenterVmGuestNetworkingInterfacesIpAddressInfo {
	return v.value
}

func (v *NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo) Set(val *VcenterVmGuestNetworkingInterfacesIpAddressInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterVmGuestNetworkingInterfacesIpAddressInfo(val *VcenterVmGuestNetworkingInterfacesIpAddressInfo) *NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo {
	return &NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo{value: val, isSet: true}
}

func (v NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterVmGuestNetworkingInterfacesIpAddressInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


