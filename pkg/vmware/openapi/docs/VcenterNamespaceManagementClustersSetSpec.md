# VcenterNamespaceManagementClustersSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SizeHint** | [**VcenterNamespaceManagementSizingHint**](VcenterNamespaceManagementSizingHint.md) |  | 
**NetworkProvider** | [**VcenterNamespaceManagementClustersNetworkProvider**](VcenterNamespaceManagementClustersNetworkProvider.md) |  | 
**NcpClusterNetworkSpec** | Pointer to [**VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec**](VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec.md) |  | [optional] 
**MasterDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on Kubernetes API server, specified in order of preference. If unset, DNS servers set on Kubernetes API server will be cleared. | [optional] 
**WorkerDNS** | Pointer to **[]string** | List of DNS server IP addresses to use on the worker nodes, specified in order of preference. If unset, DNS servers set on worker nodes will be cleared. | [optional] 
**MasterDNSSearchDomains** | Pointer to **[]string** | List of domains (for example \&quot;vmware.com\&quot;) to be searched when trying to lookup a host name on Kubernetes API server, specified in order of preference. If unset, DNS search domains set on Kubernetes API server will be cleared. | [optional] 
**MasterNTPServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use on Kubernetes API server, specified in order of preference. If unset, VMware Tools based time synchronization is enabled and any set NTP servers are cleared. | [optional] 
**MasterStoragePolicy** | **string** | Identifier of storage policy associated with Kubernetes API server. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | 
**EphemeralStoragePolicy** | **string** | Identifier of storage policy associated with ephemeral disks of all the Kubernetes Pods in the cluster. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: SpsStorageProfile. When operations return a value of this structure as a result, the field will be an identifier for the resource type: SpsStorageProfile. | 
**CnsFileConfig** | Pointer to [**VcenterNamespaceManagementCNSFileConfig**](VcenterNamespaceManagementCNSFileConfig.md) |  | [optional] 
**LoginBanner** | Pointer to **string** | Disclaimer to be displayed prior to login via the Kubectl plugin. If unset, disclaimer to be displayed prior to login via the Kubectl plugin will be cleared. | [optional] 
**ImageStorage** | [**VcenterNamespaceManagementClustersImageStorageSpec**](VcenterNamespaceManagementClustersImageStorageSpec.md) |  | 
**DefaultImageRegistry** | Pointer to [**VcenterNamespaceManagementClustersImageRegistry**](VcenterNamespaceManagementClustersImageRegistry.md) |  | [optional] 
**DefaultImageRepository** | Pointer to **string** | Default image repository to use when Kubernetes Pod container specification does not specify it as part of the container image name. If unset, default image repository will be set to Docker Hub official repository in case of Docker Hub image registry, otherwise will be set to empty string. | [optional] 
**DefaultKubernetesServiceContentLibrary** | Pointer to **string** | Identifier of the Content Library which holds the VM Images for vSphere Kubernetes Service. This Content Library should be subscribed to VMware&#39;s hosted vSphere Kubernetes Service Repository. Modifying or clearing the Content Library identifier will not affect existing vSphere Kubernetes Service clusters. However, upgrades or scale-out of existing clusters may be affected if the new Content Library doesn&#39;t have the necessary VM Images. If unset, the Content Library identifier will be cleared. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: content.Library. When operations return a value of this structure as a result, the field will be an identifier for the resource type: content.Library. | [optional] 
**WorkloadNtpServers** | Pointer to **[]string** | List of NTP server DNS names or IP addresses to use for workloads such as Tanzu Kubernetes Grid VMs, specified in order of preference. If unset, NTP for Kubernetes API servers will be used. | [optional] 

## Methods

### NewVcenterNamespaceManagementClustersSetSpec

`func NewVcenterNamespaceManagementClustersSetSpec(sizeHint VcenterNamespaceManagementSizingHint, networkProvider VcenterNamespaceManagementClustersNetworkProvider, masterStoragePolicy string, ephemeralStoragePolicy string, imageStorage VcenterNamespaceManagementClustersImageStorageSpec, ) *VcenterNamespaceManagementClustersSetSpec`

NewVcenterNamespaceManagementClustersSetSpec instantiates a new VcenterNamespaceManagementClustersSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterNamespaceManagementClustersSetSpecWithDefaults

`func NewVcenterNamespaceManagementClustersSetSpecWithDefaults() *VcenterNamespaceManagementClustersSetSpec`

NewVcenterNamespaceManagementClustersSetSpecWithDefaults instantiates a new VcenterNamespaceManagementClustersSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSizeHint

`func (o *VcenterNamespaceManagementClustersSetSpec) GetSizeHint() VcenterNamespaceManagementSizingHint`

GetSizeHint returns the SizeHint field if non-nil, zero value otherwise.

### GetSizeHintOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetSizeHintOk() (*VcenterNamespaceManagementSizingHint, bool)`

GetSizeHintOk returns a tuple with the SizeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeHint

`func (o *VcenterNamespaceManagementClustersSetSpec) SetSizeHint(v VcenterNamespaceManagementSizingHint)`

SetSizeHint sets SizeHint field to given value.


### GetNetworkProvider

`func (o *VcenterNamespaceManagementClustersSetSpec) GetNetworkProvider() VcenterNamespaceManagementClustersNetworkProvider`

GetNetworkProvider returns the NetworkProvider field if non-nil, zero value otherwise.

### GetNetworkProviderOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetNetworkProviderOk() (*VcenterNamespaceManagementClustersNetworkProvider, bool)`

GetNetworkProviderOk returns a tuple with the NetworkProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkProvider

`func (o *VcenterNamespaceManagementClustersSetSpec) SetNetworkProvider(v VcenterNamespaceManagementClustersNetworkProvider)`

SetNetworkProvider sets NetworkProvider field to given value.


### GetNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersSetSpec) GetNcpClusterNetworkSpec() VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec`

GetNcpClusterNetworkSpec returns the NcpClusterNetworkSpec field if non-nil, zero value otherwise.

### GetNcpClusterNetworkSpecOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetNcpClusterNetworkSpecOk() (*VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec, bool)`

GetNcpClusterNetworkSpecOk returns a tuple with the NcpClusterNetworkSpec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersSetSpec) SetNcpClusterNetworkSpec(v VcenterNamespaceManagementClustersNCPClusterNetworkSetSpec)`

SetNcpClusterNetworkSpec sets NcpClusterNetworkSpec field to given value.

### HasNcpClusterNetworkSpec

`func (o *VcenterNamespaceManagementClustersSetSpec) HasNcpClusterNetworkSpec() bool`

HasNcpClusterNetworkSpec returns a boolean if a field has been set.

### GetMasterDNS

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterDNS() []string`

GetMasterDNS returns the MasterDNS field if non-nil, zero value otherwise.

### GetMasterDNSOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterDNSOk() (*[]string, bool)`

GetMasterDNSOk returns a tuple with the MasterDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNS

`func (o *VcenterNamespaceManagementClustersSetSpec) SetMasterDNS(v []string)`

SetMasterDNS sets MasterDNS field to given value.

### HasMasterDNS

`func (o *VcenterNamespaceManagementClustersSetSpec) HasMasterDNS() bool`

HasMasterDNS returns a boolean if a field has been set.

### GetWorkerDNS

`func (o *VcenterNamespaceManagementClustersSetSpec) GetWorkerDNS() []string`

GetWorkerDNS returns the WorkerDNS field if non-nil, zero value otherwise.

### GetWorkerDNSOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetWorkerDNSOk() (*[]string, bool)`

GetWorkerDNSOk returns a tuple with the WorkerDNS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkerDNS

`func (o *VcenterNamespaceManagementClustersSetSpec) SetWorkerDNS(v []string)`

SetWorkerDNS sets WorkerDNS field to given value.

### HasWorkerDNS

`func (o *VcenterNamespaceManagementClustersSetSpec) HasWorkerDNS() bool`

HasWorkerDNS returns a boolean if a field has been set.

### GetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterDNSSearchDomains() []string`

GetMasterDNSSearchDomains returns the MasterDNSSearchDomains field if non-nil, zero value otherwise.

### GetMasterDNSSearchDomainsOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterDNSSearchDomainsOk() (*[]string, bool)`

GetMasterDNSSearchDomainsOk returns a tuple with the MasterDNSSearchDomains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersSetSpec) SetMasterDNSSearchDomains(v []string)`

SetMasterDNSSearchDomains sets MasterDNSSearchDomains field to given value.

### HasMasterDNSSearchDomains

`func (o *VcenterNamespaceManagementClustersSetSpec) HasMasterDNSSearchDomains() bool`

HasMasterDNSSearchDomains returns a boolean if a field has been set.

### GetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterNTPServers() []string`

GetMasterNTPServers returns the MasterNTPServers field if non-nil, zero value otherwise.

### GetMasterNTPServersOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterNTPServersOk() (*[]string, bool)`

GetMasterNTPServersOk returns a tuple with the MasterNTPServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterNTPServers

`func (o *VcenterNamespaceManagementClustersSetSpec) SetMasterNTPServers(v []string)`

SetMasterNTPServers sets MasterNTPServers field to given value.

### HasMasterNTPServers

`func (o *VcenterNamespaceManagementClustersSetSpec) HasMasterNTPServers() bool`

HasMasterNTPServers returns a boolean if a field has been set.

### GetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterStoragePolicy() string`

GetMasterStoragePolicy returns the MasterStoragePolicy field if non-nil, zero value otherwise.

### GetMasterStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetMasterStoragePolicyOk() (*string, bool)`

GetMasterStoragePolicyOk returns a tuple with the MasterStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterStoragePolicy

`func (o *VcenterNamespaceManagementClustersSetSpec) SetMasterStoragePolicy(v string)`

SetMasterStoragePolicy sets MasterStoragePolicy field to given value.


### GetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersSetSpec) GetEphemeralStoragePolicy() string`

GetEphemeralStoragePolicy returns the EphemeralStoragePolicy field if non-nil, zero value otherwise.

### GetEphemeralStoragePolicyOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetEphemeralStoragePolicyOk() (*string, bool)`

GetEphemeralStoragePolicyOk returns a tuple with the EphemeralStoragePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeralStoragePolicy

`func (o *VcenterNamespaceManagementClustersSetSpec) SetEphemeralStoragePolicy(v string)`

SetEphemeralStoragePolicy sets EphemeralStoragePolicy field to given value.


### GetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersSetSpec) GetCnsFileConfig() VcenterNamespaceManagementCNSFileConfig`

GetCnsFileConfig returns the CnsFileConfig field if non-nil, zero value otherwise.

### GetCnsFileConfigOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetCnsFileConfigOk() (*VcenterNamespaceManagementCNSFileConfig, bool)`

GetCnsFileConfigOk returns a tuple with the CnsFileConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnsFileConfig

`func (o *VcenterNamespaceManagementClustersSetSpec) SetCnsFileConfig(v VcenterNamespaceManagementCNSFileConfig)`

SetCnsFileConfig sets CnsFileConfig field to given value.

### HasCnsFileConfig

`func (o *VcenterNamespaceManagementClustersSetSpec) HasCnsFileConfig() bool`

HasCnsFileConfig returns a boolean if a field has been set.

### GetLoginBanner

`func (o *VcenterNamespaceManagementClustersSetSpec) GetLoginBanner() string`

GetLoginBanner returns the LoginBanner field if non-nil, zero value otherwise.

### GetLoginBannerOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetLoginBannerOk() (*string, bool)`

GetLoginBannerOk returns a tuple with the LoginBanner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginBanner

`func (o *VcenterNamespaceManagementClustersSetSpec) SetLoginBanner(v string)`

SetLoginBanner sets LoginBanner field to given value.

### HasLoginBanner

`func (o *VcenterNamespaceManagementClustersSetSpec) HasLoginBanner() bool`

HasLoginBanner returns a boolean if a field has been set.

### GetImageStorage

`func (o *VcenterNamespaceManagementClustersSetSpec) GetImageStorage() VcenterNamespaceManagementClustersImageStorageSpec`

GetImageStorage returns the ImageStorage field if non-nil, zero value otherwise.

### GetImageStorageOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetImageStorageOk() (*VcenterNamespaceManagementClustersImageStorageSpec, bool)`

GetImageStorageOk returns a tuple with the ImageStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageStorage

`func (o *VcenterNamespaceManagementClustersSetSpec) SetImageStorage(v VcenterNamespaceManagementClustersImageStorageSpec)`

SetImageStorage sets ImageStorage field to given value.


### GetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersSetSpec) GetDefaultImageRegistry() VcenterNamespaceManagementClustersImageRegistry`

GetDefaultImageRegistry returns the DefaultImageRegistry field if non-nil, zero value otherwise.

### GetDefaultImageRegistryOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetDefaultImageRegistryOk() (*VcenterNamespaceManagementClustersImageRegistry, bool)`

GetDefaultImageRegistryOk returns a tuple with the DefaultImageRegistry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersSetSpec) SetDefaultImageRegistry(v VcenterNamespaceManagementClustersImageRegistry)`

SetDefaultImageRegistry sets DefaultImageRegistry field to given value.

### HasDefaultImageRegistry

`func (o *VcenterNamespaceManagementClustersSetSpec) HasDefaultImageRegistry() bool`

HasDefaultImageRegistry returns a boolean if a field has been set.

### GetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersSetSpec) GetDefaultImageRepository() string`

GetDefaultImageRepository returns the DefaultImageRepository field if non-nil, zero value otherwise.

### GetDefaultImageRepositoryOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetDefaultImageRepositoryOk() (*string, bool)`

GetDefaultImageRepositoryOk returns a tuple with the DefaultImageRepository field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersSetSpec) SetDefaultImageRepository(v string)`

SetDefaultImageRepository sets DefaultImageRepository field to given value.

### HasDefaultImageRepository

`func (o *VcenterNamespaceManagementClustersSetSpec) HasDefaultImageRepository() bool`

HasDefaultImageRepository returns a boolean if a field has been set.

### GetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersSetSpec) GetDefaultKubernetesServiceContentLibrary() string`

GetDefaultKubernetesServiceContentLibrary returns the DefaultKubernetesServiceContentLibrary field if non-nil, zero value otherwise.

### GetDefaultKubernetesServiceContentLibraryOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetDefaultKubernetesServiceContentLibraryOk() (*string, bool)`

GetDefaultKubernetesServiceContentLibraryOk returns a tuple with the DefaultKubernetesServiceContentLibrary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersSetSpec) SetDefaultKubernetesServiceContentLibrary(v string)`

SetDefaultKubernetesServiceContentLibrary sets DefaultKubernetesServiceContentLibrary field to given value.

### HasDefaultKubernetesServiceContentLibrary

`func (o *VcenterNamespaceManagementClustersSetSpec) HasDefaultKubernetesServiceContentLibrary() bool`

HasDefaultKubernetesServiceContentLibrary returns a boolean if a field has been set.

### GetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersSetSpec) GetWorkloadNtpServers() []string`

GetWorkloadNtpServers returns the WorkloadNtpServers field if non-nil, zero value otherwise.

### GetWorkloadNtpServersOk

`func (o *VcenterNamespaceManagementClustersSetSpec) GetWorkloadNtpServersOk() (*[]string, bool)`

GetWorkloadNtpServersOk returns a tuple with the WorkloadNtpServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersSetSpec) SetWorkloadNtpServers(v []string)`

SetWorkloadNtpServers sets WorkloadNtpServers field to given value.

### HasWorkloadNtpServers

`func (o *VcenterNamespaceManagementClustersSetSpec) HasWorkloadNtpServers() bool`

HasWorkloadNtpServers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


