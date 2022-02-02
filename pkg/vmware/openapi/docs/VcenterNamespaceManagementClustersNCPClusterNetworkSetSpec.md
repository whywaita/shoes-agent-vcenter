# VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which Kubernetes allocates pod IP addresses. This range should not overlap with those in vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkSetSpec.ingress-cidrs, Clusters.NCPClusterNetworkSetSpec.egress-cidrs, or other services running in the datacenter. A set operation only allows for addition of new CIDR blocks to the existing list. All Pod CIDR blocks must be of at least subnet size /23. | 
**IngressCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. These ranges should not overlap with those in Clusters.NCPClusterNetworkSetSpec.pod-cidrs, vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkSetSpec.egress-cidrs, or other services running in the datacenter. A set operation only allows for addition of new CIDR blocks to the existing list. | 
**EgressCidrs** | [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. These ranges should not overlap with those in Clusters.NCPClusterNetworkSetSpec.pod-cidrs, vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkSetSpec.ingress-cidrs, or other services running in the datacenter. A set operation only allows for addition of new CIDR blocks to the existing list. | 
**DefaultIngressTlsCertificate** | **string** | PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services. | 

## Methods

### NewVcenterNamespaceManagementClustersNCPClusterNetworkSetSpec

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkSetSpec(podCidrs []VcenterNamespaceManagementIpv4Cidr, ingressCidrs []VcenterNamespaceManagementIpv4Cidr, egressCidrs []VcenterNamespaceManagementIpv4Cidr, defaultIngressTlsCertificate string, ) *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec`

NewVcenterNamespaceManagementClustersNCPClusterNetworkSetSpec instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersNCPClusterNetworkSetSpecWithDefaults

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkSetSpecWithDefaults() *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec`

NewVcenterNamespaceManagementClustersNCPClusterNetworkSetSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetPodCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetPodCidrs returns the PodCidrs field if non-nil, zero value otherwise.

### GetPodCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetPodCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetPodCidrsOk returns a tuple with the PodCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) SetPodCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetPodCidrs sets PodCidrs field to given value.


### GetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetIngressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetIngressCidrs returns the IngressCidrs field if non-nil, zero value otherwise.

### GetIngressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetIngressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetIngressCidrsOk returns a tuple with the IngressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) SetIngressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetIngressCidrs sets IngressCidrs field to given value.


### GetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetEgressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetEgressCidrs returns the EgressCidrs field if non-nil, zero value otherwise.

### GetEgressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetEgressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetEgressCidrsOk returns a tuple with the EgressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) SetEgressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetEgressCidrs sets EgressCidrs field to given value.


### GetDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetDefaultIngressTlsCertificate() string`

GetDefaultIngressTlsCertificate returns the DefaultIngressTlsCertificate field if non-nil, zero value otherwise.

### GetDefaultIngressTlsCertificateOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) GetDefaultIngressTlsCertificateOk() (*string, bool)`

GetDefaultIngressTlsCertificateOk returns a tuple with the DefaultIngressTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec) SetDefaultIngressTlsCertificate(v string)`

SetDefaultIngressTlsCertificate sets DefaultIngressTlsCertificate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


