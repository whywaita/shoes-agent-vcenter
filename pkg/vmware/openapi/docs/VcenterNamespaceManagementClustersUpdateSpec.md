# VcenterNamespaceManagementClustersUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeHint** | Pointer to [**VcenterNamespaceManagementSizingHint**](VcenterNamespaceManagementSizingHint.md) |  | [optional] 
**NetworkProvider** | Pointer to [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | [optional] 
**NcpClusterNetworkSpec** | Pointer to [**VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec**](VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec.md) |  | [optional] 
**MasterDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. If set, DNS servers set on Kubernetes API server will be replaced. Otherwise, they will not be modified. | [optional] 
**WorkerDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on the worker nodes, specified in order of preference. If set, DNS servers set on worker nodes will be replaced. Otherwise, they will not be modified. | [optional] 
**MasterDNSSearchDomains** | Pointer to **[]string** | List of domains (for example \&quot;vmware.com\&quot;) to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. If set, DNS search domains on Kubernetes API server will be replaced. Otherwise, they will not be modified. | [optional] 
**MasterNTPServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. If set, NTP servers on Kubernetes API server will be replaced. Otherwise, they will not be modified. | [optional] 
**MasterStoragePolicy** | Pointer to **string** | Identifier of storage policy associated with Kubernetes API server. If unset, storage policy associated with Kubernetes API server will not be modified. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | [optional] 
**EphemeralStoragePolicy** | Pointer to **string** | Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. If unset, storage policy associated with ephemeral disks of all the Kubernetes Pods will not be modified. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | [optional] 
**CnsFileConfig** | Pointer to [**VcenterNamespaceManagementCNSFileConfig**](VcenterNamespaceManagementCNSFileConfig.md) |  | [optional] 
**LoginBanner** | Pointer to **string** | Disclaimer to be displayed prior to login via the Kubectl plugin. If unset, disclaimer to be displayed prior to login via the Kubectl plugin will not be modified. | [optional] 
**ImageStorage** | Pointer to [**VcenterNamespaceManagementClustersImageStorageSpec**](VcenterNamespaceManagementClustersImageStorageSpec.md) |  | [optional] 
**DefaultImageRegistry** | Pointer to [**VcenterNamespaceManagementClustersImageRegistry**](VcenterNamespaceManagementClustersImageRegistry.md) |  | [optional] 
**DefaultImageRepository** | Pointer to **string** | Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. If unset, default image repository will not be modified. | [optional] 
**TlsEndpointCertificate** | Pointer to **string** | Certificate issued for Kubernetes API Server. Certificate used must be created by signing the Certificate Signing Request obtained from vcenter.namespace_management.certificates.Request.create Because a CertificateSigningRequest is created on an existing Namespaces-enabled Cluster, you must use the Clusters.UpdateSpec to specify this tlsEndpointCertificate on an existing Cluster rather than during initially enabling Namespaces on a Cluster. If unset, Kubernetes API Server certificate will not be modified. | [optional] 
**DefaultKubernetesServiceContentLibrary** | Pointer to **string** | Identifier of the Content Library which holds the VM Images for vSphere Kubernetes Service. This Content Library should be subscribed to VMware&#39;s hosted vSphere Kubernetes Service Repository. Modifying or clearing the Content Library identifier will not affect existing vSphere Kubernetes Service clusters. However, upgrades or scale-out of existing clusters may be affected if the new Content Library doesn&#39;t have the necessary VM Images. If unset, the Content Library identifier will not be modified. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: content.Library. When operations return a value of this structure as a result, the field will be an identifier for the resource type: content.Library. | [optional] 
**WorkloadNtpServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use for workloads such as Tanzu Kubernetes Grid VMs, specified in order of preference. If unset, NTP servers for workloads will be unmodified. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersUpdateSpec

`func NewVcenterNamespaceManagementClustersUpdateSpec() *VcenterNamespaceManagementClustersUpdateSpec`

NewVcenterNamespaceManagementClustersUpdateSpec instantiates a new VcenterNamespaceManagementClustersUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersUpdateSpecWithDefaults

`func NewVcenterNamespaceManagementClustersUpdateSpecWithDefaults() *VcenterNamespaceManagementClustersUpdateSpec`

NewVcenterNamespaceManagementClustersUpdateSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeHint

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetSizeHint() VcenterNamespaceManagementSizingHint`

GetSizeHint returns the SizeHint field if non-nil, zero value otherwise.

### GetSizeHintOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetSizeHintOk() (*VcenterNamespaceManagementSizingHint, bool)`

GetSizeHintOk returns a tuple with the SizeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHint

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetSizeHint(v VcenterNamespaceManagementSizingHint)`

SetSizeHint sets SizeHint field to given value.

### HasSizeHint

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasSizeHint() bool`

HasSizeHint returns a boolean if a field has been set.

### GetNetworkProvider

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.

### HasNetworkProvider

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasNetworkProvider() bool`

HasNetworkProvider returns a boolean if a field has been set.

### GetNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetNcpClusterNetworkSpec() VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec`

GetNcpClusterNetworkSpec returns the NcpClusterNetworkSpec field if non-nil, zero value otherwise.

### GetNcpClusterNetworkSpecOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetNcpClusterNetworkSpecOk() (*VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec, bool)`

GetNcpClusterNetworkSpecOk returns a tuple with the NcpClusterNetworkSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetNcpClusterNetworkSpec(v VcenterNamespaceManagementClustersNCPClusterNetworkUpdateSpec)`

SetNcpClusterNetworkSpec sets NcpClusterNetworkSpec field to given value.

### HasNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasNcpClusterNetworkSpec() bool`

HasNcpClusterNetworkSpec returns a boolean if a field has been set.

### GetMasterDNS

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterDNS() []string`

GetMasterDNS returns the MasterDNS field if non-nil, zero value otherwise.

### GetMasterDNSOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterDNSOk() (*[]string, bool)`

GetMasterDNSOk returns a tuple with the MasterDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNS

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetMasterDNS(v []string)`

SetMasterDNS sets MasterDNS field to given value.

### HasMasterDNS

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasMasterDNS() bool`

HasMasterDNS returns a boolean if a field has been set.

### GetWorkerDNS

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetWorkerDNS() []string`

GetWorkerDNS returns the WorkerDNS field if non-nil, zero value otherwise.

### GetWorkerDNSOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetWorkerDNSOk() (*[]string, bool)`

GetWorkerDNSOk returns a tuple with the WorkerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerDNS

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetWorkerDNS(v []string)`

SetWorkerDNS sets WorkerDNS field to given value.

### HasWorkerDNS

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasWorkerDNS() bool`

HasWorkerDNS returns a boolean if a field has been set.

### GetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterDNSSearchDomains() []string`

GetMasterDNSSearchDomains returns the MasterDNSSearchDomains field if non-nil, zero value otherwise.

### GetMasterDNSSearchDomainsOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterDNSSearchDomainsOk() (*[]string, bool)`

GetMasterDNSSearchDomainsOk returns a tuple with the MasterDNSSearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetMasterDNSSearchDomains(v []string)`

SetMasterDNSSearchDomains sets MasterDNSSearchDomains field to given value.

### HasMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasMasterDNSSearchDomains() bool`

HasMasterDNSSearchDomains returns a boolean if a field has been set.

### GetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterNTPServers() []string`

GetMasterNTPServers returns the MasterNTPServers field if non-nil, zero value otherwise.

### GetMasterNTPServersOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterNTPServersOk() (*[]string, bool)`

GetMasterNTPServersOk returns a tuple with the MasterNTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetMasterNTPServers(v []string)`

SetMasterNTPServers sets MasterNTPServers field to given value.

### HasMasterNTPServers

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasMasterNTPServers() bool`

HasMasterNTPServers returns a boolean if a field has been set.

### GetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterStoragePolicy() string`

GetMasterStoragePolicy returns the MasterStoragePolicy field if non-nil, zero value otherwise.

### GetMasterStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetMasterStoragePolicyOk() (*string, bool)`

GetMasterStoragePolicyOk returns a tuple with the MasterStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetMasterStoragePolicy(v string)`

SetMasterStoragePolicy sets MasterStoragePolicy field to given value.

### HasMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasMasterStoragePolicy() bool`

HasMasterStoragePolicy returns a boolean if a field has been set.

### GetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetEphemeralStoragePolicy() string`

GetEphemeralStoragePolicy returns the EphemeralStoragePolicy field if non-nil, zero value otherwise.

### GetEphemeralStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetEphemeralStoragePolicyOk() (*string, bool)`

GetEphemeralStoragePolicyOk returns a tuple with the EphemeralStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetEphemeralStoragePolicy(v string)`

SetEphemeralStoragePolicy sets EphemeralStoragePolicy field to given value.

### HasEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasEphemeralStoragePolicy() bool`

HasEphemeralStoragePolicy returns a boolean if a field has been set.

### GetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetCnsFileConfig() VcenterNamespaceManagementCNSFileConfig`

GetCnsFileConfig returns the CnsFileConfig field if non-nil, zero value otherwise.

### GetCnsFileConfigOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetCnsFileConfigOk() (*VcenterNamespaceManagementCNSFileConfig, bool)`

GetCnsFileConfigOk returns a tuple with the CnsFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetCnsFileConfig(v VcenterNamespaceManagementCNSFileConfig)`

SetCnsFileConfig sets CnsFileConfig field to given value.

### HasCnsFileConfig

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasCnsFileConfig() bool`

HasCnsFileConfig returns a boolean if a field has been set.

### GetLoginBanner

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetLoginBanner() string`

GetLoginBanner returns the LoginBanner field if non-nil, zero value otherwise.

### GetLoginBannerOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetLoginBannerOk() (*string, bool)`

GetLoginBannerOk returns a tuple with the LoginBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBanner

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetLoginBanner(v string)`

SetLoginBanner sets LoginBanner field to given value.

### HasLoginBanner

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasLoginBanner() bool`

HasLoginBanner returns a boolean if a field has been set.

### GetImageStorage

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetImageStorage() VcenterNamespaceManagementClustersImageStorageSpec`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetImageStorageOk() (*VcenterNamespaceManagementClustersImageStorageSpec, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetImageStorage(v VcenterNamespaceManagementClustersImageStorageSpec)`

SetImageStorage sets ImageStorage field to given value.

### HasImageStorage

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasImageStorage() bool`

HasImageStorage returns a boolean if a field has been set.

### GetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetDefaultImageRegistry() VcenterNamespaceManagementClustersImageRegistry`

GetDefaultImageRegistry returns the DefaultImageRegistry field if non-nil, zero value otherwise.

### GetDefaultImageRegistryOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetDefaultImageRegistryOk() (*VcenterNamespaceManagementClustersImageRegistry, bool)`

GetDefaultImageRegistryOk returns a tuple with the DefaultImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetDefaultImageRegistry(v VcenterNamespaceManagementClustersImageRegistry)`

SetDefaultImageRegistry sets DefaultImageRegistry field to given value.

### HasDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasDefaultImageRegistry() bool`

HasDefaultImageRegistry returns a boolean if a field has been set.

### GetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetDefaultImageRepository() string`

GetDefaultImageRepository returns the DefaultImageRepository field if non-nil, zero value otherwise.

### GetDefaultImageRepositoryOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetDefaultImageRepositoryOk() (*string, bool)`

GetDefaultImageRepositoryOk returns a tuple with the DefaultImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetDefaultImageRepository(v string)`

SetDefaultImageRepository sets DefaultImageRepository field to given value.

### HasDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasDefaultImageRepository() bool`

HasDefaultImageRepository returns a boolean if a field has been set.

### GetTlsEndpointCertificate

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetTlsEndpointCertificate() string`

GetTlsEndpointCertificate returns the TlsEndpointCertificate field if non-nil, zero value otherwise.

### GetTlsEndpointCertificateOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetTlsEndpointCertificateOk() (*string, bool)`

GetTlsEndpointCertificateOk returns a tuple with the TlsEndpointCertificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTlsEndpointCertificate

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetTlsEndpointCertificate(v string)`

SetTlsEndpointCertificate sets TlsEndpointCertificate field to given value.

### HasTlsEndpointCertificate

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasTlsEndpointCertificate() bool`

HasTlsEndpointCertificate returns a boolean if a field has been set.

### GetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetDefaultKubernetesServiceContentLibrary() string`

GetDefaultKubernetesServiceContentLibrary returns the DefaultKubernetesServiceContentLibrary field if non-nil, zero value otherwise.

### GetDefaultKubernetesServiceContentLibraryOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetDefaultKubernetesServiceContentLibraryOk() (*string, bool)`

GetDefaultKubernetesServiceContentLibraryOk returns a tuple with the DefaultKubernetesServiceContentLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetDefaultKubernetesServiceContentLibrary(v string)`

SetDefaultKubernetesServiceContentLibrary sets DefaultKubernetesServiceContentLibrary field to given value.

### HasDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasDefaultKubernetesServiceContentLibrary() bool`

HasDefaultKubernetesServiceContentLibrary returns a boolean if a field has been set.

### GetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetWorkloadNtpServers() []string`

GetWorkloadNtpServers returns the WorkloadNtpServers field if non-nil, zero value otherwise.

### GetWorkloadNtpServersOk

`func (o *VcenterNamespaceManagementClustersUpdateSpec) GetWorkloadNtpServersOk() (*[]string, bool)`

GetWorkloadNtpServersOk returns a tuple with the WorkloadNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersUpdateSpec) SetWorkloadNtpServers(v []string)`

SetWorkloadNtpServers sets WorkloadNtpServers field to given value.

### HasWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersUpdateSpec) HasWorkloadNtpServers() bool`

HasWorkloadNtpServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


