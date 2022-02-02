# VcenterTrustedInfrastructureKmsServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | The service&#39;s unique identifier. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.kms.Service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.kms.Service. | 
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**Group** | **string** | The group determines the Attestation Service instances this Key Provider Service can accept reports from. | 
**TrustAuthorityCluster** | **string** | The cluster specifies the Trust Authority Cluster this Key Provider service belongs to. | 

## Methods

### NewVcenterTrustedInfrastructureKmsServicesSummary

`func NewVcenterTrustedInfrastructureKmsServicesSummary(service string, address VcenterTrustedInfrastructureNetworkAddress, group string, trustAuthorityCluster string, ) *VcenterTrustedInfrastructureKmsServicesSummary`

NewVcenterTrustedInfrastructureKmsServicesSummary instantiates a new VcenterTrustedInfrastructureKmsServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureKmsServicesSummaryWithDefaults

`func NewVcenterTrustedInfrastructureKmsServicesSummaryWithDefaults() *VcenterTrustedInfrastructureKmsServicesSummary`

NewVcenterTrustedInfrastructureKmsServicesSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureKmsServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) SetService(v string)`

SetService sets Service field to given value.


### GetAddress

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureKmsServicesSummary) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


