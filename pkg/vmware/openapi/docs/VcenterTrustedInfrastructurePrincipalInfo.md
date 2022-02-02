# VcenterTrustedInfrastructurePrincipalInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificates** | [**[]VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) | The certificates used by the STS to sign tokens for this vCenter. | 
**Issuer** | **string** | The service which created and signed the security token. | 
**Principal** | [**VcenterTrustedInfrastructureStsPrincipal**](VcenterTrustedInfrastructureStsPrincipal.md) |  | 
**Name** | **string** | The user-friednly name of the vCenter. | 

## Methods

### NewVcenterTrustedInfrastructurePrincipalInfo

`func NewVcenterTrustedInfrastructurePrincipalInfo(certificates []VcenterTrustedInfrastructureX509CertChain, issuer string, principal VcenterTrustedInfrastructureStsPrincipal, name string, ) *VcenterTrustedInfrastructurePrincipalInfo`

NewVcenterTrustedInfrastructurePrincipalInfo instantiates a new VcenterTrustedInfrastructurePrincipalInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructurePrincipalInfoWithDefaults

`func NewVcenterTrustedInfrastructurePrincipalInfoWithDefaults() *VcenterTrustedInfrastructurePrincipalInfo`

NewVcenterTrustedInfrastructurePrincipalInfoWithDefaults instantiates a new VcenterTrustedInfrastructurePrincipalInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificates

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetCertificates() []VcenterTrustedInfrastructureX509CertChain`

GetCertificates returns the Certificates field if non-nil, zero value otherwise.

### GetCertificatesOk

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetCertificatesOk() (*[]VcenterTrustedInfrastructureX509CertChain, bool)`

GetCertificatesOk returns a tuple with the Certificates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificates

`func (o *VcenterTrustedInfrastructurePrincipalInfo) SetCertificates(v []VcenterTrustedInfrastructureX509CertChain)`

SetCertificates sets Certificates field to given value.


### GetIssuer

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VcenterTrustedInfrastructurePrincipalInfo) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetPrincipal

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetPrincipal() VcenterTrustedInfrastructureStsPrincipal`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetPrincipalOk() (*VcenterTrustedInfrastructureStsPrincipal, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *VcenterTrustedInfrastructurePrincipalInfo) SetPrincipal(v VcenterTrustedInfrastructureStsPrincipal)`

SetPrincipal sets Principal field to given value.


### GetName

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterTrustedInfrastructurePrincipalInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterTrustedInfrastructurePrincipalInfo) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


