# VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CertChain** | [**VcenterCertificateManagementX509CertChain**](VcenterCertificateManagementX509CertChain.md) |  | 
**Chain** | Pointer to **string** | Unique identifier for this trusted root. Client can specify at creation as long as it is unique, otherwise one will be generated. An example of a client providing the identifier would be if this trusted root is associated with a VC trust. In this case the identifier would be the domain id. A unique id will be generated if not given. | [optional] 

## Methods

### NewVcenterCertificateManagementVcenterTrustedRootChainsCreateSpec

`func NewVcenterCertificateManagementVcenterTrustedRootChainsCreateSpec(certChain VcenterCertificateManagementX509CertChain, ) *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec`

NewVcenterCertificateManagementVcenterTrustedRootChainsCreateSpec instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCertificateManagementVcenterTrustedRootChainsCreateSpecWithDefaults

`func NewVcenterCertificateManagementVcenterTrustedRootChainsCreateSpecWithDefaults() *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec`

NewVcenterCertificateManagementVcenterTrustedRootChainsCreateSpecWithDefaults instantiates a new VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertChain

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) GetCertChain() VcenterCertificateManagementX509CertChain`

GetCertChain returns the CertChain field if non-nil, zero value otherwise.

### GetCertChainOk

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) GetCertChainOk() (*VcenterCertificateManagementX509CertChain, bool)`

GetCertChainOk returns a tuple with the CertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertChain

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) SetCertChain(v VcenterCertificateManagementX509CertChain)`

SetCertChain sets CertChain field to given value.


### GetChain

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *VcenterCertificateManagementVcenterTrustedRootChainsCreateSpec) HasChain() bool`

HasChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


