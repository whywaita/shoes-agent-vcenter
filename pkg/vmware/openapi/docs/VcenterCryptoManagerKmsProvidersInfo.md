# VcenterCryptoManagerKmsProvidersInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Health** | [**VcenterCryptoManagerKmsProvidersHealth**](VcenterCryptoManagerKmsProvidersHealth.md) |  | 
**Details** | [**[]VapiStdLocalizableMessage**](VapiStdLocalizableMessage.md) | Details regarding the health status of the provider.   When the provider Providers.Health is not NONE or OK, this field will provide actionable descriptions of the issues.  | 
**Constraints** | Pointer to [**VcenterCryptoManagerKmsProvidersConstraints**](VcenterCryptoManagerKmsProvidersConstraints.md) |  | [optional] 
**Type** | [**VcenterCryptoManagerKmsProvidersType**](VcenterCryptoManagerKmsProvidersType.md) |  | 
**NativeInfo** | Pointer to [**VcenterCryptoManagerKmsProvidersNativeProviderInfo**](VcenterCryptoManagerKmsProvidersNativeProviderInfo.md) |  | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersInfo

`func NewVcenterCryptoManagerKmsProvidersInfo(health VcenterCryptoManagerKmsProvidersHealth, details []VapiStdLocalizableMessage, type_ VcenterCryptoManagerKmsProvidersType, ) *VcenterCryptoManagerKmsProvidersInfo`

NewVcenterCryptoManagerKmsProvidersInfo instantiates a new VcenterCryptoManagerKmsProvidersInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersInfoWithDefaults

`func NewVcenterCryptoManagerKmsProvidersInfoWithDefaults() *VcenterCryptoManagerKmsProvidersInfo`

NewVcenterCryptoManagerKmsProvidersInfoWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHealth

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetHealth() VcenterCryptoManagerKmsProvidersHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetHealthOk() (*VcenterCryptoManagerKmsProvidersHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterCryptoManagerKmsProvidersInfo) SetHealth(v VcenterCryptoManagerKmsProvidersHealth)`

SetHealth sets Health field to given value.


### GetDetails

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetDetails() []VapiStdLocalizableMessage`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetDetailsOk() (*[]VapiStdLocalizableMessage, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *VcenterCryptoManagerKmsProvidersInfo) SetDetails(v []VapiStdLocalizableMessage)`

SetDetails sets Details field to given value.


### GetConstraints

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetConstraints() VcenterCryptoManagerKmsProvidersConstraints`

GetConstraints returns the Constraints field if non-nil, zero value otherwise.

### GetConstraintsOk

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetConstraintsOk() (*VcenterCryptoManagerKmsProvidersConstraints, bool)`

GetConstraintsOk returns a tuple with the Constraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraints

`func (o *VcenterCryptoManagerKmsProvidersInfo) SetConstraints(v VcenterCryptoManagerKmsProvidersConstraints)`

SetConstraints sets Constraints field to given value.

### HasConstraints

`func (o *VcenterCryptoManagerKmsProvidersInfo) HasConstraints() bool`

HasConstraints returns a boolean if a field has been set.

### GetType

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetType() VcenterCryptoManagerKmsProvidersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetTypeOk() (*VcenterCryptoManagerKmsProvidersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterCryptoManagerKmsProvidersInfo) SetType(v VcenterCryptoManagerKmsProvidersType)`

SetType sets Type field to given value.


### GetNativeInfo

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetNativeInfo() VcenterCryptoManagerKmsProvidersNativeProviderInfo`

GetNativeInfo returns the NativeInfo field if non-nil, zero value otherwise.

### GetNativeInfoOk

`func (o *VcenterCryptoManagerKmsProvidersInfo) GetNativeInfoOk() (*VcenterCryptoManagerKmsProvidersNativeProviderInfo, bool)`

GetNativeInfoOk returns a tuple with the NativeInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeInfo

`func (o *VcenterCryptoManagerKmsProvidersInfo) SetNativeInfo(v VcenterCryptoManagerKmsProvidersNativeProviderInfo)`

SetNativeInfo sets NativeInfo field to given value.

### HasNativeInfo

`func (o *VcenterCryptoManagerKmsProvidersInfo) HasNativeInfo() bool`

HasNativeInfo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


