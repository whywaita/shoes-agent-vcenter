# VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the TPM CA certificate. When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.CaCertificate. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.CaCertificate. | 
**CertChain** | Pointer to [**VcenterTrustedInfrastructureX509CertChain**](VcenterTrustedInfrastructureX509CertChain.md) |  | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec(name string, ) *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetCertChain

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) GetCertChain() VcenterTrustedInfrastructureX509CertChain`

GetCertChain returns the CertChain field if non-nil, zero value otherwise.

### GetCertChainOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) GetCertChainOk() (*VcenterTrustedInfrastructureX509CertChain, bool)`

GetCertChainOk returns a tuple with the CertChain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertChain

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) SetCertChain(v VcenterTrustedInfrastructureX509CertChain)`

SetCertChain sets CertChain field to given value.

### HasCertChain

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2CaCertificatesCreateSpec) HasCertChain() bool`

HasCertChain returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


