# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Name of the provider.   A unique string chosen by the client.  When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.kms.Provider. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.trusted_infrastructure.trust_authority_clusters.kms.Provider. | 
**MasterKeyId** | **string** | Master key ID created for the provider.   A unique Key ID.  | 
**KeyServer** | [**VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec.md) |  | 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec(provider string, masterKeyId string, keyServer VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec, ) *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetMasterKeyId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.


### GetKeyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) GetKeyServer() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec`

GetKeyServer returns the KeyServer field if non-nil, zero value otherwise.

### GetKeyServerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) GetKeyServerOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec, bool)`

GetKeyServerOk returns a tuple with the KeyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersCreateSpec) SetKeyServer(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec)`

SetKeyServer sets KeyServer field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


