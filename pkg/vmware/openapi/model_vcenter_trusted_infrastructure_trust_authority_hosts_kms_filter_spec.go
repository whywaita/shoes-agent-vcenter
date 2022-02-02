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

// VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec struct for VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec
type VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec struct {
	// A set of host IDs by which to filter the services. If unset, the services will not be filtered by the hosts on which they run. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: HostSystem. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: HostSystem.
	Hosts *[]string `json:"hosts,omitempty"`
	// A set of cluster IDs by which to filter the services. If unset, the services will not be filtered by the clusters on which they run. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource.
	Clusters *[]string `json:"clusters,omitempty"`
	// The service's address. If unset, the services will not be filtered by address.
	Address *[]VcenterTrustedInfrastructureNetworkAddress `json:"address,omitempty"`
	// The group determines reports issued by which Attestation Service instances this Key Provider Service can accept. If unset, the services will not be filtered by groupId.
	Groups *[]string `json:"groups,omitempty"`
}

// NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec{}
	return &this
}

// NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec {
	this := VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec{}
	return &this
}

// GetHosts returns the Hosts field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetHosts() []string {
	if o == nil || o.Hosts == nil {
		var ret []string
		return ret
	}
	return *o.Hosts
}

// GetHostsOk returns a tuple with the Hosts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetHostsOk() (*[]string, bool) {
	if o == nil || o.Hosts == nil {
		return nil, false
	}
	return o.Hosts, true
}

// HasHosts returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasHosts() bool {
	if o != nil && o.Hosts != nil {
		return true
	}

	return false
}

// SetHosts gets a reference to the given []string and assigns it to the Hosts field.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetHosts(v []string) {
	o.Hosts = &v
}

// GetClusters returns the Clusters field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetClusters() []string {
	if o == nil || o.Clusters == nil {
		var ret []string
		return ret
	}
	return *o.Clusters
}

// GetClustersOk returns a tuple with the Clusters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetClustersOk() (*[]string, bool) {
	if o == nil || o.Clusters == nil {
		return nil, false
	}
	return o.Clusters, true
}

// HasClusters returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasClusters() bool {
	if o != nil && o.Clusters != nil {
		return true
	}

	return false
}

// SetClusters gets a reference to the given []string and assigns it to the Clusters field.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetClusters(v []string) {
	o.Clusters = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress {
	if o == nil || o.Address == nil {
		var ret []VcenterTrustedInfrastructureNetworkAddress
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given []VcenterTrustedInfrastructureNetworkAddress and assigns it to the Address field.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress) {
	o.Address = &v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetGroups() []string {
	if o == nil || o.Groups == nil {
		var ret []string
		return ret
	}
	return *o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetGroupsOk() (*[]string, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetGroups(v []string) {
	o.Groups = &v
}

func (o VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hosts != nil {
		toSerialize["hosts"] = o.Hosts
	}
	if o.Clusters != nil {
		toSerialize["clusters"] = o.Clusters
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec struct {
	value *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec
	isSet bool
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) Get() *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec {
	return v.value
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) Set(val *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec(val *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) *NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec {
	return &NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec{value: val, isSet: true}
}

func (v NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

