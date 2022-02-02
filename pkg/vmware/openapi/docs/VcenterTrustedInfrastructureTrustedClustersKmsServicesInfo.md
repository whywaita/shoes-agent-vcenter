# VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**TrustedCA** | [**VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) |  | 
**Group** | **string** | The group determines the Attestation Service instances this Key Provider Service can accept reports from. | 
**TrustAuthorityCluster** | **string** | The cluster specifies the Trust Authority Cluster this Key Provider Service belongs to. | 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesInfo

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesInfo(address VcenterTrustedInfrastructureNetworkAddress, trustedCA VcenterTrustedInfrastructureX509CertChain, group string, trustAuthorityCluster string, ) *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesInfo instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesInfoWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesInfoWithDefaults() *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesInfoWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain`

GetTrustedCA returns the TrustedCA field if non-nil, zero value otherwise.

### GetTrustedCAOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool)`

GetTrustedCAOk returns a tuple with the TrustedCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain)`

SetTrustedCA sets TrustedCA field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesInfo) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


