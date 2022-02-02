# VcenterTokenserviceTokenExchangeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The security token issued by the server in response to the token exchange request. Token is base64-encoded. | 
**IssuedTokenType** | **string** | An identifier, that indicates the type of the security token in the {@link Info#access_token} parameter. | 
**TokenType** | **string** | A case-insensitive value specifying the method of using the access token issued. | 
**ExpiresIn** | Pointer to **int64** | The validity lifetime, in seconds, of the token issued by the server. | [optional] 
**Scope** | Pointer to **string** | Scope of the issued security token. | [optional] 
**RefreshToken** | Pointer to **string** | A refresh token can be issued in cases where the client of the token exchange needs the ability to access a resource even when the original credential is no longer valid. | [optional] 

## Methods

### NewVcenterTokenserviceTokenExchangeInfo

`func NewVcenterTokenserviceTokenExchangeInfo(accessToken string, issuedTokenType string, tokenType string, ) *VcenterTokenserviceTokenExchangeInfo`

NewVcenterTokenserviceTokenExchangeInfo instantiates a new VcenterTokenserviceTokenExchangeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterTokenserviceTokenExchangeInfoWithDefaults

`func NewVcenterTokenserviceTokenExchangeInfoWithDefaults() *VcenterTokenserviceTokenExchangeInfo`

NewVcenterTokenserviceTokenExchangeInfoWithDefaults instantiates a new VcenterTokenserviceTokenExchangeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *VcenterTokenserviceTokenExchangeInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *VcenterTokenserviceTokenExchangeInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *VcenterTokenserviceTokenExchangeInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetIssuedTokenType

`func (o *VcenterTokenserviceTokenExchangeInfo) GetIssuedTokenType() string`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *VcenterTokenserviceTokenExchangeInfo) GetIssuedTokenTypeOk() (*string, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *VcenterTokenserviceTokenExchangeInfo) SetIssuedTokenType(v string)`

SetIssuedTokenType sets IssuedTokenType field to given value.


### GetTokenType

`func (o *VcenterTokenserviceTokenExchangeInfo) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *VcenterTokenserviceTokenExchangeInfo) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *VcenterTokenserviceTokenExchangeInfo) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *VcenterTokenserviceTokenExchangeInfo) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *VcenterTokenserviceTokenExchangeInfo) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *VcenterTokenserviceTokenExchangeInfo) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *VcenterTokenserviceTokenExchangeInfo) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetScope

`func (o *VcenterTokenserviceTokenExchangeInfo) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *VcenterTokenserviceTokenExchangeInfo) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *VcenterTokenserviceTokenExchangeInfo) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *VcenterTokenserviceTokenExchangeInfo) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRefreshToken

`func (o *VcenterTokenserviceTokenExchangeInfo) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *VcenterTokenserviceTokenExchangeInfo) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *VcenterTokenserviceTokenExchangeInfo) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *VcenterTokenserviceTokenExchangeInfo) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


