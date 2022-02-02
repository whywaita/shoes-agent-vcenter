# VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) | The certificates used by the vCenter STS to sign tokens. | 
**IssuerAlias** | **string** | A user-friendly alias of the service which created and signed the security token. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: esx.authentication.trust.security-token-issuer. When operations return a value of this structure as a result, the field will be an identifier for the resource type: esx.authentication.trust.security-token-issuer. | 
**Issuer** | **string** | The service which created and signed the security token. | 
**Principal** | [**VcenterTrustedInfrastructureStsPrincipal**](VcenterTrustedInfrastructureStsPrincipal.md) |  | 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec(certificates []VcenterTrustedInfrastructureX509CertChain, issuerAlias string, issuer string, principal VcenterTrustedInfrastructureStsPrincipal, ) *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetCertificates() []VcenterTrustedInfrastructureX509CertChain`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetCertificatesOk() (*[]VcenterTrustedInfrastructureX509CertChain, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) SetCertificates(v []VcenterTrustedInfrastructureX509CertChain)`

SetCertificates sets Certificates field to given value.


### GetIssuerAlias

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetIssuerAlias() string`

GetIssuerAlias returns the IssuerAlias field if non-nil, zero value otherwise.

### GetIssuerAliasOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetIssuerAliasOk() (*string, bool)`

GetIssuerAliasOk returns a tuple with the IssuerAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerAlias

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) SetIssuerAlias(v string)`

SetIssuerAlias sets IssuerAlias field to given value.


### GetIssuer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetPrincipal

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetPrincipal() VcenterTrustedInfrastructureStsPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) GetPrincipalOk() (*VcenterTrustedInfrastructureStsPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersConsumerPrincipalsCreateSpec) SetPrincipal(v VcenterTrustedInfrastructureStsPrincipal)`

SetPrincipal sets Principal field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


