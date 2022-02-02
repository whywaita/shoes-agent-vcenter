# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to [**VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType.md) |  | [optional] 
**Description** | Pointer to **string** | Description of the key server.    If unset, description will remain unchanged. | [optional] 
**ProxyServer** | Pointer to [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** | Connection timeout in seconds.    If unset, connection timeout will remain unchanged. | [optional] 
**KmipServer** | Pointer to [**VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec.md) |  | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetType() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetTypeOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) SetType(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpecType)`

SetType sets Type field to given value.

### HasType

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) HasType() bool`

HasType returns a boolean if a field has been set.

### GetDescription

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProxyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetProxyServer() VcenterTrustedInfrastructureNetworkAddress`

GetProxyServer returns the ProxyServer field if non-nil, zero value otherwise.

### GetProxyServerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetProxyServerOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetProxyServerOk returns a tuple with the ProxyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) SetProxyServer(v VcenterTrustedInfrastructureNetworkAddress)`

SetProxyServer sets ProxyServer field to given value.

### HasProxyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) HasProxyServer() bool`

HasProxyServer returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetKmipServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetKmipServer() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec`

GetKmipServer returns the KmipServer field if non-nil, zero value otherwise.

### GetKmipServerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) GetKmipServerOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec, bool)`

GetKmipServerOk returns a tuple with the KmipServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) SetKmipServer(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerUpdateSpec)`

SetKmipServer sets KmipServer field to given value.

### HasKmipServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerUpdateSpec) HasKmipServer() bool`

HasKmipServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


