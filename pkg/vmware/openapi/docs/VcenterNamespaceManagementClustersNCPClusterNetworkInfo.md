# VcenterNamespaceManagementClustersNCPClusterNetworkInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which Kubernetes allocates pod IP addresses. | 
**IngressCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. | 
**EgressCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. | 
**ClusterDistributedSwitch** | **string** | vSphere Distributed Switch used to connect this cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vSphereDistributedSwitch. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vSphereDistributedSwitch. | 
**NsxEdgeCluster** | **string** | NSX Edge Cluster to be used for Kubernetes Services of type LoadBalancer, Kubernetes Ingresses, and NSX SNAT. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: NSXEdgeCluster. When operations return a value of this structure as a result, the field will be an identifier for the resource type: NSXEdgeCluster. | 
**DefaultIngressTlsCertificate** | **string** | PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services. | 

## Methods

### NewVcenterNamespaceManagementClustersNCPClusterNetworkInfo

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkInfo(podCidrs []VcenterNamespaceManagementIpv4Cidr, ingressCidrs []VcenterNamespaceManagementIpv4Cidr, egressCidrs []VcenterNamespaceManagementIpv4Cidr, clusterDistributedSwitch string, nsxEdgeCluster string, defaultIngressTlsCertificate string, ) *VcenterNamespaceManagementClustersNCPClusterNetworkInfo`

NewVcenterNamespaceManagementClustersNCPClusterNetworkInfo instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersNCPClusterNetworkInfoWithDefaults

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkInfoWithDefaults() *VcenterNamespaceManagementClustersNCPClusterNetworkInfo`

NewVcenterNamespaceManagementClustersNCPClusterNetworkInfoWithDefaults instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetPodCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetPodCidrs returns the PodCidrs field if non-nil, zero value otherwise.

### GetPodCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetPodCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetPodCidrsOk returns a tuple with the PodCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetPodCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetPodCidrs sets PodCidrs field to given value.


### GetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetIngressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetIngressCidrs returns the IngressCidrs field if non-nil, zero value otherwise.

### GetIngressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetIngressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetIngressCidrsOk returns a tuple with the IngressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetIngressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetIngressCidrs sets IngressCidrs field to given value.


### GetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetEgressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetEgressCidrs returns the EgressCidrs field if non-nil, zero value otherwise.

### GetEgressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetEgressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetEgressCidrsOk returns a tuple with the EgressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetEgressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetEgressCidrs sets EgressCidrs field to given value.


### GetClusterDistributedSwitch

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetClusterDistributedSwitch() string`

GetClusterDistributedSwitch returns the ClusterDistributedSwitch field if non-nil, zero value otherwise.

### GetClusterDistributedSwitchOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetClusterDistributedSwitchOk() (*string, bool)`

GetClusterDistributedSwitchOk returns a tuple with the ClusterDistributedSwitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterDistributedSwitch

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetClusterDistributedSwitch(v string)`

SetClusterDistributedSwitch sets ClusterDistributedSwitch field to given value.


### GetNsxEdgeCluster

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetNsxEdgeCluster() string`

GetNsxEdgeCluster returns the NsxEdgeCluster field if non-nil, zero value otherwise.

### GetNsxEdgeClusterOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetNsxEdgeClusterOk() (*string, bool)`

GetNsxEdgeClusterOk returns a tuple with the NsxEdgeCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsxEdgeCluster

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetNsxEdgeCluster(v string)`

SetNsxEdgeCluster sets NsxEdgeCluster field to given value.


### GetDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetDefaultIngressTlsCertificate() string`

GetDefaultIngressTlsCertificate returns the DefaultIngressTlsCertificate field if non-nil, zero value otherwise.

### GetDefaultIngressTlsCertificateOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) GetDefaultIngressTlsCertificateOk() (*string, bool)`

GetDefaultIngressTlsCertificateOk returns a tuple with the DefaultIngressTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkInfo) SetDefaultIngressTlsCertificate(v string)`

SetDefaultIngressTlsCertificate sets DefaultIngressTlsCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


