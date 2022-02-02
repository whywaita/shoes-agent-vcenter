# VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hosts** | Pointer to **[]string** | A set of host IDs by which to filter the services. If unset, the services will not be filtered by the hosts on which they run. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: HostSystem. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: HostSystem. | [optional] 
**Clusters** | Pointer to **[]string** | A set of cluster IDs by which to filter the services. If unset, the services will not be filtered by the clusters on which they run. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: ClusterComputeResource. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: ClusterComputeResource. | [optional] 
**Address** | Pointer to [**[]VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) | The service&#39;s address. If unset, the services will not be filtered by address. | [optional] 
**Groups** | Pointer to **[]string** | The group determines reports issued by which Attestation Service instances this Key Provider Service can accept. If unset, the services will not be filtered by groupId. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHosts

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetHosts() []string`

GetHosts returns the Hosts field if non-nil, zero value otherwise.

### GetHostsOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetHostsOk() (*[]string, bool)`

GetHostsOk returns a tuple with the Hosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHosts

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetHosts(v []string)`

SetHosts sets Hosts field to given value.

### HasHosts

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasHosts() bool`

HasHosts returns a boolean if a field has been set.

### GetClusters

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetClusters() []string`

GetClusters returns the Clusters field if non-nil, zero value otherwise.

### GetClustersOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetClustersOk() (*[]string, bool)`

GetClustersOk returns a tuple with the Clusters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusters

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetClusters(v []string)`

SetClusters sets Clusters field to given value.

### HasClusters

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasClusters() bool`

HasClusters returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGroups

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsFilterSpec) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


