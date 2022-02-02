# VcenterCryptoManagerHostsKmsProvidersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | [**VcenterCryptoManagerHostsKmsProvidersHealth**](VcenterCryptoManagerHostsKmsProvidersHealth.md) |  | 
**Details** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Details regarding the health status of the provider. When the provider Providers.Health is not OK or NONE, this field will provide actionable descriptions of the issues. | 
**Type** | [**VcenterCryptoManagerHostsKmsProvidersType**](VcenterCryptoManagerHostsKmsProvidersType.md) |  | 
**NativeInfo** | Pointer to [**VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo**](VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo.md) |  | [optional] 

## Methods

### NewVcenterCryptoManagerHostsKmsProvidersInfo

`func NewVcenterCryptoManagerHostsKmsProvidersInfo(health VcenterCryptoManagerHostsKmsProvidersHealth, details []VapiStdLocalizableMessage, type_ VcenterCryptoManagerHostsKmsProvidersType, ) *VcenterCryptoManagerHostsKmsProvidersInfo`

NewVcenterCryptoManagerHostsKmsProvidersInfo instantiates a new VcenterCryptoManagerHostsKmsProvidersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerHostsKmsProvidersInfoWithDefaults

`func NewVcenterCryptoManagerHostsKmsProvidersInfoWithDefaults() *VcenterCryptoManagerHostsKmsProvidersInfo`

NewVcenterCryptoManagerHostsKmsProvidersInfoWithDefaults instantiates a new VcenterCryptoManagerHostsKmsProvidersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetHealth() VcenterCryptoManagerHostsKmsProvidersHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetHealthOk() (*VcenterCryptoManagerHostsKmsProvidersHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) SetHealth(v VcenterCryptoManagerHostsKmsProvidersHealth)`

SetHealth sets Health field to given value.


### GetDetails

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetDetails() []VapiStdLocalizableMessage`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetDetailsOk() (*[]VapiStdLocalizableMessage, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) SetDetails(v []VapiStdLocalizableMessage)`

SetDetails sets Details field to given value.


### GetType

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetType() VcenterCryptoManagerHostsKmsProvidersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetTypeOk() (*VcenterCryptoManagerHostsKmsProvidersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) SetType(v VcenterCryptoManagerHostsKmsProvidersType)`

SetType sets Type field to given value.


### GetNativeInfo

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetNativeInfo() VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo`

GetNativeInfo returns the NativeInfo field if non-nil, zero value otherwise.

### GetNativeInfoOk

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) GetNativeInfoOk() (*VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo, bool)`

GetNativeInfoOk returns a tuple with the NativeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeInfo

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) SetNativeInfo(v VcenterCryptoManagerHostsKmsProvidersNativeProviderInfo)`

SetNativeInfo sets NativeInfo field to given value.

### HasNativeInfo

`func (o *VcenterCryptoManagerHostsKmsProvidersInfo) HasNativeInfo() bool`

HasNativeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


