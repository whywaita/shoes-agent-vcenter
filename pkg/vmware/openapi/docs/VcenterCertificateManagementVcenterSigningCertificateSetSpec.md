# VcenterCertificateManagementVcenterSigningCertificateSetSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SigningCertChain** | [**VcenterCertificateManagementX509CertChain**](VcenterCertificateManagementX509CertChain.md) |  | 
**PrivateKey** | **string** | The corresponding unencrypted PKCS#8 private key in base64-encoded PEM format. | 

## Methods

### NewVcenterCertificateManagementVcenterSigningCertificateSetSpec

`func NewVcenterCertificateManagementVcenterSigningCertificateSetSpec(signingCertChain VcenterCertificateManagementX509CertChain, privateKey string, ) *VcenterCertificateManagementVcenterSigningCertificateSetSpec`

NewVcenterCertificateManagementVcenterSigningCertificateSetSpec instantiates a new VcenterCertificateManagementVcenterSigningCertificateSetSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterSigningCertificateSetSpecWithDefaults

`func NewVcenterCertificateManagementVcenterSigningCertificateSetSpecWithDefaults() *VcenterCertificateManagementVcenterSigningCertificateSetSpec`

NewVcenterCertificateManagementVcenterSigningCertificateSetSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterSigningCertificateSetSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSigningCertChain

`func (o *VcenterCertificateManagementVcenterSigningCertificateSetSpec) GetSigningCertChain() VcenterCertificateManagementX509CertChain`

GetSigningCertChain returns the SigningCertChain field if non-nil, zero value otherwise.

### GetSigningCertChainOk

`func (o *VcenterCertificateManagementVcenterSigningCertificateSetSpec) GetSigningCertChainOk() (*VcenterCertificateManagementX509CertChain, bool)`

GetSigningCertChainOk returns a tuple with the SigningCertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSigningCertChain

`func (o *VcenterCertificateManagementVcenterSigningCertificateSetSpec) SetSigningCertChain(v VcenterCertificateManagementX509CertChain)`

SetSigningCertChain sets SigningCertChain field to given value.


### GetPrivateKey

`func (o *VcenterCertificateManagementVcenterSigningCertificateSetSpec) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *VcenterCertificateManagementVcenterSigningCertificateSetSpec) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *VcenterCertificateManagementVcenterSigningCertificateSetSpec) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


