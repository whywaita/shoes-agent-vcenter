# VcenterIdentityProvidersOidcInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DiscoveryEndpoint** | **string** | Endpoint to retrieve the provider metadata | 
**LogoutEndpoint** | Pointer to **string** | The endpoint to use for terminating the user&#39;s session at the identity provider. This value is automatically derived from the metadata information provided by the OIDC discovery endpoint. This field is optional because it was added in a newer version than its parent node. | [optional] 
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

### NewVcenterIdentityProvidersOidcInfo

`func NewVcenterIdentityProvidersOidcInfo(discoveryEndpoint string, authEndpoint string, tokenEndpoint string, publicKeyUri string, clientId string, clientSecret string, claimMap []VcenterIdentityProvidersOauth2CreateSpecClaimMap, issuer string, authenticationMethod VcenterIdentityProvidersOauth2AuthenticationMethod, authQueryParams []VcenterIdentityProvidersCreateSpecAuthQueryParams, ) *VcenterIdentityProvidersOidcInfo`

NewVcenterIdentityProvidersOidcInfo instantiates a new VcenterIdentityProvidersOidcInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVcenterIdentityProvidersOidcInfoWithDefaults

`func NewVcenterIdentityProvidersOidcInfoWithDefaults() *VcenterIdentityProvidersOidcInfo`

NewVcenterIdentityProvidersOidcInfoWithDefaults instantiates a new VcenterIdentityProvidersOidcInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) GetDiscoveryEndpoint() string`

GetDiscoveryEndpoint returns the DiscoveryEndpoint field if non-nil, zero value otherwise.

### GetDiscoveryEndpointOk

`func (o *VcenterIdentityProvidersOidcInfo) GetDiscoveryEndpointOk() (*string, bool)`

GetDiscoveryEndpointOk returns a tuple with the DiscoveryEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscoveryEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) SetDiscoveryEndpoint(v string)`

SetDiscoveryEndpoint sets DiscoveryEndpoint field to given value.


### GetLogoutEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) GetLogoutEndpoint() string`

GetLogoutEndpoint returns the LogoutEndpoint field if non-nil, zero value otherwise.

### GetLogoutEndpointOk

`func (o *VcenterIdentityProvidersOidcInfo) GetLogoutEndpointOk() (*string, bool)`

GetLogoutEndpointOk returns a tuple with the LogoutEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) SetLogoutEndpoint(v string)`

SetLogoutEndpoint sets LogoutEndpoint field to given value.

### HasLogoutEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) HasLogoutEndpoint() bool`

HasLogoutEndpoint returns a boolean if a field has been set.

### GetAuthEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) GetAuthEndpoint() string`

GetAuthEndpoint returns the AuthEndpoint field if non-nil, zero value otherwise.

### GetAuthEndpointOk

`func (o *VcenterIdentityProvidersOidcInfo) GetAuthEndpointOk() (*string, bool)`

GetAuthEndpointOk returns a tuple with the AuthEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) SetAuthEndpoint(v string)`

SetAuthEndpoint sets AuthEndpoint field to given value.


### GetTokenEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) GetTokenEndpoint() string`

GetTokenEndpoint returns the TokenEndpoint field if non-nil, zero value otherwise.

### GetTokenEndpointOk

`func (o *VcenterIdentityProvidersOidcInfo) GetTokenEndpointOk() (*string, bool)`

GetTokenEndpointOk returns a tuple with the TokenEndpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEndpoint

`func (o *VcenterIdentityProvidersOidcInfo) SetTokenEndpoint(v string)`

SetTokenEndpoint sets TokenEndpoint field to given value.


### GetPublicKeyUri

`func (o *VcenterIdentityProvidersOidcInfo) GetPublicKeyUri() string`

GetPublicKeyUri returns the PublicKeyUri field if non-nil, zero value otherwise.

### GetPublicKeyUriOk

`func (o *VcenterIdentityProvidersOidcInfo) GetPublicKeyUriOk() (*string, bool)`

GetPublicKeyUriOk returns a tuple with the PublicKeyUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKeyUri

`func (o *VcenterIdentityProvidersOidcInfo) SetPublicKeyUri(v string)`

SetPublicKeyUri sets PublicKeyUri field to given value.


### GetClientId

`func (o *VcenterIdentityProvidersOidcInfo) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *VcenterIdentityProvidersOidcInfo) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *VcenterIdentityProvidersOidcInfo) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetClientSecret

`func (o *VcenterIdentityProvidersOidcInfo) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *VcenterIdentityProvidersOidcInfo) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *VcenterIdentityProvidersOidcInfo) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.


### GetClaimMap

`func (o *VcenterIdentityProvidersOidcInfo) GetClaimMap() []VcenterIdentityProvidersOauth2CreateSpecClaimMap`

GetClaimMap returns the ClaimMap field if non-nil, zero value otherwise.

### GetClaimMapOk

`func (o *VcenterIdentityProvidersOidcInfo) GetClaimMapOk() (*[]VcenterIdentityProvidersOauth2CreateSpecClaimMap, bool)`

GetClaimMapOk returns a tuple with the ClaimMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimMap

`func (o *VcenterIdentityProvidersOidcInfo) SetClaimMap(v []VcenterIdentityProvidersOauth2CreateSpecClaimMap)`

SetClaimMap sets ClaimMap field to given value.


### GetIssuer

`func (o *VcenterIdentityProvidersOidcInfo) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *VcenterIdentityProvidersOidcInfo) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *VcenterIdentityProvidersOidcInfo) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetAuthenticationMethod

`func (o *VcenterIdentityProvidersOidcInfo) GetAuthenticationMethod() VcenterIdentityProvidersOauth2AuthenticationMethod`

GetAuthenticationMethod returns the AuthenticationMethod field if non-nil, zero value otherwise.

### GetAuthenticationMethodOk

`func (o *VcenterIdentityProvidersOidcInfo) GetAuthenticationMethodOk() (*VcenterIdentityProvidersOauth2AuthenticationMethod, bool)`

GetAuthenticationMethodOk returns a tuple with the AuthenticationMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationMethod

`func (o *VcenterIdentityProvidersOidcInfo) SetAuthenticationMethod(v VcenterIdentityProvidersOauth2AuthenticationMethod)`

SetAuthenticationMethod sets AuthenticationMethod field to given value.


### GetAuthQueryParams

`func (o *VcenterIdentityProvidersOidcInfo) GetAuthQueryParams() []VcenterIdentityProvidersCreateSpecAuthQueryParams`

GetAuthQueryParams returns the AuthQueryParams field if non-nil, zero value otherwise.

### GetAuthQueryParamsOk

`func (o *VcenterIdentityProvidersOidcInfo) GetAuthQueryParamsOk() (*[]VcenterIdentityProvidersCreateSpecAuthQueryParams, bool)`

GetAuthQueryParamsOk returns a tuple with the AuthQueryParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthQueryParams

`func (o *VcenterIdentityProvidersOidcInfo) SetAuthQueryParams(v []VcenterIdentityProvidersCreateSpecAuthQueryParams)`

SetAuthQueryParams sets AuthQueryParams field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


