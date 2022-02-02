# VcenterTrustedInfrastructureAttestationServicesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**TrustedCA** | [**VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) |  | 
**Group** | **string** | The group specifies the Key Provider Service instances that can accept reports issued by this Attestation Service instance. | 
**TrustAuthorityCluster** | **string** | The cluster specifies the Trust Authority Cluster this Attestation Service belongs to. | 

## Methods

### NewVcenterTrustedInfrastructureAttestationServicesCreateSpec

`func NewVcenterTrustedInfrastructureAttestationServicesCreateSpec(address VcenterTrustedInfrastructureNetworkAddress, trustedCA VcenterTrustedInfrastructureX509CertChain, group string, trustAuthorityCluster string, ) *VcenterTrustedInfrastructureAttestationServicesCreateSpec`

NewVcenterTrustedInfrastructureAttestationServicesCreateSpec instantiates a new VcenterTrustedInfrastructureAttestationServicesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureAttestationServicesCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureAttestationServicesCreateSpecWithDefaults() *VcenterTrustedInfrastructureAttestationServicesCreateSpec`

NewVcenterTrustedInfrastructureAttestationServicesCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureAttestationServicesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetTrustedCA

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain`

GetTrustedCA returns the TrustedCA field if non-nil, zero value otherwise.

### GetTrustedCAOk

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool)`

GetTrustedCAOk returns a tuple with the TrustedCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCA

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain)`

SetTrustedCA sets TrustedCA field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesCreateSpec) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


