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

// VcenterNamespaceManagementClustersSummary struct for VcenterNamespaceManagementClustersSummary
type VcenterNamespaceManagementClustersSummary struct {
	// Identifier for the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.
	Cluster string `json:"cluster"`
	// Name of the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: ClusterComputeResource.name. When operations return a value of this structure as a result, the field will be an identifier for the resource type: ClusterComputeResource.name.
	ClusterName string `json:"cluster_name"`
	Stats VcenterNamespaceManagementClustersStats `json:"stats"`
	ConfigStatus VcenterNamespaceManagementClustersConfigStatus `json:"config_status"`
	KubernetesStatus VcenterNamespaceManagementClustersKubernetesStatus `json:"kubernetes_status"`
}

// NewVcenterNamespaceManagementClustersSummary instantiates a new VcenterNamespaceManagementClustersSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementClustersSummary(cluster string, clusterName string, stats VcenterNamespaceManagementClustersStats, configStatus VcenterNamespaceManagementClustersConfigStatus, kubernetesStatus VcenterNamespaceManagementClustersKubernetesStatus) *VcenterNamespaceManagementClustersSummary {
	this := VcenterNamespaceManagementClustersSummary{}
	this.Cluster = cluster
	this.ClusterName = clusterName
	this.Stats = stats
	this.ConfigStatus = configStatus
	this.KubernetesStatus = kubernetesStatus
	return &this
}

// NewVcenterNamespaceManagementClustersSummaryWithDefaults instantiates a new VcenterNamespaceManagementClustersSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementClustersSummaryWithDefaults() *VcenterNamespaceManagementClustersSummary {
	this := VcenterNamespaceManagementClustersSummary{}
	return &this
}

// GetCluster returns the Cluster field value
func (o *VcenterNamespaceManagementClustersSummary) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersSummary) GetClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Cluster, true
}

// SetCluster sets field value
func (o *VcenterNamespaceManagementClustersSummary) SetCluster(v string) {
	o.Cluster = v
}

// GetClusterName returns the ClusterName field value
func (o *VcenterNamespaceManagementClustersSummary) GetClusterName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterName
}

// GetClusterNameOk returns a tuple with the ClusterName field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersSummary) GetClusterNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterName, true
}

// SetClusterName sets field value
func (o *VcenterNamespaceManagementClustersSummary) SetClusterName(v string) {
	o.ClusterName = v
}

// GetStats returns the Stats field value
func (o *VcenterNamespaceManagementClustersSummary) GetStats() VcenterNamespaceManagementClustersStats {
	if o == nil {
		var ret VcenterNamespaceManagementClustersStats
		return ret
	}

	return o.Stats
}

// GetStatsOk returns a tuple with the Stats field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersSummary) GetStatsOk() (*VcenterNamespaceManagementClustersStats, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Stats, true
}

// SetStats sets field value
func (o *VcenterNamespaceManagementClustersSummary) SetStats(v VcenterNamespaceManagementClustersStats) {
	o.Stats = v
}

// GetConfigStatus returns the ConfigStatus field value
func (o *VcenterNamespaceManagementClustersSummary) GetConfigStatus() VcenterNamespaceManagementClustersConfigStatus {
	if o == nil {
		var ret VcenterNamespaceManagementClustersConfigStatus
		return ret
	}

	return o.ConfigStatus
}

// GetConfigStatusOk returns a tuple with the ConfigStatus field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersSummary) GetConfigStatusOk() (*VcenterNamespaceManagementClustersConfigStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConfigStatus, true
}

// SetConfigStatus sets field value
func (o *VcenterNamespaceManagementClustersSummary) SetConfigStatus(v VcenterNamespaceManagementClustersConfigStatus) {
	o.ConfigStatus = v
}

// GetKubernetesStatus returns the KubernetesStatus field value
func (o *VcenterNamespaceManagementClustersSummary) GetKubernetesStatus() VcenterNamespaceManagementClustersKubernetesStatus {
	if o == nil {
		var ret VcenterNamespaceManagementClustersKubernetesStatus
		return ret
	}

	return o.KubernetesStatus
}

// GetKubernetesStatusOk returns a tuple with the KubernetesStatus field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersSummary) GetKubernetesStatusOk() (*VcenterNamespaceManagementClustersKubernetesStatus, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.KubernetesStatus, true
}

// SetKubernetesStatus sets field value
func (o *VcenterNamespaceManagementClustersSummary) SetKubernetesStatus(v VcenterNamespaceManagementClustersKubernetesStatus) {
	o.KubernetesStatus = v
}

func (o VcenterNamespaceManagementClustersSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["cluster"] = o.Cluster
	}
	if true {
		toSerialize["cluster_name"] = o.ClusterName
	}
	if true {
		toSerialize["stats"] = o.Stats
	}
	if true {
		toSerialize["config_status"] = o.ConfigStatus
	}
	if true {
		toSerialize["kubernetes_status"] = o.KubernetesStatus
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementClustersSummary struct {
	value *VcenterNamespaceManagementClustersSummary
	isSet bool
}

func (v NullableVcenterNamespaceManagementClustersSummary) Get() *VcenterNamespaceManagementClustersSummary {
	return v.value
}

func (v *NullableVcenterNamespaceManagementClustersSummary) Set(val *VcenterNamespaceManagementClustersSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementClustersSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementClustersSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementClustersSummary(val *VcenterNamespaceManagementClustersSummary) *NullableVcenterNamespaceManagementClustersSummary {
	return &NullableVcenterNamespaceManagementClustersSummary{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementClustersSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementClustersSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


