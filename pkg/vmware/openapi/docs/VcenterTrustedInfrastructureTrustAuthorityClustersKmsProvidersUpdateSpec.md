# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKeyId** | Pointer to **string** | Master key identifier created for the provider.   A unique Key identifier.     If unset, masterKeyId will remain unchanged. | [optional] 
**KeyServer** | Pointer to [**VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKeyId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetMasterKeyId() string`

GetMasterKeyId returns the MasterKeyId field if non-nil, zero value otherwise.

### GetMasterKeyIdOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetMasterKeyIdOk() (*string, bool)`

GetMasterKeyIdOk returns a tuple with the MasterKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKeyId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) SetMasterKeyId(v string)`

SetMasterKeyId sets MasterKeyId field to given value.

### HasMasterKeyId

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) HasMasterKeyId() bool`

HasMasterKeyId returns a boolean if a field has been set.

### GetKeyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetKeyServer() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec`

GetKeyServer returns the KeyServer field if non-nil, zero value otherwise.

### GetKeyServerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) GetKeyServerOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec, bool)`

GetKeyServerOk returns a tuple with the KeyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) SetKeyServer(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec)`

SetKeyServer sets KeyServer field to given value.

### HasKeyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersUpdateSpec) HasKeyServer() bool`

HasKeyServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


