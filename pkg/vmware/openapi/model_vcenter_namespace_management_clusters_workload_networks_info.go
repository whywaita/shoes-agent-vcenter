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

// VcenterNamespaceManagementClustersWorkloadNetworksInfo struct for VcenterNamespaceManagementClustersWorkloadNetworksInfo
type VcenterNamespaceManagementClustersWorkloadNetworksInfo struct {
	SupervisorPrimaryWorkloadNetwork VcenterNamespaceManagementNetworksInfo `json:"supervisor_primary_workload_network"`
	// List of vSphere Namespaces networks associated with this cluster.
	NetworkList *[]VcenterNamespaceManagementNetworksInfo `json:"network_list,omitempty"`
}

// NewVcenterNamespaceManagementClustersWorkloadNetworksInfo instantiates a new VcenterNamespaceManagementClustersWorkloadNetworksInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementClustersWorkloadNetworksInfo(supervisorPrimaryWorkloadNetwork VcenterNamespaceManagementNetworksInfo) *VcenterNamespaceManagementClustersWorkloadNetworksInfo {
	this := VcenterNamespaceManagementClustersWorkloadNetworksInfo{}
	this.SupervisorPrimaryWorkloadNetwork = supervisorPrimaryWorkloadNetwork
	return &this
}

// NewVcenterNamespaceManagementClustersWorkloadNetworksInfoWithDefaults instantiates a new VcenterNamespaceManagementClustersWorkloadNetworksInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementClustersWorkloadNetworksInfoWithDefaults() *VcenterNamespaceManagementClustersWorkloadNetworksInfo {
	this := VcenterNamespaceManagementClustersWorkloadNetworksInfo{}
	return &this
}

// GetSupervisorPrimaryWorkloadNetwork returns the SupervisorPrimaryWorkloadNetwork field value
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetSupervisorPrimaryWorkloadNetwork() VcenterNamespaceManagementNetworksInfo {
	if o == nil {
		var ret VcenterNamespaceManagementNetworksInfo
		return ret
	}

	return o.SupervisorPrimaryWorkloadNetwork
}

// GetSupervisorPrimaryWorkloadNetworkOk returns a tuple with the SupervisorPrimaryWorkloadNetwork field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetSupervisorPrimaryWorkloadNetworkOk() (*VcenterNamespaceManagementNetworksInfo, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SupervisorPrimaryWorkloadNetwork, true
}

// SetSupervisorPrimaryWorkloadNetwork sets field value
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) SetSupervisorPrimaryWorkloadNetwork(v VcenterNamespaceManagementNetworksInfo) {
	o.SupervisorPrimaryWorkloadNetwork = v
}

// GetNetworkList returns the NetworkList field value if set, zero value otherwise.
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetNetworkList() []VcenterNamespaceManagementNetworksInfo {
	if o == nil || o.NetworkList == nil {
		var ret []VcenterNamespaceManagementNetworksInfo
		return ret
	}
	return *o.NetworkList
}

// GetNetworkListOk returns a tuple with the NetworkList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) GetNetworkListOk() (*[]VcenterNamespaceManagementNetworksInfo, bool) {
	if o == nil || o.NetworkList == nil {
		return nil, false
	}
	return o.NetworkList, true
}

// HasNetworkList returns a boolean if a field has been set.
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) HasNetworkList() bool {
	if o != nil && o.NetworkList != nil {
		return true
	}

	return false
}

// SetNetworkList gets a reference to the given []VcenterNamespaceManagementNetworksInfo and assigns it to the NetworkList field.
func (o *VcenterNamespaceManagementClustersWorkloadNetworksInfo) SetNetworkList(v []VcenterNamespaceManagementNetworksInfo) {
	o.NetworkList = &v
}

func (o VcenterNamespaceManagementClustersWorkloadNetworksInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["supervisor_primary_workload_network"] = o.SupervisorPrimaryWorkloadNetwork
	}
	if o.NetworkList != nil {
		toSerialize["network_list"] = o.NetworkList
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo struct {
	value *VcenterNamespaceManagementClustersWorkloadNetworksInfo
	isSet bool
}

func (v NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo) Get() *VcenterNamespaceManagementClustersWorkloadNetworksInfo {
	return v.value
}

func (v *NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo) Set(val *VcenterNamespaceManagementClustersWorkloadNetworksInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementClustersWorkloadNetworksInfo(val *VcenterNamespaceManagementClustersWorkloadNetworksInfo) *NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo {
	return &NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementClustersWorkloadNetworksInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


