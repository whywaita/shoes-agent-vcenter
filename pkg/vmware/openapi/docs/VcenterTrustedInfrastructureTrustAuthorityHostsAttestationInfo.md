# VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Host** | **string** | The trusted ESX on which the service runs. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | 
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**Group** | **string** | The group ID determines which Attestation Service instances this Attestation Service can communicate with. | 
**Cluster** | **string** | The opaque string identifier of the cluster in which the Attestation Service is part of. | 
**TrustedCA** | [**VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) |  | 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo(host string, address VcenterTrustedInfrastructureNetworkAddress, group string, cluster string, trustedCA VcenterTrustedInfrastructureX509CertChain, ) *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo`

NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfoWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfoWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo`

NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfoWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetHost(v string)`

SetHost sets Host field to given value.


### GetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetCluster(v string)`

SetCluster sets Cluster field to given value.


### GetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain`

GetTrustedCA returns the TrustedCA field if non-nil, zero value otherwise.

### GetTrustedCAOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool)`

GetTrustedCAOk returns a tuple with the TrustedCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationInfo) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain)`

SetTrustedCA sets TrustedCA field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


