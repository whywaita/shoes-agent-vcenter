# VcenterCryptoManagerKmsProvidersSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | Provider identifier When clients pass a value of this structure as a parameter, the field must be an identifier for the resource type: vcenter.crypto_manager.kms.provider. When operations return a value of this structure as a result, the field will be an identifier for the resource type: vcenter.crypto_manager.kms.provider. | 
**Type** | [**VcenterCryptoManagerKmsProvidersType**](VcenterCryptoManagerKmsProvidersType.md) |  | 
**Health** | [**VcenterCryptoManagerKmsProvidersHealth**](VcenterCryptoManagerKmsProvidersHealth.md) |  | 

## Methods

### NewVcenterCryptoManagerKmsProvidersSummary

`func NewVcenterCryptoManagerKmsProvidersSummary(provider string, type_ VcenterCryptoManagerKmsProvidersType, health VcenterCryptoManagerKmsProvidersHealth, ) *VcenterCryptoManagerKmsProvidersSummary`

NewVcenterCryptoManagerKmsProvidersSummary instantiates a new VcenterCryptoManagerKmsProvidersSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersSummaryWithDefaults

`func NewVcenterCryptoManagerKmsProvidersSummaryWithDefaults() *VcenterCryptoManagerKmsProvidersSummary`

NewVcenterCryptoManagerKmsProvidersSummaryWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *VcenterCryptoManagerKmsProvidersSummary) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *VcenterCryptoManagerKmsProvidersSummary) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *VcenterCryptoManagerKmsProvidersSummary) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetType

`func (o *VcenterCryptoManagerKmsProvidersSummary) GetType() VcenterCryptoManagerKmsProvidersType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VcenterCryptoManagerKmsProvidersSummary) GetTypeOk() (*VcenterCryptoManagerKmsProvidersType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VcenterCryptoManagerKmsProvidersSummary) SetType(v VcenterCryptoManagerKmsProvidersType)`

SetType sets Type field to given value.


### GetHealth

`func (o *VcenterCryptoManagerKmsProvidersSummary) GetHealth() VcenterCryptoManagerKmsProvidersHealth`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *VcenterCryptoManagerKmsProvidersSummary) GetHealthOk() (*VcenterCryptoManagerKmsProvidersHealth, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *VcenterCryptoManagerKmsProvidersSummary) SetHealth(v VcenterCryptoManagerKmsProvidersHealth)`

SetHealth sets Health field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


