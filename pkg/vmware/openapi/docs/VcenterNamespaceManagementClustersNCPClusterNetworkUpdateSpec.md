# VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidrs** | Pointer to [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which Kubernetes allocates pod IP addresses. This range should not overlap with those in vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkUpdateSpec.ingress-cidrs, Clusters.NCPClusterNetworkUpdateSpec.egress-cidrs, or other services running in the datacenter. An update operation only allows for addition of new CIDR blocks to the existing list. All Pod CIDR blocks must be of at least subnet size /23. If unset, CIDRs from which Kubernetes allocates pod IP addresses will not be modified. | [optional] 
**IngressCidrs** | Pointer to [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses for Kubernetes Ingresses and Kubernetes Services of type LoadBalancer. These ranges should not overlap with those in Clusters.NCPClusterNetworkUpdateSpec.pod-cidrs, vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkUpdateSpec.egress-cidrs, or other services running in the datacenter. An update operation only allows for addition of new CIDR blocks to the existing list. If unset, CIDRs from which Kubernetes allocates ingress IP addresses will not be modified. | [optional] 
**EgressCidrs** | Pointer to [**[]VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) | CIDR blocks from which NSX assigns IP addresses used for performing SNAT from container IPs to external IPs. These ranges should not overlap with those in Clusters.NCPClusterNetworkUpdateSpec.pod-cidrs, vcenter.namespace_management.EnableSpec#serviceCidr, Clusters.NCPClusterNetworkUpdateSpec.ingress-cidrs, or other services running in the datacenter. An update operation only allows for addition of new CIDR blocks to the existing list. If unset, CIDR from which Kubernetes allocates egress IP addresses will not be modified. | [optional] 
**DefaultIngressTlsCertificate** | Pointer to **string** | PEM-encoded x509 certificate used by NSX as a default fallback certificate for Kubernetes Ingress services. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec() *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec`

NewVcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpecWithDefaults

`func NewVcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpecWithDefaults() *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec`

NewVcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetPodCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetPodCidrs returns the PodCidrs field if non-nil, zero value otherwise.

### GetPodCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetPodCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetPodCidrsOk returns a tuple with the PodCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) SetPodCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetPodCidrs sets PodCidrs field to given value.

### HasPodCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) HasPodCidrs() bool`

HasPodCidrs returns a boolean if a field has been set.

### GetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetIngressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetIngressCidrs returns the IngressCidrs field if non-nil, zero value otherwise.

### GetIngressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetIngressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetIngressCidrsOk returns a tuple with the IngressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) SetIngressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetIngressCidrs sets IngressCidrs field to given value.

### HasIngressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) HasIngressCidrs() bool`

HasIngressCidrs returns a boolean if a field has been set.

### GetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetEgressCidrs() []VcenterNamespaceManagementIpv4Cidr`

GetEgressCidrs returns the EgressCidrs field if non-nil, zero value otherwise.

### GetEgressCidrsOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetEgressCidrsOk() (*[]VcenterNamespaceManagementIpv4Cidr, bool)`

GetEgressCidrsOk returns a tuple with the EgressCidrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) SetEgressCidrs(v []VcenterNamespaceManagementIpv4Cidr)`

SetEgressCidrs sets EgressCidrs field to given value.

### HasEgressCidrs

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) HasEgressCidrs() bool`

HasEgressCidrs returns a boolean if a field has been set.

### GetDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetDefaultIngressTlsCertificate() string`

GetDefaultIngressTlsCertificate returns the DefaultIngressTlsCertificate field if non-nil, zero value otherwise.

### GetDefaultIngressTlsCertificateOk

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) GetDefaultIngressTlsCertificateOk() (*string, bool)`

GetDefaultIngressTlsCertificateOk returns a tuple with the DefaultIngressTlsCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) SetDefaultIngressTlsCertificate(v string)`

SetDefaultIngressTlsCertificate sets DefaultIngressTlsCertificate field to given value.

### HasDefaultIngressTlsCertificate

`func (o *VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec) HasDefaultIngressTlsCertificate() bool`

HasDefaultIngressTlsCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


