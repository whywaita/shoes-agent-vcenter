# VcenterTrustedInfrastructureAttestationServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | The service&#39;s unique identifier. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.attestation.Service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.attestation.Service. | 
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**Group** | **string** | The group specifies the Key Provider Service instances that can accept reports issued by this Attestation Service instance. | 
**TrustAuthorityCluster** | **string** | The cluster specifies the Trust Authority Cluster this Attestation Service instance belongs to. | 

## Methods

### NewVcenterTrustedInfrastructureAttestationServicesSummary

`func NewVcenterTrustedInfrastructureAttestationServicesSummary(service string, address VcenterTrustedInfrastructureNetworkAddress, group string, trustAuthorityCluster string, ) *VcenterTrustedInfrastructureAttestationServicesSummary`

NewVcenterTrustedInfrastructureAttestationServicesSummary instantiates a new VcenterTrustedInfrastructureAttestationServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureAttestationServicesSummaryWithDefaults

`func NewVcenterTrustedInfrastructureAttestationServicesSummaryWithDefaults() *VcenterTrustedInfrastructureAttestationServicesSummary`

NewVcenterTrustedInfrastructureAttestationServicesSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureAttestationServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetService(v string)`

SetService sets Service field to given value.


### GetAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureAttestationServicesSummary) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


