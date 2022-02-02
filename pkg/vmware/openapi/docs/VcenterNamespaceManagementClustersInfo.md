# VcenterNamespaceManagementClustersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeHint** | Pointer to [**VcenterNamespaceManagementSizingHint**](VcenterNamespaceManagementSizingHint.md) |  | [optional] 
**StatSummary** | [**VcenterNamespaceManagementClustersStats**](VcenterNamespaceManagementClustersStats.md) |  | 
**ConfigStatus** | [**VcenterNamespaceManagementClustersConfigStatus**](VcenterNamespaceManagementClustersConfigStatus.md) |  | 
**Messages** | [**[]VcenterNamespaceManagementClustersMessage**](VcenterNamespaceManagementClustersMessage.md) | Current set of messages associated with the object. | 
**KubernetesStatus** | [**VcenterNamespaceManagementClustersKubernetesStatus**](VcenterNamespaceManagementClustersKubernetesStatus.md) |  | 
**KubernetesStatusMessages** | [**[]VcenterNamespaceManagementClustersMessage**](VcenterNamespaceManagementClustersMessage.md) | Current set of messages associated with the object. | 
**ApiServerManagementEndpoint** | **string** | Kubernetes API Server IP address on the management network. This is a floating IP and assigned to one of the control plane VMs on the management network. This endpoint is used by vSphere components. | 
**ApiServerClusterEndpoint** | **string** | Kubernetes API Server IP address via cluster network. This is the IP address of the Kubernetes LoadBalancer type service fronting the apiservers. This endpoint is the one configured in kubeconfig after login, and used for most human and application interaction with Kubernetes. | 
**ApiServers** | **[]string** | Identifier of the Kubernetes API servers. These are the IP addresses of the VM instances for the Kubernetes control plane on the management network. | 
**TlsManagementEndpointCertificate** | Pointer to **string** | PEM-encoded x509 certificate used by TLS endpoint on Kubernetes API servers when accessed from the management network, e.g. from ESX servers or VCSA. This certificate is only valid for use with the apiServerManagementEndpoint. | [optional] 
**TlsEndpointCertificate** | Pointer to **string** | PEM-encoded x509 certificate used by TLS endpoint on Kubernetes API servers when accessed via the load balancer, e.g. devops user on corporate network. This certificate is only valid for use with the apiServerClusterEndpoint. | [optional] 
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**NcpClusterNetworkInfo** | Pointer to [**VcenterNamespaceManagementClustersNCPClusterNetworkInfo**](VcenterNamespaceManagementClustersNCPClusterNetworkInfo.md) |  | [optional] 
**WorkloadNetworks** | Pointer to [**VcenterNamespaceManagementClustersWorkloadNetworksInfo**](VcenterNamespaceManagementClustersWorkloadNetworksInfo.md) |  | [optional] 
**WorkloadNtpServers** | Pointer to **[]string** | Information about NTP server DNS names or IP addresses to use for workloads such as Tanzu Kubernetes Grid VMs, specified in order of preference. This field is optional because it was added in a newer version than its parent node. | [optional] 
**LoadBalancers** | Pointer to [**[]VcenterNamespaceManagementLoadBalancersInfo**](VcenterNamespaceManagementLoadBalancersInfo.md) | Information related to the Load balancer used for provisioning virtual servers in the namespace. This field is optional and it is only relevant when the value of Clusters.Info.network-provider is VSPHERE_NETWORK. | [optional] 
**ServiceCidr** | [**VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) |  | 
**MasterManagementNetwork** | Pointer to [**VcenterNamespaceManagementClustersNetworkSpec**](VcenterNamespaceManagementClustersNetworkSpec.md) |  | [optional] 
**MasterDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. | [optional] 
**WorkerDNS** | Pointer to **[]string** | List of DNS server IP addresses to use for pods that execute on the worker nodes (which are native pods on ESXi hosts in the vSphere Namespaces Supervisor). | [optional] 
**MasterStoragePolicy** | Pointer to **string** | Identifier of storage policy associated with Kubernetes API server. This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | [optional] 
**EphemeralStoragePolicy** | Pointer to **string** | Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. This field is optional because it was added in a newer version than its parent node. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | [optional] 
**CnsFileConfig** | Pointer to [**VcenterNamespaceManagementCNSFileConfig**](VcenterNamespaceManagementCNSFileConfig.md) |  | [optional] 
**LoginBanner** | Pointer to **string** | Disclaimer to be displayed prior to login via the Kubectl plugin. If unset, just skip it. | [optional] 
**MasterDNSNames** | Pointer to **[]string** | List of additional DNS names to associate with the Kubernetes API server. These DNS names are embedded in the TLS certificate presented by the API server. If unset, no additional DNS names are embedded in the TLS certificate. | [optional] 
**ImageStorage** | Pointer to [**VcenterNamespaceManagementClustersImageStorageSpec**](VcenterNamespaceManagementClustersImageStorageSpec.md) |  | [optional] 
**DefaultImageRegistry** | Pointer to [**VcenterNamespaceManagementClustersImageRegistry**](VcenterNamespaceManagementClustersImageRegistry.md) |  | [optional] 
**DefaultImageRepository** | Pointer to **string** | Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. If unset, defaults to Docker Hub official repository in case of Docker Hub image registry, otherwise defaults to empty string. | [optional] 
**MasterDNSSearchDomains** | Pointer to **[]string** | List of domains (for example \&quot;vmware.com\&quot;) to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. | [optional] 
**MasterNTPServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. If unset, VMware Tools based time synchronization is enabled. | [optional] 
**DefaultKubernetesServiceContentLibrary** | Pointer to **string** | Identifier of the Content Library which holds the VM Images for vSphere Kubernetes Service. This Content Library should be subscribed to VMware&#39;s hosted vSphere Kubernetes Service Repository. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: content.Library. When operations return a value of this structure as a result, the field will be an identifier for the resource type: content.Library. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersInfo

`func NewVcenterNamespaceManagementClustersInfo(statSummary VcenterNamespaceManagementClustersStats, configStatus VcenterNamespaceManagementClustersConfigStatus, messages []VcenterNamespaceManagementClustersMessage, kubernetesStatus VcenterNamespaceManagementClustersKubernetesStatus, kubernetesStatusMessages []VcenterNamespaceManagementClustersMessage, apiServerManagementEndpoint string, apiServerClusterEndpoint string, apiServers []string, networkProvider VcenterNamespaceManagementClustersNetworkProvider, serviceCidr VcenterNamespaceManagementIpv4Cidr, ) *VcenterNamespaceManagementClustersInfo`

NewVcenterNamespaceManagementClustersInfo instantiates a new VcenterNamespaceManagementClustersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersInfoWithDefaults

`func NewVcenterNamespaceManagementClustersInfoWithDefaults() *VcenterNamespaceManagementClustersInfo`

NewVcenterNamespaceManagementClustersInfoWithDefaults instantiates a new VcenterNamespaceManagementClustersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeHint

`func (o *VcenterNamespaceManagementClustersInfo) GetSizeHint() VcenterNamespaceManagementSizingHint`

GetSizeHint returns the SizeHint field if non-nil, zero value otherwise.

### GetSizeHintOk

`func (o *VcenterNamespaceManagementClustersInfo) GetSizeHintOk() (*VcenterNamespaceManagementSizingHint, bool)`

GetSizeHintOk returns a tuple with the SizeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHint

`func (o *VcenterNamespaceManagementClustersInfo) SetSizeHint(v VcenterNamespaceManagementSizingHint)`

SetSizeHint sets SizeHint field to given value.

### HasSizeHint

`func (o *VcenterNamespaceManagementClustersInfo) HasSizeHint() bool`

HasSizeHint returns a boolean if a field has been set.

### GetStatSummary

`func (o *VcenterNamespaceManagementClustersInfo) GetStatSummary() VcenterNamespaceManagementClustersStats`

GetStatSummary returns the StatSummary field if non-nil, zero value otherwise.

### GetStatSummaryOk

`func (o *VcenterNamespaceManagementClustersInfo) GetStatSummaryOk() (*VcenterNamespaceManagementClustersStats, bool)`

GetStatSummaryOk returns a tuple with the StatSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatSummary

`func (o *VcenterNamespaceManagementClustersInfo) SetStatSummary(v VcenterNamespaceManagementClustersStats)`

SetStatSummary sets StatSummary field to given value.


### GetConfigStatus

`func (o *VcenterNamespaceManagementClustersInfo) GetConfigStatus() VcenterNamespaceManagementClustersConfigStatus`

GetConfigStatus returns the ConfigStatus field if non-nil, zero value otherwise.

### GetConfigStatusOk

`func (o *VcenterNamespaceManagementClustersInfo) GetConfigStatusOk() (*VcenterNamespaceManagementClustersConfigStatus, bool)`

GetConfigStatusOk returns a tuple with the ConfigStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStatus

`func (o *VcenterNamespaceManagementClustersInfo) SetConfigStatus(v VcenterNamespaceManagementClustersConfigStatus)`

SetConfigStatus sets ConfigStatus field to given value.


### GetMessages

`func (o *VcenterNamespaceManagementClustersInfo) GetMessages() []VcenterNamespaceManagementClustersMessage`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMessagesOk() (*[]VcenterNamespaceManagementClustersMessage, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *VcenterNamespaceManagementClustersInfo) SetMessages(v []VcenterNamespaceManagementClustersMessage)`

SetMessages sets Messages field to given value.


### GetKubernetesStatus

`func (o *VcenterNamespaceManagementClustersInfo) GetKubernetesStatus() VcenterNamespaceManagementClustersKubernetesStatus`

GetKubernetesStatus returns the KubernetesStatus field if non-nil, zero value otherwise.

### GetKubernetesStatusOk

`func (o *VcenterNamespaceManagementClustersInfo) GetKubernetesStatusOk() (*VcenterNamespaceManagementClustersKubernetesStatus, bool)`

GetKubernetesStatusOk returns a tuple with the KubernetesStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesStatus

`func (o *VcenterNamespaceManagementClustersInfo) SetKubernetesStatus(v VcenterNamespaceManagementClustersKubernetesStatus)`

SetKubernetesStatus sets KubernetesStatus field to given value.


### GetKubernetesStatusMessages

`func (o *VcenterNamespaceManagementClustersInfo) GetKubernetesStatusMessages() []VcenterNamespaceManagementClustersMessage`

GetKubernetesStatusMessages returns the KubernetesStatusMessages field if non-nil, zero value otherwise.

### GetKubernetesStatusMessagesOk

`func (o *VcenterNamespaceManagementClustersInfo) GetKubernetesStatusMessagesOk() (*[]VcenterNamespaceManagementClustersMessage, bool)`

GetKubernetesStatusMessagesOk returns a tuple with the KubernetesStatusMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKubernetesStatusMessages

`func (o *VcenterNamespaceManagementClustersInfo) SetKubernetesStatusMessages(v []VcenterNamespaceManagementClustersMessage)`

SetKubernetesStatusMessages sets KubernetesStatusMessages field to given value.


### GetApiServerManagementEndpoint

`func (o *VcenterNamespaceManagementClustersInfo) GetApiServerManagementEndpoint() string`

GetApiServerManagementEndpoint returns the ApiServerManagementEndpoint field if non-nil, zero value otherwise.

### GetApiServerManagementEndpointOk

`func (o *VcenterNamespaceManagementClustersInfo) GetApiServerManagementEndpointOk() (*string, bool)`

GetApiServerManagementEndpointOk returns a tuple with the ApiServerManagementEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerManagementEndpoint

`func (o *VcenterNamespaceManagementClustersInfo) SetApiServerManagementEndpoint(v string)`

SetApiServerManagementEndpoint sets ApiServerManagementEndpoint field to given value.


### GetApiServerClusterEndpoint

`func (o *VcenterNamespaceManagementClustersInfo) GetApiServerClusterEndpoint() string`

GetApiServerClusterEndpoint returns the ApiServerClusterEndpoint field if non-nil, zero value otherwise.

### GetApiServerClusterEndpointOk

`func (o *VcenterNamespaceManagementClustersInfo) GetApiServerClusterEndpointOk() (*string, bool)`

GetApiServerClusterEndpointOk returns a tuple with the ApiServerClusterEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServerClusterEndpoint

`func (o *VcenterNamespaceManagementClustersInfo) SetApiServerClusterEndpoint(v string)`

SetApiServerClusterEndpoint sets ApiServerClusterEndpoint field to given value.


### GetApiServers

`func (o *VcenterNamespaceManagementClustersInfo) GetApiServers() []string`

GetApiServers returns the ApiServers field if non-nil, zero value otherwise.

### GetApiServersOk

`func (o *VcenterNamespaceManagementClustersInfo) GetApiServersOk() (*[]string, bool)`

GetApiServersOk returns a tuple with the ApiServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiServers

`func (o *VcenterNamespaceManagementClustersInfo) SetApiServers(v []string)`

SetApiServers sets ApiServers field to given value.


### GetTlsManagementEndpointCertificate

`func (o *VcenterNamespaceManagementClustersInfo) GetTlsManagementEndpointCertificate() string`

GetTlsManagementEndpointCertificate returns the TlsManagementEndpointCertificate field if non-nil, zero value otherwise.

### GetTlsManagementEndpointCertificateOk

`func (o *VcenterNamespaceManagementClustersInfo) GetTlsManagementEndpointCertificateOk() (*string, bool)`

GetTlsManagementEndpointCertificateOk returns a tuple with the TlsManagementEndpointCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsManagementEndpointCertificate

`func (o *VcenterNamespaceManagementClustersInfo) SetTlsManagementEndpointCertificate(v string)`

SetTlsManagementEndpointCertificate sets TlsManagementEndpointCertificate field to given value.

### HasTlsManagementEndpointCertificate

`func (o *VcenterNamespaceManagementClustersInfo) HasTlsManagementEndpointCertificate() bool`

HasTlsManagementEndpointCertificate returns a boolean if a field has been set.

### GetTlsEndpointCertificate

`func (o *VcenterNamespaceManagementClustersInfo) GetTlsEndpointCertificate() string`

GetTlsEndpointCertificate returns the TlsEndpointCertificate field if non-nil, zero value otherwise.

### GetTlsEndpointCertificateOk

`func (o *VcenterNamespaceManagementClustersInfo) GetTlsEndpointCertificateOk() (*string, bool)`

GetTlsEndpointCertificateOk returns a tuple with the TlsEndpointCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEndpointCertificate

`func (o *VcenterNamespaceManagementClustersInfo) SetTlsEndpointCertificate(v string)`

SetTlsEndpointCertificate sets TlsEndpointCertificate field to given value.

### HasTlsEndpointCertificate

`func (o *VcenterNamespaceManagementClustersInfo) HasTlsEndpointCertificate() bool`

HasTlsEndpointCertificate returns a boolean if a field has been set.

### GetNetworkProvider

`func (o *VcenterNamespaceManagementClustersInfo) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementClustersInfo) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementClustersInfo) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetNcpClusterNetworkInfo

`func (o *VcenterNamespaceManagementClustersInfo) GetNcpClusterNetworkInfo() VcenterNamespaceManagementClustersNCPClusterNetworkInfo`

GetNcpClusterNetworkInfo returns the NcpClusterNetworkInfo field if non-nil, zero value otherwise.

### GetNcpClusterNetworkInfoOk

`func (o *VcenterNamespaceManagementClustersInfo) GetNcpClusterNetworkInfoOk() (*VcenterNamespaceManagementClustersNCPClusterNetworkInfo, bool)`

GetNcpClusterNetworkInfoOk returns a tuple with the NcpClusterNetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcpClusterNetworkInfo

`func (o *VcenterNamespaceManagementClustersInfo) SetNcpClusterNetworkInfo(v VcenterNamespaceManagementClustersNCPClusterNetworkInfo)`

SetNcpClusterNetworkInfo sets NcpClusterNetworkInfo field to given value.

### HasNcpClusterNetworkInfo

`func (o *VcenterNamespaceManagementClustersInfo) HasNcpClusterNetworkInfo() bool`

HasNcpClusterNetworkInfo returns a boolean if a field has been set.

### GetWorkloadNetworks

`func (o *VcenterNamespaceManagementClustersInfo) GetWorkloadNetworks() VcenterNamespaceManagementClustersWorkloadNetworksInfo`

GetWorkloadNetworks returns the WorkloadNetworks field if non-nil, zero value otherwise.

### GetWorkloadNetworksOk

`func (o *VcenterNamespaceManagementClustersInfo) GetWorkloadNetworksOk() (*VcenterNamespaceManagementClustersWorkloadNetworksInfo, bool)`

GetWorkloadNetworksOk returns a tuple with the WorkloadNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNetworks

`func (o *VcenterNamespaceManagementClustersInfo) SetWorkloadNetworks(v VcenterNamespaceManagementClustersWorkloadNetworksInfo)`

SetWorkloadNetworks sets WorkloadNetworks field to given value.

### HasWorkloadNetworks

`func (o *VcenterNamespaceManagementClustersInfo) HasWorkloadNetworks() bool`

HasWorkloadNetworks returns a boolean if a field has been set.

### GetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersInfo) GetWorkloadNtpServers() []string`

GetWorkloadNtpServers returns the WorkloadNtpServers field if non-nil, zero value otherwise.

### GetWorkloadNtpServersOk

`func (o *VcenterNamespaceManagementClustersInfo) GetWorkloadNtpServersOk() (*[]string, bool)`

GetWorkloadNtpServersOk returns a tuple with the WorkloadNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersInfo) SetWorkloadNtpServers(v []string)`

SetWorkloadNtpServers sets WorkloadNtpServers field to given value.

### HasWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersInfo) HasWorkloadNtpServers() bool`

HasWorkloadNtpServers returns a boolean if a field has been set.

### GetLoadBalancers

`func (o *VcenterNamespaceManagementClustersInfo) GetLoadBalancers() []VcenterNamespaceManagementLoadBalancersInfo`

GetLoadBalancers returns the LoadBalancers field if non-nil, zero value otherwise.

### GetLoadBalancersOk

`func (o *VcenterNamespaceManagementClustersInfo) GetLoadBalancersOk() (*[]VcenterNamespaceManagementLoadBalancersInfo, bool)`

GetLoadBalancersOk returns a tuple with the LoadBalancers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancers

`func (o *VcenterNamespaceManagementClustersInfo) SetLoadBalancers(v []VcenterNamespaceManagementLoadBalancersInfo)`

SetLoadBalancers sets LoadBalancers field to given value.

### HasLoadBalancers

`func (o *VcenterNamespaceManagementClustersInfo) HasLoadBalancers() bool`

HasLoadBalancers returns a boolean if a field has been set.

### GetServiceCidr

`func (o *VcenterNamespaceManagementClustersInfo) GetServiceCidr() VcenterNamespaceManagementIpv4Cidr`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *VcenterNamespaceManagementClustersInfo) GetServiceCidrOk() (*VcenterNamespaceManagementIpv4Cidr, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *VcenterNamespaceManagementClustersInfo) SetServiceCidr(v VcenterNamespaceManagementIpv4Cidr)`

SetServiceCidr sets ServiceCidr field to given value.


### GetMasterManagementNetwork

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterManagementNetwork() VcenterNamespaceManagementClustersNetworkSpec`

GetMasterManagementNetwork returns the MasterManagementNetwork field if non-nil, zero value otherwise.

### GetMasterManagementNetworkOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterManagementNetworkOk() (*VcenterNamespaceManagementClustersNetworkSpec, bool)`

GetMasterManagementNetworkOk returns a tuple with the MasterManagementNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterManagementNetwork

`func (o *VcenterNamespaceManagementClustersInfo) SetMasterManagementNetwork(v VcenterNamespaceManagementClustersNetworkSpec)`

SetMasterManagementNetwork sets MasterManagementNetwork field to given value.

### HasMasterManagementNetwork

`func (o *VcenterNamespaceManagementClustersInfo) HasMasterManagementNetwork() bool`

HasMasterManagementNetwork returns a boolean if a field has been set.

### GetMasterDNS

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterDNS() []string`

GetMasterDNS returns the MasterDNS field if non-nil, zero value otherwise.

### GetMasterDNSOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterDNSOk() (*[]string, bool)`

GetMasterDNSOk returns a tuple with the MasterDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNS

`func (o *VcenterNamespaceManagementClustersInfo) SetMasterDNS(v []string)`

SetMasterDNS sets MasterDNS field to given value.

### HasMasterDNS

`func (o *VcenterNamespaceManagementClustersInfo) HasMasterDNS() bool`

HasMasterDNS returns a boolean if a field has been set.

### GetWorkerDNS

`func (o *VcenterNamespaceManagementClustersInfo) GetWorkerDNS() []string`

GetWorkerDNS returns the WorkerDNS field if non-nil, zero value otherwise.

### GetWorkerDNSOk

`func (o *VcenterNamespaceManagementClustersInfo) GetWorkerDNSOk() (*[]string, bool)`

GetWorkerDNSOk returns a tuple with the WorkerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerDNS

`func (o *VcenterNamespaceManagementClustersInfo) SetWorkerDNS(v []string)`

SetWorkerDNS sets WorkerDNS field to given value.

### HasWorkerDNS

`func (o *VcenterNamespaceManagementClustersInfo) HasWorkerDNS() bool`

HasWorkerDNS returns a boolean if a field has been set.

### GetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterStoragePolicy() string`

GetMasterStoragePolicy returns the MasterStoragePolicy field if non-nil, zero value otherwise.

### GetMasterStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterStoragePolicyOk() (*string, bool)`

GetMasterStoragePolicyOk returns a tuple with the MasterStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersInfo) SetMasterStoragePolicy(v string)`

SetMasterStoragePolicy sets MasterStoragePolicy field to given value.

### HasMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersInfo) HasMasterStoragePolicy() bool`

HasMasterStoragePolicy returns a boolean if a field has been set.

### GetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersInfo) GetEphemeralStoragePolicy() string`

GetEphemeralStoragePolicy returns the EphemeralStoragePolicy field if non-nil, zero value otherwise.

### GetEphemeralStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersInfo) GetEphemeralStoragePolicyOk() (*string, bool)`

GetEphemeralStoragePolicyOk returns a tuple with the EphemeralStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersInfo) SetEphemeralStoragePolicy(v string)`

SetEphemeralStoragePolicy sets EphemeralStoragePolicy field to given value.

### HasEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersInfo) HasEphemeralStoragePolicy() bool`

HasEphemeralStoragePolicy returns a boolean if a field has been set.

### GetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersInfo) GetCnsFileConfig() VcenterNamespaceManagementCNSFileConfig`

GetCnsFileConfig returns the CnsFileConfig field if non-nil, zero value otherwise.

### GetCnsFileConfigOk

`func (o *VcenterNamespaceManagementClustersInfo) GetCnsFileConfigOk() (*VcenterNamespaceManagementCNSFileConfig, bool)`

GetCnsFileConfigOk returns a tuple with the CnsFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersInfo) SetCnsFileConfig(v VcenterNamespaceManagementCNSFileConfig)`

SetCnsFileConfig sets CnsFileConfig field to given value.

### HasCnsFileConfig

`func (o *VcenterNamespaceManagementClustersInfo) HasCnsFileConfig() bool`

HasCnsFileConfig returns a boolean if a field has been set.

### GetLoginBanner

`func (o *VcenterNamespaceManagementClustersInfo) GetLoginBanner() string`

GetLoginBanner returns the LoginBanner field if non-nil, zero value otherwise.

### GetLoginBannerOk

`func (o *VcenterNamespaceManagementClustersInfo) GetLoginBannerOk() (*string, bool)`

GetLoginBannerOk returns a tuple with the LoginBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBanner

`func (o *VcenterNamespaceManagementClustersInfo) SetLoginBanner(v string)`

SetLoginBanner sets LoginBanner field to given value.

### HasLoginBanner

`func (o *VcenterNamespaceManagementClustersInfo) HasLoginBanner() bool`

HasLoginBanner returns a boolean if a field has been set.

### GetMasterDNSNames

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterDNSNames() []string`

GetMasterDNSNames returns the MasterDNSNames field if non-nil, zero value otherwise.

### GetMasterDNSNamesOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterDNSNamesOk() (*[]string, bool)`

GetMasterDNSNamesOk returns a tuple with the MasterDNSNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNSNames

`func (o *VcenterNamespaceManagementClustersInfo) SetMasterDNSNames(v []string)`

SetMasterDNSNames sets MasterDNSNames field to given value.

### HasMasterDNSNames

`func (o *VcenterNamespaceManagementClustersInfo) HasMasterDNSNames() bool`

HasMasterDNSNames returns a boolean if a field has been set.

### GetImageStorage

`func (o *VcenterNamespaceManagementClustersInfo) GetImageStorage() VcenterNamespaceManagementClustersImageStorageSpec`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *VcenterNamespaceManagementClustersInfo) GetImageStorageOk() (*VcenterNamespaceManagementClustersImageStorageSpec, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *VcenterNamespaceManagementClustersInfo) SetImageStorage(v VcenterNamespaceManagementClustersImageStorageSpec)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *VcenterNamespaceManagementClustersInfo) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersInfo) GetDefaultImageRegistry() VcenterNamespaceManagementClustersImageRegistry`

GetDefaultImageRegistry returns the DefaultImageRegistry field if non-nil, zero value otherwise.

### GetDefaultImageRegistryOk

`func (o *VcenterNamespaceManagementClustersInfo) GetDefaultImageRegistryOk() (*VcenterNamespaceManagementClustersImageRegistry, bool)`

GetDefaultImageRegistryOk returns a tuple with the DefaultImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersInfo) SetDefaultImageRegistry(v VcenterNamespaceManagementClustersImageRegistry)`

SetDefaultImageRegistry sets DefaultImageRegistry field to given value.

### HasDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersInfo) HasDefaultImageRegistry() bool`

HasDefaultImageRegistry returns a boolean if a field has been set.

### GetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersInfo) GetDefaultImageRepository() string`

GetDefaultImageRepository returns the DefaultImageRepository field if non-nil, zero value otherwise.

### GetDefaultImageRepositoryOk

`func (o *VcenterNamespaceManagementClustersInfo) GetDefaultImageRepositoryOk() (*string, bool)`

GetDefaultImageRepositoryOk returns a tuple with the DefaultImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersInfo) SetDefaultImageRepository(v string)`

SetDefaultImageRepository sets DefaultImageRepository field to given value.

### HasDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersInfo) HasDefaultImageRepository() bool`

HasDefaultImageRepository returns a boolean if a field has been set.

### GetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterDNSSearchDomains() []string`

GetMasterDNSSearchDomains returns the MasterDNSSearchDomains field if non-nil, zero value otherwise.

### GetMasterDNSSearchDomainsOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterDNSSearchDomainsOk() (*[]string, bool)`

GetMasterDNSSearchDomainsOk returns a tuple with the MasterDNSSearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersInfo) SetMasterDNSSearchDomains(v []string)`

SetMasterDNSSearchDomains sets MasterDNSSearchDomains field to given value.

### HasMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersInfo) HasMasterDNSSearchDomains() bool`

HasMasterDNSSearchDomains returns a boolean if a field has been set.

### GetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterNTPServers() []string`

GetMasterNTPServers returns the MasterNTPServers field if non-nil, zero value otherwise.

### GetMasterNTPServersOk

`func (o *VcenterNamespaceManagementClustersInfo) GetMasterNTPServersOk() (*[]string, bool)`

GetMasterNTPServersOk returns a tuple with the MasterNTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersInfo) SetMasterNTPServers(v []string)`

SetMasterNTPServers sets MasterNTPServers field to given value.

### HasMasterNTPServers

`func (o *VcenterNamespaceManagementClustersInfo) HasMasterNTPServers() bool`

HasMasterNTPServers returns a boolean if a field has been set.

### GetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersInfo) GetDefaultKubernetesServiceContentLibrary() string`

GetDefaultKubernetesServiceContentLibrary returns the DefaultKubernetesServiceContentLibrary field if non-nil, zero value otherwise.

### GetDefaultKubernetesServiceContentLibraryOk

`func (o *VcenterNamespaceManagementClustersInfo) GetDefaultKubernetesServiceContentLibraryOk() (*string, bool)`

GetDefaultKubernetesServiceContentLibraryOk returns a tuple with the DefaultKubernetesServiceContentLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersInfo) SetDefaultKubernetesServiceContentLibrary(v string)`

SetDefaultKubernetesServiceContentLibrary sets DefaultKubernetesServiceContentLibrary field to given value.

### HasDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersInfo) HasDefaultKubernetesServiceContentLibrary() bool`

HasDefaultKubernetesServiceContentLibrary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


