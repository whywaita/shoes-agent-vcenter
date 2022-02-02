# VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Service** | **string** | The service&#39;s unique identifier. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.kms.Service. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.kms.Service. | 
**Address** | [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | 
**Group** | **string** | The group determines the Attestation Service instances this Key Provider Service can accept reports from. | 
**TrustAuthorityCluster** | **string** | The cluster specifies the Trust Authority Cluster this Key Provider Service belongs to. | 

## Methods

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesSummary

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesSummary(service string, address VcenterTrustedInfrastructureNetworkAddress, group string, trustAuthorityCluster string, ) *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesSummary instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustedClustersKmsServicesSummaryWithDefaults

`func NewVcenterTrustedInfrastructureTrustedClustersKmsServicesSummaryWithDefaults() *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary`

NewVcenterTrustedInfrastructureTrustedClustersKmsServicesSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetService

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) SetService(v string)`

SetService sets Service field to given value.


### GetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.


### GetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) SetGroup(v string)`

SetGroup sets Group field to given value.


### GetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetTrustAuthorityCluster() string`

GetTrustAuthorityCluster returns the TrustAuthorityCluster field if non-nil, zero value otherwise.

### GetTrustAuthorityClusterOk

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) GetTrustAuthorityClusterOk() (*string, bool)`

GetTrustAuthorityClusterOk returns a tuple with the TrustAuthorityCluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustAuthorityCluster

`func (o *VcenterTrustedInfrastructureTrustedClustersKmsServicesSummary) SetTrustAuthorityCluster(v string)`

SetTrustAuthorityCluster sets TrustAuthorityCluster field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


