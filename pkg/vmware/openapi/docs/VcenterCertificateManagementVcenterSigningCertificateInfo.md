# VcenterCertificateManagementVcenterSigningCertificateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveCertChain** | [**VcenterCertificateManagementX509CertChain**](VcenterCertificateManagementX509CertChain.md) |  | 
**SigningCertChains** | [**[]VcenterCertificateManagementX509CertChain**](VcenterCertificateManagementX509CertChain.md) | List of signing certificate chains for validating vCenter-issued tokens. The list contains X509 certificate chains, each of which is ordered and contains the leaf, intermediate and root certs needed for the complete chain of trust. The leaf certificate is first in the chain and should be used for verifying vCenter-issued tokens. | 

## Methods

### NewVcenterCertificateManagementVcenterSigningCertificateInfo

`func NewVcenterCertificateManagementVcenterSigningCertificateInfo(activeCertChain VcenterCertificateManagementX509CertChain, signingCertChains []VcenterCertificateManagementX509CertChain, ) *VcenterCertificateManagementVcenterSigningCertificateInfo`

NewVcenterCertificateManagementVcenterSigningCertificateInfo instantiates a new VcenterCertificateManagementVcenterSigningCertificateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterSigningCertificateInfoWithDefaults

`func NewVcenterCertificateManagementVcenterSigningCertificateInfoWithDefaults() *VcenterCertificateManagementVcenterSigningCertificateInfo`

NewVcenterCertificateManagementVcenterSigningCertificateInfoWithDefaults instantiates a new VcenterCertificateManagementVcenterSigningCertificateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveCertChain

`func (o *VcenterCertificateManagementVcenterSigningCertificateInfo) GetActiveCertChain() VcenterCertificateManagementX509CertChain`

GetActiveCertChain returns the ActiveCertChain field if non-nil, zero value otherwise.

### GetActiveCertChainOk

`func (o *VcenterCertificateManagementVcenterSigningCertificateInfo) GetActiveCertChainOk() (*VcenterCertificateManagementX509CertChain, bool)`

GetActiveCertChainOk returns a tuple with the ActiveCertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveCertChain

`func (o *VcenterCertificateManagementVcenterSigningCertificateInfo) SetActiveCertChain(v VcenterCertificateManagementX509CertChain)`

SetActiveCertChain sets ActiveCertChain field to given value.


### GetSigningCertChains

`func (o *VcenterCertificateManagementVcenterSigningCertificateInfo) GetSigningCertChains() []VcenterCertificateManagementX509CertChain`

GetSigningCertChains returns the SigningCertChains field if non-nil, zero value otherwise.

### GetSigningCertChainsOk

`func (o *VcenterCertificateManagementVcenterSigningCertificateInfo) GetSigningCertChainsOk() (*[]VcenterCertificateManagementX509CertChain, bool)`

GetSigningCertChainsOk returns a tuple with the SigningCertChains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertChains

`func (o *VcenterCertificateManagementVcenterSigningCertificateInfo) SetSigningCertChains(v []VcenterCertificateManagementX509CertChain)`

SetSigningCertChains sets SigningCertChains field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


