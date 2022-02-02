# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerNames** | Pointer to **[]string** | Names that key server must have to match the filter (see CurrentPeerCertificates.Summary.server-name). If unset or empty, key servers with any name match the filter. | [optional] 
**Trusted** | Pointer to **bool** | Trust status that server certificates must have to match the filter (see CurrentPeerCertificates.Summary.trusted). If unset, trusted and untrusted server certificates match the filter. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerNames

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) GetServerNames() []string`

GetServerNames returns the ServerNames field if non-nil, zero value otherwise.

### GetServerNamesOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) GetServerNamesOk() (*[]string, bool)`

GetServerNamesOk returns a tuple with the ServerNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNames

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) SetServerNames(v []string)`

SetServerNames sets ServerNames field to given value.

### HasServerNames

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) HasServerNames() bool`

HasServerNames returns a boolean if a field has been set.

### GetTrusted

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) GetTrusted() bool`

GetTrusted returns the Trusted field if non-nil, zero value otherwise.

### GetTrustedOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) GetTrustedOk() (*bool, bool)`

GetTrustedOk returns a tuple with the Trusted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrusted

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) SetTrusted(v bool)`

SetTrusted sets Trusted field to given value.

### HasTrusted

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCurrentPeerCertificatesFilterSpec) HasTrusted() bool`

HasTrusted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


