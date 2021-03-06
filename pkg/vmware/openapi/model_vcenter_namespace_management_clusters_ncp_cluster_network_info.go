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

// VcenterNamespaceManagementClustersNCPClusterNetworkInfo struct for VcenterNamespaceManagementClustersNCPClusterNetworkInfo
type VcenterNamespaceManagementClustersNCPClusterNetworkInfo struct {
	// CIDR blocks from which Kubernetes allocates pod IP addresses.
	PodCidrs []VcenterNamespaceManagementIpv4Cidr `json:"pod_cidrs"`
	// CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer.
	IngressCidrs []VcenterNamespaceManagementIpv4Cidr `json:"ingress_cidrs"`
	// CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs.
	EgressCidrs []VcenterNamespaceManagementIpv4Cidr `json:"egress_cidrs"`
	// vSphere Distributed Switch used to connect this cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vSphereDistributedSwitch. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vSphereDistributedSwitch.
	ClusterDistributedSwitch string `json:"cluster_distributed_switch"`
	// NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: NSXEdgeCluster. When operations return a value of this structure as a result, the field will be an identifier for the resource type: NSXEdgeCluster.
	NsxEdgeCluster string `json:"nsx_edge_cluster"`
	// PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services.
	DefaultIngressTlsCertificate string `json:"default_ingress_tls_certificate"`
}

// NewVcenterNamespaceManagementClustersNCPClusterNetworkInfo instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVcenterNamespaceManagementClustersNCPClusterNetworkInfo(podCidrs []VcenterNamespaceManagementIpv4Cidr, ingressCidrs []VcenterNamespaceManagementIpv4Cidr, egressCidrs []VcenterNamespaceManagementIpv4Cidr, clusterDistributedSwitch string, nsxEdgeCluster string, defaultIngressTlsCertificate string) *VcenterNamespaceManagementClustersNCPClusterNetworkInfo {
	this := VcenterNamespaceManagementClustersNCPClusterNetworkInfo{}
	this.PodCidrs = podCidrs
	this.IngressCidrs = ingressCidrs
	this.EgressCidrs = egressCidrs
	this.ClusterDistributedSwitch = clusterDistributedSwitch
	this.NsxEdgeCluster = nsxEdgeCluster
	this.DefaultIngressTlsCertificate = defaultIngressTlsCertificate
	return &this
}

// NewVcenterNamespaceManagementClustersNCPClusterNetworkInfoWithDefaults instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVcenterNamespaceManagementClustersNCPClusterNetworkInfoWithDefaults() *VcenterNamespaceManagementClustersNCPClusterNetworkInfo {
	this := VcenterNamespaceManagementClustersNCPClusterNetworkInfo{}
	return &this
}

// GetPodCidrs returns the PodCidrs field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetPodCidrs() []VcenterNamespaceManagementIpv4Cidr {
	if o == nil {
		var ret []VcenterNamespaceManagementIpv4Cidr
		return ret
	}

	return o.PodCidrs
}

// GetPodCidrsOk returns a tuple with the PodCidrs field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetPodCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PodCidrs, true
}

// SetPodCidrs sets field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetPodCidrs(v []VcenterNamespaceManagementIpv4Cidr) {
	o.PodCidrs = v
}

// GetIngressCidrs returns the IngressCidrs field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetIngressCidrs() []VcenterNamespaceManagementIpv4Cidr {
	if o == nil {
		var ret []VcenterNamespaceManagementIpv4Cidr
		return ret
	}

	return o.IngressCidrs
}

// GetIngressCidrsOk returns a tuple with the IngressCidrs field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetIngressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IngressCidrs, true
}

// SetIngressCidrs sets field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetIngressCidrs(v []VcenterNamespaceManagementIpv4Cidr) {
	o.IngressCidrs = v
}

// GetEgressCidrs returns the EgressCidrs field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetEgressCidrs() []VcenterNamespaceManagementIpv4Cidr {
	if o == nil {
		var ret []VcenterNamespaceManagementIpv4Cidr
		return ret
	}

	return o.EgressCidrs
}

// GetEgressCidrsOk returns a tuple with the EgressCidrs field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetEgressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EgressCidrs, true
}

// SetEgressCidrs sets field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetEgressCidrs(v []VcenterNamespaceManagementIpv4Cidr) {
	o.EgressCidrs = v
}

// GetClusterDistributedSwitch returns the ClusterDistributedSwitch field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetClusterDistributedSwitch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClusterDistributedSwitch
}

// GetClusterDistributedSwitchOk returns a tuple with the ClusterDistributedSwitch field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetClusterDistributedSwitchOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ClusterDistributedSwitch, true
}

// SetClusterDistributedSwitch sets field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetClusterDistributedSwitch(v string) {
	o.ClusterDistributedSwitch = v
}

// GetNsxEdgeCluster returns the NsxEdgeCluster field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetNsxEdgeCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NsxEdgeCluster
}

// GetNsxEdgeClusterOk returns a tuple with the NsxEdgeCluster field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetNsxEdgeClusterOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NsxEdgeCluster, true
}

// SetNsxEdgeCluster sets field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetNsxEdgeCluster(v string) {
	o.NsxEdgeCluster = v
}

// GetDefaultIngressTlsCertificate returns the DefaultIngressTlsCertificate field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetDefaultIngressTlsCertificate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultIngressTlsCertificate
}

// GetDefaultIngressTlsCertificateOk returns a tuple with the DefaultIngressTlsCertificate field value
// and a boolean to check if the value has been set.
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetDefaultIngressTlsCertificateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DefaultIngressTlsCertificate, true
}

// SetDefaultIngressTlsCertificate sets field value
func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetDefaultIngressTlsCertificate(v string) {
	o.DefaultIngressTlsCertificate = v
}

func (o VcenterNamespaceManagementClustersNCPClusterNetworkInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pod_cidrs"] = o.PodCidrs
	}
	if true {
		toSerialize["ingress_cidrs"] = o.IngressCidrs
	}
	if true {
		toSerialize["egress_cidrs"] = o.EgressCidrs
	}
	if true {
		toSerialize["cluster_distributed_switch"] = o.ClusterDistributedSwitch
	}
	if true {
		toSerialize["nsx_edge_cluster"] = o.NsxEdgeCluster
	}
	if true {
		toSerialize["default_ingress_tls_certificate"] = o.DefaultIngressTlsCertificate
	}
	return json.Marshal(toSerialize)
}

type NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo struct {
	value *VcenterNamespaceManagementClustersNCPClusterNetworkInfo
	isSet bool
}

func (v NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo) Get() *VcenterNamespaceManagementClustersNCPClusterNetworkInfo {
	return v.value
}

func (v *NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo) Set(val *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo(val *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) *NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo {
	return &NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo{value: val, isSet: true}
}

func (v NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVcenterNamespaceManagementClustersNCPClusterNetworkInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


