# VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecType**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecType.md) |  | 
**Description** | Pointer to **string** | Description of the key server. If unset, description will not be added. | [optional] 
**ProxyServer** | Pointer to [**VcenterTrustedInfrastructureNetworkAddress**](VcenterTrustedInfrastructureNetworkAddress.md) |  | [optional] 
**ConnectionTimeout** | Pointer to **int64** | Connection timeout in seconds. If unset, connection timeout will not be set. | [optional] 
**KmipServer** | Pointer to [**VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec**](VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec.md) |  | [optional] 

## Methods

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec(type_ VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecType, ) *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecWithDefaults

`func NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecWithDefaults() *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec`

NewVcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecWithDefaults instantiates a new VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetType() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetTypeOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) SetType(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpecType)`

SetType sets Type field to given value.


### GetDescription

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetProxyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetProxyServer() VcenterTrustedInfrastructureNetworkAddress`

GetProxyServer returns the ProxyServer field if non-nil, zero value otherwise.

### GetProxyServerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetProxyServerOk() (*VcenterTrustedInfrastructureNetworkAddress, bool)`

GetProxyServerOk returns a tuple with the ProxyServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) SetProxyServer(v VcenterTrustedInfrastructureNetworkAddress)`

SetProxyServer sets ProxyServer field to given value.

### HasProxyServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) HasProxyServer() bool`

HasProxyServer returns a boolean if a field has been set.

### GetConnectionTimeout

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetConnectionTimeout() int64`

GetConnectionTimeout returns the ConnectionTimeout field if non-nil, zero value otherwise.

### GetConnectionTimeoutOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetConnectionTimeoutOk() (*int64, bool)`

GetConnectionTimeoutOk returns a tuple with the ConnectionTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionTimeout

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) SetConnectionTimeout(v int64)`

SetConnectionTimeout sets ConnectionTimeout field to given value.

### HasConnectionTimeout

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) HasConnectionTimeout() bool`

HasConnectionTimeout returns a boolean if a field has been set.

### GetKmipServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetKmipServer() VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec`

GetKmipServer returns the KmipServer field if non-nil, zero value otherwise.

### GetKmipServerOk

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) GetKmipServerOk() (*VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec, bool)`

GetKmipServerOk returns a tuple with the KmipServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmipServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) SetKmipServer(v VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKmipServerCreateSpec)`

SetKmipServer sets KmipServer field to given value.

### HasKmipServer

`func (o *VcenterTrustedInfrastructureTrustAuthorityClustersKmsProvidersKeyServerCreateSpec) HasKmipServer() bool`

HasKmipServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


