# VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | The service&#39;s unique identifier. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.attestation.Service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.attestation.Service. | 
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**Group** | **string** | The group specifies the Key Provider Service instances can accept reports issued by this Attestation Service instance. | 
**TrustAuthorityCluster** | **string** | The cluster specifies the Trust Authority Cluster this Attestation Service belongs to. | 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary

`func NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary(service string, address VcenterTrustedInfrastructureNetworkAddress, group string, trustAuthorityCluster string, ) *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary`

NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary instantiates a new VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesSummaryWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesSummaryWithDefaults() *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary`

NewVcenterTrustedInfrastructureTrustedClustersAttestationServicesSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) SetService(v string)`

SetService sets Service field to given value.


### GetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersAttestationServicesSummary) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


