# VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **[]string** | Search criteria by ESX base image version numbers. version if {@term.unset} return all ESX version numbers. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.BaseImage. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.attestation.os.esx.BaseImage. | [optional] 
**DisplayName** | Pointer to **[]string** | Search criteria by ESX base image version version numbers. displayName if {@term.unset} return all ESX display version numbers. | [optional] 
**Health** | Pointer to [**[]VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth**](VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth.md) | Search criteria by health indicator. health if {@term.unset} return all health indicators. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) GetVersion() []string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) GetVersionOk() (*[]string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) SetVersion(v []string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetDisplayName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) GetDisplayName() []string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) GetDisplayNameOk() (*[]string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) SetDisplayName(v []string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHealth

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) GetHealth() []VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) GetHealthOk() (*[]VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) SetHealth(v []VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersAttestationOsEsxBaseImagesFilterSpec) HasHealth() bool`

HasHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


