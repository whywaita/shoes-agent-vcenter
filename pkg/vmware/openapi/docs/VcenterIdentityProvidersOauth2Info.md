# VcenterIdentityProvidersOauth2Info

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthEndpoint** | **string** | Authentication/authorization endpoint of the provider | 
**TokenEndpoint** | **string** | Token endpoint of the provider | 
**PublicKeyUri** | **string** | Endpoint to retrieve the provider public key for validation | 
**ClientId** | **string** | Client identifier to connect to the provider | 
**ClientSecret** | **string** | The secret shared between the client and the provider | 
**ClaimMap** | [**[]VcenterIdentityProvidersOauth2CreateSpecClaimMap**](VcenterIdentityProvidersOauth2CreateSpecClaimMap.md) | The map used to transform an OAuth2 claim to a corresponding claim that vCenter Server understands. Currently only the key \&quot;perms\&quot; is supported. The key \&quot;perms\&quot; is used for mapping the \&quot;perms\&quot; claim of incoming JWT. The value is another map with an external group as the key and a vCenter Server group as value. | 
**Issuer** | **string** | The identity provider namespace. It is used to validate the issuer in the acquired OAuth2 token | 
**AuthenticationMethod** | [**VcenterIdentityProvidersOauth2AuthenticationMethod**](VcenterIdentityProvidersOauth2AuthenticationMethod.md) |  | 
**AuthQueryParams** | [**[]VcenterIdentityProvidersCreateSpecAuthQueryParams**](VcenterIdentityProvidersCreateSpecAuthQueryParams.md) |  key/value pairs that are to be appended to the authEndpoint request.   How to append to authEndpoint request:  If the map is not empty, a \&quot;?\&quot; is added to the endpoint URL, and combination of each k and each string in the v is added with an \&quot;&amp;\&quot; delimiter. Details:    - If the value contains only one string, then the key is added with \&quot;k&#x3D;v\&quot;.    - If the value is an empty list, then the key is added without a \&quot;&#x3D;v\&quot;.    - If the value contains multiple strings, then the key is repeated in the query-string for each string in the value.  | 

## Methods

### NewVcenterIdentityProvidersOauth2Info

`func NewVcenterIdentityProvidersOauth2Info(authEndpoint string, tokenEndpoint string, publicKeyUri string, clientId string, clientSecret string, claimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap, issuer string, authenticationMethod VcenterIdentityProvidersOauth2AuthenticationMethod, authQueryParams []VcenterIdentityProvidersCreateSpecAuthQueryParams, ) *VcenterIdentityProvidersOauth2Info`

NewVcenterIdentityProvidersOauth2Info instantiates a new VcenterIdentityProvidersOauth2Info object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOauth2InfoWithDefaults

`func NewVcenterIdentityProvidersOauth2InfoWithDefaults() *VcenterIdentityProvidersOauth2Info`

NewVcenterIdentityProvidersOauth2InfoWithDefaults instantiates a new VcenterIdentityProvidersOauth2Info object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2Info) GetAuthEndpoint() string`

GetAuthEndpoint returns the AuthEndpoint field if non-nil, zero value otherwise.

### GetAuthEndpointOk

`func (o *VcenterIdentityProvidersOauth2Info) GetAuthEndpointOk() (*string, bool)`

GetAuthEndpointOk returns a tuple with the AuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEndpoint

`func (o *VcenterIdentityProvidersOauth2Info) SetAuthEndpoint(v string)`

SetAuthEndpoint sets AuthEndpoint field to given value.


### GetTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2Info) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *VcenterIdentityProvidersOauth2Info) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *VcenterIdentityProvidersOauth2Info) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetPublicKeyUri

`func (o *VcenterIdentityProvidersOauth2Info) GetPublicKeyUri() string`

GetPublicKeyUri returns the PublicKeyUri field if non-nil, zero value otherwise.

### GetPublicKeyUriOk

`func (o *VcenterIdentityProvidersOauth2Info) GetPublicKeyUriOk() (*string, bool)`

GetPublicKeyUriOk returns a tuple with the PublicKeyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyUri

`func (o *VcenterIdentityProvidersOauth2Info) SetPublicKeyUri(v string)`

SetPublicKeyUri sets PublicKeyUri field to given value.


### GetClientId

`func (o *VcenterIdentityProvidersOauth2Info) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOauth2Info) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOauth2Info) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *VcenterIdentityProvidersOauth2Info) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VcenterIdentityProvidersOauth2Info) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VcenterIdentityProvidersOauth2Info) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetClaimMap

`func (o *VcenterIdentityProvidersOauth2Info) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap`

GetClaimMap returns the ClaimMap field if non-nil, zero value otherwise.

### GetClaimMapOk

`func (o *VcenterIdentityProvidersOauth2Info) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool)`

GetClaimMapOk returns a tuple with the ClaimMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMap

`func (o *VcenterIdentityProvidersOauth2Info) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap)`

SetClaimMap sets ClaimMap field to given value.


### GetIssuer

`func (o *VcenterIdentityProvidersOauth2Info) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VcenterIdentityProvidersOauth2Info) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VcenterIdentityProvidersOauth2Info) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetAuthenticationMethod

`func (o *VcenterIdentityProvidersOauth2Info) GetAuthenticationMethod() VcenterIdentityProvidersOauth2AuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VcenterIdentityProvidersOauth2Info) GetAuthenticationMethodOk() (*VcenterIdentityProvidersOauth2AuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VcenterIdentityProvidersOauth2Info) SetAuthenticationMethod(v VcenterIdentityProvidersOauth2AuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2Info) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersOauth2Info) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersOauth2Info) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


