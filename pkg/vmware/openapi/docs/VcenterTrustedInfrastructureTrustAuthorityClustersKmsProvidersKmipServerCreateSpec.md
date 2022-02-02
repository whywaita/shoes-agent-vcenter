# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | [**[]VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer.md) | List of Key Management Interoperability Protocol (KMIP) compliant key servers.   Key servers must be configured for active-active replication. If the server port is unset, a default value for KMIP&#39;s port will be used.  | 
**Username** | Pointer to **string** | Username for authentication. If unset, no username will be added. | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec(servers []VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer, ) *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) GetServers() []VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) GetServersOk() (*[]VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) SetServers(v []VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersServer)`

SetServers sets Servers field to given value.


### GetUsername

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


