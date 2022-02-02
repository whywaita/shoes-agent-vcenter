# Oauth2TokenInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessToken** | **string** | The access token issued by the authorization server. | 
**TokenType** | **string** | A case-insensitive value specifying the method of using the access token issued. | 
**ExpiresIn** | Pointer to **int64** | The validity lifetime, in seconds, of the token issued by the server. unset if not applicable for issued token. | [optional] 
**Scope** | Pointer to **string** | Scope of the issued access token. The value of the scope parameter is expressed as a list of space- delimited, case-sensitive strings. The strings are defined by the authorization server. If the value contains multiple space-delimited strings, their order does not matter, and each string adds an additional access range to the requested scope. unset if the scope of the issued security token is identical to the scope requested by the client. | [optional] 
**RefreshToken** | Pointer to **string** | The refresh token, which can be used to obtain new access tokens. unset if not applicable to the specific request. | [optional] 
**IssuedTokenType** | Pointer to **string** | An identifier which indicates the type of the access token in the TokenInfo.access-token field. unset if not the result of a token-exchange invocation; otherwise, required. | [optional] 

## Methods

### NewOauth2TokenInfo

`func NewOauth2TokenInfo(accessToken string, tokenType string, ) *Oauth2TokenInfo`

NewOauth2TokenInfo instantiates a new Oauth2TokenInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOauth2TokenInfoWithDefaults

`func NewOauth2TokenInfoWithDefaults() *Oauth2TokenInfo`

NewOauth2TokenInfoWithDefaults instantiates a new Oauth2TokenInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessToken

`func (o *Oauth2TokenInfo) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *Oauth2TokenInfo) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *Oauth2TokenInfo) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetTokenType

`func (o *Oauth2TokenInfo) GetTokenType() string`

GetTokenType returns the TokenType field if non-nil, zero value otherwise.

### GetTokenTypeOk

`func (o *Oauth2TokenInfo) GetTokenTypeOk() (*string, bool)`

GetTokenTypeOk returns a tuple with the TokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenType

`func (o *Oauth2TokenInfo) SetTokenType(v string)`

SetTokenType sets TokenType field to given value.


### GetExpiresIn

`func (o *Oauth2TokenInfo) GetExpiresIn() int64`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *Oauth2TokenInfo) GetExpiresInOk() (*int64, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *Oauth2TokenInfo) SetExpiresIn(v int64)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *Oauth2TokenInfo) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetScope

`func (o *Oauth2TokenInfo) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Oauth2TokenInfo) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Oauth2TokenInfo) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Oauth2TokenInfo) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetRefreshToken

`func (o *Oauth2TokenInfo) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *Oauth2TokenInfo) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *Oauth2TokenInfo) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *Oauth2TokenInfo) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetIssuedTokenType

`func (o *Oauth2TokenInfo) GetIssuedTokenType() string`

GetIssuedTokenType returns the IssuedTokenType field if non-nil, zero value otherwise.

### GetIssuedTokenTypeOk

`func (o *Oauth2TokenInfo) GetIssuedTokenTypeOk() (*string, bool)`

GetIssuedTokenTypeOk returns a tuple with the IssuedTokenType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTokenType

`func (o *Oauth2TokenInfo) SetIssuedTokenType(v string)`

SetIssuedTokenType sets IssuedTokenType field to given value.

### HasIssuedTokenType

`func (o *Oauth2TokenInfo) HasIssuedTokenType() bool`

HasIssuedTokenType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


