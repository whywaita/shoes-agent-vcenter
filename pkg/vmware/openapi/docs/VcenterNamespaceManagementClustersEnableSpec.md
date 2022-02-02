# VcenterNamespaceManagementClustersEnableSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeHint** | [**VcenterNamespaceManagementSizingHint**](VcenterNamespaceManagementSizingHint.md) |  | 
**ServiceCidr** | [**VcenterNamespaceManagementIpv4Cidr**](VcenterNamespaceManagementIpv4Cidr.md) |  | 
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**NcpClusterNetworkSpec** | Pointer to [**VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec**](VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec.md) |  | [optional] 
**WorkloadNetworksSpec** | Pointer to [**VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec**](VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec.md) |  | [optional] 
**WorkloadNtpServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use for workloads such as Tanzu Kubernetes Grid VMs, specified in order of preference. If unset, NTP server for Kubernetes API servers will be used. | [optional] 
**LoadBalancerConfigSpec** | Pointer to [**VcenterNamespaceManagementLoadBalancersConfigSpec**](VcenterNamespaceManagementLoadBalancersConfigSpec.md) |  | [optional] 
**MasterManagementNetwork** | [**VcenterNamespaceManagementClustersNetworkSpec**](VcenterNamespaceManagementClustersNetworkSpec.md) |  | 
**MasterDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. If unset, no default DNS servers are set. | [optional] 
**WorkerDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on the worker nodes, specified in order of preference. If unset, no default DNS servers are set. | [optional] 
**MasterDNSSearchDomains** | Pointer to **[]string** | List of domains (for example \&quot;vmware.com\&quot;) to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. If unset, no default DNS search domains are set. | [optional] 
**MasterNTPServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. If unset, VMware Tools based time synchronization is enabled. | [optional] 
**MasterStoragePolicy** | **string** | Identifier of storage policy associated with Kubernetes API server. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | 
**EphemeralStoragePolicy** | **string** | Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | 
**CnsFileConfig** | Pointer to [**VcenterNamespaceManagementCNSFileConfig**](VcenterNamespaceManagementCNSFileConfig.md) |  | [optional] 
**LoginBanner** | Pointer to **string** | Disclaimer to be displayed prior to login via the Kubectl plugin. If unset, just skip it. | [optional] 
**MasterDNSNames** | Pointer to **[]string** | List of additional DNS names to associate with the Kubernetes API server. These DNS names are embedded in the TLS certificate presented by the API server. If unset, no additional DNS names are embedded in the TLS certificate. | [optional] 
**ImageStorage** | [**VcenterNamespaceManagementClustersImageStorageSpec**](VcenterNamespaceManagementClustersImageStorageSpec.md) |  | 
**DefaultImageRegistry** | Pointer to [**VcenterNamespaceManagementClustersImageRegistry**](VcenterNamespaceManagementClustersImageRegistry.md) |  | [optional] 
**DefaultImageRepository** | Pointer to **string** | Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. If unset, defaults to Docker Hub official repository in case of Docker Hub image registry, otherwise defaults to empty string. | [optional] 
**DefaultKubernetesServiceContentLibrary** | Pointer to **string** | Identifier of the Content Library which holds the VM Images for vSphere Kubernetes Service. This Content Library should be subscribed to VMware&#39;s hosted vSphere Kubernetes Service Repository. If unset, the Content Library identifier will not be set. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: content.Library. When operations return a value of this structure as a result, the field will be an identifier for the resource type: content.Library. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersEnableSpec

`func NewVcenterNamespaceManagementClustersEnableSpec(sizeHint VcenterNamespaceManagementSizingHint, serviceCidr VcenterNamespaceManagementIpv4Cidr, networkProvider VcenterNamespaceManagementClustersNetworkProvider, masterManagementNetwork VcenterNamespaceManagementClustersNetworkSpec, masterStoragePolicy string, ephemeralStoragePolicy string, imageStorage VcenterNamespaceManagementClustersImageStorageSpec, ) *VcenterNamespaceManagementClustersEnableSpec`

NewVcenterNamespaceManagementClustersEnableSpec instantiates a new VcenterNamespaceManagementClustersEnableSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersEnableSpecWithDefaults

`func NewVcenterNamespaceManagementClustersEnableSpecWithDefaults() *VcenterNamespaceManagementClustersEnableSpec`

NewVcenterNamespaceManagementClustersEnableSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersEnableSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeHint

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetSizeHint() VcenterNamespaceManagementSizingHint`

GetSizeHint returns the SizeHint field if non-nil, zero value otherwise.

### GetSizeHintOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetSizeHintOk() (*VcenterNamespaceManagementSizingHint, bool)`

GetSizeHintOk returns a tuple with the SizeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHint

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetSizeHint(v VcenterNamespaceManagementSizingHint)`

SetSizeHint sets SizeHint field to given value.


### GetServiceCidr

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetServiceCidr() VcenterNamespaceManagementIpv4Cidr`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetServiceCidrOk() (*VcenterNamespaceManagementIpv4Cidr, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetServiceCidr(v VcenterNamespaceManagementIpv4Cidr)`

SetServiceCidr sets ServiceCidr field to given value.


### GetNetworkProvider

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetNcpClusterNetworkSpec() VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec`

GetNcpClusterNetworkSpec returns the NcpClusterNetworkSpec field if non-nil, zero value otherwise.

### GetNcpClusterNetworkSpecOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetNcpClusterNetworkSpecOk() (*VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec, bool)`

GetNcpClusterNetworkSpecOk returns a tuple with the NcpClusterNetworkSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetNcpClusterNetworkSpec(v VcenterNamespaceManagementClustersNCPClusterNetworkEnableSpec)`

SetNcpClusterNetworkSpec sets NcpClusterNetworkSpec field to given value.

### HasNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasNcpClusterNetworkSpec() bool`

HasNcpClusterNetworkSpec returns a boolean if a field has been set.

### GetWorkloadNetworksSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetWorkloadNetworksSpec() VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec`

GetWorkloadNetworksSpec returns the WorkloadNetworksSpec field if non-nil, zero value otherwise.

### GetWorkloadNetworksSpecOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetWorkloadNetworksSpecOk() (*VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec, bool)`

GetWorkloadNetworksSpecOk returns a tuple with the WorkloadNetworksSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNetworksSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetWorkloadNetworksSpec(v VcenterNamespaceManagementClustersWorkloadNetworksEnableSpec)`

SetWorkloadNetworksSpec sets WorkloadNetworksSpec field to given value.

### HasWorkloadNetworksSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasWorkloadNetworksSpec() bool`

HasWorkloadNetworksSpec returns a boolean if a field has been set.

### GetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetWorkloadNtpServers() []string`

GetWorkloadNtpServers returns the WorkloadNtpServers field if non-nil, zero value otherwise.

### GetWorkloadNtpServersOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetWorkloadNtpServersOk() (*[]string, bool)`

GetWorkloadNtpServersOk returns a tuple with the WorkloadNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetWorkloadNtpServers(v []string)`

SetWorkloadNtpServers sets WorkloadNtpServers field to given value.

### HasWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasWorkloadNtpServers() bool`

HasWorkloadNtpServers returns a boolean if a field has been set.

### GetLoadBalancerConfigSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetLoadBalancerConfigSpec() VcenterNamespaceManagementLoadBalancersConfigSpec`

GetLoadBalancerConfigSpec returns the LoadBalancerConfigSpec field if non-nil, zero value otherwise.

### GetLoadBalancerConfigSpecOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetLoadBalancerConfigSpecOk() (*VcenterNamespaceManagementLoadBalancersConfigSpec, bool)`

GetLoadBalancerConfigSpecOk returns a tuple with the LoadBalancerConfigSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalancerConfigSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetLoadBalancerConfigSpec(v VcenterNamespaceManagementLoadBalancersConfigSpec)`

SetLoadBalancerConfigSpec sets LoadBalancerConfigSpec field to given value.

### HasLoadBalancerConfigSpec

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasLoadBalancerConfigSpec() bool`

HasLoadBalancerConfigSpec returns a boolean if a field has been set.

### GetMasterManagementNetwork

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterManagementNetwork() VcenterNamespaceManagementClustersNetworkSpec`

GetMasterManagementNetwork returns the MasterManagementNetwork field if non-nil, zero value otherwise.

### GetMasterManagementNetworkOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterManagementNetworkOk() (*VcenterNamespaceManagementClustersNetworkSpec, bool)`

GetMasterManagementNetworkOk returns a tuple with the MasterManagementNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterManagementNetwork

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetMasterManagementNetwork(v VcenterNamespaceManagementClustersNetworkSpec)`

SetMasterManagementNetwork sets MasterManagementNetwork field to given value.


### GetMasterDNS

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterDNS() []string`

GetMasterDNS returns the MasterDNS field if non-nil, zero value otherwise.

### GetMasterDNSOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterDNSOk() (*[]string, bool)`

GetMasterDNSOk returns a tuple with the MasterDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNS

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetMasterDNS(v []string)`

SetMasterDNS sets MasterDNS field to given value.

### HasMasterDNS

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasMasterDNS() bool`

HasMasterDNS returns a boolean if a field has been set.

### GetWorkerDNS

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetWorkerDNS() []string`

GetWorkerDNS returns the WorkerDNS field if non-nil, zero value otherwise.

### GetWorkerDNSOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetWorkerDNSOk() (*[]string, bool)`

GetWorkerDNSOk returns a tuple with the WorkerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerDNS

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetWorkerDNS(v []string)`

SetWorkerDNS sets WorkerDNS field to given value.

### HasWorkerDNS

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasWorkerDNS() bool`

HasWorkerDNS returns a boolean if a field has been set.

### GetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterDNSSearchDomains() []string`

GetMasterDNSSearchDomains returns the MasterDNSSearchDomains field if non-nil, zero value otherwise.

### GetMasterDNSSearchDomainsOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterDNSSearchDomainsOk() (*[]string, bool)`

GetMasterDNSSearchDomainsOk returns a tuple with the MasterDNSSearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetMasterDNSSearchDomains(v []string)`

SetMasterDNSSearchDomains sets MasterDNSSearchDomains field to given value.

### HasMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasMasterDNSSearchDomains() bool`

HasMasterDNSSearchDomains returns a boolean if a field has been set.

### GetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterNTPServers() []string`

GetMasterNTPServers returns the MasterNTPServers field if non-nil, zero value otherwise.

### GetMasterNTPServersOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterNTPServersOk() (*[]string, bool)`

GetMasterNTPServersOk returns a tuple with the MasterNTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetMasterNTPServers(v []string)`

SetMasterNTPServers sets MasterNTPServers field to given value.

### HasMasterNTPServers

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasMasterNTPServers() bool`

HasMasterNTPServers returns a boolean if a field has been set.

### GetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterStoragePolicy() string`

GetMasterStoragePolicy returns the MasterStoragePolicy field if non-nil, zero value otherwise.

### GetMasterStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterStoragePolicyOk() (*string, bool)`

GetMasterStoragePolicyOk returns a tuple with the MasterStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetMasterStoragePolicy(v string)`

SetMasterStoragePolicy sets MasterStoragePolicy field to given value.


### GetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetEphemeralStoragePolicy() string`

GetEphemeralStoragePolicy returns the EphemeralStoragePolicy field if non-nil, zero value otherwise.

### GetEphemeralStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetEphemeralStoragePolicyOk() (*string, bool)`

GetEphemeralStoragePolicyOk returns a tuple with the EphemeralStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetEphemeralStoragePolicy(v string)`

SetEphemeralStoragePolicy sets EphemeralStoragePolicy field to given value.


### GetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetCnsFileConfig() VcenterNamespaceManagementCNSFileConfig`

GetCnsFileConfig returns the CnsFileConfig field if non-nil, zero value otherwise.

### GetCnsFileConfigOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetCnsFileConfigOk() (*VcenterNamespaceManagementCNSFileConfig, bool)`

GetCnsFileConfigOk returns a tuple with the CnsFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetCnsFileConfig(v VcenterNamespaceManagementCNSFileConfig)`

SetCnsFileConfig sets CnsFileConfig field to given value.

### HasCnsFileConfig

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasCnsFileConfig() bool`

HasCnsFileConfig returns a boolean if a field has been set.

### GetLoginBanner

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetLoginBanner() string`

GetLoginBanner returns the LoginBanner field if non-nil, zero value otherwise.

### GetLoginBannerOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetLoginBannerOk() (*string, bool)`

GetLoginBannerOk returns a tuple with the LoginBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBanner

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetLoginBanner(v string)`

SetLoginBanner sets LoginBanner field to given value.

### HasLoginBanner

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasLoginBanner() bool`

HasLoginBanner returns a boolean if a field has been set.

### GetMasterDNSNames

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterDNSNames() []string`

GetMasterDNSNames returns the MasterDNSNames field if non-nil, zero value otherwise.

### GetMasterDNSNamesOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetMasterDNSNamesOk() (*[]string, bool)`

GetMasterDNSNamesOk returns a tuple with the MasterDNSNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNSNames

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetMasterDNSNames(v []string)`

SetMasterDNSNames sets MasterDNSNames field to given value.

### HasMasterDNSNames

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasMasterDNSNames() bool`

HasMasterDNSNames returns a boolean if a field has been set.

### GetImageStorage

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetImageStorage() VcenterNamespaceManagementClustersImageStorageSpec`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetImageStorageOk() (*VcenterNamespaceManagementClustersImageStorageSpec, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetImageStorage(v VcenterNamespaceManagementClustersImageStorageSpec)`

SetImageStorage sets ImageStorage field to given value.


### GetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetDefaultImageRegistry() VcenterNamespaceManagementClustersImageRegistry`

GetDefaultImageRegistry returns the DefaultImageRegistry field if non-nil, zero value otherwise.

### GetDefaultImageRegistryOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetDefaultImageRegistryOk() (*VcenterNamespaceManagementClustersImageRegistry, bool)`

GetDefaultImageRegistryOk returns a tuple with the DefaultImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetDefaultImageRegistry(v VcenterNamespaceManagementClustersImageRegistry)`

SetDefaultImageRegistry sets DefaultImageRegistry field to given value.

### HasDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasDefaultImageRegistry() bool`

HasDefaultImageRegistry returns a boolean if a field has been set.

### GetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetDefaultImageRepository() string`

GetDefaultImageRepository returns the DefaultImageRepository field if non-nil, zero value otherwise.

### GetDefaultImageRepositoryOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetDefaultImageRepositoryOk() (*string, bool)`

GetDefaultImageRepositoryOk returns a tuple with the DefaultImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetDefaultImageRepository(v string)`

SetDefaultImageRepository sets DefaultImageRepository field to given value.

### HasDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasDefaultImageRepository() bool`

HasDefaultImageRepository returns a boolean if a field has been set.

### GetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetDefaultKubernetesServiceContentLibrary() string`

GetDefaultKubernetesServiceContentLibrary returns the DefaultKubernetesServiceContentLibrary field if non-nil, zero value otherwise.

### GetDefaultKubernetesServiceContentLibraryOk

`func (o *VcenterNamespaceManagementClustersEnableSpec) GetDefaultKubernetesServiceContentLibraryOk() (*string, bool)`

GetDefaultKubernetesServiceContentLibraryOk returns a tuple with the DefaultKubernetesServiceContentLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersEnableSpec) SetDefaultKubernetesServiceContentLibrary(v string)`

SetDefaultKubernetesServiceContentLibrary sets DefaultKubernetesServiceContentLibrary field to given value.

### HasDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersEnableSpec) HasDefaultKubernetesServiceContentLibrary() bool`

HasDefaultKubernetesServiceContentLibrary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


