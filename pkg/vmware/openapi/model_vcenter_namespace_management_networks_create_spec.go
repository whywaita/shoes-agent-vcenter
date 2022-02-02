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

// VcenterNamespaceManagementNetworksCreateSpec struct for VcenterNamespaceManagementNetworksCreateSpec
type VcenterNamespaceManagementNetworksCreateSpec struct {
	// Identifier of the network. This has DNS_LABEL restrictions as specified in . This must be an alphanumeric (a-z and 0-9) string and with maximum length of 63 characters and with the '-' character allowed anywhere except the first or last character. This name must be unique within a cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.namespace_management.Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.namespace_management.Network.
	Network string `json:"network"`
	NetworkProvider VcenterNamespaceManagementClustersNetworkProvider `json:"network_provider"`
	VsphereNetwork *VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec `json:"vsphere_network,omitempty"`
}

// NewVcenterNamespaceManagementNetworksCreateSpec instantiates a new VcenterNamespaceManagementNetworksCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementNetworksCreateSpec(network string, networkProvider VcenterNamespaceManagementClustersNetworkProvider) *VcenterNamespaceManagementNetworksCreateSpec {
	this := VcenterNamespaceManagementNetworksCreateSpec{}
	this.Network = network
	this.NetworkProvider = networkProvider
	return &this
}

// NewVcenterNamespaceManagementNetworksCreateSpecWithDefaults instantiates a new VcenterNamespaceManagementNetworksCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementNetworksCreateSpecWithDefaults() *VcenterNamespaceManagementNetworksCreateSpec {
	this := VcenterNamespaceManagementNetworksCreateSpec{}
	return &this
}

// GetNetwork returns the Network field value
func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetwork() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Network
}

// GetNetworkOk returns a tuple with the Network field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetworkOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Network, true
}

// SetNetwork sets field value
func (o *VcenterNamespaceManagementNetworksCreateSpec) SetNetwork(v string) {
	o.Network = v
}

// GetNetworkProvider returns the NetworkProvider field value
func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider {
	if o == nil {
		var ret VcenterNamespaceManagementClustersNetworkProvider
		return ret
	}

	return o.NetworkProvider
}

// GetNetworkProviderOk returns a tuple with the NetworkProvider field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksCreateSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NetworkProvider, true
}

// SetNetworkProvider sets field value
func (o *VcenterNamespaceManagementNetworksCreateSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider) {
	o.NetworkProvider = v
}

// GetVsphereNetwork returns the VsphereNetwork field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementNetworksCreateSpec) GetVsphereNetwork() VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec {
	if o == nil || o.VsphereNetwork == nil {
		var ret VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec
		return ret
	}
	return *o.VsphereNetwork
}

// GetVsphereNetworkOk returns a tuple with the VsphereNetwork field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksCreateSpec) GetVsphereNetworkOk() (*VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec, bool) {
	if o == nil || o.VsphereNetwork == nil {
		return nil, false
	}
	return o.VsphereNetwork, true
}

// HasVsphereNetwork returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementNetworksCreateSpec) HasVsphereNetwork() bool {
	if o != nil && o.VsphereNetwork != nil {
		return true
	}

	return false
}

// SetVsphereNetwork gets a reference to the given VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec and assigns it to the VsphereNetwork field.
func (o *VcenterNamespaceManagementNetworksCreateSpec) SetVsphereNetwork(v VcenterNamespaceManagementNetworksVsphereDVPGNetworkCreateSpec) {
	o.VsphereNetwork = &v
}

func (o VcenterNamespaceManagementNetworksCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["network"] = o.Network
	}
	if true {
		toSerialize["network_provider"] = o.NetworkProvider
	}
	if o.VsphereNetwork != nil {
		toSerialize["vsphere_network"] = o.VsphereNetwork
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementNetworksCreateSpec struct {
	value *VcenterNamespaceManagementNetworksCreateSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementNetworksCreateSpec) Get() *VcenterNamespaceManagementNetworksCreateSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementNetworksCreateSpec) Set(val *VcenterNamespaceManagementNetworksCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementNetworksCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementNetworksCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementNetworksCreateSpec(val *VcenterNamespaceManagementNetworksCreateSpec) *NullableVcenterNamespaceManagementNetworksCreateSpec {
	return &NullableVcenterNamespaceManagementNetworksCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementNetworksCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementNetworksCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

