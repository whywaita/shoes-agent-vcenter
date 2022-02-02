# VcenterCryptoManagerKmsProvidersFilterSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Providers** | Pointer to **[]string** | Provider identifiers. If unset or empty, the result will not be filtered by provider identifier. When clients pass a value of this structure as a parameter, the field must contain identifiers for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will contain identifiers for the resource type: vcenter.crypto_manager.kms.provider. | [optional] 
**Health** | Pointer to [**[]VcenterCryptoManagerKmsProvidersHealth**](VcenterCryptoManagerKmsProvidersHealth.md) | Provider health status. If unset or empty, the result will not be filtered by provider health status. | [optional] 

## Methods

### NewVcenterCryptoManagerKmsProvidersFilterSpec

`func NewVcenterCryptoManagerKmsProvidersFilterSpec() *VcenterCryptoManagerKmsProvidersFilterSpec`

NewVcenterCryptoManagerKmsProvidersFilterSpec instantiates a new VcenterCryptoManagerKmsProvidersFilterSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersFilterSpecWithDefaults

`func NewVcenterCryptoManagerKmsProvidersFilterSpecWithDefaults() *VcenterCryptoManagerKmsProvidersFilterSpec`

NewVcenterCryptoManagerKmsProvidersFilterSpecWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersFilterSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProviders

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) GetProviders() []string`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) GetProvidersOk() (*[]string, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) SetProviders(v []string)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetHealth

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) GetHealth() []VcenterCryptoManagerKmsProvidersHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) GetHealthOk() (*[]VcenterCryptoManagerKmsProvidersHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) SetHealth(v []VcenterCryptoManagerKmsProvidersHealth)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *VcenterCryptoManagerKmsProvidersFilterSpec) HasHealth() bool`

HasHealth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


