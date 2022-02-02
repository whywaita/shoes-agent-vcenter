# VcenterCryptoManagerHostsKmsProvidersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Identifier of the provider When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.crypto_manager.kms.provider. | 
**Type** | [**VcenterCryptoManagerHostsKmsProvidersType**](VcenterCryptoManagerHostsKmsProvidersType.md) |  | 
**Health** | [**VcenterCryptoManagerHostsKmsProvidersHealth**](VcenterCryptoManagerHostsKmsProvidersHealth.md) |  | 

## Methods

### NewVcenterCryptoManagerHostsKmsProvidersSummary

`func NewVcenterCryptoManagerHostsKmsProvidersSummary(provider string, type_ VcenterCryptoManagerHostsKmsProvidersType, health VcenterCryptoManagerHostsKmsProvidersHealth, ) *VcenterCryptoManagerHostsKmsProvidersSummary`

NewVcenterCryptoManagerHostsKmsProvidersSummary instantiates a new VcenterCryptoManagerHostsKmsProvidersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerHostsKmsProvidersSummaryWithDefaults

`func NewVcenterCryptoManagerHostsKmsProvidersSummaryWithDefaults() *VcenterCryptoManagerHostsKmsProvidersSummary`

NewVcenterCryptoManagerHostsKmsProvidersSummaryWithDefaults instantiates a new VcenterCryptoManagerHostsKmsProvidersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) GetType() VcenterCryptoManagerHostsKmsProvidersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) GetTypeOk() (*VcenterCryptoManagerHostsKmsProvidersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) SetType(v VcenterCryptoManagerHostsKmsProvidersType)`

SetType sets Type field to given value.


### GetHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) GetHealth() VcenterCryptoManagerHostsKmsProvidersHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) GetHealthOk() (*VcenterCryptoManagerHostsKmsProvidersHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterCryptoManagerHostsKmsProvidersSummary) SetHealth(v VcenterCryptoManagerHostsKmsProvidersHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


