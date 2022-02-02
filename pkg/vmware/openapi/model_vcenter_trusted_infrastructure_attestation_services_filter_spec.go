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

// VcenterTrustedInfrastructureAttestationServicesFilterSpec struct for VcenterTrustedInfrastructureAttestationServicesFilterSpec
type VcenterTrustedInfrastructureAttestationServicesFilterSpec struct {
	// A set of IDs by which to filter the services. If unset, the services will not be filtered by ID. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.trusted_infrastructure.attestation.Service. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.trusted_infrastructure.attestation.Service.
	Services *[]string `json:"services,omitempty"`
	// A set of address by which to filter. If unset, the services will not be filtered by address.
	Address *[]VcenterTrustedInfrastructureNetworkAddress `json:"address,omitempty"`
	// The group specifies the Key Provider Service instances that can accept reports issued by this Attestation Service instance. If unset, the services will not be filtered by group.
	Group *[]string `json:"group,omitempty"`
	// The cluster specifies the Trust Authority Cluster this Attestation Service belongs to. If unset, the services will not be filtered by trustAuthorityCluster.
	TrustAuthorityCluster *[]string `json:"trust_authority_cluster,omitempty"`
}

// NewVcenterTrustedInfrastructureAttestationServicesFilterSpec instantiates a new VcenterTrustedInfrastructureAttestationServicesFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureAttestationServicesFilterSpec() *VcenterTrustedInfrastructureAttestationServicesFilterSpec {
	this := VcenterTrustedInfrastructureAttestationServicesFilterSpec{}
	return &this
}

// NewVcenterTrustedInfrastructureAttestationServicesFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureAttestationServicesFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureAttestationServicesFilterSpecWithDefaults() *VcenterTrustedInfrastructureAttestationServicesFilterSpec {
	this := VcenterTrustedInfrastructureAttestationServicesFilterSpec{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetServices() []string {
	if o == nil || o.Services == nil {
		var ret []string
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetServicesOk() (*[]string, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given []string and assigns it to the Services field.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetServices(v []string) {
	o.Services = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress {
	if o == nil || o.Address == nil {
		var ret []VcenterTrustedInfrastructureNetworkAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given []VcenterTrustedInfrastructureNetworkAddress and assigns it to the Address field.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress) {
	o.Address = &v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetGroup() []string {
	if o == nil || o.Group == nil {
		var ret []string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetGroupOk() (*[]string, bool) {
	if o == nil || o.Group == nil {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasGroup() bool {
	if o != nil && o.Group != nil {
		return true
	}

	return false
}

// SetGroup gets a reference to the given []string and assigns it to the Group field.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetGroup(v []string) {
	o.Group = &v
}

// GetTrustAuthorityCluster returns the TrustAuthorityCluster field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetTrustAuthorityCluster() []string {
	if o == nil || o.TrustAuthorityCluster == nil {
		var ret []string
		return ret
	}
	return *o.TrustAuthorityCluster
}

// GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetTrustAuthorityClusterOk() (*[]string, bool) {
	if o == nil || o.TrustAuthorityCluster == nil {
		return nil, false
	}
	return o.TrustAuthorityCluster, true
}

// HasTrustAuthorityCluster returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasTrustAuthorityCluster() bool {
	if o != nil && o.TrustAuthorityCluster != nil {
		return true
	}

	return false
}

// SetTrustAuthorityCluster gets a reference to the given []string and assigns it to the TrustAuthorityCluster field.
func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetTrustAuthorityCluster(v []string) {
	o.TrustAuthorityCluster = &v
}

func (o VcenterTrustedInfrastructureAttestationServicesFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Group != nil {
		toSerialize["group"] = o.Group
	}
	if o.TrustAuthorityCluster != nil {
		toSerialize["trust_authority_cluster"] = o.TrustAuthorityCluster
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec struct {
	value *VcenterTrustedInfrastructureAttestationServicesFilterSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec) Get() *VcenterTrustedInfrastructureAttestationServicesFilterSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec) Set(val *VcenterTrustedInfrastructureAttestationServicesFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureAttestationServicesFilterSpec(val *VcenterTrustedInfrastructureAttestationServicesFilterSpec) *NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec {
	return &NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureAttestationServicesFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

