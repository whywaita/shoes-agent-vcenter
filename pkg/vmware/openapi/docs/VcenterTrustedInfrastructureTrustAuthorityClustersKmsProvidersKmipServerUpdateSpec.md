# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer.md) | List of KMIP compliant key servers.   Key servers must be configured for active-active replication. If the server port is unset, a default value for KMIP&#39;s port will be used.     If unset, server configuration will remain unchanged. | [optional] 
**Username** | Pointer to **string** | Username for authentication.    If unset, username will remain unchanged. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) GetServers() []VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) GetServersOk() (*[]VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) SetServers(v []VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer)`

SetServers sets Servers field to given value.

### HasServers

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetUsername

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


