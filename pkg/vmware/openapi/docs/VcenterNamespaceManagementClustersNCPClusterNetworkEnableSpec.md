# VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which Kubernetes allocates pod IP addresses. This range should not overlap with those in vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkEnableSpec.ingress-cidrs, Clusters.NCPClusterNetworkEnableSpec.egress-cidrs, or other services running in the datacenter. All Pod CIDR blocks must be of at least subnet size /23. | 
**IngressCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. These ranges should not overlap with those in Clusters.NCPClusterNetworkEnableSpec.pod-cidrs, vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkEnableSpec.egress-cidrs, or other services running in the datacenter. | 
**EgressCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. These ranges should not overlap with those in Clusters.NCPClusterNetworkEnableSpec.pod-cidrs, vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkEnableSpec.ingress-cidrs, or other services running in the datacenter. | 
**ClusterDistributedSwitch** | Pointer to **string** | vSphere Distributed Switch used to connect this cluster. This field is required when configuring a cluster that uses NSX-T. If unset and using NSXe, the system will choose a suitable vSphere Distributed Switch. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vSphereDistributedSwitch. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vSphereDistributedSwitch. | [optional] 
**NsxEdgeCluster** | Pointer to **string** | NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT. This field is required when configuring a cluster that uses NSX-T. If unset and using NSXe, the system will choose a suitable NSX Edge Cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: NSXEdgeCluster. When operations return a value of this structure as a result, the field will be an identifier for the resource type: NSXEdgeCluster. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec(podCidrs []VcenterNamespaceManagementIpv4Cidr, ingressCidrs []VcenterNamespaceManagementIpv4Cidr, egressCidrs []VcenterNamespaceManagementIpv4Cidr, ) *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec`

NewVcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersNCPClusterNetworkEnableSpecWithDefaults

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkEnableSpecWithDefaults() *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec`

NewVcenterNamespaceManagementClustersNCPClusterNetworkEnableSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetPodCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetPodCidrs returns the PodCidrs field if non-nil, zero value otherwise.

### GetPodCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetPodCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetPodCidrsOk returns a tuple with the PodCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) SetPodCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetPodCidrs sets PodCidrs field to given value.


### GetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetIngressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetIngressCidrs returns the IngressCidrs field if non-nil, zero value otherwise.

### GetIngressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetIngressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetIngressCidrsOk returns a tuple with the IngressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) SetIngressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetIngressCidrs sets IngressCidrs field to given value.


### GetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetEgressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetEgressCidrs returns the EgressCidrs field if non-nil, zero value otherwise.

### GetEgressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetEgressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetEgressCidrsOk returns a tuple with the EgressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) SetEgressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetEgressCidrs sets EgressCidrs field to given value.


### GetClusterDistributedSwitch

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetClusterDistributedSwitch() string`

GetClusterDistributedSwitch returns the ClusterDistributedSwitch field if non-nil, zero value otherwise.

### GetClusterDistributedSwitchOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetClusterDistributedSwitchOk() (*string, bool)`

GetClusterDistributedSwitchOk returns a tuple with the ClusterDistributedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDistributedSwitch

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) SetClusterDistributedSwitch(v string)`

SetClusterDistributedSwitch sets ClusterDistributedSwitch field to given value.

### HasClusterDistributedSwitch

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) HasClusterDistributedSwitch() bool`

HasClusterDistributedSwitch returns a boolean if a field has been set.

### GetNsxEdgeCluster

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetNsxEdgeCluster() string`

GetNsxEdgeCluster returns the NsxEdgeCluster field if non-nil, zero value otherwise.

### GetNsxEdgeClusterOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) GetNsxEdgeClusterOk() (*string, bool)`

GetNsxEdgeClusterOk returns a tuple with the NsxEdgeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsxEdgeCluster

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) SetNsxEdgeCluster(v string)`

SetNsxEdgeCluster sets NsxEdgeCluster field to given value.

### HasNsxEdgeCluster

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec) HasNsxEdgeCluster() bool`

HasNsxEdgeCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


