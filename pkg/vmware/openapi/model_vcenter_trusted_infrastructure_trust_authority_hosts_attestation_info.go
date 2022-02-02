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

// VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo struct for VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo
type VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo struct {
	// The trusted ESX on which the service runs. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem.
	Host string `json:"host"`
	Address VcenterTrustedInfrastructureNetworkAddress `json:"address"`
	// The group ID determines which Attestation Service instances this Attestation Service can communicate with.
	Group string `json:"group"`
	// The opaque string identifier of the cluster in which the Attestation Service is part of.
	Cluster string `json:"cluster"`
	TrustedCA VcenterTrustedInfrastructureX509CertChain `json:"trusted_CA"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo(host string, address VcenterTrustedInfrastructureNetworkAddress, group string, cluster string, trustedCA VcenterTrustedInfrastructureX509CertChain) *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo {
	this := VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo{}
	this.Host = host
	this.Address = address
	this.Group = group
	this.Cluster = cluster
	this.TrustedCA = trustedCA
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfoWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfoWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo {
	this := VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo{}
	return &this
}

// GetHost returns the Host field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetHost(v string) {
	o.Host = v
}

// GetAddress returns the Address field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetAddress() VcenterTrustedInfrastructureNetworkAddress {
	if o == nil {
		var ret VcenterTrustedInfrastructureNetworkAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetAddress(v VcenterTrustedInfrastructureNetworkAddress) {
	o.Address = v
}

// GetGroup returns the Group field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetGroup(v string) {
	o.Group = v
}

// GetCluster returns the Cluster field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetCluster(v string) {
	o.Cluster = v
}

// GetTrustedCA returns the TrustedCA field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain {
	if o == nil {
		var ret VcenterTrustedInfrastructureX509CertChain
		return ret
	}

	return o.TrustedCA
}

// GetTrustedCAOk returns a tuple with the TrustedCA field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrustedCA, true
}

// SetTrustedCA sets field value
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain) {
	o.TrustedCA = v
}

func (o VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["trusted_CA"] = o.TrustedCA
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo struct {
	value *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) Get() *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) Set(val *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo(val *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) *NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

