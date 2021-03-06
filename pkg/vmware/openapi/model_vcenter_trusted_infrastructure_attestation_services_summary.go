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

// VcenterTrustedInfrastructureAttestationServicesSummary struct for VcenterTrustedInfrastructureAttestationServicesSummary
type VcenterTrustedInfrastructureAttestationServicesSummary struct {
	// The service's unique identifier. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.attestation.Service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.attestation.Service.
	Service string `json:"service"`
	Address VcenterTrustedInfrastructureNetworkAddress `json:"address"`
	// The group specifies the Key Provider Service instances that can accept reports issued by this Attestation Service instance.
	Group string `json:"group"`
	// The cluster specifies the Trust Authority Cluster this Attestation Service instance belongs to.
	TrustAuthorityCluster string `json:"trust_authority_cluster"`
}

// NewVcenterTrustedInfrastructureAttestationServicesSummary instantiates a new VcenterTrustedInfrastructureAttestationServicesSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureAttestationServicesSummary(service string, address VcenterTrustedInfrastructureNetworkAddress, group string, trustAuthorityCluster string) *VcenterTrustedInfrastructureAttestationServicesSummary {
	this := VcenterTrustedInfrastructureAttestationServicesSummary{}
	this.Service = service
	this.Address = address
	this.Group = group
	this.TrustAuthorityCluster = trustAuthorityCluster
	return &this
}

// NewVcenterTrustedInfrastructureAttestationServicesSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureAttestationServicesSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureAttestationServicesSummaryWithDefaults() *VcenterTrustedInfrastructureAttestationServicesSummary {
	this := VcenterTrustedInfrastructureAttestationServicesSummary{}
	return &this
}

// GetService returns the Service field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetService() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Service
}

// GetServiceOk returns a tuple with the Service field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetServiceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Service, true
}

// SetService sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetService(v string) {
	o.Service = v
}

// GetAddress returns the Address field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress {
	if o == nil {
		var ret VcenterTrustedInfrastructureNetworkAddress
		return ret
	}

	return o.Address
}

// GetAddressOk returns a tuple with the Address field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Address, true
}

// SetAddress sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress) {
	o.Address = v
}

// GetGroup returns the Group field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Group
}

// GetGroupOk returns a tuple with the Group field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetGroupOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Group, true
}

// SetGroup sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetGroup(v string) {
	o.Group = v
}

// GetTrustAuthorityCluster returns the TrustAuthorityCluster field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetTrustAuthorityCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TrustAuthorityCluster
}

// GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field value
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetTrustAuthorityClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TrustAuthorityCluster, true
}

// SetTrustAuthorityCluster sets field value
func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetTrustAuthorityCluster(v string) {
	o.TrustAuthorityCluster = v
}

func (o VcenterTrustedInfrastructureAttestationServicesSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["service"] = o.Service
	}
	if true {
		toSerialize["address"] = o.Address
	}
	if true {
		toSerialize["group"] = o.Group
	}
	if true {
		toSerialize["trust_authority_cluster"] = o.TrustAuthorityCluster
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureAttestationServicesSummary struct {
	value *VcenterTrustedInfrastructureAttestationServicesSummary
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesSummary) Get() *VcenterTrustedInfrastructureAttestationServicesSummary {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesSummary) Set(val *VcenterTrustedInfrastructureAttestationServicesSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureAttestationServicesSummary(val *VcenterTrustedInfrastructureAttestationServicesSummary) *NullableVcenterTrustedInfrastructureAttestationServicesSummary {
	return &NullableVcenterTrustedInfrastructureAttestationServicesSummary{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


