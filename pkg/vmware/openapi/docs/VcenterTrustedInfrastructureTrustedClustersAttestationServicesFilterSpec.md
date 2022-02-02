# VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | A set of IDs by which to filter the services. If unset, the services will not be filtered by ID. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.trusted_infrastructure.attestation.Service. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.trusted_infrastructure.attestation.Service. | [optional] 
**Address** | Pointer to [**[]VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) | The service&#39;s address. If unset, the services will not be filtered by address. | [optional] 
**Group** | Pointer to **[]string** | The group specifies the Key Provider Service instances can accept reports issued by this Attestation Service instance. If unset, the services will not be filtered by group. | [optional] 
**TrustAuthorityCluster** | Pointer to **[]string** | The cluster specifies the Trust Authority Cluster this Attestation Service instance belongs to. If unset, the services will not be filtered by trustAuthorityCluster. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec

`func NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec() *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec`

NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec instantiates a new VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec`

NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetTrustAuthorityCluster() []string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) GetTrustAuthorityClusterOk() (*[]string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) SetTrustAuthorityCluster(v []string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.

### HasTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesFilterSpec) HasTrustAuthorityCluster() bool`

HasTrustAuthorityCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


