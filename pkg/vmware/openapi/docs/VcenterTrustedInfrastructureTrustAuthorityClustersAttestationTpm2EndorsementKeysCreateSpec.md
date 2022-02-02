# VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the TPM endorsement key.   The unique name should be something that an administrator can use to easily identify the remote system. For example, the hostname, or hardware UUID.  When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.EndorsementKey. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.attestation.tpm2.EndorsementKey. | 
**PublicKey** | Pointer to **string** | TPM public endorsement key in PEM format. If unset EndorsementKeys.CreateSpec.certificate must be set. | [optional] 
**Certificate** | Pointer to **string** | TPM endorsement key certificate in PEM format.   When a endorsement key certificate is provided, it will be verified against the CA certificate list. Endorsement key certificates that are not signed by one of the CA certificates will be rejected.    Using this format allows for failures to be caught during configuration rather than later during attestation.  If unset EndorsementKeys.CreateSpec.public-key must be set. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec(name string, ) *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetCertificate

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationTpm2EndorsementKeysCreateSpec) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


