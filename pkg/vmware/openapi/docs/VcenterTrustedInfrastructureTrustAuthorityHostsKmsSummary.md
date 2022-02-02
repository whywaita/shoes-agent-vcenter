# VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryType** | [**VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryType**](VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryType.md) |  | 
**Host** | Pointer to **string** | The trusted ESX on which the service runs. This field is optional and it is only relevant when the value of Kms.Summary.summary-type is one of BRIEF, NORMAL, or FULL. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | [optional] 
**Address** | Pointer to [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | [optional] 
**Group** | Pointer to **string** | The group ID determines which Attestation Service instances this Key Provider Service can communicate with. This field is optional and it is only relevant when the value of Kms.Summary.summary-type is one of NORMAL or FULL. | [optional] 
**Cluster** | Pointer to **string** | The opaque string identifier of the cluster in which the Key Provider Service is part of. This field is optional and it is only relevant when the value of Kms.Summary.summary-type is one of NORMAL or FULL. | [optional] 
**TrustedCA** | Pointer to [**VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) |  | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary(summaryType VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryType, ) *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary`

NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary`

NewVcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryType

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetSummaryType() VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryType`

GetSummaryType returns the SummaryType field if non-nil, zero value otherwise.

### GetSummaryTypeOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetSummaryTypeOk() (*VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryType, bool)`

GetSummaryTypeOk returns a tuple with the SummaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryType

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) SetSummaryType(v VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummaryType)`

SetSummaryType sets SummaryType field to given value.


### GetHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain`

GetTrustedCA returns the TrustedCA field if non-nil, zero value otherwise.

### GetTrustedCAOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool)`

GetTrustedCAOk returns a tuple with the TrustedCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain)`

SetTrustedCA sets TrustedCA field to given value.

### HasTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsKmsSummary) HasTrustedCA() bool`

HasTrustedCA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


