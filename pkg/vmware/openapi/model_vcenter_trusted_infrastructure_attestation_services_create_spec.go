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

// VcenterTrustedInfrastructureAttestationServicesCreateSpec struct for VcenterTrustedInfrastructureAttestationServicesCreateSpec
type VcenterTrustedInfrastructureAttestationServicesCreateSpec struct {
	Address VcenterTrustedInfrastructureNetworkAddress `json:"address"`
	TrustedCA VcenterTrustedInfrastructureX509CertChain `json:"trusted_CA"`
	// The group specifies the Key Provider Service instances that can accept reports issued by this Attestation Service instance.
	Group string `json:"group"`
	// The cluster specifies the Trust Authority Cluster this Attestation Service belongs to.
	TrustAuthorityCluster string `json:"trust_authority_cluster"`
}

// NewVcenterTrustedInfrastructureAttestationServicesCreateSpec instantiates a new VcenterTrustedInfrastructureAttestationServicesCreateSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureAttestationServicesCreateSpec(address VcenterTrustedInfrastructureNetworkAddress, trustedCA VcenterTrustedInfrastructureX509CertChain, group string, trustAuthorityCluster string) *VcenterTrustedInfrastructureAttestationServicesCreateSpec {
	this := VcenterTrustedInfrastructureAttestationServicesCreateSpec{}
	this.Address = address
	this.TrustedCA = trustedCA
	this.Group = group
	this.TrustAuthorityCluster = trustAuthorityCluster
	return &this
}

// NewVcenterTrustedInfrastructureAttestationServicesCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureAttestationServicesCreateSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureAttestationServicesCreateSpecWithDefaults() *VcenterTrustedInfrastructureAttestationServicesCreateSpec {
	this := VcenterTrustedInfrastructureAttestationServicesCreateSpec{}
	return &this
}

// GetAddress returns the Address field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetAddress() VcenterTrustedInfrastructureNetworkAddress {
	if o == nil {
		var ret VcenterTrustedInfrastructureNetworkAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetAddress(v VcenterTrustedInfrastructureNetworkAddress) {
	o.Address = v
}

// GetTrustedCA returns the TrustedCA field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain {
	if o == nil {
		var ret VcenterTrustedInfrastructureX509CertChain
		return ret
	}

	return o.TrustedCA
}

// GetTrustedCAOk returns a tuple with the TrustedCA field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrustedCA, true
}

// SetTrustedCA sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain) {
	o.TrustedCA = v
}

// GetGroup returns the Group field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetGroup(v string) {
	o.Group = v
}

// GetTrustAuthorityCluster returns the TrustAuthorityCluster field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustAuthorityCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustAuthorityCluster
}

// GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustAuthorityClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrustAuthorityCluster, true
}

// SetTrustAuthorityCluster sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetTrustAuthorityCluster(v string) {
	o.TrustAuthorityCluster = v
}

func (o VcenterTrustedInfrastructureAttestationServicesCreateSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["trusted_CA"] = o.TrustedCA
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["trust_authority_cluster"] = o.TrustAuthorityCluster
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec struct {
	value *VcenterTrustedInfrastructureAttestationServicesCreateSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec) Get() *VcenterTrustedInfrastructureAttestationServicesCreateSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec) Set(val *VcenterTrustedInfrastructureAttestationServicesCreateSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureAttestationServicesCreateSpec(val *VcenterTrustedInfrastructureAttestationServicesCreateSpec) *NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec {
	return &NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesCreateSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


