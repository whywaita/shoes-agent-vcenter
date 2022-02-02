# VcenterTrustedInfrastructureAttestationServicesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Services** | Pointer to **[]string** | A set of IDs by which to filter the services. If unset, the services will not be filtered by ID. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.trusted_infrastructure.attestation.Service. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.trusted_infrastructure.attestation.Service. | [optional] 
**Address** | Pointer to [**[]VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) | A set of address by which to filter. If unset, the services will not be filtered by address. | [optional] 
**Group** | Pointer to **[]string** | The group specifies the Key Provider Service instances that can accept reports issued by this Attestation Service instance. If unset, the services will not be filtered by group. | [optional] 
**TrustAuthorityCluster** | Pointer to **[]string** | The cluster specifies the Trust Authority Cluster this Attestation Service belongs to. If unset, the services will not be filtered by trustAuthorityCluster. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureAttestationServicesFilterSpec

`func NewVcenterTrustedInfrastructureAttestationServicesFilterSpec() *VcenterTrustedInfrastructureAttestationServicesFilterSpec`

NewVcenterTrustedInfrastructureAttestationServicesFilterSpec instantiates a new VcenterTrustedInfrastructureAttestationServicesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureAttestationServicesFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureAttestationServicesFilterSpecWithDefaults() *VcenterTrustedInfrastructureAttestationServicesFilterSpec`

NewVcenterTrustedInfrastructureAttestationServicesFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureAttestationServicesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServices

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetServices() []string`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetServicesOk() (*[]string, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetServices(v []string)`

SetServices sets Services field to given value.

### HasServices

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetAddress() []VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetAddressOk() (*[]VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetAddress(v []VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetTrustAuthorityCluster() []string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) GetTrustAuthorityClusterOk() (*[]string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) SetTrustAuthorityCluster(v []string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.

### HasTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesFilterSpec) HasTrustAuthorityCluster() bool`

HasTrustAuthorityCluster returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


