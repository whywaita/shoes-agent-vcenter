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

// VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec struct for VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec
type VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec struct {
	// Identifier of the vSphere Distributed Portgroup backing the vSphere network object. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, modification to existing portgroup will result in the operation failing with a ResourceInUse error. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: Network. When operations return a value of this structure as a result, the field will be an identifier for the resource type: Network.
	Portgroup string `json:"portgroup"`
	// Usable IP pools on this network. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, then replacement of or modification to any existing range in this list will result in operation failing with a ResourceInUse error. To add new address ranges to the list, existing address ranges have to be passed in without modifications.
	AddressRanges []VcenterNamespaceManagementIPRange `json:"address_ranges"`
	// Gateway for the network. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, then modification to existing gateway will result in the operation failing with a ResourceInUse error.
	Gateway string `json:"gateway"`
	// Subnet mask of the network. If the network object is associated with a Namespace or is Clusters.WorkloadNetworksEnableSpec.supervisor-primary-workload-network, then modification to existing subnet mask will result in the operation failing with a ResourceInUse error.
	SubnetMask string `json:"subnet_mask"`
}

// NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec instantiates a new VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec(portgroup string, addressRanges []VcenterNamespaceManagementIPRange, gateway string, subnetMask string) *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec {
	this := VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec{}
	this.Portgroup = portgroup
	this.AddressRanges = addressRanges
	this.Gateway = gateway
	this.SubnetMask = subnetMask
	return &this
}

// NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpecWithDefaults instantiates a new VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpecWithDefaults() *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec {
	this := VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec{}
	return &this
}

// GetPortgroup returns the Portgroup field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetPortgroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Portgroup
}

// GetPortgroupOk returns a tuple with the Portgroup field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetPortgroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Portgroup, true
}

// SetPortgroup sets field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetPortgroup(v string) {
	o.Portgroup = v
}

// GetAddressRanges returns the AddressRanges field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetAddressRanges() []VcenterNamespaceManagementIPRange {
	if o == nil {
		var ret []VcenterNamespaceManagementIPRange
		return ret
	}

	return o.AddressRanges
}

// GetAddressRangesOk returns a tuple with the AddressRanges field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetAddressRangesOk() (*[]VcenterNamespaceManagementIPRange, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AddressRanges, true
}

// SetAddressRanges sets field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetAddressRanges(v []VcenterNamespaceManagementIPRange) {
	o.AddressRanges = v
}

// GetGateway returns the Gateway field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetGateway() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Gateway
}

// GetGatewayOk returns a tuple with the Gateway field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetGatewayOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Gateway, true
}

// SetGateway sets field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetGateway(v string) {
	o.Gateway = v
}

// GetSubnetMask returns the SubnetMask field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetSubnetMask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SubnetMask
}

// GetSubnetMaskOk returns a tuple with the SubnetMask field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) GetSubnetMaskOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubnetMask, true
}

// SetSubnetMask sets field value
func (o *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) SetSubnetMask(v string) {
	o.SubnetMask = v
}

func (o VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["portgroup"] = o.Portgroup
	}
	if true {
		toSerialize["address_ranges"] = o.AddressRanges
	}
	if true {
		toSerialize["gateway"] = o.Gateway
	}
	if true {
		toSerialize["subnet_mask"] = o.SubnetMask
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec struct {
	value *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec
	isSet bool
}

func (v NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) Get() *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec {
	return v.value
}

func (v *NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) Set(val *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec(val *VcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) *NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec {
	return &NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementNetworksVsphereDVPGNetworkSetSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


