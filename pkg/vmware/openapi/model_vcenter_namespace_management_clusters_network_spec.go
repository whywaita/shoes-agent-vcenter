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

// VcenterNamespaceManagementClustersNetworkSpec struct for VcenterNamespaceManagementClustersNetworkSpec
type VcenterNamespaceManagementClustersNetworkSpec struct {
	// Optionally specify the Floating IP used by the HA master cluster in the DHCP case. This field is optional and it is only relevant when the value of Clusters.NetworkSpec.mode is DHCP.
	FloatingIP *string `json:"floating_IP,omitempty"`
	// Identifier for the network. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network.
	Network string `json:"network"`
	Mode VcenterNamespaceManagementClustersNetworkSpecIpv4Mode `json:"mode"`
	AddressRange *VcenterNamespaceManagementClustersIpv4Range `json:"address_range,omitempty"`
}

// NewVcenterNamespaceManagementClustersNetworkSpec instantiates a new VcenterNamespaceManagementClustersNetworkSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementClustersNetworkSpec(network string, mode VcenterNamespaceManagementClustersNetworkSpecIpv4Mode) *VcenterNamespaceManagementClustersNetworkSpec {
	this := VcenterNamespaceManagementClustersNetworkSpec{}
	this.Network = network
	this.Mode = mode
	return &this
}

// NewVcenterNamespaceManagementClustersNetworkSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersNetworkSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementClustersNetworkSpecWithDefaults() *VcenterNamespaceManagementClustersNetworkSpec {
	this := VcenterNamespaceManagementClustersNetworkSpec{}
	return &this
}

// GetFloatingIP returns the FloatingIP field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetFloatingIP() string {
	if o == nil || o.FloatingIP == nil {
		var ret string
		return ret
	}
	return *o.FloatingIP
}

// GetFloatingIPOk returns a tuple with the FloatingIP field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetFloatingIPOk() (*string, bool) {
	if o == nil || o.FloatingIP == nil {
		return nil, false
	}
	return o.FloatingIP, true
}

// HasFloatingIP returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementClustersNetworkSpec) HasFloatingIP() bool {
	if o != nil && o.FloatingIP != nil {
		return true
	}

	return false
}

// SetFloatingIP gets a reference to the given string and assigns it to the FloatingIP field.
func (o *VcenterNamespaceManagementClustersNetworkSpec) SetFloatingIP(v string) {
	o.FloatingIP = &v
}

// GetNetwork returns the Network field value
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *VcenterNamespaceManagementClustersNetworkSpec) SetNetwork(v string) {
	o.Network = v
}

// GetMode returns the Mode field value
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetMode() VcenterNamespaceManagementClustersNetworkSpecIpv4Mode {
	if o == nil {
		var ret VcenterNamespaceManagementClustersNetworkSpecIpv4Mode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetModeOk() (*VcenterNamespaceManagementClustersNetworkSpecIpv4Mode, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *VcenterNamespaceManagementClustersNetworkSpec) SetMode(v VcenterNamespaceManagementClustersNetworkSpecIpv4Mode) {
	o.Mode = v
}

// GetAddressRange returns the AddressRange field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetAddressRange() VcenterNamespaceManagementClustersIpv4Range {
	if o == nil || o.AddressRange == nil {
		var ret VcenterNamespaceManagementClustersIpv4Range
		return ret
	}
	return *o.AddressRange
}

// GetAddressRangeOk returns a tuple with the AddressRange field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNetworkSpec) GetAddressRangeOk() (*VcenterNamespaceManagementClustersIpv4Range, bool) {
	if o == nil || o.AddressRange == nil {
		return nil, false
	}
	return o.AddressRange, true
}

// HasAddressRange returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementClustersNetworkSpec) HasAddressRange() bool {
	if o != nil && o.AddressRange != nil {
		return true
	}

	return false
}

// SetAddressRange gets a reference to the given VcenterNamespaceManagementClustersIpv4Range and assigns it to the AddressRange field.
func (o *VcenterNamespaceManagementClustersNetworkSpec) SetAddressRange(v VcenterNamespaceManagementClustersIpv4Range) {
	o.AddressRange = &v
}

func (o VcenterNamespaceManagementClustersNetworkSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FloatingIP != nil {
		toSerialize["floating_IP"] = o.FloatingIP
	}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["mode"] = o.Mode
	}
	if o.AddressRange != nil {
		toSerialize["address_range"] = o.AddressRange
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementClustersNetworkSpec struct {
	value *VcenterNamespaceManagementClustersNetworkSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementClustersNetworkSpec) Get() *VcenterNamespaceManagementClustersNetworkSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementClustersNetworkSpec) Set(val *VcenterNamespaceManagementClustersNetworkSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementClustersNetworkSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementClustersNetworkSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementClustersNetworkSpec(val *VcenterNamespaceManagementClustersNetworkSpec) *NullableVcenterNamespaceManagementClustersNetworkSpec {
	return &NullableVcenterNamespaceManagementClustersNetworkSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementClustersNetworkSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementClustersNetworkSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

