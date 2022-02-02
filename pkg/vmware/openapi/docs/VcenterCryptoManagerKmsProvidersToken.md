# VcenterCryptoManagerKmsProvidersToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | A one-time, short-lived token required in \&quot;Authorization\&quot; field of the HTTP header of the request to the url.   After the token expires, any attempt to download the configuration with said token will fail.  | 
**Expiry** | **time.Time** | Expiry time of the token | 

## Methods

### NewVcenterCryptoManagerKmsProvidersToken

`func NewVcenterCryptoManagerKmsProvidersToken(token string, expiry time.Time, ) *VcenterCryptoManagerKmsProvidersToken`

NewVcenterCryptoManagerKmsProvidersToken instantiates a new VcenterCryptoManagerKmsProvidersToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterCryptoManagerKmsProvidersTokenWithDefaults

`func NewVcenterCryptoManagerKmsProvidersTokenWithDefaults() *VcenterCryptoManagerKmsProvidersToken`

NewVcenterCryptoManagerKmsProvidersTokenWithDefaults instantiates a new VcenterCryptoManagerKmsProvidersToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *VcenterCryptoManagerKmsProvidersToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *VcenterCryptoManagerKmsProvidersToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *VcenterCryptoManagerKmsProvidersToken) SetToken(v string)`

SetToken sets Token field to given value.


### GetExpiry

`func (o *VcenterCryptoManagerKmsProvidersToken) GetExpiry() time.Time`

GetExpiry returns the Expiry field if non-nil, zero value otherwise.

### GetExpiryOk

`func (o *VcenterCryptoManagerKmsProvidersToken) GetExpiryOk() (*time.Time, bool)`

GetExpiryOk returns a tuple with the Expiry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiry

`func (o *VcenterCryptoManagerKmsProvidersToken) SetExpiry(v time.Time)`

SetExpiry sets Expiry field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


