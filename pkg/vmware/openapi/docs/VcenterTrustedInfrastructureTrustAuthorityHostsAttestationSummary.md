# VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SummaryType** | [**VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryType**](VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryType.md) |  | 
**Host** | Pointer to **string** | The trusted ESX on which the service runs. This field is optional and it is only relevant when the value of Attestation.Summary.summary-type is one of BRIEF, NORMAL, or FULL. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: HostSystem. When operations return a value of this structure as a result, the field will be an identifier for the resource type: HostSystem. | [optional] 
**Address** | Pointer to [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | [optional] 
**Group** | Pointer to **string** | The group ID determines which Attestation Service instances this Attestation Service can communicate with. This field is optional and it is only relevant when the value of Attestation.Summary.summary-type is one of NORMAL or FULL. | [optional] 
**Cluster** | Pointer to **string** | The opaque string identifier of the cluster in which the Attestation Service is part of. This field is optional and it is only relevant when the value of Attestation.Summary.summary-type is one of NORMAL or FULL. | [optional] 
**TrustedCA** | Pointer to [**VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) |  | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary(summaryType VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryType, ) *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary`

NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary`

NewVcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummaryType

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetSummaryType() VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryType`

GetSummaryType returns the SummaryType field if non-nil, zero value otherwise.

### GetSummaryTypeOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetSummaryTypeOk() (*VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryType, bool)`

GetSummaryTypeOk returns a tuple with the SummaryType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummaryType

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) SetSummaryType(v VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummaryType)`

SetSummaryType sets SummaryType field to given value.


### GetHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetAddress() VcenterTrustedInfrastructureNetworkAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetAddressOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) SetAddress(v VcenterTrustedInfrastructureNetworkAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetCluster() string`

GetCluster returns the Cluster field if non-nil, zero value otherwise.

### GetClusterOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetClusterOk() (*string, bool)`

GetClusterOk returns a tuple with the Cluster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) SetCluster(v string)`

SetCluster sets Cluster field to given value.

### HasCluster

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) HasCluster() bool`

HasCluster returns a boolean if a field has been set.

### GetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetTrustedCA() VcenterTrustedInfrastructureX509CertChain`

GetTrustedCA returns the TrustedCA field if non-nil, zero value otherwise.

### GetTrustedCAOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) GetTrustedCAOk() (*VcenterTrustedInfrastructureX509CertChain, bool)`

GetTrustedCAOk returns a tuple with the TrustedCA field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) SetTrustedCA(v VcenterTrustedInfrastructureX509CertChain)`

SetTrustedCA sets TrustedCA field to given value.

### HasTrustedCA

`func (o *VcenterTrustedInfrastructureTrustAuthorityHostsAttestationSummary) HasTrustedCA() bool`

HasTrustedCA returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


