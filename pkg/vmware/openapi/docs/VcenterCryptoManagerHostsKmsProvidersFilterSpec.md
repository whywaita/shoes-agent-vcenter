# VcenterCryptoManagerHostsKmsProvidersFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to **[]string** | Provider identifiers. If unset or empty, the result will not be filtered by provider identifier. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.crypto_manager.kms.provider. | [optional] 
**Health** | Pointer to [**[]VcenterCryptoManagerHostsKmsProvidersHealth**](VcenterCryptoManagerHostsKmsProvidersHealth.md) | Provider health status. If unset or empty, the result will not be filtered by provider health status. | [optional] 
**Types** | Pointer to [**[]VcenterCryptoManagerHostsKmsProvidersType**](VcenterCryptoManagerHostsKmsProvidersType.md) | Provider types. If unset or empty, the result will not be filtered by provider type. | [optional] 

## Methods

### NewVcenterCryptoManagerHostsKmsProvidersFilterSpec

`func NewVcenterCryptoManagerHostsKmsProvidersFilterSpec() *VcenterCryptoManagerHostsKmsProvidersFilterSpec`

NewVcenterCryptoManagerHostsKmsProvidersFilterSpec instantiates a new VcenterCryptoManagerHostsKmsProvidersFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerHostsKmsProvidersFilterSpecWithDefaults

`func NewVcenterCryptoManagerHostsKmsProvidersFilterSpecWithDefaults() *VcenterCryptoManagerHostsKmsProvidersFilterSpec`

NewVcenterCryptoManagerHostsKmsProvidersFilterSpecWithDefaults instantiates a new VcenterCryptoManagerHostsKmsProvidersFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetHealth() []VcenterCryptoManagerHostsKmsProvidersHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetHealthOk() (*[]VcenterCryptoManagerHostsKmsProvidersHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) SetHealth(v []VcenterCryptoManagerHostsKmsProvidersHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetTypes

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetTypes() []VcenterCryptoManagerHostsKmsProvidersType`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) GetTypesOk() (*[]VcenterCryptoManagerHostsKmsProvidersType, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) SetTypes(v []VcenterCryptoManagerHostsKmsProvidersType)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *VcenterCryptoManagerHostsKmsProvidersFilterSpec) HasTypes() bool`

HasTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


